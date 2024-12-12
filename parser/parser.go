package parser

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/token"
)

func Parse(l lexer.Lexer) (ast.Script, error) {
	var p = parser{l: l}

	// So that all of curr and next, next2 and next3 tokens get initialized.
	p.proceed()
	p.proceed()
	p.proceed()
	p.proceed()

	script := p.ParseScript()
	if p.Error != nil {
		return nil, p.Error
	}
	return script, nil
}

type parser struct {
	l     lexer.Lexer
	curr  token.Token
	next  token.Token
	next2 token.Token
	next3 token.Token
	Error *ParserError
}

type ParserError struct {
	Line     int
	Position int
	Message  string
}

func (err *ParserError) Error() string {
	return fmt.Sprintf("syntax error: %s. (line: %d, column: %d)", err.Message, err.Line, err.Position)
}

func (p *parser) error(msg string, args ...any) {
	if p.Error == nil {
		p.Error = &ParserError{
			Line:     p.curr.Line,
			Position: p.curr.Position,
			Message:  fmt.Sprintf(msg, args...),
		}
	}
}

func (p *parser) proceed() {
	p.curr = p.next
	p.next = p.next2
	p.next2 = p.next3
	p.next3 = p.l.NextToken()
}

func (p *parser) ParseScript() ast.Script {
	var script ast.Script

	for ; p.curr.Type != token.EOF; p.proceed() {
		switch p.curr.Type {
		case token.BLANK, token.NEWLINE:
			continue
		case token.HASH:
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		default:
			if cmdList := p.parseCommandList(); cmdList != nil {
				script = append(script, cmdList)
			} else {
				return script
			}
			if p.curr.Type == token.LEFT_PAREN || p.curr.Type == token.RIGHT_PAREN {
				p.error("token `%s` cannot be placed here", p.curr)
				return nil
			}
		}
	}

	return script
}

func (p *parser) parseCommandList() ast.Statement {
	var left ast.Statement
	pipe := p.parsePipline()

	if pipe == nil {
		return nil
	} else if len(pipe) == 1 {
		left = pipe[0].Command
	} else {
		left = pipe
	}

	for p.curr.Type == token.AND || p.curr.Type == token.OR {
		operator := p.curr.Literal
		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		var right ast.Statement
		rightPipe := p.parsePipline()
		if rightPipe == nil {
			return nil
		} else if len(rightPipe) == 1 {
			right = rightPipe[0].Command
		} else {
			right = rightPipe
		}

		left = ast.List{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	if p.curr.Type == token.AMPERSAND {
		return ast.BackgroundConstruction{Statement: left}
	}

	return left
}

func (p *parser) parsePipline() ast.Pipeline {
	var pipeline ast.Pipeline

	cmd := p.parseCommand()
	if cmd == nil {
		return nil
	}
	pipeline = append(pipeline, ast.PipelineCommand{Command: cmd})

	for i := 0; ; i++ {
		if p.curr.Type != token.PIPE && p.curr.Type != token.PIPE_AMPERSAND {
			break
		}
		var pipe ast.PipelineCommand
		pipeline[i].Stderr = p.curr.Type == token.PIPE_AMPERSAND

		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		pipe.Command = p.parseCommand()
		if pipe.Command == nil {
			return nil
		}
		pipeline = append(pipeline, pipe)
	}

	return pipeline
}

func (p *parser) parseCommand() ast.Statement {
	if p.curr.Type == token.FUNCTION {
		return p.parseFunction()
	}

	if compound := p.getCompoundParser(); compound != nil {
		return compound()
	}

	env := p.parseAssignement()
	if env != nil && (p.isControlToken() || p.curr.Type == token.EOF) {
		return env
	}

	var cmd ast.Command
	cmd.Name = p.parseExpression()
	if cmd.Name == nil {
		p.error("expected a valid command name, found `%s`", p.curr)
		return nil
	}
	cmd.Env = env

	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	if p.curr.Type == token.LEFT_PAREN {
		return p.parseNakedFunction(cmd.Name)
	}

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			break
		case p.curr.Type == token.EOF:
			break loop
		case p.curr.Type == token.HASH:
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		case p.isRedirectionToken():
			p.HandleRedirection(&cmd.Redirections)
		default:
			if p.isControlToken() {
				break loop
			}

			cmd.Args = append(cmd.Args, p.parseExpression())
		}

		if !p.isRedirectionToken() && !p.isControlToken() {
			p.proceed()
		}
	}
	return cmd
}

func (p *parser) parseExpression() ast.Expression {
	var exprs []ast.Expression

loop:
	for {
		switch p.curr.Type {
		case token.BLANK, token.EOF:
			break loop
		case token.SIMPLE_EXPANSION:
			exprs = append(exprs, ast.Var(p.curr.Literal))
		case token.SINGLE_QUOTE:
			exprs = append(exprs, p.parseLiteralString())
		case token.DOUBLE_QUOTE:
			exprs = append(exprs, p.parseString())
		case token.DOLLAR_PAREN:
			exprs = append(exprs, p.parseCommandSubstitution())
		case token.DOLLAR_DOUBLE_PAREN:
			exprs = append(exprs, p.parseArithmeticSubstitution())
		case token.GT_PAREN, token.LT_PAREN:
			exprs = append(exprs, p.parseProcessSubstitution())
		case token.DOLLAR_BRACE:
			exprs = append(exprs, p.parseParameterExpansion())
		case token.INT:
			if len(exprs) == 0 && p.isRedirectionToken() {
				break loop
			}
			fallthrough
		default:
			if p.curr.Type != token.INT && p.isRedirectionToken() || p.isControlToken() {
				break loop
			}

			exprs = append(exprs, ast.Word(p.curr.Literal))
		}

		p.proceed()
	}

	return concat(exprs, false)
}

func (p *parser) parseLiteralString() ast.Word {
	p.proceed()

	if p.curr.Type == token.SINGLE_QUOTE {
		return ast.Word("")
	}

	word := p.curr.Literal
	p.proceed()

	if p.curr.Type != token.SINGLE_QUOTE {
		p.error("a closing single quote is missing")
	}

	return ast.Word(word)
}

func (p *parser) parseString() ast.Expression {
	p.proceed()

	if p.curr.Type == token.DOUBLE_QUOTE {
		return ast.Word("")
	}

	var exprs []ast.Expression

loop:
	for {
		switch p.curr.Type {
		case token.DOUBLE_QUOTE, token.EOF:
			break loop
		case token.ESCAPED_CHAR:
			exprs = append(exprs, ast.Word("\\"+p.curr.Literal))
		case token.SIMPLE_EXPANSION:
			exprs = append(exprs, ast.Var(p.curr.Literal))
		case token.DOLLAR_BRACE:
			exprs = append(exprs, p.parseParameterExpansion())
		case token.DOLLAR_PAREN:
			exprs = append(exprs, p.parseCommandSubstitution())
		case token.DOLLAR_DOUBLE_PAREN:
			exprs = append(exprs, p.parseArithmeticSubstitution())
		default:
			exprs = append(exprs, ast.Word(p.curr.Literal))
		}

		p.proceed()
	}

	if p.curr.Type != token.DOUBLE_QUOTE {
		p.error("a closing double quote is missing")
		return nil
	}

	return concat(exprs, true)
}

func (p *parser) parseFunction() ast.Statement {
	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	nameExpr := p.parseExpression()
	if nameExpr == nil {
		p.error("function name is required")
		return nil
	}

	name, ok := nameExpr.(ast.Word)
	if !ok {
		p.error("invalid function name was supplied")
		return nil
	}
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.LEFT_PAREN {
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.RIGHT_PAREN {
			p.error("expected `)`, found `%s`", p.curr)
			return nil
		}
		p.proceed()
	}

	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	compound := p.getCompoundParser()
	if compound == nil {
		p.error("function body is expected to be a compound command, found `%s`", p.curr)
		return nil
	}

	return ast.Function{Name: string(name), Command: compound()}
}

func (p *parser) parseNakedFunction(nameExpr ast.Expression) ast.Statement {
	name, ok := nameExpr.(ast.Word)
	if !ok {
		p.error("invalid function name was supplied")
		return nil
	}

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	if p.curr.Type != token.RIGHT_PAREN {
		p.error("expected `)`, found `%s`", p.curr)
		return nil
	}
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	compound := p.getCompoundParser()
	if compound == nil {
		p.error("function body is expected to be a compound command, found `%s`", p.curr)
		return nil
	}

	return ast.Function{Name: string(name), Command: compound()}
}

func concat(n []ast.Expression, quoted bool) ast.Expression {
	var conc ast.UnquotedString
	var mergedWords ast.Word
	var hasWords bool

	for i, node := range n {

		if w, ok := node.(ast.Word); ok {
			mergedWords += w
			hasWords = true
		} else {
			if hasWords {
				conc = append(conc, mergedWords)
				mergedWords, hasWords = "", false
			}
			conc = append(conc, node)
		}

		if i == len(n)-1 && hasWords {
			conc = append(conc, mergedWords)
		}
	}

	if len(conc) == 0 {
		return nil
	}

	if quoted {
		if len(conc) == 1 {
			if w, ok := conc[0].(ast.Word); ok {
				return w
			}
		}
		return ast.QuotedString(conc)
	}

	if len(conc) == 1 {
		return conc[0]
	}

	return conc
}

func (p *parser) isControlToken() bool {
	return p.curr.Type == token.PIPE ||
		p.curr.Type == token.PIPE_AMPERSAND ||
		p.curr.Type == token.AND ||
		p.curr.Type == token.OR ||
		p.curr.Type == token.AMPERSAND ||
		p.curr.Type == token.SEMICOLON ||
		p.curr.Type == token.NEWLINE ||
		p.curr.Type == token.LEFT_PAREN ||
		p.curr.Type == token.RIGHT_PAREN
}

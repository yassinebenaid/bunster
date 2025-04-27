package parser

import (
	"strings"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/token"
)

func (p *parser) getBuiltinParser() func() ast.Statement {
	switch p.curr.Type {
	case token.BREAK:
		return p.parseBreak
	case token.EXIT:
		return p.parseExit
	case token.RETURN:
		return p.parseReturn
	case token.CONTINUE:
		return p.parseContinue
	case token.FUNCTION:
		return p.parseFunction
	case token.DEFER:
		return p.parseDefer
	case token.WAIT:
		return p.parseWait
	case token.LOCAL:
		return p.parseLocal
	case token.EXPORT:
		return p.parseExport
	case token.UNSET:
		return p.parseUnset
	case token.AT:
		if p.next.Type != token.EMBED {
			return nil
		}
		return p.parseEmbedDirective
	case token.THEN, token.ELIF, token.ELSE, token.FI, token.DO, token.DONE, token.ESAC:
		p.error("`%s` is a reserved keyword, cannot be used a command name", p.curr)
	}
	return nil
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

	var function ast.Function

	if p.curr.Type == token.LEFT_PAREN {
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		function.Flags = p.parseFunctionFlags()
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
		p.error("function body is expected, found `%s`", p.curr)
		return nil
	}
	body := compound()
	function.Name = string(name)

	switch v := body.(type) {
	case ast.Group:
		function.Body, function.Redirections = v.Body, v.Redirections
	case ast.SubShell:
		function.Body, function.Redirections, function.SubShell = v.Body, v.Redirections, true
	default:
		p.error("function body is expected to be a group or subshell")
		return nil
	}

	switch p.curr.Type {
	case token.SEMICOLON, token.NEWLINE, token.EOF, token.AND, token.OR:
	default:
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return &function
}

func (p *parser) parseFunctionFlags() []ast.Flag {
	var flags []ast.Flag

	for p.curr.Type != token.RIGHT_PAREN && p.curr.Type != token.EOF {
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		if p.curr.Type != token.MINUS && p.curr.Type != token.DECREMENT {
			p.error("unexpected token `%s`", p.curr)
		}
		var flag = ast.Flag{
			Long: p.curr.Type == token.DECREMENT,
		}
		p.proceed()

		if p.curr.Type != token.WORD {
			p.error("expected a valid flag name, found `%s`", p.curr)
		}

		if !flag.Long && len(p.curr.Literal) != 1 {
			p.error("short flags can only be one character long, found `%s`", p.curr)
		}
		flag.Name = p.curr.Literal
		p.proceed()

		if p.curr.Type == token.ASSIGN {
			flag.AcceptsValue = true
			p.proceed()
		} else if p.curr.Type == token.LEFT_BRACKET {
			flag.AcceptsValue = true
			flag.Optional = true

			if !(p.curr.Type == token.LEFT_BRACKET && p.next.Type == token.ASSIGN && p.next2.Type == token.RIGHT_BRACKET) {
				p.error("expected [=] to indicate optional value, found `%s%s%s`", p.curr, p.next, p.next2)
			}
			p.proceed()
			p.proceed()
			p.proceed()
		}

		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
		flags = append(flags, flag)
	}

	return flags
}

func (p *parser) parseDefer() ast.Statement {
	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	command := p.parseCommand()

	switch command.(type) {
	case ast.Command, ast.Group, ast.SubShell:
	default:
		p.error("expected a simple command, group or subshell after `defer`")
	}

	fn := ast.Defer{Command: command}

	return fn
}

func (p *parser) parseBreak() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return &ast.Break{}
}

func (p *parser) parseExit() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	var code ast.Expression = ast.Word("0")
	if exp := p.parseExpression(); exp != nil {
		code = exp
	}

	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return ast.Exit{
		Code: code,
	}
}

func (p *parser) parseReturn() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	var code ast.Expression = ast.Word("0")
	if exp := p.parseExpression(); exp != nil {
		code = exp
	}

	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return &ast.Return{
		Code: code,
	}
}

func (p *parser) parseContinue() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}
	return &ast.Continue{}
}

func (p *parser) parseWait() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}
	return ast.Wait{}
}

func (p *parser) parseLocal() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var assignements ast.LocalParameterAssignement

	for {
		if p.curr.Type != token.WORD {
			break
		}
		assignment := ast.Assignement{Name: p.curr.Literal}
		p.proceed()

		if p.curr.Type == token.ASSIGN {
			p.proceed()
			assignment.Value = p.parseExpression()
		}

		if p.curr.Type == token.BLANK {
			p.proceed()
		}

		assignements = append(assignements, assignment)
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if assignements == nil || (!p.isControlToken() && p.curr.Type != token.EOF) {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return assignements
}

func (p *parser) parseExport() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var assignements ast.ExportParameterAssignement

	for {
		if p.curr.Type != token.WORD {
			break
		}
		assignment := ast.Assignement{Name: p.curr.Literal}
		p.proceed()

		if p.curr.Type == token.ASSIGN {
			p.proceed()
			assignment.Value = p.parseExpression()
		}

		if p.curr.Type == token.BLANK {
			p.proceed()
		}

		assignements = append(assignements, assignment)
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if assignements == nil || (!p.isControlToken() && p.curr.Type != token.EOF) {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return assignements
}

func (p *parser) parseUnset() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var flag string

	if p.curr.Type == token.MINUS {
		p.proceed()
		if p.curr.Type == token.WORD && (p.curr.Literal == "f" || p.curr.Literal == "v") {
			flag = "-" + p.curr.Literal
			p.proceed()
		} else {
			p.error("expected a valid flag character after `-`, found `%v`", p.curr)
		}
	}

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var names []ast.Expression

	for !p.isControlToken() && p.curr.Type != token.EOF {
		name := p.parseExpression()
		names = append(names, name)

		if p.curr.Type == token.BLANK {
			p.proceed()
		}

		if p.curr.Type == token.HASH {
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		}
	}

	if names == nil || (!p.isControlToken() && p.curr.Type != token.EOF) {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return ast.Unset{
		Flag:  flag,
		Names: names,
	}
}

func (p *parser) parseEmbedDirective() ast.Statement {
	p.proceed()
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	} else {
		p.error("expected a blank after the @embed directive, found %s", p.curr)
	}

	var embed ast.Embed
	var expr ast.Expression

loop:
	for {
		switch p.curr.Type {
		case token.EOF, token.NEWLINE, token.SEMICOLON:
			break loop
		default:
			expr = p.parseExpression()

			switch v := expr.(type) {
			case ast.Word:
				if v == "" {
					p.error("expected a file path, found empty string")
				} else if strings.ContainsAny(string(v), "*\"'<>?|`\\:") {
					p.error("expected a valid file path, found %q", v)
				} else if strings.HasPrefix(string(v), "/") || strings.HasSuffix(string(v), "/") {
					p.error("path cannot start or end with slash, %q", v)
				}

				embed = append(embed, string(v))
			default:
				p.error("expected a valid file path")
				return nil
			}
		}
		if p.curr.Type != token.NEWLINE && p.curr.Type != token.SEMICOLON {
			p.proceed()
		}
	}

	if embed == nil || (p.curr.Type != token.EOF && p.curr.Type != token.NEWLINE && p.curr.Type != token.SEMICOLON) {
		p.error("unexpected token: %v", p.curr)
	}

	return embed
}

package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) getCompoundParser() func() ast.Statement {
	switch p.curr.Type {
	case token.WHILE, token.UNTIL:
		return p.parseWhileLoop
	case token.FOR:
		return p.parseForLoop
	case token.IF:
		return p.parseIf
	case token.CASE:
		return p.parseCase
	case token.LEFT_BRACE:
		return p.parseGroup
	case token.LEFT_PAREN:
		return p.parseSubShell
	case token.DOUBLE_LEFT_PAREN:
		return p.parseArithmeticCommand
	case token.DOUBLE_LEFT_BRACKET:
		return p.parseTestCommand
	case token.THEN, token.ELIF, token.ELSE, token.FI, token.DO, token.DONE, token.ESAC:
		p.error("`%s` is a reserved keyword, cannot be used a command name", p.curr)
		fallthrough
	default:
		return nil
	}
}

func (p *Parser) parseWhileLoop() ast.Statement {
	var loop ast.Loop
	loopKeyword := p.curr.Literal
	loop.Negate = p.curr.Type == token.UNTIL
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.DO && p.curr.Type != token.DONE && p.curr.Type != token.EOF {
		cmdList := p.parseCommandList()
		if cmdList == nil {
			return nil
		}
		loop.Head = append(loop.Head, cmdList)
		if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
			p.proceed()
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if loop.Head == nil {
		p.error("expected command list after `%s`", loopKeyword)
	} else if p.curr.Type != token.DO {
		p.error("expected `do`, found `%s`", p.curr)
	}

	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.DONE && p.curr.Type != token.EOF {
		cmdList := p.parseCommandList()
		if cmdList == nil {
			return nil
		}
		loop.Body = append(loop.Body, cmdList)
		if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
			p.proceed()
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if loop.Body == nil {
		p.error("expected command list after `do`")
	} else if p.curr.Type != token.DONE {
		p.error("expected `done` to close `%s` loop", loopKeyword)
	}

	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&loop.Redirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	return loop
}

func (p *Parser) parseForLoop() ast.Statement {
	var loopHead ast.ForHead
	var loopVar string
	var loopOperands []ast.Expression
	var loopBody []ast.Statement
	var loopRedirections []ast.Redirection

	p.proceed()
	for p.curr.Type == token.BLANK {
		p.proceed()
	}
	if p.curr.Type == token.DOUBLE_LEFT_PAREN {
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.SEMICOLON {
			loopHead.Init = p.parseArithmetics()
		}
		if p.curr.Type != token.SEMICOLON {
			p.error("expected a semicolon `;`, found `%s`", p.curr)
		}
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.SEMICOLON {
			loopHead.Test = p.parseArithmetics()
		}
		if p.curr.Type != token.SEMICOLON {
			p.error("expected a semicolon `;`, found `%s`", p.curr)
		}
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.RIGHT_PAREN {
			loopHead.Update = p.parseArithmetics()
		}

		if !(p.curr.Type == token.RIGHT_PAREN && p.next.Type == token.RIGHT_PAREN) {
			p.error("expected `))` to close loop head, found `%s`", p.curr)
		}
		p.proceed()
		p.proceed()
	} else {

		if p.curr.Type != token.WORD {
			p.error("expected identifier after `for`")
		}
		loopVar = p.curr.Literal
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}

		if p.curr.Type == token.IN {
			p.proceed()
			if p.curr.Type == token.BLANK {
				p.proceed()
			}
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.SEMICOLON && p.curr.Type != token.EOF {
				member := p.parseExpression()
				if member == nil {
					p.error("unexpected token `%s`", p.curr)
					break
				}
				loopOperands = append(loopOperands, member)
				if p.curr.Type == token.BLANK {
					p.proceed()
				}
			}
			if loopOperands == nil {
				p.error("missing operand after `in`")
			}
		}
	}

	if p.curr.Type == token.SEMICOLON {
		p.proceed()
	}
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	if p.curr.Type != token.DO {
		p.error("expected `do`, found `%s`", p.curr)
	}
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.DONE && p.curr.Type != token.EOF {
		cmdList := p.parseCommandList()
		if cmdList == nil {
			return nil
		}
		loopBody = append(loopBody, cmdList)
		if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
			p.proceed()
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if loopBody == nil {
		p.error("expected command list after `do`")
	} else if p.curr.Type != token.DONE {
		p.error("expected `done` to close `for` loop")
	}

	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&loopRedirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	if loopVar == "" {
		return ast.For{
			Head:         loopHead,
			Body:         loopBody,
			Redirections: loopRedirections,
		}
	}

	return ast.RangeLoop{
		Var:          loopVar,
		Operands:     loopOperands,
		Body:         loopBody,
		Redirections: loopRedirections,
	}
}

func (p *Parser) parseIf() ast.Statement {
	var cond ast.If
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.THEN && p.curr.Type != token.FI && p.curr.Type != token.EOF {
		cmdList := p.parseCommandList()
		if cmdList == nil {
			return nil
		}
		cond.Head = append(cond.Head, cmdList)
		if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
			p.proceed()
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if cond.Head == nil {
		p.error("expected command list after `if`")
	} else if p.curr.Type != token.THEN {
		p.error("expected `then`, found `%s`", p.curr)
	}

	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.FI && p.curr.Type != token.ELIF && p.curr.Type != token.ELSE && p.curr.Type != token.EOF {
		cmdList := p.parseCommandList()
		if cmdList == nil {
			return nil
		}
		cond.Body = append(cond.Body, cmdList)
		if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
			p.proceed()
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if cond.Body == nil {
		p.error("expected command list after `then`")
	}

	for p.curr.Type == token.ELIF {
		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		var elif ast.Elif

		for p.curr.Type != token.THEN && p.curr.Type != token.FI && p.curr.Type != token.EOF {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			elif.Head = append(elif.Head, cmdList)
			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
			for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
				p.proceed()
			}
		}

		if elif.Head == nil {
			p.error("expected command list after `elif`")
		} else if p.curr.Type != token.THEN {
			p.error("expected `then`, found `%s`", p.curr)
		}

		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		for p.curr.Type != token.FI && p.curr.Type != token.ELIF && p.curr.Type != token.ELSE && p.curr.Type != token.EOF {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			elif.Body = append(elif.Body, cmdList)
			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
			for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
				p.proceed()
			}
		}

		if elif.Body == nil {
			p.error("expected command list after `then`")
		}

		cond.Elifs = append(cond.Elifs, elif)
	}

	if p.curr.Type == token.ELSE {
		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
		for p.curr.Type != token.FI && p.curr.Type != token.EOF {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			cond.Alternate = append(cond.Alternate, cmdList)
			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
			for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
				p.proceed()
			}
		}
		if cond.Alternate == nil {
			p.error("expected command list after `else`")
		}
	}

	if p.curr.Type != token.FI {
		p.error("expected `fi` to close `if` command")
	}

	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&cond.Redirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	return cond
}

func (p *Parser) parseCase() ast.Statement {
	var stmt ast.Case
	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	stmt.Word = p.parseExpression()
	if stmt.Word == nil {
		p.error("incomplete `case` statement, an operand is required after `case`")
	}

	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	if p.curr.Type != token.IN {
		p.error("expected `in`, found `%s`", p.curr)
	}
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.ESAC && p.curr.Type != token.EOF {
		var item ast.CaseItem
		if p.curr.Type == token.LEFT_PAREN {
			p.proceed()
		}

		for {
			pattern := p.parseExpression()
			if pattern == nil {
				p.error("invalid pattern provided, unexpected token `%s`", p.curr)
			}
			item.Patterns = append(item.Patterns, pattern)

			if p.curr.Type == token.BLANK {
				p.proceed()
			}
			if p.curr.Type != token.PIPE {
				break
			}
			p.proceed()
			if p.curr.Type == token.BLANK {
				p.proceed()
			}
		}

		if p.curr.Type != token.RIGHT_PAREN {
			p.error("expected `)`, found `%s`", p.curr)
		}
		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		for {
			if p.curr.Type == token.ESAC || p.curr.Type == token.EOF {
				break
			}
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			item.Body = append(item.Body, cmdList)

			if (p.curr.Type == token.SEMICOLON &&
				p.next.Type != token.SEMICOLON &&
				p.next.Type != token.AMPERSAND) ||
				p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
			for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
				p.proceed()
			}

			if p.curr.Type == token.SEMICOLON && p.next.Type == token.AMPERSAND {
				p.proceed()
				p.proceed()

				item.Terminator = ";&"
				for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
					p.proceed()
				}

				break
			}

			if p.curr.Type == token.SEMICOLON && p.next.Type == token.SEMICOLON {
				p.proceed()
				p.proceed()

				item.Terminator = ";;"
				if p.curr.Type == token.AMPERSAND {
					item.Terminator = ";;&"
					p.proceed()
				}
				for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
					p.proceed()
				}

				break
			}
		}
		stmt.Cases = append(stmt.Cases, item)
	}

	if p.curr.Type != token.ESAC {
		p.error("expected `esac` to close `case` command")
	}
	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&stmt.Redirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	return stmt
}

func (p *Parser) parseGroup() ast.Statement {
	var group ast.Group
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.RIGHT_BRACE && p.curr.Type != token.EOF {
		if p.curr.Type == token.HASH {
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		} else {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			group.Body = append(group.Body, cmdList)

			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
		}

		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if len(group.Body) == 0 {
		p.error("expeceted a command list after `{`")
	}

	if p.curr.Type != token.RIGHT_BRACE {
		p.error("expected `}`, found `%s`", p.curr)
	}

	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&group.Redirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	return group
}

func (p *Parser) parseSubShell() ast.Statement {
	var shell ast.SubShell
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.RIGHT_PAREN && p.curr.Type != token.EOF {
		if p.curr.Type == token.HASH {
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		} else {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			shell.Body = append(shell.Body, cmdList)
			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if len(shell.Body) == 0 {
		p.error("expeceted a command list after `(`")
	}

	if p.curr.Type != token.RIGHT_PAREN {
		p.error("expected `)`, found `%s`", p.curr)
	}

	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&shell.Redirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	return shell
}

func (p *Parser) parseArithmeticCommand() ast.Statement {
	p.proceed()
	for p.curr.Type == token.BLANK {
		p.proceed()
	}

	var arth ast.ArithmeticCommand
	arth.Arithmetic = p.parseArithmetics()

	if !(p.curr.Type == token.RIGHT_PAREN && p.next.Type == token.RIGHT_PAREN) {
		p.error("expected `))` to close arithmetic expression, found `%s`", p.curr)
	}
	p.proceed()
	p.proceed()

loop:
	for {
		switch {
		case p.curr.Type == token.BLANK:
			p.proceed()
		case p.isRedirectionToken():
			p.HandleRedirection(&arth.Redirections)
		default:
			break loop
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
	}

	return arth
}

package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) getCompoundParser() func() ast.Node {
	switch p.curr.Type {
	case token.WHILE, token.UNTIL:
		return p.parseWhileLoop
	case token.FOR:
		return p.parseForLoop
	case token.IF:
		return p.parseIf
	case token.THEN:
		p.error("`%s` is a reserved keyword, cannot be used a command name", p.curr.Literal)
		fallthrough
	default:
		return nil
	}
}

func (p *Parser) parseWhileLoop() ast.Node {
	var loop ast.Loop
	loopKeyword := p.curr.Literal
	loop.Negate = p.curr.Type == token.UNTIL
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.DO && p.curr.Type != token.DONE && p.curr.Type != token.EOF {
		loop.Head = append(loop.Head, p.parseCommandList())
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
		p.error("expected `do`, found `%s`", p.curr.Literal)
	}

	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.DONE && p.curr.Type != token.EOF {
		loop.Body = append(loop.Body, p.parseCommandList())
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
		p.error("unexpected token `%s`", p.curr.Literal)
	}

	return loop
}

func (p *Parser) parseForLoop() ast.Node {
	var loop ast.RangeLoop
	p.proceed()
	for p.curr.Type == token.BLANK {
		p.proceed()
	}
	if p.curr.Type != token.WORD {
		p.error("expected identifier after `for`")
	}
	loop.Var = p.curr.Literal
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
			member := p.parseField()
			if member == nil {
				p.error("unexpected token `%s`", p.curr.Literal)
				break
			}
			loop.Operands = append(loop.Operands, member)
			if p.curr.Type == token.BLANK {
				p.proceed()
			}
		}
		if loop.Operands == nil {
			p.error("missing operand after `in`")
		}
	}

	if p.curr.Type == token.SEMICOLON {
		p.proceed()
	}
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	if p.curr.Type != token.DO {
		p.error("expected `do`, found `%s`", p.curr.Literal)
	}
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.DONE && p.curr.Type != token.EOF {
		loop.Body = append(loop.Body, p.parseCommandList())
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
		p.error("expected `done` to close `for` loop")
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
		p.error("unexpected token `%s`", p.curr.Literal)
	}

	return loop
}

func (p *Parser) parseIf() ast.Node {
	var cond ast.If
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.THEN && p.curr.Type != token.FI && p.curr.Type != token.EOF {
		cond.Head = append(cond.Head, p.parseCommandList())
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
		p.error("expected `then`, found `%s`", p.curr.Literal)
	}

	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.FI && p.curr.Type != token.ELIF && p.curr.Type != token.ELSE && p.curr.Type != token.EOF {
		cond.Body = append(cond.Body, p.parseCommandList())
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
			elif.Head = append(elif.Head, p.parseCommandList())
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
			p.error("expected `then`, found `%s`", p.curr.Literal)
		}

		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		for p.curr.Type != token.FI && p.curr.Type != token.ELIF && p.curr.Type != token.ELSE && p.curr.Type != token.EOF {
			elif.Body = append(elif.Body, p.parseCommandList())
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
			cond.Alternate = append(cond.Alternate, p.parseCommandList())
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
		p.error("unexpected token `%s`", p.curr.Literal)
	}

	return cond
}

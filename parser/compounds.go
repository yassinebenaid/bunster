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
			loop.Operands = append(loop.Operands, p.parseField())
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

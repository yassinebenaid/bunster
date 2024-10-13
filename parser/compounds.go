package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) getCompoundParser() func() ast.Node {
	switch p.curr.Type {
	case token.WHILE:
		return p.parseWhileLoop
	default:
		return nil
	}
}

func (p *Parser) parseWhileLoop() ast.Node {
	var loop ast.Loop
	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	for p.curr.Type != token.DO && p.curr.Type != token.EOF {
		loop.Head = append(loop.Head, p.parseCommandList())
		if p.curr.Type == token.SEMICOLON {
			p.proceed()
		}
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
	}

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	loop.Body = append(loop.Body, p.parseCommandList())
	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	return loop
}

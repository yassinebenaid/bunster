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
	return loop
}

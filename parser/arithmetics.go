package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) parseArithmetics() ast.Expression {
	p.proceed()

	var expr ast.Arithmetic
	expr.Expr = p.parseArithmeticExpresion()
	p.proceed()

	return expr
}

func (p *Parser) parseArithmeticExpresion() ast.Expression {
	prefix := p.parsePrefix()

	return prefix
}

func (p *Parser) parsePrefix() ast.Expression {
	switch p.curr.Type {
	case token.INT:
		return ast.Word(p.curr.Literal)
	default:
		return nil
	}

}

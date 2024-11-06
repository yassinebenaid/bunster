package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) parseArithmetics() ast.Expression {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var expr ast.Arithmetic
	expr.Expr = p.parseArithmeticExpresion()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

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
		exp := ast.Number(p.curr.Literal)
		p.proceed()
		return exp
	case token.SIMPLE_EXPANSION, token.WORD:
		exp := ast.Var(p.curr.Literal)
		p.proceed()
		return exp
	default:
		return nil
	}

}

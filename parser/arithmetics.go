package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

type precedence uint

const (
	_ precedence = iota
	BASIC
	ADDITION
)

var precedences = map[token.TokenType]precedence{
	token.PLUS:  ADDITION,
	token.MINUS: ADDITION,
}

func (p *Parser) parseArithmetics() ast.Expression {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var expr ast.Arithmetic
	expr.Expr = p.parseArithmeticExpresion(BASIC)

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	p.proceed()

	return expr
}

func (p *Parser) parseArithmeticExpresion(prec precedence) ast.Expression {
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	exp := p.parsePrefix()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	for prec < precedences[p.curr.Type] {
		exp = p.parseInfix(exp)
	}

	return exp
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
	case token.DOLLAR_DOUBLE_PAREN:
		exp := p.parseArithmetics()
		p.proceed()
		return exp
	case token.DOLLAR_BRACE:
		exp := p.parseParameterExpansion()
		p.proceed()
		return exp
	default:
		return nil
	}

}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	var inf = ast.InfixArithmetic{
		Left:     left,
		Operator: p.curr.Literal,
	}

	switch p.curr.Type {
	case token.PLUS, token.MINUS:
		p.proceed()
		inf.Right = p.parseArithmeticExpresion(ADDITION)
	}

	return inf
}

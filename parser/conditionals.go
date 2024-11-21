package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) parseTestCommand() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	expr := p.parseConditionals()

	if p.curr.Type != token.DOUBLE_RIGHT_BRACKET {
		p.error("expected `]]` to close conditional expression, found `%s`", p.curr)
	}
	p.proceed()

	return ast.Test{Expr: expr}
}

func (p *Parser) parseConditionals() ast.Expression {
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	exp := p.parsePrefixConditional()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	// for prec < infixPrecedences[p.curr.Type] {
	// 	exp = p.parseInfix(exp)
	// }

	// exp = p.parsePostfix(exp)

	return exp
}

func (p *Parser) parsePrefixConditional() ast.Expression {
	switch p.curr.Type {
	case token.MINUS:
		switch p.next.Literal {
		case "a", "b", "c", "d", "e", "f", "g", "h", "k", "p", "r":
			exp := ast.UnaryConditional{
				Operator: p.curr.Literal + p.next.Literal,
			}
			p.proceed()
			p.proceed()
			p.proceed()

			exp.Operand = p.parseExpression()
			return exp
		}

	case token.EOF:
		p.error("bad conditonal expression, unexpected end of file")
	default:
		p.error("bad conditonal expression, unexpected token `%s`", p.curr)
	}

	return nil
}

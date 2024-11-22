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

	if exp := p.parseUnaryConditional(); exp != nil {
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		return exp
	}

	exp := p.parseExpression()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	operator := p.parseConditionalBinaryOperator()
	if operator == "" {
		return exp
	}

	exp = ast.BinaryConditional{
		Left:     exp,
		Operator: operator,
		Right:    p.parseExpression(),
	}
	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	return exp
}

func (p *Parser) parseUnaryConditional() ast.Expression {
	if p.curr.Type == token.MINUS {
		switch p.next.Literal {
		case "a", "b", "c", "d", "e", "f", "g", "h", "k", "p", "r", "s",
			"t", "u", "w", "x", "G", "L", "N", "O", "S", "z", "n", "v":
			if p.next2.Type != token.BLANK {
				break
			}

			u := ast.UnaryConditional{
				Operator: "-" + p.next.Literal,
			}
			p.proceed()
			p.proceed()
			p.proceed()

			if p.curr.Type != token.DOUBLE_RIGHT_BRACKET {
				u.Operand = p.parseExpression()
			}
			if u.Operand == nil {
				p.error("bad conditional expression, expected an operand after %s, found `%s`", u.Operator, p.curr)
			}
			return u
		}
	}

	return nil
}

func (p *Parser) parseConditionalBinaryOperator() string {
	switch p.curr.Type {
	case token.ASSIGN, token.EQ:
		if p.next.Type != token.BLANK {
			break
		}

		operator := p.curr.Literal
		p.proceed()
		p.proceed()

		return operator
	case token.MINUS:
		switch p.next.Literal {
		case "ef", "nt", "ot":
			if p.next2.Type != token.BLANK {
				break
			}

			p.proceed()
			operator := "-" + p.curr.Literal
			p.proceed()
			p.proceed()

			return operator
		}
	}

	return ""
}

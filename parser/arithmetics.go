package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

type precedence uint

const (
	_ precedence = iota
	BASIC
	ASSIGNMENT     //  = *= /= %= += -= <<= >>= &= ^= |=
	CONDITIONAL    // expr ? expr : expr
	LOR            // ||
	LAND           // &&
	BITOR          // |
	BITXOR         // ^
	BITAND         // &
	EQUALITY       // == !=
	COMPARISON     // <= >= < >
	BINSHIFT       // << >>
	ADDITION       // + -
	MULDIVREM      // * / %
	EXPONENTIATION // **
	NEGATION       // ! ~
	UNARY          // - +
	PRE_INCREMENT  // ++id --id
	POST_INCREMENT // id++ id--
)

func (p *Parser) parseArithmetics() ast.Expression {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var expr ast.Arithmetic

	for {
		expr = append(expr, p.parseArithmeticExpresion(BASIC))

		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.COMMA {
			break
		}
		p.proceed()
	}

	if !(p.curr.Type == token.RIGHT_PAREN && p.next.Type == token.RIGHT_PAREN) {
		p.error("expected `))` to close arithmetic expression, found `%s`", p.curr.Literal)
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

	for prec < p.getArithmeticPrecedence() {
		exp = p.parseInfix(exp)
	}

	exp = p.parsePostfix(exp)

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
	case token.INCREMENT, token.DECREMENT:
		exp := ast.PreIncDecArithmetic{
			Operator: p.curr.Literal,
		}
		p.proceed()

		exp.Operand = p.parseArithmeticExpresion(PRE_INCREMENT)
		return exp
	case token.PLUS, token.MINUS:
		exp := ast.Unary{
			Operator: p.curr.Literal,
		}
		p.proceed()

		exp.Operand = p.parseArithmeticExpresion(UNARY)
		return exp
	case token.EXCLAMATION:
		p.proceed()
		exp := ast.Negation{Operand: p.parseArithmeticExpresion(NEGATION)}
		return exp
	case token.TILDE:
		p.proceed()
		exp := ast.BitFlip{Operand: p.parseArithmeticExpresion(NEGATION)}
		return exp
	default:
		return nil
	}

}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	exp := ast.InfixArithmetic{
		Left:     left,
		Operator: p.curr.Literal,
	}

	prec := p.getArithmeticPrecedence()

	p.proceed()

	if p.curr.Type == token.ASSIGN {
		exp.Operator += "="
		p.proceed()
	}

	exp.Right = p.parseArithmeticExpresion(prec)
	return exp
}

func (p *Parser) parsePostfix(left ast.Expression) ast.Expression {
	switch p.curr.Type {
	case token.INCREMENT, token.DECREMENT:
		exp := ast.PostIncDecArithmetic{
			Operand:  left,
			Operator: p.curr.Literal,
		}
		p.proceed()
		return exp
	case token.QUESTION:
		p.proceed()
		exp := ast.Conditional{Test: left}
		exp.Body = p.parseArithmeticExpresion(CONDITIONAL)
		p.proceed()
		exp.Alternate = p.parseArithmeticExpresion(CONDITIONAL)
		return exp
	case token.DOUBLE_GT, token.DOUBLE_LT, token.AMPERSAND, token.CIRCUMFLEX, token.PIPE:
		exp := ast.InfixArithmetic{
			Left:     left,
			Operator: p.curr.Literal + "=",
		}

		p.proceed()
		p.proceed()
		exp.Right = p.parseArithmeticExpresion(ASSIGNMENT)
		return exp
	case token.STAR_ASSIGN, token.SLASH_ASSIGN, token.ASSIGN, token.PLUS_ASSIGN, token.MINUS_ASSIGN:
		exp := ast.InfixArithmetic{
			Left:     left,
			Operator: p.curr.Literal,
		}
		p.proceed()
		exp.Right = p.parseArithmeticExpresion(ASSIGNMENT)
		return exp
	default:
		return left
	}
}

func (p *Parser) getArithmeticPrecedence() precedence {
	switch p.curr.Type {
	case token.OR:
		return LOR
	case token.AND:
		return LAND
	case token.PIPE:
		if p.next.Type == token.ASSIGN {
			return ASSIGNMENT
		}
		return BITOR
	case token.AMPERSAND, token.CIRCUMFLEX:
		if p.next.Type == token.ASSIGN {
			return ASSIGNMENT
		}
		return BITXOR
	case token.EQ, token.NOT_EQ:
		return EQUALITY
	case token.GT, token.LT:
		return COMPARISON
	case token.DOUBLE_GT, token.DOUBLE_LT:
		if p.next.Type == token.ASSIGN {
			return ASSIGNMENT
		}
		return BINSHIFT
	case token.STAR, token.SLASH, token.PERCENT:
		return MULDIVREM
	case token.EXPONENTIATION:
		return EXPONENTIATION
	case token.PLUS, token.MINUS:
		return ADDITION
	default:
		return BASIC
	}
}

package parser

import (
	"github.com/yassinebenaid/ryuko/ast"
	"github.com/yassinebenaid/ryuko/token"
)

type precedence uint

const (
	BASIC          precedence = iota
	LOR                       // ||
	LAND                      // &&
	BITOR                     // |
	BITXOR                    // ^
	BITAND                    // &
	EQUALITY                  // == !=
	COMPARISON                // <= >= < >
	BINSHIFT                  // << >>
	ADDITION                  // + -
	MULDIVREM                 // * / %
	EXPONENTIATION            // **
	NEGATION                  // ! ~
	UNARY                     // - +
	PRE_INCREMENT             // ++id --id
)

var infixPrecedences = map[token.TokenType]precedence{
	token.OR:             LOR,
	token.AND:            LAND,
	token.PIPE:           BITOR,
	token.CIRCUMFLEX:     BITXOR,
	token.AMPERSAND:      BITAND,
	token.EQ:             EQUALITY,
	token.NOT_EQ:         EQUALITY,
	token.GT:             COMPARISON,
	token.LT:             COMPARISON,
	token.GT_EQ:          COMPARISON,
	token.LT_EQ:          COMPARISON,
	token.DOUBLE_GT:      BINSHIFT,
	token.DOUBLE_LT:      BINSHIFT,
	token.STAR:           MULDIVREM,
	token.SLASH:          MULDIVREM,
	token.PERCENT:        MULDIVREM,
	token.EXPONENTIATION: EXPONENTIATION,
	token.PLUS:           ADDITION,
	token.MINUS:          ADDITION,
}

func (p *Parser) parseArithmeticSubstitution() ast.Expression {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	expr := p.parseArithmetics()

	if !(p.curr.Type == token.RIGHT_PAREN && p.next.Type == token.RIGHT_PAREN) {
		p.error("expected `))` to close arithmetic expression, found `%s`", p.curr)
		return nil
	}
	p.proceed()

	return expr
}

func (p *Parser) parseArithmetics() ast.Arithmetic {
	var expr ast.Arithmetic

	for {
		expr = append(expr, p.parseArithmeticExpresion(BASIC))

		if p.curr.Type != token.COMMA {
			break
		}
		p.proceed()
	}

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

	for prec < infixPrecedences[p.curr.Type] {
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
		var exp ast.Expression = ast.Var(p.curr.Literal)
		p.proceed()

		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		switch p.curr.Type {
		case token.INCREMENT, token.DECREMENT:
			exp = ast.PostIncDecArithmetic{Operand: exp, Operator: p.curr.Literal}
			p.proceed()
		}
		return exp
	case token.DOLLAR_DOUBLE_PAREN:
		exp := p.parseArithmeticSubstitution()
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
	case token.LEFT_PAREN:
		p.proceed()
		exp := p.parseArithmeticExpresion(BASIC)

		if p.curr.Type != token.RIGHT_PAREN {
			p.error("expected a closing `)`, found `%s`", p.curr)
			return nil
		}
		p.proceed()
		return exp
	case token.EOF:
		p.error("bad arithmetic expression, unexpected end of file")
		return nil
	default:
		p.error("bad arithmetic expression, unexpected token `%s`", p.curr)
		return nil
	}
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	exp := ast.Binary{
		Left:     left,
		Operator: p.curr.Literal,
	}
	prec := infixPrecedences[p.curr.Type]

	p.proceed()
	exp.Right = p.parseArithmeticExpresion(prec)

	return exp
}

func (p *Parser) parsePostfix(left ast.Expression) ast.Expression {
	switch p.curr.Type {
	case token.QUESTION:
		p.proceed()
		exp := ast.Conditional{Test: left}
		exp.Body = p.parseArithmeticExpresion(BASIC)

		if p.curr.Type != token.COLON {
			p.error("expected a colon `:`, found `%s`", p.curr)
		}

		p.proceed()
		exp.Alternate = p.parseArithmeticExpresion(BASIC)
		return exp
	case token.ASSIGN, token.STAR_ASSIGN, token.SLASH_ASSIGN, token.PLUS_ASSIGN, token.MINUS_ASSIGN,
		token.CIRCUMFLEX_ASSIGN, token.PERCENT_ASSIGN, token.DOUBLE_GT_ASSIGN, token.DOUBLE_LT_ASSIGN,
		token.AMPERSAND_ASSIGN, token.PIPE_ASSIGN:
		exp := ast.Binary{
			Left:     left,
			Operator: p.curr.Literal,
		}
		p.proceed()
		exp.Right = p.parseArithmeticExpresion(BASIC)
		return exp
	default:
		return left
	}
}

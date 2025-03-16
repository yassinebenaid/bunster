package parser

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/token"
)

type precedence uint

const (
	pBASIC          precedence = iota
	pLOR                       // ||
	pLAND                      // &&
	pBITOR                     // |
	pBITXOR                    // ^
	pBITAND                    // &
	pEQUALITY                  // == !=
	pCOMPARISON                // <= >= < >
	pBINSHIFT                  // << >>
	pADDITION                  // + -
	pMULDIVREM                 // * / %
	pEXPONENTIATION            // **
	pNEGATION                  // ! ~
	pUNARY                     // - +
	pPRE_INCREMENT             // ++id --id
)

var infixPrecedences = map[token.TokenType]precedence{
	token.OR:             pLOR,
	token.AND:            pLAND,
	token.PIPE:           pBITOR,
	token.CIRCUMFLEX:     pBITXOR,
	token.AMPERSAND:      pBITAND,
	token.EQ:             pEQUALITY,
	token.NOT_EQ:         pEQUALITY,
	token.GT:             pCOMPARISON,
	token.LT:             pCOMPARISON,
	token.GT_EQ:          pCOMPARISON,
	token.LT_EQ:          pCOMPARISON,
	token.DOUBLE_GT:      pBINSHIFT,
	token.DOUBLE_LT:      pBINSHIFT,
	token.STAR:           pMULDIVREM,
	token.SLASH:          pMULDIVREM,
	token.PERCENT:        pMULDIVREM,
	token.EXPONENTIATION: pEXPONENTIATION,
	token.PLUS:           pADDITION,
	token.MINUS:          pADDITION,
}

func (p *parser) parseArithmeticSubstitution() ast.Expression {
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

func (p *parser) parseArithmetics() ast.Arithmetic {
	var expr ast.Arithmetic

	for {
		expr = append(expr, p.parseArithmeticExpresion(pBASIC))

		if p.curr.Type != token.COMMA {
			break
		}
		p.proceed()
	}

	return expr
}

func (p *parser) parseArithmeticExpresion(prec precedence) ast.Expression {
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

func (p *parser) parsePrefix() ast.Expression {
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
			exp = ast.PostIncDecArithmetic{Operand: string(exp.(ast.Var)), Operator: p.curr.Literal}
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

		op := p.parseArithmeticExpresion(pPRE_INCREMENT)
		if v, ok := op.(ast.Var); !ok {
			p.error("expected a variable name after `%s`", exp.Operator)
		} else {
			exp.Operand = string(v)
		}
		return exp
	case token.PLUS, token.MINUS:
		exp := ast.Unary{
			Operator: p.curr.Literal,
		}
		p.proceed()

		exp.Operand = p.parseArithmeticExpresion(pUNARY)
		return exp
	case token.EXCLAMATION:
		p.proceed()
		exp := ast.Negation{Operand: p.parseArithmeticExpresion(pNEGATION)}
		return exp
	case token.TILDE:
		p.proceed()
		exp := ast.BitFlip{Operand: p.parseArithmeticExpresion(pNEGATION)}
		return exp
	case token.LEFT_PAREN:
		p.proceed()
		exp := p.parseArithmeticExpresion(pBASIC)

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

func (p *parser) parseInfix(left ast.Expression) ast.Expression {
	exp := ast.Binary{
		Left:     left,
		Operator: p.curr.Literal,
	}
	prec := infixPrecedences[p.curr.Type]

	p.proceed()
	exp.Right = p.parseArithmeticExpresion(prec)

	return exp
}

func (p *parser) parsePostfix(left ast.Expression) ast.Expression {
	switch p.curr.Type {
	case token.QUESTION:
		p.proceed()
		exp := ast.Conditional{Test: left}
		exp.Body = p.parseArithmeticExpresion(pBASIC)

		if p.curr.Type != token.COLON {
			p.error("expected a colon `:`, found `%s`", p.curr)
		}

		p.proceed()
		exp.Alternate = p.parseArithmeticExpresion(pBASIC)
		return exp
	case token.ASSIGN, token.STAR_ASSIGN, token.SLASH_ASSIGN, token.PLUS_ASSIGN, token.MINUS_ASSIGN,
		token.CIRCUMFLEX_ASSIGN, token.PERCENT_ASSIGN, token.DOUBLE_GT_ASSIGN, token.DOUBLE_LT_ASSIGN,
		token.AMPERSAND_ASSIGN, token.PIPE_ASSIGN:
		if _, ok := left.(ast.Var); !ok {
			p.error("the operator %q expects a variable name on the left", p.curr)
		}

		exp := ast.Binary{
			Left:     left,
			Operator: p.curr.Literal,
		}
		p.proceed()
		exp.Right = p.parseArithmeticExpresion(pBASIC)
		return exp
	default:
		return left
	}
}

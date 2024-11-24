package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) parseTestCommand() ast.Statement {
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}
	if p.curr.Type == token.DOUBLE_RIGHT_BRACKET {
		p.error("expected a conditional expression before `]]`")
	}
	expr := p.parseTestExpression(false)

	if expr == nil {
		p.error("bad conditional expression, unexpected token `%s`", p.curr)
	}
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	if p.curr.Type != token.DOUBLE_RIGHT_BRACKET {
		p.error("expected `]]` to close conditional expression, found `%s`", p.curr)
	}
	p.proceed()

	return ast.Test{Expr: expr}
}

func (p *Parser) parsePosixTestCommand() ast.Statement {
	testKeyword := p.curr.Type == token.TEST

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	expr := p.parsePosixTestExpression(false)

	if !testKeyword && p.curr.Type != token.RIGHT_BRACKET {
		p.error("expected `]` to close conditional expression, found `%s`", p.curr)
	}
	p.proceed()

	return ast.Test{Expr: expr}
}

func (p *Parser) parseTestExpression(prefix bool) ast.Expression {
	if p.curr.Type == token.BLANK {
		p.proceed()
	}
	var expr ast.Expression
	if p.curr.Type == token.EXCLAMATION {
		p.proceed()
		expr = ast.Negation{Operand: p.parseTestExpression(true)}
	} else if p.curr.Type == token.LEFT_PAREN {
		p.proceed()
		expr = p.parseTestExpression(false)
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.RIGHT_PAREN {
			p.error("expected a closing `)`, found `%s`", p.curr)
		}
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
	} else {
		expr = p.parseConditionals()
	}

	for !prefix && (p.curr.Type == token.AND || p.curr.Type == token.OR) {
		operator := p.curr.Literal
		p.proceed()
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}

		expr = ast.BinaryConditional{
			Left:     expr,
			Operator: operator,
			Right:    p.parseTestExpression(true),
		}
	}

	return expr
}

func (p *Parser) parsePosixTestExpression(prefix bool) ast.Expression {
	var expr ast.Expression
	if p.curr.Type == token.EXCLAMATION {
		p.proceed()
		expr = ast.Negation{Operand: p.parsePosixTestExpression(true)}
	} else if p.curr.Type == token.LEFT_PAREN {
		p.proceed()
		expr = p.parsePosixTestExpression(false)
		if p.curr.Type != token.RIGHT_PAREN {
			p.error("expected a closing `)`, found `%s`", p.curr)
		}
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
	} else {
		expr = p.parseConditionals()
	}

	for !prefix {
		operator := p.parsePosixConditionalBinaryOperator()
		if operator == "" {
			break
		}
		if operator == "-a" {
			operator = "&&"
		} else {
			operator = "||"
		}

		if p.curr.Type == token.BLANK {
			p.proceed()
		}

		expr = ast.BinaryConditional{
			Left:     expr,
			Operator: operator,
			Right:    p.parsePosixTestExpression(true),
		}
	}

	return expr
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

	bin := ast.BinaryConditional{Left: exp, Operator: operator}

	if p.curr.Type != token.DOUBLE_RIGHT_BRACKET {
		if operator == "=~" {
			bin.Right = p.parsePatternExpression()
		} else {
			bin.Right = p.parseExpression()
		}
	}

	if bin.Right == nil {
		p.error("bad conditional expression, expected an operand after `%s`, found `%s`", operator, p.curr)
	}

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	return bin
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
	case token.ASSIGN, token.EQ, token.NOT_EQ, token.GT, token.LT, token.EQ_TILDE:
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

func (p *Parser) parsePosixConditionalBinaryOperator() string {
	if p.curr.Type == token.MINUS {
		switch p.next.Literal {
		case "a", "o":
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

func (p *Parser) parsePatternExpression() ast.Expression {
	var exprs []ast.Expression

loop:
	for {
		switch p.curr.Type {
		case token.BLANK, token.NEWLINE, token.EOF:
			break loop
		case token.SIMPLE_EXPANSION:
			exprs = append(exprs, ast.Var(p.curr.Literal))
		case token.SINGLE_QUOTE:
			exprs = append(exprs, p.parseLiteralString())
		case token.DOUBLE_QUOTE:
			exprs = append(exprs, p.parseString())
		case token.DOLLAR_PAREN:
			exprs = append(exprs, p.parseCommandSubstitution())
		case token.GT_PAREN, token.LT_PAREN:
			exprs = append(exprs, p.parseProcessSubstitution())
		case token.DOLLAR_BRACE:
			exprs = append(exprs, p.parseParameterExpansion())
		default:
			exprs = append(exprs, ast.Word(p.curr.Literal))
		}

		p.proceed()
	}

	return concat(exprs, false)
}

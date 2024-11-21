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

	exp := p.parseExpression()

	if v, ok := exp.(ast.Word); ok {
		switch v {
		case "-a", "-b", "-c", "-d", "-e", "-f", "-g", "-h", "-k", "-p", "-r", "-s",
			"-t", "-u", "-w", "-x", "-G", "-L", "-N", "-O", "-S", "-z", "-n", "-v":
			u := ast.UnaryConditional{
				Operator: string(v),
			}

			if p.curr.Type == token.BLANK {
				p.proceed()
			}
			u.Operand = p.parseExpression()
			return u
		}
	}

	return exp
}

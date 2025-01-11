package parser

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/token"
)

func (p *parser) getBuiltinParser() func() ast.Statement {
	switch p.curr.Type {
	case token.BREAK:
		return p.parseBreak
	case token.FUNCTION:
		return p.parseFunction
	case token.THEN, token.ELIF, token.ELSE, token.FI, token.DO, token.DONE, token.ESAC:
		p.error("`%s` is a reserved keyword, cannot be used a command name", p.curr)
	}
	return nil
}

func (p *parser) parseFunction() ast.Statement {
	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	nameExpr := p.parseExpression()
	if nameExpr == nil {
		p.error("function name is required")
		return nil
	}

	name, ok := nameExpr.(ast.Word)
	if !ok {
		p.error("invalid function name was supplied")
		return nil
	}
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.LEFT_PAREN {
		p.proceed()
		if p.curr.Type == token.BLANK {
			p.proceed()
		}
		if p.curr.Type != token.RIGHT_PAREN {
			p.error("expected `)`, found `%s`", p.curr)
			return nil
		}
		p.proceed()
	}

	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	compound := p.getCompoundParser()
	if compound == nil {
		p.error("function body is expected to be a compound command, found `%s`", p.curr)
		return nil
	}

	return ast.Function{Name: string(name), Command: compound()}
}

func (p *parser) parseBreak() ast.Statement {
	p.proceed()
	if p.loopLevel == 0 {
		p.error("the `break` keyword cannot be used outside loops")
	}
	return ast.Break(p.loopLevel)
}

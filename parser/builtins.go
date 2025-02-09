package parser

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/token"
)

func (p *parser) getBuiltinParser() func() ast.Statement {
	switch p.curr.Type {
	case token.BREAK:
		return p.parseBreak
	case token.CONTINUE:
		return p.parseContinue
	case token.FUNCTION:
		return p.parseFunction
	case token.WAIT:
		return p.parseWait
	case token.LOCAL:
		return p.parseLocal
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

	fn := ast.Function{Name: string(name), Command: compound()}

	switch p.curr.Type {
	case token.SEMICOLON, token.NEWLINE, token.EOF, token.AND, token.OR:
	default:
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return fn
}

func (p *parser) parseBreak() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}
	return ast.Break(1)
}

func (p *parser) parseContinue() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}
	return ast.Continue(1)
}

func (p *parser) parseWait() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if !p.isControlToken() && p.curr.Type != token.EOF {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}
	return ast.Wait{}
}

func (p *parser) parseLocal() ast.Statement {
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	var assignements ast.LocalParameterAssignement

	for {
		if p.curr.Type != token.WORD {
			break
		}
		assignment := ast.Assignement{Name: p.curr.Literal}
		p.proceed()

		if p.curr.Type == token.ASSIGN {
			p.proceed()
			assignment.Value = p.parseExpression()
		}

		if p.curr.Type == token.BLANK {
			p.proceed()
		}

		assignements = append(assignements, assignment)
	}

	if p.curr.Type == token.HASH {
		for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
			p.proceed()
		}
	}

	if assignements == nil || (!p.isControlToken() && p.curr.Type != token.EOF) {
		p.error("unexpected token `%s`", p.curr)
		return nil
	}

	return assignements
}

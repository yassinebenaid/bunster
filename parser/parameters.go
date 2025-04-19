package parser

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/token"
)

func (p *parser) parseAssignement() ast.ParameterAssignement {
	var assigns ast.ParameterAssignement

	for {
		if !(p.curr.Type == token.WORD && p.next.Type == token.ASSIGN) {
			break
		}
		assignment := ast.Assignement{Name: p.curr.Literal}
		p.proceed()
		p.proceed()

		if p.curr.Type == token.LEFT_PAREN {
			p.proceed()

			for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
				p.proceed()
			}

			for {
				if p.curr.Type != token.HASH {
					break
				}
				for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
					p.proceed()
				}
				for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
					p.proceed()
				}
			}

			var al ast.ArrayLiteral
			for p.curr.Type != token.RIGHT_PAREN && p.curr.Type != token.EOF {
				al = append(al, p.parseExpression())
				for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
					p.proceed()
				}
				for {
					if p.curr.Type != token.HASH {
						break
					}
					for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
						p.proceed()
					}
					for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
						p.proceed()
					}
				}
			}

			if p.curr.Type != token.RIGHT_PAREN {
				p.error("expected a closing parenthese `)`, found `%s`", p.curr)
			}
			p.proceed()

			assignment.Value = al
		} else {
			assignment.Value = p.parseExpression()
		}

		assigns = append(assigns, assignment)

		for p.curr.Type == token.BLANK {
			p.proceed()
		}
	}

	return assigns
}

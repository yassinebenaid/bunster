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
		assignment.Value = p.parseExpression()
		assigns = append(assigns, assignment)

		for p.curr.Type == token.BLANK {
			p.proceed()
		}
	}

	return assigns
}

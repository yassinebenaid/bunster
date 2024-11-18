package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) parseAssignement() ast.ParameterAssignement {
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
	}

	return assigns
}

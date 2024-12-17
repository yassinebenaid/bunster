package ast

import "github.com/yassinebenaid/bunster/token"

type ParameterAssignement struct {
	Token        token.Token
	Assignements []Assignement
}

type Assignement struct {
	Name  string
	Value Expression
}

func (ParameterAssignement) node()                   {}
func (p ParameterAssignement) GetToken() token.Token { return p.Token }
func (ParameterAssignement) stmt()                   {}

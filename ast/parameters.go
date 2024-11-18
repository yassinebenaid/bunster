package ast

type ParameterAssignement []Assignement

type Assignement struct {
	Name  string
	Value Expression
}

func (ParameterAssignement) node() {}
func (ParameterAssignement) stmt() {}

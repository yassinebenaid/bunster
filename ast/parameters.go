package ast

type Assignement struct {
	Name  string
	Value Expression
}

type ParameterAssignement []Assignement

func (ParameterAssignement) node() {}
func (ParameterAssignement) stmt() {}

type LocalParameterAssignement []Assignement

func (LocalParameterAssignement) node() {}
func (LocalParameterAssignement) stmt() {}

type ExportParameterAssignement []Assignement

func (ExportParameterAssignement) node() {}
func (ExportParameterAssignement) stmt() {}

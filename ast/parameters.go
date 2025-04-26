package ast

type Assignement struct {
	Name  string
	Value Expression
}

type ParameterAssignement []Assignement

type LocalParameterAssignement []Assignement

type ExportParameterAssignement []Assignement

type ArrayLiteral []Expression
type PositionalSpread struct{}

func (ParameterAssignement) node()       {}
func (ParameterAssignement) stmt()       {}
func (LocalParameterAssignement) node()  {}
func (LocalParameterAssignement) stmt()  {}
func (ExportParameterAssignement) node() {}
func (ExportParameterAssignement) stmt() {}

func (ArrayLiteral) node()          {}
func (ArrayLiteral) expr()          {}
func (ArrayLiteral) string() string { return "" }

func (PositionalSpread) node()          {}
func (PositionalSpread) expr()          {}
func (PositionalSpread) string() string { return "" }

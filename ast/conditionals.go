package ast

type UnaryConditional struct {
	Operand  Expression
	Operator string
}

func (UnaryConditional) node()          {}
func (UnaryConditional) expr()          {}
func (UnaryConditional) string() string { return "" }

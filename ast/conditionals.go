package ast

type UnaryConditional struct {
	Operand  Expression
	Operator string
}

func (UnaryConditional) node()          {}
func (UnaryConditional) expr()          {}
func (UnaryConditional) string() string { return "" }

type BinaryConditional struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (BinaryConditional) node()          {}
func (BinaryConditional) expr()          {}
func (BinaryConditional) string() string { return "" }

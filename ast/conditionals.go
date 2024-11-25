package ast

type BinaryConditional struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (BinaryConditional) node()          {}
func (BinaryConditional) expr()          {}
func (BinaryConditional) string() string { return "" }

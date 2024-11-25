package ast

type Unary struct {
	Operand  Expression
	Operator string
}

func (Unary) node() {}

func (Unary) expr() {}

func (u Unary) string() string {
	return "(" + u.Operator + u.Operand.string() + ")"
}

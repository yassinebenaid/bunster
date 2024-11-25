package ast

type Unary struct {
	Operand  Expression
	Operator string
}

type Binary struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (Unary) node()  {}
func (Binary) node() {}

func (Unary) expr()  {}
func (Binary) expr() {}

func (u Unary) string() string {
	return "(" + u.Operator + u.Operand.string() + ")"
}

func (in Binary) string() string {
	return "(" + in.Left.string() + " " + in.Operator + " " + in.Right.string() + ")"
}

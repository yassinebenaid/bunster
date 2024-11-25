package ast

type Word string

type UnquotedString []Expression

type QuotedString []Expression

type Unary struct {
	Operand  Expression
	Operator string
}

type Binary struct {
	Left     Expression
	Operator string
	Right    Expression
}

type Negation struct {
	Operand Expression
}

func (Word) node()           {}
func (QuotedString) node()   {}
func (UnquotedString) node() {}
func (Unary) node()          {}
func (Binary) node()         {}
func (Negation) node()       {}

func (Word) expr()           {}
func (QuotedString) expr()   {}
func (UnquotedString) expr() {}
func (Unary) expr()          {}
func (Binary) expr()         {}
func (Negation) expr()       {}

func (Word) string() string           { return "" }
func (QuotedString) string() string   { return "" }
func (UnquotedString) string() string { return "" }
func (u Unary) string() string {
	return "(" + u.Operator + u.Operand.string() + ")"
}

func (in Binary) string() string {
	return "(" + in.Left.string() + " " + in.Operator + " " + in.Right.string() + ")"
}

func (n Negation) string() string {
	return "(!" + n.Operand.string() + ")"
}

package ast

type Number string

type Arithmetic []Expression

type InfixArithmetic struct {
	Left     Expression
	Operator string
	Right    Expression
}

type PostIncDecArithmetic struct {
	Operand  Expression
	Operator string
}

type PreIncDecArithmetic struct {
	Operand  Expression
	Operator string
}

type Unary struct {
	Operand  Expression
	Operator string
}

type Negation struct {
	Operand Expression
}

type BitFlip struct {
	Operand Expression
}

type Conditional struct {
	Test      Expression
	Body      Expression
	Alternate Expression
}

func (Arithmetic) node()           {}
func (InfixArithmetic) node()      {}
func (PostIncDecArithmetic) node() {}
func (PreIncDecArithmetic) node()  {}
func (Unary) node()                {}
func (Negation) node()             {}
func (BitFlip) node()              {}
func (Conditional) node()          {}

// Expressions
func (Arithmetic) expr()           {}
func (InfixArithmetic) expr()      {}
func (PostIncDecArithmetic) expr() {}
func (PreIncDecArithmetic) expr()  {}
func (Unary) expr()                {}
func (Negation) expr()             {}
func (BitFlip) expr()              {}
func (Conditional) expr()          {}

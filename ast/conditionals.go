package ast

type ConditionalExpression interface {
	cond()
}

type UnaryConditional struct {
	Operand  Expression
	Operator string
}

func (UnaryConditional) cond() {}

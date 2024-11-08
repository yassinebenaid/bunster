package ast

import (
	"strings"
)

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

func (Number) node()               {}
func (Arithmetic) node()           {}
func (InfixArithmetic) node()      {}
func (PostIncDecArithmetic) node() {}
func (PreIncDecArithmetic) node()  {}
func (Unary) node()                {}
func (Negation) node()             {}
func (BitFlip) node()              {}
func (Conditional) node()          {}

func (Number) expr()               {}
func (Arithmetic) expr()           {}
func (InfixArithmetic) expr()      {}
func (PostIncDecArithmetic) expr() {}
func (PreIncDecArithmetic) expr()  {}
func (Unary) expr()                {}
func (Negation) expr()             {}
func (BitFlip) expr()              {}
func (Conditional) expr()          {}

func (n Number) string() string {
	return string(n)
}

func (in InfixArithmetic) string() string {
	return "(" + in.Left.string() + " " + in.Operator + " " + in.Right.string() + ")"
}
func (PostIncDecArithmetic) string() string {
	return ""
}
func (PreIncDecArithmetic) string() string {
	return ""
}
func (Unary) string() string {
	return ""
}
func (Negation) string() string {
	return ""
}
func (BitFlip) string() string {
	return ""
}
func (c Conditional) string() string {
	return "(" + c.Test.string() + " ? " + c.Body.string() + " : " + c.Alternate.string() + ")"
}

func (Arithmetic) string() string {
	var str string

	return str
}

func (ar Arithmetic) String() string {
	var strs []string

	for _, expr := range ar {
		strs = append(strs, expr.string())
	}

	return strings.Join(strs, ", ")
}

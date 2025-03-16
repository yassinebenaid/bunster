package ast

import (
	"strings"
)

type Number string

type Arithmetic []Expression

type ArithmeticCommand struct {
	Arithmetic
	Redirections []Redirection
}

type PostIncDecArithmetic struct {
	Operand  string
	Operator string
}

type PreIncDecArithmetic struct {
	Operand  string
	Operator string
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
func (ArithmeticCommand) node()    {}
func (PostIncDecArithmetic) node() {}
func (PreIncDecArithmetic) node()  {}
func (BitFlip) node()              {}
func (Conditional) node()          {}

func (Number) expr()               {}
func (Arithmetic) expr()           {}
func (PostIncDecArithmetic) expr() {}
func (PreIncDecArithmetic) expr()  {}
func (BitFlip) expr()              {}
func (Conditional) expr()          {}

func (ArithmeticCommand) stmt() {}

func (n Number) string() string {
	return string(n)
}

func (p PostIncDecArithmetic) string() string {
	return "(" + p.Operand + p.Operator + ")"
}

func (p PreIncDecArithmetic) string() string {
	return "(" + p.Operator + p.Operand + ")"

}

func (bf BitFlip) string() string {
	return "(~" + bf.Operand.string() + ")"
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

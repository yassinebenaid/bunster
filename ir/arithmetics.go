package ir

import "fmt"

type VarIncDec struct {
	Operator string
	Operand  string
	Post     bool
}

func (c VarIncDec) togo() string {
	op := "1"
	if c.Operator == "--" {
		op = "-1"
	}

	return fmt.Sprintf(`runtime.VarIncrement(shell, %q, %s, %t)`, c.Operand, op, c.Post)
}

type ParseInt struct {
	Value Instruction
}

func (c ParseInt) togo() string {
	return fmt.Sprintf("runtime.ParseInt(%s)", c.Value.togo())
}

type UnaryArithmetic struct {
	Operator string
	Operand  Instruction
}

func (c UnaryArithmetic) togo() string {
	return fmt.Sprintf("(%s%s)", c.Operator, c.Operand.togo())
}

type BinaryArithmetic struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c BinaryArithmetic) togo() string {
	return fmt.Sprintf("(%s%s%s)", c.Left.togo(), c.Operator, c.Right.togo())
}

type NegateArithmetic struct {
	Value Instruction
}

func (c NegateArithmetic) togo() string {
	return fmt.Sprintf("runtime.NegateInt(%s)", c.Value.togo())
}

type IntPower struct {
	Operand Instruction
	Pow     Instruction
}

func (c IntPower) togo() string {
	return fmt.Sprintf("runtime.IntPower(%s, %s)", c.Operand.togo(), c.Pow.togo())
}

type CompareInt struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c CompareInt) togo() string {
	return fmt.Sprintf("runtime.CompareInt(%s, %q, %s)", c.Left.togo(), c.Operator, c.Right.togo())
}

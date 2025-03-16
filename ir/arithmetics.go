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

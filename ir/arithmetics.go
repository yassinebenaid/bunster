package ir

import "fmt"

type VarIncDec struct {
	Operator string
	Operand  string
}

func (c VarIncDec) togo() string {
	op := "1"
	if c.Operator == "--" {
		op = "-1"
	}

	return fmt.Sprintf(
		`arithmeticResult = runtime.VarAdd(shell, %q, %s)
		`, c.Operand, op,
	)
}

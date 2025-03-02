package ir

import (
	"fmt"
)

type Compare struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c Compare) togo() string {
	return fmt.Sprintf(
		`if %s %s %s {
			shell.ExitCode = 0 
		} else {
			shell.ExitCode = 1
		}
		`, c.Left.togo(), c.Operator, c.Right.togo())
}

type CompareArithmetics struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c CompareArithmetics) togo() string {
	return fmt.Sprintf(
		`if runtime.NumberCompare(%s, %q, %s) {
			shell.ExitCode = 0 
		} else {
			shell.ExitCode = 1
		}
		`, c.Left.togo(), c.Operator, c.Right.togo())
}

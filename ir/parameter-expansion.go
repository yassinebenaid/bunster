package ir

import "fmt"

type VarLength struct {
	Name string
}

func (d VarLength) togo() string {
	return fmt.Sprintf("runtime.FormatInt(len(shell.ReadVar(%q)))", d.Name)
}

type Substring struct {
	String Instruction
	Offset Instruction
	Length Instruction
}

func (d Substring) togo() string {
	return fmt.Sprintf("runtime.Substring(%s, %s, %s)", d.String.togo(), d.Offset.togo(), d.Length.togo())
}

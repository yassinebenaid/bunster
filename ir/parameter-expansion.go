package ir

import "fmt"

type VarLength struct {
	Name string
}

func (d VarLength) togo() string {
	return fmt.Sprintf("runtime.FormatInt(len([]rune(shell.ReadVar(%q))))", d.Name)
}

type Substring struct {
	String Instruction
	Offset Instruction
	Length Instruction
}

func (d Substring) togo() string {
	return fmt.Sprintf("runtime.Substring(%s, %s, %s)", d.String.togo(), d.Offset.togo(), d.Length.togo())
}

type StringToUpperCase struct {
	String  Instruction
	Pattern Instruction
	All     bool
}

func (d StringToUpperCase) togo() string {
	return fmt.Sprintf("runtime.ChangeStringCase(true, %s, %s, %t)", d.String.togo(), d.Pattern.togo(), d.All)
}

type StringToLowerCase struct {
	String  Instruction
	Pattern Instruction
	All     bool
}

func (d StringToLowerCase) togo() string {
	return fmt.Sprintf("runtime.ChangeStringCase(false, %s, %s, %t)", d.String.togo(), d.Pattern.togo(), d.All)
}

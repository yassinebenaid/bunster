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

type RemoveMatchingPrefix struct {
	String  Instruction
	Pattern Instruction
	Longest bool
}

func (d RemoveMatchingPrefix) togo() string {
	return fmt.Sprintf("runtime.RemoveMatchingPrefix(%s, %s, %t)", d.String.togo(), d.Pattern.togo(), d.Longest)
}

type RemoveMatchingSuffix struct {
	String  Instruction
	Pattern Instruction
	Longest bool
}

func (d RemoveMatchingSuffix) togo() string {
	return fmt.Sprintf("runtime.RemoveMatchingSuffix(%s, %s, %t)", d.String.togo(), d.Pattern.togo(), d.Longest)
}

type ReplaceMatching struct {
	String  Instruction
	Pattern Instruction
	Value   Instruction
	All     bool
}

func (d ReplaceMatching) togo() string {
	return fmt.Sprintf("runtime.ReplaceMatching(%s, %s, %s, %t)", d.String.togo(), d.Pattern.togo(), d.Value.togo(), d.All)
}

type ReplaceMatchingPrefix struct {
	String  Instruction
	Pattern Instruction
	Value   Instruction
}

func (d ReplaceMatchingPrefix) togo() string {
	return fmt.Sprintf("runtime.ReplaceMatchingPrefix(%s, %s, %s)", d.String.togo(), d.Pattern.togo(), d.Value.togo())
}

type ReplaceMatchingSuffix struct {
	String  Instruction
	Pattern Instruction
	Value   Instruction
}

func (d ReplaceMatchingSuffix) togo() string {
	return fmt.Sprintf("runtime.ReplaceMatchingSuffix(%s, %s, %s)", d.String.togo(), d.Pattern.togo(), d.Value.togo())
}

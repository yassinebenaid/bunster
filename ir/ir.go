package ir

import (
	"fmt"
)

type Instruction interface {
	inst()
	fmt.Stringer
}

type Program struct {
	Instructions []Instruction
}

type Declare struct {
	Name  string
	Value Instruction
}

type DeclareSlice string

type Set struct {
	Name  string
	Value Instruction
}

type Append struct {
	Name  string
	Value Instruction
}

type String string
type Literal string

type ReadVar string

type InitCommand struct {
	Name string
	Args string
}

type RunCommanOrFail struct {
	Command string
	Name    string
}

type OpenFile struct {
	Name string
	File Instruction
}

func (Declare) inst()         {}
func (DeclareSlice) inst()    {}
func (Append) inst()          {}
func (ReadVar) inst()         {}
func (Set) inst()             {}
func (String) inst()          {}
func (Literal) inst()         {}
func (InitCommand) inst()     {}
func (OpenFile) inst()        {}
func (RunCommanOrFail) inst() {}

func (p Program) String() string {
	var str = "package main\n\n"

	str += `import (
	"os/exec"

	"ryuko-build/runtime"
)`

	str += `
func Main(shell *runtime.Shell) {
		`

	for _, in := range p.Instructions {
		str += in.String()
	}

	str += `}`
	return str
}

func (d Declare) String() string {
	return fmt.Sprintf("var %s = %s\n", d.Name, d.Value.String())
}

func (d DeclareSlice) String() string {
	return fmt.Sprintf("var %s []string\n", string(d))
}

func (a Set) String() string {
	return fmt.Sprintf("%s = %s\n", a.Name, a.Value.String())
}

func (rv ReadVar) String() string {
	return fmt.Sprintf("shell.ReadVar(%q)", string(rv))
}

func (a Append) String() string {
	return fmt.Sprintf("%s = append(%s, %s)\n", a.Name, a.Name, a.Value.String())
}

func (s String) String() string {
	return fmt.Sprintf("`%s`", string(s))
}

func (s Literal) String() string {
	return fmt.Sprintf(`%s`, string(s))
}

func (ic InitCommand) String() string {
	return fmt.Sprintf("exec.Command(%s, %s...)", ic.Name, ic.Args)
}

func (rcf RunCommanOrFail) String() string {
	return fmt.Sprintf(`
		if err := %s.Run(); err != nil {
			shell.HandleError(%s, err)
		}else{
			shell.ExitCode = 0
		}
		`, rcf.Command, rcf.Name)
}

func (of OpenFile) String() string {
	return fmt.Sprintf(`
		%s, err := runtime.OpenFile(%s)
		if err != nil {
			shell.HandleError("", err)
		}else{
			shell.ExitCode = 0
		}
		`, of.Name, of.File.String())
}

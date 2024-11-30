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

type Set struct {
	Name  string
	Value Instruction
}

type String string
type Literal string

type InitCommand struct {
	Name string
	Args string
}

type RunCommanOrFail struct {
	Name    string
	OnError Instruction
}

type Panic string

func (Declare) inst()         {}
func (Set) inst()             {}
func (String) inst()          {}
func (Literal) inst()         {}
func (InitCommand) inst()     {}
func (RunCommanOrFail) inst() {}

func (p Program) String() string {
	var str = "package main\n\n"

	str += `import (
	"os"
	"os/exec"

	"ryuko-build/runtime"
)`

	str += `
func Main(shell *runtime.Shell) error {
		`

	for _, in := range p.Instructions {
		str += in.String()
	}

	str += `
		return nil
		}`
	return str
}

func (d Declare) String() string {
	return fmt.Sprintf("var %s = %s\n", d.Name, d.Value.String())
}

func (a Set) String() string {
	return fmt.Sprintf("%s = %s\n", a.Name, a.Value.String())
}

func (s String) String() string {
	return fmt.Sprintf(`"%s"`, string(s))
}
func (s Literal) String() string {
	return fmt.Sprintf(`%s`, string(s))
}
func (ic InitCommand) String() string {
	return fmt.Sprintf("exec.Command(%s)", ic.Name)
}
func (rcf RunCommanOrFail) String() string {
	return fmt.Sprintf(`
		if err := %s.Run(); err != nil {
			shell.HandleCommandRunError(err)
		}else{
			shell.ExitCode = 0
		}
		`, rcf.Name)
}

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

type Assign struct {
	Name       string
	Initialize bool
	Value      Instruction
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

func (Assign) inst()          {}
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
func Main(*runtime.Shell) error {
		`

	for _, in := range p.Instructions {
		str += in.String()
	}

	str += `
		return nil
		}`
	return str
}

func (a Assign) String() string {
	op := "="
	if a.Initialize {
		op = ":="
	}
	return fmt.Sprintf("%s %s %s\n", a.Name, op, a.Value.String())
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
			runtime.HandleCommandRunError(err)
		}
		`, rcf.Name)
}

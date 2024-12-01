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

type OpenStream struct {
	Name   string
	Target Instruction
}

type OpenOrCreateStream struct {
	Name   string
	Target Instruction
}

type OpenOrCreateStreamForAppending struct {
	Name   string
	Target Instruction
}

type NewStringStream struct {
	Target Instruction
}

type NewStreamFromFD struct {
	Fd Instruction
}

type DuplicateFD struct {
	Old, New string
}

func (Declare) inst()                        {}
func (DeclareSlice) inst()                   {}
func (Append) inst()                         {}
func (ReadVar) inst()                        {}
func (Set) inst()                            {}
func (String) inst()                         {}
func (Literal) inst()                        {}
func (InitCommand) inst()                    {}
func (OpenStream) inst()                     {}
func (OpenOrCreateStreamForAppending) inst() {}
func (OpenOrCreateStream) inst()             {}
func (NewStringStream) inst()                {}
func (RunCommanOrFail) inst()                {}
func (NewStreamFromFD) inst()                {}
func (DuplicateFD) inst()                    {}

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

func (of OpenStream) String() string {
	return fmt.Sprintf(`
		%s, err := runtime.OpenStream(%s)
		if err != nil {
			shell.HandleError("", err)
		}else{
			shell.ExitCode = 0
		}
		`, of.Name, of.Target.String())
}

func (of OpenOrCreateStream) String() string {
	return fmt.Sprintf(`
		%s, err := runtime.OpenOrCreateStream(%s)
		if err != nil {
			shell.HandleError("", err)
		}else{
			shell.ExitCode = 0
		}
		`, of.Name, of.Target.String())
}

func (of OpenOrCreateStreamForAppending) String() string {
	return fmt.Sprintf(`
		%s, err := runtime.OpenOrCreateStreamForAppending(%s)
		if err != nil {
			shell.HandleError("", err)
		}else{
			shell.ExitCode = 0
		}
		`, of.Name, of.Target.String())
}

func (of NewStringStream) String() string {
	return fmt.Sprintf("runtime.NewStringStream(%s)", of.Target.String())
}

func (nsfd NewStreamFromFD) String() string {
	return fmt.Sprintf("runtime.NewStreamFromFD(%s)", nsfd.Fd.String())
}

func (dfd DuplicateFD) String() string {
	return fmt.Sprintf(`
	if err := runtime.DuplicateFD(%s, %s); err != nil {
		shell.HandleError("", err)
	}
	`, dfd.Old, dfd.New)
}

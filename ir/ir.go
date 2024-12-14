package ir

import (
	"fmt"
)

type Instruction interface {
	togo() string
}

type Program struct {
	Instructions []Instruction
}

func (p Program) String() string {
	var str = `
		package main

		import "bunster-build/runtime"

		func Main(shell *runtime.Shell) {
		`

	for _, in := range p.Instructions {
		str += in.togo()
	}

	str += `}`
	return str
}

type Declare struct {
	Name  string
	Value Instruction
}

func (d Declare) togo() string {
	return fmt.Sprintf("var %s = %s\n", d.Name, d.Value.togo())
}

type DeclareSlice struct {
	Name string
}

func (d DeclareSlice) togo() string {
	return fmt.Sprintf("var %s []string\n", d.Name)
}

type Set struct {
	Name  string
	Value Instruction
}

func (a Set) togo() string {
	return fmt.Sprintf("%s = %s\n", a.Name, a.Value.togo())
}

type Append struct {
	Name  string
	Value Instruction
}

func (a Append) togo() string {
	return fmt.Sprintf("%s = append(%s, %s)\n", a.Name, a.Name, a.Value.togo())
}

type String string

func (s String) togo() string {
	return fmt.Sprintf("`%s`", s)
}

type Concat []Instruction

func (c Concat) togo() string {
	var str string
	for i, ins := range c {
		str += ins.togo()
		if i < len(c)-1 {
			str += "+"
		}
	}
	return str
}

type Literal string

func (s Literal) togo() string {
	return fmt.Sprintf(`%s`, s)
}

type ReadVar string

func (rv ReadVar) togo() string {
	return fmt.Sprintf("shell.ReadVar(%q)", rv)
}

type InitCommand struct {
	Name string
	Args string
}

func (ic InitCommand) togo() string {
	return fmt.Sprintf("shell.Command(%s, %s...)", ic.Name, ic.Args)
}

type RunCommand string

func (r RunCommand) togo() string {
	return fmt.Sprintf(
		`if err := %s.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = %s.ProcessState.ExitCode()
		`, r, r)
}

type StartCommand string

func (r StartCommand) togo() string {
	return fmt.Sprintf(
		`if err := %s.Start(); err != nil {
			shell.HandleError(err)
			return
		}
		`, r)
}

type Closure struct {
	Body []Instruction
}

func (c Closure) togo() string {
	var body string

	for _, ins := range c.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`func(){
			%s
		}()
	`, body)
}

type SetCmdEnv struct {
	Command string
	Key     string
	Value   Instruction
}

func (s SetCmdEnv) togo() string {
	return fmt.Sprintf("%s.Env = append(%s.Env,`%s=` + %s)\n", s.Command, s.Command, s.Key, s.Value.togo())
}

type IfLastExitCode struct {
	Zero bool
	Body Instruction
}

func (i IfLastExitCode) togo() string {
	var condition = "shell.ExitCode == 0"
	if !i.Zero {
		condition = "shell.ExitCode != 0"
	}

	return fmt.Sprintf(
		`if %s {
			%s
		}
		`, condition, i.Body.togo())
}

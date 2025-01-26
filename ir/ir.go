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

		func Main(shell *runtime.Shell, streamManager *runtime.StreamManager) {
		`

	for _, in := range p.Instructions {
		str += in.togo()
	}

	str += `}`
	return str
}

type CloneShell struct{}

func (c CloneShell) togo() string {
	return fmt.Sprintf("shell := shell.Clone()\n")
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
	return string(s)
}

type Label string

func (l Label) togo() string {
	return fmt.Sprintf("goto %s\n%s:\n", l, l)
}

type ReadVar string

func (rv ReadVar) togo() string {
	return fmt.Sprintf("shell.ReadVar(%q)", rv)
}

type SetVar struct {
	Key   string
	Value Instruction
}

func (s SetVar) togo() string {
	return fmt.Sprintf("shell.SetVar(%q, %v)\n", s.Key, s.Value.togo())
}

type ReadSpecialVar string

func (rv ReadSpecialVar) togo() string {
	return fmt.Sprintf("shell.ReadSpecialVar(%q)", rv)
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
		shell.ExitCode = %s.ExitCode
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

type Closure []Instruction

func (c Closure) togo() string {
	var body string
	for _, ins := range c {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`func(){
			%s
		}()
	`, body)
}

type Scope []Instruction

func (s Scope) togo() string {
	var body string
	for _, ins := range s {
		body += ins.togo()
	}

	return fmt.Sprintf(`{
			%s
		}
	`, body)
}

type Gorouting []Instruction

func (g Gorouting) togo() string {
	var body string
	for _, ins := range g {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`go func(){
			%s
		}()
		`, body)
}

type ExpressionClosure struct {
	Body []Instruction
	Name string
}

func (c ExpressionClosure) togo() string {
	var body string

	for _, ins := range c.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`%s, exitCode := func() (string, int) {
			%s
		}()
		if exitCode != 0 {
			shell.ExitCode = exitCode
			return
		}
		`, c.Name, body)
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
	Body []Instruction
}

func (i IfLastExitCode) togo() string {
	var condition = "shell.ExitCode == 0"
	if !i.Zero {
		condition = "shell.ExitCode != 0"
	}

	var body string
	for _, ins := range i.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`if %s {
			%s
		}
		`, condition, body)
}

type If struct {
	Condition Instruction
	Body      []Instruction
	Alternate []Instruction
}

func (i If) togo() string {
	cond := fmt.Sprintf("if %s {\n", i.Condition)
	for _, ins := range i.Body {
		cond += ins.togo()
	}

	if len(i.Alternate) > 0 {
		cond += "} else {"

		for _, ins := range i.Alternate {
			cond += ins.togo()
		}
	}

	return cond + "}\n"
}

type Loop struct {
	Condition Instruction
	Body      []Instruction
}

func (i Loop) togo() string {
	cond := fmt.Sprintf("for %s {\n", i.Condition)
	for _, ins := range i.Body {
		cond += ins.togo()
	}

	return cond + "}\n"
}

type InvertExitCode struct{}

func (i InvertExitCode) togo() string {
	return fmt.Sprintf(
		`if shell.ExitCode == 0 {
			shell.ExitCode = 1
		} else {
			shell.ExitCode = 0
		}
		`)
}

type Function struct {
	Name string
	Body []Instruction
}

func (f Function) togo() string {
	var body string
	for _, ins := range f.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		"shell.RegisterFunction(`%s`, func(shell *runtime.Shell, streamManager *runtime.StreamManager){\n%s\n})\n",
		f.Name, body,
	)
}

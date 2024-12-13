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

type DeclareSlice string

func (d DeclareSlice) togo() string {
	return fmt.Sprintf("var %s []string\n", d)
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

type RunCommanOrFail struct {
	Command string
	Name    string
}

func (rcf RunCommanOrFail) togo() string {
	return fmt.Sprintf(
		`if err := %s.Run(); err != nil {
			shell.HandleError(%s, err)
			return
		}
		shell.ExitCode = %s.ProcessState.ExitCode()
		`, rcf.Command, rcf.Name, rcf.Command)
}

const (
	FLAG_READ   = "STREAM_FLAG_READ"
	FLAG_WRITE  = "STREAM_FLAG_WRITE"
	FLAG_RW     = "STREAM_FLAG_RW"
	FLAG_APPEND = "STREAM_FLAG_APPEND"
)

type OpenStream struct {
	Name   string
	Target Instruction
	Mode   string
}

func (of OpenStream) togo() string {
	return fmt.Sprintf(
		`%s, err := runtime.OpenStream(%s, runtime.%s)
		if err != nil {
			shell.HandleError("", err)
			return
		}
		`, of.Name, of.Target.togo(), of.Mode)
}

type NewStringStream struct {
	Target Instruction
}

func (of NewStringStream) togo() string {
	return fmt.Sprintf("runtime.NewStringStream(%s)", of.Target.togo())
}

type CloneFDT string

func (c CloneFDT) togo() string {
	return fmt.Sprintf(
		`%s, err := shell.CloneFDT()
		if err != nil {
			shell.HandleError("", err)
			return
		}
		defer %s.Destroy()
		`, c, c)
}

type AddStream struct {
	FDT        string
	Fd         string
	StreamName string
}

func (as AddStream) togo() string {
	return fmt.Sprintf("%s.Add(`%s`, %s)\n", as.FDT, as.Fd, as.StreamName)
}

type GetStream struct {
	FDT string
	Fd  Instruction
}

func (as GetStream) togo() string {
	return fmt.Sprintf(`%s.Get(%s)`, as.FDT, as.Fd.togo())
}

type DuplicateStream struct {
	FDT string
	Old string
	New Instruction
}

func (as DuplicateStream) togo() string {
	return fmt.Sprintf(
		`if err := %s.Duplicate("%s", %s); err != nil {
			shell.HandleError("", err)
			return
		}
	`, as.FDT, as.Old, as.New.togo())
}

type CloseStream struct {
	FDT string
	Fd  Instruction
}

func (c CloseStream) togo() string {
	return fmt.Sprintf(
		`if err := %s.Close(%s); err != nil {
			shell.HandleError("", err)
			return
		}
	`, c.FDT, c.Fd.togo())
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

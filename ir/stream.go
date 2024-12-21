package ir

import (
	"fmt"
)

const (
	FLAG_READ   = "STREAM_FLAG_READ"
	FLAG_WRITE  = "STREAM_FLAG_WRITE"
	FLAG_RW     = "STREAM_FLAG_RW"
	FLAG_APPEND = "STREAM_FLAG_APPEND"
)

type OpenStream struct {
	FDT    string
	Name   string
	Target Instruction
	Mode   string
}

func (of OpenStream) togo() string {
	return fmt.Sprintf(
		`%s, err := %s.OpenStream(%s, runtime.%s)
		if err != nil {
			shell.HandleError(err)
			return
		}
		`, of.Name, of.FDT, of.Target.togo(), of.Mode)
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
			shell.HandleError(err)
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
			shell.HandleError(err)
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
			shell.HandleError(err)
			return
		}
	`, c.FDT, c.Fd.togo())
}

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
	Name   string
	Target Instruction
	Mode   string
}

func (of OpenStream) togo() string {
	return fmt.Sprintf(
		`%s, err := streamManager.OpenStream(%s, runtime.%s)
		if err != nil {
			shell.HandleError(err)
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

type CloneFDT struct {
	ND bool
}

func (c CloneFDT) togo() string {
	var d = "defer streamManager.Destroy()\n"
	if c.ND {
		d = ""
	}
	return fmt.Sprintf(
		`streamManager, err := streamManager.Clone()
		if err != nil {
			shell.HandleError(err)
			return
		}
		%s
		`, d)
}

type AddStream struct {
	Fd         string
	StreamName string
}

func (as AddStream) togo() string {
	return fmt.Sprintf("streamManager.Add(`%s`, %s)\n", as.Fd, as.StreamName)
}

type GetStream struct {
	Fd Instruction
}

func (as GetStream) togo() string {
	return fmt.Sprintf(`streamManager.Get(%s)`, as.Fd.togo())
}

type DuplicateStream struct {
	Old string
	New Instruction
}

func (as DuplicateStream) togo() string {
	return fmt.Sprintf(
		`if err := streamManager.Duplicate("%s", %s); err != nil {
			shell.HandleError(err)
			return
		}
	`, as.Old, as.New.togo())
}

type CloseStream struct {
	Fd Instruction
}

func (c CloseStream) togo() string {
	return fmt.Sprintf(
		`if err := streamManager.Close(%s); err != nil {
			shell.HandleError(err)
			return
		}
	`, c.Fd.togo())
}

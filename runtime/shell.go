package runtime

import (
	"fmt"
	"os"
	"os/exec"
)

type Shell struct {
	Stdin  Stream
	Stdout Stream
	Stderr Stream

	ExitCode int

	FDT FileDescriptorTable

	Main func(*Shell)
}

func (shell *Shell) Run() int {
	shell.FDT = make(FileDescriptorTable)

	shell.Main(shell)

	return shell.ExitCode
}

func (shell *Shell) ReadVar(name string) string {
	return os.Getenv(name)
}

func (shell *Shell) HandleError(cmd string, err error) {
	shell.ExitCode = 1

	switch e := err.(type) {
	case *exec.Error:
		fmt.Fprintf(shell.Stderr, "failed to recognize command %q, %v\n", cmd, e.Err)
	default:
		fmt.Fprintln(shell.Stderr, err)
	}
}

func (shell *Shell) CloneFDT() FileDescriptorTable {
	return shell.FDT
}

func (shell *Shell) AddStream(fd string, stream Stream) {
	shell.FDT.Add(fd, stream)
}

func (shell *Shell) GetStream(fd string) (Stream, error) {
	return shell.FDT.Get(fd)
}

func (shell *Shell) DuplicateStream(oldfd, newfd string) error {
	return shell.FDT.Duplicate(oldfd, newfd)
}

func (shell *Shell) CloseStream(fd string) error {
	return shell.FDT.Close(fd)
}

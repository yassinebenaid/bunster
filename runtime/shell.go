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
	shell.FDT.Add("0", shell.Stdin)
	shell.FDT.Add("1", shell.Stdout)
	shell.FDT.Add("2", shell.Stderr)

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

package runtime

import (
	"fmt"
	"io"
	"os/exec"
)

type Shell struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	ExitCode int

	Main func(*Shell)
}

func (shell *Shell) Run() int {
	shell.Main(shell)

	return shell.ExitCode
}

func (shell *Shell) HandleCommandRunError(err error) {
	shell.ExitCode = 1

	if e, ok := err.(*exec.Error); ok {
		fmt.Fprintf(shell.Stderr, "failed to recognize command %q, %v\n", e.Name, e.Err)
	}
}

package runtime

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Shell struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	ExitCode int
	Env      map[string]string

	Main func(*Shell)
}

func (shell *Shell) Run() int {
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

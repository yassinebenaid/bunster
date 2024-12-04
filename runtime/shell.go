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

	FDT map[string]Stream

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

func (shell *Shell) AddStream(fd string, stream Stream) {
	if shell.FDT == nil {
		shell.FDT = make(map[string]Stream)
	}
	shell.FDT[fd] = stream
}

func (shell *Shell) GetStream(fd string) (Stream, error) {
	if stream, ok := shell.FDT[fd]; ok {
		return stream, nil
	}

	return nil, fmt.Errorf("bad file descriptor: %s (did you open it?)", fd)
}

func (shell *Shell) DuplicateStream(oldfd, newfd string) error {
	if shell.FDT == nil {
		shell.FDT = make(map[string]Stream)
	}

	if stream, ok := shell.FDT[newfd]; !ok {
		return fmt.Errorf("bad file descriptor: %s (did you open it?)", newfd)
	} else {
		shell.FDT[oldfd] = stream
		return nil
	}
}

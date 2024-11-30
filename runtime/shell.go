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

	Main func(*Shell) error
}

func (*Shell) Run() int {

	return 0
}

func HandleCommandRunError(err error) {
	if e, ok := err.(*exec.Error); ok {
		fmt.Fprintf(os.Stderr, "failed to recognize command %q, %v", e.Name, e.Err)
	}

	os.Exit(1)
}

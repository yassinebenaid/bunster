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
	case *exec.ExitError:
		// ignore this error
	default:
		fmt.Fprintln(shell.Stderr, err)
	}
}

func (shell *Shell) CloneFDT() (FileDescriptorTable, error) {
	return shell.FDT.clone()
}

func (shell *Shell) Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

func NewPipe() (Stream, Stream, error) {
	return os.Pipe()
}

type PiplineWaitgroupItem struct {
	Wait func() error
}
type PiplineWaitgroup []PiplineWaitgroupItem

func (pw PiplineWaitgroup) Wait() error {
	for _, item := range pw {
		if err := item.Wait(); err != nil {
			return err
		}
	}
	return nil
}

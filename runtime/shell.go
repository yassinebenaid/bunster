package runtime

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

type Shell struct {
	Stdin  Stream
	Stdout Stream
	Stderr Stream

	Args []string

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

func (shell *Shell) ReadSpecialVar(name string) string {
	switch name {
	default:
		index, err := strconv.ParseUint(name, 10, 64)
		if err != nil {
			return ""
		}
		if int(index) < len(shell.Args) {
			return shell.Args[index]
		}
		return ""
	}
}

func (shell *Shell) HandleError(err error) {
	shell.ExitCode = 1

	switch e := err.(type) {
	case *exec.Error:
		fmt.Fprintf(shell.Stderr, "%q: %v\n", e.Name, e.Err)
	case *fs.PathError:
		fmt.Fprintf(shell.Stderr, "%q: %v\n", e.Path, e.Err)
	case *exec.ExitError:
		shell.ExitCode = e.ExitCode()
	default:
		fmt.Fprintln(shell.Stderr, err)
	}
}

func (shell *Shell) CloneFDT() (FileDescriptorTable, error) {
	return shell.FDT.clone()
}

func (shell *Shell) Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Env = syscall.Environ()
	return cmd
}

func NewPipe() (Stream, Stream, error) {
	return os.Pipe()
}

type PiplineWaitgroupItem struct {
	Wait func() error
}
type PiplineWaitgroup []PiplineWaitgroupItem

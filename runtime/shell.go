package runtime

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
)

type Shell struct {
	PID int

	Stdin  Stream
	Stdout Stream
	Stderr Stream

	ExitCode int

	Main func(*Shell, *StreamManager)
	Args []string

	vars sync.Map
}

func (shell *Shell) Run() int {
	streamManager := &StreamManager{
		mappings: map[string]Stream{
			"0": shell.Stdin,
			"1": shell.Stdout,
			"2": shell.Stderr,
		},
	}
	defer streamManager.Destroy()

	shell.Main(shell, streamManager)

	return shell.ExitCode
}

func (shell *Shell) ReadVar(name string) string {
	value, ok := shell.vars.Load(name)
	if ok {
		return value.(string)
	}
	return os.Getenv(name)
}

func (shell *Shell) SetVar(name string, value string) {
	shell.vars.Store(name, value)
}

func (shell *Shell) ReadSpecialVar(name string) string {
	switch name {
	case "$":
		return strconv.FormatInt(int64(shell.PID), 10)
	case "#":
		return strconv.FormatInt(int64(len(shell.Args))-1, 10) // -1 to substract the argument index 0, which is the program name.
	case "?":
		return strconv.FormatInt(int64(shell.ExitCode), 10)
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

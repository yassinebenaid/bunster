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
	parent   *Shell
	PID      int
	Stdin    Stream
	Stdout   Stream
	Stderr   Stream
	ExitCode int
	Main     func(*Shell, *StreamManager)
	Args     []string

	vars      sync.Map
	WaitGroup sync.WaitGroup
}

func (shell *Shell) Run() (exitCode int) {
	streamManager := &StreamManager{
		mappings: make(map[string]*proxyStream),
	}
	streamManager.Add("0", shell.Stdin, true)
	streamManager.Add("1", shell.Stdout, true)
	streamManager.Add("2", shell.Stderr, true)
	defer streamManager.Destroy()

	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprintf(shell.Stderr, "crash: %v\n", err)
			exitCode = 1
		}
	}()

	shell.Main(shell, streamManager)
	exitCode = shell.ExitCode

	return exitCode
}

func (shell *Shell) ReadVar(name string) string {
	value, ok := shell.vars.Load(name)
	if ok {
		return value.(string)
	}
	if shell.parent != nil {
		return shell.parent.ReadVar(name)
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
		return strconv.FormatInt(int64(len(shell.Args)), 10)
	case "?":
		return strconv.FormatInt(int64(shell.ExitCode), 10)
	default:
		index, err := strconv.ParseUint(name, 10, 64)
		if err != nil {
			return ""
		}
		if index < uint64(len(shell.Args)) {
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

func (shell *Shell) Clone() *Shell {
	return &Shell{
		parent:   shell,
		PID:      shell.PID,
		Stdin:    shell.Stdin,
		Stdout:   shell.Stdout,
		Stderr:   shell.Stderr,
		ExitCode: shell.ExitCode,
		Args:     shell.Args,
	}
}

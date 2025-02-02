package runtime

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type Shell struct {
	parent    *Shell
	PID       int
	Stdin     Stream
	Stdout    Stream
	Stderr    Stream
	ExitCode  int
	Main      func(*Shell, *StreamManager)
	Args      []string
	WaitGroup sync.WaitGroup

	vars      *sync.Map
	env       *sync.Map
	functions map[string]func(shell *Shell, stdin, stdout, stderr Stream)
}

func (shell *Shell) Run() (exitCode int) {
	shell.vars = &sync.Map{}
	shell.env = &sync.Map{}
	shell.functions = make(map[string]func(shell *Shell, stdin, stdout, stderr Stream))

	for _, env := range os.Environ() {
		envs := strings.SplitN(env, "=", 2)
		shell.env.Store(envs[0], envs[1])
	}

	streamManager := &StreamManager{
		mappings: make(map[string]*proxyStream),
	}
	streamManager.Add("0", shell.Stdin)
	streamManager.Add("1", shell.Stdout)
	streamManager.Add("2", shell.Stderr)
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
	if value, ok := shell.vars.Load(name); ok {
		return value.(string)
	}
	if value, ok := shell.env.Load(name); ok {
		return value.(string)
	}
	if shell.parent != nil {
		return shell.parent.ReadVar(name)
	}
	return ""
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

func (shell *Shell) Clone() *Shell {
	sh := &Shell{
		parent:    shell,
		PID:       shell.PID,
		Stdin:     shell.Stdin,
		Stdout:    shell.Stdout,
		Stderr:    shell.Stderr,
		ExitCode:  shell.ExitCode,
		Args:      shell.Args,
		functions: shell.functions,
		vars:      &sync.Map{},
		env:       &sync.Map{},
	}

	shell.vars.Range(func(key any, value any) bool {
		sh.vars.Store(key, value)
		return true
	})

	return sh
}

func (shell *Shell) RegisterFunction(name string, handler func(shell *Shell, stdin, stdout, stderr Stream)) {
	shell.functions[name] = handler
}

func (shell *Shell) Command(name string, args ...string) *Command {
	var command Command
	command.shell = shell
	command.Args = args

	if fn := shell.functions[name]; fn != nil {
		command.function = fn
		return &command
	}

	cmd := exec.Command(name, args...)
	command.execCmd = cmd

	return &command
}

type Command struct {
	shell  *Shell
	Args   []string
	Stdin  Stream
	Stdout Stream
	Stderr Stream
	Env    []string

	ExitCode int

	function func(shell *Shell, stdin, stdout, stderr Stream)
	execCmd  *exec.Cmd
	wg       sync.WaitGroup
}

func (cmd *Command) Run() error {
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	if cmd.function == nil {
		cmd.ExitCode = cmd.execCmd.ProcessState.ExitCode()
	}
	return nil
}

func (cmd *Command) Start() error {
	if cmd.function != nil {
		cmd.wg.Add(1)
		go func() {
			shell := Shell{
				parent:    cmd.shell,
				PID:       cmd.shell.PID,
				Stdin:     cmd.shell.Stdin,
				Stdout:    cmd.shell.Stdout,
				Stderr:    cmd.shell.Stderr,
				Args:      append(cmd.shell.Args[:1], cmd.Args...),
				functions: cmd.shell.functions,
				vars:      cmd.shell.vars,
				env:       &sync.Map{},
			}

			cmd.shell.env.Range(func(key any, value any) bool {
				shell.env.Store(key, value)
				return true
			})
			for _, env := range cmd.Env {
				envs := strings.SplitN(env, "=", 2)
				shell.env.Store(envs[0], envs[1])
			}

			cmd.function(&shell, cmd.Stdin, cmd.Stdout, cmd.Stderr)
			cmd.wg.Done()
		}()
		return nil
	}

	cmd.execCmd.Stdin = cmd.Stdin
	cmd.execCmd.Stdout = cmd.Stdout
	cmd.execCmd.Stderr = cmd.Stderr

	cmd.shell.env.Range(func(key any, value any) bool {
		cmd.execCmd.Env = append(cmd.execCmd.Env, fmt.Sprintf("%s=%s", key, value))
		return true
	})
	cmd.execCmd.Env = append(cmd.execCmd.Env, cmd.Env...)

	return cmd.execCmd.Start()
}

func (cmd *Command) Wait() error {
	if cmd.function != nil {
		cmd.wg.Wait()
		return nil
	}
	return cmd.execCmd.Wait()
}

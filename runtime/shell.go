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
	ExitCode  int
	Main      func(*Shell, *StreamManager)
	Args      []string
	WaitGroup sync.WaitGroup

	vars      *repository
	env       *repository
	localVars *repository
	functions map[string]func(shell *Shell, stdin, stdout, stderr Stream)
}

func (shell *Shell) Run(streamManager *StreamManager) (exitCode int) {
	shell.vars = newRepository()
	shell.env = newRepository()
	shell.localVars = newRepository()
	shell.functions = make(map[string]func(shell *Shell, stdin, stdout, stderr Stream))

	for _, env := range os.Environ() {
		envs := strings.SplitN(env, "=", 2)
		shell.env.set(envs[0], envs[1])
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprintf(os.Stderr, "crash: %v\n", err)
			exitCode = 1
		}
	}()

	shell.Main(shell, streamManager)
	exitCode = shell.ExitCode

	return exitCode
}

func (shell *Shell) ReadVar(name string) string {
	if value, ok := shell.getLocalVar(name); ok {
		return value
	}
	if value, ok := shell.vars.get(name); ok {
		return value
	}
	if value, ok := shell.env.get(name); ok {
		return value
	}
	if shell.parent != nil {
		return shell.parent.ReadVar(name)
	}
	return ""
}

func (shell *Shell) setLocalVar(name, value string) bool {
	if _, ok := shell.localVars.get(name); ok {
		shell.localVars.set(name, value)
		return true
	}
	if shell.parent != nil {
		return shell.parent.setLocalVar(name, value)
	}
	return false
}

func (shell *Shell) getLocalVar(name string) (string, bool) {
	if value, ok := shell.localVars.get(name); ok {
		return value, true
	}
	if shell.parent != nil {
		return shell.parent.getLocalVar(name)
	}
	return "", false
}

func (shell *Shell) SetVar(name string, value string) {
	if !shell.setLocalVar(name, value) {
		shell.vars.set(name, value)
	}
}

func (shell *Shell) SetLocalVar(name string, value string) {
	shell.localVars.set(name, value)
}

func (shell *Shell) ReadSpecialVar(name string) string {
	switch name {
	case "$":
		return strconv.FormatInt(int64(shell.PID), 10)
	case "#":
		return strconv.FormatInt(int64(len(shell.Args)-1), 10)
	case "?":
		return strconv.FormatInt(int64(shell.ExitCode), 10)
	case "*", "@":
		return strings.Join(shell.Args[1:], " ")
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

func (shell *Shell) HandleError(sm *StreamManager, err error) {
	shell.ExitCode = 1

	stderr, _err := sm.Get("2")
	if _err != nil {
		return
	}

	switch e := err.(type) {
	case *exec.Error:
		fmt.Fprintf(stderr, "%q: %v\n", e.Name, e.Err)
	case *fs.PathError:
		fmt.Fprintf(stderr, "%q: %v\n", e.Path, e.Err)
	case *exec.ExitError:
		shell.ExitCode = e.ExitCode()
	default:
		fmt.Fprintln(stderr, err)
	}
}

func (shell *Shell) Clone() *Shell {
	sh := &Shell{
		parent:    shell,
		PID:       shell.PID,
		ExitCode:  shell.ExitCode,
		Args:      shell.Args,
		functions: shell.functions,
		vars:      shell.vars.clone(),
		env:       newRepository(),
		localVars: newRepository(),
	}

	//todo: handle locals too
	return sh
}

func (shell *Shell) RegisterFunction(name string, handler func(shell *Shell, stdin, stdout, stderr Stream)) {
	shell.functions[name] = handler
}

func (shell *Shell) Command(name string, args ...string) *Command {
	var command Command
	command.shell = shell
	command.Args = args
	command.Name = name
	command.Env = make(map[string]string)

	if fn := shell.functions[name]; fn != nil {
		command.function = fn
		return &command
	}

	return &command
}

type Command struct {
	shell  *Shell
	Name   string
	Args   []string
	Stdin  Stream
	Stdout Stream
	Stderr Stream
	Env    map[string]string

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
				Args:      append(cmd.shell.Args[:1], cmd.Args...),
				functions: cmd.shell.functions,
				vars:      cmd.shell.vars,
				env:       cmd.shell.env.clone(),
				localVars: newRepository(),
				ExitCode:  cmd.shell.ExitCode,
			}

			for key, value := range cmd.Env {
				shell.env.set(key, value)
			}

			cmd.function(&shell, cmd.Stdin, cmd.Stdout, cmd.Stderr)
			cmd.wg.Done()
		}()
		return nil
	}

	cmd.execCmd = exec.Command(cmd.Name, cmd.Args...) //nolint:gosec
	cmd.execCmd.Stdin = cmd.Stdin
	cmd.execCmd.Stdout = cmd.Stdout
	cmd.execCmd.Stderr = cmd.Stderr

	cmd.shell.env.foreach(func(key string, value string) bool {
		cmd.execCmd.Env = append(cmd.execCmd.Env, fmt.Sprintf("%s=%s", key, value))
		return true
	})
	for key, value := range cmd.Env {
		cmd.execCmd.Env = append(cmd.execCmd.Env, key+"="+value)
	}

	return cmd.execCmd.Start()
}

func (cmd *Command) Wait() error {
	if cmd.function != nil {
		cmd.wg.Wait()
		return nil
	}
	return cmd.execCmd.Wait()
}

type repository struct {
	mx   sync.RWMutex
	data map[string]string
}

func newRepository() *repository {
	return &repository{
		data: make(map[string]string),
	}
}

func (r *repository) get(key string) (string, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	v, ok := r.data[key]
	return v, ok
}

func (r *repository) set(key, value string) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.data[key] = value
}

func (r *repository) clone() *repository {
	var repo = newRepository()
	for key, value := range r.data {
		repo.set(key, value)
	}
	return repo
}

func (r *repository) foreach(fn func(key, value string) bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	for key, value := range r.data {
		if !fn(key, value) {
			break
		}
	}
}

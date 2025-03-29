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

func NewShell() *Shell {
	shell := &Shell{}

	shell.vars = newRepository[string]()
	shell.env = newRepository[string]()
	shell.localVars = newRepository[string]()
	shell.exportedVars = newRepository[struct{}]()
	shell.functions = newRepository[Function]()
	shell.builtins = newRepository[Builtin]()

	for _, env := range os.Environ() {
		envs := strings.SplitN(env, "=", 2)
		shell.env.set(envs[0], envs[1])
	}

	return shell
}

type Function func(shell *Shell, streamManager *StreamManager)
type Builtin func(shell *Shell, stdin, stdout, stderr Stream)

type Shell struct {
	parent    *Shell
	PID       int
	Path      string
	ExitCode  int
	Args      []string
	WaitGroup sync.WaitGroup
	Embed     fs.FS

	vars         *repository[string]
	env          *repository[string]
	localVars    *repository[string]
	exportedVars *repository[struct{}]
	functions    *repository[Function]
	builtins     *repository[Builtin]
	defered      []func(*Shell, *StreamManager)
}

func (shell *Shell) Shift(n int) {
	if n <= len(shell.parent.Args) {
		shell.parent.Args = shell.parent.Args[n:]
	} else {
		shell.parent.Args = nil
	}
}

func (shell *Shell) Exit(ecode string) error {
	code, err := strconv.Atoi(ecode)
	if err != nil {
		return fmt.Errorf("exit: %q is not a valid code", ecode)
	}

	os.Exit(code)
	return nil
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

func (shell *Shell) VarIsSet(name string) bool {
	if _, ok := shell.getLocalVar(name); ok {
		return true
	}
	if _, ok := shell.vars.get(name); ok {
		return true
	}
	if _, ok := shell.env.get(name); ok {
		return true
	}
	return false
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

func (shell *Shell) SetExportVar(name string, value string) {
	shell.exportedVars.set(name, struct{}{})
	shell.vars.set(name, value)
}

func (shell *Shell) MarkVarAsExported(name string) {
	shell.exportedVars.set(name, struct{}{})
}

func (shell *Shell) ReadSpecialVar(name string) string {
	switch name {
	case "0":
		return shell.Path
	case "$":
		return strconv.FormatInt(int64(shell.PID), 10)
	case "#":
		return strconv.FormatInt(int64(len(shell.Args)), 10)
	case "?":
		return strconv.FormatInt(int64(shell.ExitCode), 10)
	case "*", "@":
		return strings.Join(shell.Args, " ")
	default:
		index, err := strconv.ParseUint(name, 10, 64)
		if err != nil {
			return ""
		}
		if index <= uint64(len(shell.Args)) {
			return shell.Args[index-1]
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
		PID:          shell.PID,
		Path:         shell.Path,
		ExitCode:     shell.ExitCode,
		Args:         shell.Args,
		Embed:        shell.Embed,
		functions:    shell.functions.clone(),
		builtins:     shell.builtins.clone(),
		vars:         shell.vars.clone(),
		localVars:    shell.localVars.clone(),
		env:          shell.env.clone(),
		exportedVars: shell.exportedVars.clone(),
	}

	return sh
}

func (shell *Shell) RegisterFunction(name string, handler Function) {
	shell.functions.set(name, handler)
}

func (shell *Shell) RegisterBuiltin(name string, handler Builtin) {
	shell.builtins.set(name, handler)
}

func (shell *Shell) Defer(handler func(*Shell, *StreamManager)) {
	shell.defered = append(shell.defered, handler)
}

func (shell *Shell) Terminate(streamManager *StreamManager) {
	// defered commands run in LIFO order
	for i := len(shell.defered) - 1; i >= 0; i-- {
		shell.defered[i](shell, streamManager)
	}
}

func (shell *Shell) Exec(streamManager *StreamManager, name string, args []string, env map[string]string) error {
	var childShell Shell
	function, isFunc := shell.functions.get(name)
	builtin, isBuiltin := shell.builtins.get(name)

	if isFunc || isBuiltin {
		childShell = Shell{
			parent:       shell,
			Path:         shell.Path,
			PID:          shell.PID,
			Embed:        shell.Embed,
			Args:         args,
			functions:    shell.functions,
			builtins:     shell.builtins,
			vars:         shell.vars,
			env:          shell.env.clone(),
			localVars:    newRepository[string](),
			ExitCode:     shell.ExitCode,
			exportedVars: shell.exportedVars,
		}

		for key, value := range env {
			childShell.env.set(key, value)
		}
	}

	if isFunc {
		function(&childShell, streamManager)
		shell.ExitCode = childShell.ExitCode
		return nil
	}

	stdin, err := streamManager.Get("0")
	if err != nil {
		return err
	}
	stdout, err := streamManager.Get("1")
	if err != nil {
		return err
	}
	stderr, err := streamManager.Get("2")
	if err != nil {
		return err
	}

	if isBuiltin {
		builtin(&childShell, stdin, stdout, stderr)
		shell.ExitCode = childShell.ExitCode
		return nil
	}

	execCmd := exec.Command(name, args...) //nolint:gosec
	execCmd.Stdin = stdin
	execCmd.Stdout = stdout
	execCmd.Stderr = stderr

	shell.env.foreach(func(key string, value string) bool {
		execCmd.Env = append(execCmd.Env, key+"="+value)
		return true
	})
	shell.exportedVars.foreach(func(key string, _ struct{}) bool {
		execCmd.Env = append(execCmd.Env, key+"="+shell.ReadVar(key))
		return true
	})
	for key, value := range env {
		execCmd.Env = append(execCmd.Env, key+"="+value)
	}

	if err := execCmd.Run(); err != nil {
		return err
	}

	shell.ExitCode = execCmd.ProcessState.ExitCode()
	return nil
}

type repository[T any] struct {
	mx   sync.RWMutex
	data map[string]T
}

func newRepository[T any]() *repository[T] {
	return &repository[T]{
		data: make(map[string]T),
	}
}

func (r *repository[T]) get(key string) (T, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	v, ok := r.data[key]
	return v, ok
}

func (r *repository[T]) set(key string, value T) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.data[key] = value
}

func (r *repository[T]) clone() *repository[T] {
	var repo = newRepository[T]()
	for key, value := range r.data {
		repo.set(key, value)
	}
	return repo
}

func (r *repository[T]) foreach(fn func(key string, value T) bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	for key, value := range r.data {
		if !fn(key, value) {
			break
		}
	}
}

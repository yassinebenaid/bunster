package runtime

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func NewShell() *Shell {
	shell := &Shell{
		vars:         newRepository[parameter](),
		env:          newRepository[parameter](),
		localVars:    newRepository[parameter](),
		exportedVars: newRepository[struct{}](),
		functions:    newRepository[Function](),
		builtins:     newRepository[Builtin](),
	}

	for _, env := range os.Environ() {
		envs := strings.SplitN(env, "=", 2)
		shell.env.set(envs[0], parameter{value: envs[1]})
	}

	return shell
}

type Function func(shell *Shell, streamManager *StreamManager)
type Builtin func(shell *Shell, stdin, stdout, stderr Stream)

type Shell struct {
	parent    *Shell
	PID       int
	Arg0      string
	CWD       string
	ExitCode  int
	Args      []string
	WaitGroup sync.WaitGroup
	Embed     fs.FS

	vars         *repository[parameter]
	env          *repository[parameter]
	localVars    *repository[parameter]
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

func (shell *Shell) CD(dir string) {
	shell.CWD = dir

	if shell.parent != nil {
		shell.parent.CD(dir)
	}
}

func (shell *Shell) Path(p string) string {
	if filepath.IsAbs(p) {
		return p
	}

	return filepath.Join(shell.CWD, p)
}

func (shell *Shell) Exit(ecode string) error {
	code, err := strconv.Atoi(ecode)
	if err != nil {
		return fmt.Errorf("exit: %q is not a valid code", ecode)
	}

	os.Exit(code)
	return nil
}

func (shell *Shell) UnsetFunctions(names ...string) {
	for _, name := range names {
		shell.functions.forget(name)
	}
}

func (shell *Shell) ReadSpecialVar(name string) string {
	switch name {
	case "0":
		return shell.Arg0
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
		Arg0:         shell.Arg0,
		CWD:          shell.CWD,
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

func (shell *Shell) IsFunction(name string) bool {
	_, ok := shell.functions.get(name)
	return ok
}

func (shell *Shell) IsBuiltin(name string) bool {
	_, ok := shell.builtins.get(name)
	return ok
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
			Arg0:         shell.Arg0,
			CWD:          shell.CWD,
			PID:          shell.PID,
			Embed:        shell.Embed,
			Args:         args,
			functions:    shell.functions,
			builtins:     shell.builtins,
			vars:         shell.vars,
			env:          shell.env.clone(),
			localVars:    newRepository[parameter](),
			ExitCode:     shell.ExitCode,
			exportedVars: shell.exportedVars,
		}

		for key, value := range env {
			childShell.env.set(key, parameter{value: value})
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
	execCmd.Dir = shell.CWD

	shell.env.foreach(func(key string, p parameter) bool {
		execCmd.Env = append(execCmd.Env, key+"="+p.String())
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

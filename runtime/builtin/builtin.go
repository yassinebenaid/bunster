package builtin

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/yassinebenaid/bunster/runtime"
)

func Register(shell *runtime.Shell) {
	shell.RegisterBuiltin("true", True)
	shell.RegisterBuiltin("false", False)
	shell.RegisterBuiltin("loadenv", Loadenv)
	shell.RegisterBuiltin("embed", Embed)
	shell.RegisterBuiltin("shift", Shift)
	shell.RegisterBuiltin("cd", CD)
	shell.RegisterBuiltin("pwd", Pwd)
	shell.RegisterBuiltin("which", Which)
}

func True(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	shell.ExitCode = 0
}

func False(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	shell.ExitCode = 1
}

func CD(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	var path string
	if len(shell.Args) == 0 {
		path = shell.ReadVar("HOME")
	} else if len(shell.Args) == 1 {
		path = shell.Args[0]
	} else {
		fmt.Fprintf(stderr, "cd: expectes exactly one argument\n")
		shell.ExitCode = 1
		return
	}

	path = shell.Path(path)

	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(stderr, "cd: %v\n", err)
		shell.ExitCode = 1
		return
	}

	if !info.IsDir() {
		fmt.Fprintf(stderr, "cd: %q is not a directory\n", path)
		shell.ExitCode = 1
		return
	}

	shell.CD(path)
}

func Shift(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	if len(shell.Args) > 1 {
		fmt.Fprintf(stderr, "embed: expected 1 or 0 arguments, got %d\n", len(shell.Args))
		shell.ExitCode = 1
		return
	}

	if len(shell.Args) < 1 {
		shell.Shift(1)
		return
	}

	if n, err := strconv.Atoi(shell.Args[0]); err != nil {
		fmt.Fprintf(stderr, "embed: %q is not a valid integer\n", shell.Args[0])
		shell.ExitCode = 1
		return
	} else {
		shell.Shift(n)
	}

}

func Pwd(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	fmt.Fprintln(stdout, shell.CWD)
}

func Which(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	if len(shell.Args) != 1 {
		fmt.Fprintln(stderr, "which: expected exactly 1")
		shell.ExitCode = 1
		return
	}

	if shell.IsBuiltin(shell.Args[0]) {
		fmt.Fprintln(stdout, "builtin")
		return
	}

	path, err := exec.LookPath(shell.Args[0])
	if err != nil {
		shell.ExitCode = 1
		return
	}

	fmt.Fprintln(stdout, path)
}

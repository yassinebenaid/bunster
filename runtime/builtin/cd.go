package builtin

import (
	"fmt"
	"os"

	"github.com/yassinebenaid/bunster/runtime"
)

func CD(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	var path string
	if len(shell.Args) == 0 {
		path = shell.ReadVar("HOME")
	} else if len(shell.Args) == 1 {
		path = shell.Path(shell.Args[0])
	} else {
		fmt.Fprintf(stderr, "cd: expectes exactly one argument\n")
		shell.ExitCode = 1
		return
	}

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

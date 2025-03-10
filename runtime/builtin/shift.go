package builtin

import (
	"fmt"
	"strconv"

	"github.com/yassinebenaid/bunster/runtime"
)

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

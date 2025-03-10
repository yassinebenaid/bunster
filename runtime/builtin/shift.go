package builtin

import (
	"fmt"
	"strconv"

	"github.com/yassinebenaid/bunster/runtime"
)

func Shift(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	if len(shell.Args) > 2 {
		fmt.Fprintf(stderr, "embed: expected 1 or 0 arguments, got %d\n", len(shell.Args))
		shell.ExitCode = 1
		return
	}

	if len(shell.Args) == 2 {
		if n, err := strconv.Atoi(shell.Args[1]); err != nil {
			fmt.Fprintf(stderr, "embed: bad argument, %v\n", err)
			shell.ExitCode = 1
			return
		} else {
			shell.Shift(n)
		}
	} else {
		shell.Shift(1)
	}

}

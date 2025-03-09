package builtin

import (
	"fmt"
	"io"
	"strings"

	"github.com/yassinebenaid/bunster/runtime"
)

func Embed(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	if shell.Embed == nil {
		fmt.Fprintf(stderr, "embed: no files were embedded\n")
		shell.ExitCode = 1
		return
	}

	if len(shell.Args) != 3 {
		fmt.Fprintf(stderr, "embed: expected 2 arguments, got %d\n", len(shell.Args))
		shell.ExitCode = 1
		return
	}

	command, path := shell.Args[1], shell.Args[2]

	switch command {
	case "cat":
		f, err := shell.Embed.Open(path)
		if err != nil {
			fmt.Fprintf(stderr, "embed: %v\n", err)
			shell.ExitCode = 1
			return
		}
		if _, err := io.Copy(stdout, f); err != nil {
			fmt.Fprintf(stderr, "embed: %v\n", err)
			shell.ExitCode = 1
			return
		}
	case "ls":
		de, err := shell.Embed.ReadDir(path)
		if err != nil {
			fmt.Fprintf(stderr, "embed: %v\n", err)
			shell.ExitCode = 1
			return
		}

		var files []string
		for _, entry := range de {
			files = append(files, entry.Name())
		}

		fmt.Fprintln(stdout, strings.Join(files, "\n"))
	default:
		fmt.Fprintf(stderr, "embed: %q is not a valid command\n", command)
		shell.ExitCode = 1
	}
}

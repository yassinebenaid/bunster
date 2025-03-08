package builtin

import (
	"fmt"
	"io"
	"io/fs"
	"strings"

	"github.com/yassinebenaid/bunster/runtime"
)

func Embed(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	if len(shell.Args) != 3 {
		fmt.Fprintf(stderr, "embed: expected 2 arguments, got %d\n", len(shell.Args))
		shell.ExitCode = 1
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
		var files []string
		err := fs.WalkDir(shell.Embed, path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return err
		})
		if err != nil {
			fmt.Fprintf(stderr, "embed: %v\n", err)
			shell.ExitCode = 1
			return
		}
		fmt.Fprintln(stdout, strings.Join(files, "\n"))
	default:
		fmt.Fprintf(stderr, "embed: %q is not a valid command\n", command)
		shell.ExitCode = 1
	}
}

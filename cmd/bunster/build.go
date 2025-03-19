package main

import (
	"context"
	"os"
	"path"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster/builder"
)

func buildCMD(_ context.Context, cmd *cli.Command) error {
	destination := cmd.String("o")
	if !path.IsAbs(destination) {
		currWorkdir, err := os.Getwd()
		if err != nil {
			return err
		}
		destination = path.Join(currWorkdir, destination)
	}

	builder := builder.Builder{
		Workdir:    ".test",
		Builddir:   path.Join(os.TempDir(), "bunster-build"),
		OutputFile: destination,
	}

	return builder.Build()
}

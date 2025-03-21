package main

import (
	"context"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster/builder"
)

func geneateCMD(_ context.Context, cmd *cli.Command) error {
	builder := builder.Builder{
		Workdir:  ".",
		Builddir: cmd.String("o"),
		Gofmt:    true,
	}

	return builder.Generate()
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/godump"
)

func main() {
	app := cli.Command{
		Name: "bunster",
		Commands: []*cli.Command{
			{
				Name:        "ast",
				Description: "Print the script ast",
				Action:      astCMD,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "no-ansi", Aliases: []string{"n"}},
				},
			},
			{
				Name:        "build",
				Description: "Build a script",
				Action:      buildCMD,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "o", Required: true},
					&cli.StringFlag{Name: "build-space"},
				},
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func astCMD(_ context.Context, cmd *cli.Command) error {
	filename := cmd.Args().Get(0)
	v, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	script, err := parser.Parse(
		lexer.New(v),
	)

	if err != nil {
		return err
	}

	var d godump.Dumper
	d.ShowPrimitiveNamedTypes = true

	if !cmd.Bool("no-ansi") {
		d.Theme = godump.DefaultTheme
	}

	return d.Println(script)
}

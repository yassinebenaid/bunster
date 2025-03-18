package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/godump"
)

func main() {
	app := cli.Command{
		Name:  "bunster",
		Usage: "compile shell script to self-contained executable programs",
		Commands: []*cli.Command{
			{
				Name:   "ast",
				Usage:  "Print the script ast",
				Action: astCMD,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "no-ansi", Aliases: []string{"n"}},
				},
			},
			{
				Name:   "build",
				Usage:  "Build a script",
				Action: buildCMD,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "o", Required: true},
				},
			},
			{
				Name:   "mod",
				Usage:  "Build a module",
				Action: mod,
			},
			{
				Name:   "generate",
				Usage:  "Generate the Go module out of a script",
				Action: geneateCMD,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "o", Required: true},
				},
			},
			{
				Name:  "version",
				Usage: "Print bunster version",
				Action: func(ctx context.Context, c *cli.Command) error {
					fmt.Println(strings.TrimSpace(bunster.Version))
					return nil
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
		lexer.New([]rune(string(v))),
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

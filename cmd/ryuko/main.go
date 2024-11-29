package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/ryuko"
	"github.com/yassinebenaid/ryuko/lexer"
	"github.com/yassinebenaid/ryuko/parser"
)

func main() {
	app := cli.Command{
		Name: "ryuko",
		Commands: []*cli.Command{
			{
				Name:        "tree",
				Description: "Print the script tree",
				Action:      treeCMD,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "no-ansi", Aliases: []string{"n"}},
				},
			},
			{
				Name:        "build",
				Description: "Build the script as go program",
				Action:      buildCMD,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "o", Required: true},
				},
			},
			{
				Name: "test",
				Action: func(ctx context.Context, c *cli.Command) error {
					r, err := ryuko.RuntimeFS.Open("runtime/shell.go")
					if err != nil {
						return err
					}

					_, err = io.Copy(os.Stdout, r)

					return err
				},
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func treeCMD(_ context.Context, cmd *cli.Command) error {
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

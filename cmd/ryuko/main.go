package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/ryuko/generator"
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

func buildCMD(_ context.Context, cmd *cli.Command) error {
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

	program := generator.Generate(script)

	var instructions string

	for _, ins := range program.Instructions {
		instructions += ins.String() + "\n"
	}

	var _prog = fmt.Sprintf(`package main

import "os/exec"

func main(){
	%s
}
	`, instructions)

	wd, err := os.MkdirTemp(os.TempDir(), "ryuko-build-*")
	if err != nil {
		return err
	}

	err = os.WriteFile(wd+"/main.go", []byte(_prog), 0666)
	if err != nil {
		return err
	}

	err = os.WriteFile(wd+"/go.mod", []byte("module ryuko-build\ngo 1.22.3"), 0666)
	if err != nil {
		return err
	}

	gocmd := exec.Command("go", "build", "-o", "build.bin")
	gocmd.Stdin = os.Stdin
	gocmd.Stdout = os.Stdout
	gocmd.Stderr = os.Stderr
	gocmd.Dir = wd
	if err := gocmd.Run(); err != nil {
		return err
	}

	if err := os.Rename(path.Join(wd, "build.bin"), cmd.String("o")); err != nil {
		return err
	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
)

func geneateCMD(_ context.Context, cmd *cli.Command) error {
	filename := cmd.Args().Get(0)
	if filename == "" {
		return fmt.Errorf("failname is reqired")
	}
	v, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	script, err := parser.Parse(lexer.New(v))
	if err != nil {
		return err
	}

	if err := analyser.Analyse(script); err != nil {
		return err
	}

	program := generator.Generate(script)
	workdir := cmd.String("o")

	if err := os.RemoveAll(workdir); err != nil {
		return err
	}

	if err := os.MkdirAll(workdir, 0777); err != nil {
		return err
	}

	err = os.WriteFile(path.Join(workdir, "program.go"), []byte(program.String()), 0666)
	if err != nil {
		return err
	}

	if err := cloneRuntime(workdir); err != nil {
		return err
	}

	if err := cloneStubs(workdir); err != nil {
		return err
	}

	// we ignore the error, because this is just an optional step that shouldn't stop us from building the binary
	_ = exec.Command("gofmt", "-w", workdir).Run()

	return nil
}

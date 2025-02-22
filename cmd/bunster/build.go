package main

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster"
	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
)

func buildCMD(_ context.Context, cmd *cli.Command) error {
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

	wd := path.Join(os.TempDir(), "bunster-build")
	if err := os.RemoveAll(wd); err != nil {
		return err
	}

	if err := os.MkdirAll(wd, 0700); err != nil {
		return err
	}

	err = os.WriteFile(path.Join(wd, "program.go"), []byte(program.String()), 0600)
	if err != nil {
		return err
	}

	if err := cloneRuntime(wd); err != nil {
		return err
	}

	if err := cloneStubs(wd); err != nil {
		return err
	}

	// we ignore the error, because this is just an optional step that shouldn't stop us from building the binary
	_ = exec.Command("gofmt", "-w", wd).Run()

	destination := cmd.String("o")
	if !path.IsAbs(destination) {
		currWorkdir, err := os.Getwd()
		if err != nil {
			return err
		}
		destination = path.Join(currWorkdir, destination)
	}

	gocmd := exec.Command("go", "build", "-o", destination)
	gocmd.Stdin = os.Stdin
	gocmd.Stdout = os.Stdout
	gocmd.Stderr = os.Stderr
	gocmd.Dir = wd
	if err := gocmd.Run(); err != nil {
		return err
	}

	return nil
}

func cloneRuntime(dst string) error {
	return fs.WalkDir(bunster.RuntimeFS, "runtime", func(dpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return os.MkdirAll(path.Join(dst, dpath), 0766)
		}

		if strings.HasSuffix(dpath, "_test.go") {
			return nil
		}

		content, err := bunster.RuntimeFS.ReadFile(dpath)
		if err != nil {
			return err
		}

		content = bytes.ReplaceAll(content, []byte("github.com/yassinebenaid/bunster"), []byte("bunster-build"))

		return os.WriteFile(path.Join(dst, dpath), content, 0600)
	})
}

func cloneStubs(dst string) error {
	if err := os.WriteFile(path.Join(dst, "main.go"), bunster.MainGoStub, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(dst, "go.mod"), bunster.GoModStub, 0600); err != nil {
		return err
	}

	return nil
}

package main

import (
	"context"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster"
)

func mod(_ context.Context, cmd *cli.Command) error {
	files, err := filepath.Glob("*.sh")
	if err != nil {
		return err
	}

	module := make(map[string][]rune)
	for _, file := range files {
		v, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		module[file] = []rune(string(v))
	}

	workdir := path.Join(os.TempDir(), "bunster-build")
	if err := os.RemoveAll(workdir); err != nil {
		return err
	}

	if err := os.MkdirAll(workdir, 0700); err != nil {
		return err
	}

	if err := bunster.Mod(workdir, module); err != nil {
		return err
	}

	// we ignore the error, because this is just an optional step that shouldn't stop us from building the binary
	_ = exec.Command("gofmt", "-w", workdir).Run()

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
	gocmd.Dir = workdir
	if err := gocmd.Run(); err != nil {
		return err
	}

	return nil
}

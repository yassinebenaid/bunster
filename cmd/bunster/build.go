package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/bunster"
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

	workdir := path.Join(os.TempDir(), "bunster-build")
	if err := os.RemoveAll(workdir); err != nil {
		return err
	}

	if err := os.MkdirAll(workdir, 0700); err != nil {
		return err
	}

	if err := bunster.Generate(".", workdir, v); err != nil {
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

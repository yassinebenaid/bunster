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

func geneateCMD(_ context.Context, cmd *cli.Command) error {
	filename := cmd.Args().Get(0)
	if filename == "" {
		return fmt.Errorf("failname is reqired")
	}
	v, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	program, err := bunster.Compile(v)
	if err != nil {
		return err
	}

	workdir := cmd.String("o")

	if err := os.RemoveAll(workdir); err != nil {
		return err
	}

	if err := os.MkdirAll(workdir, 0777); err != nil {
		return err
	}

	err = os.WriteFile(path.Join(workdir, "program.go"), []byte(program.String()), 0600)
	if err != nil {
		return err
	}

	if err := bunster.CloneAssets(workdir, program.Embeds); err != nil {
		return err
	}

	// we ignore the error, because this is just an optional step that shouldn't stop us from building the binary
	_ = exec.Command("gofmt", "-w", workdir).Run()

	return nil
}

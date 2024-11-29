package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/ryuko/generator"
	"github.com/yassinebenaid/ryuko/lexer"
	"github.com/yassinebenaid/ryuko/parser"
)

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

	// main_stub, err := ryuko.StubsFS.Open("stubs/main.go.stub")
	// if err != nil {
	// 	return fmt.Errorf("internal error: failed to load necessary assets, %v", err)
	// }

	var _prog = fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

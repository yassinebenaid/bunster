package generator_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/bunster/pkg/tst"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestGenerator(t *testing.T) {
	testFiles, err := filepath.Glob("./tests/*.tst")
	if err != nil {
		t.Fatalf("Failed to `Glob` test files, %v", err)
	}

	for _, testFile := range testFiles {
		t.Run(testFile, func(t *testing.T) {
			file, err := os.Open(testFile)
			if err != nil {
				t.Fatalf("Failed to open test file, %v", err)
			}

			test, err := tst.Parse(file)
			if err != nil {
				t.Fatal(err)
			}

			for i, c := range test.Cases {
				script, err := parser.Parse(lexer.New([]byte(c.Input)))
				if err != nil {
					t.Fatalf("#%d: parser error.\nError: %s", i, err)
				}

				program := generator.Generate(script)
				formattedProgram, gofmtErr, err := gofmt(program.String())
				if err != nil {
					t.Fatalf("#%d: error when trying to format the generated program.\nError: %s.\nStderr: %s", i, err, gofmtErr)
				}

				formattedTestOutput, gofmtErr, err := gofmt(c.Output)
				if err != nil {
					t.Fatalf("#%d: error when trying to format the test expected output.\nError: %s.\nStderr: %s", i, err, gofmtErr)
				}

				if formattedProgram != formattedTestOutput {
					t.Fatalf("#%d: The generated program doesn't match the expected output.\n Program:\n%s", i, dump(formattedProgram))
				}
			}
		})
	}
}

func gofmt(s string) (gofmtOut string, gofmtErr string, err error) {
	gofmt := exec.Command("gofmt")
	gofmt.Stdin = strings.NewReader(s)

	var stdout bytes.Buffer
	gofmt.Stdout = &stdout

	var stderr bytes.Buffer
	gofmt.Stderr = &stderr

	err = gofmt.Run()
	gofmtOut = stdout.String()
	gofmtErr = stderr.String()

	return
}

// func generate(s string) ir.Program {

// }

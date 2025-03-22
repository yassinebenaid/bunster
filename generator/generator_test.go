package generator_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/bunster/pkg/diff"
	"github.com/yassinebenaid/bunster/pkg/dottest"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestGenerator(t *testing.T) {
	testFiles, err := filepath.Glob("./tests/*.test")
	if err != nil {
		t.Fatalf("Failed to `Glob` test files, %v", err)
	}

	for _, testFile := range testFiles {
		t.Run(testFile, func(t *testing.T) {
			testContent, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("Failed to open test file, %v", err)
			}

			tests, err := dottest.Parse(string(testContent))
			if err != nil {
				t.Fatal(err)
			}

			for i, test := range tests {
				script, err := parser.Parse(lexer.New([]rune(test.Input)))
				if err != nil {
					t.Fatalf("\nTest: %sError: %s", dump(test.Label), dump(err.Error()))
				}

				if err := analyser.Analyse(script, true); err != nil {
					t.Fatalf("\nTest: %sError: %s", dump(test.Label), dump(err.Error()))
				}

				program := generator.Generate(script)
				formattedProgram, gofmtErr, err := gofmt(program.String())
				if err != nil {
					t.Fatalf("gofmt error in generated program -- #%sError: %sStderr: %s", dump(i), dump(err.Error()), dump(gofmtErr))
				}

				formattedTestOutput, gofmtErr, err := gofmt(test.Output)
				if err != nil {
					t.Fatalf("gofmt error in test program -- #%sError: %sStderr: %s", dump(i), dump(err.Error()), dump(gofmtErr))
				}

				if formattedProgram != formattedTestOutput {
					t.Fatalf(
						"\n#%d: The generated program doesn't match the expected output.\nTest: %s\n Program:\n%s",
						i, test.Label, diff.Diff(formattedTestOutput, formattedProgram),
					)
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

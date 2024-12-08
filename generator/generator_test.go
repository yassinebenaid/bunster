package generator_test

import (
	"bytes"
	"fmt"
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
					t.Fatalf(
						"#%d: The generated program doesn't match the expected output.\n Program:\n%s",
						i, diffStrings(formattedTestOutput, formattedProgram),
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

// ANSI color codes
const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

// diffStrings compares two strings and returns a formatted diff output
func diffStrings(original, modified string) string {
	// Split strings into lines
	originalLines := strings.Split(original, "\n")
	modifiedLines := strings.Split(modified, "\n")

	var diffOutput []string

	// Track the maximum length of lines for alignment
	maxLen := max(len(originalLines), len(modifiedLines))

	// Compare lines
	for i := 0; i < maxLen; i++ {
		var originalLine, modifiedLine string

		// Get lines if they exist
		if i < len(originalLines) {
			originalLine = originalLines[i]
		}
		if i < len(modifiedLines) {
			modifiedLine = modifiedLines[i]
		}

		// Compare lines
		if originalLine != modifiedLine {
			if i >= len(originalLines) {
				// New line added
				diffOutput = append(diffOutput, fmt.Sprintf("%s+ %s%s", colorGreen, modifiedLine, colorReset))
			} else if i >= len(modifiedLines) {
				// Line deleted
				diffOutput = append(diffOutput, fmt.Sprintf("%s- %s%s", colorRed, originalLine, colorReset))
			} else {
				// Line modified
				diffOutput = append(diffOutput, fmt.Sprintf("%s- %s%s", colorRed, originalLine, colorReset))
				diffOutput = append(diffOutput, fmt.Sprintf("%s+ %s%s", colorGreen, modifiedLine, colorReset))
			}
		} else {
			// Unchanged line
			diffOutput = append(diffOutput, fmt.Sprintf("  %s", originalLine))
		}
	}

	return strings.Join(diffOutput, "\n")
}

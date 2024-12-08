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

// diffStrings generates a git-like diff between two strings
func diffStrings(original, modified string) string {
	originalLines := strings.Split(original, "\n")
	modifiedLines := strings.Split(modified, "\n")

	// Compute the LCS and the diffs
	ops := computeDiff(originalLines, modifiedLines)

	var diffOutput []string

	for _, op := range ops {
		switch op.Type {
		case "unchanged":
			diffOutput = append(diffOutput, fmt.Sprintf("  %s", op.Line))
		case "delete":
			diffOutput = append(diffOutput, fmt.Sprintf("%s- %s%s", colorRed, op.Line, colorReset))
		case "add":
			diffOutput = append(diffOutput, fmt.Sprintf("%s+ %s%s", colorGreen, op.Line, colorReset))
		}
	}

	return strings.Join(diffOutput, "\n")
}

// DiffOperation represents a single diff operation
type DiffOperation struct {
	Type string // "add", "delete", or "unchanged"
	Line string
}

// computeDiff calculates the diff between two slices of lines using the LCS algorithm
func computeDiff(original, modified []string) []DiffOperation {
	m, n := len(original), len(modified)
	lcs := make([][]int, m+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}

	// Build the LCS table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if original[i-1] == modified[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	// Backtrack to determine the diff
	var ops []DiffOperation
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && original[i-1] == modified[j-1] {
			ops = append([]DiffOperation{{Type: "unchanged", Line: original[i-1]}}, ops...)
			i--
			j--
		} else if j > 0 && (i == 0 || lcs[i][j-1] >= lcs[i-1][j]) {
			ops = append([]DiffOperation{{Type: "add", Line: modified[j-1]}}, ops...)
			j--
		} else {
			ops = append([]DiffOperation{{Type: "delete", Line: original[i-1]}}, ops...)
			i--
		}
	}

	return ops
}

package bunster_test

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/yassinebenaid/bunster"
	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/godump"
	"gopkg.in/yaml.v3"
)

type Test struct {
	Name   string `yaml:"name"`
	Script string `yaml:"script"`
	Expect struct {
		Stdout   string `yaml:"stdout"`
		Stderr   string `yaml:"stderr"`
		ExitCode int    `yaml:"exit_code"`
	} `yaml:"expect"`
}

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestBunster(t *testing.T) {
	testFiles, err := filepath.Glob("./tests/*.yml")
	if err != nil {
		t.Fatalf("Failed to `Glob` test files, %v", err)
	}

	for _, testFile := range testFiles {
		t.Run(testFile, func(t *testing.T) {
			testContent, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("Failed to open test file, %v", err)
			}

			var tests []Test

			if err := yaml.Unmarshal(testContent, &tests); err != nil {
				t.Fatalf("Failed to parse test file, %v", err)
			}

			for i, test := range tests {
				binary, err := buildBinary([]byte(test.Script))
				if err != nil {
					t.Fatalf("\nTest(#%d): %sBuild Error: %s", i, dump(test.Name), dump(err.Error()))
				}

				var stdout, stderr bytes.Buffer

				cmd := exec.Command(binary)
				cmd.Stdout = &stdout
				cmd.Stderr = &stderr
				if err := cmd.Run(); err != nil {
					_, ok := err.(*exec.ExitError)
					if !ok {
						t.Fatalf("\nTest(#%d): %sRuntime Error: %s", i, dump(test.Name), dump(err.Error()))
					}
				}

				if test.Expect.ExitCode != cmd.ProcessState.ExitCode() {
					t.Fatalf("\nTest(#%d): %sExpected exit code of '%d', got '%d'",
						i, dump(test.Name), test.Expect.ExitCode, cmd.ProcessState.ExitCode())
				}

				if test.Expect.Stdout != stdout.String() {
					t.Fatalf("\nTest(#%d): %sExpected `STDOUT` does not match actual value\ndiff:\n%s",
						i, dump(test.Name), diffStrings(test.Expect.Stdout, stdout.String()))
				}

				if test.Expect.Stderr != stderr.String() {
					t.Fatalf("\nTest(#%d): %sExpected `STDERR` does not match actual value\ndiff:\n%s",
						i, dump(test.Name), diffStrings(test.Expect.Stderr, stderr.String()))
				}
			}
		})
	}
}

func buildBinary(s []byte) (string, error) {
	script, err := parser.Parse(lexer.New(s))
	if err != nil {
		return "", err
	}

	if err := analyser.Analyse(script); err != nil {
		return "", err
	}

	program := generator.Generate(script)

	wd, err := os.MkdirTemp("", "bunster-build-*")
	if err != nil {
		return "", err
	}

	err = os.WriteFile(path.Join(wd, "program.go"), []byte(program.String()), 0600)
	if err != nil {
		return "", err
	}

	if err := cloneRuntime(wd); err != nil {
		return "", err
	}

	if err := cloneStubs(wd); err != nil {
		return "", err
	}

	gocmd := exec.Command("go", "build", "-o", "build.bin")
	gocmd.Stdin = os.Stdin
	gocmd.Stdout = os.Stdout
	gocmd.Stderr = os.Stderr
	gocmd.Dir = wd
	if err := gocmd.Run(); err != nil {
		return "", err
	}

	return path.Join(wd, "build.bin"), nil
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

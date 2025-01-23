package bunster_test

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/yassinebenaid/bunster"
	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/bunster/pkg/diff"
	"github.com/yassinebenaid/godump"
	"gopkg.in/yaml.v3"
)

type Test struct {
	Cases []struct {
		Name   string            `yaml:"name"`
		Stdin  string            `yaml:"stdin"`
		RunsOn string            `yaml:"runs_on"`
		Env    []string          `yaml:"env"`
		Args   []string          `yaml:"args"`
		Files  map[string]string `yaml:"files"`
		Script string            `yaml:"script"`
		Expect struct {
			Stdout   string            `yaml:"stdout"`
			Stderr   string            `yaml:"stderr"`
			ExitCode int               `yaml:"exit_code"`
			Files    map[string]string `yaml:"files"`
		} `yaml:"expect"`
	}
}

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestBunster(t *testing.T) {
	filter := os.Getenv("FILTER")

	buildWorkdir, err := prepareBuildAssets()
	if err != nil {
		t.Fatalf("Failed to prepare the build assets, %v", err)
	}

	testFiles, err := filepath.Glob("./tests/*.yml")
	if err != nil {
		t.Fatalf("Failed to `Glob` test files, %v", err)
	}

	var testsHasRan int // number of tests that has ran

	for _, testFile := range testFiles {
		t.Run(testFile, func(t *testing.T) {
			testContent, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("Failed to open test file, %v", err)
			}

			var test Test

			if err := yaml.Unmarshal(testContent, &test); err != nil {
				t.Fatalf("Failed to parse test file, %v", err)
			}

			for i, testCase := range test.Cases {
				if !strings.Contains(testCase.Name, filter) {
					// we support filtering, someone would want to run specific tests.
					continue
				}
				if testCase.RunsOn != "" && testCase.RunsOn != runtime.GOOS {
					// some tests only run on spesific platforms.
					continue
				}
				testsHasRan++

				workdir, err := setupWorkdir()
				if err != nil {
					t.Fatalf("Failed to setup test workdir, %v", err)
				}

				binary, err := buildBinary(buildWorkdir, []byte(testCase.Script))
				if err != nil {
					t.Fatalf("\nTest(#%d): %sBuild Error: %s", i, dump(testCase.Name), dump(err.Error()))
				}

				for filename, content := range testCase.Files {
					if err := os.WriteFile(path.Join(workdir, filename), []byte(content), 0600); err != nil {
						t.Fatalf("\nTest(#%d): %sFailed to write file %q, %v", i, dump(testCase.Name), filename, err)
					}
				}

				var stdout, stderr bytes.Buffer

				cmd := exec.Command(binary, testCase.Args...)
				cmd.Stdin = strings.NewReader(testCase.Stdin)
				cmd.Stdout = &stdout
				cmd.Stderr = &stderr
				cmd.Dir = workdir
				cmd.Env = append(os.Environ(), testCase.Env...)
				if err := cmd.Run(); err != nil {
					_, ok := err.(*exec.ExitError)
					if !ok {
						t.Fatalf("\nTest(#%d): %sRuntime Error: %s", i, dump(testCase.Name), dump(err.Error()))
					}
				}

				if testCase.Expect.ExitCode != cmd.ProcessState.ExitCode() {
					t.Fatalf("\nTest(#%d): %sExpected exit code of '%d', got '%d'",
						i, dump(testCase.Name), testCase.Expect.ExitCode, cmd.ProcessState.ExitCode())
				}

				if testCase.Expect.Stdout != stdout.String() {
					t.Fatalf("\nTest(#%d): %sExpected `STDOUT` does not match actual value\ndiff:\n%s",
						i, dump(testCase.Name), diff.DiffBG(testCase.Expect.Stdout, stdout.String()))
				}

				if testCase.Expect.Stderr != stderr.String() {
					t.Fatalf("\nTest(#%d): %sExpected `STDERR` does not match actual value\ndiff:\n%s",
						i, dump(testCase.Name), diff.DiffBG(testCase.Expect.Stderr, stderr.String()))
				}

				files, err := filepath.Glob(workdir + "/*")
				if err != nil {
					t.Fatalf("\nTest(#%d): %sFailed to glob the working directory, %v", i, dump(testCase.Name), err)
				}
				if testCase.Expect.Files != nil && len(files) != len(testCase.Expect.Files) {
					t.Fatalf("\nTest(#%d): %sExpected files in working directory does not match actual count, files count is: %d, expected files count: %d",
						i, dump(testCase.Name), len(files), len(testCase.Expect.Files))
				}

				for filename, expectedContent := range testCase.Expect.Files {
					content, err := os.ReadFile(path.Join(workdir, filename))
					if err != nil {
						t.Fatalf("\nTest(#%d): %sFailed to read file %q, %v", i, dump(testCase.Name), filename, err)
					}

					if string(content) != expectedContent {
						t.Fatalf("\nTest(#%d): %sExpected file content doesn't match actual content\nfile: %s\ndiff:\n%s",
							i, dump(testCase.Name), filename, diff.DiffBG(expectedContent, string(content)))
					}
				}
			}
		})
	}

	if testsHasRan == 0 {
		t.Fatalf("\nNo tests has ran.")
	}
}

func buildBinary(workdir string, s []byte) (string, error) {
	script, err := parser.Parse(lexer.New(s))
	if err != nil {
		return "", err
	}

	if err := analyser.Analyse(script); err != nil {
		return "", err
	}

	program := generator.Generate(script)

	err = os.WriteFile(path.Join(workdir, "program.go"), []byte(program.String()), 0600)
	if err != nil {
		return "", err
	}

	gocmd := exec.Command("go", "build", "-o", "build.bin")
	gocmd.Stdin = os.Stdin
	gocmd.Stdout = os.Stdout
	gocmd.Stderr = os.Stderr
	gocmd.Dir = workdir
	if err := gocmd.Run(); err != nil {
		return "", err
	}

	return path.Join(workdir, "build.bin"), nil
}

func prepareBuildAssets() (string, error) {
	wd := path.Join(os.TempDir(), "bunster-build")
	if err := os.RemoveAll(wd); err != nil {
		return "", err
	}

	if err := os.MkdirAll(wd, 0700); err != nil {
		return "", err
	}

	if err := cloneRuntime(wd); err != nil {
		return "", err
	}

	if err := cloneStubs(wd); err != nil {
		return "", err
	}

	return wd, nil
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

func setupWorkdir() (string, error) {
	wd := path.Join(os.TempDir(), "bunster-testing-workdir")
	if err := os.RemoveAll(wd); err != nil {
		return "", err
	}

	if err := os.MkdirAll(wd, 0700); err != nil {
		return "", err
	}
	return wd, nil
}

package bunster

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/ir"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
)

//go:embed VERSION
var Version string

//go:embed runtime
var runtimeFS embed.FS

//go:embed stubs/go.mod.stub
var gomod []byte

//go:embed stubs/main.go.stub
var mainGo []byte

func Generate(cwd, workdir string, s []byte) error {
	program, err := compile(s)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(workdir, "program.go"), []byte(program.String()), 0600)
	if err != nil {
		return err
	}

	if err := cloneRuntime(workdir); err != nil {
		return err
	}

	if err := cloneStubs(workdir); err != nil {
		return err
	}

	if err := cloneEmbeddedFiles(cwd, workdir, program.Embeds); err != nil {
		return err
	}

	return nil
}

func compile(s []byte) (*ir.Program, error) {
	script, err := parser.Parse(lexer.New(s))
	if err != nil {
		return nil, err
	}

	if err := analyser.Analyse(script); err != nil {
		return nil, err
	}

	program := generator.Generate(script)

	return &program, nil
}

func cloneRuntime(dst string) error {
	return fs.WalkDir(runtimeFS, "runtime", func(dpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return os.MkdirAll(path.Join(dst, dpath), 0766)
		}

		if strings.HasSuffix(dpath, "_test.go") {
			return nil
		}

		content, err := runtimeFS.ReadFile(dpath)
		if err != nil {
			return err
		}

		return os.WriteFile(path.Join(dst, dpath), content, 0600)
	})
}

func cloneStubs(dst string) error {
	if err := os.WriteFile(path.Join(dst, "main.go"), mainGo, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(dst, "go.mod"), gomod, 0600); err != nil {
		return err
	}

	return nil
}

func cloneEmbeddedFiles(cwd, dst string, files []string) error {
	for _, file := range files {
		srcPath, dstPath := path.Join(cwd, file), path.Join(dst, file)

		info, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			if err := copyDir(srcPath, dst); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(_path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.MkdirAll(path.Join(dst, _path), 0766)
		}

		return copyFile(_path, path.Join(dst, _path))
	})
}

func copyFile(src, dst string) error {
	srcf, err := os.OpenFile(src, os.O_RDONLY, 000)
	if err != nil {
		return err
	}
	defer srcf.Close()

	dstf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer dstf.Close()

	if _, err := io.Copy(dstf, srcf); err != nil {
		return err
	}

	return nil
}

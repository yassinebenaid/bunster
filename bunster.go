package bunster

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
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
var RuntimeFS embed.FS

//go:embed stubs/go.mod.stub
var Gomod []byte

//go:embed stubs/main.go.stub
var MainGo []byte

func Generate(workdir string, s []rune) error {
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

	if err := cloneEmbeddedFiles(workdir, program.Embeds); err != nil {
		return err
	}

	return nil
}

func compile(s []rune) (*ir.Program, error) {
	script, err := parser.Parse(lexer.New(s))
	if err != nil {
		return nil, err
	}

	if err := analyser.Analyse(script, true); err != nil {
		return nil, err
	}

	program := generator.Generate(script)

	return &program, nil
}

func cloneRuntime(dst string) error {
	return fs.WalkDir(RuntimeFS, "runtime", func(dpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return os.MkdirAll(path.Join(dst, dpath), 0766)
		}

		if strings.HasSuffix(dpath, "_test.go") {
			return nil
		}

		content, err := RuntimeFS.ReadFile(dpath)
		if err != nil {
			return err
		}

		return os.WriteFile(path.Join(dst, dpath), content, 0600)
	})
}

func cloneStubs(dst string) error {
	if err := os.WriteFile(path.Join(dst, "main.go"), MainGo, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(dst, "go.mod"), Gomod, 0600); err != nil {
		return err
	}

	return nil
}

func cloneEmbeddedFiles(dst string, files []string) error {
	if err := os.MkdirAll(path.Join(dst, ir.EmbedDirectory), 0766); err != nil {
		return err
	}

	for _, file := range files {
		srcPath, dstPath := file, path.Join(dst, ir.EmbedDirectory, file)

		info, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			if err := copyDir(srcPath, path.Join(dst, ir.EmbedDirectory)); err != nil {
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

var specialPathRegex = regexp.MustCompile(`^(.*\.git.*)|(.*go\.mod.*)$`)

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(_path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if specialPathRegex.MatchString(_path) {
			return nil
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

	dir, _ := path.Split(dst)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

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

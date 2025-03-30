package builder

import (
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"

	"github.com/yassinebenaid/bunster"
	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/ir"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
)

type Builder struct {
	Workdir    string
	MainScript string
	Builddir   string
	OutputFile string
	Gofmt      bool
}

func (b *Builder) Build() (err error) {
	if err := b.Generate(); err != nil {
		return err
	}

	gocmd := exec.Command("go", "build", "-o", b.OutputFile) //nolint:gosec
	gocmd.Stdin = os.Stdin
	gocmd.Stdout = os.Stdout
	gocmd.Stderr = os.Stderr
	gocmd.Dir = b.Builddir
	if err := gocmd.Run(); err != nil {
		return err
	}

	return nil
}

func (b *Builder) Generate() (err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err = os.Chdir(b.Workdir); err != nil {
		return err
	}
	defer func() {
		if e := os.Chdir(cwd); e != nil {
			err = e
		}
	}()

	if err := b.prepare(); err != nil {
		return err
	}

	v, err := os.ReadFile(b.MainScript)
	if err != nil {
		return err
	}

	mainSh, err := parser.Parse(lexer.New([]rune(string(v))))
	if err != nil {
		return err
	}

	if err := analyser.Analyse(mainSh, true); err != nil {
		return err
	}

	config, err := b.loadConfig()
	if err != nil {
		return err
	}

	module, err := b.globModule(config)
	if err != nil {
		return err
	}

	for _, f := range module.Tree {
		v, err := os.ReadFile(f)
		if err != nil {
			return err
		}

		script, err := parser.Parse(lexer.New([]rune(string(v))))
		if err != nil {
			return err
		}

		if err := analyser.Analyse(script, false); err != nil {
			return err
		}

		mainSh = append(script, mainSh...)
	}

	for _, submodule := range module.Require {
		for _, f := range submodule.Tree {
			v, err := os.ReadFile(f)
			if err != nil {
				return err
			}

			script, err := parser.Parse(lexer.New([]rune(string(v))))
			if err != nil {
				return err
			}

			if err := analyser.Analyse(script, false); err != nil {
				return err
			}

			mainSh = append(script, mainSh...)
		}
	}

	program := generator.Generate(mainSh)

	err = os.WriteFile(path.Join(b.Builddir, "program.go"), []byte(program.String()), 0600)
	if err != nil {
		return err
	}

	if err := b.writeRuntime(); err != nil {
		return err
	}

	if err := b.writeStubs(); err != nil {
		return err
	}

	if err := b.writeEmbeddedFiles(program.Embeds); err != nil {
		return err
	}

	if b.Gofmt {
		// we ignore the error, because this is just an optional step that shouldn't stop us from building the binary
		_ = exec.Command("gofmt", "-w", b.Builddir).Run() //nolint:gosec
	}

	return nil
}

func (b *Builder) prepare() error {
	if err := os.RemoveAll(b.Builddir); err != nil {
		return err
	}

	if err := os.MkdirAll(b.Builddir, 0700); err != nil {
		return err
	}

	if err := os.MkdirAll(path.Join(b.Builddir, ir.EmbedDirectory), 0700); err != nil {
		return err
	}

	if b.MainScript == "" {
		b.MainScript = "main.sh"
	}

	return nil
}

func (b *Builder) writeRuntime() error {
	return fs.WalkDir(bunster.RuntimeFS, "runtime", func(dpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		dst := path.Join(b.Builddir, dpath)

		if d.IsDir() {
			return os.MkdirAll(dst, 0766)
		}

		content, err := bunster.RuntimeFS.ReadFile(dpath)
		if err != nil {
			return err
		}

		return os.WriteFile(dst, content, 0600)
	})
}

func (b *Builder) writeStubs() error {
	if err := os.WriteFile(path.Join(b.Builddir, "main.go"), bunster.MainGo, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(b.Builddir, "go.mod"), bunster.Gomod, 0600); err != nil {
		return err
	}

	return nil
}

func (b *Builder) writeEmbeddedFiles(files []string) error {
	for _, file := range files {
		src, dst := file, path.Join(b.Builddir, ir.EmbedDirectory, file)

		info, err := os.Stat(src)
		if err != nil {
			return err
		}

		if info.IsDir() {
			if err := copyDir(src, path.Join(b.Builddir, ir.EmbedDirectory)); err != nil {
				return err
			}
		} else {
			if err := copyFile(src, dst); err != nil {
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

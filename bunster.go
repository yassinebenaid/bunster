package bunster

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//go:embed runtime
var RuntimeFS embed.FS

//go:embed stubs/go.mod.stub
var GoModStub []byte

//go:embed stubs/main.go.stub
var MainGoStub []byte

//go:embed VERSION
var Version string

func CloneRuntime(dst string) error {
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

func CloneStubs(dst string) error {
	if err := os.WriteFile(path.Join(dst, "main.go"), MainGoStub, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(dst, "go.mod"), GoModStub, 0600); err != nil {
		return err
	}

	return nil
}

func CloneEmbeddedFiles(dst string, files []string) error {
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return err
		}

		if info.IsDir() {
			if err := copyDir(file, dst); err != nil {
				return err
			}
		} else {
			if err := copyFile(file, path.Join(dst, file)); err != nil {
				return err
			}
		}
	}

	return nil
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

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(_path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return os.MkdirAll(path.Join(dst, _path), 0766)
		}

		return copyFile(_path, path.Join(dst, _path))
	})
}

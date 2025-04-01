package builder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Dependency struct {
	From string `yaml:"from"`
	Rev  string `yaml:"rev"`
}

type Config struct {
	Module  string       `yaml:"module"`
	Require []Dependency `yaml:"require"`
}

type Module struct {
	Module  string
	Require []Module
	Tree    []string
}

func (b *Builder) globModule(c *Config) (*Module, error) {
	files, err := filepath.Glob("*.sh")
	if err != nil {
		return nil, err
	}

	var module Module

	for _, file := range files {
		if file != b.MainScript {
			module.Tree = append(module.Tree, file)
		}
	}

	for _, dep := range c.Require {
		var submodule Module
		submodule.Module = dep.From + "/" + dep.Rev
		module.Require = append(module.Require, submodule)
		files, err := filepath.Glob(filepath.Join(os.Getenv("HOME"), ".bunster", "pkg", dep.From, dep.Rev, "*.sh"))
		if err != nil {
			return nil, err
		}
		submodule.Tree = files

		module.Require = append(module.Require, submodule)
	}

	return &module, nil
}

const configFile = "bunster.test.yml"

func (b *Builder) loadConfig() (*Config, error) {
	var c Config

	content, err := os.ReadFile(filepath.Join(b.Workdir, configFile))
	if err != nil {
		if os.IsNotExist(err) {
			return &c, nil
		}
		return nil, err
	}

	if err := yaml.Unmarshal(content, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (b *Builder) writeConfig(config *Config) error {
	f, err := os.OpenFile(configFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	encoder := yaml.NewEncoder(f)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return encoder.Close()
}

func (b *Builder) ResolveDeps(packages []string) (err error) {
	config, err := b.loadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, p := range packages {
		pack := strings.SplitN(p, "@", 2)
		name, rev := pack[0], pack[1]

		if err := b.getPackage(name, rev); err != nil {
			return err
		}

		config.Require = append(config.Require, Dependency{
			From: name,
			Rev:  rev,
		})
	}

	return b.writeConfig(config)
}

func (b *Builder) getPackage(path, rev string) (err error) {
	tmpDir, pkgDir := filepath.Join(b.Home, "tmp"), filepath.Join(b.Home, "pkg", path, rev)

	if err := os.RemoveAll(tmpDir); err != nil {
		return err
	}
	if err := os.MkdirAll(tmpDir, 0700); err != nil {
		return err
	}

	if err := os.RemoveAll(pkgDir); err != nil {
		return err
	}
	if err := os.MkdirAll(pkgDir, 0700); err != nil {
		return err
	}

	var stderr bytes.Buffer
	sh := exec.Command("sh")
	sh.Stderr = &stderr
	sh.Dir = filepath.Join(b.Home, "tmp")
	sh.Stdin = strings.NewReader(fmt.Sprintf(`
		git init && \
			git fetch --dept=1 "https://%s" "%s" && \
			git checkout "%s" && \
			rm -rf .git && \
			cp -r . %s
	`, path, rev, rev, pkgDir))

	if err := sh.Run(); err != nil {
		return fmt.Errorf(`failed to resolve package "%s", error: "%s", stderr: "%s"`, path, err, stderr.String())
	}
	return nil
}

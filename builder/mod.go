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

type Config struct {
	Module  string            `yaml:"module"`
	Require map[string]string `yaml:"require"`
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

	for from, rev := range c.Require {
		var submodule Module
		submodule.Module = from + "/" + rev
		module.Require = append(module.Require, submodule)
		files, err := filepath.Glob(filepath.Join(os.Getenv("HOME"), ".bunster", "pkg", from, rev, "*.sh"))
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
	var c = Config{
		Require: map[string]string{},
	}

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
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFile, out, 0600); err != nil {
		return err
	}

	return nil
}

func (b *Builder) ResolveDeps(packages []string) (err error) {
	config, err := b.loadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(packages) == 0 {
		for dep, rev := range config.Require {
			packages = append(packages, dep+"@"+rev)
		}
	}

	for _, p := range packages {
		pack := strings.SplitN(p, "@", 2)
		name, rev := pack[0], pack[1]

		if err := b.getPackage(name, rev); err != nil {
			return err
		}

		config.Require[name] = rev
	}

	return b.writeConfig(config)
}

func (b *Builder) getPackage(path, rev string) (err error) {
	pkgDir := filepath.Join(b.Home, "pkg", path, rev)

	if err := os.RemoveAll(pkgDir); err != nil {
		return err
	}
	if err := os.MkdirAll(pkgDir, 0700); err != nil {
		return err
	}

	if err := _exec(pkgDir, "git", "init"); err != nil {
		return err
	}
	if err := _exec(pkgDir, "git", "fetch", "--dept=1", "https://"+path, rev); err != nil {
		return fmt.Errorf("failed to resolve package %q, either path or version are invalid", path)
	}
	if err := _exec(pkgDir, "git", "checkout", "FETCH_HEAD"); err != nil {
		return fmt.Errorf("failed to resolve package %q, unknown revision %q", path, rev)
	}
	if err := _exec(pkgDir, "rm", "-rf", ".git"); err != nil {
		return err
	}

	return nil
}

func _exec(dir string, args ...string) error {
	var stderr bytes.Buffer
	sh := exec.Command(args[0], args[1:]...)
	sh.Stderr = &stderr
	sh.Dir = dir

	if err := sh.Run(); err != nil {
		return fmt.Errorf(`exec error(%s): "%s", stderr: "%s"`, args, err, stderr.String())
	}

	return nil
}

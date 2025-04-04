package builder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Require map[string]string `yaml:"require"`
}

type Module struct {
	Path    string
	Version string
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

	for path, version := range c.Require {
		var submodule Module
		submodule.Path = path
		submodule.Version = version
		module.Require = append(module.Require, submodule)
		files, err := filepath.Glob(filepath.Join(os.Getenv("HOME"), ".bunster", "pkg", path, version, "*.sh"))
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

func (b *Builder) ResolveDeps(packages []string, missing bool) (err error) {
	config, err := b.loadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if missing {
		for module, version := range config.Require {
			packages = append(packages, module+"@"+version)
		}
	}

	for _, p := range packages {
		query, err := parseQuery(p)
		if err != nil {
			return err
		}

		if b.checkPackage(query) {
			continue
		}

		if err := b.getPackage(query); err != nil {
			return err
		}

		config.Require[query.module] = query.commit
	}

	return b.writeConfig(config)
}

func (b *Builder) getPackage(q query) (err error) {
	pkgDir := filepath.Join(b.Home, "pkg", q.module, q.commit)

	if err := os.RemoveAll(pkgDir); err != nil {
		return err
	}
	if err := os.MkdirAll(pkgDir, 0700); err != nil {
		return err
	}

	if err := _exec(pkgDir, "git", "init"); err != nil {
		return err
	}
	if err := _exec(pkgDir, "git", "fetch", "--dept=1", "https://"+q.module, q.commit); err != nil {
		return fmt.Errorf("failed to resolve package %q, either path or version are invalid", q.module)
	}
	if err := _exec(pkgDir, "git", "checkout", "FETCH_HEAD"); err != nil {
		return fmt.Errorf("failed to resolve package %q, unknown revision %q", q.module, q.commit)
	}
	if err := _exec(pkgDir, "rm", "-rf", ".git"); err != nil {
		return err
	}

	if err := b.lockPackage(q); err != nil {
		return err
	}

	return nil
}

func (b *Builder) lockPackage(q query) (err error) {
	lockFile := filepath.Join(b.Home, "pkg", q.module, q.commit, ".bunster.lock")

	return os.WriteFile(lockFile, nil, 0600)
}

func (b *Builder) checkPackage(q query) bool {
	lockFile := filepath.Join(b.Home, "pkg", q.module, q.commit, ".bunster.lock")

	_, err := os.Stat(lockFile)
	return err == nil
}

func _exec(dir string, args ...string) error {
	var stderr bytes.Buffer
	sh := exec.Command(args[0], args[1:]...) //nolint:gosec
	sh.Stderr = &stderr
	sh.Dir = dir

	if err := sh.Run(); err != nil {
		return fmt.Errorf(`exec error(%s): "%s", stderr: "%s"`, args, err, stderr.String())
	}

	return nil
}

package builder

import (
	"os"
	"path/filepath"

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

package runtime

type parameter struct {
	value any
}

func (p parameter) String() string {
	switch v := p.value.(type) {
	case string:
		return v
	case []string:
		if len(v) == 0 {
			return ""
		}
		return v[0]
	}
	return ""
}

func (p parameter) AtIndex(index int) string {
	switch v := p.value.(type) {
	case []string:
		if len(v) == 0 {
			return ""
		}
		if index >= 0 && index <= len(v)-1 {
			return v[index]
		}
	}
	return ""
}

func (p parameter) HasIndex(index int) bool {
	switch v := p.value.(type) {
	case []string:
		if len(v) == 0 {
			return false
		}
		if index >= 0 && index <= len(v)-1 {
			return true
		}
	}
	return false
}

func (shell *Shell) ReadVar(name string) string {
	p := shell.readVar(name)

	return p.String()
}

func (shell *Shell) ReadArrayVar(name string, index int) string {
	p := shell.readVar(name)

	return p.AtIndex(index)
}

func (shell *Shell) readVar(name string) parameter {
	if value, ok := shell.getLocalVar(name); ok {
		return value
	}
	if value, ok := shell.vars.get(name); ok {
		return value
	}
	if value, ok := shell.env.get(name); ok {
		return value
	}
	if shell.parent != nil {
		return shell.parent.readVar(name)
	}
	return parameter{}
}

func (shell *Shell) VarIsSet(name string) bool {
	if _, ok := shell.getLocalVar(name); ok {
		return true
	}
	if _, ok := shell.vars.get(name); ok {
		return true
	}
	if _, ok := shell.env.get(name); ok {
		return true
	}
	return false
}

func (shell *Shell) VarIndexIsSet(name string, index int) bool {
	if v, ok := shell.getLocalVar(name); ok {
		return v.HasIndex(index)
	}
	if v, ok := shell.vars.get(name); ok {
		return v.HasIndex(index)
	}
	if v, ok := shell.env.get(name); ok {
		return v.HasIndex(index)
	}
	return false
}

func (shell *Shell) setLocalVar(name string, value any) bool {
	if _, ok := shell.localVars.get(name); ok {
		shell.localVars.set(name, parameter{value: value})
		return true
	}
	if shell.parent != nil {
		return shell.parent.setLocalVar(name, value)
	}
	return false
}

func (shell *Shell) getLocalVar(name string) (parameter, bool) {
	if value, ok := shell.localVars.get(name); ok {
		return value, true
	}
	if shell.parent != nil {
		return shell.parent.getLocalVar(name)
	}
	return parameter{}, false
}

func (shell *Shell) SetVar(name string, value any) {
	if !shell.setLocalVar(name, value) {
		shell.vars.set(name, parameter{value: value})
	}
}

func (shell *Shell) SetLocalVar(name string, value string) {
	shell.localVars.set(name, parameter{value: value})
}

func (shell *Shell) SetExportVar(name string, value string) {
	shell.exportedVars.set(name, struct{}{})
	shell.vars.set(name, parameter{value: value})
}

func (shell *Shell) MarkVarAsExported(name string) {
	shell.exportedVars.set(name, struct{}{})
}

func (shell *Shell) Unset(vars_only bool, names ...string) {
	for _, name := range names {
		if shell.localVars.forget(name) ||
			shell.vars.forget(name) ||
			shell.env.forget(name) {
			continue
		}

		if !vars_only && shell.functions.forget(name) {
			continue
		}

		if shell.parent != nil {
			shell.parent.Unset(vars_only, name)
		}
	}
}

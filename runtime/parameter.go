package runtime

import (
	"strconv"
	"strings"
)

type parameter struct {
	value any
}

func (p *parameter) String() string {
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

func (p *parameter) atIndex(index int) string {
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

func (p *parameter) setIndex(index int, value string) {
	v, ok := p.value.([]string)
	if !ok {
		v = make([]string, index+1)
	}

	if index >= len(v) {
		v = append(v, make([]string, index)...)
	}

	if index >= 0 {
		v[index] = value
	}

	p.value = v
}

func (p *parameter) hasIndex(index int) bool {
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
	p, ok := shell.readVar(name)
	if !ok {
		return ""
	}
	return p.String()
}

func (shell *Shell) ReadArrayVar(name string, index int) string {
	p, ok := shell.readVar(name)
	if !ok {
		return ""
	}
	return p.atIndex(index)
}

func (shell *Shell) readVar(name string) (*parameter, bool) {
	if value, ok := shell.getLocalVar(name); ok {
		return value, true
	}
	if value, ok := shell.vars.get(name); ok {
		return value, true
	}
	if value, ok := shell.env.get(name); ok {
		return value, true
	}
	if shell.parent != nil {
		return shell.parent.readVar(name)
	}
	return nil, false
}

func (shell *Shell) VarIsSet(name string) bool {
	_, ok := shell.readVar(name)
	return ok
}

func (shell *Shell) VarIndexIsSet(name string, index int) bool {
	p, ok := shell.readVar(name)
	if !ok {
		return false
	}
	return p.hasIndex(index)
}

func (shell *Shell) setLocalVar(name string, value any) bool {
	if _, ok := shell.localVars.get(name); ok {
		shell.localVars.set(name, &parameter{value: value})
		return true
	}
	if shell.parent != nil {
		return shell.parent.setLocalVar(name, value)
	}
	return false
}

func (shell *Shell) getLocalVar(name string) (*parameter, bool) {
	if value, ok := shell.localVars.get(name); ok {
		return value, true
	}
	if shell.parent != nil {
		return shell.parent.getLocalVar(name)
	}
	return nil, false
}

func (shell *Shell) SetVar(name string, value any) {
	if !shell.setLocalVar(name, value) {
		shell.vars.set(name, &parameter{value: value})
	}
}

func (shell *Shell) SetArrayVar(name string, index int, value string) {
	p, ok := shell.readVar(name)
	if ok {
		p.setIndex(index, value)
		return
	}
	v := make([]string, index+1)
	v[index] = value
	shell.SetVar(name, v)

}

func (shell *Shell) SetLocalVar(name string, value string) {
	shell.localVars.set(name, &parameter{value: value})
}

func (shell *Shell) SetExportVar(name string, value string) {
	shell.exportedVars.set(name, struct{}{})
	shell.vars.set(name, &parameter{value: value})
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

func (shell *Shell) ReadSpecialVar(name string) string {
	switch name {
	case "0":
		return shell.Arg0
	case "$":
		return strconv.FormatInt(int64(shell.PID), 10)
	case "#":
		return strconv.FormatInt(int64(len(shell.Args)), 10)
	case "?":
		return strconv.FormatInt(int64(shell.ExitCode), 10)
	case "*", "@":
		return strings.Join(shell.Args, " ")
	default:
		index, err := strconv.ParseUint(name, 10, 64)
		if err != nil {
			return ""
		}
		if index <= uint64(len(shell.Args)) {
			return shell.Args[index-1]
		}
		return ""
	}
}

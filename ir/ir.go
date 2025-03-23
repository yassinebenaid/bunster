package ir

import (
	"fmt"
	"path"
	"strings"
)

const EmbedDirectory = "embed"

type Instruction interface {
	togo() string
}

type Program struct {
	Instructions []Instruction
	Embeds       []string
}

func (p Program) String() string {
	var str string

	if p.Embeds != nil {
		var embeds []string
		for _, embed := range p.Embeds {
			if embed == "." {
				embed = "*"
			}

			embeds = append(embeds, fmt.Sprintf("//go:embed %q", path.Join(EmbedDirectory, embed)))
		}

		str = fmt.Sprintf(`
			package main

			import (
				"embed"
				"io/fs"
				"github.com/yassinebenaid/bunster/runtime"
			)

			%s
			var embedFS embed.FS

			func Main(shell *runtime.Shell, streamManager *runtime.StreamManager) {
				defer shell.Terminate(streamManager)
				subfs, err := fs.Sub(&embedFS, %q)
				if err != nil {
					shell.HandleError(streamManager, err)
					return
				}
				shell.Embed = subfs
			`, strings.Join(embeds, "\n"), EmbedDirectory)
	} else {
		str = `
			package main
			
			import "github.com/yassinebenaid/bunster/runtime"
			
			func Main(shell *runtime.Shell, streamManager *runtime.StreamManager) {
				defer shell.Terminate(streamManager)
			`
	}

	for _, in := range p.Instructions {
		str += in.togo()
	}

	str += `}`
	return str
}

type CloneShell struct {
	DontTerminate bool
}

func (c CloneShell) togo() string {
	if c.DontTerminate {
		return "shell := shell.Clone()\n"
	}

	return `shell := shell.Clone()
		defer shell.Terminate(streamManager)
		`
}

type DeferTerminateShell struct{}

func (c DeferTerminateShell) togo() string {
	return "defer shell.Terminate(streamManager)\n"
}

type Declare struct {
	Name  string
	Value Instruction
}

func (d Declare) togo() string {
	return fmt.Sprintf("var %s = %s\n", d.Name, d.Value.togo())
}

type DeclareSlice struct {
	Name string
}

func (d DeclareSlice) togo() string {
	return fmt.Sprintf("var %s []string\n", d.Name)
}

type DeclareMap string

func (d DeclareMap) togo() string {
	return fmt.Sprintf("var %s = make(map[string]string)\n", d)
}

type Set struct {
	Name  string
	Value Instruction
}

func (a Set) togo() string {
	return fmt.Sprintf("%s = %s\n", a.Name, a.Value.togo())
}

type Append struct {
	Name  string
	Value Instruction
}

func (a Append) togo() string {
	return fmt.Sprintf("%s = append(%s, %s)\n", a.Name, a.Name, a.Value.togo())
}

type SetMap struct {
	Name  string
	Key   string
	Value Instruction
}

func (a SetMap) togo() string {
	return fmt.Sprintf("%s[%q] = %s\n", a.Name, a.Key, a.Value.togo())
}

type String string

func (s String) togo() string {
	return fmt.Sprintf("%q", s)
}

type Concat []Instruction

func (c Concat) togo() string {
	var str string
	for i, ins := range c {
		str += ins.togo()
		if i < len(c)-1 {
			str += "+"
		}
	}
	return str
}

type Literal string

func (s Literal) togo() string {
	return string(s)
}

type ReadVar string

func (rv ReadVar) togo() string {
	return fmt.Sprintf("shell.ReadVar(%q)", rv)
}

type SetVar struct {
	Key   string
	Value Instruction
}

func (s SetVar) togo() string {
	return fmt.Sprintf("shell.SetVar(%q, %v)\n", s.Key, s.Value.togo())
}

type SetLocalVar struct {
	Key   string
	Value Instruction
}

func (s SetLocalVar) togo() string {
	return fmt.Sprintf("shell.SetLocalVar(%q, %v)\n", s.Key, s.Value.togo())
}

type SetExportVar struct {
	Key   string
	Value Instruction
}

func (s SetExportVar) togo() string {
	return fmt.Sprintf("shell.SetExportVar(%q, %v)\n", s.Key, s.Value.togo())
}

type MarkVarAsExported string

func (s MarkVarAsExported) togo() string {
	return fmt.Sprintf("shell.MarkVarAsExported(%q)\n", s)
}

type ReadSpecialVar string

func (rv ReadSpecialVar) togo() string {
	return fmt.Sprintf("shell.ReadSpecialVar(%q)", rv)
}

type InitCommand struct {
	Name string
	Args string
	Env  string
}

func (ic InitCommand) togo() string {
	return fmt.Sprintf("shell.Command(%s, %s, %s)", ic.Name, ic.Args, ic.Env)
}

type RunCommand string

func (r RunCommand) togo() string {
	return fmt.Sprintf(
		`if err := %s.Run(shell, streamManager); err != nil {
			shell.HandleError(streamManager, err)
			return
		}
		shell.ExitCode = %s.ExitCode
		`, r, r)
}

type Exec struct {
	Name string
	Args string
	Env  string
}

func (e Exec) togo() string {
	return fmt.Sprintf(
		`if err := shell.Exec(streamManager, %s, %s, %s); err != nil {
			shell.HandleError(streamManager, err)
			return
		}
		`, e.Name, e.Args, e.Env,
	)
}

type Procedure struct {
	Returns []string
	Body    []Instruction
}

func (c Procedure) togo() string {
	var body string
	for _, ins := range c.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`func() (%s) {
			%s
		}`, strings.Join(c.Returns, ", "), body)
}

type Closure []Instruction

func (c Closure) togo() string {
	var body string
	for _, ins := range c {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`func(){
			%s
		}()
	`, body)
}

type Scope []Instruction

func (s Scope) togo() string {
	var body string
	for _, ins := range s {
		body += ins.togo()
	}

	return fmt.Sprintf(`{
			%s
		}
	`, body)
}

type Gorouting []Instruction

func (g Gorouting) togo() string {
	var body string
	for _, ins := range g {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`go func(){
			%s
		}()
		`, body)
}

type ExpressionClosure struct {
	Body []Instruction
	Name string
}

func (c ExpressionClosure) togo() string {
	var body string

	for _, ins := range c.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`%s, exitCode := func() (string, int) {
			%s
		}()
		if exitCode != 0 {
			shell.ExitCode = exitCode
			return
		}
		`, c.Name, body)
}

type SetCmdEnv struct {
	Command string
	Key     string
	Value   Instruction
}

func (s SetCmdEnv) togo() string {
	return fmt.Sprintf("%s.Env[%q] = %s\n", s.Command, s.Key, s.Value.togo())
}

type IfLastExitCode struct {
	Zero bool
	Body []Instruction
}

func (i IfLastExitCode) togo() string {
	var condition = "shell.ExitCode == 0"
	if !i.Zero {
		condition = "shell.ExitCode != 0"
	}

	var body string
	for _, ins := range i.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		`if %s {
			%s
		}
		`, condition, body)
}

type If struct {
	Condition Instruction
	Body      []Instruction
	Alternate []Instruction
}

func (i If) togo() string {
	cond := fmt.Sprintf("if %s {\n", i.Condition)
	for _, ins := range i.Body {
		cond += ins.togo()
	}

	if len(i.Alternate) > 0 {
		cond += "} else {"

		for _, ins := range i.Alternate {
			cond += ins.togo()
		}
	}

	return cond + "}\n"
}

type Loop struct {
	Condition Instruction
	Body      []Instruction
}

func (i Loop) togo() string {
	cond := fmt.Sprintf("for %s {\n", i.Condition)
	for _, ins := range i.Body {
		cond += ins.togo()
	}

	return cond + "}\n"
}

type RangeLoop struct {
	Var     string
	Members Instruction
	Body    []Instruction
}

func (i RangeLoop) togo() string {
	cond := fmt.Sprintf(
		`for _,member := range %s {
			shell.SetVar(%q, member)
		`, i.Members.togo(), i.Var,
	)
	for _, ins := range i.Body {
		cond += ins.togo()
	}

	return cond + "}\n"
}

type For struct {
	Init   Instruction
	Test   Instruction
	Update Instruction
	Body   []Instruction
}

func (i For) togo() string {
	cond := fmt.Sprintf("for %s; %s; %s {\n", i.Init.togo(), i.Test.togo(), i.Update.togo())
	for _, ins := range i.Body {
		cond += ins.togo()
	}

	return cond + "}\n"
}

type InvertExitCode struct{}

func (i InvertExitCode) togo() string {
	return `if shell.ExitCode == 0 {
			shell.ExitCode = 1
		} else {
			shell.ExitCode = 0
		}
		`
}

type Function struct {
	Name string
	Body []Instruction
}

func (f Function) togo() string {
	var body string
	for _, ins := range f.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		"shell.RegisterFunction(%q, func(shell *runtime.Shell, streamManager *runtime.StreamManager){"+`
			defer shell.Terminate(streamManager)
			%s
		`+"})\n",
		f.Name, body,
	)
}

type Defer struct {
	Body []Instruction
}

func (f Defer) togo() string {
	var body string
	for _, ins := range f.Body {
		body += ins.togo()
	}

	return fmt.Sprintf(
		"shell.Defer(func(shell *runtime.Shell, streamManager *runtime.StreamManager){"+`
			%s
		`+"})\n",
		body,
	)
}

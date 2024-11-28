package ir

import "fmt"

type Instruction interface {
	inst()
	fmt.Stringer
}

type Program struct {
	Instructions []Instruction
}

type Assign struct {
	Name       string
	Initialize bool
	Value      Instruction
}

type String string
type Literal string

type InitCommand struct {
	Name string
	Args string
}

type RunCommanOrFail struct {
	Name    string
	OnError Instruction
}

type Panic string

func (Assign) inst()          {}
func (String) inst()          {}
func (Literal) inst()         {}
func (InitCommand) inst()     {}
func (RunCommanOrFail) inst() {}

func (a Assign) String() string {
	op := "="
	if a.Initialize {
		op = ":="
	}
	return fmt.Sprintf("%s %s %s\n", a.Name, op, a.Value.String())
}
func (s String) String() string {
	return fmt.Sprintf(`"%s"`, string(s))
}
func (s Literal) String() string {
	return fmt.Sprintf(`%s`, string(s))
}
func (ic InitCommand) String() string {
	return fmt.Sprintf("exec.Command(%s)", ic.Name)
}
func (rcf RunCommanOrFail) String() string {
	return fmt.Sprintf(`
		if err := %s.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "command %%q not found.\n", %s.Path)
			os.Exit(1)
		}
		`, rcf.Name, rcf.Name)
}

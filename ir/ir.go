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
	Name  string
	Value Instruction
}

type String string

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
func (InitCommand) inst()     {}
func (RunCommanOrFail) inst() {}

func (a Assign) String() string {
	return fmt.Sprintf("%s := %s\n", a.Name, a.Value.String())
}
func (s String) String() string {
	return fmt.Sprintf(`"%s"`, string(s))
}
func (ic InitCommand) String() string {
	return fmt.Sprintf("exec.Command(%s)", ic.Name)
}
func (rcf RunCommanOrFail) String() string {
	return fmt.Sprintf(`
		if err := %s.Run(); err != nil {
			panic(err)
		}
		`, rcf.Name)
}

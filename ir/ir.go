package ir

type Instruction interface {
	inst()
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

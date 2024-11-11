package ast

type Node interface {
	node()
}

type Statement interface {
	Node
	stmt()
}

type Expression interface {
	Node
	expr()
	string() string
}

type Script []Statement

type BinaryConstruction struct {
	Left     Statement
	Operator string // || or &&
	Right    Statement
}

type BackgroundConstruction struct {
	Statement
}

type PipelineCommand struct {
	Stderr  bool
	Command Statement
}

type Pipeline []PipelineCommand

type Word string

type Redirection struct {
	Src    string
	Method string
	Dst    Expression
}

type Command struct {
	Name         Expression
	Args         []Expression
	Redirections []Redirection
}

type Concatination struct {
	Nodes []Expression
}

type Loop struct {
	Negate       bool
	Head         []Statement
	Body         []Statement
	Redirections []Redirection
}

type RangeLoop struct {
	Var          string
	Operands     []Expression
	Body         []Statement
	Redirections []Redirection
}

type For struct {
	Head         ForHead
	Body         []Statement
	Redirections []Redirection
}

type ForHead struct {
	Init   Arithmetic
	Test   Arithmetic
	Update Arithmetic
}

type If struct {
	Head         []Statement
	Body         []Statement
	Elifs        []Elif
	Alternate    []Statement
	Redirections []Redirection
}

type Elif struct {
	Head []Statement
	Body []Statement
}

type Case struct {
	Word         Expression
	Cases        []CaseItem
	Redirections []Redirection
}

type CaseItem struct {
	Patterns   []Expression
	Body       []Statement
	Terminator string
}

type Group struct {
	Body         []Statement
	Redirections []Redirection
}

type SubShell struct {
	Body         []Statement
	Redirections []Redirection
}

type CommandSubstitution []Statement

type ProcessSubstitution struct {
	Direction rune
	Body      []Statement
}

func (Word) node()                {}
func (Redirection) node()         {}
func (Concatination) node()       {}
func (Command) node()             {}
func (Pipeline) node()            {}
func (BinaryConstruction) node()  {}
func (Loop) node()                {}
func (RangeLoop) node()           {}
func (If) node()                  {}
func (Case) node()                {}
func (Group) node()               {}
func (SubShell) node()            {}
func (CommandSubstitution) node() {}
func (ProcessSubstitution) node() {}
func (For) node()                 {}

// Expressions
func (Word) expr()                {}
func (Redirection) expr()         {}
func (Concatination) expr()       {}
func (CommandSubstitution) expr() {}
func (ProcessSubstitution) expr() {}

func (Word) string() string                { return "" }
func (Redirection) string() string         { return "" }
func (Concatination) string() string       { return "" }
func (CommandSubstitution) string() string { return "" }
func (ProcessSubstitution) string() string { return "" }
func (For) string() string                 { return "" }

// Statements
func (Command) stmt()            {}
func (Pipeline) stmt()           {}
func (BinaryConstruction) stmt() {}
func (Loop) stmt()               {}
func (RangeLoop) stmt()          {}
func (If) stmt()                 {}
func (Case) stmt()               {}
func (Group) stmt()              {}
func (SubShell) stmt()           {}
func (For) stmt()                {}

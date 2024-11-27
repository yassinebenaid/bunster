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

type List struct {
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

type Redirection struct {
	Src    string
	Method string
	Dst    Expression
}

type Command struct {
	Name         Expression
	Args         []Expression
	Redirections []Redirection
	Env          []Assignement
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

type Function struct {
	Name    string
	Command Statement
}

type Test struct {
	Expr         Expression
	Redirections []Redirection
}

func (Redirection) node()         {}
func (Command) node()             {}
func (Pipeline) node()            {}
func (List) node()                {}
func (Loop) node()                {}
func (RangeLoop) node()           {}
func (If) node()                  {}
func (Case) node()                {}
func (Group) node()               {}
func (SubShell) node()            {}
func (CommandSubstitution) node() {}
func (ProcessSubstitution) node() {}
func (For) node()                 {}
func (Function) node()            {}
func (Test) node()                {}

// Expressions
func (Redirection) expr()         {}
func (CommandSubstitution) expr() {}
func (ProcessSubstitution) expr() {}

func (CommandSubstitution) string() string { return "" }
func (ProcessSubstitution) string() string { return "" }

// Statements
func (Command) stmt()   {}
func (Pipeline) stmt()  {}
func (List) stmt()      {}
func (Loop) stmt()      {}
func (RangeLoop) stmt() {}
func (If) stmt()        {}
func (Case) stmt()      {}
func (Group) stmt()     {}
func (SubShell) stmt()  {}
func (For) stmt()       {}
func (Function) stmt()  {}
func (Test) stmt()      {}
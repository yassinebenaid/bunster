package ast

type Node interface {
	node()
}

type Statement interface {
	Node
	stmt()
}

type BreakPointsType byte

const (
	RETURN   BreakPointsType = iota
	BREAK                    = iota
	CONTINUE                 = iota
	DECLARE                  = iota
)

type BreakPoints map[int]BreakPointsType

func (b *BreakPoints) Add(k int, t BreakPointsType) {
	if *b == nil {
		*b = make(BreakPoints)
	}
	(*b)[k] = t
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

type InvertExitCode struct {
	Statement
}

type Redirection struct {
	Src    string
	Method string
	Dst    Expression
	Close  bool
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
	BreakPoints
}

type RangeLoop struct {
	Var          string
	Operands     []Expression
	Body         []Statement
	Redirections []Redirection
}

type Break struct {
	BreakPoint int
	Type       BreakPointsType
}

type Exit struct {
	Code Expression
}

type Continue struct {
	BreakPoint int
	Type       BreakPointsType
}

type Wait struct{}

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
	BreakPoints
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
	Name         string
	SubShell     bool
	Body         []Statement
	Redirections []Redirection
}

type Defer struct {
	Command Statement
}

type Test struct {
	Expr         Expression
	Redirections []Redirection
}

type Embed []string

func (Redirection) node()         {}
func (Command) node()             {}
func (Pipeline) node()            {}
func (InvertExitCode) node()      {}
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
func (Break) node()               {}
func (Continue) node()            {}
func (Exit) node()                {}
func (Wait) node()                {}
func (Embed) node()               {}
func (Defer) node()               {}

// Expressions
func (CommandSubstitution) expr() {}
func (ProcessSubstitution) expr() {}

func (CommandSubstitution) string() string { return "" }
func (ProcessSubstitution) string() string { return "" }

// Statements
func (Command) stmt()        {}
func (Pipeline) stmt()       {}
func (InvertExitCode) stmt() {}
func (List) stmt()           {}
func (Loop) stmt()           {}
func (RangeLoop) stmt()      {}
func (If) stmt()             {}
func (Case) stmt()           {}
func (Group) stmt()          {}
func (SubShell) stmt()       {}
func (For) stmt()            {}
func (Function) stmt()       {}
func (Test) stmt()           {}
func (Break) stmt()          {}
func (Continue) stmt()       {}
func (Exit) stmt()           {}
func (Wait) stmt()           {}
func (Embed) stmt()          {}
func (Defer) stmt()          {}

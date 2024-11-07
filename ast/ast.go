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
type Number string

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

type Var string

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

type VarOrDefault struct {
	Name         string
	Default      Expression
	CheckForNull bool
}

type VarOrSet struct {
	Name    string
	Default Expression
}

type VarOrFail struct {
	Name  string
	Error Expression
}

type CheckAndUse struct {
	Name  string
	Value Expression
}

type ChangeCase struct {
	Name     string
	Operator string
	Pattern  Expression
}

type VarCount string

type MatchAndRemove struct {
	Name     string
	Operator string
	Pattern  Expression
}

type MatchAndReplace struct {
	Name     string
	Operator string
	Pattern  Expression
	Value    Expression
}

type Arithmetic []Expression

type InfixArithmetic struct {
	Left     Expression
	Operator string
	Right    Expression
}

type PostIncDecArithmetic struct {
	Operand  Expression
	Operator string
}

type PreIncDecArithmetic struct {
	Operand  Expression
	Operator string
}

type Unary struct {
	Operand  Expression
	Operator string
}

func (Word) node()                 {}
func (Number) node()               {}
func (Redirection) node()          {}
func (Var) node()                  {}
func (Concatination) node()        {}
func (Command) node()              {}
func (Pipeline) node()             {}
func (BinaryConstruction) node()   {}
func (Loop) node()                 {}
func (RangeLoop) node()            {}
func (If) node()                   {}
func (Case) node()                 {}
func (Group) node()                {}
func (SubShell) node()             {}
func (CommandSubstitution) node()  {}
func (ProcessSubstitution) node()  {}
func (VarOrDefault) node()         {}
func (VarOrSet) node()             {}
func (VarOrFail) node()            {}
func (CheckAndUse) node()          {}
func (ChangeCase) node()           {}
func (VarCount) node()             {}
func (MatchAndRemove) node()       {}
func (MatchAndReplace) node()      {}
func (Arithmetic) node()           {}
func (InfixArithmetic) node()      {}
func (PostIncDecArithmetic) node() {}
func (PreIncDecArithmetic) node()  {}
func (Unary) node()                {}

// Expressions
func (Word) expr()                 {}
func (Number) expr()               {}
func (Redirection) expr()          {}
func (Var) expr()                  {}
func (Concatination) expr()        {}
func (CommandSubstitution) expr()  {}
func (ProcessSubstitution) expr()  {}
func (VarOrDefault) expr()         {}
func (VarOrSet) expr()             {}
func (VarOrFail) expr()            {}
func (CheckAndUse) expr()          {}
func (ChangeCase) expr()           {}
func (VarCount) expr()             {}
func (MatchAndRemove) expr()       {}
func (MatchAndReplace) expr()      {}
func (Arithmetic) expr()           {}
func (InfixArithmetic) expr()      {}
func (PostIncDecArithmetic) expr() {}
func (PreIncDecArithmetic) expr()  {}
func (Unary) expr()                {}

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

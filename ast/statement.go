package ast

import "github.com/yassinebenaid/bunster/token"

type Node interface {
	node()
}

type Statement interface {
	GetToken() token.Token
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
	Token    token.Token
	Left     Statement
	Operator string // || or &&
	Right    Statement
}

type BackgroundConstruction struct {
	Token token.Token
	Statement
}

type PipelineCommand struct {
	Token   token.Token
	Stderr  bool
	Command Statement
}

type Pipeline struct {
	Token    token.Token
	Commands []PipelineCommand
}

type Redirection struct {
	Src    string
	Method string
	Dst    Expression
	Close  bool
}

type Command struct {
	Token        token.Token
	Name         Expression
	Args         []Expression
	Redirections []Redirection
	Env          []Assignement
}

type Loop struct {
	Token        token.Token
	Negate       bool
	Head         []Statement
	Body         []Statement
	Redirections []Redirection
}

type RangeLoop struct {
	Token        token.Token
	Var          string
	Operands     []Expression
	Body         []Statement
	Redirections []Redirection
}

type For struct {
	Token        token.Token
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
	Token        token.Token
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
	Token        token.Token
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
	Token        token.Token
	Body         []Statement
	Redirections []Redirection
}

type SubShell struct {
	Token        token.Token
	Body         []Statement
	Redirections []Redirection
}

type CommandSubstitution []Statement

type ProcessSubstitution struct {
	Token     token.Token
	Direction rune
	Body      []Statement
}

type Function struct {
	Token   token.Token
	Name    string
	Command Statement
}

type Test struct {
	Token        token.Token
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

func (n Command) GetToken() token.Token             { return n.Token }
func (n Pipeline) GetToken() token.Token            { return n.Token }
func (n List) GetToken() token.Token                { return n.Token }
func (n Loop) GetToken() token.Token                { return n.Token }
func (n RangeLoop) GetToken() token.Token           { return n.Token }
func (n If) GetToken() token.Token                  { return n.Token }
func (n Case) GetToken() token.Token                { return n.Token }
func (n Group) GetToken() token.Token               { return n.Token }
func (n SubShell) GetToken() token.Token            { return n.Token }
func (n ProcessSubstitution) GetToken() token.Token { return n.Token }
func (n For) GetToken() token.Token                 { return n.Token }
func (n Function) GetToken() token.Token            { return n.Token }
func (n Test) GetToken() token.Token                { return n.Token }

// Expressions
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

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

type Script struct {
	Statements []Statement
}

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

type SimpleExpansion string

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

func (Word) node()               {}
func (Redirection) node()        {}
func (SimpleExpansion) node()    {}
func (Concatination) node()      {}
func (Command) node()            {}
func (Pipeline) node()           {}
func (BinaryConstruction) node() {}
func (Loop) node()               {}
func (RangeLoop) node()          {}
func (If) node()                 {}
func (Case) node()               {}

// Expressions
func (Word) expr()            {}
func (Redirection) expr()     {}
func (SimpleExpansion) expr() {}
func (Concatination) expr()   {}

// Statements
func (Command) stmt()            {}
func (Pipeline) stmt()           {}
func (BinaryConstruction) stmt() {}
func (Loop) stmt()               {}
func (RangeLoop) stmt()          {}
func (If) stmt()                 {}
func (Case) stmt()               {}

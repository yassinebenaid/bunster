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
	Statements []Node
}

type BinaryConstruction struct {
	Left     Node
	Operator string // || or &&
	Right    Node
}

type BackgroundConstruction struct {
	Node
}

type PipelineCommand struct {
	Stderr  bool
	Command Node
}

type Pipeline []PipelineCommand

type Word string

type Redirection struct {
	Src    string
	Method string
	Dst    Node
}

type Command struct {
	Name         Node
	Args         []Node
	Redirections []Redirection
}

type SimpleExpansion string

type Concatination struct {
	Nodes []Node
}

type Loop struct {
	Negate       bool
	Head         []Node
	Body         []Node
	Redirections []Redirection
}

type RangeLoop struct {
	Var          string
	Operands     []Node
	Body         []Node
	Redirections []Redirection
}

type If struct {
	Head         []Node
	Body         []Node
	Elifs        []Elif
	Alternate    []Node
	Redirections []Redirection
}

type Elif struct {
	Head []Node
	Body []Node
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

package ast

type Node interface {
	node() // Just to distinguish it.
}

type Script struct {
	Statements []Node
}

type Word string

func (Word) node() {}

type Redirection struct {
	Src    Node
	Method string
	Dst    Node
}

func (Redirection) node() {}

type FileDescriptor string

func (FileDescriptor) node() {}

type Command struct {
	Name         Node
	Args         []Node
	Redirections []Redirection
}

func (Command) node() {}

type SimpleExpansion string

func (SimpleExpansion) node() {}

type Concatination struct {
	Nodes []Node
}

func (Concatination) node() {}

package ast

type Node interface {
	node() // Just to distinguish it.
}

type Script struct {
	Statements []Node
}

type Word string

func (Word) node() {}

type Command struct {
	Name Node
	Args []Node
}

func (Command) node() {}

type SimpleExpansion string

func (SimpleExpansion) node() {}

type Concatination struct {
	Nodes []Node
}

func (Concatination) node() {}

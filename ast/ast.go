package ast

type Node interface {
	node() // Just to distinguish it.
}

type Script struct {
	Statements []Node
}

type Command struct {
	Name string
	Args []string
}

func (cc Command) node() {}

package ast

type Node interface {
	TokenLiteral() string
}

type Program struct {
	Nodes []Node
}

type CommandCall struct {
	Command string
	Args    []string
}

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

func (cc CommandCall) TokenLiteral() string {
	return cc.Command
}

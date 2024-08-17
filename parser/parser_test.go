package parser

import (
	"testing"

	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/lexer"
)

func TestCanParseCommandCall(t *testing.T) {
	input := `
		git branch
	`

	l := lexer.New([]byte(input))
	p := New(l)
	program := p.ParseProgram()

	if len := len(program.Nodes); len != 1 {
		t.Fatalf(`expected program to have 1 expression, got "%d".`, len)
	}

	cmd, ok := program.Nodes[0].(ast.CommandCall)

	if !ok {
		t.Fatalf(`expected first node to be 'ast.CommandCall', got "%T".`, program.Nodes[0])
	}

	if cmd.Command != "git" {
		t.Fatalf(`expected command to be 'git', got "%s".`, cmd.Command)
	}

	if len(cmd.Args) != 1 {
		t.Fatalf(`expected command args length to be '1', got "%d".`, len(cmd.Args))
	}

	if cmd.Args[0] != "branch" {
		t.Fatalf(`expected first command arg to be 'branch', got "%s".`, cmd.Args[0])
	}
}

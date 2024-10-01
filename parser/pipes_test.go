package parser_test

import "github.com/yassinebenaid/nbs/ast"

var pipesTests = []testCase{
	{` cmd | cmd2 |& cmd3`, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
			},
		},
	}},
}

var pipesErrorHandlingCases = []errorHandlingTestCase{
	// {`cmd >`, "syntax error: a pipes operand was not provided after the `>`."},
}

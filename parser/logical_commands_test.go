package parser_test

import "github.com/yassinebenaid/nbs/ast"

var logicalCommandsTests = []testCase{
	{` cmd && cmd2 `, ast.Script{
		Statements: []ast.Node{
			ast.LogicalCommand{
				Left:     ast.Command{Name: ast.Word("cmd")},
				Operator: "&&",
				Right:    ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},
	{`cmd&&cmd2`, ast.Script{
		Statements: []ast.Node{
			ast.LogicalCommand{
				Left:     ast.Command{Name: ast.Word("cmd")},
				Operator: "&&",
				Right:    ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},
}

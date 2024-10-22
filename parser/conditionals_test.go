package parser_test

import "github.com/yassinebenaid/bunny/ast"

var conditionalsTests = []testCase{
	{`if cmd; then cmd2; fi`, ast.Script{
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
		},
	}},
}

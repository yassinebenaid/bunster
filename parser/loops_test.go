package parser_test

import "github.com/yassinebenaid/bunny/ast"

var loopsTests = []testCase{
	{`while cmd; do echo "Hello World"; done`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{ast.Command{Name: ast.Word("cmd")}},
				Body: []ast.Node{
					ast.Command{
						Name: ast.Word("echo"),
						Args: []ast.Node{ast.Word("Hello World")},
					},
				},
			},
		},
	}},
}

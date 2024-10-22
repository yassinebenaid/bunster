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
	{`if
		cmd;
	 then
		cmd2;
	fi`, ast.Script{
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
	{`if
		cmd1 | cmd2 && cmd3
	 then
		echo 'baz'
	fi`, ast.Script{
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.BinaryConstruction{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
			},
		},
	}},
}

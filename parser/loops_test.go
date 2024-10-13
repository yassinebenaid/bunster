package parser_test

import "github.com/yassinebenaid/bunny/ast"

var loopsTests = []testCase{
	{`while cmd1; cmd2; cmd3; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd1")},
					ast.Command{Name: ast.Word("cmd2")},
					ast.Command{Name: ast.Word("cmd3")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("bar")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
			},
		},
	}},
	{`while
		cmd1
		cmd2
		cmd3
	do
		echo "foo"
		echo bar
		echo 'baz'
	done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd1")},
					ast.Command{Name: ast.Word("cmd2")},
					ast.Command{Name: ast.Word("cmd3")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("bar")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
			},
		},
	}},
	{`while
		cmd1 | cmd2 && cmd3
	do
		echo 'baz'
	done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
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

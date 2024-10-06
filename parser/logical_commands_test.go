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
	{`cmd >foo arg <<<"foo bar" | cmd2 <input.txt 'foo bar baz' && cmd >foo $var 3<<<"foo bar" |& cmd2 "foo bar baz" <input.txt `, ast.Script{
		Statements: []ast.Node{
			ast.LogicalCommand{
				Left: ast.Pipeline{
					{
						Command: ast.Command{
							Name: ast.Word("cmd"),
							Args: []ast.Node{ast.Word("arg")},
							Redirections: []ast.Redirection{
								{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("foo")},
								{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("foo bar")},
							},
						},
						Stderr: false,
					},
					{
						Command: ast.Command{
							Name: ast.Word("cmd2"),
							Args: []ast.Node{ast.Word("foo bar baz")},
							Redirections: []ast.Redirection{
								{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("input.txt")},
							},
						},
						Stderr: false,
					},
				},
				Operator: "&&",
				Right: ast.Pipeline{
					{
						Command: ast.Command{
							Name: ast.Word("cmd"),
							Args: []ast.Node{ast.SimpleExpansion("var")},
							Redirections: []ast.Redirection{
								{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("foo")},
								{Src: ast.FileDescriptor("3"), Method: "<<<", Dst: ast.Word("foo bar")},
							},
						},
						Stderr: false,
					},
					{
						Command: ast.Command{
							Name: ast.Word("cmd2"),
							Args: []ast.Node{ast.Word("foo bar baz")},
							Redirections: []ast.Redirection{
								{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("input.txt")},
							},
						},
						Stderr: true,
					},
				},
			},
		},
	}},
	{` cmd && cmd2 && cmd3 && cmd4`, ast.Script{
		Statements: []ast.Node{
			ast.LogicalCommand{
				Left: ast.LogicalCommand{
					Left: ast.LogicalCommand{
						Left:     ast.Command{Name: ast.Word("cmd")},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd2")},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				Operator: "&&",
				Right:    ast.Command{Name: ast.Word("cmd4")},
			},
		},
	}},
	{` cmd&&cmd2&&cmd3&&cmd4`, ast.Script{
		Statements: []ast.Node{
			ast.LogicalCommand{
				Left: ast.LogicalCommand{
					Left: ast.LogicalCommand{
						Left:     ast.Command{Name: ast.Word("cmd")},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd2")},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				Operator: "&&",
				Right:    ast.Command{Name: ast.Word("cmd4")},
			},
		},
	}},
}

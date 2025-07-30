package parser_test

import "github.com/yassinebenaid/bunster/ast"

var commandListTests = []testCase{
	{` cmd && cmd2 `, ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
		},
	}},
	{`cmd&&cmd2`, ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 6}, Name: ast.Word("cmd2")},
		},
	}},
	{` cmd1 | cmd2 && cmd3 `, ast.Script{
		ast.List{
			Left: ast.Pipeline{
				{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd1")}},
				{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")}},
			},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd3")},
		},
	}},
	{`cmd >foo arg <<<"foo bar" | cmd2 <input.txt 'foo bar baz' && cmd >foo $var 3<<<"foo bar" |& cmd2 "foo bar baz" <input.txt `, ast.Script{
		ast.List{
			Left: ast.Pipeline{
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.Word("cmd"),
						Args: []ast.Expression{ast.Word("arg")},
						Redirections: []ast.Redirection{
							{Src: "1", Method: ">", Dst: ast.Word("foo")},
							{Src: "0", Method: "<<<", Dst: ast.Word("foo bar")},
						},
					},
					Stderr: false,
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 29},
						Name: ast.Word("cmd2"),
						Args: []ast.Expression{ast.Word("foo bar baz")},
						Redirections: []ast.Redirection{
							{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
						},
					},
					Stderr: false,
				},
			},
			Operator: "&&",
			Right: ast.Pipeline{
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 62},
						Name: ast.Word("cmd"),
						Args: []ast.Expression{ast.Var("var")},
						Redirections: []ast.Redirection{
							{Src: "1", Method: ">", Dst: ast.Word("foo")},
							{Src: "3", Method: "<<<", Dst: ast.Word("foo bar")},
						},
					},
					Stderr: true,
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 93},
						Name: ast.Word("cmd2"),
						Args: []ast.Expression{ast.Word("foo bar baz")},
						Redirections: []ast.Redirection{
							{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
						},
					},
					Stderr: false,
				},
			},
		},
	}},
	{` cmd && cmd2 && cmd3 && cmd4`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.List{
					Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
					Operator: "&&",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
				},
				Operator: "&&",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd3")},
			},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 25}, Name: ast.Word("cmd4")},
		},
	}},
	{` cmd&&cmd2&&cmd3&&cmd4`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.List{
					Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
					Operator: "&&",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 7}, Name: ast.Word("cmd2")},
				},
				Operator: "&&",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 13}, Name: ast.Word("cmd3")},
			},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd4")},
		},
	}},

	{" cmd && cmd2; cmd3 && cmd4\n", ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
		},
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 15}, Name: ast.Word("cmd3")},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 23}, Name: ast.Word("cmd4")},
		},
	}},

	{" cmd && cmd2 && cmd3 & cmd", ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.List{
				Left: ast.List{
					Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
					Operator: "&&",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
				},
				Operator: "&&",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd3")},
			},
		},
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 24}, Name: ast.Word("cmd")},
	}},

	{` cmd || cmd2 `, ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
		},
	}},
	{`cmd||cmd2`, ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 6}, Name: ast.Word("cmd2")},
		},
	}},
	{`cmd >foo arg <<<"foo bar" | cmd2 <input.txt 'foo bar baz' || cmd >foo $var 3<<<"foo bar" |& cmd2 "foo bar baz" <input.txt `, ast.Script{
		ast.List{
			Left: ast.Pipeline{
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.Word("cmd"),
						Args: []ast.Expression{ast.Word("arg")},
						Redirections: []ast.Redirection{
							{Src: "1", Method: ">", Dst: ast.Word("foo")},
							{Src: "0", Method: "<<<", Dst: ast.Word("foo bar")},
						},
					},
					Stderr: false,
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 29},
						Name: ast.Word("cmd2"),
						Args: []ast.Expression{ast.Word("foo bar baz")},
						Redirections: []ast.Redirection{
							{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
						},
					},
					Stderr: false,
				},
			},
			Operator: "||",
			Right: ast.Pipeline{
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 62},
						Name: ast.Word("cmd"),
						Args: []ast.Expression{ast.Var("var")},
						Redirections: []ast.Redirection{
							{Src: "1", Method: ">", Dst: ast.Word("foo")},
							{Src: "3", Method: "<<<", Dst: ast.Word("foo bar")},
						},
					},
					Stderr: true,
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 93},
						Name: ast.Word("cmd2"),
						Args: []ast.Expression{ast.Word("foo bar baz")},
						Redirections: []ast.Redirection{
							{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
						},
					},
					Stderr: false,
				},
			},
		},
	}},
	{` cmd || cmd2 || cmd3 || cmd4`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.List{
					Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
					Operator: "||",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
				},
				Operator: "||",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd3")},
			},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 25}, Name: ast.Word("cmd4")},
		},
	}},
	{` cmd||cmd2||cmd3||cmd4`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.List{
					Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
					Operator: "||",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 7}, Name: ast.Word("cmd2")},
				},
				Operator: "||",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 13}, Name: ast.Word("cmd3")},
			},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd4")},
		},
	}},
	{` cmd || cmd2 && cmd3 || cmd4 && cmd5`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.List{
					Left: ast.List{
						Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
						Operator: "||",
						Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
					},
					Operator: "&&",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd3")},
				},
				Operator: "||",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 25}, Name: ast.Word("cmd4")},
			},
			Operator: "&&",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 33}, Name: ast.Word("cmd5")},
		},
	}},
	{"cmd || \n\t\n cmd2 ", ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 2}, Name: ast.Word("cmd2")},
		},
	}},

	{" cmd || cmd2; cmd3 || cmd4\n", ast.Script{
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
		},
		ast.List{
			Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 15}, Name: ast.Word("cmd3")},
			Operator: "||",
			Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 23}, Name: ast.Word("cmd4")},
		},
	}},

	{" cmd || cmd2 || cmd3 & cmd", ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.List{
				Left: ast.List{
					Left:     ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")},
					Operator: "||",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd2")},
				},
				Operator: "||",
				Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd3")},
			},
		},
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 24}, Name: ast.Word("cmd")},
	}},
}

var commandListErrorHandlingCases = []errorHandlingTestCase{
	{`cmd &&`, "main.sh(1:7): syntax error: expected a valid command name, found `end of file`."},
	{`cmd ||`, "main.sh(1:7): syntax error: expected a valid command name, found `end of file`."},
	{`cmd || cmd && cmd ||`, "main.sh(1:21): syntax error: expected a valid command name, found `end of file`."},
	{`cmd && cmd || cmd &&`, "main.sh(1:21): syntax error: expected a valid command name, found `end of file`."},
}

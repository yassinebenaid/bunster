package parser_test

import "github.com/yassinebenaid/bunny/ast"

var logicalCommandsTests = []testCase{
	{` cmd && cmd2 `, ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "&&",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
	}},
	{`cmd&&cmd2`, ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "&&",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
	}},
	{` cmd1 | cmd2 && cmd3 `, ast.Script{
		ast.ConditionalCommand{
			Left: ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd1")}},
				{Command: ast.Command{Name: ast.Word("cmd2")}},
			},
			Operator: "&&",
			Right:    ast.Command{Name: ast.Word("cmd3")},
		},
	}},
	{`cmd >foo arg <<<"foo bar" | cmd2 <input.txt 'foo bar baz' && cmd >foo $var 3<<<"foo bar" |& cmd2 "foo bar baz" <input.txt `, ast.Script{
		ast.ConditionalCommand{
			Left: ast.Pipeline{
				{
					Command: ast.Command{
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
					Command: ast.Command{
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
					Command: ast.Command{
						Name: ast.Word("cmd"),
						Args: []ast.Expression{ast.Var("var")},
						Redirections: []ast.Redirection{
							{Src: "1", Method: ">", Dst: ast.Word("foo")},
							{Src: "3", Method: "<<<", Dst: ast.Word("foo bar")},
						},
					},
					Stderr: false,
				},
				{
					Command: ast.Command{
						Name: ast.Word("cmd2"),
						Args: []ast.Expression{ast.Word("foo bar baz")},
						Redirections: []ast.Redirection{
							{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
						},
					},
					Stderr: true,
				},
			},
		},
	}},
	{` cmd && cmd2 && cmd3 && cmd4`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
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
	}},
	{` cmd&&cmd2&&cmd3&&cmd4`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
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
	}},

	{" cmd && cmd2; cmd3 && cmd4\n", ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "&&",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd3")},
			Operator: "&&",
			Right:    ast.Command{Name: ast.Word("cmd4")},
		},
	}},

	{" cmd && cmd2 && cmd3 & cmd", ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
					Left:     ast.Command{Name: ast.Word("cmd")},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd2")},
				},
				Operator: "&&",
				Right:    ast.Command{Name: ast.Word("cmd3")},
			},
		},
		ast.Command{Name: ast.Word("cmd")},
	}},

	{` cmd || cmd2 `, ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
	}},
	{`cmd||cmd2`, ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
	}},
	{`cmd >foo arg <<<"foo bar" | cmd2 <input.txt 'foo bar baz' || cmd >foo $var 3<<<"foo bar" |& cmd2 "foo bar baz" <input.txt `, ast.Script{
		ast.ConditionalCommand{
			Left: ast.Pipeline{
				{
					Command: ast.Command{
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
					Command: ast.Command{
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
					Command: ast.Command{
						Name: ast.Word("cmd"),
						Args: []ast.Expression{ast.Var("var")},
						Redirections: []ast.Redirection{
							{Src: "1", Method: ">", Dst: ast.Word("foo")},
							{Src: "3", Method: "<<<", Dst: ast.Word("foo bar")},
						},
					},
					Stderr: false,
				},
				{
					Command: ast.Command{
						Name: ast.Word("cmd2"),
						Args: []ast.Expression{ast.Word("foo bar baz")},
						Redirections: []ast.Redirection{
							{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
						},
					},
					Stderr: true,
				},
			},
		},
	}},
	{` cmd || cmd2 || cmd3 || cmd4`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
					Left:     ast.Command{Name: ast.Word("cmd")},
					Operator: "||",
					Right:    ast.Command{Name: ast.Word("cmd2")},
				},
				Operator: "||",
				Right:    ast.Command{Name: ast.Word("cmd3")},
			},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd4")},
		},
	}},
	{` cmd||cmd2||cmd3||cmd4`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
					Left:     ast.Command{Name: ast.Word("cmd")},
					Operator: "||",
					Right:    ast.Command{Name: ast.Word("cmd2")},
				},
				Operator: "||",
				Right:    ast.Command{Name: ast.Word("cmd3")},
			},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd4")},
		},
	}},
	{` cmd || cmd2 && cmd3 || cmd4 && cmd5`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
					Left: ast.ConditionalCommand{
						Left:     ast.Command{Name: ast.Word("cmd")},
						Operator: "||",
						Right:    ast.Command{Name: ast.Word("cmd2")},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				Operator: "||",
				Right:    ast.Command{Name: ast.Word("cmd4")},
			},
			Operator: "&&",
			Right:    ast.Command{Name: ast.Word("cmd5")},
		},
	}},
	{"cmd || \n\t\n cmd2 ", ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
	}},

	{" cmd || cmd2; cmd3 || cmd4\n", ast.Script{
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd")},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd2")},
		},
		ast.ConditionalCommand{
			Left:     ast.Command{Name: ast.Word("cmd3")},
			Operator: "||",
			Right:    ast.Command{Name: ast.Word("cmd4")},
		},
	}},

	{" cmd || cmd2 || cmd3 & cmd", ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.ConditionalCommand{
				Left: ast.ConditionalCommand{
					Left:     ast.Command{Name: ast.Word("cmd")},
					Operator: "||",
					Right:    ast.Command{Name: ast.Word("cmd2")},
				},
				Operator: "||",
				Right:    ast.Command{Name: ast.Word("cmd3")},
			},
		},
		ast.Command{Name: ast.Word("cmd")},
	}},
}

var logicalCommandsErrorHandlingCases = []errorHandlingTestCase{
	{`cmd &&`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 7)"},
	{`cmd ||`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 7)"},
	{`cmd || cmd && cmd ||`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 21)"},
	{`cmd && cmd || cmd &&`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 21)"},
}

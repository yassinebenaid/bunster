package parser_test

import "github.com/yassinebenaid/chrollo/ast"

var ifCommandTests = []testCase{
	{`if cmd; then cmd2; fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},
	{`if
		cmd;
	 then
		cmd2;
	fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},
	{`if
		cmd1 | cmd2 && cmd3
	 then
		echo 'baz'
	fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
			},
		},
	}},
	{`if
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt;
	then
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt &
	fi;`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.List{
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
			},
			Body: []ast.Statement{
				ast.BackgroundConstruction{
					Statement: ast.List{
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
				},
			},
		},
	}},
	{`if cmd; then echo "foo"; fi & if cmd; then cmd; fi & cmd`, ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.If{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
		},
		ast.BackgroundConstruction{
			Statement: ast.If{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
		ast.Command{Name: ast.Word("cmd")},
	}},
	{`if cmd; then echo "foo"; fi | if cmd; then echo "foo"; fi |& if cmd; then echo "foo"; fi `, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{
				Command: ast.If{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			ast.PipelineCommand{
				Command: ast.If{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			ast.PipelineCommand{
				Stderr: true,
				Command: ast.If{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
		},
	}},
	{`if cmd; then echo "foo"; fi && if cmd; then echo "foo"; fi;`, ast.Script{
		ast.List{
			Left: ast.If{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
			Operator: "&&",
			Right: ast.If{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
		},
	}},
	// Nesting loops
	{`if
		if cmd; then echo "foo"; fi
	then
		if cmd; then echo "foo"; fi
	fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.If{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			Body: []ast.Statement{
				ast.If{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
		},
	}},
	{`if cmd; then echo "foo"; fi >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
			},
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">", Dst: ast.Word("error.txt")},
				{Src: "1", Method: ">&", Dst: ast.Word("3")},
				{Src: "1", Method: ">>", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<<<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">>", Dst: ast.Word("error.txt")},
				{Method: "&>", Dst: ast.Word("all.txt")},
				{Method: "&>>", Dst: ast.Word("all.txt")},
				{Src: "0", Method: "<&", Dst: ast.Word("4")},
				{Src: "5", Method: "<&", Dst: ast.Word("6")},
			},
		},
	}},
	{"if cmd; then cmd2; fi; if cmd; then cmd2; fi \n  if cmd; then cmd2; fi", ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},

	{`if cmd; then cmd2; else cmd3; fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
			Alternate: []ast.Statement{
				ast.Command{Name: ast.Word("cmd3")},
			},
		},
	}},
	{`if
		cmd;
	then
		cmd2;
	else
		cmd3;
	fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
			Alternate: []ast.Statement{
				ast.Command{Name: ast.Word("cmd3")},
			},
		},
	}},
	{`if cmd; then cmd2; elif cmd3; then cmd4; fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
			Elifs: []ast.Elif{
				{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd3")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd4")},
					},
				},
			},
		},
	}},
	{`if cmd; then
		cmd;
	elif cmd; then
		cmd;
	elif cmd; then
		cmd;
	elif cmd; then
		cmd;
	else
		cmd
	fi`, ast.Script{
		ast.If{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Elifs: []ast.Elif{
				{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			Alternate: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
}

var ifCommandErrorHandlingCases = []errorHandlingTestCase{
	{`if`, "syntax error: expected command list after `if`. (line: 1, column: 3)"},
	{`if then`, "syntax error: expected command list after `if`. (line: 1, column: 4)"},
	{`if; then`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 3)"},
	{`if cmd; fi`, "syntax error: expected `then`, found `fi`. (line: 1, column: 9)"},
	{`if fi`, "syntax error: expected command list after `if`. (line: 1, column: 4)"},
	{`if; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 3)"},
	{`if cmd;then fi`, "syntax error: expected command list after `then`. (line: 1, column: 13)"},
	{`if cmd;then cmd`, "syntax error: expected `fi` to close `if` command. (line: 1, column: 16)"},
	{`if cmd;then cmd; fi arg`, "syntax error: unexpected token `arg`. (line: 1, column: 21)"},
	{`if cmd;then cmd; fi <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 37)"},

	{`if cmd; then cmd;elif fi`, "syntax error: expected command list after `elif`. (line: 1, column: 23)"},
	{`if cmd; then cmd;elif cmd; fi`, "syntax error: expected `then`, found `fi`. (line: 1, column: 28)"},
	{`if cmd; then cmd;elif cmd; then fi`, "syntax error: expected command list after `then`. (line: 1, column: 33)"},

	{`if cmd;then cmd;else fi`, "syntax error: expected command list after `else`. (line: 1, column: 22)"},
	{`if cmd;then cmd;else; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 21)"},

	{`if cmd|;then cmd; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 8)"},
	{`if cmd| |;then cmd; fi`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 9)"},
	{`if cmd;then cmd|; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 17)"},
	{`if cmd;then cmd| |; fi`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 18)"},
	{`if cmd; then cmd; elif cmd|;then cmd; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 28)"},
	{`if cmd; then cmd; elif cmd| |;then cmd; fi`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 29)"},
	{`if cmd; then cmd; elif cmd;then cmd|; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 37)"},
	{`if cmd; then cmd; elif cmd;then cmd| |; fi`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 38)"},
	{`if cmd; then cmd; else cmd|; fi`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 28)"},
	{`if cmd; then cmd; else cmd| |; fi`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 29)"},

	{`then`, "syntax error: `then` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
	{`elif`, "syntax error: `elif` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
	{`else`, "syntax error: `else` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
	{`fi`, "syntax error: `fi` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
}

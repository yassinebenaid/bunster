package parser_test

import "github.com/yassinebenaid/bunny/ast"

var conditionalsTests = []testCase{
	{`if cmd; then cmd2; fi`, ast.Script{
		Statements: []ast.Statement{
			ast.If{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
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
		Statements: []ast.Statement{
			ast.If{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
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
		Statements: []ast.Statement{
			ast.If{
				Head: []ast.Statement{
					ast.BinaryConstruction{
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
		Statements: []ast.Statement{
			ast.If{
				Head: []ast.Statement{
					ast.BinaryConstruction{
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
						Statement: ast.BinaryConstruction{
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
		},
	}},
	{`if cmd; then echo "foo"; fi & if cmd; then cmd; fi & cmd`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`if cmd; then echo "foo"; fi | if cmd; then echo "foo"; fi |& if cmd; then echo "foo"; fi `, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`if cmd; then echo "foo"; fi && if cmd; then echo "foo"; fi;`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
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
		},
	}},
	// Nesting loops
	{`if
		if cmd; then echo "foo"; fi
	then
		if cmd; then echo "foo"; fi
	fi`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`if cmd; then echo "foo"; fi >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{"if cmd; then cmd2; fi; if cmd; then cmd2; fi \n  if cmd; then cmd2; fi", ast.Script{
		Statements: []ast.Statement{
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
		},
	}},

	{`if cmd; then cmd2; else cmd3; fi`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`if
		cmd;
	then
		cmd2;
	else
		cmd3;
	fi`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`if cmd; then cmd2; elif cmd3; then cmd4; fi`, ast.Script{
		Statements: []ast.Statement{
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
		Statements: []ast.Statement{
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
		},
	}},
}

var ifErrorHandlingCases = []errorHandlingTestCase{
	{`if`, "syntax error: expected command list after `if`."},
	{`if then`, "syntax error: expected command list after `if`."},
	{`if; then`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
	{`if cmd; fi`, "syntax error: expected `then`, found `fi`."},
	{`if fi`, "syntax error: expected command list after `if`."},
	{`if; fi`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
	{`if cmd;then fi`, "syntax error: expected command list after `then`."},
	{`if cmd;then cmd`, "syntax error: expected `fi` to close `if` command."},
	{`if cmd;then cmd; fi arg`, "syntax error: unexpected token `arg`."},
	{`if cmd;then cmd; fi <in >out <<<etc arg`, "syntax error: unexpected token `arg`."},

	{`if cmd; then cmd;elif fi`, "syntax error: expected command list after `elif`."},
	{`if cmd; then cmd;elif cmd; fi`, "syntax error: expected `then`, found `fi`."},
	{`if cmd; then cmd;elif cmd; then fi`, "syntax error: expected command list after `then`."},

	{`if cmd;then cmd;else fi`, "syntax error: expected command list after `else`."},
	{`if cmd;then cmd;else; fi`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},

	{`if cmd|;then cmd; fi`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
	{`if cmd| |;then cmd; fi`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},

	{`then`, "syntax error: `then` is a reserved keyword, cannot be used a command name."},
	{`elif`, "syntax error: `elif` is a reserved keyword, cannot be used a command name."},
	{`else`, "syntax error: `else` is a reserved keyword, cannot be used a command name."},
	{`fi`, "syntax error: `fi` is a reserved keyword, cannot be used a command name."},
}

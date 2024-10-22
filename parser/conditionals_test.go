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
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.BinaryConstruction{
						Left: ast.Pipeline{
							{
								Command: ast.Command{
									Name: ast.Word("cmd"),
									Args: []ast.Node{ast.Word("arg")},
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
									Args: []ast.Node{ast.Word("foo bar baz")},
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
									Args: []ast.Node{ast.SimpleExpansion("var")},
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
									Args: []ast.Node{ast.Word("foo bar baz")},
									Redirections: []ast.Redirection{
										{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
									},
								},
								Stderr: true,
							},
						},
					},
				},
				Body: []ast.Node{
					ast.BackgroundConstruction{
						Node: ast.BinaryConstruction{
							Left: ast.Pipeline{
								{
									Command: ast.Command{
										Name: ast.Word("cmd"),
										Args: []ast.Node{ast.Word("arg")},
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
										Args: []ast.Node{ast.Word("foo bar baz")},
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
										Args: []ast.Node{ast.SimpleExpansion("var")},
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
										Args: []ast.Node{ast.Word("foo bar baz")},
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
		Statements: []ast.Node{
			ast.BackgroundConstruction{
				Node: ast.If{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
			},
			ast.BackgroundConstruction{
				Node: ast.If{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			ast.Command{Name: ast.Word("cmd")},
		},
	}},
	{`if cmd; then echo "foo"; fi | if cmd; then echo "foo"; fi |& if cmd; then echo "foo"; fi `, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				ast.PipelineCommand{
					Command: ast.If{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				ast.PipelineCommand{
					Command: ast.If{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				ast.PipelineCommand{
					Stderr: true,
					Command: ast.If{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
			},
		},
	}},
	{`if cmd; then echo "foo"; fi && if cmd; then echo "foo"; fi;`, ast.Script{
		Statements: []ast.Node{
			ast.BinaryConstruction{
				Left: ast.If{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
				Operator: "&&",
				Right: ast.If{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
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
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.If{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				Body: []ast.Node{
					ast.If{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
			},
		},
	}},
	{`if cmd; then echo "foo"; fi >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
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
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
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

	{`if cmd; then cmd2; else cmd3; fi`, ast.Script{
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
				Alternate: []ast.Node{
					ast.Command{Name: ast.Word("cmd3")},
				},
			},
		},
	}},
}

var conditionalsErrorHandlingCases = []errorHandlingTestCase{
	{`if`, "syntax error: expected command list after `if`."},
	{`if then`, "syntax error: expected command list after `if`."},
	{`if; then`, "syntax error: invalid command construction."},
	{`if cmd; fi`, "syntax error: expected `then`, found `fi`."},
	{`if fi`, "syntax error: expected command list after `if`."},
	{`if; fi`, "syntax error: invalid command construction."},
	{`if cmd;then fi`, "syntax error: expected command list after `then`."},
	{`if cmd;then cmd`, "syntax error: expected `fi` to close `if` command."},
	{`if cmd;then cmd; fi arg`, "syntax error: unexpected token `arg`."},
	{`if cmd;then cmd; fi <in >out <<<etc arg`, "syntax error: unexpected token `arg`."},
}

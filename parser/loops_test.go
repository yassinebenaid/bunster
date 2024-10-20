package parser_test

import "github.com/yassinebenaid/bunny/ast"

var loopsTests = []testCase{
	//
	// WHILE LOOPS
	//
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
	{`while
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt;
	do
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt &
	done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
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
	{`while cmd; do echo "foo"; done & while cmd; do cmd; done & cmd`, ast.Script{
		Statements: []ast.Node{
			ast.BackgroundConstruction{
				Node: ast.Loop{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
			},
			ast.BackgroundConstruction{
				Node: ast.Loop{
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
	{`while cmd; do echo "foo"; done | while cmd; do echo "foo"; done |& while cmd; do echo "foo"; done `, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				ast.PipelineCommand{
					Command: ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				ast.PipelineCommand{
					Command: ast.Loop{
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
					Command: ast.Loop{
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
	{`while cmd; do echo "foo"; done && while cmd; do echo "foo"; done`, ast.Script{
		Statements: []ast.Node{
			ast.BinaryConstruction{
				Left: ast.Loop{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
				Operator: "&&",
				Right: ast.Loop{
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
	{`while
		while cmd; do echo "foo"; done
	do
		while cmd; do echo "foo"; done
	done`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				Body: []ast.Node{
					ast.Loop{
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
	{`while cmd; do echo "foo"; done >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
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
	{`\while cmd; do echo "foo"; done`, ast.Script{
		Statements: []ast.Node{
			ast.Command{Name: ast.Word("while"), Args: []ast.Node{ast.Word("cmd")}},
			ast.Command{Name: ast.Word("do"), Args: []ast.Node{ast.Word("echo"), ast.Word("foo")}},
			ast.Command{Name: ast.Word("done")},
		},
	}},
	{`while cmd; \do; do echo "foo"; \done; done`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("do")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("done")},
				},
			},
		},
	}},
	{"while cmd; do cmd2; done; while cmd; do cmd2; done \n  while cmd; do cmd2; done", ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
		},
	}},

	//-----------------------------------------------------------
	// UNTIL LOOPS
	//-----------------------------------------------------------
	{`until cmd1; cmd2; cmd3; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
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
	{`until
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
				Negate: true,
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
	{`until
		cmd1 | cmd2 && cmd3
	do
		echo 'baz'
	done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
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
	{`until
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt;
	do
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt &
	done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
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
	{`until cmd; do echo "foo"; done & until cmd; do cmd; done & cmd`, ast.Script{
		Statements: []ast.Node{
			ast.BackgroundConstruction{
				Node: ast.Loop{
					Negate: true,
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
			},
			ast.BackgroundConstruction{
				Node: ast.Loop{
					Negate: true,
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
	{`until cmd; do echo "foo"; done | until cmd; do echo "foo"; done |& until cmd; do echo "foo"; done `, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				ast.PipelineCommand{
					Command: ast.Loop{
						Negate: true,
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				ast.PipelineCommand{
					Command: ast.Loop{
						Negate: true,
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
					Command: ast.Loop{
						Negate: true,
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
	{`until cmd; do echo "foo"; done && until cmd; do echo "foo"; done`, ast.Script{
		Statements: []ast.Node{
			ast.BinaryConstruction{
				Left: ast.Loop{
					Negate: true,
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
				Operator: "&&",
				Right: ast.Loop{
					Negate: true,
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
	{`until
		until cmd; do echo "foo"; done
	do
		until cmd; do echo "foo"; done
	done`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
				Head: []ast.Node{
					ast.Loop{
						Negate: true,
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				Body: []ast.Node{
					ast.Loop{
						Negate: true,
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
	{`until cmd; do echo "foo"; done >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
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
	{`\until cmd; do echo "foo"; done`, ast.Script{
		Statements: []ast.Node{
			ast.Command{Name: ast.Word("until"), Args: []ast.Node{ast.Word("cmd")}},
			ast.Command{Name: ast.Word("do"), Args: []ast.Node{ast.Word("echo"), ast.Word("foo")}},
			ast.Command{Name: ast.Word("done")},
		},
	}},
	{`until cmd; \do; do echo "foo"; \done; done`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("do")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("done")},
				},
			},
		},
	}},
	{"until cmd; do cmd2; done; until cmd; do cmd2; done \n  until cmd; do cmd2; done", ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Negate: true,
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
			ast.Loop{
				Negate: true,
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
			ast.Loop{
				Negate: true,
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
		},
	}},

	//
	// FOR LOOPS
	//
	{`for varname; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		Statements: []ast.Node{
			ast.RangeLoop{
				Var: "varname",
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("bar")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
			},
		},
	}},
	{`
	for varname
	do
		echo "foo"
	 	echo bar;
		echo 'baz';
	done
	`, ast.Script{
		Statements: []ast.Node{
			ast.RangeLoop{
				Var: "varname",
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("bar")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
			},
		},
	}},
	{`for varname do cmd; done`, ast.Script{
		Statements: []ast.Node{
			ast.RangeLoop{
				Var: "varname",
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`for varname do cmd; done &`, ast.Script{
		Statements: []ast.Node{
			ast.BackgroundConstruction{
				Node: ast.RangeLoop{
					Var: "varname",
					Body: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`for varname do cmd; done | cmd |& for varname do cmd; done`, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{
					Command: ast.RangeLoop{
						Var: "varname",
						Body: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
				{Command: ast.Command{Name: ast.Word("cmd")}},
				{
					Stderr: true,
					Command: ast.RangeLoop{
						Var: "varname",
						Body: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`for varname do cmd; done && cmd || for varname do cmd; done`, ast.Script{
		Statements: []ast.Node{
			ast.BinaryConstruction{
				Left: ast.BinaryConstruction{
					Left: ast.RangeLoop{
						Var: "varname",
						Body: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd")},
				},
				Operator: "||",
				Right: ast.RangeLoop{
					Var: "varname",
					Body: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
}

var loopsErrorHandlingCases = []errorHandlingTestCase{
	// WHILE LOOPS
	{`while`, "syntax error: expected command list after `while`."},
	{`while do`, "syntax error: expected command list after `while`."},
	{`while; do`, "syntax error: invalid command construction."},
	{`while cmd; done`, "syntax error: expected `do`, found `done`."},
	{`while done`, "syntax error: expected command list after `while`."},
	{`while; done`, "syntax error: invalid command construction."},
	{`while cmd;do done`, "syntax error: expected command list after `do`."},
	{`while cmd;do cmd`, "syntax error: expected `done` to close `while` loop."},
	{`while cmd;do cmd; done arg`, "syntax error: unexpected token `arg`."},
	{`while cmd;do cmd; done <in >out <<<etc arg`, "syntax error: unexpected token `arg`."},

	// UNTIL LOOPS
	{`until`, "syntax error: expected command list after `until`."},
	{`until do`, "syntax error: expected command list after `until`."},
	{`until; do`, "syntax error: invalid command construction."},
	{`until cmd; done`, "syntax error: expected `do`, found `done`."},
	{`until done`, "syntax error: expected command list after `until`."},
	{`until; done`, "syntax error: invalid command construction."},
	{`until cmd;do done`, "syntax error: expected command list after `do`."},
	{`until cmd;do cmd`, "syntax error: expected `done` to close `until` loop."},
	{`until cmd;do cmd; done arg`, "syntax error: unexpected token `arg`."},
	{`until cmd;do cmd; done <in >out <<<etc arg`, "syntax error: unexpected token `arg`."},

	// FOR LOOPS (over positional arguments)
	{`for`, "syntax error: expected identifier after `for`."},
	{`for do`, "syntax error: expected identifier after `for`."},
	{`for; do`, "syntax error: expected identifier after `for`."},
	{`for var done`, "syntax error: expected `do`, found `done`."},
}

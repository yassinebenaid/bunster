package parser_test

import "github.com/yassinebenaid/bunster/ast"

var loopsTests = []testCase{
	//
	// WHILE LOOPS
	//
	{`while cmd1; cmd2; cmd3; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		ast.Loop{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
				ast.Command{Name: ast.Word("cmd3")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("bar")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
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
		ast.Loop{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
				ast.Command{Name: ast.Word("cmd3")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("bar")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
			},
		},
	}},
	{`while
		cmd1 | cmd2 && cmd3
	do
		echo 'baz'
	done;`, ast.Script{
		ast.Loop{
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
		ast.Loop{
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
							Stderr: true,
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
								Stderr: true,
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
					},
				},
			},
		},
	}},
	{`while cmd; do echo "foo"; done & while cmd; do cmd; done & cmd`, ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.Loop{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
		},
		ast.BackgroundConstruction{
			Statement: ast.Loop{
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
	{`while cmd; do echo "foo"; done | while cmd; do echo "foo"; done |& while cmd; do echo "foo"; done `, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{
				Command: ast.Loop{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			ast.PipelineCommand{
				Command: ast.Loop{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
				Stderr: true,
			},
			ast.PipelineCommand{
				Command: ast.Loop{
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
	{`while cmd; do echo "foo"; done && while cmd; do echo "foo"; done`, ast.Script{
		ast.List{
			Left: ast.Loop{
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
			Operator: "&&",
			Right: ast.Loop{
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
	{`while
		while cmd; do echo "foo"; done
	do
		while cmd; do echo "foo"; done
	done`, ast.Script{
		ast.Loop{
			Head: []ast.Statement{
				ast.Loop{
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			Body: []ast.Statement{
				ast.Loop{
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
	{`while cmd; do echo "foo"; done >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		ast.Loop{
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
	{`while cmd; \do; do echo "foo"; \done; done`, ast.Script{
		ast.Loop{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("do")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("done")},
			},
		},
	}},
	{"while cmd; do cmd2; done; while cmd; do cmd2; done \n  while cmd; do cmd2; done", ast.Script{
		ast.Loop{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
		ast.Loop{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
		ast.Loop{
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},

	//-----------------------------------------------------------
	// UNTIL LOOPS
	//-----------------------------------------------------------
	{`until cmd1; cmd2; cmd3; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
				ast.Command{Name: ast.Word("cmd3")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("bar")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
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
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
				ast.Command{Name: ast.Word("cmd3")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("bar")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
			},
		},
	}},
	{`until
		cmd1 | cmd2 && cmd3
	do
		echo 'baz'
	done;`, ast.Script{
		ast.Loop{
			Negate: true,
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
		ast.Loop{
			Negate: true,
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
							Stderr: true,
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
								Stderr: true,
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
					},
				},
			},
		},
	}},
	{`until cmd; do echo "foo"; done & until cmd; do cmd; done & cmd`, ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.Loop{
				Negate: true,
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
		},
		ast.BackgroundConstruction{
			Statement: ast.Loop{
				Negate: true,
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
	{`until cmd; do echo "foo"; done | until cmd; do echo "foo"; done |& until cmd; do echo "foo"; done `, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{
				Command: ast.Loop{
					Negate: true,
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			ast.PipelineCommand{
				Command: ast.Loop{
					Negate: true,
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
				Stderr: true,
			},
			ast.PipelineCommand{
				Command: ast.Loop{
					Negate: true,
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
	{`until cmd; do echo "foo"; done && until cmd; do echo "foo"; done`, ast.Script{
		ast.List{
			Left: ast.Loop{
				Negate: true,
				Head: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				},
			},
			Operator: "&&",
			Right: ast.Loop{
				Negate: true,
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
	{`until
		until cmd; do echo "foo"; done
	do
		until cmd; do echo "foo"; done
	done`, ast.Script{
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Loop{
					Negate: true,
					Head: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
					},
				},
			},
			Body: []ast.Statement{
				ast.Loop{
					Negate: true,
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
	{`until cmd; do echo "foo"; done >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		ast.Loop{
			Negate: true,
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
	{`until cmd; \do; do echo "foo"; \done; done`, ast.Script{
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("do")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("done")},
			},
		},
	}},
	{"until cmd; do cmd2; done; until cmd; do cmd2; done \n  until cmd; do cmd2; done", ast.Script{
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
		ast.Loop{
			Negate: true,
			Head: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},

	//
	// FOR LOOPS
	//
	{`for varname; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		ast.RangeLoop{
			Var: "varname",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("bar")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
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
		ast.RangeLoop{
			Var: "varname",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("foo")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("bar")}},
				ast.Command{Name: ast.Word("echo"), Args: []ast.Expression{ast.Word("baz")}},
			},
		},
	}},
	{`for varname do cmd; done`, ast.Script{
		ast.RangeLoop{
			Var: "varname",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`for varname do cmd; done &`, ast.Script{
		ast.BackgroundConstruction{
			Statement: ast.RangeLoop{
				Var: "varname",
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`for varname do cmd; done | cmd |& for varname do cmd; done`, ast.Script{
		ast.Pipeline{
			{
				Command: ast.RangeLoop{
					Var: "varname",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: true},
			{
				Command: ast.RangeLoop{
					Var: "varname",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`for varname do cmd; done && cmd || for varname do cmd; done`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.RangeLoop{
					Var: "varname",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right:    ast.Command{Name: ast.Word("cmd")},
			},
			Operator: "||",
			Right: ast.RangeLoop{
				Var: "varname",
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	// Nesting
	{`for varname do for varname do cmd; done; done`, ast.Script{
		ast.RangeLoop{
			Var: "varname",
			Body: []ast.Statement{
				ast.RangeLoop{
					Var: "varname",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`for varname do cmd; done >output.txt <input.txt 2>error.txt >&3 \
		 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		ast.RangeLoop{
			Var: "varname",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
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
	{`for var in foo; do cmd; done`, ast.Script{
		ast.RangeLoop{
			Var: "var",
			Operands: []ast.Expression{
				ast.Word("foo"),
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`for var in foo bar $baz "foobar" 'bar-baz'; do cmd; done`, ast.Script{
		ast.RangeLoop{
			Var: "var",
			Operands: []ast.Expression{
				ast.Word("foo"),
				ast.Word("bar"),
				ast.Var("baz"),
				ast.Word("foobar"),
				ast.Word("bar-baz"),
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`for var in foo bar $baz "foobar" 'bar-baz'
	 do
		cmd
	 done`, ast.Script{
		ast.RangeLoop{
			Var: "var",
			Operands: []ast.Expression{
				ast.Word("foo"),
				ast.Word("bar"),
				ast.Var("baz"),
				ast.Word("foobar"),
				ast.Word("bar-baz"),
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	// ALternative For
	{`for (( a ; b ; c )) do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{
				Init:   ast.Arithmetic{ast.Var("a")},
				Test:   ast.Arithmetic{ast.Var("b")},
				Update: ast.Arithmetic{ast.Var("c")},
			},
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`for ((a;b;c)); do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{
				Init:   ast.Arithmetic{ast.Var("a")},
				Test:   ast.Arithmetic{ast.Var("b")},
				Update: ast.Arithmetic{ast.Var("c")},
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`for (( ; b ; c )) do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{
				Init:   nil,
				Test:   ast.Arithmetic{ast.Var("b")},
				Update: ast.Arithmetic{ast.Var("c")},
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`for (( a ; ; c )) do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{
				Init:   ast.Arithmetic{ast.Var("a")},
				Test:   nil,
				Update: ast.Arithmetic{ast.Var("c")},
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`for (( a ; b ; )) do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{
				Init:   ast.Arithmetic{ast.Var("a")},
				Test:   ast.Arithmetic{ast.Var("b")},
				Update: nil,
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`for ((  ;  ; )) do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`for ((;;)) do cmd;done`, ast.Script{
		ast.For{
			Head: ast.ForHead{},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},

	// Break
	{`while true;do break;done`, ast.Script{
		ast.Loop{
			Head: []ast.Statement{ast.Command{Name: ast.Word("true")}},
			Body: []ast.Statement{ast.Break(1)},
		},
	}},
	{`while true; do
		while true; do
	 		break
		done
	done`, ast.Script{
		ast.Loop{
			Head: []ast.Statement{ast.Command{Name: ast.Word("true")}},
			Body: []ast.Statement{
				ast.Loop{
					Head: []ast.Statement{ast.Command{Name: ast.Word("true")}},
					Body: []ast.Statement{ast.Break(1)},
				},
			},
		},
	}},
	{`until true;do break;done`, ast.Script{
		ast.Loop{
			Negate: true,
			Head:   []ast.Statement{ast.Command{Name: ast.Word("true")}},
			Body:   []ast.Statement{ast.Break(1)},
		},
	}},
	{`until true; do
		until true; do
	 		break
		done
	done`, ast.Script{
		ast.Loop{
			Negate: true,
			Head:   []ast.Statement{ast.Command{Name: ast.Word("true")}},
			Body: []ast.Statement{
				ast.Loop{
					Negate: true,
					Head:   []ast.Statement{ast.Command{Name: ast.Word("true")}},
					Body:   []ast.Statement{ast.Break(1)},
				},
			},
		},
	}},
}

var loopsErrorHandlingCases = []errorHandlingTestCase{
	// WHILE LOOPS
	{`while`, "syntax error: expected command list after `while`. (line: 1, column: 6)"},
	{`while do`, "syntax error: expected command list after `while`. (line: 1, column: 7)"},
	{`while; do`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 6)"},
	{`while cmd; done`, "syntax error: expected `do`, found `done`. (line: 1, column: 12)"},
	{`while done`, "syntax error: expected command list after `while`. (line: 1, column: 7)"},
	{`while; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 6)"},
	{`while cmd;do done`, "syntax error: expected command list after `do`. (line: 1, column: 14)"},
	{`while cmd;do cmd`, "syntax error: expected `done` to close `while` loop. (line: 1, column: 17)"},
	{`while cmd;do cmd; done arg`, "syntax error: unexpected token `arg`. (line: 1, column: 24)"},
	{`while cmd;do cmd; done <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 40)"},
	{`while cmd |;do cmd; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 12)"},
	{`while cmd | &;do cmd; done`, "syntax error: expected a valid command name, found `&`. (line: 1, column: 13)"},
	{`while cmd;do cmd && | ; done`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 21)"},
	{`while cmd;do cmd &&; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 20)"},

	// UNTIL LOOPS
	{`until`, "syntax error: expected command list after `until`. (line: 1, column: 6)"},
	{`until do`, "syntax error: expected command list after `until`. (line: 1, column: 7)"},
	{`until; do`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 6)"},
	{`until cmd; done`, "syntax error: expected `do`, found `done`. (line: 1, column: 12)"},
	{`until done`, "syntax error: expected command list after `until`. (line: 1, column: 7)"},
	{`until; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 6)"},
	{`until cmd;do done`, "syntax error: expected command list after `do`. (line: 1, column: 14)"},
	{`until cmd;do cmd`, "syntax error: expected `done` to close `until` loop. (line: 1, column: 17)"},
	{`until cmd;do cmd; done arg`, "syntax error: unexpected token `arg`. (line: 1, column: 24)"},
	{`until cmd;do cmd; done <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 40)"},
	{`until cmd |;do cmd; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 12)"},
	{`until cmd | &;do cmd; done`, "syntax error: expected a valid command name, found `&`. (line: 1, column: 13)"},
	{`until cmd;do cmd && | ; done`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 21)"},
	{`until cmd;do cmd &&; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 20)"},

	// FOR LOOPS (over positional arguments)
	{`for`, "syntax error: expected identifier after `for`. (line: 1, column: 4)"},
	{"for \n var do done", "syntax error: expected identifier after `for`. (line: 2, column: 0)"},
	{`for do`, "syntax error: expected identifier after `for`. (line: 1, column: 5)"},
	{`for; do`, "syntax error: expected identifier after `for`. (line: 1, column: 4)"},
	{`for var done`, "syntax error: expected `do`, found `done`. (line: 1, column: 9)"},
	{`for var \do done`, "syntax error: expected `do`, found `\\d`. (line: 1, column: 9)"},
	{`for done`, "syntax error: expected identifier after `for`. (line: 1, column: 5)"},
	{`for var do done`, "syntax error: expected command list after `do`. (line: 1, column: 12)"},
	{`for var do cmd `, "syntax error: expected `done` to close `for` loop. (line: 1, column: 16)"},
	{`for var do cmd \done`, "syntax error: expected `done` to close `for` loop. (line: 1, column: 22)"},
	{`for var do cmd; done arg`, "syntax error: unexpected token `arg`. (line: 1, column: 22)"},
	{`for var do cmd; done <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 38)"},
	{`for \var do cmd; done`, "syntax error: expected identifier after `for`. (line: 1, column: 5)"},
	{`for var foo do cmd; done`, "syntax error: expected `do`, found `foo`. (line: 1, column: 9)"},
	{`for invalid-var-name do cmd; done`, "syntax error: expected `do`, found `-`. (line: 1, column: 12)"},
	{`for n in; do cmd; done`, "syntax error: missing operand after `in`. (line: 1, column: 9)"},
	{"for n in foo \n bar; do cmd; done", "syntax error: expected `do`, found `bar`. (line: 2, column: 2)"},
	{"for n in &; do cmd; done", "syntax error: unexpected token `&`. (line: 1, column: 10)"},
	{"for n in foo &; do cmd; done", "syntax error: unexpected token `&`. (line: 1, column: 14)"},
	{`for n do cmd |; done`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 15)"},
	{`for n do cmd | |; done`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 16)"},

	// C like for loops
	{`for (()) do cmd;done`, "syntax error: bad arithmetic expression, unexpected token `)`. (line: 1, column: 7)"},
	{`for ((x)) do cmd;done`, "syntax error: expected a semicolon `;`, found `)`. (line: 1, column: 8)"},
	{`for ((x;y)) do cmd;done`, "syntax error: expected a semicolon `;`, found `)`. (line: 1, column: 10)"},
	{`for ((x;y; w z)) do cmd;done`, "syntax error: expected `))` to close loop head, found `z`. (line: 1, column: 14)"},

	{`do`, "syntax error: `do` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
	{`done`, "syntax error: `done` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
	{`break`, "syntax error: the `break` keyword cannot be used outside loops. (line: 1, column: 1)"},
}

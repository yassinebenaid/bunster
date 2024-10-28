package parser_test

import "github.com/yassinebenaid/bunny/ast"

var caseTests = []testCase{
	{`case foo in bar) cmd; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar) cmd
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar )
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg'
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg';;
		baz)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg';&
		boo)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg';;&
		fab)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg'
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{ast.Word("baz")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";&",
					},
					{
						Patterns: []ast.Expression{ast.Word("boo")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;&",
					},
					{
						Patterns: []ast.Expression{ast.Word("fab")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
					},
				},
			},
		},
	}},
	{`case $foo in
		bar|'foo'|$var ) cmd "arg" arg;;
		bar    |   'foo'   |   $var   ) cmd "arg" arg;;
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
							ast.Var("var"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
							ast.Var("var"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},
	{`case $foo in
		(bar) cmd;;
		(bar | 'foo') cmd;;
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},
	{`case $foo in (bar) cmd;; (bar | 'foo') cmd;; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},
	{`case $foo in bar) cmd;; esac & case $foo in bar) cmd;; esac & cmd`, ast.Script{
		Statements: []ast.Statement{
			ast.BackgroundConstruction{
				Statement: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
			},
			ast.BackgroundConstruction{
				Statement: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
			},
			ast.Command{Name: ast.Word("cmd")},
		},
	}},
	{`case $foo in bar) cmd;; esac | case $foo in bar) cmd;; esac |& cmd`, ast.Script{
		Statements: []ast.Statement{
			ast.Pipeline{
				ast.PipelineCommand{
					Command: ast.Case{
						Word: ast.Var("foo"),
						Cases: []ast.CaseItem{
							{
								Patterns: []ast.Expression{ast.Word("bar")},
								Body: []ast.Statement{
									ast.Command{Name: ast.Word("cmd")},
								},
								Terminator: ";;",
							},
						},
					},
				},
				ast.PipelineCommand{
					Command: ast.Case{
						Word: ast.Var("foo"),
						Cases: []ast.CaseItem{
							{
								Patterns: []ast.Expression{ast.Word("bar")},
								Body: []ast.Statement{
									ast.Command{Name: ast.Word("cmd")},
								},
								Terminator: ";;",
							},
						},
					},
				},
				ast.PipelineCommand{
					Stderr:  true,
					Command: ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`case $foo in bar) cmd;; esac || case $foo in bar) cmd;; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
				Left: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
				Operator: "||",
				Right: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
			},
		},
	}},
	{`case $foo in
		bar)
			case $foo in
				bar)
					cmd;;
			esac;;
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Case{
								Word: ast.Var("foo"),
								Cases: []ast.CaseItem{
									{
										Patterns: []ast.Expression{ast.Word("bar")},
										Body: []ast.Statement{
											ast.Command{Name: ast.Word("cmd")},
										},
										Terminator: ";;",
									},
								},
							},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},

	{` case $foo in
		bar)
			cmd;;
	esac >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
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
	{`
		case $foo in	bar) cmd;;esac
		case $foo in	bar) cmd;;esac; case $foo in	bar) cmd;;esac
	`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns:   []ast.Expression{ast.Word("bar")},
						Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
						Terminator: ";;",
					},
				},
			},
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns:   []ast.Expression{ast.Word("bar")},
						Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
						Terminator: ";;",
					},
				},
			},
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns:   []ast.Expression{ast.Word("bar")},
						Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
						Terminator: ";;",
					},
				},
			},
		},
	}},

	{`case in in bar) cmd;; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("in"),
				Cases: []ast.CaseItem{
					{
						Patterns:   []ast.Expression{ast.Word("bar")},
						Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
						Terminator: ";;",
					},
				},
			},
		},
	}},

	{`case esac in bar) cmd;; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("esac"),
				Cases: []ast.CaseItem{
					{
						Patterns:   []ast.Expression{ast.Word("bar")},
						Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
						Terminator: ";;",
					},
				},
			},
		},
	}},

	{`case esac in esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("esac"),
			},
		},
	}},

	// TODO: see if we must resolve compatibility here or not
	// Inputs: `case esac in bar);; esac`, `case esac in bar);& esac`, `case esac in bar);;& esac`
}

var caseErrorHandlingCases = []errorHandlingTestCase{
	{`case`, "syntax error: incomplete `case` statement, an operand is required after `case`."},
	{`case;`, "syntax error: incomplete `case` statement, an operand is required after `case`."},
	{`case foo;`, "syntax error: expected `in`, found `;`."},
	{`case foo`, "syntax error: expected `in`, found `end of file`."},
	{`case foo in`, "syntax error: expected `esac` to close `case` command."},
	{`case foo in ) esac`, "syntax error: invalid pattern provided, unexpected token `)`."},
	{`case foo in pattern foo esac`, "syntax error: expected `)`, found `foo`."},
}

package parser_test

import "github.com/yassinebenaid/bunster/ast"

var caseTests = []testCase{
	{`case foo in bar) cmd; esac`, ast.Script{
		&ast.Case{
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
	}},
	{`case foo
	in
		bar) cmd
	esac`, ast.Script{
		&ast.Case{
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
	}},
	{`case foo
	in
		bar )
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg'
	esac`, ast.Script{
		&ast.Case{
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
		&ast.Case{
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
	}},
	{`case $foo in
		bar|'foo'|$var ) cmd "arg" arg;;
		bar    |   'foo'   |   $var   ) cmd "arg" arg;;
	esac`, ast.Script{
		&ast.Case{
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
	}},
	{`case $foo in
		(bar) cmd;;
		(bar | 'foo') cmd;;
	esac`, ast.Script{
		&ast.Case{
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
	}},
	{`case $foo in (bar) cmd;; (bar | 'foo') cmd;; esac`, ast.Script{
		&ast.Case{
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
	}},
	{`case $foo in bar) cmd;; esac & case $foo in bar) cmd;; esac & cmd`, ast.Script{
		ast.BackgroundConstruction{
			Statement: &ast.Case{
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
			Statement: &ast.Case{
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
	}},
	{`case $foo in bar) cmd;; esac | case $foo in bar) cmd;; esac |& cmd`, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{
				Command: &ast.Case{
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
				Command: &ast.Case{
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
				Stderr: true,
			},
			ast.PipelineCommand{
				Command: ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`case $foo in bar) cmd;; esac || case $foo in bar) cmd;; esac`, ast.Script{
		ast.List{
			Left: &ast.Case{
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
			Right: &ast.Case{
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
	}},
	{`case $foo in
		bar)
			case $foo in
				bar)
					cmd;;
			esac;;
	esac`, ast.Script{
		&ast.Case{
			Word: ast.Var("foo"),
			Cases: []ast.CaseItem{
				{
					Patterns: []ast.Expression{ast.Word("bar")},
					Body: []ast.Statement{
						&ast.Case{
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
	}},

	{` case $foo in
		bar)
			cmd;;
	esac >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		&ast.Case{
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
	}},
	{`
		case $foo in	bar) cmd;;esac
		case $foo in	bar) cmd;;esac; case $foo in	bar) cmd;;esac
	`, ast.Script{
		&ast.Case{
			Word: ast.Var("foo"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
		&ast.Case{
			Word: ast.Var("foo"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
		&ast.Case{
			Word: ast.Var("foo"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
	}},

	{`case in in bar) cmd;; esac`, ast.Script{
		&ast.Case{
			Word: ast.Word("in"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
	}},

	{`case esac in bar) cmd;; esac`, ast.Script{
		&ast.Case{
			Word: ast.Word("esac"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
	}},

	{`case word in
 		pattern	) cmd; ;;
   		pattern ) cmd; ;&
	 	pattern ) cmd; ;;&
	esac`, ast.Script{
		&ast.Case{
			Word: ast.Word("word"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";&",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;&",
				},
			},
		},
	}},

	{`
	# comment
	case word # comment
	# comment
	# comment
	in # comment
	   # comment
	   # comment
 		(   pattern   ) # comment
		   # comment 
		   # comment 
		   # comment 
		   cmd;   # comment
		   # comment 
		   # comment 
		;;
		   # comment 
		   # comment 
		   # comment 
   		pattern ) cmd; ;& # comment 
		# comment 
		# comment 
		pattern ) cmd; ;;& # comment 
		# comment 
		# comment 
	esac # comment
`, ast.Script{
		&ast.Case{
			Word: ast.Word("word"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";&",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
					Terminator: ";;&",
				},
			},
		},
	}},

	// TODO: see if we must resolve compatibility here or not
	// Inputs: `case esac in bar);; esac`, `case esac in bar);& esac`, `case esac in bar);;& esac`
}

var caseErrorHandlingCases = []errorHandlingTestCase{
	{`case`, "syntax error: incomplete `case` statement, an operand is required after `case`. (line: 1, column: 5)"},
	{`case;`, "syntax error: incomplete `case` statement, an operand is required after `case`. (line: 1, column: 5)"},
	{`case foo;`, "syntax error: expected `in`, found `;`. (line: 1, column: 9)"},
	{`case foo`, "syntax error: expected `in`, found `end of file`. (line: 1, column: 9)"},
	{`case foo in`, "syntax error: expected `esac` to close `case` command. (line: 1, column: 12)"},
	{`case foo in ) esac`, "syntax error: invalid pattern provided, unexpected token `)`. (line: 1, column: 13)"},
	{`case foo in pattern foo esac`, "syntax error: expected `)`, found `foo`. (line: 1, column: 21)"},
	{`case foo in pattern) foo esac`, "syntax error: expected `esac` to close `case` command. (line: 1, column: 31)"},
	{`case foo in pattern) foo;;; esac`, "syntax error: invalid pattern provided, unexpected token `;`. (line: 1, column: 27)"},
	{`case foo in pattern) foo;;;& esac`, "syntax error: invalid pattern provided, unexpected token `;`. (line: 1, column: 27)"},
	{`case foo in foo) cmd;; esac arg`, "syntax error: unexpected token `arg`. (line: 1, column: 29)"},
	{`case foo in foo) cmd;; esac <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 45)"},
	{`case foo in foo) cmd (; esac`, "syntax error: expected `)`, found `;`. (line: 1, column: 23)"},

	{`case foo in esac`, "syntax error: at least one case expected in `case` statement. (line: 1, column: 17)"},
	{`esac`, "syntax error: `esac` is a reserved keyword, cannot be used a command name. (line: 1, column: 1)"},
}

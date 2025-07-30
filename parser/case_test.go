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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 18}, Name: ast.Word("cmd")},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 8}, Name: ast.Word("cmd")},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 4, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 5, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 6, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 4, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 5, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 6, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
					},
					Terminator: ";;",
				},
				{
					Patterns: []ast.Expression{ast.Word("baz")},
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 8, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 9, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 10, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
					},
					Terminator: ";&",
				},
				{
					Patterns: []ast.Expression{ast.Word("boo")},
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 12, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 13, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 14, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
					},
					Terminator: ";;&",
				},
				{
					Patterns: []ast.Expression{ast.Word("fab")},
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 16, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 17, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 18, Col: 4}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 2, Col: 20}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 35}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 2, Col: 9}, Name: ast.Word("cmd")},
					},
					Terminator: ";;",
				},
				{
					Patterns: []ast.Expression{
						ast.Word("bar"),
						ast.Word("foo"),
					},
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 17}, Name: ast.Word("cmd")},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 20}, Name: ast.Word("cmd")},
					},
					Terminator: ";;",
				},
				{
					Patterns: []ast.Expression{
						ast.Word("bar"),
						ast.Word("foo"),
					},
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 40}, Name: ast.Word("cmd")},
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
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")},
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
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 50}, Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
				},
			},
		},
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 63}, Name: ast.Word("cmd")},
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
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")},
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
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 50}, Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
				Stderr: true,
			},
			ast.PipelineCommand{
				Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 64}, Name: ast.Word("cmd")},
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
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")},
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
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 51}, Name: ast.Word("cmd")},
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
										ast.Command{Position: ast.Position{File: "main.sh", Line: 5, Col: 6}, Name: ast.Word("cmd")},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 4}, Name: ast.Word("cmd")},
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
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 2, Col: 21}, Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
		&ast.Case{
			Word: ast.Var("foo"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 21}, Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
			},
		},
		&ast.Case{
			Word: ast.Var("foo"),
			Cases: []ast.CaseItem{
				{
					Patterns:   []ast.Expression{ast.Word("bar")},
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 50}, Name: ast.Word("cmd")}},
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
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd")}},
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
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")}},
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
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 2, Col: 14}, Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 16}, Name: ast.Word("cmd")}},
					Terminator: ";&",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 4, Col: 14}, Name: ast.Word("cmd")}},
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
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 13, Col: 6}, Name: ast.Word("cmd")}},
					Terminator: ";;",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 20, Col: 16}, Name: ast.Word("cmd")}},
					Terminator: ";&",
				},
				{
					Patterns:   []ast.Expression{ast.Word("pattern")},
					Body:       []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 23, Col: 13}, Name: ast.Word("cmd")}},
					Terminator: ";;&",
				},
			},
		},
	}},

	// TODO: see if we must resolve compatibility here or not
	// Inputs: `case esac in bar);; esac`, `case esac in bar);& esac`, `case esac in bar);;& esac`
}

var caseErrorHandlingCases = []errorHandlingTestCase{
	{`case`, "main.sh(1:5): syntax error: incomplete `case` statement, an operand is required after `case`."},
	{`case;`, "main.sh(1:5): syntax error: incomplete `case` statement, an operand is required after `case`."},
	{`case foo;`, "main.sh(1:9): syntax error: expected `in`, found `;`."},
	{`case foo`, "main.sh(1:9): syntax error: expected `in`, found `end of file`."},
	{`case foo in`, "main.sh(1:12): syntax error: expected `esac` to close `case` command."},
	{`case foo in ) esac`, "main.sh(1:13): syntax error: invalid pattern provided, unexpected token `)`."},
	{`case foo in pattern foo esac`, "main.sh(1:21): syntax error: expected `)`, found `foo`."},
	{`case foo in pattern) foo esac`, "main.sh(1:31): syntax error: expected `esac` to close `case` command."},
	{`case foo in pattern) foo;;; esac`, "main.sh(1:27): syntax error: invalid pattern provided, unexpected token `;`."},
	{`case foo in pattern) foo;;;& esac`, "main.sh(1:27): syntax error: invalid pattern provided, unexpected token `;`."},
	{`case foo in foo) cmd;; esac arg`, "main.sh(1:29): syntax error: unexpected token `arg`."},
	{`case foo in foo) cmd;; esac <in >out <<<etc arg`, "main.sh(1:45): syntax error: unexpected token `arg`."},
	{`case foo in foo) cmd (; esac`, "main.sh(1:23): syntax error: expected `)`, found `;`."},

	{`case foo in esac`, "main.sh(1:17): syntax error: at least one case expected in `case` statement."},
	{`esac`, "main.sh(1:1): syntax error: `esac` is a reserved keyword, cannot be used as a command name."},
}

package parser_test

import "github.com/yassinebenaid/bunster/ast"

var groupAndSubshellTests = []testCase{
	{`{ cmd; }`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`{ cmd; cmd; }`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`{cmd;cmd;}`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`{
		cmd

	 	cmd
	}`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`{cmd&cmd&}`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
			},
		},
	}},
	{`{cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3;}`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
			},
		},
	}},

	{`{cmd; cmd;} | {cmd; cmd;}&& {cmd; cmd;}`, ast.Script{
		ast.List{
			Left: ast.Pipeline{
				{Command: ast.Group{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					},
				}},
				{Command: ast.Group{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					},
				}},
			},
			Operator: "&&",
			Right: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`{cmd};}`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd}")},
			},
		},
	}},
	{`{cmd;} >output.txt <input.txt 2>error.txt >&3 \
		 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6 >&- 3<&4-`, ast.Script{
		ast.Group{
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
				{Src: "1", Method: ">&", Close: true},
				{Src: "3", Method: "<&", Dst: ast.Number("4"), Close: true},
			},
		},
	}},
	{`{{cmd;};}`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Group{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`{# comment
		#comment
		#comment
		cmd # comment
		#comment
		cmd2 #comment
		#comment
		} #comment`, ast.Script{
		ast.Group{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},

	{`( cmd )`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`( cmd; cmd )`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`( cmd; cmd; )`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`(cmd;cmd)`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`(
		cmd
	 	cmd
	)`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`(cmd&cmd&)`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
			},
		},
	}},
	{`(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
			},
		},
	}},

	{`(cmd; cmd) | (cmd; cmd)&& (cmd; cmd)`, ast.Script{
		ast.List{
			Left: ast.Pipeline{
				{Command: ast.SubShell{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					},
				}},
				{Command: ast.SubShell{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					},
				}},
			},
			Operator: "&&",
			Right: ast.SubShell{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`(cmd) >output.txt <input.txt 2>error.txt >&3 \
		 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6 >&- 3<&4-`, ast.Script{
		ast.SubShell{
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
				{Src: "1", Method: ">&", Close: true},
				{Src: "3", Method: "<&", Dst: ast.Number("4"), Close: true},
			},
		},
	}},
	{`( (cmd) )`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.SubShell{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`(# comment
		#comment
		#comment
		cmd # comment
		#comment
		#comment
		cmd2 #comment
		#comment
		#comment
		) #comment`, ast.Script{
		ast.SubShell{
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},
}

var groupAndSubshellErrorHandlingCases = []errorHandlingTestCase{
	{`{`, "syntax error: expeceted a command list after `{`. (line: 1, column: 2)"},
	{`{}`, "syntax error: expeceted a command list after `{`. (line: 1, column: 2)"},
	{`{cmd`, "syntax error: expected `}`, found `end of file`. (line: 1, column: 5)"},
	{`{cmd}`, "syntax error: expected `}`, found `end of file`. (line: 1, column: 6)"},
	{`{cmd |;}`, "syntax error: expected a valid command name, found `;`. (line: 1, column: 7)"},
	{`{cmd | |}`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 8)"},

	{`{cmd;} arg`, "syntax error: unexpected token `arg`. (line: 1, column: 8)"},
	{`{cmd;} <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 24)"},

	{`(`, "syntax error: expeceted a command list after `(`. (line: 1, column: 2)"},
	{`()`, "syntax error: expeceted a command list after `(`. (line: 1, column: 2)"},
	{`(cmd`, "syntax error: expected `)`, found `end of file`. (line: 1, column: 5)"},
	{`(cmd |)`, "syntax error: expected a valid command name, found `)`. (line: 1, column: 7)"},
	{`(cmd | |)`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 8)"},

	{`(cmd) arg`, "syntax error: unexpected token `arg`. (line: 1, column: 7)"},
	{`(cmd) <in >out <<<etc arg`, "syntax error: unexpected token `arg`. (line: 1, column: 23)"},
}

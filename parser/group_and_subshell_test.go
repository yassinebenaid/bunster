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
	{`{`, "main.sh(1:2): syntax error: expected a command list after `{`."},
	{`{}`, "main.sh(1:2): syntax error: expected a command list after `{`."},
	{`{cmd`, "main.sh(1:5): syntax error: expected `}`, found `end of file`."},
	{`{cmd}`, "main.sh(1:6): syntax error: expected `}`, found `end of file`."},
	{`{cmd |;}`, "main.sh(1:7): syntax error: expected a valid command name, found `;`."},
	{`{cmd | |}`, "main.sh(1:8): syntax error: expected a valid command name, found `|`."},

	{`{cmd;} arg`, "main.sh(1:8): syntax error: unexpected token `arg`."},
	{`{cmd;} <in >out <<<etc arg`, "main.sh(1:24): syntax error: unexpected token `arg`."},

	{`(`, "main.sh(1:2): syntax error: expected a command list after `(`."},
	{`()`, "main.sh(1:2): syntax error: expected a command list after `(`."},
	{`(cmd`, "main.sh(1:5): syntax error: expected `)`, found `end of file`."},
	{`(cmd |)`, "main.sh(1:7): syntax error: expected a valid command name, found `)`."},
	{`(cmd | |)`, "main.sh(1:8): syntax error: expected a valid command name, found `|`."},

	{`(cmd) arg`, "main.sh(1:7): syntax error: unexpected token `arg`."},
	{`(cmd) <in >out <<<etc arg`, "main.sh(1:23): syntax error: unexpected token `arg`."},
}

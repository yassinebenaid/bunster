package parser_test

import "github.com/yassinebenaid/bunny/ast"

var commandAndProcessSubstitutionTests = []testCase{
	{`$( cmd )`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`$( cmd; cmd )`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`$( cmd; cmd; )`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`$(
		cmd
	 	cmd
	)`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`$(cmd&cmd&)`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
			},
		},
	}},
	{`$(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.ConditionalCommand{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				ast.ConditionalCommand{
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
	{`$(cmd; cmd) arg | $(cmd; cmd)&& $(cmd; cmd)`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.Pipeline{
				{
					Command: ast.Command{
						Name: ast.CommandSubstitution{
							ast.Command{Name: ast.Word("cmd")},
							ast.Command{Name: ast.Word("cmd")},
						},
						Args: []ast.Expression{
							ast.Word("arg"),
						},
					},
				},
				{
					Command: ast.Command{
						Name: ast.CommandSubstitution{
							ast.Command{Name: ast.Word("cmd")},
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
			Operator: "&&",
			Right: ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$( $(cmd) )`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{
					Name: ast.CommandSubstitution{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`$($(cmd))`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{
					Name: ast.CommandSubstitution{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`$(# comment
		#comment

		#comment
		cmd # comment
		#comment

		cmd2

		#comment
		)`, ast.Script{
		ast.Command{
			Name: ast.CommandSubstitution{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd2")},
			},
		},
	}},

	{`<( cmd )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<( cmd; cmd )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<( cmd; cmd; )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<(
		cmd
	 	cmd
	)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<(cmd&cmd&)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
					ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				},
				Direction: '<',
			},
		},
	}},
	{`<(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.ConditionalCommand{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
					ast.ConditionalCommand{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
				},
				Direction: '<',
			},
		},
	}},
	{`<(cmd; cmd) arg | <(cmd; cmd)&& <(cmd; cmd)`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.Pipeline{
				{
					Command: ast.Command{
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
								ast.Command{Name: ast.Word("cmd")},
							},
							Direction: '<',
						},
						Args: []ast.Expression{
							ast.Word("arg"),
						},
					},
				},
				{
					Command: ast.Command{
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
								ast.Command{Name: ast.Word("cmd")},
							},
							Direction: '<',
						},
					},
				},
			},
			Operator: "&&",
			Right: ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					},
					Direction: '<',
				},
			},
		},
	}},
	{`<( <(cmd) )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Direction: '<',
				Body: []ast.Statement{
					ast.Command{
						Name: ast.ProcessSubstitution{
							Direction: '<',
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},
	{`<(<(cmd))`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Direction: '<',
				Body: []ast.Statement{
					ast.Command{
						Name: ast.ProcessSubstitution{
							Direction: '<',
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},

	{`>( cmd )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>( cmd; cmd )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>( cmd; cmd; )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>(
		cmd
	 	cmd
	)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>(cmd&cmd&)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
					ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				},
				Direction: '>',
			},
		},
	}},
	{`>(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.ConditionalCommand{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
					ast.ConditionalCommand{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
				},
				Direction: '>',
			},
		},
	}},
	{`>(cmd; cmd) arg | >(cmd; cmd)&& >(cmd; cmd)`, ast.Script{
		ast.ConditionalCommand{
			Left: ast.Pipeline{
				{
					Command: ast.Command{
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
								ast.Command{Name: ast.Word("cmd")},
							},
							Direction: '>',
						},
						Args: []ast.Expression{
							ast.Word("arg"),
						},
					},
				},
				{
					Command: ast.Command{
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
								ast.Command{Name: ast.Word("cmd")},
							},
							Direction: '>',
						},
					},
				},
			},
			Operator: "&&",
			Right: ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					},
					Direction: '>',
				},
			},
		},
	}},
	{`>( >(cmd) )`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Direction: '>',
				Body: []ast.Statement{
					ast.Command{
						Name: ast.ProcessSubstitution{
							Direction: '>',
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},
	{`>(>(cmd))`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Direction: '>',
				Body: []ast.Statement{
					ast.Command{
						Name: ast.ProcessSubstitution{
							Direction: '>',
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},
}

var CommandAndProcessSubstitutionErrorHandlingCases = []errorHandlingTestCase{
	{`$(`, "syntax error: expeceted a command list after `$(`. (line: 1, column: 3)"},
	{`$()`, "syntax error: expeceted a command list after `$(`. (line: 1, column: 3)"},
	{`$(cmd`, "syntax error: unexpected end of file, expeceted `)`. (line: 1, column: 6)"},
	{`$(cmd |)`, "syntax error: expected a valid command name, found `)`. (line: 1, column: 8)"},
	{`$(cmd | |)`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 9)"},

	{`<(`, "syntax error: expeceted a command list after `<(`. (line: 1, column: 3)"},
	{`<()`, "syntax error: expeceted a command list after `<(`. (line: 1, column: 3)"},
	{`<(cmd`, "syntax error: unexpected end of file, expeceted `)`. (line: 1, column: 6)"},
	{`<(cmd |)`, "syntax error: expected a valid command name, found `)`. (line: 1, column: 8)"},
	{`<(cmd | |)`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 9)"},

	{`>(`, "syntax error: expeceted a command list after `>(`. (line: 1, column: 3)"},
	{`>()`, "syntax error: expeceted a command list after `>(`. (line: 1, column: 3)"},
	{`>(cmd`, "syntax error: unexpected end of file, expeceted `)`. (line: 1, column: 6)"},
	{`>(cmd |)`, "syntax error: expected a valid command name, found `)`. (line: 1, column: 8)"},
	{`>(cmd | |)`, "syntax error: expected a valid command name, found `|`. (line: 1, column: 9)"},
}

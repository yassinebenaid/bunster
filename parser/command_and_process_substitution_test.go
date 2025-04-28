package parser_test

import "github.com/yassinebenaid/bunster/ast"

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
	{`$(cmd; cmd) arg | $(cmd; cmd)&& $(cmd; cmd)`, ast.Script{
		ast.List{
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
				Direction: '<',
			},
		},
	}},
	{`<(cmd; cmd) arg | <(cmd; cmd)&& <(cmd; cmd)`, ast.Script{
		ast.List{
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
				Direction: '>',
			},
		},
	}},
	{`>(cmd; cmd) arg | >(cmd; cmd)&& >(cmd; cmd)`, ast.Script{
		ast.List{
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
	{`<(# comment
		#comment

		#comment
		cmd # comment
		#comment

		cmd2

		#comment
		)`, ast.Script{
		ast.Command{
			Name: ast.ProcessSubstitution{
				Direction: '<',
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
		},
	}},
}

var CommandAndProcessSubstitutionErrorHandlingCases = []errorHandlingTestCase{
	{`$(`, "main.sh(1:3): syntax error: expected a command list after `$(`."},
	{`()`, "main.sh(1:2): syntax error: expected a command list after `(`."},
	{`$(cmd`, "main.sh(1:6): syntax error: unexpected end of file, expected `)`."},
	{`$(cmd |)`, "main.sh(1:8): syntax error: expected a valid command name, found `)`."},
	{`$(cmd | |)`, "main.sh(1:9): syntax error: expected a valid command name, found `|`."},

	{`<(`, "main.sh(1:3): syntax error: expected a command list after `<(`."},
	{`<()`, "main.sh(1:3): syntax error: expected a command list after `<(`."},
	{`<(cmd`, "main.sh(1:6): syntax error: unexpected end of file, expected `)`."},
	{`<(cmd |)`, "main.sh(1:8): syntax error: expected a valid command name, found `)`."},
	{`<(cmd | |)`, "main.sh(1:9): syntax error: expected a valid command name, found `|`."},

	{`>(`, "main.sh(1:3): syntax error: expected a command list after `>(`."},
	{`>()`, "main.sh(1:3): syntax error: expected a command list after `>(`."},
	{`>(cmd`, "main.sh(1:6): syntax error: unexpected end of file, expected `)`."},
	{`>(cmd |)`, "main.sh(1:8): syntax error: expected a valid command name, found `)`."},
	{`>(cmd | |)`, "main.sh(1:9): syntax error: expected a valid command name, found `|`."},
}

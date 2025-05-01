package parser_test

import "github.com/yassinebenaid/bunster/ast"

var commandAndProcessSubstitutionTests = []testCase{
	{`$( cmd )`, ast.Script{
		ast.Command{
			Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 4}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`$( cmd; cmd )`, ast.Script{
		ast.Command{
			Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 4}, Name: ast.Word("cmd")},
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`$( cmd; cmd; )`, ast.Script{
		ast.Command{
			Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 4}, Name: ast.Word("cmd")},
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`$(
		cmd
	 	cmd
	)`, ast.Script{
		ast.Command{
			Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 2, Col: 3}, Name: ast.Word("cmd")},
				ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 4}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`$(cmd&cmd&)`, ast.Script{
		ast.Command{
			Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.BackgroundConstruction{Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 3}, Name: ast.Word("cmd")}},
				ast.BackgroundConstruction{Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 7}, Name: ast.Word("cmd")}},
			},
		},
	}},
	{`$(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.Command{
			Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 3}, Name: ast.Word("cmd1")}},
						{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 10}, Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 18}, Name: ast.Word("cmd3")},
				},
				ast.List{
					Left: ast.Pipeline{
						{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 24}, Name: ast.Word("cmd1")}},
						{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 31}, Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 39}, Name: ast.Word("cmd3")},
				},
			},
		},
	}},
	{`$(cmd; cmd) arg | $(cmd; cmd)&& $(cmd; cmd)`, ast.Script{
		ast.List{
			Left: ast.Pipeline{
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.CommandSubstitution{
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 3}, Name: ast.Word("cmd")},
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 8}, Name: ast.Word("cmd")},
						},
						Args: []ast.Expression{
							ast.Word("arg"),
						},
					},
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19},
						Name: ast.CommandSubstitution{
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 21}, Name: ast.Word("cmd")},
							ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 26}, Name: ast.Word("cmd")},
						},
					},
				},
			},
			Operator: "&&",
			Right: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 33},
				Name: ast.CommandSubstitution{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 35}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 40}, Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$( $(cmd) )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 4},
					Name: ast.CommandSubstitution{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 6}, Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`$($(cmd))`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 3},
					Name: ast.CommandSubstitution{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 5}, Name: ast.Word("cmd")},
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
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.CommandSubstitution{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd2")},
			},
		},
	}},

	{`<( cmd )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<( cmd; cmd )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<( cmd; cmd; )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<(
		cmd
	 	cmd
	)`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '<',
			},
		},
	}},
	{`<(cmd&cmd&)`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.BackgroundConstruction{Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")}},
					ast.BackgroundConstruction{Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")}},
				},
				Direction: '<',
			},
		},
	}},
	{`<(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.List{
						Left: ast.Pipeline{
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd1")}},
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd3")},
					},
					ast.List{
						Left: ast.Pipeline{
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd1")}},
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd3")},
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
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
							Direction: '<',
						},
						Args: []ast.Expression{
							ast.Word("arg"),
						},
					},
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
							Direction: '<',
						},
					},
				},
			},
			Operator: "&&",
			Right: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					},
					Direction: '<',
				},
			},
		},
	}},
	{`<( <(cmd) )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Direction: '<',
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Direction: '<',
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},
	{`<(<(cmd))`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Direction: '<',
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Direction: '<',
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},

	{`>( cmd )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>( cmd; cmd )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>( cmd; cmd; )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>(
		cmd
	 	cmd
	)`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
				},
				Direction: '>',
			},
		},
	}},
	{`>(cmd&cmd&)`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.BackgroundConstruction{Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")}},
					ast.BackgroundConstruction{Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")}},
				},
				Direction: '>',
			},
		},
	}},
	{`>(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Body: []ast.Statement{
					ast.List{
						Left: ast.Pipeline{
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd1")}},
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd3")},
					},
					ast.List{
						Left: ast.Pipeline{
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd1")}},
							{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd3")},
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
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
							Direction: '>',
						},
						Args: []ast.Expression{
							ast.Word("arg"),
						},
					},
				},
				{
					Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
							Direction: '>',
						},
					},
				},
			},
			Operator: "&&",
			Right: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					},
					Direction: '>',
				},
			},
		},
	}},
	{`>( >(cmd) )`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Direction: '>',
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Direction: '>',
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
							},
						},
					},
				},
			},
		},
	}},
	{`>(>(cmd))`, ast.Script{
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Direction: '>',
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
						Name: ast.ProcessSubstitution{
							Direction: '>',
							Body: []ast.Statement{
								ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
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
		ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
			Name: ast.ProcessSubstitution{
				Direction: '<',
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")},
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd2")},
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

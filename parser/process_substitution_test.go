package parser_test

import "github.com/yassinebenaid/bunny/ast"

var substitutionTests = []testCase{
	{`$( cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$( cmd; cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$( cmd; cmd; )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$(
		cmd
	 	cmd
	)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$(cmd&cmd&)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
					ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				},
			},
		},
	}},
	{`$(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.BinaryConstruction{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
					ast.BinaryConstruction{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
				},
			},
		},
	}},
	{`$(cmd; cmd) arg | $(cmd; cmd)&& $(cmd; cmd)`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
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
		},
	}},
	{`$( $(cmd) )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{
						Name: ast.CommandSubstitution{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`$($(cmd))`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{
						Name: ast.CommandSubstitution{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},

	{`<( cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Direction: '<',
				},
			},
		},
	}},
	{`<( cmd; cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
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
	{`<( cmd; cmd; )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
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
	{`<(
		cmd
	 	cmd
	)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
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
	{`<(cmd&cmd&)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
						ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
					},
					Direction: '<',
				},
			},
		},
	}},
	{`<(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.BinaryConstruction{
							Left: ast.Pipeline{
								{Command: ast.Command{Name: ast.Word("cmd1")}},
								{Command: ast.Command{Name: ast.Word("cmd2")}},
							},
							Operator: "&&",
							Right:    ast.Command{Name: ast.Word("cmd3")},
						},
						ast.BinaryConstruction{
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
		},
	}},
	{`<(cmd; cmd) arg | <(cmd; cmd)&& <(cmd; cmd)`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
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
		},
	}},
	{`<( <(cmd) )`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`<(<(cmd))`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},

	{`>( cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
					Direction: '>',
				},
			},
		},
	}},
	{`>( cmd; cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
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
	{`>( cmd; cmd; )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
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
	{`>(
		cmd
	 	cmd
	)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
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
	{`>(cmd&cmd&)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
						ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
					},
					Direction: '>',
				},
			},
		},
	}},
	{`>(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.ProcessSubstitution{
					Body: []ast.Statement{
						ast.BinaryConstruction{
							Left: ast.Pipeline{
								{Command: ast.Command{Name: ast.Word("cmd1")}},
								{Command: ast.Command{Name: ast.Word("cmd2")}},
							},
							Operator: "&&",
							Right:    ast.Command{Name: ast.Word("cmd3")},
						},
						ast.BinaryConstruction{
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
		},
	}},
	{`>(cmd; cmd) arg | >(cmd; cmd)&& >(cmd; cmd)`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
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
		},
	}},
	{`>( >(cmd) )`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
	{`>(>(cmd))`, ast.Script{
		Statements: []ast.Statement{
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
		},
	}},
}

var substitutionErrorHandlingCases = []errorHandlingTestCase{
	{`$(`, "syntax error: expeceted a command list after `$(`."},
	{`$()`, "syntax error: expeceted a command list after `$(`."},
	{`$(cmd`, "syntax error: unexpected end of file, expeceted `)`."},
	{`$(cmd |)`, "syntax error: `)` has a special meaning here and cannot be used as a command name."},
	{`$(cmd | |)`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},

	{`<(`, "syntax error: expeceted a command list after `<(`."},
	{`<()`, "syntax error: expeceted a command list after `<(`."},
	{`<(cmd`, "syntax error: unexpected end of file, expeceted `)`."},
	{`<(cmd |)`, "syntax error: `)` has a special meaning here and cannot be used as a command name."},
	{`<(cmd | |)`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},

	{`>(`, "syntax error: expeceted a command list after `>(`."},
	{`>()`, "syntax error: expeceted a command list after `>(`."},
	{`>(cmd`, "syntax error: unexpected end of file, expeceted `)`."},
	{`>(cmd |)`, "syntax error: `)` has a special meaning here and cannot be used as a command name."},
	{`>(cmd | |)`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},
}
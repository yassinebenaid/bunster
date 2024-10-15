package parser_test

import "github.com/yassinebenaid/bunny/ast"

var loopsTests = []testCase{
	{`while cmd1; cmd2; cmd3; do echo "foo"; echo bar; echo 'baz'; done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd1")},
					ast.Command{Name: ast.Word("cmd2")},
					ast.Command{Name: ast.Word("cmd3")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("bar")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
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
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd1")},
					ast.Command{Name: ast.Word("cmd2")},
					ast.Command{Name: ast.Word("cmd3")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("bar")}},
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
			},
		},
	}},
	{`while
		cmd1 | cmd2 && cmd3
	do
		echo 'baz'
	done;`, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.BinaryConstruction{
						Left: ast.Pipeline{
							{Command: ast.Command{Name: ast.Word("cmd1")}},
							{Command: ast.Command{Name: ast.Word("cmd2")}},
						},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd3")},
					},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("baz")}},
				},
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
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.BinaryConstruction{
						Left: ast.Pipeline{
							{
								Command: ast.Command{
									Name: ast.Word("cmd"),
									Args: []ast.Node{ast.Word("arg")},
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
									Args: []ast.Node{ast.Word("foo bar baz")},
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
									Args: []ast.Node{ast.SimpleExpansion("var")},
									Redirections: []ast.Redirection{
										{Src: "1", Method: ">", Dst: ast.Word("foo")},
										{Src: "3", Method: "<<<", Dst: ast.Word("foo bar")},
									},
								},
								Stderr: false,
							},
							{
								Command: ast.Command{
									Name: ast.Word("cmd2"),
									Args: []ast.Node{ast.Word("foo bar baz")},
									Redirections: []ast.Redirection{
										{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
									},
								},
								Stderr: true,
							},
						},
					},
				},
				Body: []ast.Node{
					ast.BackgroundConstruction{
						Node: ast.BinaryConstruction{
							Left: ast.Pipeline{
								{
									Command: ast.Command{
										Name: ast.Word("cmd"),
										Args: []ast.Node{ast.Word("arg")},
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
										Args: []ast.Node{ast.Word("foo bar baz")},
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
										Args: []ast.Node{ast.SimpleExpansion("var")},
										Redirections: []ast.Redirection{
											{Src: "1", Method: ">", Dst: ast.Word("foo")},
											{Src: "3", Method: "<<<", Dst: ast.Word("foo bar")},
										},
									},
									Stderr: false,
								},
								{
									Command: ast.Command{
										Name: ast.Word("cmd2"),
										Args: []ast.Node{ast.Word("foo bar baz")},
										Redirections: []ast.Redirection{
											{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
										},
									},
									Stderr: true,
								},
							},
						},
					},
				},
			},
		},
	}},
	{`while cmd; do echo "foo"; done &`, ast.Script{
		Statements: []ast.Node{
			ast.BackgroundConstruction{
				Node: ast.Loop{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
			},
		},
	}},
	{`while cmd; do echo "foo"; done | while cmd; do echo "foo"; done |& while cmd; do echo "foo"; done `, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				ast.PipelineCommand{
					Command: ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				ast.PipelineCommand{
					Command: ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				ast.PipelineCommand{
					Stderr: true,
					Command: ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
			},
		},
	}},
	{`while cmd; do echo "foo"; done && while cmd; do echo "foo"; done`, ast.Script{
		Statements: []ast.Node{
			ast.BinaryConstruction{
				Left: ast.Loop{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
				Operator: "&&",
				Right: ast.Loop{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
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
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
				Body: []ast.Node{
					ast.Loop{
						Head: []ast.Node{
							ast.Command{Name: ast.Word("cmd")},
						},
						Body: []ast.Node{
							ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
						},
					},
				},
			},
		},
	}},
	{`while cmd; do echo "foo"; done >output.txt <input.txt 2>error.txt >&3 >>output.txt <<<input.txt 2>>error.txt >>&3 `, ast.Script{
		Statements: []ast.Node{
			ast.Loop{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
				},
				Redirections: []ast.Redirection{
					{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
					{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
					{Src: "2", Method: ">", Dst: ast.Word("error.txt")},
					{Src: "1", Method: ">&", Dst: ast.Word("3")},
				},
			},
		},
	}},
}

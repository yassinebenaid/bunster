package parser_test

import "github.com/yassinebenaid/bunny/ast"

var conditionalsTests = []testCase{
	{`if cmd; then cmd2; fi`, ast.Script{
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
		},
	}},
	{`if
		cmd;
	 then
		cmd2;
	fi`, ast.Script{
		Statements: []ast.Node{
			ast.If{
				Head: []ast.Node{
					ast.Command{Name: ast.Word("cmd")},
				},
				Body: []ast.Node{
					ast.Command{Name: ast.Word("cmd2")},
				},
			},
		},
	}},
	{`if
		cmd1 | cmd2 && cmd3
	 then
		echo 'baz'
	fi`, ast.Script{
		Statements: []ast.Node{
			ast.If{
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
	{`if
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt;
	then
		cmd >foo arg <<<"foo bar" |
		cmd2 <input.txt 'foo bar baz' &&
		cmd >foo $var 3<<<"foo bar" |&
		cmd2 "foo bar baz" <input.txt &
	fi;`, ast.Script{
		Statements: []ast.Node{
			ast.If{
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
	{`if cmd; then echo "foo"; fi & if cmd; then cmd; fi & cmd`, ast.Script{
		Statements: []ast.Node{
			ast.BackgroundConstruction{
				Node: ast.If{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("echo"), Args: []ast.Node{ast.Word("foo")}},
					},
				},
			},
			ast.BackgroundConstruction{
				Node: ast.If{
					Head: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
					Body: []ast.Node{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			ast.Command{Name: ast.Word("cmd")},
		},
	}},
}

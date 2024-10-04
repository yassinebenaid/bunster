package parser_test

import "github.com/yassinebenaid/nbs/ast"

var pipesTests = []testCase{
	{` cmd | cmd2 |& cmd3 | cmd4 |& cmd5`, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
				{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd5")}, Stderr: true},
			},
		},
	}},
	{`cmd|cmd2|&cmd3|cmd4|&cmd5`, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
				{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd5")}, Stderr: true},
			},
		},
	}},
	{`cmd arg| cmd2 \|`, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd"), Args: []ast.Node{ast.Word("arg")}}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd2"), Args: []ast.Node{ast.Word("|")}}, Stderr: false},
			},
		},
	}},

	{`cmd arg >foo 2>&1| cmd2 123 |&$var`, ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{
					Command: ast.Command{
						Name: ast.Word("cmd"),
						Args: []ast.Node{ast.Word("arg")},
						Redirections: []ast.Redirection{
							{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("foo")},
							{Src: ast.FileDescriptor("2"), Method: ">&", Dst: ast.Word("1")},
						},
					},
					Stderr: false,
				},
				{
					Command: ast.Command{
						Name: ast.Word("cmd2"),
						Args: []ast.Node{ast.Word("123")},
					},
					Stderr: false,
				},
				{
					Command: ast.Command{
						Name: ast.SimpleExpansion("var"),
					},
					Stderr: true,
				},
			},
		},
	}},
	{"cmd |\n\n\t cmd2", ast.Script{
		Statements: []ast.Node{
			ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
			},
		},
	}},
}

var pipesErrorHandlingCases = []errorHandlingTestCase{
	{`cmd |`, "syntax error: invalid pipeline construction, a command is missing after `|`."},
	{`cmd | foo |&`, "syntax error: invalid pipeline construction, a command is missing after `|&`."},
	{`cmd foo | cmd >foo| |&`, "syntax error: invalid pipeline construction, a command is missing after `|`."},
	{"cmd |\n\n\t <foo", "syntax error: invalid pipeline construction, a command is missing after `|`."},
}

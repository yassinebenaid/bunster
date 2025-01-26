package parser_test

import "github.com/yassinebenaid/bunster/ast"

var pipesTests = []testCase{
	{` cmd | cmd2 |& cmd3 | cmd4 |& cmd5`, ast.Script{

		ast.Pipeline{
			{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
			{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: true},
			{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: false},
			{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: true},
			{Command: ast.Command{Name: ast.Word("cmd5")}, Stderr: false},
		},
	}},
	{`cmd|cmd2|&cmd3|cmd4|&cmd5`, ast.Script{

		ast.Pipeline{
			{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
			{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: true},
			{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: false},
			{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: true},
			{Command: ast.Command{Name: ast.Word("cmd5")}, Stderr: false},
		},
	}},
	{`cmd arg| cmd2 \|`, ast.Script{

		ast.Pipeline{
			{Command: ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg")}}, Stderr: false},
			{Command: ast.Command{Name: ast.Word("cmd2"), Args: []ast.Expression{ast.Word("|")}}, Stderr: false},
		},
	}},

	{`cmd arg >foo 2>&1| cmd2 123 |&$var`, ast.Script{

		ast.Pipeline{
			{
				Command: ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Expression{ast.Word("arg")},
					Redirections: []ast.Redirection{
						{Src: "1", Method: ">", Dst: ast.Word("foo")},
						{Src: "2", Method: ">&", Dst: ast.Word("1")},
					},
				},
				Stderr: false,
			},
			{
				Command: ast.Command{
					Name: ast.Word("cmd2"),
					Args: []ast.Expression{ast.Word("123")},
				},
				Stderr: true,
			},
			{
				Command: ast.Command{
					Name: ast.Var("var"),
				},
				Stderr: false,
			},
		},
	}},
	{"cmd |\n\n\t cmd2", ast.Script{

		ast.Pipeline{
			{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
			{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
		},
	}},
	{"! cmd", ast.Script{
		ast.InvertExitCode{
			Statement: ast.Command{Name: ast.Word("cmd")},
		},
	}},
	{"! ! cmd", ast.Script{
		ast.InvertExitCode{
			Statement: ast.InvertExitCode{
				Statement: ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
}

var pipesErrorHandlingCases = []errorHandlingTestCase{
	{`cmd |`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 6)"},
	{`cmd | foo |&`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 13)"},
	{`cmd foo | cmd >foo| |&`, "syntax error: expected a valid command name, found `|&`. (line: 1, column: 21)"},
	{"cmd |\n\n\t <foo", "syntax error: expected a valid command name, found `<`. (line: 3, column: 3)"},
	{"!\n", "syntax error: expected a valid command name, found `newline`. (line: 2, column: 0)"},
}

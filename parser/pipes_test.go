package parser_test

import "github.com/yassinebenaid/bunster/ast"

var pipesTests = []testCase{
	{` cmd | cmd2 |& cmd3 | cmd4 |& cmd5`, ast.Script{
		ast.Pipeline{
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 2}, Name: ast.Word("cmd")}, Stderr: false},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 8}, Name: ast.Word("cmd2")}, Stderr: true},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 16}, Name: ast.Word("cmd3")}, Stderr: false},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 23}, Name: ast.Word("cmd4")}, Stderr: true},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 31}, Name: ast.Word("cmd5")}, Stderr: false},
		},
	}},
	{`cmd|cmd2|&cmd3|cmd4|&cmd5`, ast.Script{
		ast.Pipeline{
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")}, Stderr: false},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 5}, Name: ast.Word("cmd2")}, Stderr: true},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 11}, Name: ast.Word("cmd3")}, Stderr: false},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 16}, Name: ast.Word("cmd4")}, Stderr: true},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 22}, Name: ast.Word("cmd5")}, Stderr: false},
		},
	}},
	{`cmd arg| cmd2 \|`, ast.Script{
		ast.Pipeline{
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg")}}, Stderr: false},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 10}, Name: ast.Word("cmd2"), Args: []ast.Expression{ast.Word("|")}}, Stderr: false},
		},
	}},

	{`cmd arg >foo 2>&1| cmd2 123 |&$var`, ast.Script{
		ast.Pipeline{
			{
				Command: ast.Command{
					Position: ast.Position{File: "main.sh", Line: 1, Col: 1},
					Name:     ast.Word("cmd"),
					Args:     []ast.Expression{ast.Word("arg")},
					Redirections: []ast.Redirection{
						{Src: "1", Method: ">", Dst: ast.Word("foo")},
						{Src: "2", Method: ">&", Dst: ast.Word("1")},
					},
				},
				Stderr: false,
			},
			{
				Command: ast.Command{
					Position: ast.Position{File: "main.sh", Line: 1, Col: 20},
					Name:     ast.Word("cmd2"),
					Args:     []ast.Expression{ast.Word("123")},
				},
				Stderr: true,
			},
			{
				Command: ast.Command{
					Position: ast.Position{File: "main.sh", Line: 1, Col: 31},
					Name:     ast.Var("var"),
				},
				Stderr: false,
			},
		},
	}},
	{"cmd |\n\n\t cmd2", ast.Script{
		ast.Pipeline{
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 1}, Name: ast.Word("cmd")}, Stderr: false},
			{Command: ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 3}, Name: ast.Word("cmd2")}, Stderr: false},
		},
	}},
	{"! cmd", ast.Script{
		ast.InvertExitCode{
			Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 3}, Name: ast.Word("cmd")},
		},
	}},
	{"! ! cmd", ast.Script{
		ast.InvertExitCode{
			Statement: ast.InvertExitCode{
				Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 5}, Name: ast.Word("cmd")},
			},
		},
	}},
	{"!!cmd", ast.Script{
		ast.InvertExitCode{
			Statement: ast.InvertExitCode{
				Statement: ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 3}, Name: ast.Word("cmd")},
			},
		},
	}},
}

var pipesErrorHandlingCases = []errorHandlingTestCase{
	{`cmd |`, "main.sh(1:6): syntax error: expected a valid command name, found `end of file`."},
	{`cmd | foo |&`, "main.sh(1:13): syntax error: expected a valid command name, found `end of file`."},
	{`cmd foo | cmd >foo| |&`, "main.sh(1:21): syntax error: expected a valid command name, found `|&`."},
	{"cmd |\n\n\t <foo", "main.sh(3:3): syntax error: expected a valid command name, found `<`."},
	{"!\n", "main.sh(2:0): syntax error: expected a valid command name, found `newline`."},
}

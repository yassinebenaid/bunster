package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
)

var deferTests = []testCase{
	{`defer cmd`, ast.Script{
		ast.Defer{
			Command: ast.Command{Name: ast.Word("cmd")},
		},
	}},
	{`defer cmd arg`, ast.Script{
		ast.Defer{
			Command: ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{ast.Word("arg")},
			},
		},
	}},
	{`defer cmd >output.txt`, ast.Script{
		ast.Defer{
			Command: ast.Command{
				Name: ast.Word("cmd"),
				Redirections: []ast.Redirection{
					{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				},
			},
		},
	}},
	{`defer VAR=value cmd`, ast.Script{
		ast.Defer{
			Command: ast.Command{
				Name: ast.Word("cmd"),
				Env: []ast.Assignement{
					ast.Assignement{Name: "VAR", Value: ast.Word("value")},
				},
			},
		},
	}},
	{`defer { cmd; }`, ast.Script{
		ast.Defer{
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`defer { cmd; } >output.txt`, ast.Script{
		ast.Defer{
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Redirections: []ast.Redirection{
					{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				},
			},
		},
	}},
	{`defer{
		cmd
	}`, ast.Script{
		ast.Defer{
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`defer{cmd;}&&defer{cmd;} || defer{cmd;}`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.Defer{
					Command: ast.Group{
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
				Operator: "&&",
				Right: ast.Defer{
					Command: ast.Group{
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
			Operator: "||",
			Right: ast.Defer{
				Command: ast.Group{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`defer { cmd; } # comment`, ast.Script{
		ast.Defer{
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`defer ( cmd )`, ast.Script{
		ast.Defer{
			Command: ast.SubShell{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`defer ( cmd ) >output.txt`, ast.Script{
		ast.Defer{
			Command: ast.SubShell{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
				Redirections: []ast.Redirection{
					{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				},
			},
		},
	}},
	{`defer(
		cmd
	)`, ast.Script{
		ast.Defer{
			Command: ast.SubShell{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`defer(cmd)&&defer(cmd) || defer(cmd)`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: ast.Defer{
					Command: ast.SubShell{
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
				Operator: "&&",
				Right: ast.Defer{
					Command: ast.SubShell{
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
			Operator: "||",
			Right: ast.Defer{
				Command: ast.SubShell{
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
		},
	}},
	{`defer ( cmd ) # comment`, ast.Script{
		ast.Defer{
			Command: ast.SubShell{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
}

var deferErrorHandlingCases = []errorHandlingTestCase{
	{`defer`, "syntax error: expected a valid command name, found `end of file`. (line: 1, column: 6)"},
	{`defer name=foobar`, "syntax error: expected a simple command, group or subshell after `defer`. (line: 1, column: 18)"},
	{`defer {simple_command;} arg`, "syntax error: unexpected token `arg`. (line: 1, column: 25)"},
}

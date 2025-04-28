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
					{Name: "VAR", Value: ast.Word("value")},
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
	{`defer`, "main.sh(1:6): syntax error: expected a valid command name, found `end of file`."},
	{`defer name=foobar`, "main.sh(1:18): syntax error: expected a simple command, group or subshell after `defer`."},
	{`defer {simple_command;} arg`, "main.sh(1:25): syntax error: unexpected token `arg`."},
}

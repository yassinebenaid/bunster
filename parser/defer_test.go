package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
)

var deferTests = []testCase{
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
	{`defer`, "syntax error: expected a group or subshell after `defer`, found `end of file`. (line: 1, column: 6)"},
	{`defer simple_command`, "syntax error: expected a group or subshell after `defer`, found `simple_command`. (line: 1, column: 7)"},
	{`defer {simple_command;} arg`, "syntax error: unexpected token `arg`. (line: 1, column: 25)"},
}

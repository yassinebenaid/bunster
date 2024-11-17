package parser_test

import (
	"github.com/yassinebenaid/bunny/ast"
)

var functionsTests = []testCase{
	{`foo(){ cmd; }`, ast.Script{
		ast.Function{
			Name: "foo",
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`foo-bar-baz () { cmd; }`, ast.Script{
		ast.Function{
			Name: "foo-bar-baz",
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`"foo-bar-baz" () { cmd; }`, ast.Script{
		ast.Function{
			Name: "foo-bar-baz",
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`"foo"-"bar"-'baz' () { cmd; } `, ast.Script{
		ast.Function{
			Name: "foo-bar-baz",
			Command: ast.Group{
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`foo () { cmd; } >output.txt`, ast.Script{
		ast.Function{
			Name: "foo",
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
}

var functionsErrorHandlingCases = []errorHandlingTestCase{
	{`foo ()`, "syntax error: bad function definition, invalid token `end of file`. (line: 1, column: 7)"},
	{`foo () simple_command`, "syntax error: bad function definition, invalid token `simple_command`. (line: 1, column: 8)"},
	{`$foo () {cmd;}`, "syntax error: invalid function name was supplied. (line: 1, column: 6)"},
}

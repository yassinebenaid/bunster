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
}

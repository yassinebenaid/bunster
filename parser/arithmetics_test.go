package parser_test

import "github.com/yassinebenaid/bunny/ast"

var arithmeticsTests = []testCase{
	{`$((1))`, ast.Script{
		ast.Command{Name: ast.Arithmetic{
			Expr: ast.Word("1"),
		}},
	}},
}

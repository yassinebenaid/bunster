package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterExpansionCases = []testCase{
	{`cmd ${var} `, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Var("var"),
			},
		},
	}}},
}

package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterExpansionCases = []testCase{
	{`cmd ${var} ${var} `, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Var("var"),
				ast.Var("var"),
			},
		},
	}}},
	{`cmd ${var-default} ${var-$default} ${var:-default} ${var:-$default}`, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Name: "var", Default: ast.Word("default"), CheckForNull: false},
				ast.VarOrDefault{Name: "var", Default: ast.Var("default")},
				ast.VarOrDefault{Name: "var", Default: ast.Word("default"), CheckForNull: true},
				ast.VarOrDefault{Name: "var", Default: ast.Var("default"), CheckForNull: true},
			},
		},
	}}},
}

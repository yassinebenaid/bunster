package parser_test

import "github.com/yassinebenaid/bunny/ast"

var arithmeticsTests = []testCase{
	{`$((1)) $(( variable_name )) $(( $VARIABLE_NAME ))`, ast.Script{
		ast.Command{
			Name: ast.Arithmetic{Expr: ast.Number("1")},
			Args: []ast.Expression{
				ast.Arithmetic{Expr: ast.Var("variable_name")},
				ast.Arithmetic{Expr: ast.Var("VARIABLE_NAME")},
			},
		},
	}},
}

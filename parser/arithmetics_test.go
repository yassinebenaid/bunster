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
	{`cmd $(( $((123)) )) $(( ${var} ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					Expr: ast.Arithmetic{Expr: ast.Number("123")},
				},
				ast.Arithmetic{Expr: ast.Var("var")},
			},
		},
	}},
	{`cmd $((1+2-3)))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					Expr: ast.InfixArithmetic{
						Left: ast.InfixArithmetic{
							Left:     ast.Number("1"),
							Operator: "+",
							Right:    ast.Number("2"),
						},
						Operator: "-",
						Right:    ast.Number("3"),
					},
				},
			},
		},
	}},
}

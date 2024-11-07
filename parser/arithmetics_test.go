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
	{`cmd $(( 1+2 - 3 + 4-5)))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					Expr: ast.InfixArithmetic{
						Left: ast.InfixArithmetic{
							Left: ast.InfixArithmetic{
								Left: ast.InfixArithmetic{
									Left:     ast.Number("1"),
									Operator: "+",
									Right:    ast.Number("2"),
								},
								Operator: "-",
								Right:    ast.Number("3"),
							},
							Operator: "+",
							Right:    ast.Number("4"),
						},
						Operator: "-",
						Right:    ast.Number("5"),
					},
				},
			},
		},
	}},
	{`cmd $(( var++ )) $(( var-- )) $(( ++var )) $(( --var )) $(( --var++ )) $(( --var++ + ++var-- - 1 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					Expr: ast.PostIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "++",
					},
				},
				ast.Arithmetic{
					Expr: ast.PostIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "--",
					},
				},
				ast.Arithmetic{
					Expr: ast.PreIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "++",
					},
				},
				ast.Arithmetic{
					Expr: ast.PreIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "--",
					},
				},
				ast.Arithmetic{
					Expr: ast.PreIncDecArithmetic{
						Operand: ast.PostIncDecArithmetic{
							Operand:  ast.Var("var"),
							Operator: "++",
						},
						Operator: "--",
					},
				},
				ast.Arithmetic{
					Expr: ast.InfixArithmetic{
						Left: ast.InfixArithmetic{
							Left: ast.PreIncDecArithmetic{
								Operand: ast.PostIncDecArithmetic{
									Operand:  ast.Var("var"),
									Operator: "++",
								},
								Operator: "--",
							},
							Operator: "+",
							Right: ast.PreIncDecArithmetic{
								Operand: ast.PostIncDecArithmetic{
									Operand:  ast.Var("var"),
									Operator: "--",
								},
								Operator: "++",
							},
						},
						Operator: "-",
						Right:    ast.Number("1"),
					},
				},
			},
		},
	}},
	{`cmd $(( +var )) $(( -var ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					Expr: ast.Unary{
						Operand:  ast.Var("var"),
						Operator: "+",
					},
				},
				ast.Arithmetic{
					Expr: ast.Unary{
						Operand:  ast.Var("var"),
						Operator: "-",
					},
				},
			},
		},
	}},
}

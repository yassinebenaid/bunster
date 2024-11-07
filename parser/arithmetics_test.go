package parser_test

import "github.com/yassinebenaid/bunny/ast"

var arithmeticsTests = []testCase{
	{`$((1)) $(( variable_name )) $(( $VARIABLE_NAME ))`, ast.Script{
		ast.Command{
			Name: ast.Arithmetic{ast.Number("1")},
			Args: []ast.Expression{
				ast.Arithmetic{ast.Var("variable_name")},
				ast.Arithmetic{ast.Var("VARIABLE_NAME")},
			},
		},
	}},
	{`cmd $(( $((123)) )) $(( ${var} ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.Arithmetic{ast.Number("123")},
				},
				ast.Arithmetic{ast.Var("var")},
			},
		},
	}},
	{`cmd $(( 1+2 - 3 + 4-5)))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
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
	{`cmd $(( var++ )) $(( var-- )) $(( ++var )) $(( --var ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.PostIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "++",
					},
				},
				ast.Arithmetic{
					ast.PostIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "--",
					},
				},
				ast.Arithmetic{
					ast.PreIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "++",
					},
				},
				ast.Arithmetic{
					ast.PreIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "--",
					},
				},
			},
		},
	}},
	{`cmd $(( +var )) $(( -var )) $(( + - var )) $(( - + var ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.Unary{
						Operand:  ast.Var("var"),
						Operator: "+",
					},
				},
				ast.Arithmetic{
					ast.Unary{
						Operand:  ast.Var("var"),
						Operator: "-",
					},
				},
				ast.Arithmetic{
					ast.Unary{
						Operand: ast.Unary{
							Operand:  ast.Var("var"),
							Operator: "-",
						},
						Operator: "+",
					},
				},
				ast.Arithmetic{
					ast.Unary{
						Operand: ast.Unary{
							Operand:  ast.Var("var"),
							Operator: "+",
						},
						Operator: "-",
					},
				},
			},
		},
	}},
	{`cmd $(( !var )) $(( !$var ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{ast.Negation{Operand: ast.Var("var")}},
				ast.Arithmetic{ast.Negation{Operand: ast.Var("var")}},
			},
		},
	}},
	{`cmd $(( ~var )) $(( ~$var ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{ast.BitFlip{Operand: ast.Var("var")}},
				ast.Arithmetic{ast.BitFlip{Operand: ast.Var("var")}},
			},
		},
	}},
	{`cmd $(( 1 ** 2 )) $(( $var ** var ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "**",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Var("var"),
						Operator: "**",
						Right:    ast.Var("var"),
					},
				},
			},
		},
	}},
	{`cmd $(( 1 * 2 )) $(( 1 / 2 )) $(( 1 % 2 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "*",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "/",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "%",
						Right:    ast.Number("2"),
					},
				},
			},
		},
	}},
}

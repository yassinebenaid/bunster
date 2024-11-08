package parser_test

import (
	"testing"

	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/lexer"
	"github.com/yassinebenaid/bunny/parser"
)

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
	{`cmd $(( var++ )) $(( var-- )) $(( ++var )) $(( --var )) $(( var ++ ))`, ast.Script{
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
				ast.Arithmetic{
					ast.PostIncDecArithmetic{
						Operand:  ast.Var("var"),
						Operator: "++",
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
	{`cmd $(( 1 << 2 )) $(( 1 >> 2 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "<<",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: ">>",
						Right:    ast.Number("2"),
					},
				},
			},
		},
	}},
	{`cmd $(( 1 < 2 )) $(( 1 > 2 )) $(( 1 <= 2 )) $(( 1 >= 2 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "<",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: ">",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "<=",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: ">=",
						Right:    ast.Number("2"),
					},
				},
			},
		},
	}},
	{`cmd $(( 1 == 2 )) $(( 1 != 2 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "==",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "!=",
						Right:    ast.Number("2"),
					},
				},
			},
		},
	}},
	{`cmd $(( 1 & 2 )) $(( 1 ^ 2 )) $(( 1 | 2 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "&",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "^",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "|",
						Right:    ast.Number("2"),
					},
				},
			},
		},
	}},
	{`cmd $(( 1 && 2 )) $(( 1 || 2 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "&&",
						Right:    ast.Number("2"),
					},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{
						Left:     ast.Number("1"),
						Operator: "||",
						Right:    ast.Number("2"),
					},
				},
			},
		},
	}},
	{`cmd $(( 1 ? 2 : 3 )) $(( 1 ? 2 ? 3 : 4 : 5 ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.Conditional{
						Test:      ast.Number("1"),
						Body:      ast.Number("2"),
						Alternate: ast.Number("3"),
					},
				},
				ast.Arithmetic{
					ast.Conditional{
						Test: ast.Number("1"),
						Body: ast.Conditional{
							Test:      ast.Number("2"),
							Body:      ast.Number("3"),
							Alternate: ast.Number("4"),
						},
						Alternate: ast.Number("5"),
					},
				},
			},
		},
	}},
	{`cmd $(( x = y )) $(( x *= y )) $(( x /= y )) $(( x %= y )) $(( x += y )) $(( x -= y )) \
		$(( x <<= y )) $(( x >>= y )) $(( x &= y )) $(( x ^= y )) $(( x |= y ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "*=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "/=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "%=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "+=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "-=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "<<=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: ">>=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "&=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "^=", Right: ast.Var("y")},
				},
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "|=", Right: ast.Var("y")},
				},
			},
		},
	}},
	{`cmd $(( x = y, x + y ,x*y ))`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Arithmetic{
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "=", Right: ast.Var("y")},
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "+", Right: ast.Var("y")},
					ast.InfixArithmetic{Left: ast.Var("x"), Operator: "*", Right: ast.Var("y")},
				},
			},
		},
	}},
}

var arithmeticsPrecedenceTests = []struct {
	input    string
	expected string
}{
	{`$((1))`, `1`},
	{`$((1, 2, 3))`, `1, 2, 3`},
	{`$((a = b *= c /= d %= e += f -= g <<= h >>= i &= j ^= k |= l + 2))`,
		`(a = (b *= (c /= (d %= (e += (f -= (g <<= (h >>= (i &= (j ^= (k |= (l + 2))))))))))))`},
	{`$((a || b || c || d ))`, `(((a || b) || c) || d)`},
	{`$((a && b && c && d ))`, `(((a && b) && c) && d)`},
	{`$((a | b | c | d ))`, `(((a | b) | c) | d)`},
	{`$((a & b & c & d ))`, `(((a & b) & c) & d)`},
	{`$((a ^ b ^ c ^ d ))`, `(((a ^ b) ^ c) ^ d)`},
	{`$((a == b == c != d != e == f ))`, `(((((a == b) == c) != d) != e) == f)`},
	{`$((a <= b >= c < d > e))`, `((((a <= b) >= c) < d) > e)`},
	{`$((a << b >> c))`, `((a << b) >> c)`},
	{`$((a + b - c))`, `((a + b) - c)`},
	{`$((a * b / c % d))`, `(((a * b) / c) % d)`},
	{`$((a ** b ** c))`, `((a ** b) ** c)`},
	{`$((!a ** ~b))`, `(!a ** ~b)`},
}

func TestArithmeticsPrecedence(t *testing.T) {
	for i, tc := range arithmeticsPrecedenceTests {
		p := parser.New(
			lexer.New([]byte(tc.input)),
		)

		script := p.ParseScript()

		if p.Error != nil {
			t.Fatalf("\nCase: %s\nInput: %s\nUnexpected Error: %s\n", dump(i), dump(tc.input), dump(p.Error.Error()))
		}

		if len(script) != 1 {
			t.Fatalf("\nCase: %s\nInput: %s\nExpected a script of one statement, got %s\n", dump(i), dump(tc.input), dump(len(script)))
		}

		cmd, ok := script[0].(ast.Command)
		if !ok {
			t.Fatalf("\nCase: %s\nInput: %s\nExpected a command, got %s\n", dump(i), dump(tc.input), dump(script[0]))
		}

		name, ok := cmd.Name.(ast.Arithmetic)
		if !ok {
			t.Fatalf("\nCase: %s\nInput: %s\nExpected command name to be an arithmetic, got %s\n", dump(i), dump(tc.input), dump(cmd.Name))
		}

		nameStr := name.String()
		if nameStr != tc.expected {
			t.Fatalf("\n Case: %s Input: %s Expected: %s Got %s\n", dump(i), dump(tc.input), dump(tc.expected), dump(nameStr))
		}
	}
}

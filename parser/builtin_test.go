package parser_test

import "github.com/yassinebenaid/bunster/ast"

var builtinTests = []testCase{
	{`exit`, ast.Script{
		ast.Exit{
			Code: ast.Word("0"),
		},
	}},
	{`exit 123`, ast.Script{
		ast.Exit{
			Code: ast.Word("123"),
		},
	}},
	{`exit #comment`, ast.Script{
		ast.Exit{
			Code: ast.Word("0"),
		},
	}},

	{`exit 1 #comment`, ast.Script{
		ast.Exit{
			Code: ast.Word("1"),
		},
	}},

	{`return`, ast.Script{
		&ast.Return{
			Code: ast.Word("0"),
		},
	}},
	{`return 123`, ast.Script{
		&ast.Return{
			Code: ast.Word("123"),
		},
	}},
	{`return #comment`, ast.Script{
		&ast.Return{
			Code: ast.Word("0"),
		},
	}},

	{`return 1 #comment`, ast.Script{
		&ast.Return{
			Code: ast.Word("1"),
		},
	}},

	{`unset var`, ast.Script{
		ast.Unset{
			Names: []ast.Expression{
				ast.Word("var"),
			},
		},
	}},
	{`unset var1 $var2 var3`, ast.Script{
		ast.Unset{
			Names: []ast.Expression{
				ast.Word("var1"),
				ast.Var("var2"),
				ast.Word("var3")},
		},
	}},
	{`unset var1 # comment`, ast.Script{
		ast.Unset{
			Names: []ast.Expression{
				ast.Word("var1"),
			},
		},
	}},
	{`unset var1 && unset var1`, ast.Script{
		ast.List{
			Left: ast.Unset{
				Names: []ast.Expression{
					ast.Word("var1"),
				},
			},
			Operator: "&&",
			Right: ast.Unset{
				Names: []ast.Expression{
					ast.Word("var1"),
				},
			},
		},
	}},
	{`
	unset var1
	unset var2
	`, ast.Script{
		ast.Unset{Names: []ast.Expression{ast.Word("var1")}},
		ast.Unset{Names: []ast.Expression{ast.Word("var2")}},
	}},
	{`
	unset -f var1
	unset -v var2
	`, ast.Script{
		ast.Unset{Flag: "-f", Names: []ast.Expression{ast.Word("var1")}},
		ast.Unset{Flag: "-v", Names: []ast.Expression{ast.Word("var2")}},
	}},
}

var builtinsErrorHandlingCases = []errorHandlingTestCase{
	{`exit <foo`, "syntax error: unexpected token `<`. (line: 1, column: 6)"},
	{`return <foo`, "syntax error: unexpected token `<`. (line: 1, column: 8)"},
	{`unset`, "syntax error: unexpected token `end of file`. (line: 1, column: 6)"},
	{`unset -`, "syntax error: expected a valid flag character after `-`, found `end of file`. (line: 1, column: 8)"},
	{`unset -k`, "syntax error: expected a valid flag character after `-`, found `k`. (line: 1, column: 8)"},
}

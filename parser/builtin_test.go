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
}

var builtinsErrorHandlingCases = []errorHandlingTestCase{
	{`exit <foo`, "syntax error: unexpected token `<`. (line: 1, column: 6)"},
}

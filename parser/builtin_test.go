package parser_test

import "github.com/yassinebenaid/bunster/ast"

var builtinTests = []testCase{
	{`exit`, ast.Script{
		ast.Exit("0"),
	}},
	{`exit 123`, ast.Script{
		ast.Exit("123"),
	}},
	{`exit #comment`, ast.Script{
		ast.Exit("0"),
	}},
}

var builtinsErrorHandlingCases = []errorHandlingTestCase{
	{`exit foo`, "syntax error: unexpected token `foo`. (line: 1, column: 6)"},
	{`exit <foo`, "syntax error: unexpected token `<`. (line: 1, column: 6)"},
}

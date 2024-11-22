package parser_test

import (
	"github.com/yassinebenaid/bunny/ast"
)

var conditionalsTests = []testCase{
	{`[[ foo-bar_baz ]]`, ast.Script{
		ast.Test{
			Expr: ast.Word("foo-bar_baz"),
		},
	}},
	{`[[ -a-file ]]`, ast.Script{
		ast.Test{
			Expr: ast.Word("-a-file"),
		},
	}},
	{`[[ "-a" ]]`, ast.Script{
		ast.Test{
			Expr: ast.Word("-a"),
		},
	}},
	{`[[ -a file ]]`, ast.Script{
		ast.Test{
			Expr: ast.UnaryConditional{
				Operator: "-a",
				Operand:  ast.Word("file"),
			},
		},
	}},
}

package parser_test

import (
	"github.com/yassinebenaid/bunny/ast"
)

var conditionalsTests = []testCase{
	{`[[ -a file ]]`, ast.Script{
		ast.Test{
			Expr: ast.UnaryConditional{
				Operator: "-a",
				Operand:  ast.Word("file"),
			},
		},
	}},
}

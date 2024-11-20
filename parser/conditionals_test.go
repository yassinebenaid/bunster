package parser_test

import (
	"github.com/yassinebenaid/bunny/ast"
)

var conditionalsTests = []testCase{
	{`[[ -a file ]]`, ast.Script{
		ast.Test{
			Expression: ast.Unary{
				Operator: "-a",
				Operand:  ast.Word("file"),
			},
		},
	}},
}

package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterAssignmentTests = []testCase{
	{`var=value`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("value")},
		},
	}},
}

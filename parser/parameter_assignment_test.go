package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterAssignmentTests = []testCase{
	{`var=value var=value`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("value")},
			ast.Assignement{Name: "var", Value: ast.Word("value")},
		},
	}},
	{`var=$var`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Var("var")},
		},
	}},
}

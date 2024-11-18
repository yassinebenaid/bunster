package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterAssignmentTests = []testCase{
	{`var=value var2=value2`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("value")},
			ast.Assignement{Name: "var2", Value: ast.Word("value2")},
		},
	}},
	{`var=$var`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Var("var")},
		},
	}},
}

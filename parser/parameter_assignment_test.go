package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterAssignmentTests = []testCase{
	{`var=value var2='value2'    var3="value3"`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("value")},
			ast.Assignement{Name: "var2", Value: ast.Word("value2")},
			ast.Assignement{Name: "var3", Value: ast.Word("value3")},
		},
	}},
	{`var=$var var=${var}`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Var("var")},
			ast.Assignement{Name: "var", Value: ast.Var("var")},
		},
	}},
	{`var= var2=value var3=`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var"},
			ast.Assignement{Name: "var2", Value: ast.Word("value")},
			ast.Assignement{Name: "var3"},
		},
	}},
	{`var= var2=value var3= cmd var=value`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("var=value"),
			},
			Env: []ast.Assignement{
				ast.Assignement{Name: "var"},
				ast.Assignement{Name: "var2", Value: ast.Word("value")},
				ast.Assignement{Name: "var3"},
			},
		},
	}},
}

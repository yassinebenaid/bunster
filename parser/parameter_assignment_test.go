package parser_test

import "github.com/yassinebenaid/bunster/ast"

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
	{`var= var2=value \
		var3=`, ast.Script{
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
				{Name: "var"},
				{Name: "var2", Value: ast.Word("value")},
				{Name: "var3"},
			},
		},
	}},
	{`var=foo # comment`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("foo")},
		},
	}},
	{`var=() var2=(value) var3=('foo' $bar "baz")`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.ArrayLiteral(nil)},
			ast.Assignement{Name: "var2", Value: ast.ArrayLiteral{ast.Word("value")}},
			ast.Assignement{Name: "var3", Value: ast.ArrayLiteral{
				ast.Word("foo"),
				ast.Var("bar"),
				ast.Word("baz"),
			}},
		},
	}},
	{`var=(
		'foo' 
		$bar 
		"baz"
	)`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.ArrayLiteral{
				ast.Word("foo"),
				ast.Var("bar"),
				ast.Word("baz"),
			}},
		},
	}},
	{`var=( # comment
		'foo' # comment
		# comment
		$bar # comment
		# comment
		"baz" # comment
		# comment
	) # comment`, ast.Script{
		ast.ParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.ArrayLiteral{
				ast.Word("foo"),
				ast.Var("bar"),
				ast.Word("baz"),
			}},
		},
	}},

	{`local var=value var2='value2' var3="value3"`, ast.Script{
		ast.LocalParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("value")},
			ast.Assignement{Name: "var2", Value: ast.Word("value2")},
			ast.Assignement{Name: "var3", Value: ast.Word("value3")},
		},
	}},
	{`local var=$var var=${var}`, ast.Script{
		ast.LocalParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Var("var")},
			ast.Assignement{Name: "var", Value: ast.Var("var")},
		},
	}},
	{`local var= var2=value var3=`, ast.Script{
		ast.LocalParameterAssignement{
			ast.Assignement{Name: "var"},
			ast.Assignement{Name: "var2", Value: ast.Word("value")},
			ast.Assignement{Name: "var3"},
		},
	}},
	{`local var var2 var3`, ast.Script{
		ast.LocalParameterAssignement{
			ast.Assignement{Name: "var"},
			ast.Assignement{Name: "var2"},
			ast.Assignement{Name: "var3"},
		},
	}},
	{`local var # comment`, ast.Script{
		ast.LocalParameterAssignement{
			ast.Assignement{Name: "var"},
		},
	}},

	{`export var=value var2='value2' var3="value3"`, ast.Script{
		ast.ExportParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Word("value")},
			ast.Assignement{Name: "var2", Value: ast.Word("value2")},
			ast.Assignement{Name: "var3", Value: ast.Word("value3")},
		},
	}},
	{`export var=$var var=${var}`, ast.Script{
		ast.ExportParameterAssignement{
			ast.Assignement{Name: "var", Value: ast.Var("var")},
			ast.Assignement{Name: "var", Value: ast.Var("var")},
		},
	}},
	{`export var= var2=value var3=`, ast.Script{
		ast.ExportParameterAssignement{
			ast.Assignement{Name: "var"},
			ast.Assignement{Name: "var2", Value: ast.Word("value")},
			ast.Assignement{Name: "var3"},
		},
	}},
	{`export var var2 var3`, ast.Script{
		ast.ExportParameterAssignement{
			ast.Assignement{Name: "var"},
			ast.Assignement{Name: "var2"},
			ast.Assignement{Name: "var3"},
		},
	}},
	{`export var # comment`, ast.Script{
		ast.ExportParameterAssignement{
			ast.Assignement{Name: "var"},
		},
	}},
}

var parameterAssignmentErrorHandlingCases = []errorHandlingTestCase{
	{`var=(`, "syntax error: expected a closing parenthese `)`, found `end of file`. (line: 1, column: 6)"},
	{`local`, "syntax error: unexpected token `end of file`. (line: 1, column: 6)"},
	{`local # variable`, "syntax error: unexpected token `end of file`. (line: 1, column: 17)"},
	{`export`, "syntax error: unexpected token `end of file`. (line: 1, column: 7)"},
	{`export # variable`, "syntax error: unexpected token `end of file`. (line: 1, column: 18)"},
}

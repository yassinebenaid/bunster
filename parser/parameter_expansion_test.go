package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterExpansionCases = []testCase{
	{`cmd ${var} ${var} `, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Var("var"),
				ast.Var("var"),
			},
		},
	}}},
	{`cmd ${var-default} ${var-${default}} ${var- $foo bar "baz" | & ; 2> < }`, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Name: "var", Default: ast.Word("default")},
				ast.VarOrDefault{Name: "var", Default: ast.Var("default")},
				ast.VarOrDefault{
					Name: "var",
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
			},
		},
	}}},
	{`cmd ${var:-default} ${var:-${default}} ${var:- $foo bar "baz" | & ; 2> < }`, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Name: "var", Default: ast.Word("default"), CheckForNull: true},
				ast.VarOrDefault{Name: "var", Default: ast.Var("default"), CheckForNull: true},
				ast.VarOrDefault{
					Name: "var",
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
					CheckForNull: true,
				},
			},
		},
	}}},
	{`cmd ${var:=default} ${var:=${default}} ${var:= $foo bar "baz" | & ; 2> < }`, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrSet{Name: "var", Default: ast.Word("default")},
				ast.VarOrSet{Name: "var", Default: ast.Var("default")},
				ast.VarOrSet{
					Name: "var",
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
			},
		},
	}}},
}

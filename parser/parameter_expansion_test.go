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
	{`cmd ${var-default} ${var-${default}} ${var- $foo bar "baz" | & ; 2> < } ${var-}`, ast.Script{Statements: []ast.Statement{
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
				ast.VarOrDefault{Name: "var"},
			},
		},
	}}},
	{`cmd ${var:-default} ${var:-${default}} ${var:- $foo bar "baz" | & ; 2> < } ${var:-}`, ast.Script{Statements: []ast.Statement{
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
				ast.VarOrDefault{Name: "var", CheckForNull: true},
			},
		},
	}}},
	{`cmd ${var:=default} ${var:=${default}} ${var:= $foo bar "baz" | & ; 2> < } ${var:=}`, ast.Script{Statements: []ast.Statement{
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
				ast.VarOrSet{Name: "var"},
			},
		},
	}}},
	{`cmd ${var:?error} ${var:?${error}} ${var:? $foo bar "baz" | & ; 2> < } ${var:?}`, ast.Script{Statements: []ast.Statement{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrFail{Name: "var", Error: ast.Word("error")},
				ast.VarOrFail{Name: "var", Error: ast.Var("error")},
				ast.VarOrFail{
					Name: "var",
					Error: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrFail{Name: "var"},
			},
		},
	}}},
}

package parser_test

import "github.com/yassinebenaid/bunny/ast"

var caseTests = []testCase{
	{`case foo in bar) cmd; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar) cmd
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
}

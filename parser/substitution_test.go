package parser_test

import "github.com/yassinebenaid/bunny/ast"

var substitutionTests = []testCase{
	{`$( cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$( cmd; cmd )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$( cmd; cmd; )`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`$(
		cmd
	 	cmd
	)`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{
				Name: ast.CommandSubstitution{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	// {`$(cmd&cmd&)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.CommandSubstitution{
	// 			ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
	// 			ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
	// 		},
	// }},
	// {`$(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.CommandSubstitution{
	// 			ast.BinaryConstruction{
	// 				Left: ast.Pipeline{
	// 					{Command: ast.Command{Name: ast.Word("cmd1")}},
	// 					{Command: ast.Command{Name: ast.Word("cmd2")}},
	// 				},
	// 				Operator: "&&",
	// 				Right:    ast.Command{Name: ast.Word("cmd3")},
	// 			},
	// 			ast.BinaryConstruction{
	// 				Left: ast.Pipeline{
	// 					{Command: ast.Command{Name: ast.Word("cmd1")}},
	// 					{Command: ast.Command{Name: ast.Word("cmd2")}},
	// 				},
	// 				Operator: "&&",
	// 				Right:    ast.Command{Name: ast.Word("cmd3")},
	// 			},
	// 		},
	// }},
	// {`$(cmd; cmd) | $(cmd; cmd)&& $(cmd; cmd)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.BinaryConstruction{
	// 			Left: ast.Pipeline{
	// 				{
	// 					Command: ast.Command{
	// 						Name: ast.CommandSubstitution{
	// 							ast.Command{Name: ast.Word("cmd")},
	// 							ast.Command{Name: ast.Word("cmd")},
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Command: ast.Command{
	// 						Name: ast.CommandSubstitution{
	// 							ast.Command{Name: ast.Word("cmd")},
	// 							ast.Command{Name: ast.Word("cmd")},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Operator: "&&",
	// 			Right: ast.Command{
	// 				Name: ast.CommandSubstitution{
	// 					ast.Command{Name: ast.Word("cmd")},
	// 					ast.Command{Name: ast.Word("cmd")},
	// 				},
	// 			},
	// 		},
	// }},
}

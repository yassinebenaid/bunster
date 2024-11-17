package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterExpansionTests = []testCase{
	{`cmd ${var} ${var} `, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Var("var"),
				ast.Var("var"),
			},
		},
	}},
	{`cmd ${#var} ${#var}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarCount{Parameter: "var"},
				ast.VarCount{Parameter: "var"},
			},
		},
	}},
	{`cmd ${var-default} ${var-'default'} ${var-${default}} ${var- $foo bar "baz" | & ; 2> < } ${var-}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Parameter: "var", Default: ast.Word("default")},
				ast.VarOrDefault{Parameter: "var", Default: ast.Word("default")},
				ast.VarOrDefault{Parameter: "var", Default: ast.Var("default")},
				ast.VarOrDefault{
					Parameter: "var",
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrDefault{Parameter: "var"},
			},
		},
	}},
	{`cmd ${var:-default} ${var:-${default}} ${var:- $foo bar "baz" | & ; 2> < } ${var:-}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Parameter: "var", Default: ast.Word("default"), CheckForNull: true},
				ast.VarOrDefault{Parameter: "var", Default: ast.Var("default"), CheckForNull: true},
				ast.VarOrDefault{
					Parameter: "var",
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
					CheckForNull: true,
				},
				ast.VarOrDefault{Parameter: "var", CheckForNull: true},
			},
		},
	}},
	{`cmd ${var:=default} ${var:=${default}} ${var:= $foo bar "baz" | & ; 2> < } ${var:=}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrSet{Parameter: "var", Default: ast.Word("default")},
				ast.VarOrSet{Parameter: "var", Default: ast.Var("default")},
				ast.VarOrSet{
					Parameter: "var",
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrSet{Parameter: "var"},
			},
		},
	}},
	{`cmd ${var:?error} ${var:?${error}} ${var:? $foo bar "baz" | & ; 2> < } ${var:?}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrFail{Parameter: "var", Error: ast.Word("error")},
				ast.VarOrFail{Parameter: "var", Error: ast.Var("error")},
				ast.VarOrFail{
					Parameter: "var",
					Error: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrFail{Parameter: "var"},
			},
		},
	}},
	{`cmd ${var:+alternate} ${var:+${alternate}} ${var:+ $foo bar "baz" | & ; 2> < } ${var:+}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.CheckAndUse{Parameter: "var", Value: ast.Word("alternate")},
				ast.CheckAndUse{Parameter: "var", Value: ast.Var("alternate")},
				ast.CheckAndUse{
					Parameter: "var",
					Value: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.CheckAndUse{Parameter: "var"},
			},
		},
	}},
	{`cmd ${var^pattern} ${var^${pattern}} ${var^ $foo bar "baz" | & ; 2> < } ${var^}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: "var", Operator: "^", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: "var", Operator: "^", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: "var",
					Operator:  "^",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: "var", Operator: "^"},
			},
		},
	}},
	{`cmd ${var^^pattern} ${var^^${pattern}} ${var^^ $foo bar "baz" | & ; 2> < } ${var^^}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: "var", Operator: "^^", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: "var", Operator: "^^", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: "var",
					Operator:  "^^",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: "var", Operator: "^^"},
			},
		},
	}},
	{`cmd ${var,pattern} ${var,${pattern}} ${var, $foo bar "baz" | & ; 2> < } ${var,}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: "var", Operator: ",", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: "var", Operator: ",", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: "var",
					Operator:  ",",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: "var", Operator: ","},
			},
		},
	}},
	{`cmd ${var,,pattern} ${var,,${pattern}} ${var,, $foo bar "baz" | & ; 2> < } ${var,,}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: "var", Operator: ",,", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: "var", Operator: ",,", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: "var",
					Operator:  ",,",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: "var", Operator: ",,"},
			},
		},
	}},
	{`cmd ${var#pattern} ${var#${pattern}} ${var# $foo bar "baz" | & ; 2> < } ${var#}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: "var", Operator: "#", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: "var", Operator: "#", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: "var",
					Operator:  "#",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: "var", Operator: "#"},
			},
		},
	}},
	{`cmd ${var##pattern} ${var##${pattern}} ${var## $foo bar "baz" | & ; 2> < # } ${var##}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: "var", Operator: "##", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: "var", Operator: "##", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: "var",
					Operator:  "##",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: "var", Operator: "##"},
			},
		},
	}},
	{`cmd ${var%pattern} ${var%${pattern}} ${var% $foo bar "baz" | & ; 2> < # } ${var%}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: "var", Operator: "%", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: "var", Operator: "%", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: "var",
					Operator:  "%",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: "var", Operator: "%"},
			},
		},
	}},
	{`cmd ${var%%pattern} ${var%%${pattern}} ${var%% $foo bar "baz" | & ; 2> < # } ${var%%}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: "var", Operator: "%%", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: "var", Operator: "%%", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: "var",
					Operator:  "%%",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: "var", Operator: "%%"},
			},
		},
	}},
	{`
		cmd ${var/pattern/value} ${var/${pattern}/${value}} \
		${var/ $foo bar "baz" | & ; 2> < # / $foo bar "baz" | & ; 2> < #////} \
		${var/pattern/} ${var/pattern} ${var/}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
					Value: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < #////"),
						},
					},
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
				},
			},
		},
	}},
	{`
		cmd ${var//pattern/value} ${var//${pattern}/${value}} \
		${var// $foo bar "baz" | & ; 2> < # / $foo bar "baz" | & ; 2> < #////} \
		${var//pattern/} ${var//pattern} ${var//} ${var///} ${var//////}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
					Value: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < #////"),
						},
					},
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern:   ast.Word("/"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "//",
					Pattern:   ast.Word("/"),
					Value:     ast.Word("///"),
				},
			},
		},
	}},
	{`
		cmd ${var/#pattern/value} ${var/#${pattern}/${value}} \
		${var/# $foo bar "baz" | & ; 2> < # / $foo bar "baz" | & ; 2> < #////} \
		${var/#pattern/} ${var/#pattern} ${var/#}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/#",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/#",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/#",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
					Value: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < #////"),
						},
					},
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/#",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/#",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/#",
				},
			},
		},
	}},
	{`
		cmd ${var/%pattern/value} ${var/%${pattern}/${value}} \
		${var/% $foo bar "baz" | & ; 2> < # / $foo bar "baz" | & ; 2> < #////} \
		${var/%pattern/} ${var/%pattern} ${var/%}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/%",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/%",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/%",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
					Value: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < #////"),
						},
					},
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/%",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/%",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/%",
				},
			},
		},
	}},
	{`cmd ${var@U} ${var@u} ${var@L} ${var@Q} ${var@E} ${var@P} ${var@A} ${var@K} ${var@a} ${var@k}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Transform{Parameter: "var", Operator: "U"},
				ast.Transform{Parameter: "var", Operator: "u"},
				ast.Transform{Parameter: "var", Operator: "L"},
				ast.Transform{Parameter: "var", Operator: "Q"},
				ast.Transform{Parameter: "var", Operator: "E"},
				ast.Transform{Parameter: "var", Operator: "P"},
				ast.Transform{Parameter: "var", Operator: "A"},
				ast.Transform{Parameter: "var", Operator: "K"},
				ast.Transform{Parameter: "var", Operator: "a"},
				ast.Transform{Parameter: "var", Operator: "k"},
			},
		},
	}},
	{`cmd ${var:x:y} ${var:x} `, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Slice{
					Parameter: "var",
					Offset:    ast.Arithmetic{ast.Var("x")},
					Length:    ast.Arithmetic{ast.Var("y")},
				},
				ast.Slice{
					Parameter: "var",
					Offset:    ast.Arithmetic{ast.Var("x")},
				},
			},
		},
	}},
	{`cmd ${var/<(ls)/$(ls)}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndReplace{
					Parameter: "var",
					Operator:  "/",
					Pattern: ast.ProcessSubstitution{
						Direction: 60,
						Body: []ast.Statement{
							ast.Command{
								Name:         ast.Word("ls"),
								Args:         []ast.Expression(nil),
								Redirections: []ast.Redirection(nil),
							},
						},
					},
					Value: ast.CommandSubstitution{
						ast.Command{
							Name:         ast.Word("ls"),
							Args:         []ast.Expression(nil),
							Redirections: []ast.Redirection(nil),
						},
					},
				},
			},
		},
	}},
	// {`cmd ${var[1]}`, ast.Script{
	// 	ast.Command{
	// 		Name: ast.Word("cmd"),
	// 		Args: []ast.Expression{
	// 			ast.MatchAndReplace{
	// 				Parameter: "var",
	// 				Operator:  "/",
	// 				Pattern: ast.ProcessSubstitution{
	// 					Direction: 60,
	// 					Body: []ast.Statement{
	// 						ast.Command{
	// 							Name:         ast.Word("ls"),
	// 							Args:         []ast.Expression(nil),
	// 							Redirections: []ast.Redirection(nil),
	// 						},
	// 					},
	// 				},
	// 				Value: ast.CommandSubstitution{
	// 					ast.Command{
	// 						Name:         ast.Word("ls"),
	// 						Args:         []ast.Expression(nil),
	// 						Redirections: []ast.Redirection(nil),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }},
}

var parameterExpansionErrorHandlingCases = []errorHandlingTestCase{
	{"${", "syntax error: couldn't find a valid parameter name, found `end of file`. (line: 1, column: 3)"},
	{"${}", "syntax error: couldn't find a valid parameter name, found `}`. (line: 1, column: 3)"},
	{"${!", "syntax error: couldn't find a valid parameter name, found `!`. (line: 1, column: 3)"},
	{"${var", "syntax error: expected closing brace `}`, found `end of file`. (line: 1, column: 6)"},
	{"${#var", "syntax error: expected closing brace `}`, found `end of file`. (line: 1, column: 7)"},
	{"${#var:-default}", "syntax error: expected closing brace `}`, found `:-`. (line: 1, column: 7)"},
	{"${var:}", "syntax error: bad arithmetic expression, unexpected token `}`. (line: 1, column: 7)"},
	{"${var:x:}", "syntax error: bad arithmetic expression, unexpected token `}`. (line: 1, column: 9)"},
	{"${var@}", "syntax error: bad substitution operator `}`, possible operators are (U, u, L, Q, E, P, A, K, a, k). (line: 1, column: 7)"},
}

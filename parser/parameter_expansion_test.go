package parser_test

import "github.com/yassinebenaid/bunny/ast"

var parameterExpansionTests = []testCase{
	{`cmd ${var} ${var} `, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Var(ast.Var("var")),
				ast.Var(ast.Var("var")),
			},
		},
	}},
	{`cmd ${#var} ${#var}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarCount{Parameter: ast.Var("var")},
				ast.VarCount{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var-default} ${var-'default'} ${var-${default}} ${var- $foo bar "baz" | & ; 2> < } ${var-}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Word("default")},
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Word("default")},
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Var("default")},
				ast.VarOrDefault{
					Parameter: ast.Var("var"),
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrDefault{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var:-default} ${var:-${default}} ${var:- $foo bar "baz" | & ; 2> < } ${var:-}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Word("default"), CheckForNull: true},
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Var("default"), CheckForNull: true},
				ast.VarOrDefault{
					Parameter: ast.Var("var"),
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
					CheckForNull: true,
				},
				ast.VarOrDefault{Parameter: ast.Var("var"), CheckForNull: true},
			},
		},
	}},
	{`cmd ${var:=default} ${var:=${default}} ${var:= $foo bar "baz" | & ; 2> < } ${var:=}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrSet{Parameter: ast.Var("var"), Default: ast.Word("default")},
				ast.VarOrSet{Parameter: ast.Var("var"), Default: ast.Var("default")},
				ast.VarOrSet{
					Parameter: ast.Var("var"),
					Default: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrSet{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var:?error} ${var:?${error}} ${var:? $foo bar "baz" | & ; 2> < } ${var:?}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrFail{Parameter: ast.Var("var"), Error: ast.Word("error")},
				ast.VarOrFail{Parameter: ast.Var("var"), Error: ast.Var("error")},
				ast.VarOrFail{
					Parameter: ast.Var("var"),
					Error: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.VarOrFail{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var:+alternate} ${var:+${alternate}} ${var:+ $foo bar "baz" | & ; 2> < } ${var:+}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.CheckAndUse{Parameter: ast.Var("var"), Value: ast.Word("alternate")},
				ast.CheckAndUse{Parameter: ast.Var("var"), Value: ast.Var("alternate")},
				ast.CheckAndUse{
					Parameter: ast.Var("var"),
					Value: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.CheckAndUse{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var^pattern} ${var^${pattern}} ${var^ $foo bar "baz" | & ; 2> < } ${var^}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: "^", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: "^", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: ast.Var("var"),
					Operator:  "^",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: "^"},
			},
		},
	}},
	{`cmd ${var^^pattern} ${var^^${pattern}} ${var^^ $foo bar "baz" | & ; 2> < } ${var^^}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: "^^", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: "^^", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: ast.Var("var"),
					Operator:  "^^",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: "^^"},
			},
		},
	}},
	{`cmd ${var,pattern} ${var,${pattern}} ${var, $foo bar "baz" | & ; 2> < } ${var,}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: ",", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: ",", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: ast.Var("var"),
					Operator:  ",",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: ","},
			},
		},
	}},
	{`cmd ${var,,pattern} ${var,,${pattern}} ${var,, $foo bar "baz" | & ; 2> < } ${var,,}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: ",,", Pattern: ast.Word("pattern")},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: ",,", Pattern: ast.Var("pattern")},
				ast.ChangeCase{
					Parameter: ast.Var("var"),
					Operator:  ",,",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.ChangeCase{Parameter: ast.Var("var"), Operator: ",,"},
			},
		},
	}},
	{`cmd ${var#pattern} ${var#${pattern}} ${var# $foo bar "baz" | & ; 2> < } ${var#}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "#", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "#", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: ast.Var("var"),
					Operator:  "#",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "#"},
			},
		},
	}},
	{`cmd ${var##pattern} ${var##${pattern}} ${var## $foo bar "baz" | & ; 2> < # } ${var##}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "##", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "##", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: ast.Var("var"),
					Operator:  "##",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "##"},
			},
		},
	}},
	{`cmd ${var%pattern} ${var%${pattern}} ${var% $foo bar "baz" | & ; 2> < # } ${var%}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "%", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "%", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: ast.Var("var"),
					Operator:  "%",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "%"},
			},
		},
	}},
	{`cmd ${var%%pattern} ${var%%${pattern}} ${var%% $foo bar "baz" | & ; 2> < # } ${var%%}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "%%", Pattern: ast.Word("pattern")},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "%%", Pattern: ast.Var("pattern")},
				ast.MatchAndRemove{
					Parameter: ast.Var("var"),
					Operator:  "%%",
					Pattern: ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word(" "),
							ast.Var("foo"),
							ast.Word(" bar baz | & ; 2> < # "),
						},
					},
				},
				ast.MatchAndRemove{Parameter: ast.Var("var"), Operator: "%%"},
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
					Parameter: ast.Var("var"),
					Operator:  "/",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "/",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "//",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "//",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "//",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "//",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "//",
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "//",
					Pattern:   ast.Word("/"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "/#",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/#",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "/#",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/#",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "/%",
					Pattern:   ast.Word("pattern"),
					Value:     ast.Word("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/%",
					Pattern:   ast.Var("pattern"),
					Value:     ast.Var("value"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
					Operator:  "/%",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/%",
					Pattern:   ast.Word("pattern"),
				},
				ast.MatchAndReplace{
					Parameter: ast.Var("var"),
					Operator:  "/%",
				},
			},
		},
	}},
	{`cmd ${var@U} ${var@u} ${var@L} ${var@Q} ${var@E} ${var@P} ${var@A} ${var@K} ${var@a} ${var@k}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Transform{Parameter: ast.Var("var"), Operator: "U"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "u"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "L"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "Q"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "E"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "P"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "A"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "K"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "a"},
				ast.Transform{Parameter: ast.Var("var"), Operator: "k"},
			},
		},
	}},
	{`cmd ${var:x:y} ${var:x} `, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Slice{
					Parameter: ast.Var("var"),
					Offset:    ast.Arithmetic{ast.Var("x")},
					Length:    ast.Arithmetic{ast.Var("y")},
				},
				ast.Slice{
					Parameter: ast.Var("var"),
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
					Parameter: ast.Var("var"),
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
}

var parameterExpansionErrorHandlingCases = []errorHandlingTestCase{
	{"${", "syntax error: couldn't find a valid parameter name, found `end of file`. (line: 1)"},
	{"${}", "syntax error: couldn't find a valid parameter name, found `}`. (line: 1)"},
	{"${!", "syntax error: couldn't find a valid parameter name, found `!`. (line: 1)"},
	{"${var", "syntax error: expected closing brace `}`, found `end of file`. (line: 1)"},
	{"${#var", "syntax error: expected closing brace `}`, found `end of file`. (line: 1)"},
	{"${#var:-default}", "syntax error: expected closing brace `}`, found `:-`. (line: 1)"},
	{"${var:}", "syntax error: bad arithmetic expression, unexpected token `}`. (line: 1)"},
	{"${var:x:}", "syntax error: bad arithmetic expression, unexpected token `}`. (line: 1)"},
	{"${var@}", "syntax error: bad substitution operator `}`, possible operators are (U, u, L, Q, E, P, A, K, a, k). (line: 1)"},
}

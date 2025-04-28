package parser_test

import "github.com/yassinebenaid/bunster/ast"

var parameterExpansionTests = []testCase{
	{`cmd ${var} ${var[123]} ${123} ${*} ${@}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Var("var"),
				ast.ArrayAccess{
					Name:  "var",
					Index: ast.Arithmetic{ast.Number("123")},
				},
				ast.SpecialVar("123"),
				ast.PositionalSpread{},
				ast.PositionalSpread{},
			},
		},
	}},
	{`cmd ${#var} ${#var} ${#var[123]}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarLength{Parameter: ast.Var("var")},
				ast.VarLength{Parameter: ast.Var("var")},
				ast.VarLength{Parameter: ast.ArrayAccess{
					Name:  "var",
					Index: ast.Arithmetic{ast.Number("123")},
				}},
			},
		},
	}},
	{`cmd ${var-default} ${var-'default'} ${var-${default}} ${var- $foo bar "baz" | & ; 2> < } ${var-}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Word("default"), UnsetOnly: true},
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Word("default"), UnsetOnly: true},
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Var("default"), UnsetOnly: true},
				ast.VarOrDefault{
					Parameter: ast.Var("var"),
					Default: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
					},
					UnsetOnly: true,
				},
				ast.VarOrDefault{Parameter: ast.Var("var"), UnsetOnly: true},
			},
		},
	}},
	{`cmd ${var:-default} ${var:-${default}} ${var:- $foo bar "baz" | & ; 2> < } ${var:-}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Word("default")},
				ast.VarOrDefault{Parameter: ast.Var("var"), Default: ast.Var("default")},
				ast.VarOrDefault{
					Parameter: ast.Var("var"),
					Default: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
					},
				},
				ast.VarOrDefault{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var=default} ${var=${default}} ${var= $foo bar "baz" | & ; 2> < } ${var=}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.VarOrSet{Parameter: ast.Var("var"), Default: ast.Word("default"), UnsetOnly: true},
				ast.VarOrSet{Parameter: ast.Var("var"), Default: ast.Var("default"), UnsetOnly: true},
				ast.VarOrSet{
					Parameter: ast.Var("var"),
					Default: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
					},
					UnsetOnly: true,
				},
				ast.VarOrSet{Parameter: ast.Var("var"), UnsetOnly: true},
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
					Default: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Error: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
					},
				},
				ast.VarOrFail{Parameter: ast.Var("var")},
			},
		},
	}},
	{`cmd ${var+alternate} ${var+${alternate}} ${var+ $foo bar "baz" | & ; 2> < } ${var+}`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.CheckAndUse{Parameter: ast.Var("var"), Value: ast.Word("alternate"), UnsetOnly: true},
				ast.CheckAndUse{Parameter: ast.Var("var"), Value: ast.Var("alternate"), UnsetOnly: true},
				ast.CheckAndUse{
					Parameter: ast.Var("var"),
					Value: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
					}, UnsetOnly: true,
				},
				ast.CheckAndUse{Parameter: ast.Var("var"), UnsetOnly: true},
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
					Value: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
					},
					Value: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < #////"),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
					},
					Value: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < #////"),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
					},
					Value: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < #////"),
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
					Pattern: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < # "),
					},
					Value: ast.UnquotedString{
						ast.Word(" "),
						ast.Var("foo"),
						ast.Word(" bar baz | & ; 2> < #////"),
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
	{"${", "main.sh(1:3): syntax error: couldn't find a valid parameter name, found `end of file`."},
	{"${}", "main.sh(1:3): syntax error: couldn't find a valid parameter name, found `}`."},
	{"${!", "main.sh(1:3): syntax error: couldn't find a valid parameter name, found `!`."},
	{"${var", "main.sh(1:6): syntax error: expected closing brace `}`, found `end of file`."},
	{"${#var", "main.sh(1:7): syntax error: expected closing brace `}`, found `end of file`."},
	{"${#var:-default}", "main.sh(1:7): syntax error: expected closing brace `}`, found `:-`."},
	{"${var:}", "main.sh(1:7): syntax error: bad arithmetic expression, unexpected token `}`."},
	{"${var:x:}", "main.sh(1:9): syntax error: bad arithmetic expression, unexpected token `}`."},
	{"${var@}", "main.sh(1:7): syntax error: bad substitution operator `}`, possible operators are (U, u, L, Q, E, P, A, K, a, k)."},

	{"${var[]}", "main.sh(1:7): syntax error: bad arithmetic expression, unexpected token `]`."},
	{"${var[}", "main.sh(1:7): syntax error: bad arithmetic expression, unexpected token `}`."},
	{"${var[1}", "main.sh(1:8): syntax error: expected a closing bracket `]`, found `}`."},
	{"${1:=foo}", "main.sh(1:4): syntax error: unexpected token `:=`."},
	{"${1=foo}", "main.sh(1:4): syntax error: unexpected token `=`."},
	{"${@:=foo}", "main.sh(1:4): syntax error: unexpected token `:=`."},
	{"${*=foo}", "main.sh(1:3): syntax error: couldn't find a valid parameter name, found `*=`."},
}

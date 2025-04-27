package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
)

var functionsTests = []testCase{
	{`foo(){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo-bar-baz () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo-bar-baz" () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo"-"bar"-'baz' () { cmd; } `, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo () { cmd; } >output.txt`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
			},
		},
	}},
	{`foo (  )
	 {
		cmd
	}`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo(){cmd;}&&foo(){cmd;} || foo(){cmd;}`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name: "foo",
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},

	{`function foo(){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo-bar-baz () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo-bar-baz" () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo"-"bar"-'baz' () { cmd; } `, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo () { cmd; } >output.txt`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
			},
		},
	}},
	{`function foo (  )
	 {
		cmd
	}`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo {
		cmd
	}`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo(){cmd;}&&function foo {cmd;} || function foo(){cmd;}`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name: "foo",
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`foo() { cmd; } # comment`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo() { cmd; } # comment`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`foo()( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo-bar-baz () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo-bar-baz" () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo"-"bar"-'baz' () ( cmd ) `, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo () ( cmd ) >output.txt`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
			},
		},
	}},
	{`foo (  )
	 (
		cmd
	)`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo()(cmd)&&foo()(cmd) || foo()(cmd)`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: &ast.Function{
					Name:     "foo",
					SubShell: true,
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name:     "foo",
					SubShell: true,
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name:     "foo",
				SubShell: true,
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},

	{`function foo()( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo-bar-baz () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo-bar-baz" () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo"-"bar"-'baz' () ( cmd ) `, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo () ( cmd ) >output.txt`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
			},
		},
	}},
	{`function foo (  )
	 (
		cmd
	)`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo 
	(
		cmd
	)`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo()(cmd)&&function foo() (cmd) || function foo()(cmd)`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: &ast.Function{
					Name:     "foo",
					SubShell: true,
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name:     "foo",
					SubShell: true,
					Body: []ast.Statement{
						ast.Command{Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name:     "foo",
				SubShell: true,
				Body: []ast.Statement{
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`foo() ( cmd ) # comment`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo() ( cmd ) # comment`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	// flags
	{`function foo(-a -b -c -e= -f[=] --abc --def --igk= --lmn[=]){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Flags: []ast.Flag{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
				{Name: "e", AcceptsValue: true},
				{Name: "f", AcceptsValue: true, Optional: true},
				{Name: "abc", Long: true},
				{Name: "def", Long: true},
				{Name: "igk", Long: true, AcceptsValue: true},
				{Name: "lmn", Long: true, AcceptsValue: true, Optional: true},
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`function foo(
		-a 
		-b 
		-c 
		-e= 
		-f[=] 
		--abc 
		--def 
		--igk= 
		--lmn[=]
	){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Flags: []ast.Flag{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
				{Name: "e", AcceptsValue: true},
				{Name: "f", AcceptsValue: true, Optional: true},
				{Name: "abc", Long: true},
				{Name: "def", Long: true},
				{Name: "igk", Long: true, AcceptsValue: true},
				{Name: "lmn", Long: true, AcceptsValue: true, Optional: true},
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
	{`function foo(# comment
		# comment
		# comment
		-a# comment
		# comment
		# comment
		-b # comment
		-c 
		-e= 
		-f[=]# comment
		# comment
		--abc 
		--def 
		# comment
		--igk= 
		--lmn[=]
		# comment
		# comment
	){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Flags: []ast.Flag{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
				{Name: "e", AcceptsValue: true},
				{Name: "f", AcceptsValue: true, Optional: true},
				{Name: "abc", Long: true},
				{Name: "def", Long: true},
				{Name: "igk", Long: true, AcceptsValue: true},
				{Name: "lmn", Long: true, AcceptsValue: true, Optional: true},
			},
			Body: []ast.Statement{ast.Command{Name: ast.Word("cmd")}},
		},
	}},
}

var functionsErrorHandlingCases = []errorHandlingTestCase{
	{`foo ()`, "syntax error: function body is expected, found `end of file`. (line: 1, column: 7)"},
	{`foo () simple_command`, "syntax error: function body is expected, found `simple_command`. (line: 1, column: 8)"},
	{`$foo () {cmd;}`, "syntax error: invalid function name was supplied. (line: 1, column: 6)"},
	{`cmd (`, "syntax error: expected `)`, found `end of file`. (line: 1, column: 6)"},
	{`cmd )`, "syntax error: token `)` cannot be placed here. (line: 1, column: 5)"},
	{`cmd arg (`, "syntax error: token `(` cannot be placed here. (line: 1, column: 9)"},
	{`cmd arg )`, "syntax error: token `)` cannot be placed here. (line: 1, column: 9)"},
	{`cmd arg(arg`, "syntax error: token `(` cannot be placed here. (line: 1, column: 8)"},
	{`cmd arg)arg`, "syntax error: token `)` cannot be placed here. (line: 1, column: 8)"},
	{`func() if true; then cmd;fi`, "syntax error: function body is expected to be a group or subshell. (line: 1, column: 28)"},

	{`function`, "syntax error: function name is required. (line: 1, column: 9)"},
	{`function foo ()`, "syntax error: function body is expected, found `end of file`. (line: 1, column: 16)"},
	{`function foo () simple_command`, "syntax error: function body is expected, found `simple_command`. (line: 1, column: 17)"},
	{`function $foo () {cmd;}`, "syntax error: invalid function name was supplied. (line: 1, column: 14)"},
	{`function cmd (`, "syntax error: expected `)`, found `end of file`. (line: 1, column: 15)"},
	{`function cmd )`, "syntax error: function body is expected, found `)`. (line: 1, column: 14)"},
	{`function cmd () function foo() {cmd;}`, "syntax error: function body is expected, found `function`. (line: 1, column: 17)"},
	{`function func() {cmd;} | cat`, "syntax error: unexpected token `|`. (line: 1, column: 24)"},
	{`func() {cmd;} | cat`, "syntax error: unexpected token `|`. (line: 1, column: 15)"},
	{`function func() if true; then cmd;fi`, "syntax error: function body is expected to be a group or subshell. (line: 1, column: 37)"},
}

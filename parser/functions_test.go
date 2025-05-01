package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
)

var functionsTests = []testCase{
	{`foo(){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 8}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo-bar-baz () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 18}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo-bar-baz" () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 20}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo"-"bar"-'baz' () { cmd; } `, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 24}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo () { cmd; } >output.txt`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 10}, Name: ast.Word("cmd")},
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
				ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 3}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo(){cmd;}&&foo(){cmd;} || foo(){cmd;}`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 7}, Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 20}, Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name: "foo",
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 35}, Name: ast.Word("cmd")},
				},
			},
		},
	}},

	{`function foo(){ cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo-bar-baz () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 27}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo-bar-baz" () { cmd; }`, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 29}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo"-"bar"-'baz' () { cmd; } `, ast.Script{
		&ast.Function{
			Name: "foo-bar-baz",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 33}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo () { cmd; } >output.txt`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")},
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
				ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 3}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo {
		cmd
	}`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 2, Col: 3}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo(){cmd;}&&function foo {cmd;} || function foo(){cmd;}`, ast.Script{
		ast.List{
			Left: ast.List{
				Left: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 16}, Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name: "foo",
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 37}, Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name: "foo",
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 61}, Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`foo() { cmd; } # comment`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo() { cmd; } # comment`, ast.Script{
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 18}, Name: ast.Word("cmd")},
			},
		},
	}},

	{`foo()( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 8}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo-bar-baz () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 18}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo-bar-baz" () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 20}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`"foo"-"bar"-'baz' () ( cmd ) `, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 24}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`foo () ( cmd ) >output.txt`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 10}, Name: ast.Word("cmd")},
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
				ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 3}, Name: ast.Word("cmd")},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 7}, Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name:     "foo",
					SubShell: true,
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name:     "foo",
				SubShell: true,
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 33}, Name: ast.Word("cmd")},
				},
			},
		},
	}},

	{`function foo()( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 17}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo-bar-baz () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 27}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo-bar-baz" () ( cmd )`, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 29}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function "foo"-"bar"-'baz' () ( cmd ) `, ast.Script{
		&ast.Function{
			Name:     "foo-bar-baz",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 33}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo () ( cmd ) >output.txt`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 19}, Name: ast.Word("cmd")},
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
				ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 3}, Name: ast.Word("cmd")},
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
				ast.Command{Position: ast.Position{File: "main.sh", Line: 3, Col: 3}, Name: ast.Word("cmd")},
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
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 16}, Name: ast.Word("cmd")},
					},
				},
				Operator: "&&",
				Right: &ast.Function{
					Name:     "foo",
					SubShell: true,
					Body: []ast.Statement{
						ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 38}, Name: ast.Word("cmd")},
					},
				},
			},
			Operator: "||",
			Right: &ast.Function{
				Name:     "foo",
				SubShell: true,
				Body: []ast.Statement{
					ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 61}, Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`foo() ( cmd ) # comment`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 9}, Name: ast.Word("cmd")},
			},
		},
	}},
	{`function foo() ( cmd ) # comment`, ast.Script{
		&ast.Function{
			Name:     "foo",
			SubShell: true,
			Body: []ast.Statement{
				ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 18}, Name: ast.Word("cmd")},
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
			Body: []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 63}, Name: ast.Word("cmd")}},
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
			Body: []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 11, Col: 5}, Name: ast.Word("cmd")}},
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
			Body: []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 19, Col: 5}, Name: ast.Word("cmd")}},
		},
	}},
}

var functionsErrorHandlingCases = []errorHandlingTestCase{
	{`foo ()`, "main.sh(1:7): syntax error: function body is expected, found `end of file`."},
	{`foo () simple_command`, "main.sh(1:8): syntax error: function body is expected, found `simple_command`."},
	{`$foo () {cmd;}`, "main.sh(1:6): syntax error: invalid function name was supplied."},
	{`cmd (`, "main.sh(1:6): syntax error: expected `)`, found `end of file`."},
	{`cmd )`, "main.sh(1:5): syntax error: token `)` cannot be placed here."},
	{`cmd arg (`, "main.sh(1:9): syntax error: token `(` cannot be placed here."},
	{`cmd arg )`, "main.sh(1:9): syntax error: token `)` cannot be placed here."},
	{`cmd arg(arg`, "main.sh(1:8): syntax error: token `(` cannot be placed here."},
	{`cmd arg)arg`, "main.sh(1:8): syntax error: token `)` cannot be placed here."},
	{`func() if true; then cmd;fi`, "main.sh(1:28): syntax error: function body is expected to be a group or subshell."},

	{`function`, "main.sh(1:9): syntax error: function name is required."},
	{`function foo ()`, "main.sh(1:16): syntax error: function body is expected, found `end of file`."},
	{`function foo () simple_command`, "main.sh(1:17): syntax error: function body is expected, found `simple_command`."},
	{`function $foo () {cmd;}`, "main.sh(1:14): syntax error: invalid function name was supplied."},
	{`function cmd (`, "main.sh(1:15): syntax error: expected `)`, found `end of file`."},
	{`function cmd )`, "main.sh(1:14): syntax error: function body is expected, found `)`."},
	{`function cmd () function foo() {cmd;}`, "main.sh(1:17): syntax error: function body is expected, found `function`."},
	{`function func() {cmd;} | cat`, "main.sh(1:24): syntax error: unexpected token `|`."},
	{`func() {cmd;} | cat`, "main.sh(1:15): syntax error: unexpected token `|`."},
	{`function func() if true; then cmd;fi`, "main.sh(1:37): syntax error: function body is expected to be a group or subshell."},

	{`function func( ; ){ cmd; }`, "main.sh(1:16): syntax error: unexpected token `;`."},
	{`function func( -; ){ cmd; }`, "main.sh(1:17): syntax error: expected a valid flag name, found `;`."},
	{`function func( --; ){ cmd; }`, "main.sh(1:18): syntax error: expected a valid flag name, found `;`."},
	{`function func( -abc ){ cmd; }`, "main.sh(1:17): syntax error: short flags can only be one character long, found `abc`."},
	{`function func( -a[ ){ cmd; }`, "main.sh(1:18): syntax error: expected [=] to indicate optional value, found `[blank)`."},
	{`function func( -a -a ){ cmd; }`, "main.sh(1:22): syntax error: flag declared twice: `a`."},
	{`function func( --foo --foo ){ cmd; }`, "main.sh(1:28): syntax error: flag declared twice: `foo`."},
	{`function func( -a --a ){ cmd; }`, "main.sh(1:23): syntax error: flag declared twice: `a`."},
}

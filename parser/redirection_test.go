package parser_test

import "github.com/yassinebenaid/nbs/ast"

var redirectionTests = []testCase{
	{`cmd>'file.ext' arg > file>/foo/bar arg2 >"$var" arg345>xyz 645 >file 3>foo.bar 45> /foo/bar 12.34>baz`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
					ast.Word("arg345"),
					ast.Word("645"),
					ast.Word("12.34"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file.ext")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.SimpleExpansion("var")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("xyz")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("3"), Method: ">", Dst: ast.Word("foo.bar")},
					{Src: ast.FileDescriptor("45"), Method: ">", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("baz")},
				},
			},
		},
	}},
	{`cmd>>'file.ext' arg >> file>>/foo/bar arg2 >>"$var" arg345>>xyz 123 >>file 3>>foo.bar 12.34>>baz`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
					ast.Word("arg345"),
					ast.Word("123"),
					ast.Word("12.34"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("file.ext")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.SimpleExpansion("var")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("xyz")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("3"), Method: ">>", Dst: ast.Word("foo.bar")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("baz")},
				},
			},
		},
	}},
	{`cmd&>'file.ext' arg &> file&>/foo/bar arg2 &>"$var" 3&>xyz`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
					ast.Word("3"),
				}, Redirections: []ast.Redirection{
					{Src: nil, Method: "&>", Dst: ast.Word("file.ext")},
					{Src: nil, Method: "&>", Dst: ast.Word("file")},
					{Src: nil, Method: "&>", Dst: ast.Word("/foo/bar")},
					{Src: nil, Method: "&>", Dst: ast.SimpleExpansion("var")},
					{Src: nil, Method: "&>", Dst: ast.Word("xyz")},
				},
			},
		},
	}},
	{`cmd>&1 arg >&2 arg>&3 arg345>&4 5>&6 985 >&19 12.34>& 7 8>& 9 >& $FD 3>&$FD`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg"),
					ast.Word("arg345"),
					ast.Word("985"),
					ast.Word("12.34"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.Word("1")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.Word("2")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.Word("3")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.Word("4")},
					{Src: ast.FileDescriptor("5"), Method: ">&", Dst: ast.Word("6")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.Word("19")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.Word("7")},
					{Src: ast.FileDescriptor("8"), Method: ">&", Dst: ast.Word("9")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.SimpleExpansion("FD")},
					{Src: ast.FileDescriptor("3"), Method: ">&", Dst: ast.SimpleExpansion("FD")},
				},
			},
		},
	}},
	{`cmd<'file.ext' arg < file</foo/bar arg123<foo 3<bar 928 <bar 282 <&123 <&3 4<&5 6<& 7 <& "$FD" <&'9'`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg123"),
					ast.Word("928"),
					ast.Word("282"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("file.ext")},
					{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("foo")},
					{Src: ast.FileDescriptor("3"), Method: "<", Dst: ast.Word("bar")},
					{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("bar")},
					{Src: ast.FileDescriptor("0"), Method: "<&", Dst: ast.Word("123")},
					{Src: ast.FileDescriptor("0"), Method: "<&", Dst: ast.Word("3")},
					{Src: ast.FileDescriptor("4"), Method: "<&", Dst: ast.Word("5")},
					{Src: ast.FileDescriptor("6"), Method: "<&", Dst: ast.Word("7")},
					{Src: ast.FileDescriptor("0"), Method: "<&", Dst: ast.SimpleExpansion("FD")},
					{Src: ast.FileDescriptor("0"), Method: "<&", Dst: ast.Word("9")},
				},
			},
		},
	}},
	{`cmd<<<'foo bar' arg <<< foo<<<foo-bar arg2 <<<"$var" 3<<<foobar <<<123 4<<<123 5<<< 	776`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("foo bar")},
					{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("foo")},
					{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("foo-bar")},
					{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.SimpleExpansion("var")},
					{Src: ast.FileDescriptor("3"), Method: "<<<", Dst: ast.Word("foobar")},
					{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("123")},
					{Src: ast.FileDescriptor("4"), Method: "<<<", Dst: ast.Word("123")},
					{Src: ast.FileDescriptor("5"), Method: "<<<", Dst: ast.Word("776")},
				},
			},
		},
	}},
}

var redirectionErrorHandlingCases = []errorHandlingTestCase{
	{`cmd >`, "syntax error: a redirection operand was not provided after the `>`."},
	{`cmd > >file.txt`, "syntax error: a redirection operand was not provided after the `>`."},
	{`cmd >>`, "syntax error: a redirection operand was not provided after the `>>`."},
	{`cmd >> >>foo`, "syntax error: a redirection operand was not provided after the `>>`."},
	{`cmd >& `, "syntax error: a redirection operand was not provided after the `>&`."},
	{`cmd >& >&$foo`, "syntax error: a redirection operand was not provided after the `>&`."},

	{`cmd 1>`, "syntax error: a redirection operand was not provided after the `>`."},
	{`cmd 1>1>x`, "syntax error: a redirection operand was not provided after the `>`."},
	{`cmd 1>>`, "syntax error: a redirection operand was not provided after the `>>`."},
	{`cmd 1>>1>>x`, "syntax error: a redirection operand was not provided after the `>>`."},
	{`cmd 1>& `, "syntax error: a redirection operand was not provided after the `>&`."},
	{`cmd 1>&1>&2`, "syntax error: a redirection operand was not provided after the `>&`."},

	{`cmd <`, "syntax error: a redirection operand was not provided after the `<`."},
	{`cmd < <foo`, "syntax error: a redirection operand was not provided after the `<`."},
	{`cmd 1<`, "syntax error: a redirection operand was not provided after the `<`."},
	{`cmd 1<1<`, "syntax error: a redirection operand was not provided after the `<`."},
	{`cmd 1<&`, "syntax error: a redirection operand was not provided after the `<&`."},
	{`cmd 1<&2<foo`, "syntax error: a redirection operand was not provided after the `<&`."},

	{`cmd &>`, "syntax error: a redirection operand was not provided after the `&>`."},
	{`cmd &>12>foo`, "syntax error: a redirection operand was not provided after the `&>`."},

	{`cmd <<<`, "syntax error: a redirection operand was not provided after the `<<<`."},
	{`cmd <<<<<<foo`, "syntax error: a redirection operand was not provided after the `<<<`."},
	{`cmd 2<<<`, "syntax error: a redirection operand was not provided after the `<<<`."},
	{`cmd <<<2<<<foo`, "syntax error: a redirection operand was not provided after the `<<<`."},
}

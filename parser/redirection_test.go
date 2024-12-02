package parser_test

import "github.com/yassinebenaid/ryuko/ast"

var redirectionTests = []testCase{
	{`cmd>'file.ext' arg > file>/foo/bar arg2 >"$var" arg345>xyz 645 >file 3>foo.bar 45> /foo/bar 12.34>baz`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg2"),
				ast.Word("arg345"),
				ast.Word("645"),
				ast.Word("12.34"),
			}, Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("file.ext")},
				{Src: "1", Method: ">", Dst: ast.Word("file")},
				{Src: "1", Method: ">", Dst: ast.Word("/foo/bar")},
				{Src: "1", Method: ">", Dst: ast.QuotedString{ast.Var("var")}},
				{Src: "1", Method: ">", Dst: ast.Word("xyz")},
				{Src: "1", Method: ">", Dst: ast.Word("file")},
				{Src: "3", Method: ">", Dst: ast.Word("foo.bar")},
				{Src: "45", Method: ">", Dst: ast.Word("/foo/bar")},
				{Src: "1", Method: ">", Dst: ast.Word("baz")},
			},
		},
	}},
	{`cmd>|'file.ext' arg >| file>|/foo/bar arg2 >|"$var" arg345>|xyz 645 >|file 3>|foo.bar 45>| /foo/bar 12.34>|baz`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg2"),
				ast.Word("arg345"),
				ast.Word("645"),
				ast.Word("12.34"),
			}, Redirections: []ast.Redirection{
				{Src: "1", Method: ">|", Dst: ast.Word("file.ext")},
				{Src: "1", Method: ">|", Dst: ast.Word("file")},
				{Src: "1", Method: ">|", Dst: ast.Word("/foo/bar")},
				{Src: "1", Method: ">|", Dst: ast.QuotedString{ast.Var("var")}},
				{Src: "1", Method: ">|", Dst: ast.Word("xyz")},
				{Src: "1", Method: ">|", Dst: ast.Word("file")},
				{Src: "3", Method: ">|", Dst: ast.Word("foo.bar")},
				{Src: "45", Method: ">|", Dst: ast.Word("/foo/bar")},
				{Src: "1", Method: ">|", Dst: ast.Word("baz")},
			},
		},
	}},
	{`cmd>>'file.ext' arg >> file>>/foo/bar arg2 >>"$var" arg345>>xyz 123 >>file 3>>foo.bar 12.34>>baz`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg2"),
				ast.Word("arg345"),
				ast.Word("123"),
				ast.Word("12.34"),
			}, Redirections: []ast.Redirection{
				{Src: "1", Method: ">>", Dst: ast.Word("file.ext")},
				{Src: "1", Method: ">>", Dst: ast.Word("file")},
				{Src: "1", Method: ">>", Dst: ast.Word("/foo/bar")},
				{Src: "1", Method: ">>", Dst: ast.QuotedString{ast.Var("var")}},
				{Src: "1", Method: ">>", Dst: ast.Word("xyz")},
				{Src: "1", Method: ">>", Dst: ast.Word("file")},
				{Src: "3", Method: ">>", Dst: ast.Word("foo.bar")},
				{Src: "1", Method: ">>", Dst: ast.Word("baz")},
			},
		},
	}},
	{`cmd&>'file.ext' arg &> file&>/foo/bar arg2 &>"$var" 3&>xyz`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg2"),
				ast.Word("3"),
			}, Redirections: []ast.Redirection{
				{Method: "&>", Dst: ast.Word("file.ext")},
				{Method: "&>", Dst: ast.Word("file")},
				{Method: "&>", Dst: ast.Word("/foo/bar")},
				{Method: "&>", Dst: ast.QuotedString{ast.Var("var")}},
				{Method: "&>", Dst: ast.Word("xyz")},
			},
		},
	}},
	{`cmd&>>'file.ext' arg &>> file&>>/foo/bar arg2 &>>"$var" 3&>>xyz`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg2"),
				ast.Word("3"),
			}, Redirections: []ast.Redirection{
				{Method: "&>>", Dst: ast.Word("file.ext")},
				{Method: "&>>", Dst: ast.Word("file")},
				{Method: "&>>", Dst: ast.Word("/foo/bar")},
				{Method: "&>>", Dst: ast.QuotedString{ast.Var("var")}},
				{Method: "&>>", Dst: ast.Word("xyz")},
			},
		},
	}},
	{`cmd>&1 arg >&2 arg>&3 arg345>&4 5>&6 985 >&19 12.34>& 7 8>& 9 >& $FD 3>&$FD`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg"),
				ast.Word("arg345"),
				ast.Word("985"),
				ast.Word("12.34"),
			}, Redirections: []ast.Redirection{
				{Src: "1", Method: ">&", Dst: ast.Word("1")},
				{Src: "1", Method: ">&", Dst: ast.Word("2")},
				{Src: "1", Method: ">&", Dst: ast.Word("3")},
				{Src: "1", Method: ">&", Dst: ast.Word("4")},
				{Src: "5", Method: ">&", Dst: ast.Word("6")},
				{Src: "1", Method: ">&", Dst: ast.Word("19")},
				{Src: "1", Method: ">&", Dst: ast.Word("7")},
				{Src: "8", Method: ">&", Dst: ast.Word("9")},
				{Src: "1", Method: ">&", Dst: ast.Var("FD")},
				{Src: "3", Method: ">&", Dst: ast.Var("FD")},
			},
		},
	}},
	{`cmd<'file.ext' arg < file</foo/bar arg123<foo 3<bar 928 <bar 282 <&123 <&3 4<&5 6<& 7 <& "$var" <&'9'`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg123"),
				ast.Word("928"),
				ast.Word("282"),
			}, Redirections: []ast.Redirection{
				{Src: "0", Method: "<", Dst: ast.Word("file.ext")},
				{Src: "0", Method: "<", Dst: ast.Word("file")},
				{Src: "0", Method: "<", Dst: ast.Word("/foo/bar")},
				{Src: "0", Method: "<", Dst: ast.Word("foo")},
				{Src: "3", Method: "<", Dst: ast.Word("bar")},
				{Src: "0", Method: "<", Dst: ast.Word("bar")},
				{Src: "0", Method: "<&", Dst: ast.Word("123")},
				{Src: "0", Method: "<&", Dst: ast.Word("3")},
				{Src: "4", Method: "<&", Dst: ast.Word("5")},
				{Src: "6", Method: "<&", Dst: ast.Word("7")},
				{Src: "0", Method: "<&", Dst: ast.QuotedString{ast.Var("var")}},
				{Src: "0", Method: "<&", Dst: ast.Word("9")},
			},
		},
	}},
	{`cmd<<<'foo bar' arg <<< foo<<<foo-bar arg2 <<<"$var" 3<<<foobar <<<123 4<<<123 5<<< 	776`, ast.Script{

		ast.Command{
			Name: ast.Word("cmd"),
			Args: []ast.Expression{
				ast.Word("arg"),
				ast.Word("arg2"),
			}, Redirections: []ast.Redirection{
				{Src: "0", Method: "<<<", Dst: ast.Word("foo bar")},
				{Src: "0", Method: "<<<", Dst: ast.Word("foo")},
				{Src: "0", Method: "<<<", Dst: ast.Word("foo-bar")},
				{Src: "0", Method: "<<<", Dst: ast.QuotedString{ast.Var("var")}},
				{Src: "3", Method: "<<<", Dst: ast.Word("foobar")},
				{Src: "0", Method: "<<<", Dst: ast.Word("123")},
				{Src: "4", Method: "<<<", Dst: ast.Word("123")},
				{Src: "5", Method: "<<<", Dst: ast.Word("776")},
			},
		},
	}},
	{`cmd <&- 2<&- >&-`, ast.Script{
		ast.Command{
			Name: ast.Word("cmd"),
			Redirections: []ast.Redirection{
				{Src: "0", Method: "<&", Close: true},
				{Src: "2", Method: "<&", Close: true},
				{Src: "1", Method: ">&", Close: true},
			},
		},
	}},
}

var redirectionErrorHandlingCases = []errorHandlingTestCase{
	{`cmd >`, "syntax error: a redirection operand was not provided after the `>`. (line: 1, column: 6)"},
	{`cmd > >file.txt`, "syntax error: a redirection operand was not provided after the `>`. (line: 1, column: 7)"},
	{`cmd >>`, "syntax error: a redirection operand was not provided after the `>>`. (line: 1, column: 7)"},
	{`cmd >> >>foo`, "syntax error: a redirection operand was not provided after the `>>`. (line: 1, column: 8)"},
	{`cmd >& `, "syntax error: a redirection operand was not provided after the `>&`. (line: 1, column: 8)"},
	{`cmd >& >&$foo`, "syntax error: a redirection operand was not provided after the `>&`. (line: 1, column: 8)"},

	{`cmd 1>`, "syntax error: a redirection operand was not provided after the `>`. (line: 1, column: 7)"},
	{`cmd 1>1>x`, "syntax error: a redirection operand was not provided after the `>`. (line: 1, column: 7)"},
	{`cmd 1>>`, "syntax error: a redirection operand was not provided after the `>>`. (line: 1, column: 8)"},
	{`cmd 1>>1>>x`, "syntax error: a redirection operand was not provided after the `>>`. (line: 1, column: 8)"},
	{`cmd 1>& `, "syntax error: a redirection operand was not provided after the `>&`. (line: 1, column: 9)"},
	{`cmd 1>&1>&2`, "syntax error: a redirection operand was not provided after the `>&`. (line: 1, column: 8)"},

	{`cmd >|`, "syntax error: a redirection operand was not provided after the `>|`. (line: 1, column: 7)"},
	{`cmd >|>|foo`, "syntax error: a redirection operand was not provided after the `>|`. (line: 1, column: 7)"},
	{`cmd 1>|`, "syntax error: a redirection operand was not provided after the `>|`. (line: 1, column: 8)"},
	{`cmd 1>|2>|foo`, "syntax error: a redirection operand was not provided after the `>|`. (line: 1, column: 8)"},

	{`cmd <`, "syntax error: a redirection operand was not provided after the `<`. (line: 1, column: 6)"},
	{`cmd < <foo`, "syntax error: a redirection operand was not provided after the `<`. (line: 1, column: 7)"},
	{`cmd 1<`, "syntax error: a redirection operand was not provided after the `<`. (line: 1, column: 7)"},
	{`cmd 1<1<`, "syntax error: a redirection operand was not provided after the `<`. (line: 1, column: 7)"},
	{`cmd 1<&`, "syntax error: a redirection operand was not provided after the `<&`. (line: 1, column: 8)"},
	{`cmd 1<&2<foo`, "syntax error: a redirection operand was not provided after the `<&`. (line: 1, column: 8)"},

	{`cmd &>`, "syntax error: a redirection operand was not provided after the `&>`. (line: 1, column: 7)"},
	{`cmd &>12>foo`, "syntax error: a redirection operand was not provided after the `&>`. (line: 1, column: 7)"},

	{`cmd <<<`, "syntax error: a redirection operand was not provided after the `<<<`. (line: 1, column: 8)"},
	{`cmd <<<<<<foo`, "syntax error: a redirection operand was not provided after the `<<<`. (line: 1, column: 8)"},
	{`cmd 2<<<`, "syntax error: a redirection operand was not provided after the `<<<`. (line: 1, column: 9)"},
	{`cmd <<<2<<<foo`, "syntax error: a redirection operand was not provided after the `<<<`. (line: 1, column: 8)"},
}

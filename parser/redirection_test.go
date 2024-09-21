package parser_test

import "github.com/yassinebenaid/nbs/ast"

var redirectionTests = []testCase{
	{`cmd>'file.ext' arg > file>/foo/bar arg2 >"$var" arg345>xyz 3>foo.bar 12.34>baz`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
					ast.Word("arg345"),
					ast.Word("12.34"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file.ext")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.SimpleExpansion("var")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("xyz")},
					{Src: ast.FileDescriptor("3"), Method: ">", Dst: ast.Word("foo.bar")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("baz")},
				},
			},
		},
	}},
	{`cmd>>'file.ext' arg >> file>>/foo/bar arg2 >>"$var" arg345>>xyz 3>>foo.bar 12.34>>baz`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
					ast.Word("arg345"),
					ast.Word("12.34"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("file.ext")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.SimpleExpansion("var")},
					{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("xyz")},
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
	{`cmd>&1 arg >&2 arg>&3 arg345>&4 5>&6 12.34>&7`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg"),
					ast.Word("arg345"),
					ast.Word("12.34"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.FileDescriptor("1")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.FileDescriptor("2")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.FileDescriptor("3")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.FileDescriptor("4")},
					{Src: ast.FileDescriptor("5"), Method: ">&", Dst: ast.FileDescriptor("6")},
					{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.FileDescriptor("7")},
				},
			},
		},
	}},
	{`cmd<'file.ext' arg`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("file.ext")},
				},
			},
		},
	}},
	// {`cmd<file.txt 2<file.txt 3<&5 arg<foo`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: []ast.Node{
	// 				ast.Word("arg"),
	// 			},
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("file.txt")},
	// 				{Src: ast.FileDescriptor("2"), Method: "<", Dst: ast.Word("file.txt")},
	// 				{Src: ast.FileDescriptor("3"), Method: "<&", Dst: ast.FileDescriptor("5")},
	// 				{Src: ast.FileDescriptor("0"), Method: "<", Dst: ast.Word("foo")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd&>file`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: nil,
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.StdoutStderr{}, Method: "&>", Dst: ast.Word("file")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd <<< foo arg<<<"foo bar"`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: []ast.Node{
	// 				ast.Word("arg"),
	// 			},
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("foo")},
	// 				{Src: ast.FileDescriptor("0"), Method: "<<<", Dst: ast.Word("foo bar")},
	// 			},
	// 		},
	// 	},
	// }},
}

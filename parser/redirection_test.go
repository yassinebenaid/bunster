package parser_test

import "github.com/yassinebenaid/nbs/ast"

var redirectionTests = []testCase{
	{`cmd>file.ext > file arg>/foo/bar arg2 >$var arg345>xyz`, ast.Script{
		Statements: []ast.Node{
			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Node{
					ast.Word("arg"),
					ast.Word("arg2"),
					ast.Word("arg345"),
				}, Redirections: []ast.Redirection{
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file.ext")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("file")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("/foo/bar")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.SimpleExpansion("var")},
					{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("xyz")},
				},
			},
		},
	}},
	// {`cmd > /dev/null arg1 >foobar.txt arg2 >$var arg3 >'some string goes here' arg>/dev/null`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: []ast.Node{
	// 				ast.Word("arg1"),
	// 				ast.Word("arg2"),
	// 				ast.Word("arg3"),
	// 				ast.Word("arg"),
	// 			},
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("/dev/null")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("foobar.txt")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.SimpleExpansion("var")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("some string goes here")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("/dev/null")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd>"foo $bar baz">'boo'`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: nil,
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Concatination{
	// 					Nodes: []ast.Node{
	// 						ast.Word("foo "),
	// 						ast.SimpleExpansion("bar"),
	// 						ast.Word(" baz"),
	// 					},
	// 				}},
	// 				{Src: ast.FileDescriptor("1"), Method: ">", Dst: ast.Word("boo")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd >> /dev/null arg1 >>foobar.txt arg2 >>$var arg3 >>'some string goes here' arg4`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: []ast.Node{
	// 				ast.Word("arg1"),
	// 				ast.Word("arg2"),
	// 				ast.Word("arg3"),
	// 				ast.Word("arg4"),
	// 			},
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("/dev/null")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("foobar.txt")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.SimpleExpansion("var")},
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("some string goes here")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd arg>>/dev/null`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: []ast.Node{
	// 				ast.Word("arg"),
	// 			},
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("/dev/null")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd>>"foo $bar baz">>'boo'`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: nil,
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Concatination{
	// 					Nodes: []ast.Node{
	// 						ast.Word("foo "),
	// 						ast.SimpleExpansion("bar"),
	// 						ast.Word(" baz"),
	// 					},
	// 				}},
	// 				{Src: ast.FileDescriptor("1"), Method: ">>", Dst: ast.Word("boo")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd 3>file.txt 123>>$foo 123>>xyz`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: nil,
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("3"), Method: ">", Dst: ast.Word("file.txt")},
	// 				{Src: ast.FileDescriptor("123"), Method: ">>", Dst: ast.SimpleExpansion("foo")},
	// 				{Src: ast.FileDescriptor("123"), Method: ">>", Dst: ast.Word("xyz")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`cmd>&2 3>&5`, ast.Script{
	// 	Statements: []ast.Node{
	// 		ast.Command{
	// 			Name: ast.Word("cmd"),
	// 			Args: nil,
	// 			Redirections: []ast.Redirection{
	// 				{Src: ast.FileDescriptor("1"), Method: ">&", Dst: ast.FileDescriptor("2")},
	// 				{Src: ast.FileDescriptor("3"), Method: ">&", Dst: ast.FileDescriptor("5")},
	// 			},
	// 		},
	// 	},
	// }},
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

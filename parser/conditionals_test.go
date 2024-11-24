package parser_test

import (
	"github.com/yassinebenaid/bunny/ast"
)

var conditionalsTests = []testCase{
	{`[[ foo-bar_baz ]]`, ast.Script{
		ast.Test{
			Expr: ast.Word("foo-bar_baz"),
		},
	}},
	{`[[ -a-file ]]`, ast.Script{
		ast.Test{
			Expr: ast.Word("-a-file"),
		},
	}},
	{`[[ "-a" ]]`, ast.Script{
		ast.Test{
			Expr: ast.Word("-a"),
		},
	}},
	{`
		[[  -a  file ]]
		[[  -b  file ]]
		[[  -c  file ]]
		[[  -d  file ]]
		[[  -e  file ]]
		[[  -f  file ]]
		[[  -g  file ]]
		[[  -h  file ]]
		[[  -k  file ]]
		[[  -p  file ]]
		[[  -r  file ]]
		[[  -s  file ]]
		[[  -t  file ]]
		[[  -u  file ]]
		[[  -w  file ]]
		[[  -x  file ]]
		[[  -G  file ]]
		[[  -L  file ]]
		[[  -N  file ]]
		[[  -O  file ]]
		[[  -S  file ]]
		[[  -z  file ]]
		[[  -n  file ]]
		[[  -v  file ]]
	`, ast.Script{
		ast.Test{Expr: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-d", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-e", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-f", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-g", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-h", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-k", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-p", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-r", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-s", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-t", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-u", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-w", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-x", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-G", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-L", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-N", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-O", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-S", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-z", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-n", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-v", Operand: ast.Word("file")}},
	}},
	{`
		[[ file1 -ef file2 ]]
		[[ file1 -nt file2 ]]
		[[ file1 -ot file2 ]]
		[[ file1 = file2 ]]
		[[ file1 == file2 ]]
		[[ file1 != file2 ]]
		[[ file1 < file2 ]]
		[[ file1 > file2 ]]
		[[ file1 =~ ^(.*)/([[:alnum:]]+)-(.*)$ ]]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-nt", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ot", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "==", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "!=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "<", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: ">", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "=~", Right: ast.Word("^(.*)/([[:alnum:]]+)-(.*)$")}},
	}},
	{`
		[[ file1 && file2 ]]
		[[ file1 || file2 ]]
		[[ file1 && file2 || file3 ]]
		[[ file1 || file2 && file3 ]]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "||", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.Word("file1"),
				Operator: "&&",
				Right:    ast.Word("file2")},
			Operator: "||",
			Right:    ast.Word("file3"),
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.Word("file1"),
				Operator: "||",
				Right:    ast.Word("file2")},
			Operator: "&&",
			Right:    ast.Word("file3"),
		}},
	}},
	{`
		[[ -a file1 && -b file2 ]]
		[[ -a file1 || -b file2 ]]
		[[ -a file1 && -b file2 || -c file3 ]]
		[[ -a file1 || -b file2 && -c file3 ]]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "&&",
				Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file3")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "||",
				Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file3")},
		}},
	}},
	{`
		[[ file1 -ef file2 && file1 -ef file2 ]]
		[[ file1 -ef file2 || file1 -ef file2 ]]
		[[ file1 -ef file2 && file1 -ef file2 || file1 -ef file2 ]]
		[[ file1 -ef file2 || file1 -ef file2 && file1 -ef file2 ]]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "&&",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "||",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "&&",
				Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "||",
				Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
	}},
	{`
		[[ !file1 ]]
		[[ !file1 && !file2 ]]
		[[ ! -a file1 && ! -b file2 ]]
		[[ ! file1 -ef file2 && !file1 -ef file2 ]]

		[[ !file1 || !file2 ]]
		[[ ! -a file1 || ! -b file2 ]]
		[[ ! file1 -ef file2 || !file1 -ef file2 ]]
	`, ast.Script{
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
	}},
	{`
		[[ (file1) ]]
		[[ (!file1) ]]
		[[ !(file1) ]]
		[[ !(file1 && file2) ]]
		[[ ! (   file1  ) ]]
	`, ast.Script{
		ast.Test{Expr: ast.Word("file1")},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}}},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
	}},
	{`
		[[
			-a file1
		]]
		[[
			-a file1 &&
			-b file2
		]]
	`, ast.Script{
		ast.Test{Expr: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
	}},

	// POSIX compatible variant.
	{`[ foo-bar_baz ]`, ast.Script{
		ast.Test{
			Expr: ast.Word("foo-bar_baz"),
		},
	}},
	{`[ -a-file ]`, ast.Script{
		ast.Test{
			Expr: ast.Word("-a-file"),
		},
	}},
	{`[ "-a" ]`, ast.Script{
		ast.Test{
			Expr: ast.Word("-a"),
		},
	}},
	{`
		[  -a  file ]
		[  -b  file ]
		[  -c  file ]
		[  -d  file ]
		[  -e  file ]
		[  -f  file ]
		[  -g  file ]
		[  -h  file ]
		[  -k  file ]
		[  -p  file ]
		[  -r  file ]
		[  -s  file ]
		[  -t  file ]
		[  -u  file ]
		[  -w  file ]
		[  -x  file ]
		[  -G  file ]
		[  -L  file ]
		[  -N  file ]
		[  -O  file ]
		[  -S  file ]
		[  -z  file ]
		[  -n  file ]
		[  -v  file ]
	`, ast.Script{
		ast.Test{Expr: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-d", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-e", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-f", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-g", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-h", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-k", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-p", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-r", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-s", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-t", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-u", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-w", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-x", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-G", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-L", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-N", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-O", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-S", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-z", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-n", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-v", Operand: ast.Word("file")}},
	}},
	{`
		[ file1 -ef file2 ]
		[ file1 -nt file2 ]
		[ file1 -ot file2 ]
		[ file1 = file2 ]
		[ file1 == file2 ]
		[ file1 != file2 ]
		[ file1 < file2 ]
		[ file1 > file2 ]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-nt", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ot", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "==", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "!=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "<", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: ">", Right: ast.Word("file2")}},
	}},
	{`
		[ file1 -a file2 ]
		[ file1 -o file2 ]
		[ file1 -a file2 -o file3 ]
		[ file1 -o file2 -a file3 ]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "||", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.Word("file1"),
				Operator: "&&",
				Right:    ast.Word("file2")},
			Operator: "||",
			Right:    ast.Word("file3"),
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.Word("file1"),
				Operator: "||",
				Right:    ast.Word("file2")},
			Operator: "&&",
			Right:    ast.Word("file3"),
		}},
	}},
	{`
		[ -a file1 -a -b file2 ]
		[ -a file1 -o -b file2 ]
		[ -a file1 -a -b file2 -o -c file3 ]
		[ -a file1 -o -b file2 -a -c file3 ]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "&&",
				Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file3")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "||",
				Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file3")},
		}},
	}},
	{`
		[ file1 -ef file2 -a file1 -ef file2 ]
		[ file1 -ef file2 -o file1 -ef file2 ]
		[ file1 -ef file2 -a file1 -ef file2 -o file1 -ef file2 ]
		[ file1 -ef file2 -o file1 -ef file2 -a file1 -ef file2 ]
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "&&",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "||",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "&&",
				Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "||",
				Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
	}},
	{`
		[ !file1 ]
		[ !file1 -a !file2 ]
		[ ! -a file1 -a ! -b file2 ]
		[ ! file1 -ef file2 -a !file1 -ef file2 ]

		[ !file1 -o !file2 ]
		[ ! -a file1 -o ! -b file2 ]
		[ ! file1 -ef file2 -o !file1 -ef file2 ]
	`, ast.Script{
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
	}},
	{`
		[ (file1) ]
		[ (!file1) ]
		[ !(file1) ]
		[ !(file1 -a file2) ]
	`, ast.Script{
		ast.Test{Expr: ast.Word("file1")},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}}},
	}},
	{`test foo-bar_baz `, ast.Script{
		ast.Test{
			Expr: ast.Word("foo-bar_baz"),
		},
	}},
	{`test -a-file `, ast.Script{
		ast.Test{
			Expr: ast.Word("-a-file"),
		},
	}},
	{`test "-a" `, ast.Script{
		ast.Test{
			Expr: ast.Word("-a"),
		},
	}},
	{`
		test  -a  file
		test  -b  file
		test  -c  file
		test  -d  file
		test  -e  file
		test  -f  file
		test  -g  file
		test  -h  file
		test  -k  file
		test  -p  file
		test  -r  file
		test  -s  file
		test  -t  file
		test  -u  file
		test  -w  file
		test  -x  file
		test  -G  file
		test  -L  file
		test  -N  file
		test  -O  file
		test  -S  file
		test  -z  file
		test  -n  file
		test  -v  file
	`, ast.Script{
		ast.Test{Expr: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-d", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-e", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-f", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-g", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-h", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-k", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-p", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-r", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-s", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-t", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-u", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-w", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-x", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-G", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-L", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-N", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-O", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-S", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-z", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-n", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.UnaryConditional{Operator: "-v", Operand: ast.Word("file")}},
	}},
	{`
		test file1 -ef file2
		test file1 -nt file2
		test file1 -ot file2
		test file1 = file2
		test file1 == file2
		test file1 != file2
		test file1 < file2
		test file1 > file2
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-nt", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ot", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "==", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "!=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "<", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: ">", Right: ast.Word("file2")}},
	}},
	{`
		test file1 -a file2
		test file1 -o file2
		test file1 -a file2 -o file3
		test file1 -o file2 -a file3
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "||", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.Word("file1"),
				Operator: "&&",
				Right:    ast.Word("file2")},
			Operator: "||",
			Right:    ast.Word("file3"),
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.Word("file1"),
				Operator: "||",
				Right:    ast.Word("file2")},
			Operator: "&&",
			Right:    ast.Word("file3"),
		}},
	}},
	{`
		test -a file1 -a -b file2
		test -a file1 -o -b file2
		test -a file1 -a -b file2 -o -c file3
		test -a file1 -o -b file2 -a -c file3
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "&&",
				Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file3")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "||",
				Right:    ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.UnaryConditional{Operator: "-c", Operand: ast.Word("file3")},
		}},
	}},
	{`
		test file1 -ef file2 -a file1 -ef file2
		test file1 -ef file2 -o file1 -ef file2
		test file1 -ef file2 -a file1 -ef file2 -o file1 -ef file2
		test file1 -ef file2 -o file1 -ef file2 -a file1 -ef file2
	`, ast.Script{
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "&&",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "||",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "&&",
				Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left: ast.BinaryConditional{
				Left:     ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "||",
				Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
	}},
	{`
		test !file1
		test !file1 -a !file2
		test ! -a file1 -a ! -b file2
		test ! file1 -ef file2 -a !file1 -ef file2

		test !file1 -o !file2
		test ! -a file1 -o ! -b file2
		test ! file1 -ef file2 -o !file1 -ef file2
	`, ast.Script{
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.UnaryConditional{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.UnaryConditional{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.BinaryConditional{
			Left:     ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
	}},
	{`
		test (file1)
		test (!file1)
		test !(file1)
		test !(file1 -a file2)
	`, ast.Script{
		ast.Test{Expr: ast.Word("file1")},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Negation{Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Negation{Operand: ast.BinaryConditional{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}}},
	}},
}

var conditionalsErrorHandlingCases = []errorHandlingTestCase{
	{`[[`, "syntax error: bad conditional expression, unexpected token `end of file`. (line: 1, column: 3)"},
	{`[[ &`, "syntax error: bad conditional expression, unexpected token `&`. (line: 1, column: 4)"},
	{`[[]]`, "syntax error: expected a conditional expression before `]]`. (line: 1, column: 3)"},
	{`[[  ]]`, "syntax error: expected a conditional expression before `]]`. (line: 1, column: 5)"},
	{`[[  file `, "syntax error: expected `]]` to close conditional expression, found `end of file`. (line: 1, column: 10)"},
	{`[[  -a ]] `, "syntax error: bad conditional expression, expected an operand after -a, found `]]`. (line: 1, column: 8)"},
	{`[[ file file ]] `, "syntax error: expected `]]` to close conditional expression, found `file`. (line: 1, column: 9)"},
	{`[[ file = ]] `, "syntax error: bad conditional expression, expected an operand after `=`, found `]]`. (line: 1, column: 11)"},
	{`[[ file = & ]] `, "syntax error: bad conditional expression, expected an operand after `=`, found `&`. (line: 1, column: 11)"},
	{`[[ file && ]] `, "syntax error: bad conditional expression, unexpected token `]]`. (line: 1, column: 12)"},
	{`[[ ! ]] `, "syntax error: bad conditional expression, unexpected token `]]`. (line: 1, column: 6)"},
	{`[[ ( ]] `, "syntax error: bad conditional expression, unexpected token `]]`. (line: 1, column: 6)"},
	{`[[ (exp ]] `, "syntax error: expected a closing `)`, found `]]`. (line: 1, column: 9)"},

	{`[`, "syntax error: bad conditional expression, unexpected token `end of file`. (line: 1, column: 2)"},
	{`[ &`, "syntax error: bad conditional expression, unexpected token `&`. (line: 1, column: 3)"},
	{`[]`, "syntax error: expected a conditional expression before `]`. (line: 1, column: 2)"},
	{`[  ]`, "syntax error: expected a conditional expression before `]`. (line: 1, column: 4)"},
	{`[  file `, "syntax error: expected `]` to close conditional expression, found `end of file`. (line: 1, column: 9)"},
	{`[  -a ] `, "syntax error: bad conditional expression, expected an operand after -a, found `]`. (line: 1, column: 7)"},
	{`[ file file ] `, "syntax error: expected `]` to close conditional expression, found `file`. (line: 1, column: 8)"},
	{`[ file = ] `, "syntax error: bad conditional expression, expected an operand after `=`, found `]`. (line: 1, column: 10)"},
	{`[ file = & ] `, "syntax error: bad conditional expression, expected an operand after `=`, found `&`. (line: 1, column: 10)"},
	{`[ file -a ] `, "syntax error: bad conditional expression, unexpected token `]`. (line: 1, column: 11)"},
	{`[ ! ] `, "syntax error: bad conditional expression, unexpected token `]`. (line: 1, column: 5)"},
	{`[ ( ] `, "syntax error: bad conditional expression, unexpected token `]`. (line: 1, column: 5)"},
	{`[ (exp ] `, "syntax error: expected a closing `)`, found `]`. (line: 1, column: 8)"},

	{`test`, "syntax error: bad conditional expression, unexpected token `end of file`. (line: 1, column: 5)"},
	{`test &`, "syntax error: bad conditional expression, unexpected token `&`. (line: 1, column: 6)"},
	{`test  -a  `, "syntax error: bad conditional expression, expected an operand after -a, found `end of file`. (line: 1, column: 11)"},
	{`test file file`, "syntax error: bad conditional expected, unexpected token `file`. (line: 1, column: 11)"},
	{`test file =  `, "syntax error: bad conditional expression, expected an operand after `=`, found `end of file`. (line: 1, column: 14)"},
	{`test file = &  `, "syntax error: bad conditional expression, expected an operand after `=`, found `&`. (line: 1, column: 13)"},
	{`test file -a  `, "syntax error: bad conditional expression, unexpected token `end of file`. (line: 1, column: 15)"},
	{`test !  `, "syntax error: bad conditional expression, unexpected token `end of file`. (line: 1, column: 9)"},
	// {`test ( `, "syntax error: bad conditional expression, unexpected token `]`. (line: 1, column: 5)"},
	// {`test (exp `, "syntax error: expected a closing `)`, found `]`. (line: 1, column: 8)"},
}

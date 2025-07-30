package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
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
		ast.Test{Expr: ast.Unary{Operator: "-a", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-b", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-c", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-d", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-e", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-f", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-g", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-h", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-k", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-p", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-r", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-s", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-t", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-u", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-w", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-x", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-G", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-L", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-N", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-O", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-S", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-z", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-n", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-v", Operand: ast.Word("file")}},
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

		[[ 1 -eq 2 ]]
		[[ 1 -ne 2 ]]
		[[ 1 -lt 2 ]]
		[[ 1 -le 2 ]]
		[[ 1 -gt 2 ]]
		[[ 1 -ge 2 ]]
	`, ast.Script{
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-nt", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-ot", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "==", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "!=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "<", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: ">", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "=~", Right: ast.Word("^(.*)/([[:alnum:]]+)-(.*)$")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-eq", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-ne", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-lt", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-le", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-gt", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-ge", Right: ast.Word("2")}},
	}},
	{`
		[[ file1 && file2 ]]
		[[ file1 || file2 ]]
		[[ file1 && file2 || file3 ]]
		[[ file1 || file2 && file3 ]]
	`, ast.Script{
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "||", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Word("file1"),
				Operator: "&&",
				Right:    ast.Word("file2")},
			Operator: "||",
			Right:    ast.Word("file3"),
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
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
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "&&",
				Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.Unary{Operator: "-c", Operand: ast.Word("file3")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "||",
				Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-c", Operand: ast.Word("file3")},
		}},
	}},
	{`
		[[ file1 -ef file2 && file1 -ef file2 ]]
		[[ file1 -ef file2 || file1 -ef file2 ]]
		[[ file1 -ef file2 && file1 -ef file2 || file1 -ef file2 ]]
		[[ file1 -ef file2 || file1 -ef file2 && file1 -ef file2 ]]
	`, ast.Script{
		ast.Test{Expr: ast.Binary{
			Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "&&",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "||",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "&&",
				Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "||",
				Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
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
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Unary{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Unary{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
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
		ast.Test{Expr: ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}}},
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
		ast.Test{Expr: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
	}},
	{`[[ str ]]|[[ str ]] |& [[ str ]]
	`, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}},
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}, Stderr: true},
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}},
		},
	}},
	{`[[ str ]] >output.txt <input.txt 2>error.txt >&3 \
		 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6
	`, ast.Script{
		ast.Test{
			Expr: ast.Word("str"),
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">", Dst: ast.Word("error.txt")},
				{Src: "1", Method: ">&", Dst: ast.Word("3")},
				{Src: "1", Method: ">>", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<<<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">>", Dst: ast.Word("error.txt")},
				{Method: "&>", Dst: ast.Word("all.txt")},
				{Method: "&>>", Dst: ast.Word("all.txt")},
				{Src: "0", Method: "<&", Dst: ast.Word("4")},
				{Src: "5", Method: "<&", Dst: ast.Word("6")},
			},
		},
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
		ast.Test{Expr: ast.Unary{Operator: "-a", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-b", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-c", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-d", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-e", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-f", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-g", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-h", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-k", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-p", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-r", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-s", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-t", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-u", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-w", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-x", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-G", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-L", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-N", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-O", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-S", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-z", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-n", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-v", Operand: ast.Word("file")}},
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

		[ 1 -eq 2 ]
		[ 1 -ne 2 ]
		[ 1 -lt 2 ]
		[ 1 -le 2 ]
		[ 1 -gt 2 ]
		[ 1 -ge 2 ]
	`, ast.Script{
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-nt", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-ot", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "==", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "!=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "<", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: ">", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-eq", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-ne", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-lt", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-le", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-gt", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-ge", Right: ast.Word("2")}},
	}},
	{`
		[ file1 -a file2 ]
		[ file1 -o file2 ]
		[ file1 -a file2 -o file3 ]
		[ file1 -o file2 -a file3 ]
	`, ast.Script{
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "||", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Word("file1"),
				Operator: "&&",
				Right:    ast.Word("file2")},
			Operator: "||",
			Right:    ast.Word("file3"),
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
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
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "&&",
				Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.Unary{Operator: "-c", Operand: ast.Word("file3")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "||",
				Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-c", Operand: ast.Word("file3")},
		}},
	}},
	{`
		[ file1 -ef file2 -a file1 -ef file2 ]
		[ file1 -ef file2 -o file1 -ef file2 ]
		[ file1 -ef file2 -a file1 -ef file2 -o file1 -ef file2 ]
		[ file1 -ef file2 -o file1 -ef file2 -a file1 -ef file2 ]
	`, ast.Script{
		ast.Test{Expr: ast.Binary{
			Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "&&",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "||",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "&&",
				Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "||",
				Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
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
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Unary{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Unary{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
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
		ast.Test{Expr: ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}}},
	}},
	{`[ str ]|[ str ] |& [ str ]
	`, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}},
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}, Stderr: true},
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}},
		},
	}},
	{`[ str ] >output.txt <input.txt 2>error.txt >&3 \
		 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6
	`, ast.Script{
		ast.Test{
			Expr: ast.Word("str"),
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">", Dst: ast.Word("error.txt")},
				{Src: "1", Method: ">&", Dst: ast.Word("3")},
				{Src: "1", Method: ">>", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<<<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">>", Dst: ast.Word("error.txt")},
				{Method: "&>", Dst: ast.Word("all.txt")},
				{Method: "&>>", Dst: ast.Word("all.txt")},
				{Src: "0", Method: "<&", Dst: ast.Word("4")},
				{Src: "5", Method: "<&", Dst: ast.Word("6")},
			},
		},
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
		ast.Test{Expr: ast.Unary{Operator: "-a", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-b", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-c", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-d", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-e", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-f", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-g", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-h", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-k", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-p", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-r", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-s", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-t", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-u", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-w", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-x", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-G", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-L", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-N", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-O", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-S", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-z", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-n", Operand: ast.Word("file")}},
		ast.Test{Expr: ast.Unary{Operator: "-v", Operand: ast.Word("file")}},
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

		test  1  -eq  2
		test  1  -ne  2
		test  1  -lt  2
		test  1  -le  2
		test  1  -gt  2
		test  1  -ge  2
	`, ast.Script{
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-nt", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "-ot", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "==", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "!=", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "<", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: ">", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-eq", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-ne", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-lt", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-le", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-gt", Right: ast.Word("2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("1"), Operator: "-ge", Right: ast.Word("2")}},
	}},
	{`
		test file1 -a file2
		test file1 -o file2
		test file1 -a file2 -o file3
		test file1 -o file2 -a file3
	`, ast.Script{
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{Left: ast.Word("file1"), Operator: "||", Right: ast.Word("file2")}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Word("file1"),
				Operator: "&&",
				Right:    ast.Word("file2")},
			Operator: "||",
			Right:    ast.Word("file3"),
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
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
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "&&",
				Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.Unary{Operator: "-c", Operand: ast.Word("file3")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Unary{Operator: "-a", Operand: ast.Word("file1")},
				Operator: "||",
				Right:    ast.Unary{Operator: "-b", Operand: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.Unary{Operator: "-c", Operand: ast.Word("file3")},
		}},
	}},
	{`
		test file1 -ef file2 -a file1 -ef file2
		test file1 -ef file2 -o file1 -ef file2
		test file1 -ef file2 -a file1 -ef file2 -o file1 -ef file2
		test file1 -ef file2 -o file1 -ef file2 -a file1 -ef file2
	`, ast.Script{
		ast.Test{Expr: ast.Binary{
			Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "&&",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			Operator: "||",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "&&",
				Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "||",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left: ast.Binary{
				Left:     ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
				Operator: "||",
				Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
			},
			Operator: "&&",
			Right:    ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")},
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
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Unary{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "&&",
			Right:    ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Word("file1")},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Word("file2")},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Unary{Operator: "-a", Operand: ast.Word("file1")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Unary{Operator: "-b", Operand: ast.Word("file2")}},
		}},
		ast.Test{Expr: ast.Binary{
			Left:     ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
			Operator: "||",
			Right:    ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "-ef", Right: ast.Word("file2")}},
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
		ast.Test{Expr: ast.Negation{Operand: ast.Binary{Left: ast.Word("file1"), Operator: "&&", Right: ast.Word("file2")}}},
	}},
	{`test str|test str |& test str
	`, ast.Script{
		ast.Pipeline{
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}},
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}, Stderr: true},
			ast.PipelineCommand{Command: ast.Test{Expr: ast.Word("str")}},
		},
	}},
	{`test str>output.txt <input.txt 2>error.txt >&3 \
		 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6
	`, ast.Script{
		ast.Test{
			Expr: ast.Word("str"),
			Redirections: []ast.Redirection{
				{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">", Dst: ast.Word("error.txt")},
				{Src: "1", Method: ">&", Dst: ast.Word("3")},
				{Src: "1", Method: ">>", Dst: ast.Word("output.txt")},
				{Src: "0", Method: "<<<", Dst: ast.Word("input.txt")},
				{Src: "2", Method: ">>", Dst: ast.Word("error.txt")},
				{Method: "&>", Dst: ast.Word("all.txt")},
				{Method: "&>>", Dst: ast.Word("all.txt")},
				{Src: "0", Method: "<&", Dst: ast.Word("4")},
				{Src: "5", Method: "<&", Dst: ast.Word("6")},
			},
		},
	}},

	{`[[ str =~ $var ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.Var("var"),
	}}}},
	{`[[ str =~ $(cmd) ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.CommandSubstitution{
			ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 13}, Name: ast.Word("cmd")},
		},
	}}}},
	{`[[ str =~ <(cmd) ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.ProcessSubstitution{
			Direction: '<', Body: []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 13}, Name: ast.Word("cmd")}},
		},
	}}}},
	{`[[ str =~ >(cmd) ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.ProcessSubstitution{
			Direction: '>', Body: []ast.Statement{ast.Command{Position: ast.Position{File: "main.sh", Line: 1, Col: 13}, Name: ast.Word("cmd")}},
		},
	}}}},
	{`[[ str =~ "str" ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.Word("str"),
	}}}},
	{`[[ str =~ 'str' ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.Word("str"),
	}}}},
	{`[[ str =~ ${var} ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.Var("var"),
	}}}},
	{`[[ str =~ $((var)) ]]`, ast.Script{ast.Test{Expr: ast.Binary{
		Left: ast.Word("str"), Operator: "=~", Right: ast.Arithmetic{ast.Var("var")},
	}}}},

	{`[[ -f file ]] # comment`, ast.Script{ast.Test{Expr: ast.Unary{
		Operand: ast.Word("file"), Operator: "-f"},
	}}},
	{`[ -f file ] # comment`, ast.Script{ast.Test{Expr: ast.Unary{
		Operand: ast.Word("file"), Operator: "-f"},
	}}},
	{`test -f file # comment`, ast.Script{ast.Test{Expr: ast.Unary{
		Operand: ast.Word("file"), Operator: "-f"},
	}}},
}

var conditionalsErrorHandlingCases = []errorHandlingTestCase{
	{`[[`, "main.sh(1:3): syntax error: bad conditional expression, unexpected token `end of file`."},
	{`[[ &`, "main.sh(1:4): syntax error: bad conditional expression, unexpected token `&`."},
	{`[[]]`, "main.sh(1:3): syntax error: expected a conditional expression before `]]`."},
	{`[[  ]]`, "main.sh(1:5): syntax error: expected a conditional expression before `]]`."},
	{`[[  file `, "main.sh(1:10): syntax error: expected `]]` to close conditional expression, found `end of file`."},
	{`[[  -a ]] `, "main.sh(1:8): syntax error: bad conditional expression, expected an operand after -a, found `]]`."},
	{`[[ file file ]] `, "main.sh(1:9): syntax error: expected `]]` to close conditional expression, found `file`."},
	{`[[ file = ]] `, "main.sh(1:11): syntax error: bad conditional expression, expected an operand after `=`, found `]]`."},
	{`[[ file = & ]] `, "main.sh(1:11): syntax error: bad conditional expression, expected an operand after `=`, found `&`."},
	{`[[ file && ]] `, "main.sh(1:12): syntax error: bad conditional expression, unexpected token `]]`."},
	{`[[ ! ]] `, "main.sh(1:6): syntax error: bad conditional expression, unexpected token `]]`."},
	{`[[ ( ]] `, "main.sh(1:6): syntax error: bad conditional expression, unexpected token `]]`."},
	{`[[ (exp ]] `, "main.sh(1:9): syntax error: expected a closing `)`, found `]]`."},
	{`[[ file1 -ef-file2 ]] `, "main.sh(1:10): syntax error: expected `]]` to close conditional expression, found `-`."},
	{`[[ file1 =file2 ]] `, "main.sh(1:10): syntax error: expected `]]` to close conditional expression, found `=`."},
	{`[[ file1 ]] arg`, "main.sh(1:13): syntax error: unexpected token `arg`."},

	{`[`, "main.sh(1:2): syntax error: bad conditional expression, unexpected token `end of file`."},
	{`[ &`, "main.sh(1:3): syntax error: bad conditional expression, unexpected token `&`."},
	{`[]`, "main.sh(1:2): syntax error: expected a conditional expression before `]`."},
	{`[  ]`, "main.sh(1:4): syntax error: expected a conditional expression before `]`."},
	{`[  file `, "main.sh(1:9): syntax error: expected `]` to close conditional expression, found `end of file`."},
	{`[  -a ] `, "main.sh(1:7): syntax error: bad conditional expression, expected an operand after -a, found `]`."},
	{`[ file file ] `, "main.sh(1:8): syntax error: expected `]` to close conditional expression, found `file`."},
	{`[ file = ] `, "main.sh(1:10): syntax error: bad conditional expression, expected an operand after `=`, found `]`."},
	{`[ file = & ] `, "main.sh(1:10): syntax error: bad conditional expression, expected an operand after `=`, found `&`."},
	{`[ file -a ] `, "main.sh(1:11): syntax error: bad conditional expression, unexpected token `]`."},
	{`[ ! ] `, "main.sh(1:5): syntax error: bad conditional expression, unexpected token `]`."},
	{`[ ( ] `, "main.sh(1:5): syntax error: bad conditional expression, unexpected token `]`."},
	{`[ (exp ] `, "main.sh(1:8): syntax error: expected a closing `)`, found `]`."},
	{`[ exp -o-exp ] `, "main.sh(1:7): syntax error: expected `]` to close conditional expression, found `-`."},
	{`[ file1 ] arg`, "main.sh(1:11): syntax error: bad conditional expression, unexpected token `arg`."},

	{`test`, "main.sh(1:5): syntax error: bad conditional expression, unexpected token `end of file`."},
	{`test &`, "main.sh(1:6): syntax error: bad conditional expression, unexpected token `&`."},
	{`test  -a  `, "main.sh(1:11): syntax error: bad conditional expression, expected an operand after -a, found `end of file`."},
	{`test file file`, "main.sh(1:11): syntax error: bad conditional expression, unexpected token `file`."},
	{`test file =  `, "main.sh(1:14): syntax error: bad conditional expression, expected an operand after `=`, found `end of file`."},
	{`test file = &  `, "main.sh(1:13): syntax error: bad conditional expression, expected an operand after `=`, found `&`."},
	{`test file -a  `, "main.sh(1:15): syntax error: bad conditional expression, unexpected token `end of file`."},
	{`test !  `, "main.sh(1:9): syntax error: bad conditional expression, unexpected token `end of file`."},
	{`test ( `, "main.sh(1:8): syntax error: bad conditional expression, unexpected token `end of file`."},
	{`test (exp `, "main.sh(1:11): syntax error: expected a closing `)`, found `end of file`."},
}

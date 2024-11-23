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
}

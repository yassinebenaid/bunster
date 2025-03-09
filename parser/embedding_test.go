package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
)

var embeddingTests = []testCase{
	{`@embed file`, ast.Script{ast.Embed{"file"}}},
	{`@embed "file"`, ast.Script{ast.Embed{"file"}}},
	{`@embed 'file'`, ast.Script{ast.Embed{"file"}}},
	{`@embed file 'file' "file"`, ast.Script{ast.Embed{
		"file",
		"file",
		"file",
	}}},
	{`
		@embed file
		@embed file

		command
		@ embed
		\@embed
	`, ast.Script{
		ast.Embed{"file"},
		ast.Embed{"file"},
		ast.Command{Name: ast.Word("command")},
		ast.Command{Name: ast.Word("@"), Args: []ast.Expression{ast.Word("embed")}},
		ast.Command{Name: ast.Word("@embed")},
	}},
	{`
		@embed file; @embed file;

	`, ast.Script{
		ast.Embed{"file"},
		ast.Embed{"file"},
	}},
	{`
	function foo(){
		
		@embed file; @embed file;
		
	}`, ast.Script{
		ast.Function{Name: "foo", Command: ast.Group{
			Body: []ast.Statement{
				ast.Embed{"file"},
				ast.Embed{"file"},
			},
		}},
	}},
}

var embeddingErrorHandlingCases = []errorHandlingTestCase{
	{`@embed`, "syntax error: expected a blank after the @embed directive, found end of file. (line: 1, column: 7)"},
	{`@embed `, "syntax error: unexpected token: end of file. (line: 1, column: 8)"},
	{`@embed $var`, "syntax error: expected a valid file path. (line: 1, column: 12)"},
	{`@embed "$var foo"`, "syntax error: expected a valid file path. (line: 1, column: 18)"},
	{`@embed file | cmd`, "syntax error: expected a valid file path. (line: 1, column: 13)"},
	{`@embed /foo/bar`, `syntax error: path cannot start or end with slash, "/foo/bar". (line: 1, column: 16)`},
	{`@embed foo/bar/`, `syntax error: path cannot start or end with slash, "foo/bar/". (line: 1, column: 16)`},
	{`@embed .`, `syntax error: expected a valid file path, found ".". (line: 1, column: 9)`},
	{`@embed \"`, `syntax error: expected a valid file path, found "\"". (line: 1, column: 10)`},
	{`@embed *`, `syntax error: expected a valid file path, found "*". (line: 1, column: 9)`},
	{`@embed \<`, `syntax error: expected a valid file path, found "<". (line: 1, column: 10)`},
	{`@embed \>`, `syntax error: expected a valid file path, found ">". (line: 1, column: 10)`},
	{`@embed \?`, `syntax error: expected a valid file path, found "?". (line: 1, column: 10)`},
	{"@embed `", "syntax error: expected a valid file path, found \"`\". (line: 1, column: 9)"},
	{`@embed \'`, `syntax error: expected a valid file path, found "'". (line: 1, column: 10)`},
	{`@embed \|`, `syntax error: expected a valid file path, found "|". (line: 1, column: 10)`},
	{`@embed \\`, `syntax error: expected a valid file path, found "\\". (line: 1, column: 10)`},
	{`@embed :`, `syntax error: expected a valid file path, found ":". (line: 1, column: 9)`},
	{`@embed ""`, `syntax error: expected a file path, found empty string. (line: 1, column: 10)`},
	{`@embed ''`, `syntax error: expected a file path, found empty string. (line: 1, column: 10)`},
}

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
		&ast.Function{
			Name: "foo",
			Body: []ast.Statement{
				ast.Embed{"file"},
				ast.Embed{"file"},
			},
		},
	}},
}

var embeddingErrorHandlingCases = []errorHandlingTestCase{
	{`@embed`, "main.sh(1:7): syntax error: expected a blank after the @embed directive, found end of file."},
	{`@embed `, "main.sh(1:8): syntax error: unexpected token: end of file."},
	{`@embed $var`, "main.sh(1:12): syntax error: expected a valid file path."},
	{`@embed "$var foo"`, "main.sh(1:18): syntax error: expected a valid file path."},
	{`@embed file | cmd`, "main.sh(1:13): syntax error: expected a valid file path."},
	{`@embed /foo/bar`, `main.sh(1:16): syntax error: path cannot start or end with slash, "/foo/bar".`},
	{`@embed foo/bar/`, `main.sh(1:16): syntax error: path cannot start or end with slash, "foo/bar/".`},
	{`@embed \"`, `main.sh(1:10): syntax error: expected a valid file path, found "\"".`},
	{`@embed *`, `main.sh(1:9): syntax error: expected a valid file path, found "*".`},
	{`@embed \<`, `main.sh(1:10): syntax error: expected a valid file path, found "<".`},
	{`@embed \>`, `main.sh(1:10): syntax error: expected a valid file path, found ">".`},
	{`@embed \?`, `main.sh(1:10): syntax error: expected a valid file path, found "?".`},
	{"@embed `", "main.sh(1:9): syntax error: expected a valid file path, found \"`\"."},
	{`@embed \'`, `main.sh(1:10): syntax error: expected a valid file path, found "'".`},
	{`@embed \|`, `main.sh(1:10): syntax error: expected a valid file path, found "|".`},
	{`@embed \\`, `main.sh(1:10): syntax error: expected a valid file path, found "\\".`},
	{`@embed :`, `main.sh(1:9): syntax error: expected a valid file path, found ":".`},
	{`@embed ""`, `main.sh(1:10): syntax error: expected a file path, found empty string.`},
	{`@embed ''`, `main.sh(1:10): syntax error: expected a file path, found empty string.`},
}

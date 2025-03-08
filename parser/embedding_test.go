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
	`, ast.Script{
		ast.Embed{"file"},
		ast.Embed{"file"},
		ast.Command{Name: ast.Word("command")},
	}},
}

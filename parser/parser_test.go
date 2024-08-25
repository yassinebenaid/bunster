package parser_test

import (
	"reflect"
	"testing"

	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/parser"
)

var dump = (&godump.Dumper{
	Theme: godump.DefaultTheme,
}).Sprintln

func TestCanParseCommandCall(t *testing.T) {
	testCases := []struct {
		input    string
		expected ast.Script
	}{
		{`git`, ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word{Value: "git"}},
			},
		}},
		{`foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word{Value: "foo"},
					Args: []ast.Node{
						ast.Word{Value: "bar"},
						ast.Word{Value: "baz"},
					},
				},
			},
		}},
		{`foo $bar $FOO_BAR_1234567890`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word{Value: "foo"},
					Args: []ast.Node{
						ast.SimpleExpansion{Name: "bar"},
						ast.SimpleExpansion{Name: "FOO_BAR_1234567890"},
					},
				},
			},
		}},
		{`/usr/bin/foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word{Value: "/usr/bin/foo"},
					Args: []ast.Node{
						ast.Word{Value: "bar"},
						ast.Word{Value: "baz"},
					},
				},
			},
		}},
		{`/usr/bin/foo-bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word{Value: "/usr/bin/foo-bar"},
					Args: []ast.Node{
						ast.Word{Value: "baz"},
					},
				},
			},
		}},
		{`/usr/bin/$BINARY_NAME --option -f --do=something`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Concatination{
						Nodes: []ast.Node{
							ast.Word{Value: "/usr/bin/"},
							ast.SimpleExpansion{Name: "BINARY_NAME"},
						},
					},
					Args: []ast.Node{
						ast.Word{Value: "--option"},
						ast.Word{Value: "-f"},
						ast.Word{Value: "--do=something"},
					},
				},
			},
		}},
	}

	for i, tc := range testCases {
		p := parser.New(
			lexer.New([]byte(tc.input)),
		)

		script := p.ParseScript()

		if !reflect.DeepEqual(script, tc.expected) {
			t.Fatalf("\nCase #%d: the script is not as expected:\n\nwant:\n%s\ngot:\n%s", i, dump(tc.expected), dump(script))
		}

	}
}

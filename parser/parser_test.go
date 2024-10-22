package parser_test

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/lexer"
	"github.com/yassinebenaid/bunny/parser"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

type testCase struct {
	input    string
	expected ast.Script
}

var testCases = []struct {
	label string
	cases []testCase
}{
	{"Simle Command calls", []testCase{
		{``, ast.Script{}},
		{`	 	`, ast.Script{}},
		{"\n	\n \n ", ast.Script{}},
		{`git`, ast.Script{Statements: []ast.Node{ast.Command{Name: ast.Word("git")}}}},
		{`foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("foo"),
					Args: []ast.Node{ast.Word("bar"), ast.Word("baz")},
				},
			},
		}},
		{`foo $bar $FOO_BAR_1234567890`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("foo"),
					Args: []ast.Node{
						ast.SimpleExpansion("bar"),
						ast.SimpleExpansion("FOO_BAR_1234567890"),
					},
				},
			},
		}},
		{`/usr/bin/foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("/usr/bin/foo"),
					Args: []ast.Node{
						ast.Word("bar"),
						ast.Word("baz"),
					},
				},
			},
		}},
		{`/usr/bin/foo-bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("/usr/bin/foo-bar"),
					Args: []ast.Node{ast.Word("baz")},
				},
			},
		}},
		{"cmd1 \n cmd2", ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
			},
		}},
	}},
	{"Simle Command calls", []testCase{
		{``, ast.Script{}},
		{`	 	`, ast.Script{}},
		{"\n	\n \n ", ast.Script{}},
		{`git`, ast.Script{Statements: []ast.Node{ast.Command{Name: ast.Word("git")}}}},
		{`foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("foo"),
					Args: []ast.Node{ast.Word("bar"), ast.Word("baz")},
				},
			},
		}},
		{`foo $bar $FOO_BAR_1234567890`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("foo"),
					Args: []ast.Node{
						ast.SimpleExpansion("bar"),
						ast.SimpleExpansion("FOO_BAR_1234567890"),
					},
				},
			},
		}},
		{`/usr/bin/foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("/usr/bin/foo"),
					Args: []ast.Node{
						ast.Word("bar"),
						ast.Word("baz"),
					},
				},
			},
		}},
		{`/usr/bin/foo-bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("/usr/bin/foo-bar"),
					Args: []ast.Node{ast.Word("baz")},
				},
			},
		}},
		{"cmd1\n cmd2\ncmd3\n cmd4 arg1 arg2\ncmd5", ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
				ast.Command{Name: ast.Word("cmd3")},
				ast.Command{Name: ast.Word("cmd4"), Args: []ast.Node{ast.Word("arg1"), ast.Word("arg2")}},
				ast.Command{Name: ast.Word("cmd5")},
			},
		}},
		{"cmd1; cmd2;cmd3; cmd4 arg1 arg2;cmd5", ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("cmd1")},
				ast.Command{Name: ast.Word("cmd2")},
				ast.Command{Name: ast.Word("cmd3")},
				ast.Command{Name: ast.Word("cmd4"), Args: []ast.Node{ast.Word("arg1"), ast.Word("arg2")}},
				ast.Command{Name: ast.Word("cmd5")},
			},
		}},
	}},

	{"Strings", []testCase{
		{`cmd 'hello world'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("hello world"),
					},
				},
			},
		}},
		{`cmd 'if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset"),
					},
				},
			},
		}},
		{`cmd '+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word(`+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~`),
					},
				},
			},
		}},
		{`cmd '' '\' '$foo'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word(""),
						ast.Word(`\`),
						ast.Word(`$foo`),
					},
				},
			},
		}},
		{`cmd ""`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word(""),
					},
				},
			},
		}},
		{`cmd "Hello World" "name is: $NAME and path is $DIR/$FILE"`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("Hello World"),
						ast.Concatination{
							Nodes: []ast.Node{
								ast.Word("name is: "),
								ast.SimpleExpansion("NAME"),
								ast.Word(" and path is "),
								ast.SimpleExpansion("DIR"),
								ast.Word("/"),
								ast.SimpleExpansion("FILE"),
							},
						},
					},
				},
			},
		}},
		{`cmd "\"" "\$ESCAPED_VAR" "\foo\bar\\" \$var \" \foo`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word(`"`),
						ast.Word(`$ESCAPED_VAR`),
						ast.Word(`\foo\bar\`),
						ast.Word(`$var`),
						ast.Word(`"`),
						ast.Word(`foo`),
					},
				},
			},
		}},
		{"cmd \"\\\nfoo\"", ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("cmd"), Args: []ast.Node{ast.Word(`foo`)}},
			},
		}},
		{`/usr/bin/$BINARY_NAME --path=/home/$USER/dir --option -f --do=something $HOME$DIR_NAME$PKG_NAME/foo`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Concatination{
						Nodes: []ast.Node{
							ast.Word("/usr/bin/"),
							ast.SimpleExpansion("BINARY_NAME"),
						},
					},
					Args: []ast.Node{
						ast.Concatination{
							Nodes: []ast.Node{
								ast.Word("--path=/home/"),
								ast.SimpleExpansion("USER"),
								ast.Word("/dir"),
							},
						},
						ast.Word("--option"),
						ast.Word("-f"),
						ast.Word("--do=something"),
						ast.Concatination{
							Nodes: []ast.Node{
								ast.SimpleExpansion("HOME"),
								ast.SimpleExpansion("DIR_NAME"),
								ast.SimpleExpansion("PKG_NAME"),
								ast.Word("/foo"),
							},
						},
					},
				},
			},
		}},
		{`cmd 'foo''bar' "foo""bar" "foo"'bar' "'foo'"`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("foobar"),
						ast.Word("foobar"),
						ast.Word("foobar"),
						ast.Word("'foo'"),
					},
				},
			},
		}},
		{"cmd \"\n\"", ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("\n"),
					},
				},
			},
		}},
	}},
	{"Comments", []testCase{
		{`#`, ast.Script{}},
		{`# foo bar`, ast.Script{}},
		{`	 # foo bar`, ast.Script{}},
		{"# foo bar    \n    \t # baz", ast.Script{}},
		{"cmd # comment", ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("cmd")},
			},
		}},
		{"cmd#not-comment arg#not-comment", ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("cmd#not-comment"), Args: []ast.Node{ast.Word("arg#not-comment")}},
			},
		}},
	}},
	{"Redirections", redirectionTests},
	{"Piplines", pipesTests},
	{"Binary Constructions", logicalCommandsTests},
	{"Background Constructions", []testCase{
		{`cmd & cmd2`, ast.Script{
			Statements: []ast.Node{
				ast.BackgroundConstruction{
					Node: ast.Command{Name: ast.Word("cmd")},
				},
				ast.Command{Name: ast.Word("cmd2")},
			},
		}},
		{`cmd && cmd2 & cmd3 && cmd4&`, ast.Script{
			Statements: []ast.Node{
				ast.BackgroundConstruction{
					Node: ast.BinaryConstruction{
						Left:     ast.Command{Name: ast.Word("cmd")},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd2")},
					},
				},
				ast.BackgroundConstruction{
					Node: ast.BinaryConstruction{
						Left:     ast.Command{Name: ast.Word("cmd3")},
						Operator: "&&",
						Right:    ast.Command{Name: ast.Word("cmd4")},
					},
				},
			},
		}},
		{` cmd | cmd2 |& cmd3 | cmd4 |& cmd5 foo& cmd | cmd2 |& cmd3 | cmd4 |& cmd5`, ast.Script{
			Statements: []ast.Node{
				ast.BackgroundConstruction{
					Node: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
						{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
						{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
						{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: false},
						{Command: ast.Command{Name: ast.Word("cmd5"), Args: []ast.Node{ast.Word("foo")}}, Stderr: true},
					},
				},
				ast.Pipeline{
					{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
					{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
					{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
					{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: false},
					{Command: ast.Command{Name: ast.Word("cmd5")}, Stderr: true},
				},
			},
		}},
	}},
	{"Loops", loopsTests},
	{"Conditionals", conditionalsTests},
}

func TestParser(t *testing.T) {
	tgroup, tcase := os.Getenv("TEST_GROUP"), os.Getenv("TEST_CASE")

	for _, group := range testCases {
		if tgroup != "" && !strings.Contains(strings.ToLower(group.label), tgroup) {
			continue
		}

		for i, tc := range group.cases {
			if tcase != "" && fmt.Sprint(i) != tcase {
				continue
			}

			p := parser.New(
				lexer.New([]byte(tc.input)),
			)

			script := p.ParseScript()

			if p.Error != nil {
				t.Fatalf("\nGroup: %sCase: %s\nUnexpected Error: %s\n", dump(group.label), dump(i), dump(p.Error.Error()))
			}

			if !reflect.DeepEqual(script, tc.expected) {
				t.Fatalf("\nGroup: %sCase: %s\nwant:\n%s\ngot:\n%s", dump(group.label), dump(i), dump(tc.expected), dump(script))
			}
		}
	}
}

type errorHandlingTestCase struct {
	input string
	err   string
}

var errorHandlingTestCases = []struct {
	label string
	cases []errorHandlingTestCase
}{
	{"Simple Commands", []errorHandlingTestCase{
		{`|`, `syntax error: invalid command construction.`},
		{`>`, `syntax error: invalid command construction.`},
		{`>>`, `syntax error: invalid command construction.`},
		{`1>>`, `syntax error: invalid command construction.`},
		{`<<<`, `syntax error: invalid command construction.`},
		{`1<<<`, `syntax error: invalid command construction.`},
		{`1>`, `syntax error: invalid command construction.`},
		{`1<`, `syntax error: invalid command construction.`},
		{`1>&`, `syntax error: invalid command construction.`},
		{`1<&`, `syntax error: invalid command construction.`},
		{`&& cmd2`, `syntax error: invalid command construction.`},
		{`|| cmd2`, `syntax error: invalid command construction.`},

		{`& cmd2`, `syntax error: invalid command construction.`},
		{`cmd & || cmd2`, `syntax error: invalid command construction.`},
		{`cmd & && cmd2`, `syntax error: invalid command construction.`},
		{`cmd & | cmd2`, `syntax error: invalid command construction.`},
		{`cmd || & cmd2`, `syntax error: invalid command construction.`},
		{`cmd && & cmd2`, `syntax error: invalid command construction.`},
		{`cmd | & cmd2`, "syntax error: invalid pipeline construction, a command is missing after `|`."},

		{"cmd \n || cmd2", `syntax error: invalid command construction.`},
		{"cmd \n && cmd2", `syntax error: invalid command construction.`},
		{"cmd \n | cmd2", `syntax error: invalid command construction.`},

		{`; cmd2`, `syntax error: invalid command construction.`},
		{`cmd ; || cmd2`, `syntax error: invalid command construction.`},
		{`cmd ; && cmd2`, `syntax error: invalid command construction.`},
		{`cmd ; | cmd2`, `syntax error: invalid command construction.`},
		{`cmd || ; cmd2`, `syntax error: invalid command construction.`},
		{`cmd && ; cmd2`, `syntax error: invalid command construction.`},
		{`cmd | ; cmd2`, "syntax error: invalid pipeline construction, a command is missing after `|`."},
		{`cmd ;;`, "syntax error: invalid command construction."},
	}},
	{"Quotes", []errorHandlingTestCase{
		{`cmd 'foo bar`, `syntax error: a closing single quote is missing.`},
		{`cmd "foo bar'`, `syntax error: a closing double quote is missing.`},
	}},
	{"Redirections", redirectionErrorHandlingCases},
	{"Pipes", pipesErrorHandlingCases},
	{"Logical Constructions", logicalCommandsErrorHandlingCases},
	{"Loops", loopsErrorHandlingCases},
	{"Conditionals", conditionalsErrorHandlingCases},
}

func TestParserErrorHandling(t *testing.T) {
	tgroup, tcase := os.Getenv("TEST_GROUP"), os.Getenv("TEST_CASE")

	for _, group := range errorHandlingTestCases {
		if tgroup != "" && !strings.Contains(strings.ToLower(group.label), tgroup) {
			continue
		}

		for i, tc := range group.cases {
			if tcase != "" && fmt.Sprint(i) != tcase {
				continue
			}

			p := parser.New(
				lexer.New([]byte(tc.input)),
			)

			p.ParseScript()

			if p.Error == nil {
				t.Fatalf("\nGroup: %s\nCase#%d: Expected Error, got nil\n", group.label, i)
			}

			if p.Error.Error() != tc.err {
				t.Fatalf("\nGroup: %s\nCase: %s\nwant:\n%s\ngot:\n%s", dump(group.label), dump(i), dump(tc.err), dump(p.Error.Error()))
			}
		}
	}
}

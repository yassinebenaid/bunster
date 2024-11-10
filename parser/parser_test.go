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
		{``, nil},
		{`	 	`, nil},
		{"\n	\n \n ", nil},
		{`git`, ast.Script{ast.Command{Name: ast.Word("git")}}},
		{`foo bar baz`, ast.Script{

			ast.Command{
				Name: ast.Word("foo"),
				Args: []ast.Expression{ast.Word("bar"), ast.Word("baz")},
			},
		}},
		{`foo $bar $FOO_BAR_1234567890`, ast.Script{

			ast.Command{
				Name: ast.Word("foo"),
				Args: []ast.Expression{
					ast.Var("bar"),
					ast.Var("FOO_BAR_1234567890"),
				},
			},
		}},
		{`/usr/bin/foo bar baz`, ast.Script{

			ast.Command{
				Name: ast.Word("/usr/bin/foo"),
				Args: []ast.Expression{
					ast.Word("bar"),
					ast.Word("baz"),
				},
			},
		}},
		{`/usr/bin/foo-bar baz`, ast.Script{

			ast.Command{
				Name: ast.Word("/usr/bin/foo-bar"),
				Args: []ast.Expression{ast.Word("baz")},
			},
		}},
		{"cmd1 \n cmd2", ast.Script{

			ast.Command{Name: ast.Word("cmd1")},
			ast.Command{Name: ast.Word("cmd2")},
		}},
		{"cmd1\n cmd2\ncmd3\n cmd4 arg1 arg2\ncmd5", ast.Script{

			ast.Command{Name: ast.Word("cmd1")},
			ast.Command{Name: ast.Word("cmd2")},
			ast.Command{Name: ast.Word("cmd3")},
			ast.Command{Name: ast.Word("cmd4"), Args: []ast.Expression{ast.Word("arg1"), ast.Word("arg2")}},
			ast.Command{Name: ast.Word("cmd5")},
		}},
		{"cmd1; cmd2;cmd3; cmd4 arg1 arg2;cmd5", ast.Script{

			ast.Command{Name: ast.Word("cmd1")},
			ast.Command{Name: ast.Word("cmd2")},
			ast.Command{Name: ast.Word("cmd3")},
			ast.Command{Name: ast.Word("cmd4"), Args: []ast.Expression{ast.Word("arg1"), ast.Word("arg2")}},
			ast.Command{Name: ast.Word("cmd5")},
		}},
	}},

	{"Strings", []testCase{
		{`cmd 'hello world'`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word("hello world"),
				},
			},
		}},
		{`cmd 'if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset'`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word("if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset"),
				},
			},
		}},
		{`cmd '+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~'`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word(`+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~`),
				},
			},
		}},
		{`cmd '' '\' '$foo'`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word(""),
					ast.Word(`\`),
					ast.Word(`$foo`),
				},
			},
		}},
		{`cmd ""`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word(""),
				},
			},
		}},
		{`cmd "Hello World" "name is: $NAME and path is $DIR/$FILE"`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word("Hello World"),
					ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word("name is: "),
							ast.Var("NAME"),
							ast.Word(" and path is "),
							ast.Var("DIR"),
							ast.Word("/"),
							ast.Var("FILE"),
						},
					},
				},
			},
		}},
		{`cmd "\"" "\$ESCAPED_VAR" "\foo\bar\\" \$var \" \foo`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word(`"`),
					ast.Word(`$ESCAPED_VAR`),
					ast.Word(`\foo\bar\`),
					ast.Word(`$var`),
					ast.Word(`"`),
					ast.Word(`foo`),
				},
			},
		}},
		{"cmd \"\\\nfoo\"", ast.Script{

			ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word(`foo`)}},
		}},
		{`/usr/bin/$BINARY_NAME --path=/home/$USER/dir --option -f --do=something $HOME$DIR_NAME$PKG_NAME/foo`, ast.Script{

			ast.Command{
				Name: ast.Concatination{
					Nodes: []ast.Expression{
						ast.Word("/usr/bin/"),
						ast.Var("BINARY_NAME"),
					},
				},
				Args: []ast.Expression{
					ast.Concatination{
						Nodes: []ast.Expression{
							ast.Word("--path=/home/"),
							ast.Var("USER"),
							ast.Word("/dir"),
						},
					},
					ast.Word("--option"),
					ast.Word("-f"),
					ast.Word("--do=something"),
					ast.Concatination{
						Nodes: []ast.Expression{
							ast.Var("HOME"),
							ast.Var("DIR_NAME"),
							ast.Var("PKG_NAME"),
							ast.Word("/foo"),
						},
					},
				},
			},
		}},
		{`cmd 'foo''bar' "foo""bar" "foo"'bar' "'foo'"`, ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word("foobar"),
					ast.Word("foobar"),
					ast.Word("foobar"),
					ast.Word("'foo'"),
				},
			},
		}},
		{"cmd \"\n\"", ast.Script{

			ast.Command{
				Name: ast.Word("cmd"),
				Args: []ast.Expression{
					ast.Word("\n"),
				},
			},
		}},
	}},
	{"Comments", []testCase{
		{`#`, nil},
		{`# foo bar`, nil},
		{`	 # foo bar`, nil},
		{"# foo bar    \n    \t # baz", nil},
		{"cmd # comment", ast.Script{

			ast.Command{Name: ast.Word("cmd")},
		}},
		{"cmd#not-comment arg#not-comment", ast.Script{

			ast.Command{Name: ast.Word("cmd#not-comment"), Args: []ast.Expression{ast.Word("arg#not-comment")}},
		}},
	}},
	{"Redirections", redirectionTests},
	{"Piplines", pipesTests},
	{"Binary Constructions", logicalCommandsTests},
	{"Background Constructions", []testCase{
		{`cmd & cmd2`, ast.Script{

			ast.BackgroundConstruction{
				Statement: ast.Command{Name: ast.Word("cmd")},
			},
			ast.Command{Name: ast.Word("cmd2")},
		}},
		{`cmd && cmd2 & cmd3 && cmd4&`, ast.Script{

			ast.BackgroundConstruction{
				Statement: ast.BinaryConstruction{
					Left:     ast.Command{Name: ast.Word("cmd")},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd2")},
				},
			},
			ast.BackgroundConstruction{
				Statement: ast.BinaryConstruction{
					Left:     ast.Command{Name: ast.Word("cmd3")},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd4")},
				},
			},
		}},
		{` cmd | cmd2 |& cmd3 | cmd4 |& cmd5 foo& cmd | cmd2 |& cmd3 | cmd4 |& cmd5`, ast.Script{

			ast.BackgroundConstruction{
				Statement: ast.Pipeline{
					{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
					{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
					{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
					{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: false},
					{Command: ast.Command{Name: ast.Word("cmd5"), Args: []ast.Expression{ast.Word("foo")}}, Stderr: true},
				},
			},
			ast.Pipeline{
				{Command: ast.Command{Name: ast.Word("cmd")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd2")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd3")}, Stderr: true},
				{Command: ast.Command{Name: ast.Word("cmd4")}, Stderr: false},
				{Command: ast.Command{Name: ast.Word("cmd5")}, Stderr: true},
			},
		}},
	}},
	{"Loops", loopsTests},
	{"Conditionals", conditionalsTests},
	{"Case", caseTests},

	{"Command Group", groupingTests},
	{"Subsitutions", substitutionTests},
	{"Parameter Expansion", parameterExpansionTests},
	{"Arithmetics", arithmeticsTests},
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
		{`)`, "syntax error: `)` has a special meaning here and cannot be used as a command name."},
		{`|`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},
		{`>`, "syntax error: `>` has a special meaning here and cannot be used as a command name."},
		{`>>`, "syntax error: `>>` has a special meaning here and cannot be used as a command name."},
		{`1>>`, "syntax error: `1` has a special meaning here and cannot be used as a command name."},
		{`<<<`, "syntax error: `<<<` has a special meaning here and cannot be used as a command name."},
		{`1<<<`, "syntax error: `1` has a special meaning here and cannot be used as a command name."},
		{`1>`, "syntax error: `1` has a special meaning here and cannot be used as a command name."},
		{`1<`, "syntax error: `1` has a special meaning here and cannot be used as a command name."},
		{`1>&`, "syntax error: `1` has a special meaning here and cannot be used as a command name."},
		{`1<&`, "syntax error: `1` has a special meaning here and cannot be used as a command name."},
		{`&& cmd2`, "syntax error: `&&` has a special meaning here and cannot be used as a command name."},
		{`|| cmd2`, "syntax error: `||` has a special meaning here and cannot be used as a command name."},

		{`& cmd2`, "syntax error: `&` has a special meaning here and cannot be used as a command name."},
		{`cmd & || cmd2`, "syntax error: `||` has a special meaning here and cannot be used as a command name."},
		{`cmd & && cmd2`, "syntax error: `&&` has a special meaning here and cannot be used as a command name."},
		{`cmd & | cmd2`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},
		{`cmd || & cmd2`, "syntax error: `&` has a special meaning here and cannot be used as a command name."},
		{`cmd && & cmd2`, "syntax error: `&` has a special meaning here and cannot be used as a command name."},
		{`cmd | & cmd2`, "syntax error: `&` has a special meaning here and cannot be used as a command name."},

		{"cmd \n || cmd2", "syntax error: `||` has a special meaning here and cannot be used as a command name."},
		{"cmd \n && cmd2", "syntax error: `&&` has a special meaning here and cannot be used as a command name."},
		{"cmd \n | cmd2", "syntax error: `|` has a special meaning here and cannot be used as a command name."},

		{`; cmd2`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
		{`cmd ; || cmd2`, "syntax error: `||` has a special meaning here and cannot be used as a command name."},
		{`cmd ; && cmd2`, "syntax error: `&&` has a special meaning here and cannot be used as a command name."},
		{`cmd ; | cmd2`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},
		{`cmd || ; cmd2`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
		{`cmd && ; cmd2`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
		{`cmd | ; cmd2`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
		{`cmd ;;`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
	}},
	{"Quotes", []errorHandlingTestCase{
		{`cmd 'foo bar`, `syntax error: a closing single quote is missing.`},
		{`cmd "foo bar'`, `syntax error: a closing double quote is missing.`},
	}},
	{"Redirections", redirectionErrorHandlingCases},
	{"Pipes", pipesErrorHandlingCases},
	{"Logical Constructions", logicalCommandsErrorHandlingCases},
	{"Loops", loopsErrorHandlingCases},
	{"Conditionals", ifErrorHandlingCases},
	{"Case", caseErrorHandlingCases},
	{"Command Group", groupingErrorHandlingCases},
	{"Substitution", substitutionErrorHandlingCases},
	{"Arithmetics", arithmeticsErrorHandlingCases},
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

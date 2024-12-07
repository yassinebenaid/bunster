package generator_test

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/ir"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

type testCase struct {
	input    string
	expected ir.Program
}

var testCases = []struct {
	label string
	cases []testCase
}{
	{"Simple Commands", []testCase{}},
}

func TestGenerator(t *testing.T) {
	tgroup, tcase := os.Getenv("TEST_GROUP"), os.Getenv("TEST_CASE")

	for _, group := range testCases {
		if tgroup != "" && !strings.Contains(strings.ToLower(group.label), tgroup) {
			continue
		}

		for i, tc := range group.cases {
			if tcase != "" && fmt.Sprint(i) != tcase {
				continue
			}

			script, err := parser.Parse(
				lexer.New([]byte(tc.input)),
			)

			if err != nil {
				t.Fatalf("\nGroup: %sCase: %sInput: %s\nUnexpected Error: %s\n", dump(group.label), dump(i), dump(tc.input), dump(err.Error()))
			}

			program := generator.Generate(script)

			if !reflect.DeepEqual(program, tc.expected) {
				t.Fatalf("\nGroup: %sCase: %sInput: %s\nWant:\n%s\nGot:\n%s", dump(group.label), dump(i), dump(tc.input), dump(tc.expected), dump(program))
			}
		}
	}
}

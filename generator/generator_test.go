package generator_test

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/ryuko/generator"
	"github.com/yassinebenaid/ryuko/ir"
	"github.com/yassinebenaid/ryuko/lexer"
	"github.com/yassinebenaid/ryuko/parser"
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
	{"Simple Commands", []testCase{
		{"cmd", ir.Program{Instructions: []ir.Instruction{
			ir.Assign{Name: "cmd_1_name", Value: ir.String("cmd")},
			ir.InitCommand{Name: "cmd_1"},
			ir.RunCommanOrFail{Name: "cmd_1"},
		}}},
	}},
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

			script, err := parser.Parse(
				lexer.New([]byte(tc.input)),
			)

			if err != nil {
				t.Fatalf("\nGroup: %sCase: %sInput: %s\nUnexpected Error: %s\n", dump(group.label), dump(i), dump(tc.input), dump(err.Error()))
			}

			program := generator.Generate(script)

			if !reflect.DeepEqual(program, tc.expected) {
				t.Fatalf("\nGroup: %sCase: %sInput: %s\nwant:\n%s\ngot:\n%s", dump(group.label), dump(i), dump(tc.input), dump(tc.expected), dump(script))
			}
		}
	}
}

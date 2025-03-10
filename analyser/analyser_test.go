package analyser_test

import (
	"testing"

	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

type testCase struct {
	input string
	error string
}

var testCases = []testCase{
	{`name=foo | cmd`, "semantic error: using shell parameters within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines. (line: 0, column: 0)"},
	{`break`, "semantic error: The `break` keyword cannot be used here. (line: 0, column: 0)"},
	{`continue`, "semantic error: The `continue` keyword cannot be used here. (line: 0, column: 0)"},
	{`local var`, "semantic error: The `local` keyword cannot be used outside functions. (line: 0, column: 0)"},
	{`wait | cmd`, "semantic error: using 'wait' command within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines. (line: 0, column: 0)"},
	{`func(){ local var | cmd; }`, "semantic error: using 'local' command within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines. (line: 0, column: 0)"},
	{`export var | cmd;`, "semantic error: using 'export' command within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines. (line: 0, column: 0)"},
	{`
	func(){
		@embed file
		}
		`, "semantic error: using '@embed' directive is only valid in global scope. (line: 0, column: 0)"},
	{`@embed ../file`, "semantic error: the path \"../file\" is not local. (line: 0, column: 0)"},
}

func TestErrors(t *testing.T) {
	for i, tc := range testCases {
		script, err := parser.Parse(
			lexer.New([]byte(tc.input)),
		)

		if err != nil {
			t.Fatalf("\nCase: %sInput: %s\nUnexpected Error: %s\n", dump(i), dump(tc.input), dump(err.Error()))
		}

		analysisError := analyser.Analyse(script)

		if analysisError == nil {
			t.Fatalf("\nCase: %sInput: %s\nExpected Error, got nil\n", dump(i), dump(tc.input))
		}

		if analysisError.Error() != tc.error {
			t.Fatalf("\nCase: %sInput: %s\nwant:\n%s\ngot:\n%s", dump(i), dump(tc.input), dump(tc.error), dump(analysisError.Error()))
		}
	}
}

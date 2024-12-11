package tst_test

import (
	"reflect"
	"testing"

	"github.com/yassinebenaid/bunster/pkg/tst"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestParser_Parse(t *testing.T) {
	testCases := []struct {
		input    string
		expected []tst.Test
		err      error
	}{
		{
			input: `
#(TEST: foo bar)

foo bar

  baz

#(RESULT)

foo bar
	whatever

#(ENDTEST)`,
		},
	}

	for i, tc := range testCases {
		tests, parseErr := tst.Parse([]byte(tc.input))

		if tc.err != nil {
			if parseErr == tc.err {
				t.Errorf("expected:\n%sgot:\n%s", dump(tc.err), dump(parseErr))
			}
			return
		}

		if parseErr != nil {
			t.Errorf("unexpected error: %v", parseErr)
			return
		}

		if !reflect.DeepEqual(tc.expected, tests) {
			t.Errorf("Case: %d\nExpected:\n%sGot:\n%s", i, dump(tc.expected), dump(tests))
		}
	}
}

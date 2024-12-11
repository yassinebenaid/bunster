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
		err      string
	}{
		{
			input: `
#(TEST: foo bar)

foo bar
 
#(RESULT)
 
foo bar
 
#(ENDTEST)`,
			expected: []tst.Test{{
				Label:  `foo bar`,
				Input:  "\nfoo bar\n \n",
				Output: " \nfoo bar\n \n",
			}},
		},
	}

	for i, tc := range testCases {
		tests, parseErr := tst.Parse(tc.input)

		if tc.err != "" {
			if parseErr == nil || parseErr.Error() != tc.err {
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

func TestParseErrors(t *testing.T) {
	testCases := []struct {
		input string
		err   string
	}{
		{
			input: `foo bar`,
			err:   `bad test syntax, coundl't find test header '#(TEST: ...)', found "foo bar"`,
		},
		{
			input: `#(TEST: foo bar `,
			err:   `bad test syntax, unclosed test header '#(TEST: ...)'`,
		},
		{
			input: `#(TEST: foo bar)`,
			err:   "bad test syntax, coundl't find #(RESULT) section",
		},
		{
			input: `#(TEST: foo bar) extra-string`,
			err:   "bad test syntax, unclosed test header '#(TEST: ...)'",
		},
	}

	for _, tc := range testCases {
		_, parseErr := tst.Parse(tc.input)

		if parseErr == nil || parseErr.Error() != tc.err {
			t.Errorf("expected:\n%sgot:\n%s", dump(tc.err), dump(parseErr))
		}
	}
}

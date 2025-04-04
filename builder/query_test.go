package builder

import (
	"testing"

	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestParserQuery(t *testing.T) {

	cases := []struct {
		input string
		query query
		err   string
	}{
		{input: "", err: `module path "" is not in an expected format`},
		{input: "foo", err: `module path "foo" is not in an expected format`},
		{input: "foo/bar", err: `module path "foo/bar" is not in an expected format`},
		{input: "foo.com/bar/baz", err: `module path "foo.com/bar/baz" is not in an expected format`},
		{input: "foo.com/bar/baz@", err: `module path "foo.com/bar/baz@" is not in an expected format`},
		{input: "foo.com/bar/baz@bad-commit-hash", err: `module path "foo.com/bar/baz@bad-commit-hash" is not in an expected format`},
		{
			input: "foo.com/bar/baz@a86d7b5f7d93c8a6263979dc9e7a2beb1124acce",
			query: query{
				module: "foo.com/bar/baz",
				commit: "a86d7b5f7d93c8a6263979dc9e7a2beb1124acce",
			},
		},
	}

	for _, tc := range cases {
		q, err := parseQuery(tc.input)
		if err != nil {
			if err.Error() != tc.err {
				t.Fatalf("Input: %v Unexpected Error: %v", dump(tc.input), dump(err.Error()))
			}
			continue
		}

		if tc.err != "" {
			t.Fatalf("Input: %v Expected Error, go nil", dump(tc.input))
		}

		if q != tc.query {
			t.Fatalf("Input: %v Expected: %v Got: %v", dump(tc.input), dump(tc.query), dump(q))
		}
	}
}

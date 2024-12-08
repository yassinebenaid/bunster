package tst_test

import (
	"strings"
	"testing"

	"github.com/yassinebenaid/bunster/pkg/tst"
)

func TestParser_Parse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *tst.Test
		wantErr  bool
	}{
		{
			name: "Valid single case",
			input: `Test: Simple Commands
------------
------input------
go mod tidy
------output------
package main
func main(){ } `,
			expected: &tst.Test{
				Label: "Simple Commands",
				Cases: []tst.TestCase{
					{
						Input:  "go mod tidy",
						Output: "package main\nfunc main(){ } ",
					},
				},
			},
		},
		{
			name: "Valid multiple cases",
			input: `Test: Multiple Commands
------------
------input------
go mod tidy
------output------
package main
func main(){ }
------------
------input------
go run main.go
------output------
Hello, World!`,
			expected: &tst.Test{
				Label: "Multiple Commands",
				Cases: []tst.TestCase{
					{
						Input:  "go mod tidy",
						Output: "package main\nfunc main(){ }",
					},
					{
						Input:  "go run main.go",
						Output: "Hello, World!",
					},
				},
			},
		},
		{
			name: "Invalid - missing Test:",
			input: `Simple Commands
------------
------input------
go mod tidy`,
			wantErr: true,
		},
		{
			name: "Invalid - output before input",
			input: `Test: Invalid Command
------------
------output------
package main
------input------
go mod tidy`,
			wantErr: true,
		},
	}

	parser := &tst.Parser{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parser.Parse(strings.NewReader(tc.input))

			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Compare results
			if result.Label != tc.expected.Label {
				t.Errorf("label mismatch: got %q, want %q", result.Label, tc.expected.Label)
			}

			if len(result.Cases) != len(tc.expected.Cases) {
				t.Errorf("number of cases mismatch: got %d, want %d", len(result.Cases), len(tc.expected.Cases))
				return
			}

			for i, case_ := range result.Cases {
				if case_.Input != tc.expected.Cases[i].Input {
					t.Errorf("input mismatch for case %d: got %q, want %q", i, case_.Input, tc.expected.Cases[i].Input)
				}
				if case_.Output != tc.expected.Cases[i].Output {
					t.Errorf("output mismatch for case %d: got %q, want %q", i, case_.Output, tc.expected.Cases[i].Output)
				}
			}
		})
	}
}

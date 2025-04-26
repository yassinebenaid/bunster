package flag

import (
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		flagSetup     func(*Parser)
		args          []string
		expectedFlags map[string]any
		expectedArgs  []string
		expectErr     string
	}{
		{
			name: "basic short boolean flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, false)
				p.AddShortFlag("b", Boolean, false)
				p.AddShortFlag("c", Boolean, false)
			},
			args: []string{"-a", "-c"},
			expectedFlags: map[string]any{
				"a": true,
				"b": false,
				"c": true,
			},
		},
		{
			name: "boolean short flags are optional",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, true)
				p.AddShortFlag("b", Boolean, true)
				p.AddShortFlag("c", Boolean, true)
			},
			args: []string{"-a", "-c"},
			expectedFlags: map[string]any{
				"a": true,
				"b": false,
				"c": true,
			},
		},
		{
			name: "can group boolean flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, true)
				p.AddShortFlag("b", Boolean, true)
				p.AddShortFlag("c", Boolean, true)
				p.AddShortFlag("d", Boolean, true)
			},
			args: []string{"-abd"},
			expectedFlags: map[string]any{
				"a": true,
				"b": true,
				"c": false,
				"d": true,
			},
		},
		{
			name: "short string flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", String, false)
			},
			args: []string{"-a", "foo", "-b", "bar"},
			expectedFlags: map[string]any{
				"a": "foo",
				"b": "bar",
			},
		},
		{
			name: "optional string flags can be omited",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", String, false)
			},
			args: []string{"-b", "bar"},
			expectedFlags: map[string]any{
				"b": "bar",
			},
		},
		{
			name: "missing required flags throws an error",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", String, true)
				p.AddShortFlag("c", String, false)
			},
			args:      []string{},
			expectErr: "required flag not provided: b",
		},
		{
			name: "can group string flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, true)
				p.AddShortFlag("b", String, true)
				p.AddShortFlag("c", String, true)
			},
			args: []string{"-cba", "foo", "bar", "baz"},
			expectedFlags: map[string]any{
				"a": "baz",
				"b": "bar",
				"c": "foo",
			},
		},
		{
			name: "can group string and boolean flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, true)
				p.AddShortFlag("b", Boolean, true)
				p.AddShortFlag("c", String, true)
			},
			args: []string{"-cba", "foo", "bar"},
			expectedFlags: map[string]any{
				"a": "bar",
				"b": true,
				"c": "foo",
			},
		},
		{
			name: "missing arguments to flags throws an error",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", String, false)
			},
			args:      []string{"-a", "foo", "-b"},
			expectErr: "not enough arguments for flags in group: -b",
		},
		{
			name: "basic boolean long flags",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", Boolean, true)
				p.AddLongFlag("bar", Boolean, true)
				p.AddLongFlag("baz", Boolean, true)
			},
			args: []string{"--foo", "--bar", "--baz"},
			expectedFlags: map[string]any{
				"foo": true,
				"bar": true,
				"baz": true,
			},
		},
		{
			name: "boolean long flags are optional",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", Boolean, true)
				p.AddLongFlag("bar", Boolean, true)
				p.AddLongFlag("baz", Boolean, true)
			},
			args: []string{"--bar"},
			expectedFlags: map[string]any{
				"foo": false,
				"bar": true,
				"baz": false,
			},
		},
		{
			name: "basic string long flags",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo-bar", String, true)
				p.AddLongFlag("baz", String, true)
			},
			args: []string{"--foo-bar", "boo", "--baz", "pie"},
			expectedFlags: map[string]any{
				"foo-bar": "boo",
				"baz":     "pie",
			},
		},
		{
			name: "string long flags require an argument",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", String, true)
				p.AddLongFlag("bar", String, true)
			},
			args:      []string{"--foo", "boo", "--bar"},
			expectErr: "missing value for flag: bar",
		},
		{
			name: "mixed short and long flags",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", String, true)
				p.AddLongFlag("baz", Boolean, true)
				p.AddShortFlag("a", Boolean, true)
				p.AddShortFlag("c", String, true)
			},
			args: []string{"-a", "--foo", "bar", "-c", "zik", "--baz"},
			expectedFlags: map[string]any{
				"foo": "bar",
				"baz": true,
				"a":   true,
				"c":   "zik",
			},
		},
		{
			name: "arguments are returned back",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", String, false)
			},
			args:          []string{"foo", "bar", "baz"},
			expectedFlags: map[string]any{},
			expectedArgs:  []string{"foo", "bar", "baz"},
		},
		{
			name: "mixed flags and arguments",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", String, true)
				p.AddLongFlag("baz", Boolean, true)
				p.AddShortFlag("a", Boolean, true)
				p.AddShortFlag("c", String, true)
			},
			args: []string{"-a", "abc", "--foo", "bar", "xyz", "-c", "zik", "vbn", "--baz", "pop"},
			expectedFlags: map[string]any{
				"foo": "bar",
				"baz": true,
				"a":   true,
				"c":   "zik",
			}, expectedArgs: []string{"abc", "xyz", "vbn", "pop"},
		},
		{
			name: "missing required flags throws an error",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("abc", String, true)
			},
			expectErr: "required flag not provided: abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			tt.flagSetup(p)

			result, err := p.Parse(tt.args)

			if tt.expectErr != "" || err != nil {
				if tt.expectErr != "" && err == nil {
					t.Fatalf("expected error: %q, got nil", tt.expectErr)
				}

				if err != nil && tt.expectErr != err.Error() {
					t.Fatalf("expected error: %q, got: %q", tt.expectErr, err)
				}
				return
			}

			if !reflect.DeepEqual(result.Flags, tt.expectedFlags) {
				t.Fatalf("Parse() flags = %v, want %v", result.Flags, tt.expectedFlags)
			}

			if !reflect.DeepEqual(result.Args, tt.expectedArgs) {
				t.Fatalf("Parse() args = %v, want %v", result.Args, tt.expectedArgs)
			}
		})
	}
}

func TestParser_AddFlags(t *testing.T) {
	p := NewParser()

	// Valid short flag
	err := p.AddShortFlag("a", Boolean, false)
	if err != nil {
		t.Errorf("AddShortFlag() unexpected error = %v", err)
	}

	// Invalid short flag (more than one character)
	err = p.AddShortFlag("abc", Boolean, false)
	if err == nil {
		t.Error("AddShortFlag() expected error for multi-character flag")
	}

	// Valid long flag
	err = p.AddLongFlag("verbose", Boolean, false)
	if err != nil {
		t.Errorf("AddLongFlag() unexpected error = %v", err)
	}

	// Invalid long flag (only one character)
	err = p.AddLongFlag("v", Boolean, false)
	if err == nil {
		t.Error("AddLongFlag() expected error for single-character flag")
	}

}

// {
// 	name: "Basic short boolean flags",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", Boolean, false)
// 		p.AddShortFlag("b", Boolean, false)
// 		p.AddShortFlag("c", Boolean, false)
// 	},
// 	args: []string{"-a", "-c"},
// 	expectedFlags: map[string]any{
// 		"a": true,
// 		"b": false,
// 		"c": true,
// 	},
// },
// {
// 	name: "Short boolean flags with grouping",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", Boolean, false)
// 		p.AddShortFlag("b", Boolean, false)
// 		p.AddShortFlag("c", Boolean, false)
// 	},
// 	args: []string{"-abc"},
// 	expectedFlags: map[string]any{
// 		"a": true,
// 		"b": true,
// 		"c": true,
// 	},
// },
// {
// 	name: "Short string flags",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", String, false)
// 		p.AddShortFlag("b", Boolean, false)
// 		p.AddShortFlag("c", String, false)
// 	},
// 	args: []string{"-a", "foo", "-b", "-c", "bar"},
// 	expectedFlags: map[string]any{
// 		"a": "foo",
// 		"b": true,
// 		"c": "bar",
// 	},
// },
// {
// 	name: "Mixed short flags with grouping",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", String, false)
// 		p.AddShortFlag("b", Boolean, false)
// 		p.AddShortFlag("c", String, false)
// 	},
// 	args: []string{"-abc", "foo", "bar"},
// 	expectedFlags: map[string]any{
// 		"a": "foo",
// 		"b": true,
// 		"c": "bar",
// 	},
// },
// {
// 	name: "Basic long boolean flags",
// 	flagSetup: func(p *Parser) {
// 		p.AddLongFlag("foo", Boolean, false)
// 		p.AddLongFlag("bar", Boolean, false)
// 		p.AddLongFlag("baz", Boolean, false)
// 	},
// 	args: []string{"--foo", "--baz"},
// 	expectedFlags: map[string]any{
// 		"foo": true,
// 		"bar": false,
// 		"baz": true,
// 	},
// },
// {
// 	name: "Long string flags",
// 	flagSetup: func(p *Parser) {
// 		p.AddLongFlag("name", String, false)
// 		p.AddLongFlag("verbose", Boolean, false)
// 		p.AddLongFlag("output", String, false)
// 	},
// 	args: []string{"--name", "John", "--verbose", "--output", "file.txt"},
// 	expectedFlags: map[string]any{
// 		"name":    "John",
// 		"verbose": true,
// 		"output":  "file.txt",
// 	},
// },
// {
// 	name: "Mixed short and long flags",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", String, false)
// 		p.AddShortFlag("b", Boolean, false)
// 		p.AddLongFlag("verbose", Boolean, false)
// 		p.AddLongFlag("name", String, false)
// 	},
// 	args: []string{"-a", "foo", "--verbose", "-b", "--name", "John"},
// 	expectedFlags: map[string]any{
// 		"a":       "foo",
// 		"b":       true,
// 		"verbose": true,
// 		"name":    "John",
// 	},
// },
// {
// 	name: "Long flags with dashes",
// 	flagSetup: func(p *Parser) {
// 		p.AddLongFlag("flag-with-dashes", Boolean, false)
// 		p.AddLongFlag("another-flag", String, false)
// 	},
// 	args: []string{"--flag-with-dashes", "--another-flag", "value"},
// 	expectedFlags: map[string]any{
// 		"flag-with-dashes": true,
// 		"another-flag":     "value",
// 	},
// },
// {
// 	name: "Required flags mix",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", String, true)
// 		p.AddLongFlag("required", Boolean, true)
// 		p.AddLongFlag("optional", String, false)
// 	},
// 	args: []string{"-a", "foo", "--required"},
// 	expectedFlags: map[string]any{
// 		"a":        "foo",
// 		"required": true,
// 	},
// },
// {
// 	name: "Missing required long flag",
// 	flagSetup: func(p *Parser) {
// 		p.AddLongFlag("required", String, true)
// 		p.AddShortFlag("a", Boolean, false)
// 	},
// 	args:          []string{"-a"},
// 	expectedFlags: nil,
// 	expectedArgs:  nil,
// 	expectErr:     true,
// },
// {
// 	name: "Missing value for long string flag",
// 	flagSetup: func(p *Parser) {
// 		p.AddLongFlag("name", String, false)
// 	},
// 	args:          []string{"--name"},
// 	expectedFlags: nil,
// 	expectedArgs:  nil,
// 	expectErr:     true,
// },
// {
// 	name: "Unknown long flag",
// 	flagSetup: func(p *Parser) {
// 		p.AddLongFlag("known", Boolean, false)
// 	},
// 	args:          []string{"--unknown"},
// 	expectedFlags: nil,
// 	expectedArgs:  nil,
// 	expectErr:     true,
// },
// {
// 	name: "With non-flag arguments",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", Boolean, false)
// 		p.AddLongFlag("name", String, false)
// 	},
// 	args: []string{"file1.txt", "-a", "--name", "value", "file2.txt", "file3.txt"},
// 	expectedFlags: map[string]any{
// 		"a":    true,
// 		"name": "value",
// 	},
// 	expectedArgs: []string{"file1.txt", "file2.txt", "file3.txt"},
// },
// {
// 	name: "All non-flag arguments",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", Boolean, false)
// 		p.AddLongFlag("verbose", Boolean, false)
// 	},
// 	args: []string{"file1.txt", "file2.txt", "file3.txt"},
// 	expectedFlags: map[string]any{
// 		"a":       false,
// 		"verbose": false,
// 	},
// 	expectedArgs: []string{"file1.txt", "file2.txt", "file3.txt"},
// },
// {
// 	name: "Flags and non-flags mixed",
// 	flagSetup: func(p *Parser) {
// 		p.AddShortFlag("a", String, false)
// 		p.AddLongFlag("verbose", Boolean, false)
// 	},
// 	args: []string{"-a", "foo", "command", "--verbose", "arg1", "arg2"},
// 	expectedFlags: map[string]any{
// 		"a":       "foo",
// 		"verbose": true,
// 	},
// 	expectedArgs: []string{"command", "arg1", "arg2"},
// },

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
		expectedFlags map[string]interface{}
		expectedArgs  []string
		expectErr     bool
	}{
		{
			name: "Basic short boolean flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, false)
				p.AddShortFlag("b", Boolean, false)
				p.AddShortFlag("c", Boolean, false)
			},
			args: []string{"-a", "-c"},
			expectedFlags: map[string]interface{}{
				"a": true,
				"b": false,
				"c": true,
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Short boolean flags with grouping",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, false)
				p.AddShortFlag("b", Boolean, false)
				p.AddShortFlag("c", Boolean, false)
			},
			args: []string{"-abc"},
			expectedFlags: map[string]interface{}{
				"a": true,
				"b": true,
				"c": true,
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Short string flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", Boolean, false)
				p.AddShortFlag("c", String, false)
			},
			args: []string{"-a", "foo", "-b", "-c", "bar"},
			expectedFlags: map[string]interface{}{
				"a": "foo",
				"b": true,
				"c": "bar",
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Mixed short flags with grouping",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", Boolean, false)
				p.AddShortFlag("c", String, false)
			},
			args: []string{"-abc", "foo", "bar"},
			expectedFlags: map[string]interface{}{
				"a": "foo",
				"b": true,
				"c": "bar",
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Basic long boolean flags",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("foo", Boolean, false)
				p.AddLongFlag("bar", Boolean, false)
				p.AddLongFlag("baz", Boolean, false)
			},
			args: []string{"--foo", "--baz"},
			expectedFlags: map[string]interface{}{
				"foo": true,
				"bar": false,
				"baz": true,
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Long string flags",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("name", String, false)
				p.AddLongFlag("verbose", Boolean, false)
				p.AddLongFlag("output", String, false)
			},
			args: []string{"--name", "John", "--verbose", "--output", "file.txt"},
			expectedFlags: map[string]interface{}{
				"name":    "John",
				"verbose": true,
				"output":  "file.txt",
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Mixed short and long flags",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddShortFlag("b", Boolean, false)
				p.AddLongFlag("verbose", Boolean, false)
				p.AddLongFlag("name", String, false)
			},
			args: []string{"-a", "foo", "--verbose", "-b", "--name", "John"},
			expectedFlags: map[string]interface{}{
				"a":       "foo",
				"b":       true,
				"verbose": true,
				"name":    "John",
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Long flags with dashes",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("flag-with-dashes", Boolean, false)
				p.AddLongFlag("another-flag", String, false)
			},
			args: []string{"--flag-with-dashes", "--another-flag", "value"},
			expectedFlags: map[string]interface{}{
				"flag-with-dashes": true,
				"another-flag":     "value",
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Required flags mix",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, true)
				p.AddLongFlag("required", Boolean, true)
				p.AddLongFlag("optional", String, false)
			},
			args: []string{"-a", "foo", "--required"},
			expectedFlags: map[string]interface{}{
				"a":        "foo",
				"required": true,
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Missing required long flag",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("required", String, true)
				p.AddShortFlag("a", Boolean, false)
			},
			args:          []string{"-a"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "Missing value for long string flag",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("name", String, false)
			},
			args:          []string{"--name"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "Unknown long flag",
			flagSetup: func(p *Parser) {
				p.AddLongFlag("known", Boolean, false)
			},
			args:          []string{"--unknown"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "With non-flag arguments",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, false)
				p.AddLongFlag("name", String, false)
			},
			args: []string{"file1.txt", "-a", "--name", "value", "file2.txt", "file3.txt"},
			expectedFlags: map[string]interface{}{
				"a":    true,
				"name": "value",
			},
			expectedArgs: []string{"file1.txt", "file2.txt", "file3.txt"},
			expectErr:    false,
		},
		{
			name: "All non-flag arguments",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", Boolean, false)
				p.AddLongFlag("verbose", Boolean, false)
			},
			args: []string{"file1.txt", "file2.txt", "file3.txt"},
			expectedFlags: map[string]interface{}{
				"a":       false,
				"verbose": false,
			},
			expectedArgs: []string{"file1.txt", "file2.txt", "file3.txt"},
			expectErr:    false,
		},
		{
			name: "Flags and non-flags mixed",
			flagSetup: func(p *Parser) {
				p.AddShortFlag("a", String, false)
				p.AddLongFlag("verbose", Boolean, false)
			},
			args: []string{"-a", "foo", "command", "--verbose", "arg1", "arg2"},
			expectedFlags: map[string]interface{}{
				"a":       "foo",
				"verbose": true,
			},
			expectedArgs: []string{"command", "arg1", "arg2"},
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			tt.flagSetup(p)

			result, err := p.Parse(tt.args)

			if (err != nil) != tt.expectErr {
				t.Errorf("Parse() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			if !tt.expectErr {
				if !reflect.DeepEqual(result.Flags, tt.expectedFlags) {
					t.Errorf("Parse() flags = %v, want %v", result.Flags, tt.expectedFlags)
				}

				if !reflect.DeepEqual(result.Args, tt.expectedArgs) {
					t.Errorf("Parse() args = %v, want %v", result.Args, tt.expectedArgs)
				}
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

	// Legacy AddFlag method should work as AddShortFlag
	err = p.AddFlag("z", Boolean, false)
	if err != nil {
		t.Errorf("AddFlag() (legacy) unexpected error = %v", err)
	}
}

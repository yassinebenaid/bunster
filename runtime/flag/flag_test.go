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
			name: "Basic boolean flags",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", Boolean, false)
				p.AddFlag("b", Boolean, false)
				p.AddFlag("c", Boolean, false)
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
			name: "Boolean flags with grouping",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", Boolean, false)
				p.AddFlag("b", Boolean, false)
				p.AddFlag("c", Boolean, false)
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
			name: "String flags",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, false)
				p.AddFlag("b", Boolean, false)
				p.AddFlag("c", String, false)
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
			name: "Mixed flags with grouping",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, false)
				p.AddFlag("b", Boolean, false)
				p.AddFlag("c", String, false)
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
			name: "String flag in middle of group",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", Boolean, false)
				p.AddFlag("b", String, false)
				p.AddFlag("c", Boolean, false)
			},
			args: []string{"-abc", "value"},
			expectedFlags: map[string]interface{}{
				"a": true,
				"b": "value",
				"c": true,
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Complex grouping with multiple string flags",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, false)
				p.AddFlag("b", Boolean, false)
				p.AddFlag("c", String, false)
				p.AddFlag("d", Boolean, false)
				p.AddFlag("e", String, false)
			},
			args: []string{"-abcde", "avalueC", "cvalueC", "evalueC"},
			expectedFlags: map[string]interface{}{
				"a": "avalueC",
				"b": true,
				"c": "cvalueC",
				"d": true,
				"e": "evalueC",
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Required flags",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, true)
				p.AddFlag("b", Boolean, true)
				p.AddFlag("c", String, false)
			},
			args: []string{"-a", "foo", "-b"},
			expectedFlags: map[string]interface{}{
				"a": "foo",
				"b": true,
			},
			expectedArgs: []string{},
			expectErr:    false,
		},
		{
			name: "Missing required flag",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, true)
				p.AddFlag("b", Boolean, false)
			},
			args:          []string{"-b"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "Missing value for string flag",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, false)
			},
			args:          []string{"-a"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "Not enough values for grouped string flags",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, false)
				p.AddFlag("b", Boolean, false)
				p.AddFlag("c", String, false)
			},
			args:          []string{"-abc", "foo"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "Unknown flag",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", Boolean, false)
			},
			args:          []string{"-x"},
			expectedFlags: nil,
			expectedArgs:  nil,
			expectErr:     true,
		},
		{
			name: "With non-flag arguments",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", Boolean, false)
				p.AddFlag("b", String, false)
			},
			args: []string{"file1.txt", "-a", "-b", "value", "file2.txt", "file3.txt"},
			expectedFlags: map[string]interface{}{
				"a": true,
				"b": "value",
			},
			expectedArgs: []string{"file1.txt", "file2.txt", "file3.txt"},
			expectErr:    false,
		},
		{
			name: "All non-flag arguments",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", Boolean, false)
				p.AddFlag("b", Boolean, false)
			},
			args: []string{"file1.txt", "file2.txt", "file3.txt"},
			expectedFlags: map[string]interface{}{
				"a": false,
				"b": false,
			},
			expectedArgs: []string{"file1.txt", "file2.txt", "file3.txt"},
			expectErr:    false,
		},
		{
			name: "Flags and non-flags mixed",
			flagSetup: func(p *Parser) {
				p.AddFlag("a", String, false)
				p.AddFlag("b", Boolean, false)
			},
			args: []string{"-a", "foo", "command", "-b", "arg1", "arg2"},
			expectedFlags: map[string]interface{}{
				"a": "foo",
				"b": true,
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

func TestParser_AddFlag(t *testing.T) {
	p := NewParser()

	// Valid flag
	err := p.AddFlag("a", Boolean, false)
	if err != nil {
		t.Errorf("AddFlag() unexpected error = %v", err)
	}

	// Invalid flag (more than one character)
	err = p.AddFlag("abc", Boolean, false)
	if err == nil {
		t.Error("AddFlag() expected error for multi-character flag")
	}
}

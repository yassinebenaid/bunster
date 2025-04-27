package runtime_test

import (
	"reflect"
	"testing"

	"github.com/yassinebenaid/bunster/runtime"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		flagSetup     func(*runtime.FlagParser)
		args          []string
		expectedFlags map[string]any
		expectedArgs  []string
		expectErr     string
	}{
		{
			name: "basic short boolean flags",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.BooleanFlag, false)
				p.AddShortFlag("b", runtime.BooleanFlag, false)
				p.AddShortFlag("c", runtime.BooleanFlag, false)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.BooleanFlag, true)
				p.AddShortFlag("b", runtime.BooleanFlag, true)
				p.AddShortFlag("c", runtime.BooleanFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.BooleanFlag, true)
				p.AddShortFlag("b", runtime.BooleanFlag, true)
				p.AddShortFlag("c", runtime.BooleanFlag, true)
				p.AddShortFlag("d", runtime.BooleanFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, false)
				p.AddShortFlag("b", runtime.StringFlag, false)
			},
			args: []string{"-a", "foo", "-b", "bar"},
			expectedFlags: map[string]any{
				"a": "foo",
				"b": "bar",
			},
		},
		{
			name: "optional string flags can be omited",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, false)
				p.AddShortFlag("b", runtime.StringFlag, false)
			},
			args: []string{"-b", "bar"},
			expectedFlags: map[string]any{
				"b": "bar",
			},
		},
		{
			name: "missing required flags throws an error",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, false)
				p.AddShortFlag("b", runtime.StringFlag, true)
				p.AddShortFlag("c", runtime.StringFlag, false)
			},
			args:      []string{},
			expectErr: "required flag not provided: b",
		},
		{
			name: "can group string flags",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, true)
				p.AddShortFlag("b", runtime.StringFlag, true)
				p.AddShortFlag("c", runtime.StringFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, true)
				p.AddShortFlag("b", runtime.BooleanFlag, true)
				p.AddShortFlag("c", runtime.StringFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, false)
				p.AddShortFlag("b", runtime.StringFlag, false)
			},
			args:      []string{"-b", "-a", "foo"},
			expectErr: "missing value for flag: b",
		},
		{
			name: "missing arguments to flags that show last throws an error",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, false)
				p.AddShortFlag("b", runtime.StringFlag, false)
				p.AddShortFlag("c", runtime.StringFlag, false)
			},
			args:      []string{"-abc"},
			expectErr: "missing value for flag: a",
		},
		{
			name: "basic boolean long flags",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.BooleanFlag, true)
				p.AddLongFlag("bar", runtime.BooleanFlag, true)
				p.AddLongFlag("baz", runtime.BooleanFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.BooleanFlag, true)
				p.AddLongFlag("bar", runtime.BooleanFlag, true)
				p.AddLongFlag("baz", runtime.BooleanFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo-bar", runtime.StringFlag, true)
				p.AddLongFlag("baz", runtime.StringFlag, true)
				p.AddLongFlag("xyz", runtime.StringFlag, true)
			},
			args: []string{"--foo-bar", "boo", "--baz", "pie", "--xyz=pop", "booyah"},
			expectedFlags: map[string]any{
				"foo-bar": "boo",
				"baz":     "pie",
				"xyz":     "pop",
			},
			expectedArgs: []string{"booyah"},
		},
		{
			name: "string long flags require an argument",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.StringFlag, true)
				p.AddLongFlag("bar", runtime.StringFlag, true)
			},
			args:      []string{"--bar", "--foo", "boo"},
			expectErr: "missing value for flag: bar",
		},
		{
			name: "inline string long flags require an argument",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.StringFlag, true)
				p.AddLongFlag("bar", runtime.StringFlag, true)
			},
			args:      []string{"--bar=", "--foo", "boo"},
			expectErr: "missing value for flag: bar",
		},
		{
			name: "string long flags require an argument when appeare at end",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("bar", runtime.StringFlag, true)
			},
			args:      []string{"--bar"},
			expectErr: "missing value for flag: bar",
		},
		{
			name: "mixed short and long flags",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.StringFlag, true)
				p.AddLongFlag("baz", runtime.BooleanFlag, true)
				p.AddShortFlag("a", runtime.BooleanFlag, true)
				p.AddShortFlag("c", runtime.StringFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.StringFlag, false)
			},
			args:          []string{"foo", "bar", "baz"},
			expectedFlags: map[string]any{},
			expectedArgs:  []string{"foo", "bar", "baz"},
		},
		{
			name: "mixed flags and arguments",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.StringFlag, true)
				p.AddLongFlag("baz", runtime.BooleanFlag, true)
				p.AddShortFlag("a", runtime.BooleanFlag, true)
				p.AddShortFlag("c", runtime.StringFlag, true)
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
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("abc", runtime.StringFlag, true)
			},
			expectErr: "required flag not provided: abc",
		},
		{
			name:      "an error occurs when pasing unknown short flags",
			flagSetup: func(p *runtime.FlagParser) {},
			args:      []string{"-x"},
			expectErr: "unknown short flag: x",
		},
		{
			name:      "an error occurs when pasing unknown long flags",
			flagSetup: func(p *runtime.FlagParser) {},
			args:      []string{"--foo"},
			expectErr: "unknown long flag: foo",
		},
		{
			name:      "an error occurs when passing unknown long flags with value",
			flagSetup: func(p *runtime.FlagParser) {},
			args:      []string{"--foo=bar"},
			expectErr: "unknown long flag: foo",
		},
		{
			name: "an error occurs when passing a value to boolean long flags",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.BooleanFlag, false)
			},
			args:      []string{"--foo=bar"},
			expectErr: "passing value to a flag that doesn't expect it: foo",
		},
		{
			name:      "an error occurs when passing a single dash",
			flagSetup: func(p *runtime.FlagParser) {},
			args:      []string{"c", "-", "a"},
			expectErr: "invalid short flag format: -",
		},
		{
			name: "double dash indicate end of flags",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("foo", runtime.StringFlag, false)
				p.AddLongFlag("bar", runtime.StringFlag, false)
			},
			args:          []string{"foo", "--", "--foo", "--bar", "--baz"},
			expectedFlags: map[string]any{},
			expectedArgs:  []string{"foo", "--foo", "--bar", "--baz"},
		},
		{
			name: "an error occurs when the same short flag appears too many times",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.BooleanFlag, false)
			},
			args:      []string{"-a", "-a"},
			expectErr: "flag supplied too many times: a",
		},
		{
			name: "an error occurs when the same short string flag appears too many times",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddShortFlag("a", runtime.StringFlag, false)
			},
			args:      []string{"-a", "vv", "-a", "xx"},
			expectErr: "flag supplied too many times: a",
		},
		{
			name: "an error occurs when the same long flag appears too many times",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("abc", runtime.BooleanFlag, false)
			},
			args:      []string{"--abc", "--abc"},
			expectErr: "flag supplied too many times: abc",
		},
		{
			name: "an error occurs when the same long string flag appears too many times",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("abc", runtime.StringFlag, false)
			},
			args:      []string{"--abc", "vv", "--abc", "xx"},
			expectErr: "flag supplied too many times: abc",
		},
		{
			name: "an error occurs when the same long inline string flag appears too many times",
			flagSetup: func(p *runtime.FlagParser) {
				p.AddLongFlag("abc", runtime.StringFlag, false)
			},
			args:      []string{"--abc=vv", "--abc=xx"},
			expectErr: "flag supplied too many times: abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runtime.NewFlagParser()
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
	// Valid short flag
	_, err := runtime.NewFlagParser().AddShortFlag("a", runtime.BooleanFlag, false).Parse(nil)
	if err != nil {
		t.Errorf("AddShortFlag() unexpected error = %v", err)
	}

	// Invalid short flag (more than one character)
	_, err = runtime.NewFlagParser().AddShortFlag("abc", runtime.BooleanFlag, false).Parse(nil)
	if err == nil {
		t.Error("AddShortFlag() expected error for multi-character flag")
	}

	// Valid long flag
	_, err = runtime.NewFlagParser().AddLongFlag("verbose", runtime.BooleanFlag, false).Parse(nil)
	if err != nil {
		t.Errorf("AddLongFlag() unexpected error = %v", err)
	}

	// Invalid long flag (only one character)
	_, err = runtime.NewFlagParser().AddLongFlag("v", runtime.BooleanFlag, false).Parse(nil)
	if err == nil {
		t.Error("AddLongFlag() expected error for single-character flag")
	}

}

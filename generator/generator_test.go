package generator_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yassinebenaid/bunster/pkg/tst"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestGenerator(t *testing.T) {
	testFiles, err := filepath.Glob("./tests/*.tst")
	if err != nil {
		t.Fatalf("Failed to `Glob` test files, %v", err)
	}

	for _, testFile := range testFiles {
		file, err := os.Open(testFile)
		if err != nil {
			t.Fatalf("Failed to open test file %s, %v", testFile, err)
		}

		cases, err := tst.Parse(file)
		if err != nil {
			t.Fatalf("%v, (%s)", err, testFile)
		}

		_ = cases
	}
}

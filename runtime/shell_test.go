package runtime_test

import (
	"testing"

	"github.com/yassinebenaid/bunster/runtime"
)

func TestShell_Run(t *testing.T) {
	shell := runtime.Shell{
		Stdin:  &testStream{},
		Stdout: &testStream{},
		Stderr: &testStream{},
		Main: func(s *runtime.Shell, sm *runtime.StreamManager) {

			s.ExitCode = 123
		},
	}

	if code := shell.Run(); code != 123 {
		t.Fatalf("shell.Run() must return %d, returns %d", 123, code)
	}
}

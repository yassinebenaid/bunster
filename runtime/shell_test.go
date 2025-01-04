package runtime

import (
	"testing"
)

func TestShell_Run(t *testing.T) {
	shell := Shell{
		Stdin:  &testStream{},
		Stdout: &testStream{},
		Stderr: &testStream{},
		Main: func(s *Shell, sm *StreamManager) {

			s.ExitCode = 123
		},
	}

	if code := shell.Run(); code != 123 {
		t.Fatalf("shell.Run() must return %d, returns %d", 123, code)
	}
}

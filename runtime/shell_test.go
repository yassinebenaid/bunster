package runtime_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yassinebenaid/bunster/runtime"
)

type testStream struct {
	closed bool
	buf    bytes.Buffer
}

func (ts *testStream) Read(b []byte) (int, error) {
	return ts.Read(b)
}

func (ts *testStream) Write(b []byte) (int, error) {
	return ts.Write(b)
}

func (ts *testStream) Close() error {
	if ts.closed {
		return fmt.Errorf("closing closed stream")
	}
	return nil
}

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

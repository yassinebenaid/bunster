package runtime_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yassinebenaid/bunster/runtime"
)

type testStream struct {
	id     string // we just use this id for assertions
	closed bool
	buf    bytes.Buffer
}

func (ts *testStream) Read(b []byte) (int, error) {
	return ts.buf.Read(b)
}

func (ts *testStream) Write(b []byte) (int, error) {
	return ts.buf.Write(b)
}

func (ts *testStream) Close() error {
	if ts.closed {
		return fmt.Errorf("closing closed stream")
	}
	return nil
}

func TestStreamManager_Add_And_Get(t *testing.T) {
	sm := runtime.NewStreamManager()
	originalStream := &testStream{id: "foobar"}

	sm.Add("123", originalStream, false)

	returnedStream, err := sm.Get("123")
	if err != nil {
		t.Fatalf("StreamManager.Get() returns error where it shouldn't, error: %s", err)
	}

	concreteReturnedStream, ok := returnedStream.(*testStream)
	if !ok {
		t.Fatalf("StreamManager.Get() returns the wrong stream type, we added stream of type *testStream, it returns %T", returnedStream)
	}

	if concreteReturnedStream.id != originalStream.id {
		t.Fatalf("StreamManager.Get() returns the wrong stream")
	}
}

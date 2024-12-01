package runtime

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Stream interface {
	io.Reader
	io.Writer
	io.Closer
}

func OpenStream(name string) (Stream, error) {
	return os.OpenFile(name, os.O_RDWR, 0644)
}

func OpenOrCreateStream(name string) (Stream, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
}

func OpenOrCreateStreamForAppending(name string) (Stream, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

type stringStream struct {
	buf *bytes.Buffer
}

func (s *stringStream) Close() error {
	if s.buf == nil {
		return fmt.Errorf("cannot close closed stream")
	}
	s.buf = nil
	return nil
}

func (s *stringStream) Read(p []byte) (n int, err error) {
	if s.buf == nil {
		return 0, fmt.Errorf("cannot read from closed stream")
	}
	return s.buf.Read(p)
}

func (s *stringStream) Write(p []byte) (n int, err error) {
	if s.buf == nil {
		return 0, fmt.Errorf("cannot write to closed stream")
	}
	return s.buf.Write(p)
}

func NewStringStream(s string) *stringStream {
	return &stringStream{buf: bytes.NewBufferString(s)}
}

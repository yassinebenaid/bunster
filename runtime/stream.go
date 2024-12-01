package runtime

import (
	"bytes"
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
	buf bytes.Buffer
}

func (s *stringStream) Close() error {
	s.buf = bytes.Buffer{}
	return nil
}

func NewStringStream(s string) stringStream {
	return stringStream{buf: *bytes.NewBufferString(s)}
}

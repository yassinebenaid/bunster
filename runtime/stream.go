package runtime

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	STREAM_FLAG_READ   = os.O_RDONLY
	STREAM_FLAG_WRITE  = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	STREAM_FLAG_RW     = os.O_RDWR | os.O_CREATE
	STREAM_FLAG_APPEND = os.O_WRONLY | os.O_APPEND | os.O_CREATE
)

type Stream interface {
	io.Reader
	io.Writer
	io.Closer
}

type Buffer struct {
	buf      *bytes.Buffer
	readonly bool
}

func (s *Buffer) Close() error {
	if s.buf == nil {
		return fmt.Errorf("cannot close closed stream")
	}
	s.buf = nil
	return nil
}

func (s *Buffer) Read(p []byte) (n int, err error) {
	if s.buf == nil {
		return 0, fmt.Errorf("bad file descriptor, cannot read from closed stream")
	}
	return s.buf.Read(p)
}

func (s *Buffer) Write(p []byte) (n int, err error) {
	if s.buf == nil {
		return 0, fmt.Errorf("bad file descriptor, cannot write to closed stream")
	}
	if s.readonly {
		return 0, fmt.Errorf("bad file descriptor, cannot write to read-only stream")
	}
	return s.buf.Write(p)
}

func (s *Buffer) String(trim_leading_newline bool) string {
	v := s.buf.String()
	if trim_leading_newline {
		return strings.TrimRight(v, "\n")
	}
	return v
}

func NewBuffer(s string, readonly bool) *Buffer {
	return &Buffer{
		buf:      bytes.NewBufferString(s),
		readonly: readonly,
	}
}

type proxyStream struct {
	original Stream
	closed   bool
}

func (s *proxyStream) Close() error {
	if s.closed {
		return fmt.Errorf("cannot close closed stream")
	}
	s.closed = true
	return nil
}

func (s *proxyStream) Read(p []byte) (n int, err error) { return 0, nil }

func (s *proxyStream) Write(p []byte) (n int, err error) { return 0, nil }
func (s *proxyStream) getOriginal() (Stream, error) {
	if s.closed {
		return nil, fmt.Errorf("file descriptor is closed")
	}

	if o, ok := s.original.(*proxyStream); ok {
		return o.getOriginal()
	}

	return s.original, nil
}

type StreamManager struct {
	openStreams []Stream
	mappings    map[string]*proxyStream
}

func (sm *StreamManager) OpenStream(name string, flag int) (Stream, error) {
	switch name {
	case "/dev/stdin":
		proxy, ok := sm.mappings["0"]
		if !ok {
			return nil, fmt.Errorf("file descriptor %q is not open", "0")
		}
		return &proxyStream{original: proxy.original}, nil
	case "/dev/stdout":
		proxy, ok := sm.mappings["1"]
		if !ok {
			return nil, fmt.Errorf("file descriptor %q is not open", "1")
		}
		return &proxyStream{original: proxy.original}, nil
	case "/dev/stderr":
		proxy, ok := sm.mappings["2"]
		if !ok {
			return nil, fmt.Errorf("file descriptor %q is not open", "2")
		}
		return &proxyStream{original: proxy.original}, nil
	default:
		return os.OpenFile(name, flag, 0644)
	}
}

func (sm *StreamManager) Add(fd string, stream Stream, _ bool) {
	if proxy, ok := stream.(*proxyStream); ok {
		sm.mappings[fd] = proxy
		return
	}

	sm.openStreams = append(sm.openStreams, stream)
	proxy := &proxyStream{original: stream}
	sm.mappings[fd] = proxy
}

func (sm *StreamManager) Get(fd string) (Stream, error) {
	proxy, ok := sm.mappings[fd]
	if !ok {
		return nil, fmt.Errorf("file descriptor %q is not open", fd)
	}

	if stream, err := proxy.getOriginal(); err != nil {
		return nil, fmt.Errorf("bad file descriptor %q, %w", fd, err)
	} else {
		return stream, nil
	}
}

func (sm *StreamManager) Duplicate(newfd, oldfd string) error {
	if proxy, ok := sm.mappings[oldfd]; !ok {
		return fmt.Errorf("trying to duplicate bad file descriptor: %s", oldfd)
	} else if proxy.closed {
		return fmt.Errorf("trying to duplicate closed file descriptor: %s", oldfd)
	} else {
		sm.mappings[newfd] = &proxyStream{
			original: proxy.original,
		}
		return nil
	}
}

func (sm *StreamManager) Close(fd string) error {
	if stream, ok := sm.mappings[fd]; !ok {
		return fmt.Errorf("trying to close bad file descriptor: %s", fd)
	} else {
		return stream.Close()
	}
}

func (sm *StreamManager) Destroy() {
	for _, stream := range sm.openStreams {
		stream.Close()
	}
	for _, stream := range sm.mappings {
		stream.Close()
	}
}

func (sm *StreamManager) Clone() *StreamManager {
	clone := &StreamManager{
		mappings: make(map[string]*proxyStream),
	}

	for fd, stream := range sm.mappings {
		clone.mappings[fd] = &proxyStream{
			original: stream,
		}
	}
	return clone
}

func NewPipe() (Stream, Stream, error) {
	return os.Pipe()
}

func NewBufferedStream(s string) (Stream, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create buffered stream, %w", err)
	}

	go func() {
		_, _ = w.Write([]byte(s))
		_ = w.Close()
	}()

	return r, nil
}

package runtime

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"syscall"
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
		return 0, fmt.Errorf("bad file descriptor, cannot read from closed stream")
	}
	return s.buf.Read(p)
}

func (s *stringStream) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("bad file descriptor, cannot write to read only stream")
}

func NewStringStream(s string) Stream {
	return &stringStream{buf: bytes.NewBufferString(s)}
}

type StreamManager struct {
	parent   *StreamManager
	mappings map[string]Stream
	cleaners []func() error
}

func (sm *StreamManager) OpenStream(name string, flag int) (Stream, error) {
	switch name {
	case "/dev/stdin":
		return sm.Get("0"), nil
	case "/dev/stdout":
		return sm.Get("1"), nil
	case "/dev/stderr":
		return sm.Get("2"), nil
	default:
		return os.OpenFile(name, flag, 0644)
	}
}

func (sm *StreamManager) Add(fd string, stream Stream) {
	// If this stream is already open, we need to close it. otherwise, Its handler will be lost and leak.
	// This is related to pipelines in particular. when instantiating a new pipeline, we add its ends to the FDT. but if
	// a redirection happened afterwards, it will cause the pipline handler to be lost and kept open.
	if sm.mappings[fd] != nil {
		sm.cleaners = append(sm.cleaners, sm.mappings[fd].Close)
	}

	sm.mappings[fd] = stream
}

func (sm *StreamManager) Get(fd string) Stream {
	stream, ok := sm.mappings[fd]
	if !ok && sm.parent != nil {
		return sm.parent.Get(fd)
	}
	return stream
}

func (sm *StreamManager) Duplicate(newfd, oldfd string) error {
	if stream := sm.Get(oldfd); stream == nil {
		return fmt.Errorf("trying to duplicate bad file descriptor: %s", oldfd)
	} else {
		// when trying to duplicate a file descriptor to it self (eg: 3>&3 ), we just return.
		if newfd == oldfd {
			return nil
		}

		// If the new fd is already open, we need to close it. otherwise, Its handler will be lost and leak. and remain open forever.
		// for example: "3<file.txt 3<&0", we don't explicitly close 3. Thus, it is going to remain open forever, unless we implicitly close it here.
		if sm.mappings[newfd] != nil {
			sm.cleaners = append(sm.cleaners, sm.mappings[newfd].Close)
		}

		switch stream := stream.(type) {
		case *stringStream:
			newbuf := &bytes.Buffer{}
			_, err := io.Copy(newbuf, stream)
			if err != nil {
				return fmt.Errorf("failed to duplicate file descriptor '%s', %w", oldfd, err)
			}
			sm.mappings[newfd] = &stringStream{buf: newbuf}
		case *os.File:
			dupFd, err := syscall.Dup(int(stream.Fd()))
			if err != nil {
				return fmt.Errorf("failed to duplicate file descriptor '%s', %w", oldfd, err)
			}
			sm.mappings[newfd] = os.NewFile(uintptr(dupFd), stream.Name())
		default:
			panic(fmt.Sprintf("failed to clone (%s), unhandled stream type: %T", oldfd, stream))
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
	for _, cleanup := range sm.cleaners {
		cleanup()
	}
	for _, stream := range sm.mappings {
		stream.Close()
	}
}

func (sm *StreamManager) Clone() *StreamManager {
	return &StreamManager{
		parent:   sm,
		mappings: make(map[string]Stream),
	}
}

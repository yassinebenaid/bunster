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
	Fd() uintptr
}

func OpenStream(name string, flag int) (Stream, error) {
	return os.OpenFile(name, flag, 0644)
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

func (*stringStream) Fd() uintptr { return 0 }

func NewStringStream(s string) Stream {
	return &stringStream{buf: bytes.NewBufferString(s)}
}

type FileDescriptorTable map[string]Stream

func (fdt FileDescriptorTable) Add(fd string, stream Stream) {
	fdt[fd] = stream
}

func (fdt FileDescriptorTable) Get(fd string) Stream {
	stream, ok := fdt[fd]
	if !ok {
		// I'm not sure if we need to handle this case because we only use this function
		// to read 0, 1 and 2 file descriptors. Which are always available.
		panic("FIXME: an error handler is needed here.")
	}
	return stream
}

func (fdt FileDescriptorTable) Duplicate(newfd, oldfd string) error {
	if stream, ok := fdt[oldfd]; !ok {
		return fmt.Errorf("trying to duplicate bad file descriptor: %s", oldfd)
	} else {
		switch stream := stream.(type) {
		case *stringStream:
			newbuf := &bytes.Buffer{}
			_, err := io.Copy(newbuf, stream)
			if err != nil {
				return fmt.Errorf("failed to duplicate file descriptor '%s', %w", oldfd, err)
			}
			fdt[newfd] = &stringStream{buf: newbuf}
		case *os.File:
			dupFd, err := syscall.Dup(int(stream.Fd()))
			if err != nil {
				return fmt.Errorf("failed to duplicate file descriptor '%s', %w", oldfd, err)
			}
			fdt[newfd] = os.NewFile(uintptr(dupFd), stream.Name())
		default:
			panic(fmt.Sprintf("failed to clone (%s), unhandled stream type: %T", oldfd, stream))
		}

		return nil
	}
}

func (fdt FileDescriptorTable) Close(fd string) error {
	if stream, ok := fdt[fd]; !ok {
		return fmt.Errorf("trying to close bad file descriptor: %s", fd)
	} else {
		return stream.Close()
	}
}

func (fdt FileDescriptorTable) Destroy() {
	for _, stream := range fdt {
		stream.Close()
	}
}

func (fdt FileDescriptorTable) clone() (FileDescriptorTable, error) {
	clone := make(FileDescriptorTable, len(fdt))

	for fd, stream := range fdt {
		switch stream := stream.(type) {
		case *stringStream:
			newbuf := &bytes.Buffer{}
			_, err := io.Copy(newbuf, stream)
			if err != nil {
				return nil, fmt.Errorf("failure when trying to inherit the FileDescriptorTable, failed to duplicate file descriptor '%s', %w", fd, err)
			}
			clone[fd] = &stringStream{buf: newbuf}
		case *os.File:
			dupFd, err := syscall.Dup(int(stream.Fd()))
			if err != nil {
				return nil, fmt.Errorf("failure when trying to inherit the FileDescriptorTable, failed to duplicate file descriptor '%s', %w", fd, err)
			}
			clone[fd] = os.NewFile(uintptr(dupFd), stream.Name())
		default:
			panic(fmt.Sprintf("failed to clone FDT, unhandled stream type: %T (%s)", stream, fd))
		}
	}

	return clone, nil
}

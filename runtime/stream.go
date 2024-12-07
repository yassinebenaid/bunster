package runtime

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
)

type Stream interface {
	io.Reader
	io.Writer
	io.Closer
	Fd() uintptr
}

func OpenReadableStream(name string) (Stream, error) {
	return os.OpenFile(name, os.O_RDONLY, 0)
}

func OpenWritableStream(name string) (Stream, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}

func OpenReadWritableStream(name string) (Stream, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
}

func OpenAppendableStream(name string) (Stream, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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

func NewStreamFromFD(fds string) Stream {
	fd, err := strconv.ParseUint(fds, 10, 10)
	if err != nil {
		return nil
	}
	return os.NewFile(uintptr(fd), fmt.Sprintf("fd%d", fd))
}

func DuplicateFD(old int, new int) error {
	return syscall.Dup2(old, int(new))
}

type FileDescriptorTable map[string]Stream

func (fdt FileDescriptorTable) Add(fd string, stream Stream) {
	fdt[fd] = stream
}

func (fdt FileDescriptorTable) Get(fd string) (Stream, error) {
	if stream, ok := fdt[fd]; ok {
		return stream, nil
	}

	return nil, fmt.Errorf("bad file descriptor: %s", fd)
}

func (fdt FileDescriptorTable) Duplicate(oldfd, newfd string) error {
	if stream, ok := fdt[newfd]; !ok {
		return fmt.Errorf("trying to duplicate bad file descriptor: %s", newfd)
	} else {
		fdt[oldfd] = stream
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

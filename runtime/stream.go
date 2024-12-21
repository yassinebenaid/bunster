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

func (*stringStream) Fd() uintptr { return 0 }

func NewStringStream(s string) Stream {
	return &stringStream{buf: bytes.NewBufferString(s)}
}

type FileDescriptorTable map[string]Stream

func (fdt FileDescriptorTable) OpenStream(name string, flag int) (Stream, error) {
	switch name {
	case "/dev/stdin":
		return fdt.Get("0"), nil
	case "/dev/stdout":
		return fdt.Get("1"), nil
	case "/dev/stderr":
		return fdt.Get("2"), nil
	default:
		return os.OpenFile(name, flag, 0644)
	}
}

func (fdt FileDescriptorTable) Add(fd string, stream Stream) {
	// If this stream is already open, we need to close it. otherwise, Its handler will be lost and leak.
	// This is related to pipelines in particular. when instantiating a new pipeline, we add its ends to the FDT. but if
	// a redirection happened afterwards, it will cause the pipline handler to be lost and kept open.
	if fdt[fd] != nil {
		_ = fdt[fd].Close() // error here is not important.
	}

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
		// when trying to duplicate a file descriptor to it self (eg: 3>&3 ), we just return.
		if newfd == oldfd {
			return nil
		}

		// If the new fd is already open, we need to close it. otherwise, Its handler will be lost and leak. and remain open forever.
		// for example: "3<file.txt 3<&0", we don't explicitly close 3. Thus, it is going to remain open forever, unless we implicitly close it here.
		if fdt[newfd] != nil {
			_ = fdt[newfd].Close() // error here is not important.
		}

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

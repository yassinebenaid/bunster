//go:build linux || freebsd || openbsd || netbsd

package runtime

import (
	"syscall"
	"unsafe"
)

func FileDescriptorIsTerminal(sm *StreamManager, fd string) bool {
	original, err := sm.Get(fd)
	if err != nil {
		return false
	}

	file, ok := original.(interface{ Fd() uintptr })
	if !ok {
		return false
	}

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		file.Fd(),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(&syscall.Termios{})),
	)
	return errno == 0
}

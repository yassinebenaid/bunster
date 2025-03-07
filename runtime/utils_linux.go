package runtime

import (
	"os"
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

func FileHasBeenModifiedSinceLastRead(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return stat.Mtim.Nsec > stat.Atim.Nsec
}

package runtime

import (
	"os"
	"strconv"
	"syscall"
)

func NumberCompare(x, op, y string) bool {
	xv, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		return false
	}
	yv, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return false
	}

	switch op {
	case "==":
		return xv == yv
	case "!=":
		return xv != yv
	case "<":
		return xv < yv
	case ">":
		return xv > yv
	case "<=":
		return xv <= yv
	case ">=":
		return xv >= yv
	default:
		panic("unsupported arithmetic operator: " + op)
	}
}

func FilesHaveSameDevAndIno(file1, file2 string) bool {
	file1Info, err := os.Lstat(file1)
	if err != nil {
		return false
	}

	file2Info, err := os.Lstat(file2)
	if err != nil {
		return false
	}

	file1Stat, ok := file1Info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	file2Stat, ok := file2Info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return file1Stat.Dev == file2Stat.Dev && file1Stat.Ino == file2Stat.Ino
}

func FileIsOlderThan(file1, file2 string) bool {
	file2Info, err := os.Lstat(file2)
	if err != nil {
		return false
	}

	file1Info, err := os.Lstat(file1)
	if err != nil {
		return os.IsNotExist(err)
	}

	return file1Info.ModTime().Before(file2Info.ModTime())

}

func FileExists(file string) bool {
	_, err := os.Lstat(file)
	return err == nil || (!os.IsNotExist(err) && !os.IsPermission(err))
}

func DirectoryExists(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func BlockSpecialFileExists(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return (stat.Mode & syscall.S_IFMT) == syscall.S_IFBLK
}

func CharacterSpecialFileExists(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}

func RegularFileExists(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode().IsRegular()
}

func FileSGIDIsSet(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSetgid != 0
}

func FileSUIDIsSet(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSetuid != 0
}

func FileIsOwnedByEffectiveGroup(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return stat.Gid == uint32(os.Getgid())
}

func FileIsOwnedByEffectiveUser(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return stat.Uid == uint32(os.Getuid())
}

func FileIsSymbolic(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSymlink != 0
}

func FileIsSticky(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSticky != 0
}

func FileIsFIFO(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeNamedPipe != 0
}

func FileIsReadable(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&0400 != 0
}

func FileIsWritable(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&0200 != 0
}

func FileIsExecutable(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&0100 != 0
}

func FileHasAPositiveSize(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return stat.Size > 0
}

func FileHasBeenModifiedSinceLastRead(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return stat.Mtim.Nsec > stat.Atim.Nsec
}

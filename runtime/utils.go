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
	file1Info, err := os.Stat(file1)
	if err != nil {
		return false
	}

	file2Info, err := os.Stat(file2)
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

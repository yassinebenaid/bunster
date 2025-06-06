package runtime

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/yassinebenaid/bunster/runtime/pattern"
)

func NumberCompare(x, op, y string) bool {
	xv, err := strconv.Atoi(x)
	if err != nil {
		return false
	}
	yv, err := strconv.Atoi(y)
	if err != nil {
		return false
	}

	return CompareInt(xv, op, yv) == 1
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

func FileIsOlderThan(file1, file2 string) bool {
	file2Info, err := os.Stat(file2)
	if err != nil {
		return false
	}

	file1Info, err := os.Stat(file1)
	if err != nil {
		return os.IsNotExist(err)
	}

	return file1Info.ModTime().Before(file2Info.ModTime())

}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || (!os.IsNotExist(err) && !os.IsPermission(err))
}

func DirectoryExists(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func BlockSpecialFileExists(file string) bool {
	info, err := os.Stat(file)
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
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}

func RegularFileExists(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode().IsRegular()
}

func FileSGIDIsSet(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSetgid != 0
}

func FileSUIDIsSet(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSetuid != 0
}

func FileIsOwnedByEffectiveGroup(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return int(stat.Gid) == os.Getgid()
}

func FileIsOwnedByEffectiveUser(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return int(stat.Uid) == os.Getuid()
}

func FileIsSymbolic(file string) bool {
	info, err := os.Lstat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSymlink != 0
}

func FileIsSticky(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSticky != 0
}

func FileIsFIFO(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeNamedPipe != 0
}

func FileIsReadable(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&0400 != 0
}

func FileIsWritable(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&0200 != 0
}

func FileIsExecutable(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&0100 != 0
}

func FileHasAPositiveSize(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return stat.Size > 0
}

func FileIsSocket(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSocket != 0
}

func PatternMatch(str string, p string) bool {
	rx, err := pattern.Regexp(p, pattern.EntireString)
	if err != nil {
		return false
	}
	return regexp.MustCompile(rx).MatchString(str)
}

func Substring(str string, offset, length int) string {
	runes := []rune(str)
	n := len(runes)

	// Handle negative offset (from end)
	if offset < 0 {
		offset = n + offset
	}
	if offset < 0 || offset > n {
		return ""
	}

	// Handle negative length like Bash: treat as endpoint (from end)
	var end int
	if length < 0 {
		end = n + length
	} else {
		end = offset + length
	}

	// Clamp bounds
	if end < offset {
		return ""
	}
	if end > n {
		end = n
	}

	return string(runes[offset:end])
}

func ChangeStringCase(upper bool, str string, _pattern string, all bool) string {
	rx, err := pattern.Regexp(_pattern, pattern.Filenames|pattern.EntireString)
	if err != nil {
		return str
	}
	regx := regexp.MustCompile(rx)

	var result []rune
	var matched = false

	for _, ch := range str {
		var s = string(ch)
		if (all || !matched) && regx.MatchString(s) {
			matched = true
			if upper {
				s = strings.ToUpper(s)
			} else {
				s = strings.ToLower(s)
			}
		}
		result = append(result, []rune(s)...)
	}

	return string(result)
}

func RemoveMatchingPrefix(str string, _pattern string, longest bool) string {
	mode := pattern.Shortest
	if longest {
		mode = 0
	}
	rx, err := pattern.Regexp(_pattern, mode)
	if err != nil {
		return str
	}
	regx := regexp.MustCompile("^" + rx)

	match := regx.FindString(str)

	return strings.TrimPrefix(str, match)
}

func RemoveMatchingSuffix(str string, _pattern string, longest bool) string {
	mode := pattern.Shortest
	if longest {
		mode = 0
	}
	rx, err := pattern.Regexp(_pattern, mode)
	if err != nil {
		return str
	}
	regx := regexp.MustCompile(rx + "$")

	matches := regx.FindAllString(str, -1)
	if len(matches) == 0 {
		matches = append(matches, "")
	}

	return strings.TrimSuffix(str, matches[len(matches)-1])
}

func ReplaceMatching(str string, _pattern string, replace string, all bool) string {
	if str == "" {
		return str
	}
	rx, err := pattern.Regexp(_pattern, 0)
	if err != nil {
		return str
	}
	regx := regexp.MustCompile(rx)
	regx.Longest()

	if all {
		return regx.ReplaceAllString(str, replace)
	}

	first := regx.FindString(str)

	return strings.Replace(str, first, replace, 1)
}

func ReplaceMatchingPrefix(str string, _pattern string, replace string) string {
	rx, err := pattern.Regexp(_pattern, 0)
	if err != nil {
		return str
	}
	regx := regexp.MustCompile("^" + rx)
	regx.Longest()
	first := regx.FindString(str)
	if first == "" {
		return str
	}

	return strings.Replace(str, first, replace, 1)
}

func ReplaceMatchingSuffix(str string, _pattern string, replace string) string {
	rx, err := pattern.Regexp(_pattern, 0)
	if err != nil {
		return str
	}
	regx := regexp.MustCompile(rx + "$")
	regx.Longest()

	matches := regx.FindAllString(str, -1)
	if len(matches) == 0 {
		return str
	}

	return strings.TrimSuffix(str, matches[len(matches)-1]) + replace
}

package builtin

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/yassinebenaid/bunster/runtime"
)

func Loadenv(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	fs := flag.NewFlagSet("loadenv", flag.ContinueOnError)
	fs.SetOutput(stderr)
	export := fs.Bool("X", false, "mark variables as exported")
	if err := fs.Parse(shell.Args); err != nil {
		shell.ExitCode = 1
		return
	}

	files := fs.Args()
	if len(files) == 0 {
		files = append(files, ".env")
	}

	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(stderr, "loadenv: %s\n", err)
			shell.ExitCode = 1
			return
		}

		vars, err := parse(content)
		if err != nil {
			fmt.Fprintf(stderr, "loadenv: %s\n", err)
			shell.ExitCode = 1
			return
		}

		for key, value := range vars {
			shell.SetVar(key, value)
			if *export {
				shell.MarkVarAsExported(key)
			}
		}
	}

	shell.ExitCode = 0
}

func parse(src []byte) (map[string]string, error) {
	out := make(map[string]string)

	src = bytes.Replace(src, []byte("\r\n"), []byte("\n"), -1)
	cutset := src
	for {
		cutset = getStatementStart(cutset)
		if cutset == nil {
			// reached end of file
			break
		}

		key, left, err := locateKeyName(cutset)
		if err != nil {
			return nil, err
		}

		value, left, err := extractVarValue(left, out)
		if err != nil {
			return nil, err
		}

		out[key] = value
		cutset = left
	}

	return out, nil
}

// getStatementPosition returns position of statement begin.
//
// It skips any comment line or non-whitespace character.
func getStatementStart(src []byte) []byte {
	pos := indexOfNonSpaceChar(src)
	if pos == -1 {
		return nil
	}

	src = src[pos:]
	if src[0] != '#' {
		return src
	}

	// skip comment section
	pos = bytes.IndexFunc(src, isCharFunc('\n'))
	if pos == -1 {
		return nil
	}

	return getStatementStart(src[pos:])
}

// locateKeyName locates and parses key name and returns rest of slice
func locateKeyName(src []byte) (key string, cutset []byte, err error) {
	// trim "export" and space at beginning
	src = bytes.TrimLeftFunc(src, isSpace)
	if bytes.HasPrefix(src, []byte("export")) {
		trimmed := bytes.TrimPrefix(src, []byte("export"))
		if bytes.IndexFunc(trimmed, isSpace) == 0 {
			src = bytes.TrimLeftFunc(trimmed, isSpace)
		}
	}

	// locate key name end and validate it in single loop
	offset := 0
loop:
	for i, char := range src {
		rchar := rune(char)
		if isSpace(rchar) {
			continue
		}

		switch char {
		case '=', ':':
			// library also supports yaml-style value declaration
			key = string(src[0:i])
			offset = i + 1
			break loop
		case '_':
		default:
			// variable name should match [A-Za-z0-9_.]
			if unicode.IsLetter(rchar) || unicode.IsNumber(rchar) || rchar == '.' {
				continue
			}

			return "", nil, fmt.Errorf(
				`unexpected character %q in variable name near %q`,
				string(char), string(src))
		}
	}

	if len(src) == 0 {
		return "", nil, errors.New("zero length string")
	}

	// trim whitespace
	key = strings.TrimRightFunc(key, unicode.IsSpace)
	cutset = bytes.TrimLeftFunc(src[offset:], isSpace)
	return key, cutset, nil
}

// extractVarValue extracts variable value and returns rest of slice
func extractVarValue(src []byte, vars map[string]string) (value string, rest []byte, err error) {
	quote, hasPrefix := hasQuotePrefix(src)
	if !hasPrefix {
		// unquoted value - read until end of line
		endOfLine := bytes.IndexFunc(src, isLineEnd)

		// Hit EOF without a trailing newline
		if endOfLine == -1 {
			endOfLine = len(src)

			if endOfLine == 0 {
				return "", nil, nil
			}
		}

		// Convert line to rune away to do accurate countback of runes
		line := []rune(string(src[0:endOfLine]))

		// Assume end of line is end of var
		endOfVar := len(line)
		if endOfVar == 0 {
			return "", src[endOfLine:], nil
		}

		// Work backwards to check if the line ends in whitespace then
		// a comment, ie: foo=bar # baz # other
		for i := 0; i < endOfVar; i++ {
			if line[i] == '#' && i < endOfVar {
				if isSpace(line[i-1]) {
					endOfVar = i
					break
				}
			}
		}

		trimmed := strings.TrimFunc(string(line[0:endOfVar]), isSpace)

		return expand(trimmed, vars), src[endOfLine:], nil
	}

	// lookup quoted string terminator
	for i := 1; i < len(src); i++ {
		if char := src[i]; char != quote {
			continue
		}

		// skip escaped quote symbol (\" or \', depends on quote)
		if prevChar := src[i-1]; prevChar == '\\' {
			continue
		}

		// trim quotes
		trimFunc := isCharFunc(rune(quote))
		value = string(bytes.TrimLeftFunc(bytes.TrimRightFunc(src[0:i], trimFunc), trimFunc))
		if quote == '"' {
			// unescape newlines for double quote (this is compat feature)
			// and expand environment variables
			value = expand(expandEscapes(value), vars)
		}

		return value, src[i+1:], nil
	}

	// return formatted error if quoted string is not terminated
	valEndIndex := bytes.IndexFunc(src, isCharFunc('\n'))
	if valEndIndex == -1 {
		valEndIndex = len(src)
	}

	return "", nil, fmt.Errorf("unterminated quoted value %s", src[:valEndIndex])
}

func expandEscapes(str string) string {
	out := escapeRegex.ReplaceAllStringFunc(str, func(match string) string {
		c := strings.TrimPrefix(match, `\`)
		switch c {
		case "n":
			return "\n"
		case "r":
			return "\r"
		default:
			return match
		}
	})
	return unescapeCharsRegex.ReplaceAllString(out, "$1")
}

func indexOfNonSpaceChar(src []byte) int {
	return bytes.IndexFunc(src, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
}

// hasQuotePrefix reports whether charset starts with single or double quote and returns quote character
func hasQuotePrefix(src []byte) (prefix byte, isQuored bool) {
	if len(src) == 0 {
		return 0, false
	}

	switch prefix := src[0]; prefix {
	case '"', '\'':
		return prefix, true
	default:
		return 0, false
	}
}

func isCharFunc(char rune) func(rune) bool {
	return func(v rune) bool {
		return v == char
	}
}

// isSpace reports whether the rune is a space character but not line break character
//
// this differs from unicode.IsSpace, which also applies line break as space
func isSpace(r rune) bool {
	switch r {
	case '\t', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

func isLineEnd(r rune) bool {
	if r == '\n' || r == '\r' {
		return true
	}
	return false
}

var (
	escapeRegex        = regexp.MustCompile(`\\.`)
	expandVarRegex     = regexp.MustCompile(`(\\)?(\$)(\()?\{?([A-Z0-9_]+)?\}?`)
	unescapeCharsRegex = regexp.MustCompile(`\\([^$])`)
)

func expand(v string, m map[string]string) string {
	return expandVarRegex.ReplaceAllStringFunc(v, func(s string) string {
		submatch := expandVarRegex.FindStringSubmatch(s)

		if submatch == nil {
			return s
		}
		if submatch[1] == "\\" || submatch[2] == "(" {
			return submatch[0][1:]
		} else if submatch[4] != "" {
			if val, ok := m[submatch[4]]; ok {
				return val
			}
			if val, ok := os.LookupEnv(submatch[4]); ok {
				return val
			}
			return m[submatch[4]]
		}
		return s
	})
}

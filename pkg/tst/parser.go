package tst

import (
	"errors"
	"strings"
)

// TestCase represents a single test case with input and output
type Test struct {
	Label  string
	Input  string
	Output string
}

const (
	START  = iota
	INPUT  = iota
	OUTPUT = iota
)

// Parse reads from an io.Reader and parses the test case format
func Parse(in []byte) ([]Test, error) {
	var lines = splitIntoLines(string(in))

	var step = START
	var tests []Test
	var test Test

	for _, line := range lines {
		if step == START {
			if strings.TrimSpace(line) == "" {
				continue
			}

			label := strings.TrimSpace(line)

			label, ok := strings.CutPrefix(label, "#(TEST:")
			if !ok {
				panic("expected '#(TEST:', found " + label)
			}

			label, ok = strings.CutSuffix(label, ")")
			if !ok {
				panic("expected ')' to close test label, found " + label)
			}

			test.Label = strings.TrimSpace(label)
			step = INPUT
			continue
		}

		if step == INPUT {
			if strings.TrimSpace(line) == "#(RESULT)" {
				step = OUTPUT
				continue
			}

			test.Input += line
			continue
		}

		if step == OUTPUT {
			if strings.TrimSpace(line) == "#(ENDTEST)" {
				step = START
				tests = append(tests, test)
				continue
			}

			test.Output += line
			continue
		}
	}

	if step == INPUT {
		return nil, errors.New("bad test syntax, coundl't find #(RESULT) section")
	}

	return tests, nil
}

// SplitIntoLines splits a string into a slice of lines, preserving newlines.
func splitIntoLines(input string) []string {
	var lines []string
	var currentLine strings.Builder

	for _, char := range input {
		currentLine.WriteRune(char)
		if char == '\n' {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}
	}

	// Append the last line if it doesn't end with a newline
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
	}

	return lines
}

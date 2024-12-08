package tst

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// TestCase represents a single test case with input and output
type TestCase struct {
	Input  string
	Output string
}

// Test represents the entire test structure
type Test struct {
	Label string
	Cases []TestCase
}

// Parser for the specific test case text format
type Parser struct{}

// Parse reads from an io.Reader and parses the test case format
func (p *Parser) Parse(reader io.Reader) (*Test, error) {
	scanner := bufio.NewScanner(reader)

	// First line must be "Test:" followed by label
	if !scanner.Scan() {
		return nil, errors.New("empty input")
	}
	firstLine := scanner.Text()
	if !strings.HasPrefix(firstLine, "Test:") {
		return nil, fmt.Errorf("invalid format: must start with 'Test:', got %q", firstLine)
	}

	// Extract label
	label := strings.TrimSpace(strings.TrimPrefix(firstLine, "Test:"))

	// Second line must be separator
	if !scanner.Scan() || strings.TrimSpace(scanner.Text()) != "------------" {
		return nil, errors.New("missing separator after label")
	}

	test := &Test{
		Label: label,
		Cases: []TestCase{},
	}

	var currentCase *TestCase
	var state string // "start", "waiting_input", "input", "waiting_output", "output"

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, "------input------"):
			if currentCase != nil && (currentCase.Input == "" || currentCase.Output == "") {
				return nil, errors.New("incomplete previous test case")
			}
			currentCase = &TestCase{}
			state = "waiting_input"
			continue

		case strings.Contains(line, "------output------"):
			if currentCase == nil {
				return nil, errors.New("output marker before input marker")
			}
			if currentCase.Input == "" {
				return nil, errors.New("output marker before input content")
			}
			state = "waiting_output"
			continue

		case strings.TrimSpace(line) == "------------":
			if currentCase == nil {
				continue
			}
			// Validate that current case has both input and output
			if currentCase.Input == "" || currentCase.Output == "" {
				return nil, errors.New("incomplete test case before delimiter")
			}
			test.Cases = append(test.Cases, *currentCase)
			currentCase = nil
			state = "start"
			continue
		}

		// Append content based on current state
		switch state {
		case "waiting_input":
			if strings.TrimSpace(line) != "" {
				currentCase.Input = line
				state = "input"
			}
		case "input":
			currentCase.Input += "\n" + line
		case "waiting_output":
			if strings.TrimSpace(line) != "" {
				currentCase.Output = line
				state = "output"
			}
		case "output":
			currentCase.Output += "\n" + line
		}
	}

	// Handle last case
	if currentCase != nil {
		if currentCase.Input == "" || currentCase.Output == "" {
			return nil, errors.New("incomplete final test case")
		}
		test.Cases = append(test.Cases, *currentCase)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Validate test has at least one case
	if len(test.Cases) == 0 {
		return nil, errors.New("no test cases found")
	}

	return test, nil
}

package flag

import (
	"errors"
	"fmt"
	"strings"
)

// FlagType defines the type of flag
type FlagType int

const (
	// Boolean flags don't require an argument
	Boolean FlagType = iota
	// String flags require an argument
	String
)

// Flag represents a command line flag
type Flag struct {
	Name     string
	Type     FlagType
	Required bool
	Value    interface{}
}

// ParseResult contains the parsed flags and remaining arguments
type ParseResult struct {
	Flags map[string]interface{}
	Args  []string
}

// Parser contains the flag definitions and parsing logic
type Parser struct {
	flags map[string]*Flag
}

// NewParser creates a new flag parser
func NewParser() *Parser {
	return &Parser{
		flags: make(map[string]*Flag),
	}
}

// AddFlag adds a new flag to the parser
func (p *Parser) AddFlag(name string, flagType FlagType, required bool) error {
	if len(name) != 1 {
		return errors.New("flag name must be exactly one character")
	}

	p.flags[name] = &Flag{
		Name:     name,
		Type:     flagType,
		Required: required,
	}

	return nil
}

// Parse parses the command line arguments
func (p *Parser) Parse(args []string) (*ParseResult, error) {
	result := &ParseResult{
		Flags: make(map[string]interface{}),
		Args:  []string{},
	}

	// Initialize all Boolean flags to false
	for name, flag := range p.flags {
		if flag.Type == Boolean {
			result.Flags[name] = false
		}
	}

	i := 0
	for i < len(args) {
		arg := args[i]
		i++

		// Check if this is a flag argument
		if !strings.HasPrefix(arg, "-") || len(arg) <= 1 {
			// This is not a flag, add it to remaining args
			result.Args = append(result.Args, arg)
			continue
		}

		// Remove the leading dash
		flagGroup := arg[1:]

		// First pass: count how many string arguments we need
		stringArgsNeeded := 0
		flagsInGroup := make([]string, 0, len(flagGroup))

		for _, ch := range flagGroup {
			name := string(ch)
			flagsInGroup = append(flagsInGroup, name)

			flag, exists := p.flags[name]
			if !exists {
				return nil, fmt.Errorf("unknown flag: %s", name)
			}

			if flag.Type == String {
				stringArgsNeeded++
			}
		}

		// Check if we have enough remaining arguments for the string flags
		if i+stringArgsNeeded > len(args) {
			return nil, fmt.Errorf("not enough arguments for flags in group: %s", arg)
		}

		// Second pass: set the values
		for _, name := range flagsInGroup {
			flag := p.flags[name]

			if flag.Type == Boolean {
				result.Flags[name] = true
			} else { // String flag
				if i >= len(args) {
					return nil, fmt.Errorf("missing value for flag: %s", name)
				}

				result.Flags[name] = args[i]
				i++ // Consume the next argument as the value
			}
		}
	}

	// Check if all required flags are provided
	for name, flag := range p.flags {
		if flag.Required {
			if _, exists := result.Flags[name]; !exists {
				return nil, fmt.Errorf("required flag not provided: %s", name)
			}
		}
	}

	return result, nil
}

// GetValue returns the parsed value for a flag
func (p *Parser) GetValue(name string) (interface{}, bool) {
	flag, exists := p.flags[name]
	if !exists {
		return nil, false
	}
	return flag.Value, true
}

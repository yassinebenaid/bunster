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
}

// ParseResult contains the parsed flags and remaining arguments
type ParseResult struct {
	Flags map[string]interface{}
	Args  []string
}

// Parser contains the flag definitions and parsing logic
type Parser struct {
	shortFlags map[string]*Flag
	longFlags  map[string]*Flag
}

// NewParser creates a new flag parser
func NewParser() *Parser {
	return &Parser{
		shortFlags: make(map[string]*Flag),
		longFlags:  make(map[string]*Flag),
	}
}

// AddShortFlag adds a new single-character flag to the parser
func (p *Parser) AddShortFlag(name string, flagType FlagType, required bool) error {
	if len(name) != 1 {
		return errors.New("short flag name must be exactly one character")
	}

	p.shortFlags[name] = &Flag{
		Name:     name,
		Type:     flagType,
		Required: required,
	}

	return nil
}

// AddLongFlag adds a new multi-character flag to the parser
func (p *Parser) AddLongFlag(name string, flagType FlagType, required bool) error {
	if len(name) <= 1 {
		return errors.New("long flag name must be more than one character")
	}

	p.longFlags[name] = &Flag{
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
	}

	// Initialize all Boolean flags to false
	for name, flag := range p.shortFlags {
		if flag.Type == Boolean {
			result.Flags[name] = false
		}
	}

	for name, flag := range p.longFlags {
		if flag.Type == Boolean {
			result.Flags[name] = false
		}
	}

	i := 0
	for i < len(args) {
		arg := args[i]
		i++

		// Check if this is a flag argument
		if !strings.HasPrefix(arg, "-") {
			// This is not a flag, add it to remaining args
			result.Args = append(result.Args, arg)
			continue
		}

		// Check if it's a long flag (--flag)
		if strings.HasPrefix(arg, "--") {
			if len(arg) <= 2 {
				result.Args = append(result.Args, args[i:]...)
				break
			}

			flagName := arg[2:] // trim the -- from the beginning

			// if it's value is associated, like --flag=value
			if strings.Contains(flagName, "=") {
				fields := strings.SplitN(flagName, "=", 2)
				flag, exists := p.longFlags[fields[0]]

				if !exists {
					return nil, fmt.Errorf("unknown long flag: %s", fields[0])
				}

				if flag.Type == Boolean {
					return nil, fmt.Errorf("passing value to a flag that doesn't expect it: %s", fields[0])
				} else { // String flag
					if len(fields) != 2 || fields[1] == "" {
						return nil, fmt.Errorf("missing value for flag: %s", fields[0])
					}

					result.Flags[fields[0]] = fields[1]
				}
				continue
			}

			flag, exists := p.longFlags[flagName]

			if !exists {
				return nil, fmt.Errorf("unknown long flag: %s", flagName)
			}

			if flag.Type == Boolean {
				result.Flags[flagName] = true
			} else { // String flag
				if i >= len(args) {
					return nil, fmt.Errorf("missing value for flag: %s", flagName)
				}

				if strings.HasPrefix(args[i], "-") {
					return nil, fmt.Errorf("missing value for flag: %s", flagName)
				}

				result.Flags[flagName] = args[i]
				i++ // Consume the next argument as the value
			}

			continue
		}

		// Short flag processing (-f or -abc)
		if len(arg) <= 1 {
			return nil, fmt.Errorf("invalid short flag format: %s", arg)
		}

		// Remove the leading dash
		flagGroup := arg[1:]

		// First pass: count how many string arguments we need
		stringArgsNeeded := 0
		flagsInGroup := make([]string, 0, len(flagGroup))

		for _, ch := range flagGroup {
			name := string(ch)
			flagsInGroup = append(flagsInGroup, name)

			flag, exists := p.shortFlags[name]
			if !exists {
				return nil, fmt.Errorf("unknown short flag: %s", name)
			}

			if flag.Type == String {
				stringArgsNeeded++
			}
		}

		// Second pass: set the values
		for _, name := range flagsInGroup {
			flag := p.shortFlags[name]

			if flag.Type == Boolean {
				result.Flags[name] = true
			} else { // String flag
				if i >= len(args) {
					return nil, fmt.Errorf("missing value for flag: %s", name)
				}

				if strings.HasPrefix(args[i], "-") {
					return nil, fmt.Errorf("missing value for flag: %s", name)
				}

				result.Flags[name] = args[i]
				i++ // Consume the next argument as the value
			}
		}
	}

	// Check if all required flags are provided
	for name, flag := range p.shortFlags {
		if flag.Required {
			if _, exists := result.Flags[name]; !exists {
				return nil, fmt.Errorf("required flag not provided: %s", name)
			}
		}
	}

	for name, flag := range p.longFlags {
		if flag.Required {
			if _, exists := result.Flags[name]; !exists {
				return nil, fmt.Errorf("required flag not provided: %s", name)
			}
		}
	}

	return result, nil
}

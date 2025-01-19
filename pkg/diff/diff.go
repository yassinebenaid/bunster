package diff

import (
	"fmt"
	"strings"
)

// ANSI color codes
const (
	ansiFourgroundRed   = "\033[31m"
	ansiFourgroundGreen = "\033[32m"
	ansiBackgroundRed   = "\033[41m"
	ansiBackgroundGreen = "\033[42m"
	ansiReset           = "\033[0m"
)

// Diff generates a git-like diff between two strings.
func Diff(original, modified string) string {
	return diffStrings(original, modified, ansiFourgroundRed, ansiFourgroundGreen)
}

// DiffBG  is same as Diff but uses background colors.
func DiffBG(original, modified string) string {
	return diffStrings(original, modified, ansiBackgroundRed, ansiBackgroundGreen)
}

func diffStrings(original, modified, colorRed, colorGreen string) string {
	originalLines := strings.Split(original, "\n")
	modifiedLines := strings.Split(modified, "\n")

	// Compute the LCS and the diffs
	ops := computeDiff(originalLines, modifiedLines)

	var diffOutput []string

	for _, op := range ops {
		switch op.Type {
		case "unchanged":
			diffOutput = append(diffOutput, fmt.Sprintf("  %s", op.Line))
		case "delete":
			diffOutput = append(diffOutput, fmt.Sprintf("%s- %s%s", colorRed, op.Line, ansiReset))
		case "add":
			diffOutput = append(diffOutput, fmt.Sprintf("%s+ %s%s", colorGreen, op.Line, ansiReset))
		}
	}

	return strings.Join(diffOutput, "\n")
}

// DiffOperation represents a single diff operation
type DiffOperation struct {
	Type string // "add", "delete", or "unchanged"
	Line string
}

// computeDiff calculates the diff between two slices of lines using the LCS algorithm
func computeDiff(original, modified []string) []DiffOperation {
	m, n := len(original), len(modified)
	lcs := make([][]int, m+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}

	// Build the LCS table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if original[i-1] == modified[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	// Backtrack to determine the diff
	var ops []DiffOperation
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && original[i-1] == modified[j-1] {
			ops = append([]DiffOperation{{Type: "unchanged", Line: original[i-1]}}, ops...)
			i--
			j--
		} else if j > 0 && (i == 0 || lcs[i][j-1] >= lcs[i-1][j]) {
			ops = append([]DiffOperation{{Type: "add", Line: modified[j-1]}}, ops...)
			j--
		} else {
			ops = append([]DiffOperation{{Type: "delete", Line: original[i-1]}}, ops...)
			i--
		}
	}

	return ops
}

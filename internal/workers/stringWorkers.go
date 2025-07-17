package workers

import (
	"regexp"
	"strconv"
)

var (
	bracketRegex = regexp.MustCompile(`([\[\]])\w*`)
	numberRegex  = regexp.MustCompile(`[\d]`)
	colonRegex   = regexp.MustCompile(`([\:])\w*`)
)

// HasBrackets checks if a string contains brackets.
func HasBrackets(s string) bool {
	return bracketRegex.MatchString(s)
}

// RemoveBrackets removes brackets from a string.
func RemoveBrackets(s string) string {
	return bracketRegex.ReplaceAllString(s, "")
}

// NumStringToInt converts a number string to an integer.
func NumStringToInt(s string) int {
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	default:
		i, err := strconv.Atoi(s)
		if err != nil {
			return 1
		}
		return i
	}
}

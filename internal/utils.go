package internal

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gabrielfmcoelho/platform-core/domain"
)

// Parse number in string type and returns a uint
func ParseUint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, domain.ErrInvalidNumberToParse
	}
	return uint(i), nil
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// parseDelimitedStrings splits a semicolon-delimited string (e.g. "foo;bar;baz")
// into a []string{"foo", "bar", "baz"}. Trims spaces as well.
func ParseDelimitedStrings(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ";")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// Parse time.Duration to int (seconds)
func ToSeconds(d time.Duration) int {
	return int(d.Seconds())
}

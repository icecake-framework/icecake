package helper

import (
	"strings"
)

func Normalize(s string) string {
	return strings.ToLower(strings.Trim(s, " "))
}

func NormalizeUp(s string) string {
	return strings.ToUpper(strings.Trim(s, " "))
}

// FillLen fills s with blanks to make it a string of length l.
// If l > len(s) then return s
func FillString(s string, l int) string {
	filler := l - len(s)
	if filler <= 0 {
		return s
	}
	s += strings.Repeat(" ", filler)
	return s
}

package helper

import "strings"

func Normalize(s string) string {
	return strings.ToLower(strings.Trim(s, " "))
}

func NormalizeUp(s string) string {
	return strings.ToUpper(strings.Trim(s, " "))
}

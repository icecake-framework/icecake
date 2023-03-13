package helper

import "strings"

func Normalize(s string) string {
	return strings.ToLower(strings.Trim(s, " "))
}

package browser

import "strings"

func normalize(s string) string {
	return strings.ToLower(strings.Trim(s, " "))
}

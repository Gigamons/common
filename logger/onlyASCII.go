package logger

import (
	"regexp"
)

// OnlyASCII = Fix for console exploit.
func OnlyASCII(s string) string {
	return regexp.MustCompile("[[:^ascii:]]").ReplaceAllLiteralString(s, "")
}

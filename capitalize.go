package goUtils

import (
	"strings"
)

func Capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

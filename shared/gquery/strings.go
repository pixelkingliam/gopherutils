package gquery

import (
	"strings"
)

func EndsWith(str string, compare string) bool {
	if len(str) < len(compare) {
		return false
	}
	if strings.Compare(str[len(str)-len(compare):], compare) == 0 {
		return true
	}
	return false
}
func StartsWith(str string, compare string) bool {
	if len(str) < len(compare) {
		return false
	}
	if strings.Compare(str[:len(compare)], compare) == 0 {
		return true
	}
	return false
}

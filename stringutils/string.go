package stringutils

import (
	"strings"
	"unicode"
)

// TrimSpace 去除字符串首尾空格
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// IsAllUpper 判断字符串是否全为大写
func IsAllUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAllLower 判断字符串是否全为小写
func IsAllLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
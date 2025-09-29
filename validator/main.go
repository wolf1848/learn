package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}

func IsMinLen(value string, length int) bool {
	return utf8.RuneCountInString(value) < length
}

func IsMaxLen(value string, length int) bool {
	return utf8.RuneCountInString(value) > length
}

func IsEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	return err == nil && matched
}

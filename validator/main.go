package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type ValidateErrors map[string][]string

type Validator struct {
	data ValidateErrors
}

func (v *Validator) IsValid() bool {
	return len(v.data) == 0
}

func (v *Validator) AddProblem(field, problem string) {
	v.data[field] = append(v.data[field], problem)
}

func (v *Validator) GetProblems() *ValidateErrors {
	return &v.data
}

func NewValidator() *Validator {
	return &Validator{
		data: make(map[string][]string),
	}
}

func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}

func IsMinLen(value string, length int) bool {
	return utf8.RuneCountInString(value) < length
}

func IsMaxLen(value string, length int) bool {
	return utf8.RuneCountInString(value) > length
}

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsEmail(email string) bool {
	/*pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	return err == nil && matched*/

	return emailPattern.MatchString(email)
}

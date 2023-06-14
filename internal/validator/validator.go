package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

var EmailRegex = regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddFieldError(key, value string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = value
	}
}

func (v *Validator) CheckValid(ok bool, key, value string) {
	if !ok {
		v.AddFieldError(key, value)
	}
}

func FieldNotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func FieldMaxSize(value string, length int) bool {
	return utf8.RuneCountInString(value) <= length
}

func FieldMinSize(value string, length int) bool {
	return utf8.RuneCountInString(value) >= length
}

func MatchesToRegex(value string, rgx *regexp.Regexp) bool {
	return rgx.MatchString(value)
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func (v *Validator) AddNonFieldError(value string) {
	v.NonFieldErrors = append(v.NonFieldErrors, value)
}

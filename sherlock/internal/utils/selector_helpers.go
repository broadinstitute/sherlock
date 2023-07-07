package utils

// selector_helpers.go contains simple but frequently-used functions helpful for different data types
// analyzing and validating their individual selectors

import (
	"strings"
	"unicode"
)

func CanonicalizeSelector(selector string) string {
	return strings.Trim(selector, "/")
}

func IsNumeric(selector string) bool {
	for _, r := range selector {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func IsAlphaNumeric(selector string) bool {
	for _, r := range selector {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLowerAlphaNumeric(selector string) bool {
	for _, r := range selector {
		if !unicode.IsDigit(r) && !(unicode.IsLetter(r) && unicode.IsLower(r)) {
			return false
		}
	}
	return true
}

func IsAlphaNumericWithHyphens(selector string) bool {
	for _, r := range selector {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) && r != '-' {
			return false
		}
	}
	return true
}

func IsStartingWithLetter(selector string) bool {
	return unicode.IsLetter(rune(selector[0]))
}

func IsEndingWithAlphaNumeric(selector string) bool {
	r := rune(selector[len(selector)-1])
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

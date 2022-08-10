package v2models

// selector_helpers.go contains simple but frequently-used functions helpful for different data types
// analyzing and validating their individual selectors

import "unicode"

func isNumeric(selector string) bool {
	for _, r := range selector {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isAlphaNumericWithHyphens(selector string) bool {
	for _, r := range selector {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) && r != '-' {
			return false
		}
	}
	return true
}

func isStartingWithLetter(selector string) bool {
	return unicode.IsLetter(rune(selector[0]))
}

func isEndingWithAlphaNumeric(selector string) bool {
	r := rune(selector[len(selector)-1])
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

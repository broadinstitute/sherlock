package sherlock

import "strings"

func canonicalizeSelector(selector string) string {
	return strings.Trim(selector, "/")
}

package utils

import "strings"

func SubstituteSuffix(s string, suffixesToReplace []string, replacement string) string {
	for _, suffix := range suffixesToReplace {
		if strings.HasSuffix(s, suffix) {
			s = strings.TrimSuffix(s, suffix) + replacement
			break
		}
	}
	return s
}

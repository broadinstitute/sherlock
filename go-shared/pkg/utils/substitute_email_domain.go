package utils

import "strings"

func SubstituteEmailDomain(email string, oldDomains []string, newDomain string) string {
	for _, domain := range oldDomains {
		if strings.HasSuffix(email, "@"+domain) {
			email = strings.TrimSuffix(email, domain) + newDomain
			break
		}
	}
	return email
}

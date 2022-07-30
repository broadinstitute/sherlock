package auth

import "strings"

func emailToFirecloudEmail(email string) string {
	if strings.HasSuffix(email, "@broadinstitute.org") {
		email = strings.TrimSuffix(email, "@broadinstitute.org") + "@firecloud.org"
	}
	return email
}

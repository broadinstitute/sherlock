package auth

import "strings"

const firecloudDomain = "firecloud.org"

var firecloudGroups = struct {
	FcAdmins               string
	FirecloudProjectOwners string
}{
	FcAdmins:               "fc-admins@firecloud.org",
	FirecloudProjectOwners: "firecloud-project-owners@firecloud.org",
}

func emailToFirecloudEmail(email string) string {
	if strings.HasSuffix(email, "@broadinstitute.org") {
		email = strings.TrimSuffix(email, "@broadinstitute.org") + "@firecloud.org"
	}
	return email
}

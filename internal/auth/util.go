package auth

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"strings"
)

func emailToFirecloudEmail(email string) string {
	if strings.HasSuffix(email, "@"+config.Config.MustString("auth.broadinstitute.domain")) {
		email = strings.TrimSuffix(email, config.Config.MustString("auth.broadinstitute.domain")) + config.Config.MustString("auth.firecloud.domain")
	}
	return email
}

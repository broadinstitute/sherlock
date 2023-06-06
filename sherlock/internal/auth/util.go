package auth

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v2models"
	"strings"
)

func emailToFirecloudEmail(email string) string {
	if strings.HasSuffix(email, "@"+config.Config.MustString("auth.broadinstitute.domain")) {
		email = strings.TrimSuffix(email, config.Config.MustString("auth.broadinstitute.domain")) + config.Config.MustString("auth.firecloud.domain")
	}
	return email
}

func storedUserToUser(storedUser v2models.User, authMethod auth_models.AuthMethod, via *auth_models.User) *auth_models.User {
	return &auth_models.User{
		ID:                         storedUser.ID,
		StoredControlledUserFields: storedUser.StoredControlledUserFields,
		StoredMutableUserFields:    storedUser.StoredMutableUserFields,
		InferredUserFields: auth_models.InferredUserFields{
			MatchedFirecloudAccount: cachedFirecloudAccounts[emailToFirecloudEmail(storedUser.Email)],
			MatchedExtraPermissions: cachedExtraPermissions[storedUser.Email],
		},
		AuthMethod: authMethod,
		Via:        via,
	}
}

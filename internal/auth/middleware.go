package auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/auth/iap_auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/gin-gonic/gin"
)

const contextUserKey = "SherlockUser"

func IapUserMiddleware(userStore *v2models.UserMiddlewareStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, googleID, err := iap_auth.ParseIAP(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}

		storedUser, err := userStore.GetOrCreateUser(email, googleID)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}

		ctx.Set(contextUserKey, &auth_models.User{
			StoredUserFields: storedUser.StoredUserFields,
			InferredUserFields: auth_models.InferredUserFields{
				MatchedFirecloudAccount: cachedFirecloudAccounts[emailToFirecloudEmail(email)],
				MatchedExtraPermissions: cachedExtraPermissions[email],
			},
		})

		ctx.Next()
	}
}

func FakeUserMiddleware(userStore *v2models.UserMiddlewareStore) gin.HandlerFunc {
	email := "fake@broadinstitute.org"
	googleID := "some id would go here"
	return func(ctx *gin.Context) {
		var firecloudAccount *auth_models.FirecloudAccount
		if ctx.GetHeader("Suitable") == "false" {
			firecloudAccount = nil
		} else {
			firecloudAccount = &auth_models.FirecloudAccount{
				Email:               email,
				AcceptedGoogleTerms: true,
				EnrolledIn2fa:       true,
				Suspended:           false,
				Archived:            false,
				SuspensionReason:    "",
				Groups: &auth_models.FirecloudGroupMembership{
					FcAdmins:               true,
					FirecloudProjectOwners: true,
				},
			}
		}

		storedUser, err := userStore.GetOrCreateUser(email, googleID)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}

		ctx.Set(contextUserKey, &auth_models.User{
			StoredUserFields: storedUser.StoredUserFields,
			InferredUserFields: auth_models.InferredUserFields{
				MatchedFirecloudAccount: firecloudAccount,
			},
		})

		ctx.Next()
	}
}

// ExtractUserFromContext is the counterpart to the middlewares provided by this package:
// handlers can call it to extract a User from the context.
func ExtractUserFromContext(ctx *gin.Context) (*auth_models.User, error) {
	userValue, exists := ctx.Get(contextUserKey)
	if !exists {
		return nil, fmt.Errorf("(%s) authentication middleware not present", errors.InternalServerError)
	}
	user, ok := userValue.(*auth_models.User)
	if !ok {
		return nil, fmt.Errorf("(%s) authentication middleware misconfigured: suitability represented as %T", errors.InternalServerError, userValue)
	}
	return user, nil
}

package auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/auth/gha_oidc_auth"
	"github.com/broadinstitute/sherlock/internal/auth/iap_auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const contextUserKey = "SherlockUser"

func IapUserMiddleware(userStore *v2models.MiddlewareUserStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, googleID, err := iap_auth.ParseHeader(ctx)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}

		storedIapUser, err := userStore.GetOrCreateUser(email, googleID)
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(err))
			return
		}

		iapUser := storedUserToUser(storedIapUser, auth_models.AuthMethodIAP, nil)

		headerPresent, githubUsername, githubID, err := gha_oidc_auth.ParseHeader(ctx)
		if headerPresent {
			if err != nil {
				ctx.JSON(errors.ErrorToApiResponse(err))
				return
			}

			storedGhaUser, err := userStore.GetGithubUserIfExists(githubID)
			if err != nil {
				ctx.JSON(errors.ErrorToApiResponse(err))
				return
			}

			if storedGhaUser != nil {
				ghaUser := storedUserToUser(*storedGhaUser, auth_models.AuthMethodGHA, iapUser)
				log.Info().Msgf("AUTH | substituted GHA OIDC JWT user %s over IAP JWT user %s", ghaUser.Email, iapUser.Email)
				ctx.Set(contextUserKey, ghaUser)
			} else {
				log.Info().Msgf("AUTH | ignored GHA OIDC JWT for unknown github user %s, still using IAP JWT user %s", githubUsername, iapUser.Email)
				ctx.Set(contextUserKey, iapUser)
			}
		} else {
			ctx.Set(contextUserKey, iapUser)
		}

		ctx.Next()
	}
}

func FakeUserMiddleware(userStore *v2models.MiddlewareUserStore) gin.HandlerFunc {
	email := "fake@broadinstitute.org"
	googleID := "fakeGoogleID"
	return func(ctx *gin.Context) {
		var firecloudAccount *auth_models.FirecloudAccount
		if ctx.GetHeader("Suitable") == "false" {
			firecloudAccount = nil
		} else {
			firecloudAccount = &auth_models.FirecloudAccount{
				Email:               emailToFirecloudEmail(email),
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
			ID:                         storedUser.ID,
			StoredControlledUserFields: storedUser.StoredControlledUserFields,
			StoredMutableUserFields:    storedUser.StoredMutableUserFields,
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

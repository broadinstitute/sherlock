package auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const contextUserKey = "SherlockUser"

func IapUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		iapJWT := ctx.GetHeader("X-Goog-IAP-JWT-Assertion")
		if iapJWT == "" {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) no 'X-Goog-IAP-JWT-Assertion' header set, IAP authentication required", errors.ProxyAuthenticationRequired)))
			return
		}
		// Sherlock is deployed behind an Apache proxy that checks that it is correctly wrapped by IAP, so we don't
		// actually care about exhaustively validating here (and we lack all the audience information to do so),
		// this is just the easiest way to decode the JWT payload.
		payload, err := idtoken.Validate(ctx, iapJWT, "")
		if err != nil {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) failed to validate IAP JWT in 'X-Goog-IAP-JWT-Assertion' header: %v", errors.ProxyAuthenticationRequired, err)))
			return
		}
		emailValue := payload.Claims["email"]
		email, ok := emailValue.(string)
		if !ok || email == "" {
			ctx.JSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) IAP JWT seemed to pass validation but lacked an email claim", errors.ProxyAuthenticationRequired)))
		}

		ctx.Set(contextUserKey, &User{
			AuthenticatedEmail:      email,
			MatchedFirecloudAccount: cachedFirecloudAccounts[emailToFirecloudEmail(email)],
		})

		ctx.Next()
	}
}

func FakeUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var firecloudAccount *FirecloudAccount
		if ctx.GetHeader("Suitable") == "false" {
			firecloudAccount = nil
		} else {
			firecloudAccount = &FirecloudAccount{
				Email:               "fake@broadinstitute.org",
				AcceptedGoogleTerms: true,
				EnrolledIn2fa:       true,
				Suspended:           false,
				Archived:            false,
				SuspensionReason:    "",
				Groups: &FirecloudGroupMembership{
					FcAdmins:               true,
					FirecloudProjectOwners: true,
				},
			}
		}
		ctx.Set(contextUserKey, &User{
			AuthenticatedEmail:      "fake@broadinstitute.org",
			MatchedFirecloudAccount: firecloudAccount,
		})

		ctx.Next()
	}
}

// ExtractUserFromContext is the counterpart to the middlewares provided by this package:
// handlers can call it to extract a User from the context.
func ExtractUserFromContext(ctx *gin.Context) (*User, error) {
	userValue, exists := ctx.Get(contextUserKey)
	if !exists {
		return nil, fmt.Errorf("(%s) authentication middleware not present", errors.InternalServerError)
	}
	user, ok := userValue.(*User)
	if !ok {
		return nil, fmt.Errorf("(%s) authentication middleware misconfigured: suitability represented as %T", errors.InternalServerError, userValue)
	}
	return user, nil
}

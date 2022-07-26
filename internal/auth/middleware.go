package auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const ContextUserKey = "UserSuitable"

func IdentityAwareProxyAuthentication() gin.HandlerFunc {
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

		ctx.Set(ContextUserKey, User{
			AuthenticatedEmail: email,
			suitabilityInfo:    getUserSuitabilityInfo(email),
		})

		ctx.Next()
	}
}

func DummyAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		suitability := &suitabilityInfo{
			acceptedWorkspaceTos: true,
			enrolledIn2fa:        true,
			suspended:            false,
			archived:             false,
			suspensionReason:     "",
		}
		if ctx.GetHeader("Suitable") == "false" {
			suitability = nil
		}
		ctx.Set(ContextUserKey, User{
			AuthenticatedEmail: "fake@broadinstitute.org",
			suitabilityInfo:    suitability,
		})

		ctx.Next()
	}
}

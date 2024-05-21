package iap

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const (
	header     = "X-Goog-IAP-JWT-Assertion"
	emailClaim = "email"
	subClaim   = "sub"
)

func ParseHeader(ctx *gin.Context) (email string, googleID string, err error) {
	iapJWT := ctx.GetHeader(header)
	if iapJWT == "" {
		return "", "", fmt.Errorf("(%s) no '%s' header set, IAP authentication required", errors.ProxyAuthenticationRequired, header)
	}

	// Sherlock is deployed behind an Apache proxy that checks that it is correctly wrapped by IAP, so we don't
	// actually care about exhaustively validating here (and we lack all the audience information to do so),
	// this is just the easiest way to decode the JWT payload.
	payload, err := idtoken.Validate(ctx, iapJWT, "")
	if err != nil {
		return "", "", fmt.Errorf("(%s) failed to validate IAP JWT in '%s' header: %w", errors.ProxyAuthenticationRequired, header, err)
	} else if payload == nil {
		return "", "", fmt.Errorf("(%s) IAP JWT seemed to pass validation but payload was nil", errors.ProxyAuthenticationRequired)
	}

	emailValue := payload.Claims[emailClaim]
	email, ok := emailValue.(string)
	if !ok || email == "" {
		return "", "", fmt.Errorf("(%s) IAP JWT seemed to pass validation but lacked an '%s' claim", errors.ProxyAuthenticationRequired, emailClaim)
	}

	subValue := payload.Claims[subClaim]
	googleID, ok = subValue.(string)
	if !ok || googleID == "" {
		return "", "", fmt.Errorf("(%s) IAP JWT seemed to pass validation but lacked a '%s' claim", errors.ProxyAuthenticationRequired, subClaim)
	}

	return email, googleID, nil
}

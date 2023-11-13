package local_user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TODO (Jack): Make this package actually grab and cache the ADC credentials, so we use the user's real email/ID

const (
	emailControlHeader       = "X-SHERLOCK-DEBUG-EMAIL"
	fallbackEmail            = "fallback-fake-sherlock-user@broadinstitute.org"
	googleIDControlHeader    = "X-SHERLOCK-DEBUG-GOOGLE-ID"
	fallbackGoogleID         = "1234fallback1234"
	suitabilityControlHeader = "X-SHERLOCK-DEBUG-SUITABLE"
)

var (
	LocalUserEmail    = fallbackEmail
	LocalUserGoogleID = fallbackGoogleID
	LocalUserSuitable = true
)

func ParseHeader(ctx *gin.Context) (email string, googleID string, err error) {
	if emailHeader := ctx.GetHeader(emailControlHeader); emailHeader != "" {
		LocalUserEmail = emailHeader
	}

	if googleIDHeader := ctx.GetHeader(googleIDControlHeader); googleIDHeader != "" {
		LocalUserGoogleID = googleIDHeader
	}

	if suitabilityHeader := ctx.GetHeader(suitabilityControlHeader); suitabilityHeader == "" {
		LocalUserSuitable = true
	} else if LocalUserSuitable, err = strconv.ParseBool(suitabilityHeader); err != nil {
		return "", "", fmt.Errorf("failed to parse boolean from %s header: %w", suitabilityControlHeader, err)
	}

	return LocalUserEmail, LocalUserGoogleID, nil
}

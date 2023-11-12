package local_user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TODO (Jack): Make this package actually grab and cache the ADC credentials, so we use the user's real email/ID

const (
	EmailControlHeader       = "X-SHERLOCK-DEBUG-EMAIL"
	FallbackEmail            = "fallback-fake-sherlock-user@broadinstitute.org"
	GoogleIDControlHeader    = "X-SHERLOCK-DEBUG-GOOGLE-ID"
	FallbackGoogleID         = "1234fallback1234"
	SuitabilityControlHeader = "X-SHERLOCK-DEBUG-SUITABLE"
)

var (
	LocalUserEmail    = FallbackEmail
	LocalUserGoogleID = FallbackGoogleID
	LocalUserSuitable = true
)

func ParseHeader(ctx *gin.Context) (email string, googleID string, err error) {
	if emailHeader := ctx.GetHeader(EmailControlHeader); emailHeader != "" {
		LocalUserEmail = emailHeader
	}

	if googleIDHeader := ctx.GetHeader(GoogleIDControlHeader); googleIDHeader != "" {
		LocalUserGoogleID = googleIDHeader
	}

	if suitabilityHeader := ctx.GetHeader(SuitabilityControlHeader); suitabilityHeader == "" {
		LocalUserSuitable = true
	} else if LocalUserSuitable, err = strconv.ParseBool(suitabilityHeader); err != nil {
		return "", "", fmt.Errorf("failed to parse boolean from %s header: %w", SuitabilityControlHeader, err)
	}

	return LocalUserEmail, LocalUserGoogleID, nil
}

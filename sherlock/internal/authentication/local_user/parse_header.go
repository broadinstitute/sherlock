package local_user

import (
	"github.com/gin-gonic/gin"
	"strings"
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

func ParseHeader(ctx *gin.Context) (email string, googleID string) {
	emailHeader := ctx.GetHeader(EmailControlHeader)
	if emailHeader != "" {
		LocalUserEmail = emailHeader
	}
	googleIDHeader := ctx.GetHeader(GoogleIDControlHeader)
	if googleIDHeader != "" {
		LocalUserGoogleID = googleIDHeader
	}
	suitabilityHeader := ctx.GetHeader(SuitabilityControlHeader)
	if strings.ToLower(suitabilityHeader) == "false" {
		LocalUserSuitable = false
	} else {
		LocalUserSuitable = true
	}
	return LocalUserEmail, LocalUserGoogleID
}

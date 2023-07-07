package test_users

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	SuitabilityControlHeader    = "X-SHERLOCK-TEST-SUITABLE"
	SuitableTestUserEmail       = "suitable-test-email@broadinstitute.org"
	SuitableTestUserGoogleID    = "12341234"
	NonSuitableTestUserEmail    = "non-suitable-test-email@broadinstitute.org"
	NonSuitableTestUserGoogleID = "67896789"
)

func ParseHeader(ctx *gin.Context) (email string, googleID string) {
	header := ctx.GetHeader(SuitabilityControlHeader)
	if strings.ToLower(header) == "false" {
		return NonSuitableTestUserEmail, NonSuitableTestUserGoogleID
	} else {
		return SuitableTestUserEmail, SuitableTestUserGoogleID
	}
}

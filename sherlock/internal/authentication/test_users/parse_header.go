package test_users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	SuitabilityControlHeader    = "X-SHERLOCK-TEST-SUITABLE"
	SuitableTestUserEmail       = "suitable-test-email@broadinstitute.org"
	SuitableTestUserGoogleID    = "12341234"
	NonSuitableTestUserEmail    = "non-suitable-test-email@broadinstitute.org"
	NonSuitableTestUserGoogleID = "67896789"
)

func ParseHeader(ctx *gin.Context) (email string, googleID string, err error) {
	if header := ctx.GetHeader(SuitabilityControlHeader); header == "" {
		return SuitableTestUserEmail, SuitableTestUserGoogleID, nil
	} else if suitable, err := strconv.ParseBool(header); err != nil {
		return "", "", fmt.Errorf("failed to parse boolean from %s header: %w", SuitabilityControlHeader, err)
	} else if suitable {
		return SuitableTestUserEmail, SuitableTestUserGoogleID, nil
	} else {
		return NonSuitableTestUserEmail, NonSuitableTestUserGoogleID, nil
	}
}

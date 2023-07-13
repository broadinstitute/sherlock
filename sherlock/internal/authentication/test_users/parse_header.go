package test_users

import (
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

func ParseHeader(ctx *gin.Context) (email string, googleID string) {
	header := ctx.GetHeader(SuitabilityControlHeader)
	if header == "" {
		return SuitableTestUserEmail, SuitableTestUserGoogleID
	}
	suitable, err := strconv.ParseBool(header)
	if err != nil {
		panic(err)
	}
	if suitable {
		return SuitableTestUserEmail, SuitableTestUserGoogleID
	} else {
		return NonSuitableTestUserEmail, NonSuitableTestUserGoogleID
	}
}

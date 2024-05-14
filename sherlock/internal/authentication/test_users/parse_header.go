package test_users

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	suitableControlHeader = "X-SHERLOCK-TEST-SUITABLE"
)

func MakeHeaderParser(suitableUser, nonSuitableUser models.User) func(ctx *gin.Context) (email string, googleID string, err error) {
	return func(ctx *gin.Context) (email string, googleID string, err error) {
		if header := ctx.GetHeader(suitableControlHeader); header == "" {
			return suitableUser.Email, suitableUser.GoogleID, nil
		} else if suitable, err := strconv.ParseBool(header); err != nil {
			return "", "", fmt.Errorf("failed to parse boolean from %v suitableControlHeader: %w", suitable, err)
		} else if suitable {
			return suitableUser.Email, suitableUser.GoogleID, nil
		} else {
			return nonSuitableUser.Email, nonSuitableUser.GoogleID, nil
		}
	}
}

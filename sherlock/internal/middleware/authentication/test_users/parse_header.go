package test_users

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	superAdminControlHeader = "X-SHERLOCK-TEST-SUPER-ADMIN"
	suitableControlHeader   = "X-SHERLOCK-TEST-SUITABLE"
)

func MakeHeaderParser(superAdminUser, suitableUser, nonSuitableUser models.User) func(ctx *gin.Context) (email string, googleID string, err error) {
	return func(ctx *gin.Context) (email string, googleID string, err error) {
		var shouldBeSuperAdmin, shouldBeSuitable bool

		if superAdminHeader := ctx.GetHeader(superAdminControlHeader); superAdminHeader == "" {
			shouldBeSuperAdmin = false // Default
		} else if shouldBeSuperAdmin, err = strconv.ParseBool(superAdminHeader); err != nil {
			return "", "", fmt.Errorf("failed to parse boolean from %v superAdminControlHeader: %w", superAdminHeader, err)
		}

		if suitableHeader := ctx.GetHeader(suitableControlHeader); suitableHeader == "" {
			shouldBeSuitable = true // Default
		} else if shouldBeSuitable, err = strconv.ParseBool(suitableHeader); err != nil {
			return "", "", fmt.Errorf("failed to parse boolean from %v suitableControlHeader: %w", suitableHeader, err)
		}

		if shouldBeSuperAdmin && shouldBeSuitable {
			return superAdminUser.Email, superAdminUser.GoogleID, nil
		} else if shouldBeSuperAdmin && !shouldBeSuitable {
			return "", "", fmt.Errorf("super admin cannot be non-suitable (this case isn't implemented in the test middleware right now)")
		} else if !shouldBeSuperAdmin && shouldBeSuitable {
			return suitableUser.Email, suitableUser.GoogleID, nil
		} else {
			return nonSuitableUser.Email, nonSuitableUser.GoogleID, nil
		}
	}
}

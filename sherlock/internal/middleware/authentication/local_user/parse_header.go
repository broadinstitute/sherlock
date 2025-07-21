package local_user

import (
	"fmt"
	"strconv"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/self"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	emailControlHeader       = "X-SHERLOCK-DEBUG-EMAIL"
	googleIDControlHeader    = "X-SHERLOCK-DEBUG-GOOGLE-ID"
	suitabilityControlHeader = "X-SHERLOCK-DEBUG-SUITABLE"
)

func MakeHeaderParser(db *gorm.DB) func(ctx *gin.Context) (email string, googleID string, err error) {
	return func(ctx *gin.Context) (email string, googleID string, err error) {
		if emailHeader := ctx.GetHeader(emailControlHeader); emailHeader != "" {
			email = emailHeader
		} else {
			email = self.Email
		}

		if googleIDHeader := ctx.GetHeader(googleIDControlHeader); googleIDHeader != "" {
			googleID = googleIDHeader
		} else {
			googleID = self.GoogleID
		}

		suitable := true
		if suitabilityHeader := ctx.GetHeader(suitabilityControlHeader); suitabilityHeader != "" {
			if suitable, err = strconv.ParseBool(suitabilityHeader); err != nil {
				return "", "", fmt.Errorf("failed to parse boolean from %s header: %w", suitabilityControlHeader, err)
			}
		}
		superUserDB := models.SetCurrentUserForDB(db, models.SelfUser)
		if err = superUserDB.
			Where(&models.Suitability{
				Email: &email,
			}).
			Assign(&models.Suitability{
				Suitable:    &suitable,
				Description: utils.PointerTo("set by local_user package, override via X-SHERLOCK-DEBUG-SUITABLE header"),
			}).
			FirstOrCreate(&models.Suitability{}).Error; err != nil {
			return "", "", fmt.Errorf("failed to set suitability for %s: %w", email, err)
		}
		return
	}
}

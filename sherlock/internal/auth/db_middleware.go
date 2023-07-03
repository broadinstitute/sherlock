package auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const contextDbKey = "SherlockDB"
const gormSettingsUserKey = contextUserKey

func DbProviderMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := MustUseUser(ctx)
		if err != nil {
			return
		}
		ctx.Set(contextDbKey, db.Set(gormSettingsUserKey, user))

		ctx.Next()
	}
}

func MustUseDB(ctx *gin.Context) (*gorm.DB, error) {
	dbValue, exists := ctx.Get(contextDbKey)
	if !exists {
		err := fmt.Errorf("(%s) database middleware not present", errors.InternalServerError)
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return nil, err
	}
	db, ok := dbValue.(*gorm.DB)
	if !ok {
		err := fmt.Errorf("(%s) database middleware misconfigured: represented as %T", errors.InternalServerError, dbValue)
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return nil, err
	}
	return db, nil
}

func GetGormUser(db *gorm.DB) (*auth_models.User, error) {
	userValue, exists := db.Get(gormSettingsUserKey)
	if !exists {
		return nil, fmt.Errorf("(%s) database user authentication misconfigured, user value not present", errors.InternalServerError)
	}
	user, ok := userValue.(*auth_models.User)
	if !ok {
		return nil, fmt.Errorf("(%s) database user authentication misconfigured: represented as %T", errors.InternalServerError, userValue)
	}
	return user, nil
}

package authentication

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const ctxDbFieldName = "SherlockDB"

// DbMiddleware must strictly come after UserMiddleware, its call to MustUseUser will fail otherwise.
func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := MustUseUser(ctx)
		if err != nil {
			return
		}
		dbWithUser := models.SetCurrentUserForDB(db, user)
		ctx.Set(ctxDbFieldName, dbWithUser)
	}
}

func ShouldUseDB(ctx *gin.Context) (*gorm.DB, error) {
	dbValue, exists := ctx.Get(ctxDbFieldName)
	if !exists {
		return nil, fmt.Errorf("(%s) database authentication middleware not present", errors.InternalServerError)
	}
	db, ok := dbValue.(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("(%s) database authentication middleware misconfigured: db represented as %T", errors.InternalServerError, dbValue)
	}
	if db == nil {
		return nil, fmt.Errorf("(%s) database reference was nil", errors.InternalServerError)
	}
	return db, nil
}

func MustUseDB(ctx *gin.Context) (*gorm.DB, error) {
	db, err := ShouldUseDB(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
	}
	return db, nil
}

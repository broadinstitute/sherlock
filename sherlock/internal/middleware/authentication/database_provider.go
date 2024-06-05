package authentication

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const ctxDbFieldName = "SherlockDB"

func setDatabaseWithUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := MustUseUser(ctx)
		if err != nil {
			return
		}
		dbWithUser := models.SetCurrentUserForDB(db, user)
		ctx.Set(ctxDbFieldName, dbWithUser)
	}
}

// ShouldUseDB returns a non-nil *gorm.DB with the calling user accessible via models.GetCurrentUserForDB, or an error
// if that isn't possible.
func ShouldUseDB(ctx *gin.Context) (*gorm.DB, error) {
	dbValue, exists := ctx.Get(ctxDbFieldName)
	if !exists {
		return nil, fmt.Errorf("(%s) database reference not present; database authentication middleware likely not present", errors.InternalServerError)
	}
	db, ok := dbValue.(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("(%s) database authentication middleware likely misconfigured: represented as %T", errors.InternalServerError, dbValue)
	}
	if db == nil {
		return nil, fmt.Errorf("(%s) database authentication middleware likely misconfigured: database reference was nil", errors.InternalServerError)
	}
	return db, nil
}

// MustUseDB is like ShouldUseDB except it calls errors.AbortRequest if there was an error so the caller doesn't have to.
func MustUseDB(ctx *gin.Context) (*gorm.DB, error) {
	db, err := ShouldUseDB(ctx)
	if err != nil {
		errors.AbortRequest(ctx, err)
	}
	return db, nil
}

package login

import (
	"database/sql"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/oidc_models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// LoginGet is meant to handle redirects to /login?id=... from the OIDC subsystem,
// read the IAP info from the request, "log the user in", and redirect back to the
// OIDC subsystem.
//
// This isn't in Swagger as it's not an API really -- it's meant to play nice with
// browsers.
func LoginGet(ctx *gin.Context) {
	user, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}

	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	authRequestID := ctx.Query("id")
	if authRequestID == "" {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) no auth request ID passed", errors.BadRequest))
		return
	}

	parsedAuthRequestID, err := uuid.Parse(authRequestID)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) invalid auth request ID passed", errors.BadRequest))
		return
	}

	err = db.Where(&oidc_models.AuthRequest{
		ID: parsedAuthRequestID,
	}).Updates(&oidc_models.AuthRequest{
		UserID: &user.ID,
		DoneAt: sql.NullTime{Time: time.Now(), Valid: true},
	}).Error
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("could not update auth request: %w", err))
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprintf("/oidc/callback?id=%s", authRequestID))
}

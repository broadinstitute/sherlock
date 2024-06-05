package authentication

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

const (
	ctxUserFieldName         = "SherlockUser"
	ctxGithubClaimsFieldName = "GithubActionsClaims"
)

// setUserWhoConnected returns a gin.HandlerFunc that finds or inserts the models.User of the connecting client.
// It sets ctxUserFieldName.
//
// When configured with a parser like test_users.ParseHeader, it essentially creates the test users out of thin air
// as requests are made.
// When configured with a parser like iap.ParseHeader, it enforces that Sherlock be running behind Google's
// Identity-Aware Proxy.
func setUserWhoConnected(db *gorm.DB, headerParser func(ctx *gin.Context) (email string, googleID string, err error), method authentication_method.Method) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if email, googleID, err := headerParser(ctx); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("failed to parse the connecting user from headers: %w", err))
			return
		} else if err = db.Where(&models.User{Email: email, GoogleID: googleID}).FirstOrCreate(&user).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("failed to find or insert the connecting user in database: %w", err))
			return
		} else if err = db.Scopes(models.ReadUserScope).First(&user, user.ID).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("failed to read the connecting user's associations from database: %w", err))
			return
		} else {
			user.AuthenticationMethod = method
			ctx.Set(ctxUserFieldName, &user)
		}
	}
}

// setGithubClaimsAndEscalateUser returns a gin.HandlerFunc that attempts to parse gha_oidc.Claims from the request if
// it came from GitHub Actions. If those claims can be tied to a models.User it will use that as the primary user, too.
// It possibly sets ctxGithubClaimsFieldName, and if so, possibly ctxUserFieldName.
//
// gha_oidc.ParseHeader validates that the OIDC JWT from GitHub is authentic, so any information in gha_oidc.Claims will
// be true.
// Meanwhile, each models.User will only have a GitHub identity set based on that user sending Sherlock an access token
// for that GitHub identity.
// This means we can know for certain the models.User responsible for the call to Sherlock, so we can safely substitute
// that information over top of the (presumably service account) models.User used to access IAP and thus parsed by
// setUserWhoConnected.
func setGithubClaimsAndEscalateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if claims, err := gha_oidc.ParseHeader(ctx); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("failed to parse GitHub Actions claims: %w", err))
			return
		} else if claims != nil {
			ctx.Set(ctxGithubClaimsFieldName, claims)

			if claims.ActorID != "" {
				var matchingGithubUsers []models.User
				if err = db.Where(&models.User{GithubID: &claims.ActorID}).Limit(1).Find(&matchingGithubUsers).Error; err != nil {
					errors.AbortRequest(ctx, fmt.Errorf("failed to query for users matching GitHub Actions claims: %w", err))
					return
				} else if len(matchingGithubUsers) == 1 {
					var user models.User
					if err = db.Scopes(models.ReadUserScope).First(&user, matchingGithubUsers[0].ID).Error; err != nil {
						errors.AbortRequest(ctx, fmt.Errorf("failed to read the GitHub Actions user's associations from database: %w", err))
						return
					} else if oldUser, err := ShouldUseUser(ctx); err != nil {
						log.Warn().Err(err).Msg("AUTH | was unable to read old user to escalate user based on GitHub claims")
					} else {
						user.AuthenticationMethod = authentication_method.GHA
						user.Via = oldUser
						ctx.Set(ctxUserFieldName, &user)
						log.Debug().Str("workflow", claims.WorkflowURL()).Msgf("AUTH | recognized %s connecting through GitHub Actions via %s", user.Email, oldUser.Email)
					}
				} else {
					log.Debug().Str("workflow", claims.WorkflowURL()).Msgf("AUTH | ignoring unknown user '%s' referenced by GitHub Actions claims", claims.Actor)
				}
			} else {
				log.Debug().Str("workflow", claims.WorkflowURL()).Msgf("AUTH | ignoring GitHub Actions claims for user escalation since actor ID was empty")
			}
		}
	}
}

// ShouldUseUser returns a non-nil *models.User who made the request, or an error if that isn't possible.
func ShouldUseUser(ctx *gin.Context) (*models.User, error) {
	userValue, exists := ctx.Get(ctxUserFieldName)
	if !exists {
		return nil, fmt.Errorf("(%s) user not present; authentication middleware likely not present", errors.InternalServerError)
	}
	user, ok := userValue.(*models.User)
	if !ok {
		return nil, fmt.Errorf("(%s) user authentication middleware likely misconfigured: represented as %T", errors.InternalServerError, userValue)
	}
	if user == nil {
		return nil, fmt.Errorf("(%s) user authentication middleware likely misconfigured: user was nil", errors.InternalServerError)
	}
	return user, nil
}

// MustUseUser is like ShouldUseUser except it calls errors.AbortRequest if there was an error so the
// caller doesn't have to.
func MustUseUser(ctx *gin.Context) (*models.User, error) {
	user, err := ShouldUseUser(ctx)
	if err != nil {
		errors.AbortRequest(ctx, err)
	}
	return user, err
}

// ShouldUseGithubClaims returns non-nil *gha_oidc.Claims associated with the request, or an error if that isn't
// possible.
//
// Note that the Actor/ActorID fields of the gha_oidc.Claims must not be correlated to a models.User. That correlation
// is the responsibility of the authentication package; use ShouldUseUser or MustUseUser to access the models.User
// associated with the request.
func ShouldUseGithubClaims(ctx *gin.Context) (*gha_oidc_claims.Claims, error) {
	claimsValue, exists := ctx.Get(ctxGithubClaimsFieldName)
	if !exists {
		return nil, fmt.Errorf("(%s) GitHub OIDC claims were not present", errors.InternalServerError)
	}
	claims, ok := claimsValue.(*gha_oidc_claims.Claims)
	if !ok {
		return nil, fmt.Errorf("(%s) GitHub OIDC claims middleware may be misconfigured: represented as %T", errors.InternalServerError, claimsValue)
	}
	if claims == nil {
		return nil, fmt.Errorf("(%s) GitHub OIDC claims middleware may be misconfigured: claims were nil", errors.InternalServerError)
	}
	return claims, nil
}

package authentication

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/iap"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/local_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

const ctxUserFieldName = "SherlockUser"

func UserMiddleware(db *gorm.DB) gin.HandlerFunc {
	if config.Config.String("mode") == "debug" {
		if config.Config.Bool("auth.createTestUsersInMiddleware") {
			log.Info().Msgf("AUTH | using test authentication; will create test users from middleware")
			return fakeUserMiddleware(db, test_users.ParseHeader, authentication_method.TEST)
		} else {
			log.Info().Msgf("AUTH | using local authentication; will create a local user from middleware")
			return fakeUserMiddleware(db, local_user.ParseHeader, authentication_method.LOCAL)
		}
	} else {
		return realUserMiddleware(db)
	}
}

func realUserMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, googleID, err := iap.ParseHeader(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}

		var iapUser models.User
		if err = db.Where(&models.User{Email: email, GoogleID: googleID}).FirstOrCreate(&iapUser).Error; err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		iapUser.AuthenticationMethod = authentication_method.IAP

		headerPresent, githubUsername, githubID, err := gha_oidc.ParseHeader(ctx)
		if headerPresent {
			if err != nil {
				ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
				return
			}

			var ghaUser models.User
			if err = db.Where(&models.User{GithubID: &githubID}).First(&ghaUser).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					log.Info().Msgf("AUTH | ignored GHA OIDC JWT for unknown github user %s, still using IAP JWT user %s", githubUsername, iapUser.Email)
					ctx.Set(ctxUserFieldName, &iapUser)
				} else {
					ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
					return
				}
			} else {
				ghaUser.AuthenticationMethod = authentication_method.GHA
				ghaUser.Via = &iapUser
				log.Info().Msgf("AUTH | substituted GHA OIDC JWT user %s (github %s) over IAP JWT user %s", ghaUser.Email, githubUsername, iapUser.Email)
				ctx.Set(ctxUserFieldName, &ghaUser)
			}
		} else {
			ctx.Set(ctxUserFieldName, &iapUser)
		}
	}
}

// fakeUserMiddleware has the same out-facing functionality as realUserMiddleware but it basically blindly uses the given header parser method.
// This means that "nonsense" (insecure) header parsers can be passed, to effectively create Sherlock users out of thin air.
func fakeUserMiddleware(db *gorm.DB, headerParser func(ctx *gin.Context) (email string, googleID string), method authentication_method.Method) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, googleID := headerParser(ctx)
		var fakeUser models.User
		if err := db.Where(&models.User{Email: email, GoogleID: googleID}).FirstOrCreate(&fakeUser).Error; err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		fakeUser.AuthenticationMethod = method
		ctx.Set(ctxUserFieldName, &fakeUser)
	}
}

func ShouldUseUser(ctx *gin.Context) (*models.User, error) {
	userValue, exists := ctx.Get(ctxUserFieldName)
	if !exists {
		return nil, fmt.Errorf("(%s) user authentication middleware not present", errors.InternalServerError)
	}
	user, ok := userValue.(*models.User)
	if !ok {
		return nil, fmt.Errorf("(%s) user authentication middleware misconfigured: represented as %T", errors.InternalServerError, userValue)
	}
	if user == nil {
		return nil, fmt.Errorf("(%s) user was nil", errors.InternalServerError)
	}
	return user, nil
}

func MustUseUser(ctx *gin.Context) (*models.User, error) {
	user, err := ShouldUseUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
	}
	return user, err
}

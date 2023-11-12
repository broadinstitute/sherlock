package authentication

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/iap"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/local_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Middleware returns an ordered list of middleware to authenticate requests, enabling request handlers to use functions
// like ShouldUseDB, ShouldUseUser, and ShouldUseGithubClaims.
func Middleware(db *gorm.DB) gin.HandlersChain {
	if config.Config.String("mode") == "debug" {
		if config.Config.Bool("auth.createTestUsersInMiddleware") {
			log.Info().Msgf("AUTH | using test authentication; will create test users from middleware")
			return gin.HandlersChain{
				setUserWhoConnected(db, test_users.ParseHeader, authentication_method.TEST),
				setDatabaseWithUser(db),
			}
		} else {
			log.Info().Msgf("AUTH | using local authentication; will create a local user from middleware")
			return gin.HandlersChain{
				setUserWhoConnected(db, local_user.ParseHeader, authentication_method.LOCAL),
				setDatabaseWithUser(db),
			}
		}
	} else {
		return gin.HandlersChain{
			setUserWhoConnected(db, iap.ParseHeader, authentication_method.IAP),
			setGithubClaimsAndEscalateUser(db),
			setDatabaseWithUser(db),
		}
	}
}

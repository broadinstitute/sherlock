package authentication

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/authentication_method"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/iap"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/local_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Middleware returns an ordered list of middleware to authenticate requests, enabling request handlers to use functions
// like ShouldUseDB, ShouldUseUser, and ShouldUseGithubClaims.
func Middleware(db *gorm.DB) gin.HandlersChain {
	var userWhoConnectedMiddleware gin.HandlerFunc
	if config.Config.String("mode") == "debug" {
		log.Info().Msgf("AUTH | using local authentication; will control your local user from middleware")
		userWhoConnectedMiddleware = setUserWhoConnected(db, local_user.MakeHeaderParser(db), authentication_method.LOCAL)
	} else {
		userWhoConnectedMiddleware = setUserWhoConnected(db, iap.ParseHeader, authentication_method.IAP)
	}
	return gin.HandlersChain{
		userWhoConnectedMiddleware,
		setGithubClaimsAndEscalateUser(db),
		setDatabaseWithUser(db),
	}
}

// TestMiddleware is like Middleware but requires models.TestData (so it can only be used from tests)
// and references models.TestData's User_Suitable and User_NonSuitable instances of models.User.
func TestMiddleware(db *gorm.DB, td models.TestData) gin.HandlersChain {
	return gin.HandlersChain{
		setUserWhoConnected(db, test_users.MakeHeaderParser(td.User_SuperAdmin(), td.User_Suitable(), td.User_NonSuitable()), authentication_method.TEST),
		setGithubClaimsAndEscalateUser(db),
		setDatabaseWithUser(db),
	}
}

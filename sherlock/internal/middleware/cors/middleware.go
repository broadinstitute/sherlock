package cors

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors does the CORS thing. IAP requires its cookie to be sent cross-origin, so origins must be
// specifically enumerated for cross-origin AJAX to work.
//
// If no origins are specified, this middleware will allow all origins since that's at least
// useful for local development where credentials aren't in play (browsers require origins to be
// enumerated to send credentials, so operation behind IAP should enumerate origins).
func Cors() gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowHeaders = append(c.AllowHeaders,
		// "X-Requested-With" being set to "XMLHttpRequest" helps IAP understand the
		// AJAX nature of the request so it can return a 401 rather than a 302.
		// https://cloud.google.com/iap/docs/sessions-howto#understanding_the_response
		"X-Requested-With")
	if origins := config.Config.Strings("cors.allowOrigins"); len(origins) > 0 {
		c.AllowCredentials = true
		c.AllowOrigins = origins
	} else {
		c.AllowCredentials = false
		c.AllowAllOrigins = true
	}
	return cors.New(c)
}

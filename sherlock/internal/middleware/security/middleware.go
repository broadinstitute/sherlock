package security

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	c := secure.Config{
		// TLS is terminated at the proxy and/or load balancer, so we
		// don't enable any TLS-related behavior. Some of the rest of
		// this is just a freebie so might as well.
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
		ContentSecurityPolicy: "default-src 'self'; " +
			"img-src 'self' data:; " +
			"style-src 'self' 'unsafe-inline'; " +
			"frame-ancestors 'none'; ",
		IENoOpen:       true,
		ReferrerPolicy: "strict-origin-when-cross-origin",
	}
	return secure.New(c)
}

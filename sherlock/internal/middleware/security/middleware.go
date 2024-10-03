package security

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"strings"
)

var styleAttributeValuesToAllow = []string{
	"position:absolute;width:0;height:0",
}

func Security() gin.HandlerFunc {
	styleAttributeHashes := make([]string, len(styleAttributeValuesToAllow))
	for i, value := range styleAttributeValuesToAllow {
		hash := sha256.Sum256([]byte(value))
		styleAttributeHashes[i] = fmt.Sprintf("'sha256-%s'", base64.StdEncoding.EncodeToString(hash[:]))
	}
	styleDirective := fmt.Sprintf("style-src 'self' 'unsafe-hashes' %s; ", strings.Join(styleAttributeHashes, " "))

	c := secure.Config{
		// TLS is terminated at the proxy and/or load balancer, so we
		// don't enable any TLS-related behavior. Some of the rest of
		// this is just a freebie so might as well.
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
		ContentSecurityPolicy: "default-src 'self'; " +
			"img-src 'self' data:; " +
			styleDirective +
			"frame-ancestors 'none'; ",
		IENoOpen:       true,
		ReferrerPolicy: "strict-origin-when-cross-origin",
	}
	return secure.New(c)
}

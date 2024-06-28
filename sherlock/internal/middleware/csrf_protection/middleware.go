package csrf_protection

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strings"
)

// CsrfProtection is a layer of defense against Cross-Site-Request-Forgery attacks.
//
// It checks three things:
//  1. The content type of the request is application/json or not passed at all. HTML forms are a common
//     vector for CSRF attacks, and they won't ever be application/json without JavaScript being involved
//     in sending the request.
//  2. The Origin header is in the list of allowed origins or not passed at all. Assuming JavaScript is
//     involved in sending the request, we're in the land of JavaScript's same-origin-policy, so we can
//     filter based on this header. An attacker won't be able to convince a modern browser *not* to send
//     this header cross-origin.
//  3. We do the same check for the Referer header as we do for the Origin header too. It's less reliable
//     than the Origin header but it has older support.
//
// This isn't perfect -- it's not bulletproof like a double-submit cookie can be -- but for an API
// it's fairly strong.
func CsrfProtection() gin.HandlerFunc {
	origins := config.Config.Strings("origins")
	return func(ctx *gin.Context) {
		if ct := ctx.ContentType(); ct != "" && ct != "application/json" {
			log.Warn().Str("content-type", ct).Msgf("unsupported content type %s observed, rejecting request for CSRF protection", ct)
			errors.AbortRequest(ctx, fmt.Errorf("(%s) unsupported content type; see logs for more details", errors.BadRequest))
			return
		} else if len(origins) > 0 {
			if origin := ctx.GetHeader("Origin"); origin != "" && !utils.Contains(origins, origin) {
				log.Warn().Str("origin", origin).Msgf("origin %s not allowed, rejecting request for CSRF protection", origin)
				errors.AbortRequest(ctx, fmt.Errorf("(%s) origin not allowed; see logs for more details", errors.Forbidden))
				return
			}
			if referer := ctx.GetHeader("Referer"); referer != "" && !refererHasOriginAsPrefix(referer, origins) {
				log.Warn().Str("referer", referer).Msgf("referer %s not allowed, rejecting request for CSRF protection", referer)
				errors.AbortRequest(ctx, fmt.Errorf("(%s) referer not allowed; see logs for more details", errors.Forbidden))
				return
			}
		}
	}
}

func refererHasOriginAsPrefix(referer string, origins []string) bool {
	for _, origin := range origins {
		if strings.HasPrefix(referer, origin) {
			return true
		}
	}
	return false
}

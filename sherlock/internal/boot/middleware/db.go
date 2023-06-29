package middleware

import "github.com/broadinstitute/sherlock/sherlock/internal/auth"

// DB is just a re-export of auth.DbProviderMiddleware so that all the middleware come from this package.
// Pedantic, yes, but we *kinda* do the same thing for the auth middleware too.
var DB = auth.DbProviderMiddleware

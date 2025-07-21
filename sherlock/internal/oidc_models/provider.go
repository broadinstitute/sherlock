package oidc_models

import (
	"encoding/hex"
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/zitadel/oidc/v3/pkg/op"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var Provider op.OpenIDProvider

func initProvider(db *gorm.DB) error {
	key, err := hex.DecodeString(config.Config.String("oidc.encryptionKeyHex"))
	if err != nil {
		return fmt.Errorf("could not decode oidc.encryptionKeyHex: %w", err)
	} else if len(key) != 32 {
		return fmt.Errorf("oidc.encryptionKeyHex must be 32 bytes long; got %d", len(key))
	}

	storage := &storageImpl{db: db}
	conf := &op.Config{
		CryptoKey:                ([32]byte)(key),
		DefaultLogoutRedirectURI: "/static/logged-out.html",
		CodeMethodS256:           true, // Enable PKCE and S256 code challenge method
		GrantTypeRefreshToken:    true, // Allow refresh token grant user
		SupportedUILocales:       []language.Tag{language.AmericanEnglish},
		SupportedClaims:          append(op.DefaultSupportedClaims, groupsClaim), // Technically more than we provide but better a superset than subset
	}
	options := []op.Option{
		op.WithCORSOptions(nil),
		// If we expand our usage of the OIDC subsystem such that users need to connect to other endpoints,
		// we may need to add custom URLs for those endpoints here.
		op.WithCustomAuthEndpoint(op.NewEndpointWithURL("/authorize", config.Config.String("oidc.publicIssuerUrl")+"/authorize")),
	}
	if config.Config.String("mode") == "debug" {
		options = append(options, op.WithAllowInsecure())
	}

	Provider, err = op.NewProvider(conf, storage, op.StaticIssuer(config.Config.String("oidc.issuerUrl")), options...)
	return err
}

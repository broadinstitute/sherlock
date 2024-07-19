package oidc_models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/zitadel/oidc/v3/pkg/op"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var Provider op.OpenIDProvider

func initProvider(db *gorm.DB) error {
	key := config.Config.Bytes("oidc.encryptionKey")
	if len(key) != 32 {
		return fmt.Errorf("oidc.encryptionKey must be 32 bytes long; got %d", len(key))
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
	var options []op.Option
	if config.Config.String("mode") == "debug" {
		options = append(options, op.WithAllowInsecure())
	}

	var err error
	Provider, err = op.NewProvider(conf, storage, op.StaticIssuer(config.Config.String("oidc.issuerUrl")), options...)
	return err
}

package oidc_models

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"time"

	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

type TestClientHelper struct{}

func (h TestClientHelper) GeneratedClientRedirectURI() string {
	return "http://localhost:8080/test/fake/redirect"
}

func (h TestClientHelper) GeneratedClientPostLogoutRedirectURI() string {
	return "http://localhost:8080/test/fake/postlogout"
}

func (h TestClientHelper) GenerateClient(db *gorm.DB) (clientID string, clientSecret string, err error) {
	clientIDBytes := make([]byte, 16)
	_, err = rand.Read(clientIDBytes)
	if err != nil {
		return "", "", err
	}
	clientID = hex.EncodeToString(clientIDBytes)

	clientSecretBytes := make([]byte, 32)
	_, err = rand.Read(clientSecretBytes)
	if err != nil {
		return "", "", err
	}
	clientSecret = hex.EncodeToString(clientSecretBytes)

	clientSecretSalt := make([]byte, 32)
	_, err = rand.Read(clientSecretSalt)
	if err != nil {
		return "", "", err
	}

	testClientHashIterations := 1_000 // Low low low value!!! Needs to be 210_000+ in production, this just makes tests run faster
	clientSecretHash := pbkdf2.Key([]byte(clientSecret), clientSecretSalt, testClientHashIterations, 32, sha512.New)

	client := Client{
		ID:                           clientID,
		ClientSecretHash:             clientSecretHash,
		ClientSecretSalt:             clientSecretSalt,
		ClientSecretIterations:       testClientHashIterations,
		ClientRedirectURIs:           oidc.SpaceDelimitedArray{h.GeneratedClientRedirectURI()},
		ClientPostLogoutRedirectURIs: oidc.SpaceDelimitedArray{h.GeneratedClientPostLogoutRedirectURI()},
		ClientApplicationType:        op.ApplicationTypeWeb,
		ClientAuthMethod:             oidc.AuthMethodBasic,
		ClientIDTokenLifetime:        (15 * time.Minute).Nanoseconds(),
		ClientDevMode:                true,
		ClientClockSkew:              (15 * time.Second).Nanoseconds(),
	}

	err = db.Create(&client).Error
	return clientID, clientSecret, err
}

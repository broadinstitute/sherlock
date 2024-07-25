package oidc_models

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"gorm.io/gorm"
	"time"
)

var (
	kmsKey    string
	kmsClient *kms.KeyManagementClient
)

func Init(ctx context.Context, db *gorm.DB) error {
	if config.Config.Bool("oidc.signingKeyEncryptionKMSEnable") {
		kmsKey = config.Config.String("oidc.signingKeyEncryptionKMSKeyName")
		var err error
		kmsClient, err = kms.NewKeyManagementClient(ctx)
		if err != nil {
			return fmt.Errorf("error creating KMS client: %w", err)
		}
		response, err := kmsClient.GetCryptoKey(ctx, &kmspb.GetCryptoKeyRequest{
			Name: kmsKey,
		})
		if err != nil {
			return fmt.Errorf("error getting KMS key '%s': %w", kmsKey, err)
		} else if response.Purpose != kmspb.CryptoKey_ENCRYPT_DECRYPT {
			return fmt.Errorf("KMS key '%s' is not an encrypt/decrypt key", kmsKey)
		}
	} else if config.Config.String("mode") != "debug" {
		return fmt.Errorf("oidc.signingKeyEncryptionKMSEnable is false, but mode is not debug")
	}

	if err := rotateSigningKeys(ctx, db); err != nil {
		return fmt.Errorf("error rotating oidc signing keys: %w", err)
	}

	return initProvider(db)
}

func KeepSigningKeysRotated(ctx context.Context, db *gorm.DB) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(15 * time.Minute):
			err := rotateSigningKeys(ctx, db)
			if err != nil {
				slack.ReportError(ctx, "error rotating oidc signing keys", err)
			}
		}
	}
}

func KeepExpiringRefreshTokens(ctx context.Context, db *gorm.DB) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			err := expireRefreshTokens(db)
			if err != nil {
				slack.ReportError(ctx, "error expiring refresh tokens", err)
			}
		}
	}
}

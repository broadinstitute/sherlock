package oidc_models

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"time"

	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/go-jose/go-jose/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/zitadel/oidc/v3/pkg/op"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SigningKey struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	CreatedAt  time.Time
	PublicKey  []byte
	PrivateKey []byte
}

func rotateSigningKeys(ctx context.Context, db *gorm.DB) error {
	var validKeys []SigningKey
	err := db.
		Where("created_at > ?", time.Now().
			Add(-config.Config.MustDuration("oidc.signingKeyPrimaryDuration"))).
		Omit(clause.Associations).
		Find(&validKeys).Error
	if err != nil {
		return fmt.Errorf("error loading valid signing keys: %w", err)
	}
	if len(validKeys) == 0 {
		newKey, err := saveNewSigningKey(ctx, db)
		if err != nil {
			return fmt.Errorf("error generating new signing key: %w", err)
		} else {
			log.Info().Msgf("OIDC | generated new signing key with ID %s", newKey.ID)
		}
	}

	var expiredKeys []SigningKey
	err = db.
		Where("created_at <= ?", time.Now().
			Add(-(config.Config.MustDuration("oidc.signingKeyPrimaryDuration"))).
			Add(-config.Config.MustDuration("oidc.signingKeyPostRotationDuration"))).
		Omit(clause.Associations).
		Find(&expiredKeys).Error
	if err != nil {
		return fmt.Errorf("error loading expired signing keys: %w", err)
	}
	for _, key := range expiredKeys {
		err = db.Omit(clause.Associations).Delete(&key).Error
		if err != nil {
			return fmt.Errorf("error deleting expired signing key: %w", err)
		}
		log.Info().Msgf("OIDC | deleted expired signing key with ID %s", key.ID)
	}
	return nil
}

func saveNewSigningKey(ctx context.Context, db *gorm.DB) (*SigningKey, error) {
	keyModel := &SigningKey{
		ID: uuid.New(),
	}

	// Generate new private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("error generating private key: %w", err)
	}

	// Store public key plaintext in database for easy access
	keyModel.PublicKey = elliptic.MarshalCompressed(privateKey.PublicKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y) //nolint:staticcheck // QF1008

	// Store private key, encrypting with KMS if configured
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("error marshaling private key: %w", err)
	}
	if kmsKey != "" {
		// Store encrypted
		response, err := kmsClient.Encrypt(ctx, &kmspb.EncryptRequest{
			Name:      kmsKey,
			Plaintext: privateKeyBytes,
		})
		if err != nil {
			return nil, fmt.Errorf("error encrypting private key with KMS: %w", err)
		}
		keyModel.PrivateKey = response.Ciphertext
	} else {
		// Store plaintext
		keyModel.PrivateKey = privateKeyBytes
	}

	// Save database row
	err = db.Create(keyModel).Error
	if err != nil {
		return nil, fmt.Errorf("error saving generated signing key to database: %w", err)
	}
	return keyModel, nil
}

var _ op.Key = &publicSigningKey{}

type publicSigningKey struct {
	SigningKey
}

func (p *publicSigningKey) ID() string {
	return p.SigningKey.ID.String()
}

func (p *publicSigningKey) Algorithm() jose.SignatureAlgorithm {
	return jose.ES512
}

func (p *publicSigningKey) Use() string {
	return "sig"
}

func (p *publicSigningKey) Key() any {
	x, y := elliptic.UnmarshalCompressed(elliptic.P521(), p.PublicKey)
	return &ecdsa.PublicKey{
		Curve: elliptic.P521(),
		X:     x,
		Y:     y,
	}
}

var _ op.SigningKey = &decryptedPrivateSigningKey{}

type decryptedPrivateSigningKey struct {
	SigningKey
	unencryptedPrivateKey *ecdsa.PrivateKey
}

func (p *decryptedPrivateSigningKey) ID() string {
	return p.SigningKey.ID.String()
}

func (p *decryptedPrivateSigningKey) SignatureAlgorithm() jose.SignatureAlgorithm {
	return jose.ES512
}

func (p *decryptedPrivateSigningKey) Key() any {
	return p.unencryptedPrivateKey
}

func decryptPrivateSigningKey(key *SigningKey) (*decryptedPrivateSigningKey, error) {
	bytesToUnmarshall := key.PrivateKey
	if kmsKey != "" {
		response, err := kmsClient.Decrypt(context.Background(), &kmspb.DecryptRequest{
			Name:       kmsKey,
			Ciphertext: bytesToUnmarshall,
		})
		if err != nil {
			return nil, fmt.Errorf("error decrypting private signing key with KMS: %w", err)
		}
		bytesToUnmarshall = response.Plaintext
	}
	privateKey, err := x509.ParseECPrivateKey(bytesToUnmarshall)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling private signing key: %w", err)
	}
	return &decryptedPrivateSigningKey{
		SigningKey:            *key,
		unencryptedPrivateKey: privateKey,
	}, nil
}

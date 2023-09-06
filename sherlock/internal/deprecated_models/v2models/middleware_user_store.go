package v2models

import (
	go_errors "errors"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/auth_models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func NewMiddlewareUserStore(db *gorm.DB) *MiddlewareUserStore {
	return &MiddlewareUserStore{
		modelStore: &ModelStore[User]{db: db, internal: InternalUserStore},
	}
}

type MiddlewareUserStore struct {
	modelStore *ModelStore[User]
}

func (s *MiddlewareUserStore) GetOrCreateUser(email, googleID string) (User, error) {
	query := User{
		StoredControlledUserFields: auth_models.StoredControlledUserFields{
			Email: email,
		},
	}
	existing, err := s.modelStore.internal.GetIfExists(s.modelStore.db, query)
	if err != nil {
		return User{}, err
	} else if existing != nil {
		if existing.GoogleID != googleID {
			return User{}, fmt.Errorf("(%s) incoming google ID '%s' mismatched with stored '%s', please contact DevOps", errors.BadRequest, googleID, existing.GoogleID)
		} else {
			return *existing, nil
		}
	} else {
		log.Info().Msgf("AUTH | automatically adding new user %s (ID %s)", email, googleID)
		user, _, err := s.modelStore.internal.Create(s.modelStore.db, User{
			StoredControlledUserFields: auth_models.StoredControlledUserFields{
				Email:    email,
				GoogleID: googleID,
			},
		}, nil)
		if err != nil {
			var pgErr *pgconn.PgError
			if go_errors.As(err, &pgErr) && pgErr != nil && pgerrcode.UniqueViolation == pgErr.Code {
				existing, err = s.modelStore.internal.GetIfExists(s.modelStore.db, query)
				if err != nil {
					return User{}, err
				} else if existing != nil {
					if existing.GoogleID != googleID {
						return User{}, fmt.Errorf("(%s) incoming google ID '%s' mismatched with stored '%s', please contact DevOps", errors.BadRequest, googleID, existing.GoogleID)
					} else {
						return *existing, nil
					}
				}
			}
			return User{}, err
		}
		return user, nil
	}
}

func (s *MiddlewareUserStore) GetGithubUserIfExists(githubID string) (*User, error) {
	query := User{
		StoredControlledUserFields: auth_models.StoredControlledUserFields{
			GithubID: &githubID,
		},
	}
	return s.modelStore.internal.GetIfExists(s.modelStore.db, query)
}

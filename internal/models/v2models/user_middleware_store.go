package v2models

import (
	go_errors "errors"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

func NewUserMiddlewareStore(db *gorm.DB) *UserMiddlewareStore {
	return &UserMiddlewareStore{
		ModelStore: &ModelStore[User]{db: db, internalModelStore: userStore},
	}
}

type UserMiddlewareStore struct {
	*ModelStore[User]
}

func (s *UserMiddlewareStore) GetOrCreateUser(email, googleID string) (User, error) {
	query := User{
		StoredUserFields: auth_models.StoredUserFields{
			Email: email,
		},
	}
	existing, err := s.getIfExists(s.db, query)
	if err != nil {
		return User{}, err
	} else if existing != nil {
		if existing.GoogleID != googleID {
			return User{}, fmt.Errorf("(%s) incoming google ID '%s' mismatched with stored '%s', please contact DevOps", errors.BadRequest, googleID, existing.GoogleID)
		} else {
			return *existing, nil
		}
	} else {
		user, _, err := s.create(s.db, User{
			StoredUserFields: auth_models.StoredUserFields{
				Email:    email,
				GoogleID: googleID,
			},
		}, nil)
		if err != nil {
			var pgErr *pgconn.PgError
			if go_errors.As(err, &pgErr) && pgErr != nil && pgerrcode.UniqueViolation == pgErr.Code {
				existing, err = s.getIfExists(s.db, query)
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

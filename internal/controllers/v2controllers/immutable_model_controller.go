package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/creasty/defaults"
	"time"
)

type ReadableBaseType struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Readable interface {
}

type Creatable[R Readable] interface {
	toReadable() R
}

type ImmutableModelController[M v2models.Model, R Readable, C Creatable[R]] struct {
	primaryStore    v2models.Store[M]
	allStores       v2models.StoreSet
	modelToReadable func(model M) *R
	readableToModel func(readable R, stores v2models.StoreSet) (M, error)

	// setDynamicDefaults is an optional function to set context-dynamic defaults during creation.
	//
	// Normally, github.com/creasty/defaults is run to respect the same `default` struct tags that
	// github.com/swaggo/swag understands for documentation. creasty/defaults also respects any no-argument, no-output
	// "SetDefaults" pointer-methods on any type it comes across, allowing for some basic dynamic behavior.
	//
	// However, sometimes this isn't enough. When setDynamicDefaults is present, it will get run first, before
	// creasty/defaults. It needn't worry about handling those `default` struct tags--it can be provided to allow a type
	// to set defaults before creation that are dynamic based on other existing data or the calling user. When this
	// behavior is taken advantage of, the type should have comments either on the struct or the field itself--those
	// comments will be picked up by swaggo/swag.
	//
	// Environment is an example type that takes advantage of this capability.
	setDynamicDefaults func(readable *R, stores v2models.StoreSet, user auth.User) error
}

func (c ImmutableModelController[M, R, C]) Create(creatable C, user auth.User) (R, error) {
	readable := creatable.toReadable()
	if c.setDynamicDefaults != nil {
		if err := c.setDynamicDefaults(&readable, c.allStores, user); err != nil {
			return readable, fmt.Errorf("(%s) error setting dynamic default values for %T: %v", errors.InternalServerError, readable, err)
		}
	}
	if err := defaults.Set(&readable); err != nil {
		return readable, fmt.Errorf("(%s) error setting static default values for %T: %v", errors.InternalServerError, readable, err)
	}
	model, err := c.readableToModel(readable, c.allStores)
	if err != nil {
		return readable, err
	}
	result, err := c.primaryStore.Create(model, user.Suitable)
	return *c.modelToReadable(result), err
}

func (c ImmutableModelController[M, R, C]) ListAllMatching(query R, limit int) ([]R, error) {
	model, err := c.readableToModel(query, c.allStores)
	if err != nil {
		return []R{}, err
	}
	results, err := c.primaryStore.ListAllMatching(model, limit)
	var readables []R
	for _, result := range results {
		readables = append(readables, *c.modelToReadable(result))
	}
	return readables, err
}

func (c ImmutableModelController[M, R, C]) Get(selector string) (R, error) {
	result, err := c.primaryStore.Get(selector)
	return *c.modelToReadable(result), err
}

func (c ImmutableModelController[M, R, C]) GetOtherValidSelectors(selector string) ([]string, error) {
	return c.primaryStore.GetOtherValidSelectors(selector)
}

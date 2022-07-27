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

type Editable[R Readable, C Creatable[R]] interface {
	toCreatable() C
}

type ModelController[M v2models.Model, R Readable, C Creatable[R], E Editable[R, C]] struct {
	primaryStore    v2models.Store[M]
	allStores       v2models.StoreSet
	modelToReadable func(model M) *R
	readableToModel func(readable R, stores v2models.StoreSet) (M, error)

	// setDynamicDefaults is an optional function to set context-dynamic defaults during creation.
	//
	// Normally, github.com/creasty/defaults is run to respect the same `default` struct tags that
	// github.com/swaggo/swag understands for documentation.
	//
	// However, sometimes this isn't enough. When setDynamicDefaults is present, it will get run first, before
	// creasty/defaults. It needn't worry about handling those `default` struct tags--it can be provided to allow a type
	// to set defaults before creation that are dynamic based on other existing data or the calling user.
	//
	// (Implementation note:
	setDynamicDefaults func(readable *R, stores v2models.StoreSet, user *auth.User) error
}

func (c ModelController[M, R, C, E]) Create(creatable C, user *auth.User) (R, error) {
	readable := creatable.toReadable()
	// Handle dynamic defaults, like making an environment from a template
	if c.setDynamicDefaults != nil {
		if err := c.setDynamicDefaults(&readable, c.allStores, user); err != nil {
			return readable, fmt.Errorf("(%s) error setting dynamic default values for %T: %v", errors.InternalServerError, readable, err)
		}
	}
	// Handle static struct default tags, which both swaggo/swag and this creasty/defaults.Set function respect
	if err := defaults.Set(&readable); err != nil {
		return readable, fmt.Errorf("(%s) error setting static default values for %T: %v", errors.InternalServerError, readable, err)
	}
	model, err := c.readableToModel(readable, c.allStores)
	if err != nil {
		return readable, err
	}
	result, err := c.primaryStore.Create(model, user)
	return *c.modelToReadable(result), err
}

func (c ModelController[M, R, C, E]) ListAllMatching(query R, limit int) ([]R, error) {
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

func (c ModelController[M, R, C, E]) Get(selector string) (R, error) {
	result, err := c.primaryStore.Get(selector)
	return *c.modelToReadable(result), err
}

func (c ModelController[M, R, C, E]) GetOtherValidSelectors(selector string) ([]string, error) {
	return c.primaryStore.GetOtherValidSelectors(selector)
}

func (c ModelController[M, R, C, E]) Edit(selector string, editable E, user *auth.User) (R, error) {
	readable := editable.toCreatable().toReadable()
	model, err := c.readableToModel(readable, c.allStores)
	if err != nil {
		return readable, err
	}
	result, err := c.primaryStore.Edit(selector, model, user)
	return *c.modelToReadable(result), err
}

func (c ModelController[M, R, C, E]) Delete(selector string, user *auth.User) (R, error) {
	result, err := c.primaryStore.Delete(selector, user)
	return *c.modelToReadable(result), err
}

package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
)

type Editable[R Readable, C Creatable[R]] interface {
	toCreatable() C
}

type MutableModelController[M v2models.Model, R Readable, C Creatable[R], E Editable[R, C]] struct {
	ImmutableModelController[M, R, C]
}

func (c MutableModelController[M, R, C, E]) Edit(selector string, editable E, user auth.User) (R, error) {
	readable := editable.toCreatable().toReadable()
	model, err := c.readableToModel(readable, c.allStores)
	if err != nil {
		return readable, err
	}
	result, err := c.primaryStore.Edit(selector, model, user.Suitable)
	return *c.modelToReadable(result), err
}

func (c MutableModelController[M, R, C, E]) Delete(selector string, user auth.User) (R, error) {
	result, err := c.primaryStore.Delete(selector, user.Suitable)
	return *c.modelToReadable(result), err
}

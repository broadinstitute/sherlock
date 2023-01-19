package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/broadinstitute/sherlock/internal/pagerduty"
	"github.com/creasty/defaults"
	"time"
)

type ReadableBaseType struct {
	ID        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" format:"date-time"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" format:"date-time"`
}

// Readable represents the full set of fields that can be read (or queried for) by a user.
type Readable[M v2models.Model] interface {
	toModel(stores *v2models.StoreSet) (M, error)
}

// Creatable represents the set of fields that can be set upon creation by a user.
type Creatable[M v2models.Model] interface {
	toModel(stores *v2models.StoreSet) (M, error)
}

// Editable represents the set of fields that can be mutated by a user.
type Editable[M v2models.Model] interface {
	toModel(stores *v2models.StoreSet) (M, error)
}

// ModelController exposes the same "verbs" exposed by a v2models.internalStore, but it adds the user-type to database-type
// mapping that provides type safety for what fields can be read/queried, created, and edited. ModelController also
// handles setting defaults--even complex ones, like from template Environment entries.
//
// Implementation note: this mapping behavior exists at the controller level (rather than in serializers, etc. written
// elsewhere) because going from a user-type to a database-type actually itself requires a database connection, so it
// can resolve associations. For example, a user-type would allow an association to be referenced by name, ID, or any
// other selector, but a database-type would specifically use the ID as the foreign key. ModelController is responsible
// for doing that translation. A bonus of defining the controller in terms of user-types is that defaults can be
// handled in terms of the user-type, making for simpler documentation and more obvious behavior.
type ModelController[M v2models.Model, R Readable[M], C Creatable[M], E Editable[M]] struct {
	// primaryStore is the part of the model that this is a controller for.
	primaryStore *v2models.ModelStore[M]
	// allStores is a reference to the entire model, so that readableToModel and setDynamicDefaults can work
	// with associations if they need to.
	allStores *v2models.StoreSet
	// modelToReadable is responsible for converting database results to user-readable responses.
	modelToReadable func(model *M) *R

	// Optional:

	// setDynamicDefaults allows setting context-dynamic defaults during creation.
	//
	// Normally, github.com/creasty/defaults is run to respect the same `default` struct tags that
	// github.com/swaggo/swag understands for documentation.
	//
	// However, sometimes this isn't enough. When setDynamicDefaults is present, it will get run first, before
	// creasty/defaults. It shouldn't worry about handling those `default` struct tags--it can be provided to allow a
	// type to set defaults before creation that are dynamic based on other existing data or the calling user.
	setDynamicDefaults func(creatable *C, stores *v2models.StoreSet, user *auth.User) error

	// extractPagerdutyIntegrationKey allows a data type to declare how to go from the database type to a Pagerduty key.
	//
	// Defining this function and beehiveUrlFormatString enables the TriggerPagerdutyIncident method.
	extractPagerdutyIntegrationKey func(model *M) *string

	// beehiveUrlFormatString is a format string to go from a selector to a Beehive link.
	//
	// Defining this and extractPagerdutyIntegrationKey enables the TriggerPagerdutyIncident method.
	beehiveUrlFormatString string
}

func (c ModelController[M, R, C, E]) Create(creatable C, user *auth.User) (R, bool, error) {
	var empty R
	// Handle dynamic defaults, like making an environment from a template
	if c.setDynamicDefaults != nil {
		if err := c.setDynamicDefaults(&creatable, c.allStores, user); err != nil {
			return empty, false, fmt.Errorf("error setting dynamic default values for %T: %v", creatable, err)
		}
	}
	// Handle static struct default tags, which both swaggo/swag and this creasty/defaults.Set function respect
	if err := defaults.Set(&creatable); err != nil {
		return empty, false, fmt.Errorf("(%s) error setting static default values for %T: %v", errors.InternalServerError, creatable, err)
	}
	model, err := creatable.toModel(c.allStores)
	if err != nil {
		return empty, false, err
	}
	result, created, err := c.primaryStore.Create(model, user)
	return *c.modelToReadable(&result), created, err
}

func (c ModelController[M, R, C, E]) ListAllMatching(filter R, limit int) ([]R, error) {
	model, err := filter.toModel(c.allStores)
	if err != nil {
		return []R{}, fmt.Errorf("error parsing filter to a %T that can be queried against the database: %v", model, err)
	}
	results, err := c.primaryStore.ListAllMatchingByUpdated(model, limit)
	readables := make([]R, 0)
	for _, result := range results {
		readables = append(readables, *c.modelToReadable(&result))
	}
	return readables, err
}

func (c ModelController[M, R, C, E]) Get(selector string) (R, error) {
	result, err := c.primaryStore.Get(selector)
	return *c.modelToReadable(&result), err
}

func (c ModelController[M, R, C, E]) GetOtherValidSelectors(selector string) ([]string, error) {
	return c.primaryStore.GetOtherValidSelectors(selector)
}

func (c ModelController[M, R, C, E]) Edit(selector string, editable E, user *auth.User) (R, error) {
	var empty R
	model, err := editable.toModel(c.allStores)
	if err != nil {
		return empty, err
	}
	result, err := c.primaryStore.Edit(selector, model, user)
	return *c.modelToReadable(&result), err
}

// Upsert is "dumb": it tries to edit, and if there's an error, it tries to create. Edit will always error if the
// selector didn't match, but it could error for other reasons too. We're relying on Create also error-ing in the case
// of those other reasons, which is reasonable since the same validation functions get called.
func (c ModelController[M, R, C, E]) Upsert(selector string, creatable C, editable E, user *auth.User) (R, bool, error) {
	ret, err := c.Edit(selector, editable, user)
	if err != nil {
		return c.Create(creatable, user)
	} else {
		return ret, false, nil
	}
}

func (c ModelController[M, R, C, E]) Delete(selector string, user *auth.User) (R, error) {
	result, err := c.primaryStore.Delete(selector, user)
	return *c.modelToReadable(&result), err
}

func (c ModelController[M, R, C, E]) TriggerPagerdutyIncident(selector string, summary pagerduty.AlertSummary) (pagerduty.SendAlertResponse, error) {
	if c.extractPagerdutyIntegrationKey == nil {
		var empty R
		return pagerduty.SendAlertResponse{}, fmt.Errorf("(%s) Pagerduty incidents can't be triggered via %Ts, no extractPagerdutyIntegrationKey functionality configured", errors.InternalServerError, empty)
	}
	if c.beehiveUrlFormatString == "" {
		var empty R
		return pagerduty.SendAlertResponse{}, fmt.Errorf("(%s) Pagerduty incidents can't be triggered via %Ts, no beehiveUrlFormatString configured", errors.InternalServerError, empty)
	}
	match, err := c.primaryStore.Get(selector)
	if err != nil {
		return pagerduty.SendAlertResponse{}, err
	}
	if key := c.extractPagerdutyIntegrationKey(&match); key != nil {
		return pagerduty.SendAlert(*key, summary, fmt.Sprintf(c.beehiveUrlFormatString, selector))
	} else {
		return pagerduty.SendAlertResponse{}, fmt.Errorf("(%s) no Pagerduty integration configured for %T '%s'", errors.BadRequest, match, selector)
	}
}

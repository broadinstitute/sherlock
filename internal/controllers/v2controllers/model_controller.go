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
	ID        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

// Readable represents the full set of fields that can be read (or queried for) by a user.
// Generally, a Readable will at east embed a ReadableBaseType and a Creatable inside it, but it can have additional
// read-only fields. A Readable should map to some database v2models.Model type; see ModelController for more context.
type Readable interface{}

// Creatable represents the set of fields that can be set upon creation by a user.
// The fields available on a Creatable are a subset of those available on a Readable. Generally, a Creatable will
// embed an Editable, since all fields that can be edited can be set upon creation.
type Creatable[R Readable] interface {
	// toReadable is a type-safe hoisting method, to the superset Readable type.
	// This allows ModelController methods to accept Creatable types as input and hoist them into full Readable types
	// to map them to database v2models.Model types.
	// Note that Go linters currently (as of August 2022) have trouble detecting when a generic method is unused, so
	// implementers may need to use a `//nolint:unused` directive.
	toReadable() R
}

// Editable represents the set of fields that can be mutated by a user.
// The fields available on an Editable are a subset of those available on a Creatable.
type Editable[R Readable, C Creatable[R]] interface {
	// toCreatable is a type-safe hoisting method, to the superset Creatable type.
	// This allows ModelController methods to accept Editable types as input and hoist them into Creatable (and then
	// Readable) types to map them to database v2models.Model types.
	// Note that Go linters currently (as of August 2022) have trouble detecting when a generic method is unused, so
	// implementers may need to use a `//nolint:unused` directive.
	toCreatable() C
}

// ModelController exposes the same "verbs" exposed by a v2models.Store, but it adds the user-type to database-type
// mapping that provides type safety for what fields can be read/queried, created, and edited. ModelController also
// handles setting defaults--even complex ones, like from template Environment entries.
//
// Implementation note: this mapping behavior exists at the controller level (rather than in serializers, etc. written
// elsewhere) because going from a user-type to a database-type actually itself requires a database connection, so it
// can resolve associations. For example, a user-type would allow an association to be referenced by name, ID, or any
// other selector, but a database-type would specifically use the ID as the foreign key. ModelController is responsible
// for doing that translation. A bonus of defining the controller in terms of user-types is that defaults can be
// handled in terms of the user-type, making for simpler documentation and more obvious behavior.
type ModelController[M v2models.Model, R Readable, C Creatable[R], E Editable[R, C]] struct {
	// primaryStore is the part of the model that this is a controller for.
	primaryStore *v2models.Store[M]

	// allStores is a reference to the entire model, so that readableToModel and setDynamicDefaults can work
	// with associations if they need to.
	allStores *v2models.StoreSet

	// modelToReadable is a required half of the Readable-v2models.Model mapping that the ModelController establishes.
	//
	// Since this direction is coming from the database type, it can be done offline and without errors.
	modelToReadable func(model M) *R

	// readableToModel is a required half of the Readable-v2models.Model mapping that the ModelController establishes.
	//
	// Since this direction is coming from the user type, it may take advantage of the database stores to load
	// associations, and it may throw errors if there are problems doing that. It should not perform other validation,
	// since that is the job of the model and this function can be correctly called with invalid Readable types (for
	// example the filters given to ListAllMatching).
	readableToModel func(readable R, stores *v2models.StoreSet) (M, error)

	// setDynamicDefaults is an optional function to set context-dynamic defaults during creation.
	//
	// Normally, github.com/creasty/defaults is run to respect the same `default` struct tags that
	// github.com/swaggo/swag understands for documentation.
	//
	// However, sometimes this isn't enough. When setDynamicDefaults is present, it will get run first, before
	// creasty/defaults. It shouldn't worry about handling those `default` struct tags--it can be provided to allow a
	// type to set defaults before creation that are dynamic based on other existing data or the calling user.
	setDynamicDefaults func(readable *R, stores *v2models.StoreSet, user *auth.User) error
}

func (c ModelController[M, R, C, E]) Create(creatable C, user *auth.User) (R, bool, error) {
	readable := creatable.toReadable()
	// Handle dynamic defaults, like making an environment from a template
	if c.setDynamicDefaults != nil {
		if err := c.setDynamicDefaults(&readable, c.allStores, user); err != nil {
			return readable, false, fmt.Errorf("error setting dynamic default values for %T: %v", readable, err)
		}
	}
	// Handle static struct default tags, which both swaggo/swag and this creasty/defaults.Set function respect
	if err := defaults.Set(&readable); err != nil {
		return readable, false, fmt.Errorf("(%s) error setting static default values for %T: %v", errors.InternalServerError, readable, err)
	}
	model, err := c.readableToModel(readable, c.allStores)
	if err != nil {
		return readable, false, err
	}
	result, err := c.primaryStore.Create(model, user)
	return *c.modelToReadable(result), err == nil, err
}

func (c ModelController[M, R, C, E]) ListAllMatching(filter R, limit int) ([]R, error) {
	model, err := c.readableToModel(filter, c.allStores)
	if err != nil {
		return []R{}, fmt.Errorf("error parsing filter to a %T that can be queried against the database: %v", model, err)
	}
	results, err := c.primaryStore.ListAllMatching(model, limit)
	readables := make([]R, 0)
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

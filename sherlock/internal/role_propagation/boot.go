package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/google/uuid"
)

var propagators []propagator

// Init sets up the propagators to be used during normal operation. They will
// run in the given order, which can be important if one creates accounts that
// a later one will attempt to put into groups.
func Init(ctx context.Context) error {
	propagators = []propagator{

		&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			configKey: "devFirecloudGroup",
			getGrant:  func(role models.Role) *string { return role.GrantsDevFirecloudGroup },
			engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
		},

		&propagatorImpl[uuid.UUID, propagation_engines.AzureGroupIdentifier, propagation_engines.AzureGroupFields]{
			configKey: "devAzureGroup",
			getGrant:  func(role models.Role) *uuid.UUID { return role.GrantsDevAzureGroup },
			engine:    &propagation_engines.AzureGroupEngine{},
		},
	}
	for _, p := range propagators {
		if err := p.Init(ctx); err != nil {
			return err
		}
	}
	return nil
}

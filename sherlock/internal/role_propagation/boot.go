package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
)

var propagators []propagator

// Init sets up the propagators to be used during normal operation. They will
// run in the given order, which can be important if one creates accounts that
// a later one will attempt to put into groups.
func Init(ctx context.Context) error {
	propagators = []propagator{

		&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			configKey: "prodFirecloudGroup",
			getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsProdFirecloudGroup) },
			engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
		},
		&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			configKey: "qaFirecloudGroup",
			getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsQaFirecloudGroup) },
			engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
		},
		&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			configKey: "devFirecloudGroup",
			getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsDevFirecloudGroup) },
			engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
		},

		&propagatorImpl[string, propagation_engines.AzureGroupIdentifier, propagation_engines.AzureGroupFields]{
			configKey: "prodAzureGroup",
			getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsProdAzureGroup) },
			engine:    &propagation_engines.AzureGroupEngine{},
		},
		&propagatorImpl[string, propagation_engines.AzureGroupIdentifier, propagation_engines.AzureGroupFields]{
			configKey: "devAzureGroup",
			getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsDevAzureGroup) },
			engine:    &propagation_engines.AzureGroupEngine{},
		},

		&propagatorImpl[string, propagation_engines.NonAdminGoogleGroupIdentifier, propagation_engines.NonAdminGoogleGroupFields]{
			configKey: "broadInstituteGroup",
			getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsBroadInstituteGroup) },
			engine:    &propagation_engines.NonAdminGoogleGroupEngine{},
		},
	}
	for _, p := range propagators {
		if err := p.Init(ctx); err != nil {
			return err
		}
	}
	return nil
}

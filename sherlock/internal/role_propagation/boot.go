package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
)

// propagators will run in sequence. Use parallelizingPropagator instances to
// make parallel steps.
var propagators []propagator

// Init sets up the propagators to be used during normal operation. They will
// run in the given order, which can be important if one creates accounts that
// a later one will attempt to put into groups.
func Init(ctx context.Context) error {
	propagators = []propagator{

		// This step is for provisioning new accounts.
		&parallelizingPropagator{
			parallelPropagators: []propagator{
				&propagatorImpl[bool, propagation_engines.AzureAccountIdentifier, propagation_engines.AzureAccountFields]{
					configKey: "devAzureAccount",
					getGrants: func(role models.Role) []*bool { return []*bool{role.GrantsDevAzureAccount} },
					engine:    &propagation_engines.AzureAccountEngine{},
				},
				&propagatorImpl[bool, propagation_engines.AzureAccountIdentifier, propagation_engines.AzureAccountFields]{
					configKey: "prodAzureAccount",
					getGrants: func(role models.Role) []*bool { return []*bool{role.GrantsProdAzureAccount} },
					engine:    &propagation_engines.AzureAccountEngine{},
				},
			},
		},

		// This step is for inviting or importing accounts provisioned above to
		// other organizations or tenants, required before they can potentially
		// be granted more localized access in the next step.
		&parallelizingPropagator{
			parallelPropagators: []propagator{
				&propagatorImpl[bool, propagation_engines.AzureInvitedAccountIdentifier, propagation_engines.AzureInvitedAccountFields]{
					configKey: "devAzureInvitedB2CAccount",
					getGrants: func(role models.Role) []*bool { return []*bool{role.GrantsDevAzureAccount} },
					engine:    &propagation_engines.AzureInvitedAccountEngine{},
				},
				&propagatorImpl[bool, propagation_engines.AzureInvitedAccountIdentifier, propagation_engines.AzureInvitedAccountFields]{
					configKey: "prodAzureInvitedB2CAccount",
					getGrants: func(role models.Role) []*bool { return []*bool{role.GrantsProdAzureAccount} },
					engine:    &propagation_engines.AzureInvitedAccountEngine{},
				},
			},
		},

		// This step is for granting pre-existing accounts access to stuff.
		&parallelizingPropagator{
			parallelPropagators: []propagator{
				&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
					configKey: "devFirecloudGroup",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsDevFirecloudGroup) },
					engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
				},
				&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
					configKey: "qaFirecloudGroup",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsQaFirecloudGroup) },
					engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
				},
				&propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
					configKey: "prodFirecloudGroup",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsProdFirecloudGroup) },
					engine:    &propagation_engines.GoogleWorkspaceGroupEngine{},
				},

				&propagatorImpl[string, propagation_engines.GoogleCloudFolderRoleIdentifier, propagation_engines.GoogleCloudFolderRoleFields]{
					configKey: "devFirecloudFolderOwner",
					getGrants: func(role models.Role) []*string {
						return splitStringPointerOnCommas(role.GrantsDevFirecloudFolderOwner)
					},
					engine: &propagation_engines.GoogleCloudFolderRoleEngine{Role: propagation_engines.GoogleCloudOwnerRole},
				},
				&propagatorImpl[string, propagation_engines.GoogleCloudFolderRoleIdentifier, propagation_engines.GoogleCloudFolderRoleFields]{
					configKey: "qaFirecloudFolderOwner",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsQaFirecloudFolderOwner) },
					engine:    &propagation_engines.GoogleCloudFolderRoleEngine{Role: propagation_engines.GoogleCloudOwnerRole},
				},
				&propagatorImpl[string, propagation_engines.GoogleCloudFolderRoleIdentifier, propagation_engines.GoogleCloudFolderRoleFields]{
					configKey: "prodFirecloudFolderOwner",
					getGrants: func(role models.Role) []*string {
						return splitStringPointerOnCommas(role.GrantsProdFirecloudFolderOwner)
					},
					engine: &propagation_engines.GoogleCloudFolderRoleEngine{Role: propagation_engines.GoogleCloudOwnerRole},
				},

				&propagatorImpl[string, propagation_engines.AzureGroupIdentifier, propagation_engines.AzureGroupFields]{
					configKey: "devAzureGroup",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsDevAzureGroup) },
					engine:    &propagation_engines.AzureGroupEngine{},
				},
				&propagatorImpl[string, propagation_engines.AzureGroupIdentifier, propagation_engines.AzureGroupFields]{
					configKey: "prodAzureGroup",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsProdAzureGroup) },
					engine:    &propagation_engines.AzureGroupEngine{},
				},

				&propagatorImpl[bool, propagation_engines.AzureDirectoryRoleIdentifier, propagation_engines.AzureDirectoryRoleFields]{
					configKey: "devAzureB2CReader",
					getGrants: func(role models.Role) []*bool { return []*bool{role.GrantsDevAzureDirectoryRoles} },
					engine:    &propagation_engines.AzureDirectoryRoleEngine{RoleTemplateID: propagation_engines.AzureGlobalReaderRoleTemplateID},
				},
				&propagatorImpl[bool, propagation_engines.AzureDirectoryRoleIdentifier, propagation_engines.AzureDirectoryRoleFields]{
					configKey: "prodAzureB2CReader",
					getGrants: func(role models.Role) []*bool { return []*bool{role.GrantsProdAzureDirectoryRoles} },
					engine:    &propagation_engines.AzureDirectoryRoleEngine{RoleTemplateID: propagation_engines.AzureGlobalReaderRoleTemplateID},
				},

				&propagatorImpl[string, propagation_engines.NonAdminGoogleGroupIdentifier, propagation_engines.NonAdminGoogleGroupFields]{
					configKey: "broadInstituteGroup",
					getGrants: func(role models.Role) []*string { return splitStringPointerOnCommas(role.GrantsBroadInstituteGroup) },
					engine:    &propagation_engines.NonAdminGoogleGroupEngine{},
				},
			},
		},
	}
	for _, p := range propagators {
		if err := p.Init(ctx); err != nil {
			return fmt.Errorf("failed to initialize propagator: %s%w", p.LogPrefix(), err)
		}
	}
	return nil
}

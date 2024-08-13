package propagation_engines

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/knadh/koanf"
)

// PropagationEngine represents a mechanism to propagate role assignments out into the cloud. You can think of a
// propagationEngine implementation as handling one "type" of grant field on a models.Role, like "Firecloud.org group"
// or "Firecloud.org account." There might be multiple of those fields on models.Role, like for dev/QA/prod
// Firecloud.org, and the idea is there'd be an engine instance per each one of those (just with different config).
//
// See propagators in ./boot.go for examples. Implementations of this interface should generally use pointer receivers.
type PropagationEngine[
	// Grant is what we're trying to, well, grant. In many cases, this'll be a string, representing the name of the
	// group we want to add people to. It could be a UUID for granting some Azure role or permission. The weirdest
	// likely case is that it could be a boolean, if the propagator were meant to grant an actual account or something.
	//
	// A good rule of thumb for understanding this type is that it's however the grant is stored on the models.Role.
	Grant any,
	// Identifier is a struct containing how the engine identifies users on the cloud provider. The key thing is that
	// the engine should be able to read this from the cloud provider, so that means "Sherlock user ID" should almost
	// definitely not be in here. For granting a Google group, this might just be the email address of the user.
	//
	// A good rule of thumb for understanding this type is that it provides just enough information for us to "get"
	// a user on the cloud provider.
	Identifier intermediary_user.Identifier,
	// Fields is a struct containing non-identifying but still Sherlock-manipulated data for the user on the cloud
	// provider. For example, if we're granting a Firecloud.org account, the fields might contain the user's name,
	// since we don't consider that unique but we do want to control it. Fields can also be how an engine represents
	// suspensions.
	//
	// A good rule of thumb for understanding this type is that it contains all the information that we want to be
	// able to change for a user on the cloud provider.
	Fields intermediary_user.Fields,
] interface {
	// Init runs any instance-specific setup for this engine. Errors returned here will abort Sherlock's boot process.
	// The *koanf.Koanf is the instance-specific configuration.
	Init(ctx context.Context, k *koanf.Koanf) error

	// LoadCurrentState loads and returns the current state of the grant on the cloud provider, like who all is in the
	// group. This function shouldn't make any judgement about whether the remote state is correct or not -- it just
	// tells us what it is right now.
	LoadCurrentState(ctx context.Context, grant Grant) ([]intermediary_user.IntermediaryUser[Identifier, Fields], error)

	// GenerateDesiredState assembles the set of intermediary users that should have the grant. This function may
	// return fewer results than the input map, for example if an input entry has no corresponding intermediary user
	// or if an input entry is suspended and the engine handles that by not giving the intermediary user the grant
	// at all.
	GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[Identifier, Fields], error)

	// Add directs the engine to give the grant to the intermediary user (and set the given initial fields).
	// It won't be called if the engine has reported the identifier as already having the grant.
	//
	// It should return a string that will be logged as the result of the operation.
	Add(ctx context.Context, grant Grant, identifier Identifier, fields Fields) (string, error)

	// Update directs the engine to update the fields of the intermediary user on the given grant.
	// It will only be called if getGrantState and translateRoleAssignments return intermediary users
	// with equal identifiers but different fields (if the identifier isn't present in both or the fields
	// are the same, this function won't be called). It's safe to leave this function unimplemented if the
	// fields will always be the same (perhaps because it is an empty struct).
	//
	// It should return a string that will be logged as the result of the operation.
	Update(ctx context.Context, grant Grant, identifier Identifier, oldFields Fields, newFields Fields) (string, error)

	// Remove directs the engine to remove the grant from the intermediary user.
	// It won't be called if the engine hasn't reported the identifier as having the grant.
	//
	// It should return a string that will be logged as the result of the operation.
	Remove(ctx context.Context, grant Grant, identifier Identifier) (string, error)
}

// ToleratedUserCalculator is an interface that a PropagationEngine may OPTIONALLY implement.
// If implemented, CalculateToleratedUsers will be called after the engine's initialization to
// load dynamic but non-grant-specific identities that should be globally tolerated by the
// propagator.
//
// Put simply, if an engine implements this interface, it won't ever be asked to remove any of
// the identities it returns from CalculateToleratedUsers. CalculateToleratedUsers is called
// a single time when Sherlock boots.
//
// This is mainly helpful when Sherlock itself will show up in the current state of a grant.
// We obviously don't want to remove Sherlock itself from a grant (i.e. for
// NonAdminGoogleGroupEngine, this is how it actually has permissions to modify the group), but
// it's a bit of a sharp edge to require us to hardcode Sherlock's identity into the config file
// as a tolerated identity. This interface allows the engine to figure out its identity and add
// it to that all-important list.
type ToleratedUserCalculator[
	Identifier intermediary_user.Identifier,
] interface {
	CalculateToleratedUsers(ctx context.Context) ([]Identifier, error)
}

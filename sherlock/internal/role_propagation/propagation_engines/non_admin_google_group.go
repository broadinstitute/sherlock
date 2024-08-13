package propagation_engines

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/bits_data_warehouse"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/self"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/knadh/koanf"
	cloudidentity "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/option"
)

type NonAdminGoogleGroupIdentifier struct {
	Email string `koanf:"email"`

	// resourceName is like "groups/<ID>/memberships/<ID>". It's output-only from the Cloud Identity API.
	// We need it to delete or update membership, but it obviously can't exist before creation. This
	// field here is basically a note-taking field. In NonAdminGoogleGroupEngine.LoadCurrentState we
	// learn of the value and we write it down. In NonAdminGoogleGroupEngine.Remove we check to see
	// if we already have a value. If not, we can make an extra round trip to look it up.
	//
	// This isn't an exported field because it's an unpredictable value we're comfortable encapsulating
	// within this engine. It makes sense to try to carry it between LoadCurrentState and Remove, but
	// it isn't something relevant for determining if someone should be in a group or not (so we don't
	// want it to be checked in EqualTo, for example).
	resourceName *string
}

func (n NonAdminGoogleGroupIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case NonAdminGoogleGroupIdentifier:
		return n.Email == other.Email
	default:
		return false
	}
}

type NonAdminGoogleGroupFields struct{}

func (n NonAdminGoogleGroupFields) EqualTo(other intermediary_user.Fields) bool {
	switch other.(type) {
	case NonAdminGoogleGroupFields:
		return true
	default:
		return false
	}
}

// NonAdminGoogleGroupEngine is like GoogleWorkspaceGroupEngine but it's designed to operate without workspace-level
// admin privileges.
type NonAdminGoogleGroupEngine struct {
	service *cloudidentity.Service
}

func (n *NonAdminGoogleGroupEngine) Init(ctx context.Context, k *koanf.Koanf) error {
	var err error
	n.service, err = cloudidentity.NewService(ctx, option.WithScopes(cloudidentity.CloudIdentityGroupsScope))
	return err
}

func (n *NonAdminGoogleGroupEngine) CalculateToleratedUsers(_ context.Context) ([]NonAdminGoogleGroupIdentifier, error) {
	return []NonAdminGoogleGroupIdentifier{
		{
			Email: self.Email,
		},
	}, nil
}

func (n *NonAdminGoogleGroupEngine) LoadCurrentState(ctx context.Context, grant string) ([]intermediary_user.IntermediaryUser[NonAdminGoogleGroupIdentifier, NonAdminGoogleGroupFields], error) {
	lookupResponse, err := n.service.Groups.Lookup().GroupKeyId(grant).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to lookup cloudidentity group ID for %s: %w", grant, err)
	}

	currentState := make([]intermediary_user.IntermediaryUser[NonAdminGoogleGroupIdentifier, NonAdminGoogleGroupFields], 0)
	err = n.service.Groups.Memberships.List(lookupResponse.Name).Pages(ctx, func(memberships *cloudidentity.ListMembershipsResponse) error {
		for _, membership := range memberships.Memberships {
			resourceName := membership.Name
			currentState = append(currentState, intermediary_user.IntermediaryUser[NonAdminGoogleGroupIdentifier, NonAdminGoogleGroupFields]{
				Identifier: NonAdminGoogleGroupIdentifier{
					Email:        membership.PreferredMemberKey.Id,
					resourceName: &resourceName,
				},
				Fields: NonAdminGoogleGroupFields{},
			})
		}
		return nil
	})
	return currentState, err
}

func (n *NonAdminGoogleGroupEngine) GenerateDesiredState(_ context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[NonAdminGoogleGroupIdentifier, NonAdminGoogleGroupFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[NonAdminGoogleGroupIdentifier, NonAdminGoogleGroupFields])
	for id, roleAssignment := range roleAssignments {
		if !roleAssignment.IsActive() {
			// There's no concept of a suspended group member, so we just exclude any non-active users
			continue
		}

		if _, found, err := bits_data_warehouse.GetPerson(roleAssignment.User.Email); err != nil {
			return nil, fmt.Errorf("bits_data_warehouse error: %w", err)
		} else if found {
			// If the user is found in bits_data_warehouse, that means their account exists enough
			// that it's reasonable for us to add them to the desired state
			desiredState[id] = intermediary_user.IntermediaryUser[NonAdminGoogleGroupIdentifier, NonAdminGoogleGroupFields]{
				Identifier: NonAdminGoogleGroupIdentifier{Email: roleAssignment.User.Email},
				Fields:     NonAdminGoogleGroupFields{},
			}
		}
	}
	return desiredState, nil
}

func (n *NonAdminGoogleGroupEngine) Add(ctx context.Context, grant string, identifier NonAdminGoogleGroupIdentifier, _ NonAdminGoogleGroupFields) (string, error) {
	lookupResponse, err := n.service.Groups.Lookup().GroupKeyId(grant).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("failed to add %s to %s due to cloudidentity lookup error: %w", identifier.Email, grant, err)
	}

	_, err = n.service.Groups.Memberships.Create(lookupResponse.Name, &cloudidentity.Membership{
		PreferredMemberKey: &cloudidentity.EntityKey{
			Id: identifier.Email,
		},
		Roles: []*cloudidentity.MembershipRole{
			{
				Name: "MEMBER",
			},
		},
	}).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("failed to add %s to %s: %w", identifier.Email, grant, err)
	} else {
		return fmt.Sprintf("added %s to %s", identifier.Email, grant), nil
	}

}

func (n *NonAdminGoogleGroupEngine) Update(_ context.Context, _ string, _ NonAdminGoogleGroupIdentifier, _ NonAdminGoogleGroupFields, _ NonAdminGoogleGroupFields) (string, error) {
	return "", fmt.Errorf("%T.Update not umplemented; %T.EqualTo should always equal true", n, NonAdminGoogleGroupFields{})
}

func (n *NonAdminGoogleGroupEngine) Remove(ctx context.Context, grant string, identifier NonAdminGoogleGroupIdentifier) (string, error) {
	var resourceName string

	if identifier.resourceName != nil {
		resourceName = *identifier.resourceName
	} else {
		// Not sure we can hit this case but the code is trivial and it's easier to reason about with it being independent like this
		groupLookupResponse, err := n.service.Groups.Lookup().GroupKeyId(grant).Context(ctx).Do()
		if err != nil {
			return "", fmt.Errorf("failed to remove %s from %s due to cloudidentity lookup error: %w", identifier.Email, grant, err)
		}
		membershipLookupResponse, err := n.service.Groups.Memberships.Lookup(groupLookupResponse.Name).MemberKeyId(identifier.Email).Context(ctx).Do()
		if err != nil {
			return "", fmt.Errorf("failed to remove %s from %s due to cloudidentity lookup error: %w", identifier.Email, grant, err)
		}
		resourceName = membershipLookupResponse.Name
	}

	_, err := n.service.Groups.Memberships.Delete(resourceName).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("failed to remove %s from %s: %w", identifier.Email, grant, err)
	} else {
		return fmt.Sprintf("removed %s from %s", identifier.Email, grant), nil
	}
}

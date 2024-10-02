package propagation_engines

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/knadh/koanf"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/directoryroles"
	"github.com/microsoftgraph/msgraph-sdk-go/directoryroleswithroletemplateid"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"strings"
)

const AzureGlobalReaderRoleTemplateID = "f2ef992c-3afb-46b9-b7cf-a126ee74c451"

type AzureDirectoryRoleIdentifier struct {
	// ID is technically a UUID, but this is an intermediary type so we'd just be more brittle if we enforced that.
	// Also makes parsing via koanf more complicated, mapstructure requires a tiny bit of custom code to handle UUIDs.
	ID string `koanf:"id"`
}

func (a AzureDirectoryRoleIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case AzureDirectoryRoleIdentifier:
		return a.ID == other.ID
	default:
		return false
	}
}

type AzureDirectoryRoleFields struct{}

func (a AzureDirectoryRoleFields) EqualTo(other intermediary_user.Fields) bool {
	switch other.(type) {
	case AzureDirectoryRoleFields:
		return true
	default:
		return false
	}
}

var _ PropagationEngine[bool, AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields] = &AzureDirectoryRoleEngine{}

// AzureDirectoryRoleEngine notably uses a boolean grant. This means that there's exactly one non-zero
// grant value -- so exactly one Sherlock role will be able to grant the Azure role an instance of this
// engine provides.
//
// The logic here, and why this is different from GoogleCloudFolderRoleEngine, is that our actual auth
// mechanism changes based on the directory we're granting a role in. At present, we just don't have a
// complex enough directory setup to justify setting client IDs and on-disk projected tokens paths through
// Sherlock's API. It's simplest for these things to be set in config, and while that means this engine
// is a bit "dumb", it does the bare minimum that we need right now.
//
// See below for why Role isn't configurable either.
type AzureDirectoryRoleEngine struct {
	// RoleTemplateID is defined in Sherlock's source code where this engine is instantiated, not in
	// configuration. This is because it helps avoid a foot-gun with accidentally misconfiguring the engine.
	// When this engine is pointed at a directory, it wants to "own" all members of the given role. It will
	// remove members that don't correlate to Sherlock role assignments. If we allowed easy configuration of
	// the role, it would be easy to accidentally cause problems by pointing this engine at a role that some
	// other system also wanted to manage or relied on.
	RoleTemplateID string

	directoryRoleID            string
	memberEmailSuffix          string
	userEmailSuffixesToReplace []string
	client                     *msgraphsdk.GraphServiceClient
}

func (a *AzureDirectoryRoleEngine) Init(ctx context.Context, k *koanf.Koanf) error {
	if a.RoleTemplateID == "" {
		return fmt.Errorf("role template ID must be set")
	}

	a.memberEmailSuffix = k.String("memberEmailSuffix")
	a.userEmailSuffixesToReplace = k.Strings("userEmailSuffixesToReplace")

	credentials, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		ClientID:      k.String("clientID"),
		TenantID:      k.String("tenantID"),
		TokenFilePath: k.String("tokenFilePath"),
	})
	if err != nil {
		return err
	}

	a.client, err = msgraphsdk.NewGraphServiceClientWithCredentials(credentials, nil)
	if err != nil {
		return fmt.Errorf("failed to instantiate MS Graph client: %w", err)
	}

	matchingDirectoryRoles, err := a.client.DirectoryRolesWithRoleTemplateId(&a.RoleTemplateID).Get(ctx, &directoryroleswithroletemplateid.DirectoryRolesWithRoleTemplateIdRequestBuilderGetRequestConfiguration{
		QueryParameters: &directoryroleswithroletemplateid.DirectoryRolesWithRoleTemplateIdRequestBuilderGetQueryParameters{
			Select: []string{"id"},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to fetch directory role ID for role template ID %s: %w", a.RoleTemplateID, err)
	} else if id := matchingDirectoryRoles.GetId(); id == nil {
		return fmt.Errorf("no directory role found for role template ID %s (but no error returned either -- ID was nil)", a.RoleTemplateID)
	} else {
		a.directoryRoleID = *id
	}

	return nil
}

func (a *AzureDirectoryRoleEngine) LoadCurrentState(ctx context.Context, _ bool) ([]intermediary_user.IntermediaryUser[AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields], error) {
	currentState := make([]intermediary_user.IntermediaryUser[AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields], 0)

	headers := abstractions.NewRequestHeaders()
	configuration := &directoryroles.ItemMembersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &directoryroles.ItemMembersRequestBuilderGetQueryParameters{
			Select: []string{"id"},
			// No top because that's actually not supported here
		},
	}

	result, err := a.client.DirectoryRoles().ByDirectoryRoleId(a.directoryRoleID).Members().Get(ctx, configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch directory role members: %w", err)
	}

	pageIterator, err := msgraphgocore.NewPageIterator[graphmodels.DirectoryObjectable](result,
		a.client.GetAdapter(), graphmodels.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create page iterator for directory role members: %w", err)
	}
	pageIterator.SetHeaders(headers)

	err = pageIterator.Iterate(ctx, func(pageItem graphmodels.DirectoryObjectable) bool {
		if id := pageItem.GetId(); id != nil {
			currentState = append(currentState, intermediary_user.IntermediaryUser[AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields]{
				Identifier: AzureDirectoryRoleIdentifier{ID: *id},
				Fields:     AzureDirectoryRoleFields{},
			})
		}
		return true
	})
	return currentState, err
}

func (a *AzureDirectoryRoleEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields])
	for sherlockUserID, roleAssignment := range roleAssignments {
		if !roleAssignment.IsActive() {
			// There's no concept of a suspended group member, so we just exclude any non-active users
			continue
		}

		email := utils.SubstituteSuffix(roleAssignment.User.Email, a.userEmailSuffixesToReplace, a.memberEmailSuffix)
		if !strings.HasSuffix(email, a.memberEmailSuffix) {
			// We can short-circuit here, we know that the user is not in the expected member domain so we won't bother looking
			continue
		}

		usersResponse, err := a.client.Users().Get(ctx, &users.UsersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
				Select: []string{"id"},
				Filter: utils.PointerTo(fmt.Sprintf("userPrincipalName eq '%s'", email)),
				Top:    utils.PointerTo[int32](1),
			},
		})
		if err != nil {
			return nil, err
		} else {
			for _, user := range usersResponse.GetValue() {
				if id := user.GetId(); id != nil {
					desiredState[sherlockUserID] = intermediary_user.IntermediaryUser[AzureDirectoryRoleIdentifier, AzureDirectoryRoleFields]{
						Identifier: AzureDirectoryRoleIdentifier{ID: *user.GetId()},
						Fields:     AzureDirectoryRoleFields{},
					}
				}
			}
		}
	}
	return desiredState, nil
}

func (a *AzureDirectoryRoleEngine) Add(ctx context.Context, _ bool, identifier AzureDirectoryRoleIdentifier, _ AzureDirectoryRoleFields) (string, error) {
	body := graphmodels.NewReferenceCreate()
	body.SetOdataId(utils.PointerTo(fmt.Sprintf("https://graph.microsoft.com/v1.0/directoryObjects/%s", identifier.ID)))
	err := a.client.DirectoryRoles().ByDirectoryRoleId(a.directoryRoleID).Members().Ref().Post(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to add user %s to role %s: %w", identifier.ID, a.directoryRoleID, err)
	} else {
		return fmt.Sprintf("added user %s to role %s", identifier.ID, a.directoryRoleID), nil
	}
}

func (a *AzureDirectoryRoleEngine) Update(_ context.Context, _ bool, _ AzureDirectoryRoleIdentifier, _ AzureDirectoryRoleFields, _ AzureDirectoryRoleFields) (string, error) {
	return "", fmt.Errorf("%T.Update not implemented, %T.EqualTo should always return true", a, AzureDirectoryRoleFields{})
}

func (a *AzureDirectoryRoleEngine) Remove(ctx context.Context, _ bool, identifier AzureDirectoryRoleIdentifier) (string, error) {
	err := a.client.DirectoryRoles().ByDirectoryRoleId(a.directoryRoleID).Members().ByDirectoryObjectId(identifier.ID).Ref().Delete(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to remove user %s from role %s: %w", identifier.ID, a.directoryRoleID, err)
	} else {
		return fmt.Sprintf("removed user %s from role %s", identifier.ID, a.directoryRoleID), nil
	}
}

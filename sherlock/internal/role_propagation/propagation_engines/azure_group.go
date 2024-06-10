package propagation_engines

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/google/uuid"
	"github.com/knadh/koanf"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"strings"
)

type AzureGroupIdentifier struct {
	// ID is technically a UUID, but this is an intermediary type so we'd just be more brittle if we enforced that.
	// Also makes parsing via koanf more complicated, mapstructure requires a tiny bit of custom code to handle UUIDs.
	ID string `koanf:"id"`
}

func (a AzureGroupIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case AzureGroupIdentifier:
		return a.ID == other.ID
	default:
		return false
	}
}

type AzureGroupFields struct{}

func (a AzureGroupFields) EqualTo(other intermediary_user.Fields) bool {
	switch other.(type) {
	case AzureGroupFields:
		return true
	default:
		return false
	}
}

type AzureGroupEngine struct {
	memberEmailSuffix          string
	userEmailSuffixesToReplace []string
	client                     *msgraphsdk.GraphServiceClient
}

func (a *AzureGroupEngine) Init(_ context.Context, k *koanf.Koanf) error {
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

	a.client, err = msgraphsdk.NewGraphServiceClientWithCredentials(credentials, []string{"Group.ReadWrite.All", "User.ReadWrite.All"})
	return err
}

func (a *AzureGroupEngine) LoadCurrentState(ctx context.Context, grant uuid.UUID) ([]intermediary_user.IntermediaryUser[AzureGroupIdentifier, AzureGroupFields], error) {
	currentState := make([]intermediary_user.IntermediaryUser[AzureGroupIdentifier, AzureGroupFields], 0)
	groupMembersResponse, err := a.client.Groups().ByGroupId(grant.String()).Members().Get(ctx, nil)
	if err != nil {
		return nil, err
	} else {
		for _, directoryObject := range groupMembersResponse.GetValue() {
			if id := directoryObject.GetId(); id != nil {
				currentState = append(currentState, intermediary_user.IntermediaryUser[AzureGroupIdentifier, AzureGroupFields]{
					Identifier: AzureGroupIdentifier{ID: *id},
					Fields:     AzureGroupFields{},
				})
			}
		}
	}
	return currentState, nil
}

func (a *AzureGroupEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[AzureGroupIdentifier, AzureGroupFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[AzureGroupIdentifier, AzureGroupFields])
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
					desiredState[sherlockUserID] = intermediary_user.IntermediaryUser[AzureGroupIdentifier, AzureGroupFields]{
						Identifier: AzureGroupIdentifier{ID: *user.GetId()},
						Fields:     AzureGroupFields{},
					}
				}
			}
		}
	}
	return desiredState, nil
}

func (a *AzureGroupEngine) Add(ctx context.Context, grant uuid.UUID, identifier AzureGroupIdentifier, _ AzureGroupFields) (string, error) {
	body := graphmodels.NewReferenceCreate()
	body.SetOdataId(utils.PointerTo(fmt.Sprintf("https://graph.microsoft.com/v1.0/directoryObjects/%s", identifier.ID)))
	err := a.client.Groups().ByGroupId(grant.String()).Members().Ref().Post(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to add user %s to group %s: %w", identifier.ID, grant, err)
	} else {
		return fmt.Sprintf("added user %s to group %s", identifier.ID, grant), nil
	}
}

func (a *AzureGroupEngine) Update(_ context.Context, _ uuid.UUID, _ AzureGroupIdentifier, _ AzureGroupFields, _ AzureGroupFields) (string, error) {
	return "", fmt.Errorf("%T.Update not implemented, %T.EqualTo should always return true", a, AzureGroupFields{})
}

func (a *AzureGroupEngine) Remove(ctx context.Context, grant uuid.UUID, identifier AzureGroupIdentifier) (string, error) {
	err := a.client.Groups().ByGroupId(grant.String()).Members().ByDirectoryObjectId(identifier.ID).Ref().Delete(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to remove user %s from group %s: %w", identifier.ID, grant, err)
	} else {
		return fmt.Sprintf("removed user %s from group %s", identifier.ID, grant), nil
	}
}

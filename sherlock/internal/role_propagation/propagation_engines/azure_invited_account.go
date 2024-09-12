package propagation_engines

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/knadh/koanf"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"strings"
)

type AzureInvitedAccountIdentifier struct {
	Email string `koanf:"email"`
}

func (a AzureInvitedAccountIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case AzureInvitedAccountIdentifier:
		return a.Email == other.Email
	default:
		return false
	}
}

type AzureInvitedAccountFields struct {
	Name string
}

func (a AzureInvitedAccountFields) EqualTo(other intermediary_user.Fields) bool {
	switch other := other.(type) {
	case AzureInvitedAccountFields:
		return a.Name == other.Name
	default:
		return false
	}
}

// MayConsiderAsAlreadyRemoved always returns true for AzureInvitedAccountFields. This is because we never
// remove invited user records. We don't remove users at all -- we just disable them, and the way we do that
// is by disabling the user in their home directory (presumably via AzureCreatedAccountEngine). That prevents
// them from logging in to their account, blocking access to any tenant, home or invited.
func (a AzureInvitedAccountFields) MayConsiderAsAlreadyRemoved() bool {
	return true
}

var _ PropagationEngine[bool, AzureInvitedAccountIdentifier, AzureInvitedAccountFields] = &AzureInvitedAccountEngine{}

type AzureInvitedAccountEngine struct {
	homeTenantEmailSuffix      string
	userEmailSuffixesToReplace []string
	inviteTenantName           string

	homeTenantClient   *msgraphsdk.GraphServiceClient
	inviteTenantClient *msgraphsdk.GraphServiceClient
}

func (a *AzureInvitedAccountEngine) Init(ctx context.Context, k *koanf.Koanf) error {
	a.homeTenantEmailSuffix = k.String("homeTenantEmailSuffix")
	a.userEmailSuffixesToReplace = k.Strings("userEmailSuffixesToReplace")
	a.inviteTenantName = k.String("inviteTenantName")

	homeCredentials, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		ClientID:      k.String("homeTenantClientID"),
		TenantID:      k.String("homeTenantID"),
		TokenFilePath: k.String("homeTenantTokenFilePath"),
	})
	if err != nil {
		return err
	}

	a.homeTenantClient, err = msgraphsdk.NewGraphServiceClientWithCredentials(homeCredentials, nil)
	if err != nil {
		return err
	}

	inviteCredentials, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		ClientID:      k.String("inviteTenantClientID"),
		TenantID:      k.String("inviteTenantID"),
		TokenFilePath: k.String("inviteTenantTokenFilePath"),
	})
	if err != nil {
		return err
	}

	a.inviteTenantClient, err = msgraphsdk.NewGraphServiceClientWithCredentials(inviteCredentials, nil)

	return err
}

func (a *AzureInvitedAccountEngine) LoadCurrentState(ctx context.Context, _ bool) ([]intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields], error) {
	currentState := make([]intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields], 0)
	usersResponse, err := a.inviteTenantClient.Users().Get(ctx, &users.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Select: []string{"userPrincipalName", "displayName"},
			Filter: utils.PointerTo(fmt.Sprintf("endsWith(userPrincipalName, '%s') and creationType eq 'Invitation'", a.homeTenantEmailSuffix)),
		},
	})
	if err != nil {
		return nil, err
	} else {
		for _, directoryObject := range usersResponse.GetValue() {
			if userPrincipalName := directoryObject.GetUserPrincipalName(); userPrincipalName != nil {
				var fields AzureInvitedAccountFields
				if name := directoryObject.GetDisplayName(); name != nil {
					fields.Name = *name
				}
				currentState = append(currentState, intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields]{
					Identifier: AzureInvitedAccountIdentifier{Email: *userPrincipalName},
					Fields:     fields,
				})
			}
		}
	}
	return currentState, nil
}

func (a *AzureInvitedAccountEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields])
	for sherlockUserID, roleAssignment := range roleAssignments {
		// We explicitly aren't calling roleAssignment.IsActive() here. *This* propagator actually doesn't care about
		// suspension! We don't have a notion of suspending an invited account, but we actually don't delete them either!
		// We rely on the propagator that manages the home tenant identity to disable the user there, which suspends their
		// ability to log in here too.
		// We choose to still propagate the user here because we want to keep the user's name up to date in the invite tenant.

		email := utils.SubstituteSuffix(roleAssignment.User.Email, a.userEmailSuffixesToReplace, a.homeTenantEmailSuffix)
		if !strings.HasSuffix(email, a.homeTenantEmailSuffix) {
			// We can short-circuit here, we know that the user doesn't have an email suffix we'd expect in the home tenant
			// so we won't bother looking
			continue
		}

		usersResponse, err := a.homeTenantClient.Users().Get(ctx, &users.UsersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
				Select: []string{"userPrincipalName"},
				Filter: utils.PointerTo(fmt.Sprintf("userPrincipalName eq '%s'", email)),
				Top:    utils.PointerTo[int32](1),
			},
		})
		if err != nil {
			return nil, err
		} else {
			for _, user := range usersResponse.GetValue() {
				if userPrincipalName := user.GetUserPrincipalName(); userPrincipalName != nil {
					desiredState[sherlockUserID] = intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields]{
						Identifier: AzureInvitedAccountIdentifier{Email: *userPrincipalName},
						Fields:     AzureInvitedAccountFields{Name: roleAssignment.User.NameOrUsername()},
					}
				}
			}
		}
	}
	return desiredState, nil
}

func (a *AzureInvitedAccountEngine) inviteMessageBody(identifier AzureInvitedAccountIdentifier) (inviteMessageBody string, identifyingString string, err error) {
	randomBytes := make([]byte, 8)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate random identifying string for inviting %s: %w", identifier.Email, err)
	}
	identifyingString = hex.EncodeToString(randomBytes)

	inviteMessageBody = "This invitation has been generated by the DSP DevOps platform via Microsoft Graph API. " +
		fmt.Sprintf("This invitation is meant to grant your %s Microsoft account access to %s. ", identifier.Email, a.inviteTenantName) +
		"You should reach out to DevOps to confirm the origin of this message before clicking the link. " +
		fmt.Sprintf("They can match this message to a security event with this identifying string: %s. ", identifyingString)

	return inviteMessageBody, identifyingString, nil
}

func (a *AzureInvitedAccountEngine) Add(ctx context.Context, _ bool, identifier AzureInvitedAccountIdentifier, fields AzureInvitedAccountFields) (string, error) {
	inviteMessageBody, identifyingString, err := a.inviteMessageBody(identifier)
	if err != nil {
		return "", err
	}

	body := graphmodels.NewInvitation()
	body.SetInvitedUserEmailAddress(utils.PointerTo(identifier.Email))
	body.SetInviteRedirectUrl(utils.PointerTo("https://portal.azure.com"))
	body.SetInvitedUserType(utils.PointerTo("member"))
	body.SetInvitedUserDisplayName(utils.PointerTo(fields.Name))
	body.SetSendInvitationMessage(utils.PointerTo(true))
	invitedUserMessageInfo := graphmodels.NewInvitedUserMessageInfo()
	invitedUserMessageInfo.SetCustomizedMessageBody(utils.PointerTo(inviteMessageBody))
	body.SetInvitedUserMessageInfo(invitedUserMessageInfo)
	_, err = a.inviteTenantClient.Invitations().Post(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to invite %s: %w", identifier.Email, err)
	} else {
		return fmt.Sprintf("invited %s (invite email sent with identifying string `%s`)", identifier.Email, identifyingString), nil
	}
}

func (a *AzureInvitedAccountEngine) Update(ctx context.Context, _ bool, identifier AzureInvitedAccountIdentifier, oldFields AzureInvitedAccountFields, newFields AzureInvitedAccountFields) (string, error) {
	// We can't update an invitation (an email has already been sent), but if the fields are different that means the name has changed...
	// and we can update that on the user's identity on the "invite" tenant we're controlling here.
	// Our identifier doesn't have the user ID -- it has the email instead -- but again, if the fields are different, that means the
	// email actually matches a user principal name in the invite tenant, so we can use that directly.
	body := graphmodels.NewUser()
	body.SetDisplayName(utils.PointerTo(newFields.Name))
	_, err := a.inviteTenantClient.Users().ByUserId(identifier.Email).Patch(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to update user %s's display name from `%s` to `%s`: %w", identifier.Email, oldFields.Name, newFields.Name, err)
	} else {
		return fmt.Sprintf("updated user %s's display name from `%s` to `%s`", identifier.Email, oldFields.Name, newFields.Name), nil
	}
}

func (a *AzureInvitedAccountEngine) Remove(_ context.Context, _ bool, _ AzureInvitedAccountIdentifier) (string, error) {
	return "", fmt.Errorf("%T.Remove not implemented, %T.MayConsiderAsAlreadyRemoved should always return true", a, AzureInvitedAccountFields{})
}

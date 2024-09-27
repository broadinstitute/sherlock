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
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"reflect"
	"strings"
)

type AzureInvitedAccountIdentifier struct {
	UserPrincipalName string `koanf:"userPrincipalName"`
}

func (a AzureInvitedAccountIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case AzureInvitedAccountIdentifier:
		return a.UserPrincipalName == other.UserPrincipalName
	default:
		return false
	}
}

// AzureInvitedAccountFields has a lot that we don't actually directly specify when inviting
// a user. That's okay, because after we invite the user we update the account fields like
// AzureAccountFields / AzureAccountEngine -- in other words, these fields are here to keep
// the account updated after creation.
type AzureInvitedAccountFields struct {
	// Email controls the "mail" field of the user, which can technically be different from
	// the "userPrincipalName". For this account type, the userPrincipalName is the email and
	// should be the same as this field. We still have this as a field here so that Sherlock
	// will correct it should it get out of sync somehow (it is mutable in the UI).
	Email string
	// DisplayName is the human-readable name of the user
	DisplayName string
	// MailNickname is the prefix of the UPN before the @ symbol. It's here so Sherlock
	// can correct it if it gets mutated (and because we do have to set it during creation)
	MailNickname string
	// OtherMails is a list of other email addresses associated with the user. Critically,
	// this list must include the user's Broad email address, as this is how invites end up
	// reaching people.
	OtherMails []string
}

func (a AzureInvitedAccountFields) EqualTo(other intermediary_user.Fields) bool {
	switch other := other.(type) {
	case AzureInvitedAccountFields:
		return a.Email == other.Email &&
			a.DisplayName == other.DisplayName &&
			a.MailNickname == other.MailNickname &&
			reflect.DeepEqual(a.OtherMails, other.OtherMails)
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
	homeTenantEmailDomain      string
	inviteTenantIdentityDomain string
	userEmailDomainsToReplace  []string

	homeTenantClient   *msgraphsdk.GraphServiceClient
	inviteTenantClient *msgraphsdk.GraphServiceClient
}

func (a *AzureInvitedAccountEngine) Init(_ context.Context, k *koanf.Koanf) error {
	a.homeTenantEmailDomain = k.String("homeTenantEmailDomain")
	a.inviteTenantIdentityDomain = k.String("inviteTenantIdentityDomain")
	a.userEmailDomainsToReplace = k.Strings("userEmailDomainsToReplace")

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

	headers := abstractions.NewRequestHeaders()
	configuration := &users.UsersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Select: []string{"userPrincipalName", "accountEnabled", "mail", "displayName", "mailNickname", "otherMails"},
			// Since this is a B2C tenant, we can't do fancy filter things to also check the #EXT# thingy in the
			// user principal name. The creation type does a very good job of cutting down the response so we can
			// safely check the suffix as we iterate.
			Filter: utils.PointerTo("creationType eq 'Invitation'"),
			Top:    utils.PointerTo[int32](25),
		},
	}

	result, err := a.inviteTenantClient.Users().Get(ctx, configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to get invited users: %w", err)
	}

	pageIterator, err := msgraphgocore.NewPageIterator[graphmodels.Userable](result,
		a.inviteTenantClient.GetAdapter(), graphmodels.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create page iterator for invited users: %w", err)
	}
	pageIterator.SetHeaders(headers)

	err = pageIterator.Iterate(ctx, func(pageItem graphmodels.Userable) bool {
		if userPrincipalName := pageItem.GetUserPrincipalName(); userPrincipalName != nil &&
			strings.HasSuffix(*userPrincipalName, fmt.Sprintf("%s#EXT#@%s", a.homeTenantEmailDomain, a.inviteTenantIdentityDomain)) {
			var fields AzureInvitedAccountFields
			if mail := pageItem.GetMail(); mail != nil {
				fields.Email = *mail
			}
			if displayName := pageItem.GetDisplayName(); displayName != nil {
				fields.DisplayName = *displayName
			}
			if mailNickname := pageItem.GetMailNickname(); mailNickname != nil {
				fields.MailNickname = *mailNickname
			}
			if otherMails := pageItem.GetOtherMails(); otherMails != nil {
				fields.OtherMails = otherMails
			}
			currentState = append(currentState, intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields]{
				Identifier: AzureInvitedAccountIdentifier{UserPrincipalName: *userPrincipalName},
				Fields:     fields,
			})
		}
		return true
	})
	return currentState, err
}

func (a *AzureInvitedAccountEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields])
	for sherlockUserID, roleAssignment := range roleAssignments {
		// We explicitly aren't calling roleAssignment.IsActive() here. *This* propagator actually doesn't care about
		// suspension! We don't have a notion of suspending an invited account, but we actually don't delete them either!
		// We rely on the propagator that manages the home tenant identity to disable the user there, which suspends their
		// ability to log in here too.
		// We choose to still propagate the user here because we want to keep the user's name up to date in the invite tenant.

		email := utils.SubstituteSuffix(roleAssignment.User.Email, a.userEmailDomainsToReplace, a.homeTenantEmailDomain)
		if !strings.HasSuffix(email, a.homeTenantEmailDomain) {
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
					// The user principal name here is the UPN *from the home tenant*. For how we've got these
					// tenants set up, it's the user's home email address. We'll need to format out our actual
					// UPN as it'll look in the invite tenant.
					upn := fmt.Sprintf("%s#EXT#@%s", strings.ReplaceAll(*userPrincipalName, "@", "_"), a.inviteTenantIdentityDomain)
					desiredState[sherlockUserID] = intermediary_user.IntermediaryUser[AzureInvitedAccountIdentifier, AzureInvitedAccountFields]{
						Identifier: AzureInvitedAccountIdentifier{
							UserPrincipalName: upn,
						},
						Fields: AzureInvitedAccountFields{
							Email:        email,
							DisplayName:  roleAssignment.User.NameOrUsername(),
							MailNickname: strings.Split(upn, "@")[0],
							OtherMails:   []string{roleAssignment.User.Email},
						},
					}
				}
			}
		}
	}
	return desiredState, nil
}

func (a *AzureInvitedAccountEngine) Add(ctx context.Context, _ bool, identifier AzureInvitedAccountIdentifier, fields AzureInvitedAccountFields) (string, error) {
	inviteMessageBody, identifyingString, err := a.inviteMessageBody(identifier)
	if err != nil {
		return "", err
	}

	body := graphmodels.NewInvitation()
	body.SetInvitedUserEmailAddress(utils.PointerTo(fields.Email))
	body.SetInviteRedirectUrl(utils.PointerTo("https://portal.azure.com"))
	body.SetInvitedUserType(utils.PointerTo("member"))
	body.SetInvitedUserDisplayName(utils.PointerTo(fields.DisplayName))
	body.SetSendInvitationMessage(utils.PointerTo(true))
	invitedUserMessageInfo := graphmodels.NewInvitedUserMessageInfo()
	invitedUserMessageInfo.SetCustomizedMessageBody(utils.PointerTo(inviteMessageBody))
	body.SetInvitedUserMessageInfo(invitedUserMessageInfo)
	response, err := a.inviteTenantClient.Invitations().Post(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to invite %s: %w", identifier.UserPrincipalName, err)
	} else if response.GetInvitedUser() == nil || response.GetInvitedUser().GetId() == nil {
		return "", fmt.Errorf("failed to invite %s: no user ID returned", identifier.UserPrincipalName)
	}
	// Now we have to mutate the user that just got created to set their otherEmails field.
	// This is key for making sure the invite email goes to the BI email.
	postCreationEditBody := graphmodels.NewUser()
	postCreationEditBody.SetOtherMails(fields.OtherMails)
	_, err = a.inviteTenantClient.Users().ByUserId(*response.GetInvitedUser().GetId()).Patch(ctx, postCreationEditBody, nil)
	if err != nil {
		return "", fmt.Errorf("failed to set otherMails for newly invited user %s: %w", identifier.UserPrincipalName, err)
	}

	return fmt.Sprintf("invited %s (invite email sent with identifying string `%s`)", identifier.UserPrincipalName, identifyingString), nil
}

func (a *AzureInvitedAccountEngine) inviteMessageBody(identifier AzureInvitedAccountIdentifier) (inviteMessageBody string, identifyingString string, err error) {
	randomBytes := make([]byte, 8)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate random identifying string for inviting %s: %w", identifier.UserPrincipalName, err)
	}
	identifyingString = hex.EncodeToString(randomBytes)

	inviteMessageBody = "This invitation has been generated by the DSP DevOps platform via Microsoft Graph API. " +
		fmt.Sprintf("This invitation is meant to grant your %s Microsoft account access to %s. ", identifier.UserPrincipalName, a.inviteTenantIdentityDomain) +
		"You should reach out to DevOps to confirm the origin of this message before clicking the link. " +
		fmt.Sprintf("They can match this message to a security event with this identifying string: %s. ", identifyingString)

	return inviteMessageBody, identifyingString, nil
}

func (a *AzureInvitedAccountEngine) Update(ctx context.Context, _ bool, identifier AzureInvitedAccountIdentifier, oldFields AzureInvitedAccountFields, newFields AzureInvitedAccountFields) (string, error) {
	// We can't update an invitation (an email has already been sent), but if the fields are different that means fields (and thus the user)
	// already exist on the remote.
	body := graphmodels.NewUser()
	body.SetMail(utils.PointerTo(newFields.Email))
	body.SetDisplayName(utils.PointerTo(newFields.DisplayName))
	body.SetMailNickname(utils.PointerTo(newFields.MailNickname))
	body.SetOtherMails(newFields.OtherMails)
	_, err := a.inviteTenantClient.Users().ByUserId(identifier.UserPrincipalName).Patch(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to update user %s (%s): %w", identifier.UserPrincipalName, a.describeDiff(oldFields, newFields), err)
	} else {
		return fmt.Sprintf("updated user %s (%s)", identifier.UserPrincipalName, a.describeDiff(oldFields, newFields)), nil
	}
}

func (a *AzureInvitedAccountEngine) describeDiff(oldFields AzureInvitedAccountFields, newFields AzureInvitedAccountFields) string {
	if oldFields.Email != newFields.Email || oldFields.MailNickname != newFields.MailNickname || !reflect.DeepEqual(oldFields.OtherMails, newFields.OtherMails) {
		return "update account email info" // This is really, *really* unlikely to happen but we'll at least handle it
	} else if oldFields.DisplayName != newFields.DisplayName {
		return fmt.Sprintf("update display name from `%s` to `%s`", oldFields.DisplayName, newFields.DisplayName)
	} else {
		return "no changes"
	}
}

func (a *AzureInvitedAccountEngine) Remove(_ context.Context, _ bool, _ AzureInvitedAccountIdentifier) (string, error) {
	return "", fmt.Errorf("%T.Remove not implemented, %T.MayConsiderAsAlreadyRemoved should always return true", a, AzureInvitedAccountFields{})
}

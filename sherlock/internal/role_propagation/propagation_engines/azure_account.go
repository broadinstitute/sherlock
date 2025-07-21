package propagation_engines

import (
	"context"
	"fmt"
	"reflect"
	"strings"

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
	"github.com/sethvargo/go-password/password"
)

type AzureAccountIdentifier struct {
	UserPrincipalName string `koanf:"userPrincipalName"`
}

func (a AzureAccountIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case AzureAccountIdentifier:
		return a.UserPrincipalName == other.UserPrincipalName
	default:
		return false
	}
}

type AzureAccountFields struct {
	// AccountEnabled controls whether this account can be signed in to (to access any tenant,
	// home or invited)
	AccountEnabled bool
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

func (a AzureAccountFields) EqualTo(other intermediary_user.Fields) bool {
	switch other := other.(type) {
	case AzureAccountFields:
		return a.AccountEnabled == other.AccountEnabled &&
			a.Email == other.Email &&
			a.DisplayName == other.DisplayName &&
			a.MailNickname == other.MailNickname &&
			reflect.DeepEqual(a.OtherMails, other.OtherMails)
	default:
		return false
	}
}

// MayConsiderAsAlreadyRemoved returns true if the account is disabled, since that's how we
// remove accounts (we suspend rather than delete accounts).
func (a AzureAccountFields) MayConsiderAsAlreadyRemoved() bool {
	return !a.AccountEnabled
}

var _ PropagationEngine[bool, AzureAccountIdentifier, AzureAccountFields] = &AzureAccountEngine{}

type AzureAccountEngine struct {
	tenantEmailSuffix          string
	userEmailSuffixesToReplace []string
	client                     *msgraphsdk.GraphServiceClient
	passwordGenerator          *password.Generator
}

func (a *AzureAccountEngine) Init(_ context.Context, k *koanf.Koanf) error {
	a.tenantEmailSuffix = k.String("tenantEmailSuffix")
	a.userEmailSuffixesToReplace = k.Strings("userEmailSuffixesToReplace")
	credentials, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		ClientID:      k.String("clientID"),
		TenantID:      k.String("tenantID"),
		TokenFilePath: k.String("tokenFilePath"),
	})
	if err != nil {
		return fmt.Errorf("failed to create workload identity credential: %w", err)
	}

	a.client, err = msgraphsdk.NewGraphServiceClientWithCredentials(credentials, nil)
	if err != nil {
		return fmt.Errorf("failed to create graph service client: %w", err)
	}

	a.passwordGenerator, err = password.NewGenerator(nil)
	if err != nil {
		return fmt.Errorf("failed to create password generator: %w", err)
	}

	return nil
}

func (a *AzureAccountEngine) LoadCurrentState(ctx context.Context, _ bool) ([]intermediary_user.IntermediaryUser[AzureAccountIdentifier, AzureAccountFields], error) {
	currentState := make([]intermediary_user.IntermediaryUser[AzureAccountIdentifier, AzureAccountFields], 0)

	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")
	configuration := &users.UsersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			// Select a non-standard set of fields
			Select: []string{"userPrincipalName", "accountEnabled", "mail", "displayName", "mailNickname", "otherMails"},
			// This filter is efficient but is an advanced query, so needs an eventual consistency level header
			// and the count query parameter
			// https://learn.microsoft.com/en-us/graph/aad-advanced-queries?tabs=go
			Filter: utils.PointerTo(fmt.Sprintf("endsWith(userPrincipalName, '%s')", a.tenantEmailSuffix)),
			// Count is needed because we're using filter. It provides the count in the response *in addition to*
			// the rest of the response. It's a bit weird but I think it's to help guide you towards debugging
			// issues with the eventual consistency level. Since Sherlock is doing a reconciliation loop of its
			// own I'm making a blanket assumption that eventual consistency on the part of a cloud provider is
			// okay.
			Count: utils.PointerTo(true),
			// We want to get all the users, so interestingly we need to explicitly say we want just the first
			// N, and then that basically sets our page size for the page iterator we use below. At the time of
			// writing this, when you don't set this you seem to get the first 30 entries only, but with either
			// no @odata.nextLink property or one that the page iterator below doesn't understand -- and this
			// behavior doesn't really seem documented. The docs on paging do mention that the API sometimes has
			// a default page size, but it implies that an @odata.nextLink property will be present.
			// Anyway, if you remove just this one parameter, even if you do the page iterator below you won't
			// actually end up iterating over all the users.
			// The value of 25 is arbitrary and doesn't really matter since we end up iterating over everything.
			// It was picked as a non-30 number to try to differentiate from the discovered default.
			Top: utils.PointerTo[int32](25),
		},
	}

	result, err := a.client.Users().Get(ctx, configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	pageIterator, err := msgraphgocore.NewPageIterator[graphmodels.Userable](result, a.client.GetAdapter(), graphmodels.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create page iterator for users: %w", err)
	}
	pageIterator.SetHeaders(headers)

	err = pageIterator.Iterate(ctx, func(pageItem graphmodels.Userable) bool {
		if userPrincipalName := pageItem.GetUserPrincipalName(); userPrincipalName != nil && strings.HasSuffix(*userPrincipalName, a.tenantEmailSuffix) {
			var fields AzureAccountFields
			if accountEnabled := pageItem.GetAccountEnabled(); accountEnabled != nil {
				fields.AccountEnabled = *accountEnabled
			}
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
			currentState = append(currentState, intermediary_user.IntermediaryUser[AzureAccountIdentifier, AzureAccountFields]{
				Identifier: AzureAccountIdentifier{UserPrincipalName: *userPrincipalName},
				Fields:     fields,
			})
		}
		return true
	})
	return currentState, err
}

func (a *AzureAccountEngine) GenerateDesiredState(_ context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[AzureAccountIdentifier, AzureAccountFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[AzureAccountIdentifier, AzureAccountFields])
	for sherlockUserID, roleAssignment := range roleAssignments {
		email := utils.SubstituteSuffix(roleAssignment.User.Email, a.userEmailSuffixesToReplace, a.tenantEmailSuffix)
		if !strings.HasSuffix(email, a.tenantEmailSuffix) {
			// We can short circuit here because we're only responsible for creating accounts with the
			// given email suffix
			continue
		}
		desiredState[sherlockUserID] = intermediary_user.IntermediaryUser[AzureAccountIdentifier, AzureAccountFields]{
			Identifier: AzureAccountIdentifier{UserPrincipalName: email},
			Fields: AzureAccountFields{
				AccountEnabled: roleAssignment.IsActive(),
				Email:          email,
				DisplayName:    roleAssignment.User.NameOrUsername(),
				MailNickname:   strings.Split(email, "@")[0],
				OtherMails:     []string{roleAssignment.User.Email},
			},
		}
	}
	return desiredState, nil
}

func (a *AzureAccountEngine) Add(ctx context.Context, _ bool, identifier AzureAccountIdentifier, fields AzureAccountFields) (string, error) {
	randomPassword, err := a.passwordGenerator.Generate(256, 32, 32, false, true)
	if err != nil {
		return "", fmt.Errorf("failed to generate random password: %w", err)
	}

	body := graphmodels.NewUser()
	body.SetUserPrincipalName(utils.PointerTo(identifier.UserPrincipalName))
	body.SetMail(utils.PointerTo(fields.Email))
	body.SetDisplayName(utils.PointerTo(fields.DisplayName))
	body.SetMailNickname(utils.PointerTo(fields.MailNickname))
	body.SetAccountEnabled(utils.PointerTo(fields.AccountEnabled))
	body.SetOtherMails(fields.OtherMails)
	passwordProfile := graphmodels.NewPasswordProfile()
	passwordProfile.SetForceChangePasswordNextSignIn(utils.PointerTo(true))
	// We don't send this password anywhere, forcing the user to go through the password reset flow.
	passwordProfile.SetPassword(utils.PointerTo(randomPassword))
	body.SetPasswordProfile(passwordProfile)

	_, err = a.client.Users().Post(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create user %s: %w", identifier.UserPrincipalName, err)
	} else {
		return fmt.Sprintf("created user %s", identifier.UserPrincipalName), nil
	}
}

func (a *AzureAccountEngine) Update(ctx context.Context, _ bool, identifier AzureAccountIdentifier, oldFields AzureAccountFields, newFields AzureAccountFields) (string, error) {
	body := graphmodels.NewUser()
	body.SetMail(utils.PointerTo(newFields.Email))
	body.SetDisplayName(utils.PointerTo(newFields.DisplayName))
	body.SetMailNickname(utils.PointerTo(newFields.MailNickname))
	body.SetAccountEnabled(utils.PointerTo(newFields.AccountEnabled))
	body.SetOtherMails(newFields.OtherMails)
	_, err := a.client.Users().ByUserId(identifier.UserPrincipalName).Patch(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to update user %s (%s): %w", identifier.UserPrincipalName, a.describeDiff(oldFields, newFields), err)
	} else {
		return fmt.Sprintf("updated user %s (%s)", identifier.UserPrincipalName, a.describeDiff(oldFields, newFields)), nil
	}
}

func (a *AzureAccountEngine) describeDiff(oldFields AzureAccountFields, newFields AzureAccountFields) string {
	if oldFields.AccountEnabled && !newFields.AccountEnabled {
		return "disable account"
	} else if !oldFields.AccountEnabled && newFields.AccountEnabled {
		return "enable account"
	} else if oldFields.Email != newFields.Email || oldFields.MailNickname != newFields.MailNickname || !reflect.DeepEqual(oldFields.OtherMails, newFields.OtherMails) {
		return "update account email info" // This is really, *really* unlikely to happen but we'll at least handle it
	} else if oldFields.DisplayName != newFields.DisplayName {
		return fmt.Sprintf("update display name from `%s` to `%s`", oldFields.DisplayName, newFields.DisplayName)
	} else {
		return "no changes"
	}
}

func (a *AzureAccountEngine) Remove(ctx context.Context, _ bool, identifier AzureAccountIdentifier) (string, error) {
	body := graphmodels.NewUser()
	body.SetAccountEnabled(utils.PointerTo(false))
	_, err := a.client.Users().ByUserId(identifier.UserPrincipalName).Patch(ctx, body, nil)
	if err != nil {
		return "", fmt.Errorf("failed to disable user %s: %w", identifier.UserPrincipalName, err)
	} else {
		return fmt.Sprintf("disabled user %s", identifier.UserPrincipalName), nil
	}
}

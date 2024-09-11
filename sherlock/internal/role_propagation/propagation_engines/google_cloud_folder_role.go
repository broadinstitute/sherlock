package propagation_engines

import (
	"cloud.google.com/go/iam/apiv1/iampb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/knadh/koanf"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"strings"
)

const GoogleCloudOwnerRole = "roles/owner"

type GoogleCloudFolderRoleIdentifier struct {
	Email string `koanf:"email"`
}

func (g GoogleCloudFolderRoleIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case GoogleCloudFolderRoleIdentifier:
		return g.Email == other.Email
	default:
		return false
	}
}

type GoogleCloudFolderRoleFields struct{}

func (g GoogleCloudFolderRoleFields) EqualTo(other intermediary_user.Fields) bool {
	switch other.(type) {
	case GoogleCloudFolderRoleFields:
		return true
	default:
		return false
	}
}

var _ PropagationEngine[string, GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields] = &GoogleCloudFolderRoleEngine{}

// GoogleCloudFolderRoleEngine can grant roles outside the workspace domain. The workspace domain and
// admin directory access is used to determine what users exist, not to restrict what those users can be
// granted access to.
type GoogleCloudFolderRoleEngine struct {
	// Role is defined in Sherlock's source code where this engine is instantiated, not in configuration.
	// This is because it helps avoid a foot-gun with accidentally misconfiguring the engine. When this
	// engine is pointed at a folder, it wants to "own" all user-level bindings for this role. It will remove
	// user-level bindings that don't correlate to Sherlock role assignments. If we allowed easy configuration
	// of the role, it would be easy to accidentally cause problems by pointing this engine at a role that some
	// other system also wanted to manage or relied on.
	Role string

	workspaceDomain            string
	userEmailSuffixesToReplace []string
	adminService               *admin.Service
	foldersClient              *resourcemanager.FoldersClient
}

func (g *GoogleCloudFolderRoleEngine) Init(ctx context.Context, k *koanf.Koanf) error {
	if g.Role == "" {
		return fmt.Errorf("role must be set")
	}

	g.workspaceDomain = k.String("workspaceDomain")
	g.userEmailSuffixesToReplace = k.Strings("userEmailSuffixesToReplace")
	var err error
	g.adminService, err = admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryUserReadonlyScope))
	if err != nil {
		return err
	}
	g.foldersClient, err = resourcemanager.NewFoldersClient(ctx)
	return err
}

func (g *GoogleCloudFolderRoleEngine) LoadCurrentState(ctx context.Context, grant string) ([]intermediary_user.IntermediaryUser[GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields], error) {
	policy, err := g.foldersClient.GetIamPolicy(ctx, &iampb.GetIamPolicyRequest{
		Resource: grant,
	})
	if err != nil {
		return nil, err
	} else {
		currentState := make([]intermediary_user.IntermediaryUser[GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields], 0)
		for _, binding := range policy.Bindings {
			if binding.Role == GoogleCloudOwnerRole {
				for _, member := range binding.Members {
					if strings.HasPrefix(member, "user:") {
						currentState = append(currentState, intermediary_user.IntermediaryUser[GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields]{
							Identifier: GoogleCloudFolderRoleIdentifier{Email: strings.TrimPrefix(member, "user:")},
							Fields:     GoogleCloudFolderRoleFields{},
						})
					}
				}
			}
		}
		return currentState, nil
	}
}

func (g *GoogleCloudFolderRoleEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields])
	for id, roleAssignment := range roleAssignments {
		if !roleAssignment.IsActive() {
			// There's no concept of a suspended group member, so we just exclude any non-active users
			continue
		}

		email := utils.SubstituteSuffix(roleAssignment.User.Email, g.userEmailSuffixesToReplace, "@"+g.workspaceDomain)
		if !strings.HasSuffix(email, "@"+g.workspaceDomain) {
			// We can short-circuit here, we know that the user is not in the workspace domain so we won't bother looking
			continue
		}

		err := g.adminService.Users.List().
			Domain(g.workspaceDomain).
			Query("email="+email).
			Fields("users(primaryEmail)").
			MaxResults(1).
			Pages(ctx, func(workspaceUsers *admin.Users) error {
				for _, workspaceUser := range workspaceUsers.Users {
					if workspaceUser.PrimaryEmail == email {
						desiredState[id] = intermediary_user.IntermediaryUser[GoogleCloudFolderRoleIdentifier, GoogleCloudFolderRoleFields]{
							Identifier: GoogleCloudFolderRoleIdentifier{Email: email},
							Fields:     GoogleCloudFolderRoleFields{},
						}
					}
				}
				return nil
			})
		if err != nil {
			return nil, err
		}
	}
	return desiredState, nil
}

func (g *GoogleCloudFolderRoleEngine) Add(ctx context.Context, grant string, identifier GoogleCloudFolderRoleIdentifier, fields GoogleCloudFolderRoleFields) (string, error) {
	policy, err := g.foldersClient.GetIamPolicy(ctx, &iampb.GetIamPolicyRequest{
		Resource: grant,
	})
	if err != nil {
		return "", err
	}

	for _, binding := range policy.Bindings {
		if binding.Role == GoogleCloudOwnerRole && binding.Condition == nil {
			binding.Members = append(binding.Members, "user:"+identifier.Email)
		}
	}

	_, err = g.foldersClient.SetIamPolicy(ctx, &iampb.SetIamPolicyRequest{
		Resource: grant,
		Policy:   policy,
	})
	if err != nil {
		return "", fmt.Errorf("failed to grant %s %s on %s: %w", identifier.Email, GoogleCloudOwnerRole, grant, err)
	} else {
		return fmt.Sprintf("granted %s %s on %s", identifier.Email, GoogleCloudOwnerRole, grant), nil
	}
}

func (g *GoogleCloudFolderRoleEngine) Update(_ context.Context, _ string, _ GoogleCloudFolderRoleIdentifier, _ GoogleCloudFolderRoleFields, _ GoogleCloudFolderRoleFields) (string, error) {
	return "", fmt.Errorf("%T.Update not implemented; %T.EqualTo should always equal true", g, GoogleCloudFolderRoleFields{})
}

func (g *GoogleCloudFolderRoleEngine) Remove(ctx context.Context, grant string, identifier GoogleCloudFolderRoleIdentifier) (string, error) {
	policy, err := g.foldersClient.GetIamPolicy(ctx, &iampb.GetIamPolicyRequest{
		Resource: grant,
	})
	if err != nil {
		return "", err
	}

	for _, binding := range policy.Bindings {
		if binding.Role == GoogleCloudOwnerRole && binding.Condition == nil {
			members := make([]string, 0, len(binding.Members)-1)
			for _, member := range binding.Members {
				if member != "user:"+identifier.Email {
					members = append(members, member)
				}
			}
			binding.Members = members
		}
	}

	_, err = g.foldersClient.SetIamPolicy(ctx, &iampb.SetIamPolicyRequest{
		Resource: grant,
		Policy:   policy,
	})
	if err != nil {
		return "", fmt.Errorf("failed to revoke %s's %s on %s: %w", identifier.Email, GoogleCloudOwnerRole, grant, err)
	} else {
		return fmt.Sprintf("revoke %s's %s on %s", identifier.Email, GoogleCloudOwnerRole, grant), nil
	}
}

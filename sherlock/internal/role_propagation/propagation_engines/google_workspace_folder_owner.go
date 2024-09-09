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

// We hardcode the owner role here because it avoids a foot-gun with configuring this propagation engine.
// When this propagation engine is pointed at a folder, it wants to "own" all user-level bindings for this
// role. If we allowed configuration of this, it would be easier to accidentally cause problems by
// pointing this engine at a role that some other system also wanted to manage.
const ownerRole = "roles/owner"

type GoogleWorkspaceFolderOwnerIdentifier struct {
	Email string `koanf:"email"`
}

func (g GoogleWorkspaceFolderOwnerIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case GoogleWorkspaceFolderOwnerIdentifier:
		return g.Email == other.Email
	default:
		return false
	}
}

type GoogleWorkspaceFolderOwnerFields struct{}

func (g GoogleWorkspaceFolderOwnerFields) EqualTo(other intermediary_user.Fields) bool {
	switch other.(type) {
	case GoogleWorkspaceFolderOwnerFields:
		return true
	default:
		return false
	}
}

var _ PropagationEngine[string, GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields] = &GoogleWorkspaceFolderOwnerEngine{}

// GoogleWorkspaceFolderOwnerEngine can grant roles outside the workspace domain. The workspace domain and
// admin directory access is used to determine what users exist, not to restrict what those users can be
// granted access to.
type GoogleWorkspaceFolderOwnerEngine struct {
	workspaceDomain            string
	userEmailSuffixesToReplace []string
	adminService               *admin.Service
	foldersClient              *resourcemanager.FoldersClient
}

func (g *GoogleWorkspaceFolderOwnerEngine) Init(ctx context.Context, k *koanf.Koanf) error {
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

func (g *GoogleWorkspaceFolderOwnerEngine) LoadCurrentState(ctx context.Context, grant string) ([]intermediary_user.IntermediaryUser[GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields], error) {
	policy, err := g.foldersClient.GetIamPolicy(ctx, &iampb.GetIamPolicyRequest{
		Resource: grant,
	})
	if err != nil {
		return nil, err
	} else {
		currentState := make([]intermediary_user.IntermediaryUser[GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields], 0)
		for _, binding := range policy.Bindings {
			if binding.Role == ownerRole {
				for _, member := range binding.Members {
					if strings.HasPrefix(member, "user:") {
						currentState = append(currentState, intermediary_user.IntermediaryUser[GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields]{
							Identifier: GoogleWorkspaceFolderOwnerIdentifier{Email: strings.TrimPrefix(member, "user:")},
							Fields:     GoogleWorkspaceFolderOwnerFields{},
						})
					}
				}
			}
		}
		return currentState, nil
	}
}

func (g *GoogleWorkspaceFolderOwnerEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields])
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
						desiredState[id] = intermediary_user.IntermediaryUser[GoogleWorkspaceFolderOwnerIdentifier, GoogleWorkspaceFolderOwnerFields]{
							Identifier: GoogleWorkspaceFolderOwnerIdentifier{Email: email},
							Fields:     GoogleWorkspaceFolderOwnerFields{},
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

func (g *GoogleWorkspaceFolderOwnerEngine) Add(ctx context.Context, grant string, identifier GoogleWorkspaceFolderOwnerIdentifier, fields GoogleWorkspaceFolderOwnerFields) (string, error) {
	policy, err := g.foldersClient.GetIamPolicy(ctx, &iampb.GetIamPolicyRequest{
		Resource: grant,
	})
	if err != nil {
		return "", err
	}

	for _, binding := range policy.Bindings {
		if binding.Role == ownerRole && binding.Condition == nil {
			binding.Members = append(binding.Members, "user:"+identifier.Email)
		}
	}

	_, err = g.foldersClient.SetIamPolicy(ctx, &iampb.SetIamPolicyRequest{
		Resource: grant,
		Policy:   policy,
	})
	if err != nil {
		return "", fmt.Errorf("failed to grant %s %s on %s: %w", identifier.Email, ownerRole, grant, err)
	} else {
		return fmt.Sprintf("granted %s %s on %s", identifier.Email, ownerRole, grant), nil
	}
}

func (g *GoogleWorkspaceFolderOwnerEngine) Update(_ context.Context, _ string, _ GoogleWorkspaceFolderOwnerIdentifier, _ GoogleWorkspaceFolderOwnerFields, _ GoogleWorkspaceFolderOwnerFields) (string, error) {
	return "", fmt.Errorf("%T.Update not implemented; %T.EqualTo should always equal true", g, GoogleWorkspaceFolderOwnerFields{})
}

func (g *GoogleWorkspaceFolderOwnerEngine) Remove(ctx context.Context, grant string, identifier GoogleWorkspaceFolderOwnerIdentifier) (string, error) {
	policy, err := g.foldersClient.GetIamPolicy(ctx, &iampb.GetIamPolicyRequest{
		Resource: grant,
	})
	if err != nil {
		return "", err
	}

	for _, binding := range policy.Bindings {
		if binding.Role == ownerRole && binding.Condition == nil {
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
		return "", fmt.Errorf("failed to revoke %s's %s on %s: %w", identifier.Email, ownerRole, grant, err)
	} else {
		return fmt.Sprintf("revoke %s's %s on %s", identifier.Email, ownerRole, grant), nil
	}
}

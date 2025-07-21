package propagation_engines

import (
	"context"
	"fmt"
	"strings"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/knadh/koanf"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type GoogleWorkspaceGroupIdentifier struct {
	Email string `koanf:"email"`
}

func (g GoogleWorkspaceGroupIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	switch other := other.(type) {
	case GoogleWorkspaceGroupIdentifier:
		return g.Email == other.Email
	default:
		return false
	}
}

type GoogleWorkspaceGroupFields struct{}

func (g GoogleWorkspaceGroupFields) EqualTo(other intermediary_user.Fields) bool {
	switch other.(type) {
	case GoogleWorkspaceGroupFields:
		return true
	default:
		return false
	}
}

type GoogleWorkspaceGroupEngine struct {
	workspaceDomain            string
	userEmailSuffixesToReplace []string
	adminService               *admin.Service
}

func (g *GoogleWorkspaceGroupEngine) Init(ctx context.Context, k *koanf.Koanf) error {
	g.workspaceDomain = k.String("workspaceDomain")
	g.userEmailSuffixesToReplace = k.Strings("userEmailSuffixesToReplace")
	var err error
	g.adminService, err = admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryUserScope, admin.AdminDirectoryGroupMemberScope))
	return err
}

func (g *GoogleWorkspaceGroupEngine) LoadCurrentState(ctx context.Context, grant string) ([]intermediary_user.IntermediaryUser[GoogleWorkspaceGroupIdentifier, GoogleWorkspaceGroupFields], error) {
	currentState := make([]intermediary_user.IntermediaryUser[GoogleWorkspaceGroupIdentifier, GoogleWorkspaceGroupFields], 0)
	err := g.adminService.Members.List(grant).Pages(ctx, func(members *admin.Members) error {
		for _, member := range members.Members {
			currentState = append(currentState, intermediary_user.IntermediaryUser[GoogleWorkspaceGroupIdentifier, GoogleWorkspaceGroupFields]{
				Identifier: GoogleWorkspaceGroupIdentifier{Email: member.Email},
				Fields:     GoogleWorkspaceGroupFields{},
			})
		}
		return nil
	})
	return currentState, err
}

func (g *GoogleWorkspaceGroupEngine) GenerateDesiredState(ctx context.Context, roleAssignments map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[GoogleWorkspaceGroupIdentifier, GoogleWorkspaceGroupFields], error) {
	desiredState := make(map[uint]intermediary_user.IntermediaryUser[GoogleWorkspaceGroupIdentifier, GoogleWorkspaceGroupFields])
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
						desiredState[id] = intermediary_user.IntermediaryUser[GoogleWorkspaceGroupIdentifier, GoogleWorkspaceGroupFields]{
							Identifier: GoogleWorkspaceGroupIdentifier{Email: email},
							Fields:     GoogleWorkspaceGroupFields{},
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

func (g *GoogleWorkspaceGroupEngine) Add(ctx context.Context, grant string, identifier GoogleWorkspaceGroupIdentifier, _ GoogleWorkspaceGroupFields) (string, error) {
	response, err := g.adminService.Members.Insert(grant, &admin.Member{
		Role:  "MEMBER",
		Email: identifier.Email,
	}).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("failed to add %s to %s: %w", identifier.Email, grant, err)
	} else {
		return fmt.Sprintf("added %s to %s", response.Email, grant), nil
	}
}

func (g *GoogleWorkspaceGroupEngine) Update(_ context.Context, _ string, _ GoogleWorkspaceGroupIdentifier, _ GoogleWorkspaceGroupFields, _ GoogleWorkspaceGroupFields) (string, error) {
	return "", fmt.Errorf("%T.Update not implemented; %T.EqualTo should always equal true", g, GoogleWorkspaceGroupFields{})
}

func (g *GoogleWorkspaceGroupEngine) Remove(ctx context.Context, grant string, identifier GoogleWorkspaceGroupIdentifier) (string, error) {
	err := g.adminService.Members.Delete(grant, identifier.Email).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("failed to remove %s from %s: %w", identifier.Email, grant, err)
	} else {
		return fmt.Sprintf("removed %s from %s", identifier.Email, grant), nil
	}
}

package auth

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"strings"
)

type User struct {
	AuthenticatedEmail      string            `json:"authenticatedEmail"`
	MatchedFirecloudAccount *FirecloudAccount `json:"matchedFirecloudAccount,omitempty"`
}

type FirecloudAccount struct {
	Email               string                    `json:"email"`
	AcceptedGoogleTerms bool                      `json:"acceptedGoogleTerms"`
	EnrolledIn2fa       bool                      `json:"enrolledIn2Fa"`
	Suspended           bool                      `json:"suspended"`
	Archived            bool                      `json:"archived"`
	SuspensionReason    string                    `json:"suspensionReason,omitempty"`
	Groups              *FirecloudGroupMembership `json:"groups"`
}

func (f *FirecloudAccount) parseWorkspaceUser(user *admin.User) {
	f.Email = user.PrimaryEmail
	f.AcceptedGoogleTerms = user.AgreedToTerms
	f.EnrolledIn2fa = user.IsEnrolledIn2Sv
	f.Suspended = user.Suspended
	f.Archived = user.Archived
	f.SuspensionReason = user.SuspensionReason
}

type FirecloudGroupMembership struct {
	FcAdmins               bool `json:"fc-admins"`
	FirecloudProjectOwners bool `json:"firecloud-project-owners"`
}

func (u *User) Username() string {
	return strings.Split(u.AuthenticatedEmail, "@")[0]
}

func (u *User) isKnownSuitable() bool {
	return u.MatchedFirecloudAccount != nil &&
		u.MatchedFirecloudAccount.AcceptedGoogleTerms &&
		u.MatchedFirecloudAccount.EnrolledIn2fa &&
		!u.MatchedFirecloudAccount.Suspended &&
		!u.MatchedFirecloudAccount.Archived &&
		u.MatchedFirecloudAccount.Groups.FcAdmins &&
		u.MatchedFirecloudAccount.Groups.FirecloudProjectOwners
}

func (u *User) describeSuitability() string {
	if u.MatchedFirecloudAccount == nil {
		return fmt.Sprintf("%s is not known suitable as a matching Firecloud account wasn't found", u.Username())
	} else {
		var problems []string
		if !u.MatchedFirecloudAccount.AcceptedGoogleTerms {
			problems = append(problems, "user hasn't accepted Google Workspace terms (suggesting they've never logged in)")
		}
		if !u.MatchedFirecloudAccount.EnrolledIn2fa {
			problems = append(problems, "user hasn't enrolled in two-factor authentication")
		}
		if u.MatchedFirecloudAccount.Suspended {
			if u.MatchedFirecloudAccount.SuspensionReason == "" {
				problems = append(problems, "user is currently suspended (no reason given)")
			} else {
				problems = append(problems, fmt.Sprintf("user is currently suspended (%s)", u.MatchedFirecloudAccount.SuspensionReason))
			}
		}
		if u.MatchedFirecloudAccount.Archived {
			problems = append(problems, "user is currently archived")
		}

		if !u.MatchedFirecloudAccount.Groups.FcAdmins {
			problems = append(problems, fmt.Sprintf("user is not in the %s group", firecloudGroups.FcAdmins))
		}
		if !u.MatchedFirecloudAccount.Groups.FirecloudProjectOwners {
			problems = append(problems, fmt.Sprintf("user is not in the %s group", firecloudGroups.FirecloudProjectOwners))
		}

		if len(problems) > 0 {
			return fmt.Sprintf("%s may be suitable but the matching Firecloud account has issues: %s", u.Username(), strings.Join(problems, ", "))
		} else {
			return fmt.Sprintf("%s is known suitable", u.Username())
		}
	}
}

func (u *User) SuitableOrError() error {
	if u.isKnownSuitable() {
		return nil
	} else {
		log.Debug().Msgf("AUTH | %s might not be suitable, refreshing before denying", u.AuthenticatedEmail)
		adminService, err := admin.NewService(context.Background(), option.WithScopes(admin.AdminDirectoryUserReadonlyScope, admin.AdminDirectoryGroupMemberReadonlyScope))
		if err != nil {
			return fmt.Errorf("%s [Sherlock also failed to authenticate to Google to refresh this info: %v]", u.describeSuitability(), err)
		}

		var email string
		if u.MatchedFirecloudAccount == nil {
			email = emailToFirecloudEmail(u.AuthenticatedEmail)
			u.MatchedFirecloudAccount = &FirecloudAccount{}
		} else {
			email = u.MatchedFirecloudAccount.Email
		}

		workspaceUser, err := adminService.Users.Get(email).Do()
		if err != nil {
			return fmt.Errorf("%s [Sherlock also failed to get refreshed user info from Google Workspace: %v]", u.describeSuitability(), err)
		}
		u.MatchedFirecloudAccount.parseWorkspaceUser(workspaceUser)

		if u.MatchedFirecloudAccount.Groups == nil {
			u.MatchedFirecloudAccount.Groups = &FirecloudGroupMembership{}
		}

		fcAdminsMembership, err := adminService.Members.HasMember(firecloudGroups.FcAdmins, email).Do()
		if err != nil {
			return fmt.Errorf("%s [Sherlock also failed to refresh user's membership in %s: %v", u.describeSuitability(), firecloudGroups.FcAdmins, email)
		}
		u.MatchedFirecloudAccount.Groups.FcAdmins = fcAdminsMembership.IsMember

		firecloudProjectOwnersMembership, err := adminService.Members.HasMember(firecloudGroups.FirecloudProjectOwners, email).Do()
		if err != nil {
			return fmt.Errorf("%s [Sherlock also failed to refresh user's membership in %s: %v", u.describeSuitability(), firecloudGroups.FirecloudProjectOwners, email)
		}
		u.MatchedFirecloudAccount.Groups.FirecloudProjectOwners = firecloudProjectOwnersMembership.IsMember

		if u.isKnownSuitable() {
			return nil
		} else {
			return fmt.Errorf("%s", u.describeSuitability())
		}
	}
}

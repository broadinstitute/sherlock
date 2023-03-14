package auth_models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/config"
	"strings"
	"unicode"
)

type User struct {
	ID uint `json:"id" form:"id"`
	StoredControlledUserFields
	StoredMutableUserFields
	InferredUserFields
	AuthMethod AuthMethod `json:"authMethod" form:"authMethod"`
	Via        *User      `json:"via,omitempty" form:"-"`
}

type StoredControlledUserFields struct {
	Email          string  `json:"email" form:"email" gorm:"not null;default:null;unique"`
	GoogleID       string  `json:"googleID" form:"googleID" gorm:"not null;default:null;unique"`
	GithubUsername *string `json:"githubUsername" form:"githubUsername"`
	GithubID       *string `json:"githubID" form:"githubID"`
}

type StoredMutableUserFields struct {
	Name                   *string `json:"name,omitempty" form:"name"`
	NameInferredFromGithub *bool   `json:"nameInferredFromGithub" form:"nameInferredFromGithub"`
}

type InferredUserFields struct {
	MatchedFirecloudAccount *FirecloudAccount `json:"matchedFirecloudAccount,omitempty"`
	MatchedExtraPermissions *ExtraPermissions `json:"matchedExtraPermissions,omitempty"`
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

type FirecloudGroupMembership struct {
	FcAdmins               bool `json:"fc-admins"`
	FirecloudProjectOwners bool `json:"firecloud-project-owners"`
}

type ExtraPermissions struct {
	Suitable bool `json:"suitable"`
}

func (u *User) Username() string {
	return strings.Split(u.Email, "@")[0]
}

func (u *User) AlphaNumericHyphenatedUsername() string {
	var ret []rune
	for _, r := range u.Username() {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			ret = append(ret, r)
		} else if r == '.' || r == '-' || r == '_' {
			ret = append(ret, '-')
		}
	}
	return string(ret)
}

func (u *User) SuitableOrError() error {
	if u.isKnownSuitable() {
		return nil
	} else {
		return fmt.Errorf("%s", u.describeSuitability())
	}
}

// IsFromAuthMethod checks if this User is entirely derived from the provided auth method. This function is recursive on
// the "Via" User field.
func (u *User) IsFromAuthMethod(authMethod AuthMethod) bool {
	if u.AuthMethod != authMethod {
		return false
	} else if u.Via != nil {
		return u.Via.IsFromAuthMethod(authMethod)
	} else {
		return true
	}
}

func (u *User) isKnownSuitable() bool {
	return (u.MatchedFirecloudAccount != nil &&
		u.MatchedFirecloudAccount.AcceptedGoogleTerms &&
		u.MatchedFirecloudAccount.EnrolledIn2fa &&
		!u.MatchedFirecloudAccount.Suspended &&
		!u.MatchedFirecloudAccount.Archived &&
		u.MatchedFirecloudAccount.Groups.FcAdmins &&
		u.MatchedFirecloudAccount.Groups.FirecloudProjectOwners) ||
		(u.MatchedExtraPermissions != nil && u.MatchedExtraPermissions.Suitable)
}

func (u *User) describeSuitability() string {
	if u.MatchedExtraPermissions != nil && u.MatchedExtraPermissions.Suitable {
		return fmt.Sprintf("%s is known suitable via extra Sherlock-only permissions", u.Email)
	} else if u.MatchedFirecloudAccount == nil {
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

		if u.MatchedFirecloudAccount.Groups == nil {
			problems = append(problems, "user has no reported group information")
		} else {
			if !u.MatchedFirecloudAccount.Groups.FcAdmins {
				problems = append(problems, fmt.Sprintf("user is not in the %s group", config.Config.MustString("auth.firecloud.groups.fcAdmins")))
			}
			if !u.MatchedFirecloudAccount.Groups.FirecloudProjectOwners {
				problems = append(problems, fmt.Sprintf("user is not in the %s group", config.Config.MustString("auth.firecloud.groups.firecloudProjectOwners")))
			}
		}

		if len(problems) > 0 {
			return fmt.Sprintf("%s may be suitable but the matching Firecloud account has issues: %s", u.Username(), strings.Join(problems, ", "))
		} else {
			return fmt.Sprintf("%s is known suitable", u.Username())
		}
	}
}

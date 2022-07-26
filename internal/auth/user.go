package auth

import (
	"fmt"
	"strings"
)

type User struct {
	AuthenticatedEmail string
	suitabilityInfo    *suitabilityInfo
}

type suitabilityInfo struct {
	acceptedWorkspaceTos bool
	enrolledIn2fa        bool
	suspended            bool
	archived             bool
	suspensionReason     string
}

func (u *User) UsernameSlug() string {
	return strings.Split(u.AuthenticatedEmail, "@")[0]
}

func (u *User) EvaluateSuitability() bool {
	return u.suitabilityInfo != nil && u.suitabilityInfo.acceptedWorkspaceTos && u.suitabilityInfo.enrolledIn2fa && !u.suitabilityInfo.suspended && !u.suitabilityInfo.archived
}

func (u *User) DescribeSuitability() string {
	if u.suitabilityInfo == nil {
		return fmt.Sprintf("user %s did not have a matching suitable Firecloud account", u.UsernameSlug())
	} else {
		problems := []string{}
		if !u.suitabilityInfo.acceptedWorkspaceTos {
			problems = append(problems, "user has no accepted the Google Workspace Terms of Service")
		}
		if !u.suitabilityInfo.enrolledIn2fa {
			problems = append(problems, "user has not enrolled in two-factor authentication")
		}
		if u.suitabilityInfo.suspended {
			if u.suitabilityInfo.suspensionReason == "" {
				problems = append(problems, "user's account is suspended (no reason given)")
			} else {
				problems = append(problems, fmt.Sprintf("user's account is suspended (%s)", u.suitabilityInfo.suspensionReason))
			}
		}
		if u.suitabilityInfo.archived {
			problems = append(problems, "user's account is archived")
		}

		if len(problems) > 0 {
			return fmt.Sprintf("user's %s@firecloud.org account has problems preventing them from being currently suitable: %s", u.UsernameSlug(), strings.Join(problems, ", "))
		} else {
			return fmt.Sprintf("%s@firecloud.org is suitable", u.UsernameSlug())
		}
	}
}

func (u *User) SuitableOrError() error {
	if u.EvaluateSuitability() {
		return nil
	} else {
		return fmt.Errorf("user was not suitable: %s", u.DescribeSuitability())
	}
}

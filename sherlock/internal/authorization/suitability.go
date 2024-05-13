package authorization

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/local_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"strings"
)

type SuitabilitySource uint

const (
	NONE SuitabilitySource = iota
	FIRECLOUD
	CONFIG
)

type Suitability struct {
	// We go to the effort of internal fields with exported getters for the extra little bit of defensive
	// design, so that Suitability objects can't be modified by anything outside this package.
	suitable    bool
	description string
	source      SuitabilitySource
}

func (s *Suitability) Suitable() bool {
	return s.suitable
}

func (s *Suitability) Description() string {
	return s.description
}

func (s *Suitability) Source() SuitabilitySource {
	return s.source
}

func (s *Suitability) SuitableOrError() error {
	if !s.suitable || s.source == NONE {
		return fmt.Errorf("%s", s.description)
	} else {
		return nil
	}
}

// GetSuitabilityFor does what it says on the tin, but you probably don't need to call it.
// models.User has a .DeprecatedSuitability() method that you should use instead, it's more
// performant with less room for error (it calls this internally and caches the result).
func GetSuitabilityFor(email string) *Suitability {
	if suitability, present := cachedFirecloudSuitability[email]; present {
		return suitability
	} else if suitability, present = cachedConfigSuitability[email]; present {
		return suitability
	} else if config.Config.MustString("mode") == "debug" && (email == test_users.SuitableTestUserEmail || email == test_users.NonSuitableTestUserEmail) {
		// We don't gate respecting test users on the auth.createTestUsersInMiddleware config. In other words, if
		// we are in debug mode and we are queried about a test user being suitable, recognize the test user.
		// auth.createTestUsersInMiddleware just provides a shortcut so tests might not have to manually create
		// the test users themselves (but sometimes they want to).
		return &Suitability{
			suitable:    email == test_users.SuitableTestUserEmail,
			description: fmt.Sprintf("test user; email %s equal to suitable %s: %v", email, test_users.SuitableTestUserEmail, email == test_users.SuitableTestUserEmail),
			source:      CONFIG,
		}
	} else if config.Config.MustString("mode") == "debug" && email == local_user.LocalUserEmail {
		return &Suitability{
			suitable:    local_user.LocalUserSuitable,
			description: fmt.Sprintf("local user; suitable: %v", local_user.LocalUserSuitable),
			source:      CONFIG,
		}
	}

	// If we get here, the user isn't suitable.
	notSuitable := &Suitability{
		suitable:    false,
		description: fmt.Sprintf("user %s lacks production suitability", email),
		source:      NONE,
	}

	// However... we can detect a misconfiguration here, and save someone a bunch of debugging time.
	// Sherlock used to try to be smart and associate BI emails to FC emails internally. That wasn't
	// wrong, but a more robust way to do it is to record both the primary (FC) and recovery (BI)
	// emails on the account itself. This better handles edge cases around break-glass accounts,
	// where now Sherlock won't blindly assume that a same-named BI email exists.
	// Since we provide initial access to FC accounts via the BI recovery email, we can reasonably
	// assume that the recovery email is correct.
	// But there's a chance that it isn't. Sherlock has discovered human misconfigurations in the
	// FC workspace before, and it will again. If the user connects with a BI account that isn't
	// in the FC suitability cache, *but a same-named FC account is in the cache*, it means that
	// the recovery email for their FC user is incorrect. Rather than sending the champions person
	// on a wild goose chase to figure out that that's what's going on, we'll just explain it in
	// the suitability description (which will make its way into the error message).
	if strings.HasSuffix(email, "@broadinstitute.org") {
		firecloudEmail := fmt.Sprintf("%s@firecloud.org", strings.TrimSuffix(email, "@broadinstitute.org"))
		if _, present := cachedFirecloudSuitability[firecloudEmail]; present {
			notSuitable.description = fmt.Sprintf("SUPER IMPORTANT: %s isn't recognized, but the same-named %s is. That almost definitely means the %s Firecloud user account is misconfigured and should have %s set as the associated recovery email. DevOps needs to fix this. Once they do, you'll need to wait %d minutes for Sherlock to pick up the update.",
				email, firecloudEmail, firecloudEmail, email, config.Config.MustInt("auth.updateIntervalMinutes"))
		}
	}
	return notSuitable
}

// Package advisory_locks offers constants to differentiate PostgreSQL advisory locks that Sherlock uses.
// The constants here are meant to be used as "key1" arguments to the various advisory lock functions [1].
// Each constant documents what the "key2" argument should be.
//
// Note that it is critical that the constants in this package not be renumbered or removed, as that would
// be a breaking change for how Sherlock uses its database (would cause problems with old replicas upon
// deployment).
//
// [1]: https://www.postgresql.org/docs/current/functions-admin.html#FUNCTIONS-ADVISORY-LOCKS
package advisory_locks

// We disable unused detection here because `none` is intentionally unused. We have to put these comments
// on the entire const block for them to be parsed correctly.
//
//nolint:unused
//goland:noinspection GoUnusedConst
const (
	// none is a placeholder that should never be used. It exists so that the actual exported constants
	// begin at 1, not 0. This helps limit blast radius if we accidentally use an unset integer as "key1",
	// because then at least we won't be conflicting with any correct usages.
	none int = iota

	// ROLE_PROPAGATION locks models.Role records for propagation to cloud providers. The "key2" argument
	// should be the ID of the models.Role.
	//
	// This lock exists so that we don't try to propagate the same models.Role concurrently. This lock
	// should be acquired before determining if a models.Role should be propagated.
	ROLE_PROPAGATION

	// FIRECLOUD_ACCOUNT_MANAGER locks a firecloud_account_manager instance entry while it is actively
	// looking at the Google Workspace to suspend accounts. The "key2" argument should be 1 plus the
	// index of the firecloud_account_manager.Config in actual config.Config. There can be multiple
	// Firecloud account manager configs so we can manage multiple Fireclouds.
	//
	// This lock exists so that we don't try to suspend accounts in the same Firecloud concurrently.
	// That could cause duplicate notifications. This lock should be acquired before determining what
	// accounts need to be suspended.
	FIRECLOUD_ACCOUNT_MANAGER
)

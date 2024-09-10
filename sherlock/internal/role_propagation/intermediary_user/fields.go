package intermediary_user

// Fields represents the data that Sherlock should control about a user on the cloud provider (as opposed to the data
// that identifies the user on the cloud provider -- see Identifier).
//
// Implementations should generally be in propagation_engines (to be coupled to an engine) and should use non-pointer
// receivers.
type Fields interface {
	// EqualTo returns true if the two Fields are equal. This is used to determine if a user's fields have changed and
	// need to be updated.
	EqualTo(other Fields) bool
}

// MayBePresentWhileRemovedFields is an optional superset of Fields. It should be implemented when a user's data on
// a cloud provider may be present even after the user has been removed. It helps Sherlock know that it shouldn't
// infinitely keep trying to remove remnants of the user's data.
//
// Just like Fields, implementations should be in propagation_engines and should use non-pointer receivers.
type MayBePresentWhileRemovedFields interface {
	Fields

	// MayConsiderAsAlreadyRemoved should return true if Sherlock should skip attempting to remove the user from the
	// grant on the cloud provider. This method should only be called before removing a user -- it's output isn't
	// instructive for adding or updating a user.
	//
	// An example would be for a grant that provisions Firecloud accounts for users. We create accounts normally,
	// update fields normally, but we don't remove normally. When Sherlock goes to remove a user, we want it to
	// merely suspend the user's account, not delete it. The problem is that the user's account is still present:
	// Sherlock will keep removing it over and over again, each time the propagation runs. This method fixes that.
	// It could return true if the fields indicated the user was suspended, telling Sherlock that the user was
	// already in an acceptably removed state.
	MayConsiderAsAlreadyRemoved() bool
}

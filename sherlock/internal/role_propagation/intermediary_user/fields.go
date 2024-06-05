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

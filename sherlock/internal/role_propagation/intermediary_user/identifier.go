package intermediary_user

// Identifier represents the data that Sherlock should use to identify users on the cloud provider (as opposed to the
// data that Sherlock should control about the user on the cloud provider -- see Fields).
//
// Implementations should generally be in propagation_engines (to be coupled to an engine) and should use non-pointer
// receivers.
type Identifier interface {
	// EqualTo returns true if the two Identifiers are equal. This is used to match users in the cloud provider to
	// users in Sherlock.
	EqualTo(other Identifier) bool
}

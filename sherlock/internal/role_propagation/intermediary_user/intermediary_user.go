// Package intermediary_user helps represent the intersection of "what Sherlock cares about" and "what the cloud
// provider cares about" for the purposes of role propagation. The Identifier is a struct with whatever's needed to
// identify the user uniquely in the cloud provider, and the Fields is a struct with anything else Sherlock should
// control about the user in the cloud provider.
//
// Implementations of Identifier and Fields should be in propagation_engines, not here (since they should be tightly
// coupled to the propagation engine that uses them, and even if they aren't it just helps to be consistent).
package intermediary_user

type IntermediaryUser[
	I Identifier,
	F Fields,
] struct {
	Identifier I
	Fields     F
}

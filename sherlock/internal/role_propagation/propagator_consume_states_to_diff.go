package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
)

// consumeStatesToDiff returns a list of functions that, when called, will align the currentState with the
// desiredState. This function uses the state inputs as accumulators, so the inputs given to this function should not
// be used afterward.
//
// This function has a lot of loops that generate functions that will be called later. We explicitly initialize new
// variables to store the values of the loop variables so that the functions generated will capture the correct values.
// We just do this pattern everywhere in this function because it's always correct and it's easier to reason about.
func (p *propagatorImpl[Grant, Identifier, Fields]) consumeStatesToDiff(
	ctx context.Context,
	grant Grant,
	currentState []intermediary_user.IntermediaryUser[Identifier, Fields],
	desiredState map[uint]intermediary_user.IntermediaryUser[Identifier, Fields],
) (alignmentOperations []func() (string, error)) {

currentlyGrantedUserLoop:
	for _, unsafeCurrentlyGrantedUser := range currentState {
		currentlyGrantedUser := unsafeCurrentlyGrantedUser

		// Seek match from desiredState
		for unsafeDesiredSherlockUserID, unsafeDesiredUser := range desiredState {
			desiredSherlockUserID := unsafeDesiredSherlockUserID
			desiredUser := unsafeDesiredUser

			if currentlyGrantedUser.Identifier.EqualTo(desiredUser.Identifier) {
				// Match! If fields are different we update; either way we move on to the next currently granted user.
				// We actually remove the entry from desiredState so we know what's left over and needs to be added
				// at the end.
				if !currentlyGrantedUser.Fields.EqualTo(desiredUser.Fields) {
					alignmentOperations = append(alignmentOperations, func() (string, error) {
						return p.engine.Update(ctx, grant, desiredUser.Identifier, currentlyGrantedUser.Fields, desiredUser.Fields)
					})
				}

				delete(desiredState, desiredSherlockUserID)

				continue currentlyGrantedUserLoop
			}
		}

		// No match from desiredState! Let's seek a match in the users we are configured to tolerate.
		for _, toleratedUser := range p._toleratedUsers {
			if currentlyGrantedUser.Identifier.EqualTo(toleratedUser) {
				// Match! Let's move on to the next currently granted user, we'll leave this one alone.
				continue currentlyGrantedUserLoop
			}
		}

		// No match in desiredState or toleratedUsers! Remove the grant from the currently granted user.
		alignmentOperations = append(alignmentOperations, func() (string, error) {
			return p.engine.Remove(ctx, grant, currentlyGrantedUser.Identifier)
		})
	}

	// If there are any desired users left, add them.
	for _, unsafeDesiredUser := range desiredState {
		desiredUser := unsafeDesiredUser
		alignmentOperations = append(alignmentOperations, func() (string, error) {
			return p.engine.Add(ctx, grant, desiredUser.Identifier, desiredUser.Fields)
		})
	}
	return alignmentOperations
}

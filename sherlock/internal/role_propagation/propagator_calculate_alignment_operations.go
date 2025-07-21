package role_propagation

import (
	"context"
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
)

// calculateAlignmentOperations returns a list of functions that, when called, will align the currentState with the
// desiredState.
//
// This function has a lot of loops that generate functions that will be called later. We explicitly initialize new
// variables to store the values of the loop variables so that the functions generated will capture the correct values.
// We just do this pattern everywhere in this function because it's always correct and it's easier to reason about.
func (p *propagatorImpl[Grant, Identifier, Fields]) calculateAlignmentOperations(
	ctx context.Context,
	grant Grant,
	currentState []intermediary_user.IntermediaryUser[Identifier, Fields],
	desiredState map[uint]intermediary_user.IntermediaryUser[Identifier, Fields],
) (alignmentOperations []func() (string, error)) {

	// We don't want to mutate the state inputs we were given, so we assemble a copy of the desired state map to mess
	// with as we go.
	copyOfDesiredState := make(map[uint]intermediary_user.IntermediaryUser[Identifier, Fields])
	for userID, intermediaryUser := range desiredState {
		copyOfDesiredState[userID] = intermediaryUser
	}

currentlyGrantedUserLoop:
	for _, unsafeCurrentlyGrantedUser := range currentState {
		currentlyGrantedUser := unsafeCurrentlyGrantedUser

		// Let's first check if we're set to ignore the currently granted user completely.
		for _, ignoredUser := range p._ignoredUsers {
			if currentlyGrantedUser.Identifier.EqualTo(ignoredUser) {
				// Match! Let's move on to the next currently granted user, we'll leave this one alone.
				continue currentlyGrantedUserLoop
			}
		}

		// Seek match from copyOfDesiredState
		for unsafeDesiredSherlockUserID, unsafeDesiredUser := range copyOfDesiredState {
			desiredSherlockUserID := unsafeDesiredSherlockUserID
			desiredUser := unsafeDesiredUser

			if currentlyGrantedUser.Identifier.EqualTo(desiredUser.Identifier) {
				// Match! If fields are different we update; either way we move on to the next currently granted user.
				// We actually remove the entry from copyOfDesiredState so we know what's left over and needs to be
				// added at the end.
				if !currentlyGrantedUser.Fields.EqualTo(desiredUser.Fields) {
					alignmentOperations = append(alignmentOperations, func() (string, error) {
						return p.updateOperation()(ctx, grant, desiredUser.Identifier, currentlyGrantedUser.Fields, desiredUser.Fields)
					})
				}

				delete(copyOfDesiredState, desiredSherlockUserID)

				continue currentlyGrantedUserLoop
			}
		}

		// No match from copyOfDesiredState! Let's seek a match in the users we are configured to tolerate.
		for _, toleratedUser := range p._toleratedUsers {
			if currentlyGrantedUser.Identifier.EqualTo(toleratedUser) {
				// Match! Let's move on to the next currently granted user, we'll leave this one alone.
				continue currentlyGrantedUserLoop
			}
		}

		// No match in desiredState or toleratedUsers! Let's check if we may consider the user as being already
		// effectively removed. We need to check two things:
		// 1. If the fields are the *type* that may still be present while the user is effectively removed.
		// 2. If those fields indicate *this user* can be considered as effectively already removed.
		//
		// (We can't type-assert on a type parameter, so we convert to the interface supertype and then assert on that.)
		if mayBePresentWhileRemovedFields, ok := intermediary_user.Fields(currentlyGrantedUser.Fields).(intermediary_user.MayBePresentWhileRemovedFields); ok &&
			mayBePresentWhileRemovedFields.MayConsiderAsAlreadyRemoved() {
			continue currentlyGrantedUserLoop
		}

		// No match in desiredState or toleratedUsers, and our check if we could treat the user as being effectively
		// already removed didn't pass. We remove the grant from the currently granted user.
		alignmentOperations = append(alignmentOperations, func() (string, error) {
			return p.removeOperation()(ctx, grant, currentlyGrantedUser.Identifier)
		})
	}

	// If there are any desired users left, add them.
desiredUserLoop:
	for _, unsafeDesiredUser := range copyOfDesiredState {
		desiredUser := unsafeDesiredUser

		// Let's first check if we're set to ignore the desired user completely.
		for _, ignoredUser := range p._ignoredUsers {
			if desiredUser.Identifier.EqualTo(ignoredUser) {
				// Match! Let's move on to the next desired user, we'll leave this one alone.
				continue desiredUserLoop
			}
		}

		// If we get here, we want to add the user.
		alignmentOperations = append(alignmentOperations, func() (string, error) {
			return p.addOperation()(ctx, grant, desiredUser.Identifier, desiredUser.Fields)
		})
	}
	return alignmentOperations
}

func (p *propagatorImpl[Grant, Identifier, Fields]) addOperation() func(ctx context.Context, grant Grant, identifier Identifier, fields Fields) (string, error) {
	if p._dryRun {
		return func(ctx context.Context, grant Grant, identifier Identifier, fields Fields) (string, error) {
			return fmt.Sprintf("DRY-RUN: called for adding of %+v with fields %+v", identifier, fields), nil
		}
	}
	return p.engine.Add
}

func (p *propagatorImpl[Grant, Identifier, Fields]) updateOperation() func(ctx context.Context, grant Grant, identifier Identifier, oldFields Fields, newFields Fields) (string, error) {
	if p._dryRun {
		return func(ctx context.Context, grant Grant, identifier Identifier, oldFields Fields, newFields Fields) (string, error) {
			return fmt.Sprintf("DRY-RUN: called for update of %+v from %+v to %+v", identifier, oldFields, newFields), nil
		}
	}
	return p.engine.Update
}

func (p *propagatorImpl[Grant, Identifier, Fields]) removeOperation() func(ctx context.Context, grant Grant, identifier Identifier) (string, error) {
	if p._dryRun {
		return func(ctx context.Context, grant Grant, identifier Identifier) (string, error) {
			return fmt.Sprintf("DRY-RUN: called for removal of %+v", identifier), nil
		}
	}
	return p.engine.Remove
}

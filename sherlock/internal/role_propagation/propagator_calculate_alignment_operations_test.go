package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user/intermediary_user_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines/propagation_engines_mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testIdentifier struct {
	identifier string
}

func (t testIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	return t.identifier == other.(testIdentifier).identifier
}

type testFields struct {
	field string
}

func (t testFields) EqualTo(other intermediary_user.Fields) bool {
	return t.field == other.(testFields).field
}

type testIntermediaryUser = intermediary_user.IntermediaryUser[testIdentifier, testFields]

func Test_propagatorImpl_calculateAlignmentOperations(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, testIdentifier, testFields](t)
	p := &propagatorImpl[string, testIdentifier, testFields]{
		engine: engine,
	}

	ctx := context.Background()
	grant := "grant"
	currentState := make([]testIntermediaryUser, 0)
	desiredState := make(map[uint]testIntermediaryUser)

	// User in both current and desired, no changes, expect nothing to be called
	identifier1 := testIdentifier{"user1"}
	fields1 := testFields{"field1"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier1, Fields: fields1})
	desiredState[1] = testIntermediaryUser{Identifier: identifier1, Fields: fields1}

	// User in both current and desired, fields differ, expect Update to be called
	identifier2 := testIdentifier{"user2"}
	fields2 := testFields{"field2"}
	fields2Desired := testFields{"field2Desired"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier2, Fields: fields2})
	desiredState[2] = testIntermediaryUser{Identifier: identifier2, Fields: fields2Desired}
	engine.EXPECT().Update(ctx, grant, identifier2, fields2, fields2Desired).Return("", nil).Once()

	// User in current, not in desired, expect Remove to be called
	identifier3 := testIdentifier{"user3"}
	fields3 := testFields{"field3"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier3, Fields: fields3})
	engine.EXPECT().Remove(ctx, grant, identifier3).Return("", nil).Once()

	// User in current, not in desired, but in tolerated, expect nothing to be called
	identifier4 := testIdentifier{"user4"}
	fields4 := testFields{"field4"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier4, Fields: fields4})
	p._toleratedUsers = append(p._toleratedUsers, identifier4)

	// User in tolerated, not in current or desired, expect nothing to be called
	identifier5 := testIdentifier{"user5"}
	p._toleratedUsers = append(p._toleratedUsers, identifier5)

	// User in desired, not in current, expect Add to be called
	identifier6 := testIdentifier{"user6"}
	fields6 := testFields{"field6"}
	desiredState[6] = testIntermediaryUser{Identifier: identifier6, Fields: fields6}
	engine.EXPECT().Add(ctx, grant, identifier6, fields6).Return("", nil).Once()

	currentStateLen := len(currentState)
	desiredStateLen := len(desiredState)

	alignmentOperations := p.calculateAlignmentOperations(ctx, grant, currentState, desiredState)
	for _, alignmentOperation := range alignmentOperations {
		// These operations are pure mocks so there's no point to testing their return values,
		// we're just calling the outputs so the mock observes the calls
		_, _ = alignmentOperation()
	}

	// Make sure that the above operations didn't mutate the state parameters
	assert.Len(t, currentState, currentStateLen)
	assert.Len(t, desiredState, desiredStateLen)

	engine.AssertExpectations(t)
}

func Test_propagatorImpl_calculateAlignmentOperations_dryRun(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, testIdentifier, testFields](t)
	p := &propagatorImpl[string, testIdentifier, testFields]{
		engine:  engine,
		_dryRun: true,
	}

	ctx := context.Background()
	grant := "grant"
	currentState := make([]testIntermediaryUser, 0)
	desiredState := make(map[uint]testIntermediaryUser)
	results := make([]string, 0)

	// User in both current and desired, no changes, expect nothing to be called
	identifier1 := testIdentifier{"user1"}
	fields1 := testFields{"field1"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier1, Fields: fields1})
	desiredState[1] = testIntermediaryUser{Identifier: identifier1, Fields: fields1}

	// User in both current and desired, fields differ, expect Update to be called
	identifier2 := testIdentifier{"user2"}
	fields2 := testFields{"field2"}
	fields2Desired := testFields{"field2Desired"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier2, Fields: fields2})
	desiredState[2] = testIntermediaryUser{Identifier: identifier2, Fields: fields2Desired}
	results = append(results, fmt.Sprintf("DRY-RUN: called for update of %+v from %+v to %+v", identifier2, fields2, fields2Desired))

	// User in current, not in desired, expect Remove to be called
	identifier3 := testIdentifier{"user3"}
	fields3 := testFields{"field3"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier3, Fields: fields3})
	results = append(results, fmt.Sprintf("DRY-RUN: called for removal of %+v", identifier3))

	// User in current, not in desired, but in tolerated, expect nothing to be called
	identifier4 := testIdentifier{"user4"}
	fields4 := testFields{"field4"}
	currentState = append(currentState, testIntermediaryUser{Identifier: identifier4, Fields: fields4})
	p._toleratedUsers = append(p._toleratedUsers, identifier4)

	// User in tolerated, not in current or desired, expect nothing to be called
	identifier5 := testIdentifier{"user5"}
	p._toleratedUsers = append(p._toleratedUsers, identifier5)

	// User in desired, not in current, expect Add to be called
	identifier6 := testIdentifier{"user6"}
	fields6 := testFields{"field6"}
	desiredState[6] = testIntermediaryUser{Identifier: identifier6, Fields: fields6}
	results = append(results, fmt.Sprintf("DRY-RUN: called for adding of %+v with fields %+v", identifier6, fields6))

	currentStateLen := len(currentState)
	desiredStateLen := len(desiredState)

	alignmentOperations := p.calculateAlignmentOperations(ctx, grant, currentState, desiredState)
	actualResults := make([]string, 0)
	for _, alignmentOperation := range alignmentOperations {
		result, err := alignmentOperation()
		assert.NoError(t, err)
		actualResults = append(actualResults, result)
	}

	assert.ElementsMatch(t, results, actualResults)

	// Make sure that the above operations didn't mutate the state parameters
	assert.Len(t, currentState, currentStateLen)
	assert.Len(t, desiredState, desiredStateLen)

	engine.AssertExpectations(t)
}

func Test_propagatorImpl_calculateAlignmentOperations_empty(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, testIdentifier, testFields](t)
	p := &propagatorImpl[string, testIdentifier, testFields]{
		engine: engine,
	}

	// This is a sorta dumb test but we're checking that we don't somehow fail on empty inputs -- everything else
	// is pretty thoroughly covered by the main test above
	alignmentOperations := p.calculateAlignmentOperations(context.Background(), "", nil, nil)
	for _, alignmentOperation := range alignmentOperations {
		_, _ = alignmentOperation()
	}

	engine.AssertExpectations(t)
}

type testMayConsiderAsAlreadyRemovedIntermediaryUser = intermediary_user.IntermediaryUser[testIdentifier, *intermediary_user_mocks.MockMayBePresentWhileRemovedFields]

func Test_propagatorImpl_calculateAlignmentOperations_mayConsiderAsAlreadyRemoved(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, testIdentifier, *intermediary_user_mocks.MockMayBePresentWhileRemovedFields](t)
	p := &propagatorImpl[string, testIdentifier, *intermediary_user_mocks.MockMayBePresentWhileRemovedFields]{
		engine: engine,
	}

	ctx := context.Background()
	grant := "grant"
	currentState := make([]testMayConsiderAsAlreadyRemovedIntermediaryUser, 0)
	desiredState := make(map[uint]testMayConsiderAsAlreadyRemovedIntermediaryUser)

	// A user that should be removed
	identifier1 := testIdentifier{"user1"}
	fields1 := intermediary_user_mocks.NewMockMayBePresentWhileRemovedFields(t)
	fields1.EXPECT().MayConsiderAsAlreadyRemoved().Return(false).Once()
	currentState = append(currentState, testMayConsiderAsAlreadyRemovedIntermediaryUser{Identifier: identifier1, Fields: fields1})
	engine.EXPECT().Remove(ctx, grant, identifier1).Return("", nil).Once()

	// A user that should be considered as already removed -- expect no Remove to be called
	identifier2 := testIdentifier{"user2"}
	fields2 := intermediary_user_mocks.NewMockMayBePresentWhileRemovedFields(t)
	fields2.EXPECT().MayConsiderAsAlreadyRemoved().Return(true).Once()
	currentState = append(currentState, testMayConsiderAsAlreadyRemovedIntermediaryUser{Identifier: identifier2, Fields: fields2})

	alignmentOperations := p.calculateAlignmentOperations(ctx, grant, currentState, desiredState)
	for _, alignmentOperation := range alignmentOperations {
		// These operations are pure mocks so there's no point to testing their return values,
		// we're just calling the outputs so the mock observes the calls
		_, _ = alignmentOperation()
	}

	fields1.AssertExpectations(t)
	fields2.AssertExpectations(t)
	engine.AssertExpectations(t)
}

package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
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

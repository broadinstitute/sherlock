package authentication_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestEmptyMethod is a dumb test -- it just cares that an uninitialized Method
// is UNKNOWN and definitely, definitely not IAP. There shouldn't be anything
// in Sherlock relying on uninitialized Method values, but it's an extra layer
// of defense to just assert that the uninitialized value doesn't have a
// potentially dangerous meaning.
// In other words, this test exists to help catch problematic refactorings,
// because the order of the lines in enum.go matters.
func TestEmptyMethod(t *testing.T) {
	var anEmptyAuthenticationMethod Method
	assert.Equal(t, UNKNOWN, anEmptyAuthenticationMethod)
	assert.NotEqual(t, IAP, anEmptyAuthenticationMethod)
	assert.NotEqual(t, GHA, anEmptyAuthenticationMethod)
}

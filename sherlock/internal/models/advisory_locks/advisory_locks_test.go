package advisory_locks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// These tests exist because these consts are used as integers in the database -- if they were to
// change, that would be breaking.

func Test_none(t *testing.T) {
	assert.Equal(t, 0, none)
}

func Test_ROLE_PROPAGATION(t *testing.T) {
	assert.Equal(t, 1, ROLE_PROPAGATION)
}

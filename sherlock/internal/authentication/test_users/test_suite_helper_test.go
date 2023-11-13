package test_users

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestUseSuitableUserFor(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	TestUserHelper{}.UseSuitableUserFor(req)
	suitable, err := strconv.ParseBool(req.Header.Get(suitableControlHeader))
	assert.NoError(t, err)
	assert.True(t, suitable)
}

func TestUseNonSuitableUserFor(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	TestUserHelper{}.UseNonSuitableUserFor(req)
	suitable, err := strconv.ParseBool(req.Header.Get(suitableControlHeader))
	assert.NoError(t, err)
	assert.False(t, suitable)
}

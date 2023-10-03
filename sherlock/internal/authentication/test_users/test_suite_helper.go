package test_users

import (
	"net/http"
	"strconv"
)

// TestUserHelper offers helper functions to customize how a http.Request will be converted to a user email
// by ParseHeader. Closer integration with the package being tested would have to be done in that package,
// to avoid circular imports.
// This is a struct so that it can be embedded into the test suites of other packages, so these functions
// are just that little bit easier to call (to help nudge towards using them, instead of manually setting
// headers from tests)
type TestUserHelper struct{}

// UseSuitableUserFor sets SuitabilityControlHeader such that ParseHeader will supply SuitableTestUserEmail.
// This is ParseHeader's default behavior, but this function can be helpful for clarity or undoing
// UseNonSuitableUserFor.
func (h TestUserHelper) UseSuitableUserFor(req *http.Request) *http.Request {
	return h.selectUserForRequestBySuitability(req, true)
}

// UseNonSuitableUserFor sets SuitabilityControlHeader such that ParseHeader will supply NonSuitableTestUserEmail.
func (h TestUserHelper) UseNonSuitableUserFor(req *http.Request) *http.Request {
	return h.selectUserForRequestBySuitability(req, false)
}

func (_ TestUserHelper) selectUserForRequestBySuitability(req *http.Request, suitable bool) *http.Request {
	req.Header.Set(SuitabilityControlHeader, strconv.FormatBool(suitable))
	return req
}

package testutils

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// helpers.go contains general purpose resuable helper functions that
// can be helpful for unit and functional tests

const (
	StringNumberTooBigForInt = "9999999999999999999999999999"
)

func AssertNoDiff(t *testing.T, want any, got any) {
	t.Helper()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("expected equality, got mismatch (-want +got):\n%s", diff)
	}
}

// PointerTo returns a pointer to whatever you give it, so you don't need to
// define a bunch of temporary variables in tests. The fun generics make it
// always agree with the type system.
func PointerTo[T any](val T) *T {
	return &val
}

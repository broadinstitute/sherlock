package testutils

import (
	"strconv"
	"testing"
)

// TestStringNumberTooBigForInt is a dumb test but a lot of tests elsewhere actually use this const to
// check that they create errors correctly. If there's a Go update or something, this test will fail if
// some architectural change makes this string parseable.
//
// (For a bit more context, strconv.Atoi throws errors, so when we use it elsewhere in Sherlock, it means
// those parts of Sherlock can throw errors, which means errors can appear from weird places. This const
// is used to check that we correctly handle those cases. If it suddenly started parsing, it would take
// longer to find what caused a bunch of tests to fail than it took Jack to write this test and comment)
func TestStringNumberTooBigForInt(t *testing.T) {
	result, err := strconv.Atoi(StringNumberTooBigForInt)
	if err == nil {
		t.Errorf("StringNumberTooBigForInt ('%s') was successfully parsed to an int? (%d)", StringNumberTooBigForInt, result)
	}
}

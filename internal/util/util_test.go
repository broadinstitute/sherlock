package util

import "testing"

func TestMultBy3(t *testing.T) {
	input := 20
	want := 60
	got := multBy3(input)
	if got != want {
		t.Errorf("Expecting: %v, got: %v", want, got)
	}
}

func TestPrintTriple(t *testing.T) {
	// this is where some additional test would go
	PrintTriple(3)
}

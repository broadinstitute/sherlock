package models

import (
	"math/rand"
	"reflect"
	"time"
)

// TestTestDataItself iterates through the TestData's underlying type and executes each
// method in a random order. This helps flag when the issue is with the TestData and not
// the test calling it.
func (s *modelSuite) TestTestDataItself() {
	testDataValue := reflect.ValueOf(s.TestData)
	// Rather than iterating through the methods in order, use rand.Perm to effectively
	// shuffle the order we process each index.
	// https://pkg.go.dev/math/rand#Rand.Perm
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(testDataValue.NumMethod()) {
		methodTypeValue := testDataValue.Type().Method(i)
		// Method's receiver counts as an argument
		if methodTypeValue.Type.NumIn() == 1 {
			s.Run(methodTypeValue.Name, func() {
				// Call the method accessed from the Value directly, not the Type's Value,
				// so we don't have to pass the receiver ourselves
				testDataValue.Method(i).Call([]reflect.Value{})
			})
		}
	}
}

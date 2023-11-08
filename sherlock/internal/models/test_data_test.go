package models

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// TestBulkData uses reflection to iterate through all available methods on TestData, followed
// by a run of UpdateMetrics to at least validate that the batch metrics cron doesn't error
func (s *modelSuite) TestBulkData() {
	testDataValue := reflect.ValueOf(s.TestData)
	// Rather than iterating through the methods in order, use rand.Perm to effectively
	// shuffle the order we process each index.
	// https://pkg.go.dev/math/rand#Rand.Perm
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(testDataValue.NumMethod()) {
		methodTypeValue := testDataValue.Type().Method(i)
		// Method's receiver counts as an argument
		if methodTypeValue.Type.NumIn() == 1 {
			s.Run(fmt.Sprintf("run TestData.%s()", methodTypeValue.Name), func() {
				// Call the method accessed from the Value directly, not the Type's Value,
				// so we don't have to pass the receiver ourselves
				testDataValue.Method(i).Call([]reflect.Value{})
			})
		}
	}
	s.Run("run batch metrics cron", func() {
		s.NoError(UpdateMetrics(context.Background(), s.DB))
	})
}

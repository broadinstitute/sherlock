package services

import "fmt"

// MockServiceStore is an abstraction
// intended to represent the ability to select services
// from a database for use in unit testing
type MockServiceStore struct {
	services []Service
}

// Select implements the db.Selector interface to support using a mock database with the
// service listing
func (m *MockServiceStore) Select(dest interface{}, query string, args ...interface{}) error {
	if query == selectAll {
		switch d := dest.(type) {
		case *[]Service:
			result := make([]Service, len(m.services))
			copy(result, m.services)
			*d = result
		default:
			return fmt.Errorf("Cannot copy services into provided destination. Invalid type")
		}
	} else {
		return fmt.Errorf("invalid query: %s", query)
	}
	return nil
}

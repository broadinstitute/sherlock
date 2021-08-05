package services

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// )

// func TestListAllServices(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		services []Service
// 	}{
// 		{
// 			name:     "no existing services",
// 			services: []Service{},
// 		},
// 		{
// 			name: "one existing service",
// 			services: []Service{
// 				{
// 					ID:        1,
// 					Name:      "cromwell",
// 					RepoURL:   "https://github.com/broadinstitute/cromwell",
// 					CreatedAt: "2021-08-03T17:22:41.86241Z",
// 				},
// 			},
// 		},
// 		{
// 			name: "multiple existing services",
// 			services: []Service{
// 				{
// 					ID:        1,
// 					Name:      "cromwell",
// 					RepoURL:   "https://github.com/broadinstitute/cromwell",
// 					CreatedAt: "2021-08-03T17:22:41.86241Z",
// 				},
// 				{
// 					ID:        2,
// 					Name:      "leonardo",
// 					RepoURL:   "https://github.com/databiosphere/leonardo",
// 					CreatedAt: "2021-08-01T17:22:41.86241Z",
// 				},
// 				{
// 					ID:        3,
// 					Name:      "buffer",
// 					RepoURL:   "https://github.com/databiosphere/buffer",
// 					CreatedAt: "2021-08-02T17:22:41.86241Z",
// 				},
// 			},
// 		},
// 	}

// 	for _, testCase := range cases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			expectedServices := testCase.services
// 			gotServices, err := ListAll(&mockServiceStore{services: expectedServices})
// 			if err != nil {
// 				t.Errorf("recieved unexpected error %v\n", err)
// 			}

// 			if diff := cmp.Diff(expectedServices, gotServices); diff != "" {
// 				t.Errorf("got unexpected services: %v\n", diff)
// 			}
// 		})
// 	}

// 	t.Run("test failure mode on ListAll", func(t *testing.T) {
// 		_, err := ListAll(&failingServiceStore{})

// 		if err == nil {
// 			t.Errorf("expected to receive an error but didn't")
// 		}
// 	})
// }

// // MockServiceStore is an abstraction
// // intended to represent the ability to select services
// // from a database for use in unit testing
// type mockServiceStore struct {
// 	services []Service
// }

// // Select implements the db.Selector interface to support using a mock database with the
// // service listing
// func (m *mockServiceStore) Select(dest interface{}, query string, args ...interface{}) error {
// 	if query == selectAll {
// 		switch d := dest.(type) {
// 		case *[]Service:
// 			result := make([]Service, len(m.services))
// 			copy(result, m.services)
// 			*d = result
// 		default:
// 			return fmt.Errorf("Cannot copy services into provided destination. Invalid type")
// 		}
// 	} else {
// 		return fmt.Errorf("invalid query: %s", query)
// 	}
// 	return nil
// }

// // failing service store and its associated method are intended to unit test the failure mode
// // of operations interacting with services in some data store
// type failingServiceStore struct{}

// func (f *failingServiceStore) Select(dest interface{}, query string, args ...interface{}) error {
// 	return fmt.Errorf("Always fails")
// }

package services

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListAllServices(t *testing.T) {
	cases := []struct {
		name     string
		services []Service
	}{
		{
			name:     "no existing services",
			services: []Service{},
		},
		{
			name: "one existing service",
			services: []Service{
				{
					ID:        1,
					Name:      "cromwell",
					RepoURL:   "https://github.com/broadinstitute/cromwell",
					CreatedAt: "2021-08-03T17:22:41.86241Z",
				},
			},
		},
		{
			name: "multiple existing services",
			services: []Service{
				{
					ID:        1,
					Name:      "cromwell",
					RepoURL:   "https://github.com/broadinstitute/cromwell",
					CreatedAt: "2021-08-03T17:22:41.86241Z",
				},
				{
					ID:        2,
					Name:      "leonardo",
					RepoURL:   "https://github.com/databiosphere/leonardo",
					CreatedAt: "2021-08-01T17:22:41.86241Z",
				},
				{
					ID:        3,
					Name:      "buffer",
					RepoURL:   "https://github.com/databiosphere/buffer",
					CreatedAt: "2021-08-02T17:22:41.86241Z",
				},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedServices := testCase.services
			gotServices, err := ListAll(&MockServiceStore{services: expectedServices})
			if err != nil {
				t.Errorf("recieved unexpected error %v\n", err)
			}

			if diff := cmp.Diff(expectedServices, gotServices); diff != "" {
				t.Errorf("got unexpected services: %v\n", diff)
			}
		})
	}
}

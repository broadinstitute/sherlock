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
					ID:      1,
					Name:    "cromwell",
					RepoURL: "https://github.com/broadinstitute/cromwell",
				},
			},
		},
		{
			name: "multiple existing services",
			services: []Service{
				{
					ID:      1,
					Name:    "cromwell",
					RepoURL: "https://github.com/broadinstitute/cromwell",
				},
				{
					ID:      2,
					Name:    "leonardo",
					RepoURL: "https://github.com/databiosphere/leonardo",
				},
				{
					ID:      3,
					Name:    "buffer",
					RepoURL: "https://github.com/databiosphere/buffer",
				},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedServices := testCase.services
			gotServices, err := ListAll(&mockServiceStore{services: expectedServices})
			if err != nil {
				t.Errorf("recieved unexpected error %v\n", err)
			}

			if diff := cmp.Diff(expectedServices, gotServices); diff != "" {
				t.Errorf("got unexpected services: %v\n", diff)
			}
		})
	}

	t.Run("test failure mode on ListAll", func(t *testing.T) {
		_, err := ListAll(&failingServiceStore{})

		if err == nil {
			t.Errorf("expected to receive an error but didn't")
		}
	})
}

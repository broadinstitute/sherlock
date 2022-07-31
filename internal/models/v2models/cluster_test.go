package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_clusterToSelectors(t *testing.T) {
	type args struct {
		cluster Cluster
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{},
			want: nil,
		},
		{
			name: "with name",
			args: args{cluster: Cluster{
				Name: "foobar",
			}},
			want: []string{"foobar"},
		},
		{
			name: "with id",
			args: args{cluster: Cluster{
				Model: gorm.Model{ID: 123},
			}},
			want: []string{"123"},
		},
		{
			name: "with name and id",
			args: args{cluster: Cluster{
				Model: gorm.Model{ID: 123},
				Name:  "foobar",
			}},
			want: []string{"foobar", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clusterToSelectors(tt.args.cluster)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

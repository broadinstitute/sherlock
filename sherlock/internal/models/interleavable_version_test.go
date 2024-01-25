package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestInterleaveVersions(t *testing.T) {
	now := time.Now()
	type args struct {
		appVersions   []*AppVersion
		chartVersions []*ChartVersion
	}
	tests := []struct {
		name string
		args args
		want []InterleavableVersion
	}{
		{
			name: "interleaved",
			args: args{
				appVersions: []*AppVersion{
					{Model: gorm.Model{ID: 1, CreatedAt: now.Add(-3 * time.Hour)}},
					{Model: gorm.Model{ID: 2, CreatedAt: now.Add(-1 * time.Hour)}},
				},
				chartVersions: []*ChartVersion{
					{Model: gorm.Model{ID: 1, CreatedAt: now.Add(-4 * time.Hour)}},
					{Model: gorm.Model{ID: 2, CreatedAt: now.Add(-2 * time.Hour)}},
				},
			},
			want: []InterleavableVersion{
				&ChartVersion{Model: gorm.Model{ID: 1, CreatedAt: now.Add(-4 * time.Hour)}},
				&AppVersion{Model: gorm.Model{ID: 1, CreatedAt: now.Add(-3 * time.Hour)}},
				&ChartVersion{Model: gorm.Model{ID: 2, CreatedAt: now.Add(-2 * time.Hour)}},
				&AppVersion{Model: gorm.Model{ID: 2, CreatedAt: now.Add(-1 * time.Hour)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, InterleaveVersions(tt.args.appVersions, tt.args.chartVersions), "InterleaveVersions(%v, %v)", tt.args.appVersions, tt.args.chartVersions)
		})
	}
}

package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestChangeset_Summarize(t *testing.T) {
	type fields struct {
		Model            gorm.Model
		CiIdentifier     *CiIdentifier
		ChartRelease     *ChartRelease
		ChartReleaseID   uint
		From             ChartReleaseVersion
		To               ChartReleaseVersion
		AppliedAt        *time.Time
		SupersededAt     *time.Time
		NewAppVersions   []*AppVersion
		NewChartVersions []*ChartVersion
		PlannedBy        *User
		PlannedByID      *uint
		AppliedBy        *User
		AppliedByID      *uint
	}
	type args struct {
		includeFrom bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{includeFrom: false},
			want:   "configuration change",
		},
		{
			name: "app change, no include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: false},
			want: "app⭢1.2.3",
		},
		{
			name: "app change, include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: true},
			want: "app 1.2.2⭢1.2.3",
		},
		{
			name: "app change, include from but no data",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact: utils.PointerTo("1.2.3"),
				},
			},
			args: args{includeFrom: true},
			want: "app⭢1.2.3",
		},
		{
			name: "chart change, no include from",
			fields: fields{
				To: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: false},
			want: "chart⭢1.2.3",
		},
		{
			name: "chart change, include from",
			fields: fields{
				To: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: true},
			want: "chart 1.2.2⭢1.2.3",
		},
		{
			name: "chart change, include from but no data",
			fields: fields{
				To: ChartReleaseVersion{
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
			},
			args: args{includeFrom: true},
			want: "chart⭢1.2.3",
		},
		{
			name: "app and chart change, no include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.3"),
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.2"),
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: false},
			want: "app⭢1.2.3, chart⭢1.2.3",
		},
		{
			name: "app and chart change, include from",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.3"),
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
				From: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.2"),
					ChartVersionExact: utils.PointerTo("1.2.2"),
				},
			},
			args: args{includeFrom: true},
			want: "app 1.2.2⭢1.2.3, chart 1.2.2⭢1.2.3",
		},
		{
			name: "app and chart change, include from but no data",
			fields: fields{
				To: ChartReleaseVersion{
					AppVersionExact:   utils.PointerTo("1.2.3"),
					ChartVersionExact: utils.PointerTo("1.2.3"),
				},
			},
			args: args{includeFrom: true},
			want: "app⭢1.2.3, chart⭢1.2.3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Changeset{
				Model:            tt.fields.Model,
				CiIdentifier:     tt.fields.CiIdentifier,
				ChartRelease:     tt.fields.ChartRelease,
				ChartReleaseID:   tt.fields.ChartReleaseID,
				From:             tt.fields.From,
				To:               tt.fields.To,
				AppliedAt:        tt.fields.AppliedAt,
				SupersededAt:     tt.fields.SupersededAt,
				NewAppVersions:   tt.fields.NewAppVersions,
				NewChartVersions: tt.fields.NewChartVersions,
				PlannedBy:        tt.fields.PlannedBy,
				PlannedByID:      tt.fields.PlannedByID,
				AppliedBy:        tt.fields.AppliedBy,
				AppliedByID:      tt.fields.AppliedByID,
			}
			assert.Equalf(t, tt.want, c.Summarize(tt.args.includeFrom), "Summarize(%v)", tt.args.includeFrom)
		})
	}
}

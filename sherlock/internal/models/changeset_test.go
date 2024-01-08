package models

import (
	"cmp"
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

func TestCompareChangesetsByName(t *testing.T) {
	type args struct {
		a Changeset
		b Changeset
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "both with nil chart releases",
			args: args{
				a: Changeset{},
				b: Changeset{},
			},
			want: 0,
		},
		{
			name: "a with nil chart release",
			args: args{
				a: Changeset{},
				b: Changeset{ChartRelease: &ChartRelease{Name: "b"}},
			},
			want: -1,
		},
		{
			name: "b with nil chart release",
			args: args{
				a: Changeset{ChartRelease: &ChartRelease{Name: "a"}},
				b: Changeset{},
			},
			want: 1,
		},
		{
			name: "compare chart release names",
			args: args{
				a: Changeset{ChartRelease: &ChartRelease{Name: "a"}},
				b: Changeset{ChartRelease: &ChartRelease{Name: "b"}},
			},
			want: cmp.Compare("a", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CompareChangesetsByName(tt.args.a, tt.args.b), "CompareChangesetsByName(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestUsersFromChangesets(t *testing.T) {
	type args struct {
		changesets []Changeset
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "empty",
			args: args{changesets: []Changeset{}},
			want: nil,
		},
		{
			name: "sample case",
			args: args{changesets: []Changeset{
				{AppliedBy: &User{Model: gorm.Model{ID: 1}}},
				{AppliedBy: &User{Model: gorm.Model{ID: 2}}, PlannedBy: &User{Model: gorm.Model{ID: 1}}},
				{AppliedBy: &User{Model: gorm.Model{ID: 1}}, PlannedBy: &User{Model: gorm.Model{ID: 1}}},
				{AppliedBy: &User{Model: gorm.Model{ID: 3}}, PlannedBy: &User{Model: gorm.Model{ID: 4}}},
				{PlannedBy: &User{Model: gorm.Model{ID: 4}}},
			}},
			want: []User{
				{Model: gorm.Model{ID: 1}},
				{Model: gorm.Model{ID: 2}},
				{Model: gorm.Model{ID: 3}},
				{Model: gorm.Model{ID: 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UsersFromChangesets(tt.args.changesets), "UsersFromChangesets(%v)", tt.args.changesets)
		})
	}
}

package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCiIdentifierV3_toModel(t *testing.T) {
	t1 := time.Now().Add(-time.Minute)
	t2 := time.Now().Add(-time.Second)
	type fields struct {
		commonFields commonFields
		CiRuns       []CiRunV3
		ResourceType string
		ResourceID   uint
	}
	tests := []struct {
		name   string
		fields fields
		want   models.CiIdentifier
	}{
		{
			name: "equal",
			fields: fields{
				commonFields: commonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				CiRuns:       nil,
				ResourceType: "type",
				ResourceID:   2,
			},
			want: models.CiIdentifier{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				ResourceType: "type",
				ResourceID:   2,
				CiRuns:       nil,
			},
		},
		{
			name: "ignores ci runs",
			fields: fields{
				commonFields: commonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				CiRuns: []CiRunV3{
					{
						commonFields: commonFields{
							ID: 3,
						},
					},
				},
				ResourceType: "type",
				ResourceID:   2,
			},
			want: models.CiIdentifier{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				ResourceType: "type",
				ResourceID:   2,
				CiRuns:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CiIdentifierV3{
				commonFields: tt.fields.commonFields,
				CiRuns:       tt.fields.CiRuns,
				ResourceType: tt.fields.ResourceType,
				ResourceID:   tt.fields.ResourceID,
			}
			assert.Equalf(t, tt.want, c.toModel(), "toModel()")
		})
	}
}

func Test_ciIdentifierFromModel(t *testing.T) {
	t1 := time.Now().Add(-time.Minute)
	t2 := time.Now().Add(-time.Second)
	type args struct {
		model models.CiIdentifier
	}
	tests := []struct {
		name string
		args args
		want CiIdentifierV3
	}{
		{
			name: "equal",
			args: args{model: models.CiIdentifier{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
					DeletedAt: gorm.DeletedAt{
						Valid: true,
						Time:  time.Now(),
					},
				},
				ResourceType: "type",
				ResourceID:   2,
				CiRuns:       nil,
			}},
			want: CiIdentifierV3{
				commonFields: commonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				CiRuns:       nil,
				ResourceType: "type",
				ResourceID:   2,
			},
		},
		{
			name: "respects runs",
			args: args{model: models.CiIdentifier{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
					DeletedAt: gorm.DeletedAt{
						Valid: true,
						Time:  time.Now(),
					},
				},
				ResourceType: "type",
				ResourceID:   2,
				CiRuns: []models.CiRun{
					{
						Model: gorm.Model{
							ID: 3,
						},
					},
				},
			}},
			want: CiIdentifierV3{
				commonFields: commonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				CiRuns: []CiRunV3{
					{
						commonFields: commonFields{
							ID: 3,
						},
					},
				},
				ResourceType: "type",
				ResourceID:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ciIdentifierFromModel(tt.args.model), "ciIdentifierFromModel(%v)", tt.args.model)
		})
	}
}

func Test_ciIdentifierModelFromOldModel(t *testing.T) {
	type args struct {
		source v2models.CiIdentifiable
	}
	tests := []struct {
		name string
		args args
		want models.CiIdentifier
	}{
		{
			name: "by id",
			args: args{source: v2models.Chart{CiIdentifier: &v2models.CiIdentifier{Model: gorm.Model{ID: 1}}}},
			want: models.CiIdentifier{Model: gorm.Model{ID: 1}},
		},
		{
			name: "by resource",
			args: args{source: v2models.Chart{Model: gorm.Model{ID: 2}}},
			want: models.CiIdentifier{ResourceType: "chart", ResourceID: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ciIdentifierModelFromOldModel(tt.args.source), "ciIdentifierModelFromOldModel(%v)", tt.args.source)
		})
	}
}

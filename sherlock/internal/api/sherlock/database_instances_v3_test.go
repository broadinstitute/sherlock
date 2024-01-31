package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (s *handlerSuite) TestDatabaseInstanceV3_toModel() {
	now := time.Now()
	type fields struct {
		CommonFields             CommonFields
		ChartReleaseInfo         *ChartReleaseV3
		DatabaseInstanceV3Create DatabaseInstanceV3Create
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.DatabaseInstance
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			fields:  fields{},
			wantErr: assert.NoError,
			want:    models.DatabaseInstance{},
		},
		{
			name: "everything",
			fields: fields{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				DatabaseInstanceV3Create: DatabaseInstanceV3Create{
					DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
						Platform:        utils.PointerTo("kubernetes"),
						GoogleProject:   utils.PointerTo("google-project"),
						InstanceName:    utils.PointerTo("instance-name"),
						DefaultDatabase: utils.PointerTo("default-database"),
					},
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
			},
			wantErr: assert.NoError,
			want: models.DatabaseInstance{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				ChartReleaseID:  s.TestData.ChartRelease_LeonardoDev().ID,
				Platform:        utils.PointerTo("kubernetes"),
				GoogleProject:   utils.PointerTo("google-project"),
				InstanceName:    utils.PointerTo("instance-name"),
				DefaultDatabase: utils.PointerTo("default-database"),
			},
		},
		{
			name: "chart release invalid selector",
			fields: fields{
				DatabaseInstanceV3Create: DatabaseInstanceV3Create{
					ChartRelease: "!!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "chart release not found",
			fields: fields{
				DatabaseInstanceV3Create: DatabaseInstanceV3Create{
					ChartRelease: "does-not-exist",
				},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			d := DatabaseInstanceV3{
				CommonFields:             tt.fields.CommonFields,
				ChartReleaseInfo:         tt.fields.ChartReleaseInfo,
				DatabaseInstanceV3Create: tt.fields.DatabaseInstanceV3Create,
			}
			got, err := d.toModel(s.DB)
			if !tt.wantErr(s.T(), err, "toModel()") {
				return
			}
			s.Equalf(tt.want, got, "toModel()")
		})
	}
}

func Test_databaseInstanceFromModel(t *testing.T) {
	now := time.Now()
	type args struct {
		model models.DatabaseInstance
	}
	tests := []struct {
		name string
		args args
		want DatabaseInstanceV3
	}{
		{
			name: "empty",
			args: args{},
			want: DatabaseInstanceV3{},
		},
		{
			name: "everything",
			args: args{
				model: models.DatabaseInstance{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: now.Add(-time.Hour),
						UpdatedAt: now.Add(-time.Minute),
					},
					ChartRelease:    &models.ChartRelease{Name: "chart-release"},
					ChartReleaseID:  2,
					Platform:        utils.PointerTo("kubernetes"),
					GoogleProject:   utils.PointerTo("google-project"),
					InstanceName:    utils.PointerTo("instance-name"),
					DefaultDatabase: utils.PointerTo("default-database"),
				},
			},
			want: DatabaseInstanceV3{
				CommonFields: CommonFields{ID: 1, CreatedAt: now.Add(-time.Hour), UpdatedAt: now.Add(-time.Minute)},
				ChartReleaseInfo: &ChartReleaseV3{
					ChartReleaseV3Create: ChartReleaseV3Create{Name: "chart-release"},
				},
				DatabaseInstanceV3Create: DatabaseInstanceV3Create{
					ChartRelease: "chart-release",
					DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
						Platform:        utils.PointerTo("kubernetes"),
						GoogleProject:   utils.PointerTo("google-project"),
						InstanceName:    utils.PointerTo("instance-name"),
						DefaultDatabase: utils.PointerTo("default-database"),
					},
				},
			},
		},
		{
			name: "chart release id case",
			args: args{
				model: models.DatabaseInstance{
					ChartReleaseID: 2,
				},
			},
			want: DatabaseInstanceV3{
				DatabaseInstanceV3Create: DatabaseInstanceV3Create{
					ChartRelease: "2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, databaseInstanceFromModel(tt.args.model), "databaseInstanceFromModel(%v)", tt.args.model)
		})
	}
}

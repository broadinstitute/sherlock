package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestIncidentV3_toModel(t *testing.T) {
	now := time.Now()
	type fields struct {
		CommonFields     CommonFields
		IncidentV3Create IncidentV3Create
	}
	tests := []struct {
		name   string
		fields fields
		want   models.Incident
	}{
		{
			name: "normal",
			fields: fields{
				CommonFields: CommonFields{
					ID: 1,
				},
				IncidentV3Create: IncidentV3Create{
					IncidentV3Edit: IncidentV3Edit{
						Ticket:            utils.PointerTo("ticket"),
						Description:       utils.PointerTo("description"),
						StartedAt:         utils.PointerTo(now.Add(-1 * time.Hour)),
						RemediatedAt:      utils.PointerTo(now.Add(-30 * time.Minute)),
						ReviewCompletedAt: utils.PointerTo(now.Add(-15 * time.Minute)),
					},
				},
			},
			want: models.Incident{
				Model: gorm.Model{
					ID: 1,
				},
				Ticket:            utils.PointerTo("ticket"),
				Description:       utils.PointerTo("description"),
				StartedAt:         utils.PointerTo(now.Add(-1 * time.Hour)),
				RemediatedAt:      utils.PointerTo(now.Add(-30 * time.Minute)),
				ReviewCompletedAt: utils.PointerTo(now.Add(-15 * time.Minute)),
			},
		},
		{
			name:   "empty",
			fields: fields{},
			want:   models.Incident{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := IncidentV3{
				CommonFields:     tt.fields.CommonFields,
				IncidentV3Create: tt.fields.IncidentV3Create,
			}
			assert.Equalf(t, tt.want, i.toModel(), "toModel()")
		})
	}
}

func Test_incidentFromModel(t *testing.T) {
	now := time.Now()
	type args struct {
		model models.Incident
	}
	tests := []struct {
		name string
		args args
		want IncidentV3
	}{
		{
			name: "normal",
			args: args{
				model: models.Incident{
					Model:             gorm.Model{ID: 1},
					Ticket:            utils.PointerTo("PROD-1"),
					Description:       utils.PointerTo("An incident last month"),
					StartedAt:         utils.PointerTo(now.Add(-(24*time.Hour + 40*(24*time.Hour)))),
					RemediatedAt:      utils.PointerTo(now.Add(-(23*time.Hour + 40*(24*time.Hour)))),
					ReviewCompletedAt: utils.PointerTo(now.Add(-(22*time.Hour + 38*(24*time.Hour)))),
				},
			},
			want: IncidentV3{
				CommonFields: CommonFields{
					ID: 1,
				},
				IncidentV3Create: IncidentV3Create{
					IncidentV3Edit: IncidentV3Edit{
						Ticket:            utils.PointerTo("PROD-1"),
						Description:       utils.PointerTo("An incident last month"),
						StartedAt:         utils.PointerTo(now.Add(-(24*time.Hour + 40*(24*time.Hour)))),
						RemediatedAt:      utils.PointerTo(now.Add(-(23*time.Hour + 40*(24*time.Hour)))),
						ReviewCompletedAt: utils.PointerTo(now.Add(-(22*time.Hour + 38*(24*time.Hour)))),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, incidentFromModel(tt.args.model), "incidentFromModel(%v)", tt.args.model)
		})
	}
}

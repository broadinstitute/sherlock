package sherlock

import (
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestService_AlertV3_toModel(t *testing.T) {
	now := time.Now()
	type fields struct {
		CommonFields         CommonFields
		Uuid                 *string
		ServiceAlertV3Create ServiceAlertV3Create
	}
	tests := []struct {
		name   string
		fields fields
		want   models.ServiceAlert
	}{
		{
			name: "normal",
			fields: fields{
				CommonFields: CommonFields{
					ID: 1,
				},
				ServiceAlertV3Create: ServiceAlertV3Create{
					OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
					ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
						Title:        utils.PointerTo("Title"),
						AlertMessage: utils.PointerTo("Alert Message"),
						Link:         utils.PointerTo("NA"),
						Severity:     utils.PointerTo("minor"),
					},
				},
			},
			want: models.ServiceAlert{
				Model: gorm.Model{
					ID: 1,
				},
				utils.PointerTo(s.TestData.Environment_Dev().Name),
				Title:        utils.PointerTo("Title"),
				AlertMessage: utils.PointerTo("Alert Message"),
				Link:         utils.PointerTo("NA"),
				Severity:     utils.PointerTo("minor"),
			},
		},
		{
			name:   "empty",
			fields: fields{},
			want:   models.ServiceAlert{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := IncidentV3{
				CommonFields:         tt.fields.CommonFields,
				ServiceAlertV3Create: tt.fields.ServiceAlertV3Create,
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

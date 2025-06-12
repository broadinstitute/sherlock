package sherlock

import (
	"reflect"
	"testing"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func (s *handlerSuite) TestServiceAlertV3ToModel(t *testing.T) {
	type fields struct {
		CommonFields         CommonFields
		Uuid                 *string
		ServiceAlertV3Create ServiceAlertV3Create
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
						Title:        utils.PointerTo(*s.TestData.ServiceAlert_1().Title),
						AlertMessage: utils.PointerTo(*s.TestData.ServiceAlert_1().AlertMessage),
						Link:         utils.PointerTo(*s.TestData.ServiceAlert_1().Link),
						Severity:     utils.PointerTo(*s.TestData.ServiceAlert_1().Severity),
					},
				},
			},
			want: models.ServiceAlert{
				Model: gorm.Model{
					ID: 1,
				},
				Title:        utils.PointerTo(*s.TestData.ServiceAlert_1().Title),
				AlertMessage: utils.PointerTo(*s.TestData.ServiceAlert_1().AlertMessage),
				Link:         utils.PointerTo(*s.TestData.ServiceAlert_1().Link),
				Severity:     utils.PointerTo(*s.TestData.ServiceAlert_1().Severity),
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
			i := ServiceAlertV3{
				CommonFields:         tt.fields.CommonFields,
				Uuid:                 tt.fields.Uuid,
				ServiceAlertV3Create: tt.fields.ServiceAlertV3Create,
			}
			assert.Equalf(t, tt.want, i.toModel(tt.args.db), "toModel()")
		})
	}
}

func (s *handlerSuite) TestServiceAlertFromModel(t *testing.T) {
	type args struct {
		model models.ServiceAlert
	}
	tests := []struct {
		name string
		args args
		want ServiceAlertV3
	}{
		{
			name: "normal",
			args: args{
				model: models.ServiceAlert{
					Model:           gorm.Model{ID: 1},
					OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
					Uuid:            utils.PointerTo(*s.TestData.ServiceAlert_1().Uuid),
					Title:           utils.PointerTo(*s.TestData.ServiceAlert_1().Title),
					AlertMessage:    utils.PointerTo(*s.TestData.ServiceAlert_1().AlertMessage),
					Link:            utils.PointerTo(*s.TestData.ServiceAlert_1().Link),
					Severity:        utils.PointerTo(*s.TestData.ServiceAlert_1().Severity),
				},
			},
			want: ServiceAlertV3{
				CommonFields: CommonFields{
					ID: 1,
				},
				ServiceAlertV3Create: ServiceAlertV3Create{
					ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
						Title:        utils.PointerTo(*s.TestData.ServiceAlert_1().Title),
						AlertMessage: utils.PointerTo(*s.TestData.ServiceAlert_1().AlertMessage),
						Link:         utils.PointerTo(*s.TestData.ServiceAlert_1().Link),
						Severity:     utils.PointerTo(*s.TestData.ServiceAlert_1().Severity),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ServiceAlertFromModel(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceAlertFromModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestPagerdutyIntegrationV3Create_toModel(t *testing.T) {
	type fields struct {
		PagerdutyID                string
		PagerdutyIntegrationV3Edit PagerdutyIntegrationV3Edit
	}
	tests := []struct {
		name   string
		fields fields
		want   models.PagerdutyIntegration
	}{
		{
			name: "normal case",
			fields: fields{
				PagerdutyID: "pagerduty-id",
				PagerdutyIntegrationV3Edit: PagerdutyIntegrationV3Edit{
					Name: utils.PointerTo("name"),
					Type: utils.PointerTo("type"),
					Key:  utils.PointerTo("key"),
				},
			},
			want: models.PagerdutyIntegration{
				PagerdutyID: "pagerduty-id",
				Name:        utils.PointerTo("name"),
				Type:        utils.PointerTo("type"),
				Key:         utils.PointerTo("key"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PagerdutyIntegrationV3Create{
				PagerdutyID:                tt.fields.PagerdutyID,
				PagerdutyIntegrationV3Edit: tt.fields.PagerdutyIntegrationV3Edit,
			}
			assert.Equalf(t, tt.want, p.toModel(), "toModel()")
		})
	}
}

func TestPagerdutyIntegrationV3Edit_toModel(t *testing.T) {
	type fields struct {
		Name *string
		Key  *string
		Type *string
	}
	tests := []struct {
		name   string
		fields fields
		want   models.PagerdutyIntegration
	}{
		{
			name: "normal case",
			fields: fields{
				Name: utils.PointerTo("name"),
				Key:  utils.PointerTo("key"),
				Type: utils.PointerTo("type"),
			},
			want: models.PagerdutyIntegration{
				Name: utils.PointerTo("name"),
				Key:  utils.PointerTo("key"),
				Type: utils.PointerTo("type"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PagerdutyIntegrationV3Edit{
				Name: tt.fields.Name,
				Key:  tt.fields.Key,
				Type: tt.fields.Type,
			}
			assert.Equalf(t, tt.want, p.toModel(), "toModel()")
		})
	}
}

func TestPagerdutyIntegrationV3_toModel(t *testing.T) {
	type fields struct {
		CommonFields CommonFields
		PagerdutyID  string
		Name         *string
		Type         *string
	}
	tests := []struct {
		name   string
		fields fields
		want   models.PagerdutyIntegration
	}{
		{
			name: "normal case",
			fields: fields{
				CommonFields: CommonFields{
					ID: 1,
				},
				PagerdutyID: "pagerduty-id",
				Name:        utils.PointerTo("name"),
				Type:        utils.PointerTo("type"),
			},
			want: models.PagerdutyIntegration{
				Model:       gorm.Model{ID: 1},
				PagerdutyID: "pagerduty-id",
				Name:        utils.PointerTo("name"),
				Type:        utils.PointerTo("type"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PagerdutyIntegrationV3{
				CommonFields: tt.fields.CommonFields,
				PagerdutyID:  tt.fields.PagerdutyID,
				Name:         tt.fields.Name,
				Type:         tt.fields.Type,
			}
			assert.Equalf(t, tt.want, p.toModel(), "toModel()")
		})
	}
}

func Test_pagerdutyIntegrationFromModel(t *testing.T) {
	type args struct {
		model models.PagerdutyIntegration
	}
	tests := []struct {
		name string
		args args
		want PagerdutyIntegrationV3
	}{
		{
			name: "normal case",
			args: args{
				model: models.PagerdutyIntegration{
					Model:       gorm.Model{ID: 1},
					PagerdutyID: "pagerduty-id",
					Name:        utils.PointerTo("name"),
					Type:        utils.PointerTo("type"),
					Key:         utils.PointerTo("key"),
				},
			},
			want: PagerdutyIntegrationV3{
				CommonFields: CommonFields{
					ID: 1,
				},
				PagerdutyID: "pagerduty-id",
				Name:        utils.PointerTo("name"),
				Type:        utils.PointerTo("type"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pagerdutyIntegrationFromModel(tt.args.model), "pagerdutyIntegrationFromModel(%v)", tt.args.model)
		})
	}
}

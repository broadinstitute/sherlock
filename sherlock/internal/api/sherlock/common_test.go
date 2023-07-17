package sherlock

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_commonFieldsFromGormModel(t *testing.T) {
	t1 := time.Now().Add(-time.Minute)
	t2 := time.Now().Add(-time.Second)
	type args struct {
		model gorm.Model
	}
	tests := []struct {
		name string
		args args
		want commonFields
	}{
		{
			name: "equal",
			args: args{model: gorm.Model{
				ID:        1,
				CreatedAt: t1,
				UpdatedAt: t2,
			}},
			want: commonFields{
				ID:        1,
				CreatedAt: t1,
				UpdatedAt: t2,
			},
		},
		{
			name: "ignores deleted at",
			args: args{model: gorm.Model{
				ID:        1,
				CreatedAt: t1,
				UpdatedAt: t2,
				DeletedAt: gorm.DeletedAt{
					Valid: true,
					Time:  time.Now(),
				},
			}},
			want: commonFields{
				ID:        1,
				CreatedAt: t1,
				UpdatedAt: t2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, commonFieldsFromGormModel(tt.args.model), "commonFieldsFromGormModel(%v)", tt.args.model)
		})
	}
}

func Test_commonFields_toGormModel(t *testing.T) {
	t1 := time.Now().Add(-time.Minute)
	t2 := time.Now().Add(-time.Second)
	type fields struct {
		ID        uint
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   gorm.Model
	}{
		{
			name: "equal",
			fields: fields{
				ID:        1,
				CreatedAt: t1,
				UpdatedAt: t2,
			},
			want: gorm.Model{
				ID:        1,
				CreatedAt: t1,
				UpdatedAt: t2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := commonFields{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			assert.Equalf(t, tt.want, f.toGormModel(), "toGormModel()")
		})
	}
}

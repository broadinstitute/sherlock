package v2models

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_ciIdentifierSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    CiIdentifier
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid!"},
			wantErr: true,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			want: CiIdentifier{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ciIdentifierSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ciIdentifierSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_ciIdentifierToSelectors(t *testing.T) {
	type args struct {
		ciIdentifier *CiIdentifier
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil",
			want: nil,
		},
		{
			name: "none",
			args: args{ciIdentifier: &CiIdentifier{}},
			want: nil,
		},
		{
			name: "id",
			args: args{ciIdentifier: &CiIdentifier{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "resource type + resource id",
			args: args{ciIdentifier: &CiIdentifier{ResourceType: "type", ResourceID: 456}},
			want: []string{"type/456"},
		},
		{
			name: "id, resource type + resource id",
			args: args{ciIdentifier: &CiIdentifier{
				Model:        gorm.Model{ID: 123},
				ResourceType: "type", ResourceID: 456,
			}},
			want: []string{"123", "type/456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ciIdentifierToSelectors(tt.args.ciIdentifier)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_rejectDuplicateCiIdentifier(t *testing.T) {
	type args struct {
		existing *CiIdentifier
		new      *CiIdentifier
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				existing: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
				},
				new: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
				},
			},
			wantErr: false,
		},
		{
			name: "mismatched type",
			args: args{
				existing: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
				},
				new: &CiIdentifier{
					ResourceType: "b",
					ResourceID:   123,
				},
			},
			wantErr: true,
		},
		{
			name: "mismatched id",
			args: args{
				existing: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
				},
				new: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   124,
				},
			},
			wantErr: true,
		},
		{
			name: "runs in new are invalid",
			args: args{
				existing: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
				},
				new: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
					CiRuns: []*CiRun{
						{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "runs in existing are okay",
			args: args{
				existing: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
					CiRuns: []*CiRun{
						{},
					},
				},
				new: &CiIdentifier{
					ResourceType: "a",
					ResourceID:   123,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rejectDuplicateCiIdentifier(tt.args.existing, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("rejectDuplicateCiIdentifier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateCiIdentifier(t *testing.T) {
	type args struct {
		ciIdentifier *CiIdentifier
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{ciIdentifier: nil},
			wantErr: true,
		},
		{
			name: "no resource type",
			args: args{ciIdentifier: &CiIdentifier{
				ResourceID: 123,
			}},
			wantErr: true,
		},
		{
			name: "no resource ID",
			args: args{ciIdentifier: &CiIdentifier{
				ResourceType: "type",
			}},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{ciIdentifier: &CiIdentifier{
				ResourceID:   123,
				ResourceType: "type",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCiIdentifier(tt.args.ciIdentifier); (err != nil) != tt.wantErr {
				t.Errorf("validateCiIdentifier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ciIdentifierErrorIfForbidden(t *testing.T) {
	type args struct {
		db           *gorm.DB
		ciIdentifier *CiIdentifier
		actionType   model_actions.ActionType
		user         *auth_models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid if create",
			args:    args{ciIdentifier: &CiIdentifier{}, actionType: model_actions.CREATE},
			wantErr: false,
		},
		{
			name:    "valid if edit",
			args:    args{ciIdentifier: &CiIdentifier{}, actionType: model_actions.EDIT},
			wantErr: false,
		},
		{
			name:    "invalid if delete",
			args:    args{ciIdentifier: &CiIdentifier{}, actionType: model_actions.DELETE},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ciIdentifierErrorIfForbidden(tt.args.db, tt.args.ciIdentifier, tt.args.actionType, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("ciIdentifierErrorIfForbidden() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

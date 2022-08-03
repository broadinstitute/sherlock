package v2models

import (
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_environmentSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
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
			want: Environment{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name: "valid name",
			args: args{selector: "foo-bar-2"},
			want: Environment{Name: "foo-bar-2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := environmentSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("environmentSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_environmentToSelectors(t *testing.T) {
	type args struct {
		environment Environment
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{},
			want: nil,
		},
		{
			name: "with name",
			args: args{environment: Environment{
				Name: "foobar",
			}},
			want: []string{"foobar"},
		},
		{
			name: "with id",
			args: args{environment: Environment{
				Model: gorm.Model{ID: 123},
			}},
			want: []string{"123"},
		},
		{
			name: "with name and id",
			args: args{environment: Environment{
				Model: gorm.Model{ID: 123},
				Name:  "foobar",
			}},
			want: []string{"foobar", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := environmentToSelectors(tt.args.environment)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_environmentRequiresSuitability(t *testing.T) {
	tru := true
	fal := false
	type args struct {
		db          *gorm.DB
		environment Environment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "does require",
			args: args{environment: Environment{RequiresSuitability: &tru}},
			want: true,
		},
		{
			name: "does not require",
			args: args{environment: Environment{RequiresSuitability: &fal}},
			want: false,
		},
		{
			name: "fail safe",
			args: args{environment: Environment{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := environmentRequiresSuitability(tt.args.db, tt.args.environment)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateEnvironment(t *testing.T) {
	type args struct {
		environment Environment
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid template",
			args: args{environment: Environment{
				Name:      "foobar",
				Lifecycle: "template",
			}},
			wantErr: false,
		},
		{
			name: "valid dynamic",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "dynamic",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: false,
		},
		{
			name: "valid static",
			args: args{environment: Environment{
				Name:                "foobar",
				Lifecycle:           "static",
				Base:                "base",
				DefaultClusterID:    testutils.PointerTo[uint](456),
				DefaultNamespace:    testutils.PointerTo("namespace"),
				Owner:               testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: false,
		},
		{
			name: "no name",
			args: args{environment: Environment{
				Lifecycle: "template",
			}},
			wantErr: true,
		},
		{
			name: "no lifecycle",
			args: args{environment: Environment{
				Name: "foobar",
			}},
			wantErr: true,
		},
		{
			name: "template with template",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "template",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
			}},
			wantErr: true,
		},
		{
			name: "dynamic without template",
			args: args{environment: Environment{
				Name:                "foobar",
				Lifecycle:           "dynamic",
				Base:                "base",
				DefaultClusterID:    testutils.PointerTo[uint](456),
				DefaultNamespace:    testutils.PointerTo("namespace"),
				Owner:               testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability: testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no base",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "dynamic",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no default cluster",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "dynamic",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no default namespace",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "dynamic",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no owner",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "dynamic",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no requires suitability",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "dynamic",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
			}},
			wantErr: true,
		},
		{
			name: "static no base",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "static",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "static no default cluster",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "static",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "static no default namespace",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "static",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "static no owner",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "static",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				RequiresSuitability:   testutils.PointerTo(false),
			}},
			wantErr: true,
		},
		{
			name: "static no requires suitability",
			args: args{environment: Environment{
				Name:                  "foobar",
				Lifecycle:             "static",
				TemplateEnvironmentID: testutils.PointerTo[uint](123),
				Base:                  "base",
				DefaultClusterID:      testutils.PointerTo[uint](456),
				DefaultNamespace:      testutils.PointerTo("namespace"),
				Owner:                 testutils.PointerTo("fake@broadinstitute.org"),
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateEnvironment(tt.args.environment); (err != nil) != tt.wantErr {
				t.Errorf("validateEnvironment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

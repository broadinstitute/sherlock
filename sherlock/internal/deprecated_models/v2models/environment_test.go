package v2models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models/environment"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
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
		{
			name:    "invalid name",
			args:    args{selector: "foooooooooooooooooooooooooooooooo"},
			wantErr: true,
		},
		{
			name: "valid resource prefix",
			args: args{selector: "resource-prefix/a2k3"},
			want: Environment{UniqueResourcePrefix: "a2k3"},
		},
		{
			name:    "invalid resource prefix start",
			args:    args{selector: "blah/a2k3"},
			wantErr: true,
		},
		{
			name:    "invalid resource prefix end, caps",
			args:    args{selector: "resource-prefix/Caps"},
			wantErr: true,
		},
		{
			name:    "invalid resource prefix end, len",
			args:    args{selector: "resource-prefix/toolong"},
			wantErr: true,
		},
		{
			name:    "invalid resource prefix end, number start",
			args:    args{selector: "resource-prefix/1234"},
			wantErr: true,
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
		environment *Environment
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
			args: args{environment: &Environment{
				Name: "foobar",
			}},
			want: []string{"foobar"},
		},
		{
			name: "with id",
			args: args{environment: &Environment{
				Model: gorm.Model{ID: 123},
			}},
			want: []string{"123"},
		},
		{
			name: "with resource prefix",
			args: args{environment: &Environment{
				UniqueResourcePrefix: "a1b2",
			}},
			want: []string{"resource-prefix/a1b2"},
		},
		{
			name: "with name and id",
			args: args{environment: &Environment{
				Model: gorm.Model{ID: 123},
				Name:  "foobar",
			}},
			want: []string{"foobar", "123"},
		},
		{
			name: "with name and id and resource prefix",
			args: args{environment: &Environment{
				Model:                gorm.Model{ID: 123},
				Name:                 "foobar",
				UniqueResourcePrefix: "a1b2",
			}},
			want: []string{"foobar", "123", "resource-prefix/a1b2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := environmentToSelectors(tt.args.environment)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateEnvironment(t *testing.T) {
	type args struct {
		environment *Environment
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid template",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				DefaultNamespace:           "terra-foobar",
				OwnerID:                    utils.PointerTo[uint](987),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: false,
		},
		{
			name: "valid dynamic",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: false,
		},
		{
			name: "valid static",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: false,
		},
		{
			name: "valid template with legacy owner",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				DefaultNamespace:           "terra-foobar",
				LegacyOwner:                utils.PointerTo("fake@broadinstitute.org"),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: false,
		},
		{
			name: "valid dynamic with legacy owner",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				LegacyOwner:                utils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: false,
		},
		{
			name: "valid static with legacy owner",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				LegacyOwner:                utils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: false,
		},
		{
			name: "no name",
			args: args{environment: &Environment{
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "invalid name",
			args: args{environment: &Environment{
				Name:                       "CAPITAL-LETTERS-ARE-BAD",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "no helmfileRef",
			args: args{environment: &Environment{
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				Name:                       "foobar",
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "no lifecycle",
			args: args{environment: &Environment{
				Name:                       "foobar",
				UniqueResourcePrefix:       "a1b2",
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "no default firecloud develop ref",
			args: args{environment: &Environment{
				Name:                 "foobar",
				UniqueResourcePrefix: "a1b2",
				HelmfileRef:          utils.PointerTo("HEAD"),
			}},
			wantErr: true,
		},
		{
			name: "template with template",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "dynamic without template",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no base",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no default cluster",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no default namespace",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no owner",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "dynamic no requires suitability",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "static no base",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "static no default cluster",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "static no default namespace",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "static no owner",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "static no requires suitability",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "no resource prefix",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			}},
			wantErr: true,
		},
		{
			name: "template with prevent deletion",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				DefaultNamespace:           "terra-foobar",
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				PreventDeletion:            utils.PointerTo(true),
			}},
			wantErr: true,
		},
		{
			name: "static with auto-delete",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				AutoDelete: utils.PointerTo(environment.AutoDelete{
					Enabled: true,
					After:   time.Now(),
				}),
			}},
			wantErr: true,
		},
		{
			name: "dynamic with prevent deletion",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				PreventDeletion:            utils.PointerTo(true),
			}},
			wantErr: false,
		},
		{
			name: "dynamic with auto-delete",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				AutoDelete: utils.PointerTo(environment.AutoDelete{
					Enabled: true,
					After:   time.Now(),
				}),
			}},
			wantErr: false,
		},
		{
			name: "dynamic with prevent deletion and auto-delete",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				PreventDeletion:            utils.PointerTo(true),
				AutoDelete: utils.PointerTo(environment.AutoDelete{
					Enabled: true,
					After:   time.Now(),
				}),
			}},
			wantErr: true,
		},
		{
			name: "invalid static with offline true",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(true),
			}},
			wantErr: true,
		},
		{
			name: "valid static with offline false",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(false),
			}},
			wantErr: false,
		},
		{
			name: "invalid template with offline true",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(true),
			}},
			wantErr: true,
		},
		{
			name: "valid template with offline false",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "template",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(false),
			}},
			wantErr: false,
		},
		{
			name: "valid dynamic with offline true",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(true),
			}},
			wantErr: false,
		},
		{
			name: "valid dynamic with offline false",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(false),
			}},
			wantErr: false,
		},
		{
			name: "invalid static with begin-offline schedule true",
			args: args{environment: &Environment{
				Name:                        "foobar",
				Lifecycle:                   "static",
				UniqueResourcePrefix:        "a1b2",
				Base:                        "base",
				DefaultClusterID:            utils.PointerTo[uint](456),
				DefaultNamespace:            "namespace",
				LegacyOwner:                 utils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:         utils.PointerTo(false),
				HelmfileRef:                 utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef:  utils.PointerTo("dev"),
				Offline:                     utils.PointerTo(false),
				OfflineScheduleBeginEnabled: utils.PointerTo(true),
				OfflineScheduleBeginTime:    utils.PointerTo(time.Now().Format(time.RFC3339)),
			}},
			wantErr: true,
		},
		{
			name: "invalid static with end-offline schedule true",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "static",
				UniqueResourcePrefix:       "a1b2",
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(false),
				OfflineScheduleEndEnabled:  utils.PointerTo(true),
				OfflineScheduleEndTime:     utils.PointerTo(time.Now().Format(time.RFC3339)),
			}},
			wantErr: true,
		},
		{
			name: "invalid dynamic with begin-offline schedule true",
			args: args{environment: &Environment{
				Name:                        "foobar",
				Lifecycle:                   "dynamic",
				UniqueResourcePrefix:        "a1b2",
				TemplateEnvironmentID:       utils.PointerTo[uint](123),
				Base:                        "base",
				DefaultClusterID:            utils.PointerTo[uint](456),
				DefaultNamespace:            "namespace",
				LegacyOwner:                 utils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:         utils.PointerTo(false),
				HelmfileRef:                 utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef:  utils.PointerTo("dev"),
				Offline:                     utils.PointerTo(false),
				OfflineScheduleBeginEnabled: utils.PointerTo(true),
			}},
			wantErr: true,
		},
		{
			name: "invalid dynamic with end-offline schedule true",
			args: args{environment: &Environment{
				Name:                       "foobar",
				Lifecycle:                  "dynamic",
				UniqueResourcePrefix:       "a1b2",
				TemplateEnvironmentID:      utils.PointerTo[uint](123),
				Base:                       "base",
				DefaultClusterID:           utils.PointerTo[uint](456),
				DefaultNamespace:           "namespace",
				OwnerID:                    utils.PointerTo[uint](987),
				RequiresSuitability:        utils.PointerTo(false),
				HelmfileRef:                utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
				Offline:                    utils.PointerTo(false),
				OfflineScheduleEndEnabled:  utils.PointerTo(true),
			}},
			wantErr: true,
		},
		{
			name: "valid dynamic with offline schedule true",
			args: args{environment: &Environment{
				Name:                        "foobar",
				Lifecycle:                   "dynamic",
				UniqueResourcePrefix:        "a1b2",
				TemplateEnvironmentID:       utils.PointerTo[uint](123),
				Base:                        "base",
				DefaultClusterID:            utils.PointerTo[uint](456),
				DefaultNamespace:            "namespace",
				LegacyOwner:                 utils.PointerTo("fake@broadinstitute.org"),
				RequiresSuitability:         utils.PointerTo(false),
				HelmfileRef:                 utils.PointerTo("HEAD"),
				DefaultFirecloudDevelopRef:  utils.PointerTo("dev"),
				Offline:                     utils.PointerTo(false),
				OfflineScheduleBeginEnabled: utils.PointerTo(true),
				OfflineScheduleBeginTime:    utils.PointerTo(time.Now().Format(time.RFC3339)),
				OfflineScheduleEndEnabled:   utils.PointerTo(true),
				OfflineScheduleEndTime:      utils.PointerTo(time.Now().Format(time.RFC3339)),
			}},
			wantErr: false,
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

func TestEnvironment_GetCiIdentifier(t *testing.T) {
	type fields struct {
		Model                       gorm.Model
		CiIdentifier                *CiIdentifier
		Base                        string
		Lifecycle                   string
		Name                        string
		NamePrefix                  string
		TemplateEnvironment         *Environment
		TemplateEnvironmentID       *uint
		ValuesName                  string
		AutoPopulateChartReleases   *bool
		UniqueResourcePrefix        string
		DefaultNamespace            string
		DefaultCluster              *Cluster
		DefaultClusterID            *uint
		DefaultFirecloudDevelopRef  *string
		Owner                       *User
		OwnerID                     *uint
		LegacyOwner                 *string
		RequiresSuitability         *bool
		BaseDomain                  *string
		NamePrefixesDomain          *bool
		HelmfileRef                 *string
		PreventDeletion             *bool
		AutoDelete                  *environment.AutoDelete
		Description                 *string
		PagerdutyIntegration        *PagerdutyIntegration
		PagerdutyIntegrationID      *uint
		Offline                     *bool
		OfflineScheduleBeginEnabled *bool
		OfflineScheduleBeginTime    *string
		OfflineScheduleEndEnabled   *bool
		OfflineScheduleEndTime      *string
		OfflineScheduleEndWeekends  *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *CiIdentifier
	}{
		{
			name: "returns existing",
			fields: fields{
				CiIdentifier: &CiIdentifier{
					Model: gorm.Model{
						ID: 123,
					},
				},
			},
			want: &CiIdentifier{
				Model: gorm.Model{
					ID: 123,
				},
			},
		},
		{
			name: "returns generated if no existing",
			fields: fields{
				Model: gorm.Model{
					ID: 123,
				},
			},
			want: &CiIdentifier{
				ResourceType: "environment",
				ResourceID:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Environment{
				Model:                       tt.fields.Model,
				CiIdentifier:                tt.fields.CiIdentifier,
				Base:                        tt.fields.Base,
				Lifecycle:                   tt.fields.Lifecycle,
				Name:                        tt.fields.Name,
				NamePrefix:                  tt.fields.NamePrefix,
				TemplateEnvironment:         tt.fields.TemplateEnvironment,
				TemplateEnvironmentID:       tt.fields.TemplateEnvironmentID,
				ValuesName:                  tt.fields.ValuesName,
				AutoPopulateChartReleases:   tt.fields.AutoPopulateChartReleases,
				UniqueResourcePrefix:        tt.fields.UniqueResourcePrefix,
				DefaultNamespace:            tt.fields.DefaultNamespace,
				DefaultCluster:              tt.fields.DefaultCluster,
				DefaultClusterID:            tt.fields.DefaultClusterID,
				DefaultFirecloudDevelopRef:  tt.fields.DefaultFirecloudDevelopRef,
				Owner:                       tt.fields.Owner,
				OwnerID:                     tt.fields.OwnerID,
				LegacyOwner:                 tt.fields.LegacyOwner,
				RequiresSuitability:         tt.fields.RequiresSuitability,
				BaseDomain:                  tt.fields.BaseDomain,
				NamePrefixesDomain:          tt.fields.NamePrefixesDomain,
				HelmfileRef:                 tt.fields.HelmfileRef,
				PreventDeletion:             tt.fields.PreventDeletion,
				AutoDelete:                  tt.fields.AutoDelete,
				Description:                 tt.fields.Description,
				PagerdutyIntegration:        tt.fields.PagerdutyIntegration,
				PagerdutyIntegrationID:      tt.fields.PagerdutyIntegrationID,
				Offline:                     tt.fields.Offline,
				OfflineScheduleBeginEnabled: tt.fields.OfflineScheduleBeginEnabled,
				OfflineScheduleBeginTime:    tt.fields.OfflineScheduleBeginTime,
				OfflineScheduleEndEnabled:   tt.fields.OfflineScheduleEndEnabled,
				OfflineScheduleEndTime:      tt.fields.OfflineScheduleEndTime,
				OfflineScheduleEndWeekends:  tt.fields.OfflineScheduleEndWeekends,
			}
			if got := e.GetCiIdentifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCiIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

package v2models

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"testing"
)

func Test_userSelectorToQuery(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    User
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
			want: User{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name: "email",
			args: args{selector: "foo@bar.com"},
			want: User{StoredControlledUserFields: auth_models.StoredControlledUserFields{Email: "foo@bar.com"}},
		},
		{
			name: "google subject id",
			args: args{selector: "google-id/blah"},
			want: User{StoredControlledUserFields: auth_models.StoredControlledUserFields{GoogleID: "blah"}},
		},
		{
			name: "github username",
			args: args{selector: "github/blah"},
			want: User{StoredControlledUserFields: auth_models.StoredControlledUserFields{GithubUsername: testutils.PointerTo("blah")}},
		},
		{
			name: "github id",
			args: args{selector: "github-id/blah"},
			want: User{StoredControlledUserFields: auth_models.StoredControlledUserFields{GithubID: testutils.PointerTo("blah")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userSelectorToQuery(nil, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_userToSelectors(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{user: &User{}},
			want: nil,
		},
		{
			name: "id",
			args: args{user: &User{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "id, google id",
			args: args{user: &User{
				Model: gorm.Model{ID: 123},
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					GoogleID: "some google id here",
				},
			}},
			want: []string{"123", "google-id/some google id here"},
		},
		{
			name: "id, email",
			args: args{user: &User{
				Model: gorm.Model{ID: 123},
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email: "fake@example.com",
				},
			}},
			want: []string{"123", "fake@example.com"},
		},
		{
			name: "id, email, google id",
			args: args{user: &User{
				Model: gorm.Model{ID: 123},
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:    "fake@example.com",
					GoogleID: "some google id here",
				},
			}},
			want: []string{"123", "fake@example.com", "google-id/some google id here"},
		},
		{
			name: "id, email, google id, github id",
			args: args{user: &User{
				Model: gorm.Model{ID: 123},
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:    "fake@example.com",
					GoogleID: "some google id here",
					GithubID: testutils.PointerTo("some github id here"),
				},
			}},
			want: []string{"123", "fake@example.com", "google-id/some google id here", "github-id/some github id here"},
		},
		{
			name: "id, email, google id, github",
			args: args{user: &User{
				Model: gorm.Model{ID: 123},
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:          "fake@example.com",
					GoogleID:       "some google id here",
					GithubUsername: testutils.PointerTo("fake-github"),
				},
			}},
			want: []string{"123", "fake@example.com", "google-id/some google id here", "github/fake-github"},
		},
		{
			name: "id, email, google id, github, github id",
			args: args{user: &User{
				Model: gorm.Model{ID: 123},
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:          "fake@example.com",
					GoogleID:       "some google id here",
					GithubUsername: testutils.PointerTo("fake-github"),
					GithubID:       testutils.PointerTo("some github id here"),
				},
			}},
			want: []string{"123", "fake@example.com", "google-id/some google id here", "github/fake-github", "github-id/some github id here"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userToSelectors(tt.args.user)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_userErrorIfForbidden(t *testing.T) {
	type args struct {
		modelUser *User
		action    model_actions.ActionType
		user      *auth_models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create disallowed if user passed",
			args: args{
				modelUser: nil,
				action:    model_actions.CREATE,
				user:      &auth_models.User{},
			},
			wantErr: true,
		},
		{
			name: "create allowed if no user passed",
			args: args{
				modelUser: nil,
				action:    model_actions.CREATE,
				user:      nil,
			},
			wantErr: false,
		},
		{
			name: "edit disallowed if emails don't match",
			args: args{
				modelUser: &User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
				},
				action: model_actions.EDIT,
				user: &auth_models.User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "bar@example.com",
					},
					AuthMethod: auth_models.AuthMethodIAP,
				},
			},
			wantErr: true,
		},
		{
			name: "edit disallowed if GHA involved",
			args: args{
				modelUser: &User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
				},
				action: model_actions.EDIT,
				user: &auth_models.User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
					AuthMethod: auth_models.AuthMethodGHA,
					Via: &auth_models.User{
						StoredControlledUserFields: auth_models.StoredControlledUserFields{
							Email: "foo@example.com",
						},
						AuthMethod: auth_models.AuthMethodIAP,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "edit disallowed if GHA involved, other order",
			args: args{
				modelUser: &User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
				},
				action: model_actions.EDIT,
				user: &auth_models.User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
					AuthMethod: auth_models.AuthMethodIAP,
					Via: &auth_models.User{
						StoredControlledUserFields: auth_models.StoredControlledUserFields{
							Email: "foo@example.com",
						},
						AuthMethod: auth_models.AuthMethodGHA,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "edit allowed",
			args: args{
				modelUser: &User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
				},
				action: model_actions.EDIT,
				user: &auth_models.User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
					AuthMethod: auth_models.AuthMethodIAP,
				},
			},
			wantErr: false,
		},
		{
			name: "delete always not allowed",
			args: args{
				modelUser: &User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
				},
				action: model_actions.DELETE,
				user: &auth_models.User{
					StoredControlledUserFields: auth_models.StoredControlledUserFields{
						Email: "foo@example.com",
					},
					AuthMethod: auth_models.AuthMethodIAP,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userErrorIfForbidden(nil, tt.args.modelUser, tt.args.action, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userErrorIfForbidden() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateUser(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{user: nil},
			wantErr: true,
		},
		{
			name: "no email",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					GoogleID: "some id here",
				},
			}},
			wantErr: true,
		},
		{
			name: "bad email",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:    "fake",
					GoogleID: "some id here",
				},
			}},
			wantErr: true,
		},
		{
			name: "no google id",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email: "fake@broadinstitute.org",
				},
			}},
			wantErr: true,
		},
		{
			name: "valid without github",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:    "fake@broadinstitute.org",
					GoogleID: "some id here",
				},
			}},
			wantErr: false,
		},
		{
			name: "only half of github",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:          "fake@broadinstitute.org",
					GoogleID:       "some id here",
					GithubUsername: testutils.PointerTo("blah"),
				},
			}},
			wantErr: true,
		},
		{
			name: "the other half of github",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:    "fake@broadinstitute.org",
					GoogleID: "some id here",
					GithubID: testutils.PointerTo("bleh"),
				},
			}},
			wantErr: true,
		},
		{
			name: "valid with github",
			args: args{user: &User{
				StoredControlledUserFields: auth_models.StoredControlledUserFields{
					Email:          "fake@broadinstitute.org",
					GoogleID:       "some id here",
					GithubUsername: testutils.PointerTo("blah"),
					GithubID:       testutils.PointerTo("bleh"),
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("validateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

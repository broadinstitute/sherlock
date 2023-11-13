package gha_oidc_claims

import "testing"

func TestClaims_TrimmedRepositoryName(t *testing.T) {
	type fields struct {
		Actor           string
		ActorID         string
		RepositoryOwner string
		Repository      string
		WorkflowRef     string
		RunID           string
		RunAttempt      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "trims owner",
			fields: fields{
				RepositoryOwner: "abc",
				Repository:      "abc/def",
			},
			want: "def",
		},
		{
			name: "if owner empty",
			fields: fields{
				Repository: "abc/def",
			},
			want: "abc/def",
		},
		{
			name: "if repository empty",
			fields: fields{
				RepositoryOwner: "abc",
			},
			want: "",
		},
		{
			name: "if mismatched",
			fields: fields{
				RepositoryOwner: "abc",
				Repository:      "def/ghi",
			},
			want: "def/ghi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Claims{
				Actor:           tt.fields.Actor,
				ActorID:         tt.fields.ActorID,
				RepositoryOwner: tt.fields.RepositoryOwner,
				Repository:      tt.fields.Repository,
				WorkflowRef:     tt.fields.WorkflowRef,
				RunID:           tt.fields.RunID,
				RunAttempt:      tt.fields.RunAttempt,
			}
			if got := c.TrimmedRepositoryName(); got != tt.want {
				t.Errorf("TrimmedRepositoryName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaims_TrimmedWorkflowPath(t *testing.T) {
	type fields struct {
		Actor           string
		ActorID         string
		RepositoryOwner string
		Repository      string
		WorkflowRef     string
		RunID           string
		RunAttempt      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "trims repository",
			fields: fields{
				Repository:  "abc/def",
				WorkflowRef: "abc/def/ghi",
			},
			want: "ghi",
		},
		{
			name: "trims after @",
			fields: fields{
				WorkflowRef: "123@456",
			},
			want: "123",
		},
		{
			name: "together",
			fields: fields{
				Repository:  "broadinstitute/terra-github-workflows",
				WorkflowRef: "broadinstitute/terra-github-workflows/.github/workflows/bee-create.yaml@refs/heads/main",
			},
			want: ".github/workflows/bee-create.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Claims{
				Actor:           tt.fields.Actor,
				ActorID:         tt.fields.ActorID,
				RepositoryOwner: tt.fields.RepositoryOwner,
				Repository:      tt.fields.Repository,
				WorkflowRef:     tt.fields.WorkflowRef,
				RunID:           tt.fields.RunID,
				RunAttempt:      tt.fields.RunAttempt,
			}
			if got := c.TrimmedWorkflowPath(); got != tt.want {
				t.Errorf("TrimmedWorkflowPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaims_WorkflowURL(t *testing.T) {
	type fields struct {
		Actor           string
		ActorID         string
		RepositoryOwner string
		Repository      string
		WorkflowRef     string
		RunID           string
		RunAttempt      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   "https://github.com//actions/runs//attempts/",
		},
		{
			name: "normal",
			fields: fields{
				Repository: "broadinstitute/terra-github-workflows",
				RunID:      "123456",
				RunAttempt: "1",
			},
			want: "https://github.com/broadinstitute/terra-github-workflows/actions/runs/123456/attempts/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Claims{
				Actor:           tt.fields.Actor,
				ActorID:         tt.fields.ActorID,
				RepositoryOwner: tt.fields.RepositoryOwner,
				Repository:      tt.fields.Repository,
				WorkflowRef:     tt.fields.WorkflowRef,
				RunID:           tt.fields.RunID,
				RunAttempt:      tt.fields.RunAttempt,
			}
			if got := c.WorkflowURL(); got != tt.want {
				t.Errorf("WorkflowURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

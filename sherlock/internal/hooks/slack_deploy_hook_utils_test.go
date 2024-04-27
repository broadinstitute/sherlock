package hooks

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
	"testing"
)

func Test_slackDeployHookParseStatus(t *testing.T) {
	type args struct {
		rawStatus string
	}
	tests := []struct {
		name           string
		args           args
		wantStatus     string
		wantEmoji      string
		wantHasFailure bool
	}{
		{
			name: "empty",
			args: args{
				rawStatus: "",
			},
			wantStatus:     "",
			wantEmoji:      config.Config.String("slack.emoji.unknown"),
			wantHasFailure: false,
		},
		{
			name: "queued",
			args: args{
				rawStatus: "queued",
			},
			wantStatus:     "Waiting...",
			wantEmoji:      config.Config.String("slack.emoji.beehiveWaiting"),
			wantHasFailure: false,
		},
		{
			name: "running",
			args: args{
				rawStatus: "running: Progressing: blah blah",
			},
			wantStatus:     "Progressing...",
			wantEmoji:      config.Config.String("slack.emoji.beehiveLoading"),
			wantHasFailure: false,
		},
		{
			name: "running with progress",
			args: args{
				rawStatus: "running: Progressing: blah blah 1 out of 3 replicas are good to go pal",
			},
			wantStatus:     "Progressing... 33%",
			wantEmoji:      config.Config.String("slack.emoji.beehiveLoading"),
			wantHasFailure: false,
		},
		{
			name: "running with slash progress",
			args: args{
				rawStatus: "running: Progressing: blah blah 1/3 replicas are good to go pal",
			},
			wantStatus:     "Progressing... 33%",
			wantEmoji:      config.Config.String("slack.emoji.beehiveLoading"),
			wantHasFailure: false,
		},
		{
			name: "running with 'of' progress",
			args: args{
				rawStatus: "running: Progressing: blah blah 1 of 3 replicas are good to go pal",
			},
			wantStatus:     "Progressing... 33%",
			wantEmoji:      config.Config.String("slack.emoji.beehiveLoading"),
			wantHasFailure: false,
		},
		{
			name: "running with a division by zero progress",
			args: args{
				rawStatus: "running: Progressing: putting garbage in 3 out of 0 times",
			},
			wantStatus:     "Progressing...",
			wantEmoji:      config.Config.String("slack.emoji.beehiveLoading"),
			wantHasFailure: false,
		},
		{
			name: "in_progress",
			args: args{
				rawStatus: "in_progress",
			},
			wantStatus:     "Progressing...",
			wantEmoji:      config.Config.String("slack.emoji.beehiveLoading"),
			wantHasFailure: false,
		},
		{
			name: "success",
			args: args{
				rawStatus: "success: blah blah",
			},
			wantStatus:     "Success",
			wantEmoji:      config.Config.String("slack.emoji.succeeded"),
			wantHasFailure: false,
		},
		{
			name: "error",
			args: args{
				rawStatus: "error: blah blah",
			},
			wantStatus:     "Failed",
			wantEmoji:      config.Config.String("slack.emoji.failed"),
			wantHasFailure: true,
		},
		{
			name: "failure",
			args: args{
				rawStatus: "failure",
			},
			wantStatus:     "Failed",
			wantEmoji:      config.Config.String("slack.emoji.failed"),
			wantHasFailure: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, gotEmoji, gotHasFailure := slackDeployHookParseStatus(tt.args.rawStatus)
			if gotStatus != tt.wantStatus {
				t.Errorf("slackDeployHookParseStatus() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotEmoji != tt.wantEmoji {
				t.Errorf("slackDeployHookParseStatus() gotEmoji = %v, want %v", gotEmoji, tt.wantEmoji)
			}
			if gotHasFailure != tt.wantHasFailure {
				t.Errorf("slackDeployHookParseStatus() gotHasFailure = %v, want %v", gotHasFailure, tt.wantHasFailure)
			}
		})
	}
}

func Test_slackDeployHookBeehiveLink(t *testing.T) {
	type args struct {
		changesets []models.Changeset
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				changesets: []models.Changeset{
					{Model: gorm.Model{ID: 1}},
					{Model: gorm.Model{ID: 2}},
				},
			},
			want: "https://beehive.dsp-devops-prod.broadinstitute.org/review-changesets?changeset=1&changeset=2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slackDeployHookBeehiveLink(tt.args.changesets); got != tt.want {
				t.Errorf("slackDeployHookBeehiveLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slackDeployHookSummarizeUsers(t *testing.T) {
	type args struct {
		users         []models.User
		mentionPeople bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "mention people",
			args: args{
				users: []models.User{
					{Email: "test-1@example.com"},
					{Email: "test-2@example.com", SlackID: utils.PointerTo("test-2-id")},
					{Email: "test-3@gserviceaccount.com"},
					{Email: "test-4@gserviceaccount.com", SlackID: utils.PointerTo("test-4-id")},
				},
				mentionPeople: true,
			},
			want: "By <https://broad.io/beehive/r/user/test-1@example.com|test-1>, <@test-2-id>, <https://broad.io/beehive/r/user/test-3@gserviceaccount.com|test-3> (service account), <@test-4-id> (service account)",
		},
		{
			name: "don't mention people",
			args: args{
				users: []models.User{
					{Email: "test-1@example.com"},
					{Email: "test-2@example.com", SlackID: utils.PointerTo("test-2-id")},
					{Email: "test-3@gserviceaccount.com"},
					{Email: "test-4@gserviceaccount.com", SlackID: utils.PointerTo("test-4-id")},
				},
				mentionPeople: false,
			},
			want: "By <https://broad.io/beehive/r/user/test-1@example.com|test-1>, <https://broad.io/beehive/r/user/test-2@example.com|test-2>, <https://broad.io/beehive/r/user/test-3@gserviceaccount.com|test-3> (service account), <https://broad.io/beehive/r/user/test-4@gserviceaccount.com|test-4> (service account)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slackDeployHookSummarizeUsers(tt.args.users, tt.args.mentionPeople); got != tt.want {
				t.Errorf("slackDeployHookSummarizeUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slackDeployHookChangesetsToChangelogSections(t *testing.T) {
	type args struct {
		changesets    []models.Changeset
		mentionPeople bool
		beehiveUrl    string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test",
			args: args{
				changesets: []models.Changeset{
					{
						ChartRelease: &models.ChartRelease{
							Name: "cr-1",
						},
						From: models.ChartReleaseVersion{
							AppVersionExact: utils.PointerTo("1.0.0"),
						},
						To: models.ChartReleaseVersion{
							AppVersionExact: utils.PointerTo("1.0.1"),
						},
						NewAppVersions: []*models.AppVersion{
							{Description: "a description", AppVersion: "1.0.1"},
						},
					},
					{
						ChartRelease: &models.ChartRelease{
							Name: "cr-2",
						},
					},
				},
				mentionPeople: true,
				beehiveUrl:    "beehive URL",
			},
			want: [][]string{
				{
					"*<https://beehive.dsp-devops-prod.broadinstitute.org/r/chart-release/cr-1|cr-1>* [app 1.0.0⭢1.0.1]",
					"• *app 1.0.1*: a description",
				},
				{
					"*<https://beehive.dsp-devops-prod.broadinstitute.org/r/chart-release/cr-2|cr-2>* [configuration change]",
					"• *No changelog entries found;* <beehive URL|Beehive might have more information>",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slackDeployHookChangesetsToChangelogSections(tt.args.changesets, tt.args.mentionPeople, tt.args.beehiveUrl)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("slackDeployHookChangesetsToChangelogSections(): %s", diff)
			}
		})
	}
}

func Test_slackDeployHookChangelogTitle(t *testing.T) {
	type args struct {
		hasFailure      bool
		destinationLink string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "has failure",
			args: args{
				hasFailure:      true,
				destinationLink: "test",
			},
			want: "Failures deploying to *test*; changelog preview:",
		},
		{
			name: "no failure",
			args: args{
				hasFailure:      false,
				destinationLink: "test",
			},
			want: "Successfully deployed to *test*; changelog preview:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slackDeployHookChangelogTitle(tt.args.hasFailure, tt.args.destinationLink); got != tt.want {
				t.Errorf("slackDeployHookChangelogTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

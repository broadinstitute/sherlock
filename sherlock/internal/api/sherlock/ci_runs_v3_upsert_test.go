package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc/gha_oidc_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks/hooks_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"slices"
	"time"
)

func (s *handlerSuite) TestCiRunsV3Upsert_error() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiRunsV3Upsert_edits() {
	startedAt := time.Now().Add(-time.Minute)
	var got1 CiRunV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				StartedAt:                  &startedAt,
				NotifySlackCustomIcon:      utils.PointerTo(""),
			},
		}),
		&got1)
	s.Equal(http.StatusCreated, code)
	s.WithinDuration(startedAt, *got1.StartedAt, time.Second) // Database stores with less precision
	s.Nil(got1.NotifySlackCustomIcon)
	var got2 CiRunV3
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				Status:                     utils.PointerTo("in_progress"),
				NotifySlackCustomIcon:      utils.PointerTo(":smiley:"),
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal("in_progress", *got2.Status)
	s.Equal(got1.StartedAt, got2.StartedAt)
	s.Equal(got1.ID, got2.ID)
	s.NotEqual(got1.UpdatedAt, got2.UpdatedAt)
	if s.NotNil(got2.NotifySlackCustomIcon) {
		s.Equal(":smiley:", *got2.NotifySlackCustomIcon)
	}
}

func (s *handlerSuite) TestCiRunsV3Upsert_fieldValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
			},
			RelateToChangesetNewVersions: "some invalid value",
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Contains(got.Message, "RelateToChangesetNewVersions")
}

func (s *handlerSuite) TestCiRunsV3Upsert_identifiers() {

	s.Run("basic upsert of identifiers", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				Charts:        []string{s.TestData.Chart_Leonardo().Name},
				ChartVersions: []string{s.TestData.Chart_Leonardo().Name + "/" + s.TestData.ChartVersion_Leonardo_V1().ChartVersion},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 2)
	})

	cluster := s.TestData.Cluster_TerraDev()
	environment := s.TestData.Environment_Dev()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	templateChartRelease := s.TestData.ChartRelease_LeonardoSwatomation()

	s.Run("chart release identifiers", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         123123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				ChartReleases: []string{s.TestData.ChartRelease_LeonardoDev().Name},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 3)
	})

	changesetReq := &ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease:           s.TestData.ChartRelease_LeonardoDev().Name,
					ToAppVersionResolver:   utils.PointerTo("exact"),
					ToAppVersionExact:      utils.PointerTo(s.TestData.AppVersion_Leonardo_V1().AppVersion),
					ToChartVersionResolver: utils.PointerTo("exact"),
					ToChartVersionExact:    utils.PointerTo(s.TestData.ChartVersion_Leonardo_V1().ChartVersion),
				},
			},
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease:           s.TestData.ChartRelease_LeonardoSwatomation().Name,
					ToAppVersionResolver:   utils.PointerTo("exact"),
					ToAppVersionExact:      utils.PointerTo(s.TestData.AppVersion_Leonardo_V1().AppVersion),
					ToChartVersionResolver: utils.PointerTo("exact"),
					ToChartVersionExact:    utils.PointerTo(s.TestData.ChartVersion_Leonardo_V1().ChartVersion),
				},
			},
		},
	}
	changesets, err := changesetReq.parseChartReleaseEntries(s.DB)
	s.NoError(err)
	createdIDs, err := models.PlanChangesets(s.DB, changesets)
	s.NoError(err)
	s.Len(createdIDs, 2)

	var changeset, templateChangeset models.Changeset
	err = s.DB.Scopes(models.ReadChangesetScope).Find(&changesets, createdIDs).Error
	s.NoError(err)
	for _, c := range changesets {
		if c.ChartReleaseID == chartRelease.ID {
			changeset = c
		} else if c.ChartReleaseID == templateChartRelease.ID {
			templateChangeset = c
		}
	}
	s.Len(changeset.NewAppVersions, 1)
	s.Len(changeset.NewChartVersions, 1)
	s.Len(templateChangeset.NewAppVersions, 1)
	s.Len(templateChangeset.NewChartVersions, 1)

	s.Run("more advanced upsert of identifiers", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 2,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(changeset.ID)},
				RelateToChangesetNewVersions: "never",
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 4)
	})
	s.Run("more advanced upsert of identifiers with more spreading", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 2,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets: []string{utils.UintToString(changeset.ID)},
				// Use default for RelateToChangesetNewVersions, which spreads for static environments
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 6)
		s.Run("put loads identifiers", func() {
			var gotAgain CiRunV3
			code = s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                   "github-actions",
						GithubActionsOwner:         "owner",
						GithubActionsRepo:          "repo",
						GithubActionsRunID:         1,
						GithubActionsAttemptNumber: 2,
						GithubActionsWorkflowPath:  "workflow",
					},
				}),
				&gotAgain)
			s.Equal(http.StatusCreated, code)
			s.Len(gotAgain.RelatedResources, 6)
		})
		s.Run("get loads identifiers too", func() {
			var gotAgain CiRunV3
			code = s.HandleRequest(
				s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/%d", got.ID), nil),
				&gotAgain)
			s.Equal(http.StatusOK, code)
			s.Len(gotAgain.RelatedResources, 6)
		})
	})
	s.Run("changeset spreading for non-static only when set to always", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(templateChangeset.ID)},
				RelateToChangesetNewVersions: "never",
			}),
			&got)
		s.Run("never", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 4)
		})
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets: []string{utils.UintToString(templateChangeset.ID)},
			}),
			&got)
		s.Run("default", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 4)
		})
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(templateChangeset.ID)},
				RelateToChangesetNewVersions: "when-static",
			}),
			&got)
		s.Run("explicit when-static", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 4)
		})
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(templateChangeset.ID)},
				RelateToChangesetNewVersions: "always",
			}),
			&got)
		s.Run("always", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 6)
		})
	})
	s.Run("eliminates duplicates and spreads downwards", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 5,
					GithubActionsWorkflowPath:  "workflow",
				},
				Environments: []string{"dev", "dev", utils.UintToString(environment.ID)},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 3)
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 5,
					GithubActionsWorkflowPath:  "workflow",
				},
				Environments: []string{"dev"},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 3)
		var envPresent, chartReleasePresent, clusterPresent bool
		for _, ciIdentifier := range got.RelatedResources {
			if ciIdentifier.ResourceType == "environment" {
				s.Equal(environment.ID, ciIdentifier.ResourceID)
				envPresent = true
			}
			if ciIdentifier.ResourceType == "chart-release" {
				s.Equal(chartRelease.ID, ciIdentifier.ResourceID)
				chartReleasePresent = true
			}
			if ciIdentifier.ResourceType == "cluster" {
				s.Equal(cluster.ID, ciIdentifier.ResourceID)
				clusterPresent = true
			}
		}
		s.True(envPresent)
		s.True(chartReleasePresent)
		s.True(clusterPresent)
	})
}

func (s *handlerSuite) TestCiRunsV3Upsert_identifiersInvalid() {
	s.TestData.Chart_Leonardo()

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
			},
			Charts: []string{"leonardo-that-doesn't-exist"},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(got.Type, errors.BadRequest)
}

func (s *handlerSuite) TestCiRunsV3Upsert_identifiersInvalidIgnore() {
	chart := s.TestData.Chart_Leonardo()

	s.Run("flat ignore", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				Charts:             []string{"leonardo-that-doesn't-exist"},
				IgnoreBadSelectors: true,
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 0)
		s.NotZero(got.ID)
	})
	s.Run("opportunistic when ignoring", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				Charts:             []string{"leonardo-that-doesn't-exist", "leonardo"},
				IgnoreBadSelectors: true,
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 1)
		s.Equal(chart.ID, got.RelatedResources[0].ResourceID)
		s.NotZero(got.ID)
	})
}

func (s *handlerSuite) TestCiRunsV3Upsert_slackNotifications() {
	s.Run("failure", func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-failure-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1234,
						GithubActionsAttemptNumber:     1,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("failure"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
						NotifySlackChannelsUponRetry:   []string{"#my-retry-channel"},
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.Equal([]string{"#my-retry-channel"}, got.NotifySlackChannelsUponRetry)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
	s.Run("success", func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-success-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1233,
						GithubActionsAttemptNumber:     1,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("success"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
						NotifySlackChannelsUponRetry:   []string{"#my-retry-channel"},
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.Equal([]string{"#my-retry-channel"}, got.NotifySlackChannelsUponRetry)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
	s.Run("retry", func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-success-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-retry-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1232,
						GithubActionsAttemptNumber:     2,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("success"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
						NotifySlackChannelsUponRetry:   []string{"#my-retry-channel"},
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.Equal([]string{"#my-retry-channel"}, got.NotifySlackChannelsUponRetry)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
	s.Run("retry dedupe", func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-success-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-retry-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1239,
						GithubActionsAttemptNumber:     2,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("success"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
						NotifySlackChannelsUponRetry:   []string{"#my-retry-channel", "#my-success-channel"},
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.Equal([]string{"#my-retry-channel", "#my-success-channel"}, got.NotifySlackChannelsUponRetry)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
	s.Run("with against string", func() {
		chartRelease := s.TestData.ChartRelease_LeonardoDev()
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				// Unfortunately mockery isn't smart enough for us to actually test the text output here but at least
				// we can validate that this scenario doesn't blow things up.
				SendMessageContext(mock.Anything, "#my-success-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1231,
						GithubActionsAttemptNumber:     1,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("success"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
						NotifySlackChannelsUponRetry:   []string{"#my-retry-channel"},
					},
					ChartReleases: []string{chartRelease.Name},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.Equal([]string{"#my-retry-channel"}, got.NotifySlackChannelsUponRetry)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
}

// TestCiRunsV3Upsert_retryChannelInheritance simulates a complex form of an issue described in [Slack].
// We now inherit retry notification channels from previous runs (where possible) to smooth over this
// UX snag.
//
// [Slack]: https://broadinstitute.slack.com/archives/C029LTN5L80/p1713879397983369?thread_ts=1713453869.706689&cid=C029LTN5L80
func (s *handlerSuite) TestCiRunsV3Upsert_retryChannelInheritance() {
	var got CiRunV3
	var code int

	// First attempt, no retry channels, *yet*
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                       "github-actions",
				GithubActionsOwner:             "owner",
				GithubActionsRepo:              "repo",
				GithubActionsRunID:             1234,
				GithubActionsAttemptNumber:     1,
				GithubActionsWorkflowPath:      "workflow",
				StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
				Status:                         utils.PointerTo("in_progress"),
				NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
				NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
				NotifySlackChannelsUponRetry:   []string{},
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Nil(got.NotifySlackChannelsUponRetry)

	// Suppose a second attempt starts, has a retry channel
	// No retry channels on the first yet, so we don't get any inherited
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                       "github-actions",
				GithubActionsOwner:             "owner",
				GithubActionsRepo:              "repo",
				GithubActionsRunID:             1234,
				GithubActionsAttemptNumber:     2,
				GithubActionsWorkflowPath:      "workflow",
				StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
				Status:                         utils.PointerTo("in_progress"),
				NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
				NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
				NotifySlackChannelsUponRetry:   []string{"#my-second-attempt-retry-channel"},
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal([]string{"#my-second-attempt-retry-channel"}, got.NotifySlackChannelsUponRetry)

	// Suppose the first workflow progresses more and now does add a retry channel
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                       "github-actions",
				GithubActionsOwner:             "owner",
				GithubActionsRepo:              "repo",
				GithubActionsRunID:             1234,
				GithubActionsAttemptNumber:     1,
				GithubActionsWorkflowPath:      "workflow",
				StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
				Status:                         utils.PointerTo("in_progress"),
				NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
				NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
				NotifySlackChannelsUponRetry:   []string{"#my-first-attempt-retry-channel"},
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal([]string{"#my-first-attempt-retry-channel"}, got.NotifySlackChannelsUponRetry)

	// Now suppose we have a third workflow, marked as completed as soon as we hear about it. It should
	// send Slack messages to the retry channels from the first and second attempts in addition to its
	// own.
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().
			SendMessageContext(mock.Anything, "#my-first-attempt-retry-channel", mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", nil)
		c.EXPECT().
			SendMessageContext(mock.Anything, "#my-second-attempt-retry-channel", mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", nil)
		c.EXPECT().
			SendMessageContext(mock.Anything, "#my-third-attempt-retry-channel", mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", nil)
	}, func() {
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                     "github-actions",
					GithubActionsOwner:           "owner",
					GithubActionsRepo:            "repo",
					GithubActionsRunID:           1234,
					GithubActionsAttemptNumber:   3,
					GithubActionsWorkflowPath:    "workflow",
					StartedAt:                    utils.PointerTo(time.Now().Add(-time.Minute)),
					TerminalAt:                   utils.PointerTo(time.Now()),
					Status:                       utils.PointerTo("success"),
					NotifySlackChannelsUponRetry: []string{"#my-third-attempt-retry-channel"},
				},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		slices.Sort(got.NotifySlackChannelsUponRetry)
		s.Equal([]string{"#my-first-attempt-retry-channel", "#my-second-attempt-retry-channel", "#my-third-attempt-retry-channel"}, got.NotifySlackChannelsUponRetry)
	})
}

func (s *handlerSuite) TestCiRunsV3Upsert_githubActionsClaimDefaults() {
	// Note that the request body is empty!
	// Normally this would result in an error due to missing fields, but suppose a GHA OIDC JWT was passed...
	request := s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{})
	request.Header.Set(gha_oidc.Header, "some GHA OIDC token")

	var got CiRunV3
	var code int
	gha_oidc.UseMockedVerifier(s.T(), func(v *gha_oidc_mocks.MockMockableVerifier) {
		v.EXPECT().VerifyAndParseClaims(mock.AnythingOfType("*gin.Context"), "some GHA OIDC token").
			Return(gha_oidc_claims.Claims{
				RepositoryOwner: "broadinstitute",
				Repository:      "broadinstitute/terra-github-workflows",
				WorkflowRef:     "broadinstitute/terra-github-workflows/.github/workflows/bee-create.yaml@refs/heads/main",
				RunID:           "123456",
				RunAttempt:      "1",
			}, nil)
	}, func() {
		code = s.HandleRequest(request, &got)
	})

	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)
	s.Equal(got.Platform, "github-actions")
	s.Equal(got.GithubActionsOwner, "broadinstitute")
	s.Equal(got.GithubActionsRepo, "terra-github-workflows")
	s.Equal(got.GithubActionsRunID, uint(123456))
	s.Equal(got.GithubActionsAttemptNumber, uint(1))
	s.Equal(got.GithubActionsWorkflowPath, ".github/workflows/bee-create.yaml")
}

func (s *handlerSuite) TestCiRunsV3Upsert_dispatch() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()

	var got CiRunV3
	var code int

	// 1. Suppose a CiRun is upserted via webhook
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "broadinstitute",
					GithubActionsRepo:          "terra-github-workflows",
					GithubActionsRunID:         123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
					StartedAt:                  utils.PointerTo(changeset.AppliedAt.Add(30 * time.Second)),
					Status:                     utils.PointerTo("in_progress"),
				},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Status) {
		s.Equal("in_progress", *got.Status)
	}

	// 2. Suppose the relation to the changeset is reported by the action itself, plus a notification channel
	//    (now deploy hooks start firing because we can match the resources)
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
		d.EXPECT().DispatchSlackDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                       "github-actions",
					GithubActionsOwner:             "broadinstitute",
					GithubActionsRepo:              "terra-github-workflows",
					GithubActionsRunID:             123,
					GithubActionsAttemptNumber:     1,
					GithubActionsWorkflowPath:      ".github/workflows/sync-release.yaml",
					NotifySlackChannelsUponSuccess: []string{"#workbench-resilience-dev"},
					NotifySlackChannelsUponFailure: []string{"#workbench-resilience-dev", "#ap-k8s-monitor"},
				},
				Changesets: []string{utils.UintToString(changeset.ID)},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)

	s.Run("check CiIdentifier creation", func() {
		relatedResourceCounts := map[string]int{}
		for _, rr := range got.RelatedResources {
			if s.NotZero(rr.ResourceType) && s.NotZero(rr.ResourceID) && s.NotZero(rr.ID) {
				relatedResourceCounts[rr.ResourceType]++
			}
			s.Nil(rr.ResourceStatus)
		}
		s.Equal(1, relatedResourceCounts["changeset"])
		s.Equal(1, relatedResourceCounts["chart-release"])
		s.Equal(1, relatedResourceCounts["cluster"])
		s.Equal(1, relatedResourceCounts["environment"])
		s.Equal(len(changeset.NewAppVersions), relatedResourceCounts["app-version"])
		s.Equal(len(changeset.NewChartVersions), relatedResourceCounts["chart-version"])
	})

	// 3. Suppose Thelma reported the status for the chart release itself
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
		d.EXPECT().DispatchSlackDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "broadinstitute",
					GithubActionsRepo:          "terra-github-workflows",
					GithubActionsRunID:         123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
				},
				ChartReleaseStatuses: map[string]string{
					chartRelease.Name: "success: healthy",
				},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)

	s.Run("check resource statuses", func() {
		for _, rr := range got.RelatedResources {
			if rr.ResourceType == "chart-release" || rr.ResourceType == "changeset" || rr.ResourceType == "app-version" || rr.ResourceType == "chart-version" {
				if s.NotNil(rr.ResourceStatus) {
					s.Equal("success: healthy", *rr.ResourceStatus)
				}
			} else {
				s.Nil(rr.ResourceStatus)
			}
		}
	})

	// 4. Suppose CiRun marked as completed by via webhook
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
		d.EXPECT().DispatchSlackDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		d.EXPECT().DispatchGithubActionsDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		d.EXPECT().DispatchSlackCompletionNotification(mock.Anything, "#workbench-resilience-dev", mock.Anything, true, mock.AnythingOfType("*string")).Return(nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "broadinstitute",
					GithubActionsRepo:          "terra-github-workflows",
					GithubActionsRunID:         123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
					TerminalAt:                 utils.PointerTo(changeset.AppliedAt.Add(time.Minute)),
					Status:                     utils.PointerTo("success"),
				},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Status) {
		s.Equal("success", *got.Status)
	}
	s.NotNil(got.TerminationHooksDispatchedAt)
}

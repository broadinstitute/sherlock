package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestSlackDeployHooksV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/deploy-hooks/slack/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/deploy-hooks/slack/v3/123", gin.H{
			"slackChannel": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "slackChannel")
}

func (s *handlerSuite) TestSlackDeployHooksV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/deploy-hooks/slack/v3/123", SlackDeployHookV3Edit{
			SlackDeployHookFields: SlackDeployHookFields{
				SlackChannel: utils.PointerTo("some channel"),
			},
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Edit_sqlValidation() {
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), SlackDeployHookV3Edit{
			SlackDeployHookFields: SlackDeployHookFields{
				SlackChannel: utils.PointerTo(""),
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "slack_channel_present")
}

func (s *handlerSuite) TestSlackDeployHooksV3Edit_forbidden() {
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Prod().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), SlackDeployHookV3Edit{}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Edit() {
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	s.Run("simple case", func() {
		var got SlackDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), SlackDeployHookV3Edit{
				deployHookTriggerConfigV3EditableFields: deployHookTriggerConfigV3EditableFields{
					OnSuccess: utils.PointerTo(true),
				},
				SlackDeployHookFields: SlackDeployHookFields{
					SlackChannel:  utils.PointerTo("different channel"),
					MentionPeople: utils.PointerTo(true),
				},
			}),
			&got)
		s.Equal(http.StatusOK, code)
		if s.NotNil(got.SlackChannel) {
			s.Equal("different channel", *got.SlackChannel)
		}
		if s.NotNil(got.MentionPeople) {
			s.True(*got.MentionPeople)
		}
		if s.NotNil(got.OnSuccess) {
			s.True(*got.OnSuccess)
		}
		if s.NotNil(got.OnEnvironment) {
			s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
		}
	})
}

func (s *handlerSuite) TestSlackDeployHooksV3Edit_SpuriousDuplicates() {
	// DDO-3402

	bee := s.TestData.Environment_Swatomation_DevBee()

	// Create a deploy hook on the BEE
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &bee.ID,
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	// Number of chart releases in the BEE
	var chartReleasesInBee []models.ChartRelease
	s.NoError(s.DB.Unscoped().Where(&models.ChartRelease{EnvironmentID: &bee.ID}).Find(&chartReleasesInBee).Error)
	startingChartReleases := len(chartReleasesInBee)
	s.Greater(startingChartReleases, 0)

	// No-op deploy hook edit
	var got SlackDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), SlackDeployHookV3Edit{}),
		&got)
	s.Equal(http.StatusOK, code)

	// No duplicate chart releases
	s.NoError(s.DB.Unscoped().Where(&models.ChartRelease{EnvironmentID: &bee.ID}).Find(&chartReleasesInBee).Error)
	s.Len(chartReleasesInBee, startingChartReleases)
}

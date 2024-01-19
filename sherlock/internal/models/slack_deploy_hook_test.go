package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
	"testing"
	"time"
)

func (s *modelSuite) TestSlackDeployHookEnvironment() {
	environment := s.TestData.Environment_Dev()
	hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"),
		Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestSlackDeployHookChartRelease() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"),
		Trigger: DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestSlackDeployHookNoChannel() {
	environment := s.TestData.Environment_Dev()
	hook := SlackDeployHook{Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "slack_channel_present")
}

func (s *modelSuite) TestSlackDeployHookSuitability() {
	environment := s.TestData.Environment_Prod()
	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.NoError(s.DB.Create(&hook).Error)
	})
	s.Run("when not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.ErrorContains(s.DB.Create(&hook).Error, errors.Forbidden)
	})
}

func (s *modelSuite) TestSlackDeployHookFlow() {
	environment := s.TestData.Environment_Dev()
	hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
	s.NotZero(hook.ID)
	s.NotZero(hook.Trigger.ID)

	s.Run("edits hook", func() {
		s.NoError(s.DB.Model(&hook).Updates(&SlackDeployHook{SlackChannel: utils.PointerTo("different channel")}).Error)
		if s.NotNil(hook.SlackChannel) {
			s.Equal("different channel", *hook.SlackChannel)
		}

		s.Run("read", func() {
			hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
			s.NoError(s.DB.Take(&hookAgain).Error)
			if s.NotNil(hookAgain.SlackChannel) {
				s.Equal("different channel", *hookAgain.SlackChannel)
			}
		})

	})

	s.Run("edits trigger", func() {
		s.NoError(s.DB.Model(&hook.Trigger).Updates(&DeployHookTriggerConfig{OnSuccess: utils.PointerTo(true)}).Error)
		if s.NotNil(hook.Trigger.OnSuccess) {
			s.True(*hook.Trigger.OnSuccess)
		}

		s.Run("read", func() {
			s.Run("empty without association", func() {
				hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Take(&hookAgain).Error)
				s.Nil(hookAgain.Trigger.OnSuccess)
			})
			s.Run("filled with association", func() {
				hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Preload("Trigger").Take(&hookAgain).Error)
				if s.NotNil(hookAgain.Trigger.OnSuccess) {
					s.True(*hookAgain.Trigger.OnSuccess)
				}
			})
			s.Run("filled with preload all", func() {
				hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Preload(clause.Associations).Take(&hookAgain).Error)
				if s.NotNil(hookAgain.Trigger.OnSuccess) {
					s.True(*hookAgain.Trigger.OnSuccess)
				}
			})
		})
	})

	s.Run("delete", func() {
		s.NoError(s.DB.Select("Trigger").Delete(&SlackDeployHook{Model: gorm.Model{ID: hook.ID}}).Error)

		s.Run("hook gone", func() {
			var matchingHooks []SlackDeployHook
			s.NoError(s.DB.Where(&SlackDeployHook{Model: gorm.Model{ID: hook.ID}}).Find(&matchingHooks).Error)
			s.Len(matchingHooks, 0)
		})
		s.Run("trigger gone", func() {
			var matchingTriggers []DeployHookTriggerConfig
			s.NoError(s.DB.Where(&DeployHookTriggerConfig{Model: gorm.Model{ID: hook.Trigger.ID}}).Find(&matchingTriggers).Error)
			s.Len(matchingTriggers, 0)
		})
	})
}

func TestDeduplicateSlackDeployHooks(t *testing.T) {
	now := time.Now()
	type args struct {
		hooks []SlackDeployHook
	}
	tests := []struct {
		name string
		args args
		want []SlackDeployHook
	}{
		{
			name: "empty",
			args: args{hooks: []SlackDeployHook{}},
			want: []SlackDeployHook{},
		},
		{
			name: "one",
			args: args{hooks: []SlackDeployHook{{SlackChannel: utils.PointerTo("channel")}}},
			want: []SlackDeployHook{{SlackChannel: utils.PointerTo("channel")}},
		},
		{
			name: "two different",
			args: args{hooks: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel")},
				{SlackChannel: utils.PointerTo("different channel")},
			}},
			want: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel")},
				{SlackChannel: utils.PointerTo("different channel")},
			},
		},
		{
			name: "two same",
			args: args{hooks: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel")},
				{SlackChannel: utils.PointerTo("channel")},
			}},
			want: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel")},
			},
		},
		{
			name: "prioritizes environments",
			args: args{hooks: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}},
				{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnChartReleaseID: utils.PointerTo[uint](2)}},
				{SlackChannel: utils.PointerTo("channel")},
			}},
			want: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}},
			},
		},
		{
			name: "prioritizes mentioning people",
			args: args{hooks: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Model: gorm.Model{ID: 1}},
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(false), Model: gorm.Model{ID: 2}},
				{SlackChannel: utils.PointerTo("channel"), Model: gorm.Model{ID: 3}},
			}},
			want: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Model: gorm.Model{ID: 1}},
			},
		},
		{
			name: "prioritizes createdAt",
			args: args{hooks: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 3)}},
				{SlackChannel: utils.PointerTo("channel"), Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 2)}},
				{SlackChannel: utils.PointerTo("channel")},
			}},
			want: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 2)}},
			},
		},
		{
			name: "prioritizes environments, then mentioning people, then createdAt",
			args: args{hooks: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Model: gorm.Model{ID: 1}},
				{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}, Model: gorm.Model{ID: 2}},
				{SlackChannel: utils.PointerTo("channel"), Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 2), ID: 3}},
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}, Model: gorm.Model{ID: 2}},
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 2), ID: 3}},
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}, Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 2), ID: 2}},
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}, Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 1), ID: 2}},
			}},
			want: []SlackDeployHook{
				{SlackChannel: utils.PointerTo("channel"), MentionPeople: utils.PointerTo(true), Trigger: DeployHookTriggerConfig{OnEnvironmentID: utils.PointerTo[uint](1)}, Model: gorm.Model{CreatedAt: now.Add(-time.Hour * 1), ID: 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, DeduplicateSlackDeployHooks(tt.args.hooks), "DeduplicateSlackDeployHooks(%v)", tt.args.hooks)
		})
		slices.Reverse(tt.args.hooks)
		t.Run(tt.name+" STABLE", func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, DeduplicateSlackDeployHooks(tt.args.hooks), "DeduplicateSlackDeployHooks(%v)", tt.args.hooks)
		})
	}
}

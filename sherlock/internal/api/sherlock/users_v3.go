package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type UserV3 struct {
	CommonFields
	Email                  string              `json:"email" form:"email"`
	GoogleID               string              `json:"googleID" form:"googleID"`
	GithubUsername         *string             `json:"githubUsername,omitempty" form:"githubUsername"`
	GithubID               *string             `json:"githubID,omitempty" form:"githubID"`
	SlackUsername          *string             `json:"slackUsername,omitempty" form:"slackUsername"`
	SlackID                *string             `json:"slackID,omitempty" form:"slackID"`
	Suitable               *bool               `json:"suitable,omitempty" form:"suitable"`                             // Available only in responses; indicates whether the user is production-suitable
	SuitabilityDescription *string             `json:"suitabilityDescription,omitempty" form:"suitabilityDescription"` // Available only in responses; describes the user's production-suitability
	DeactivatedAt          *time.Time          `json:"deactivatedAt,omitempty" form:"deactivatedAt"`                   // If set, indicates that the user is currently deactivated
	Assignments            []*RoleAssignmentV3 `json:"assignments,omitempty" form:"-"`
	userDirectlyEditableFields
}

type userDirectlyEditableFields struct {
	Name     *string `json:"name,omitempty" form:"name"`
	NameFrom *string `json:"nameFrom,omitempty" form:"nameFrom" enums:"sherlock,github,slack" binding:"omitempty,oneof=sherlock github slack"`
}

func (u UserV3) toModel() models.User {
	ret := models.User{
		Model:          u.toGormModel(),
		Email:          u.Email,
		GoogleID:       u.GoogleID,
		GithubUsername: u.GithubUsername,
		GithubID:       u.GithubID,
		SlackUsername:  u.SlackUsername,
		SlackID:        u.SlackID,
		Name:           u.Name,
		NameFrom:       u.NameFrom,
		DeactivatedAt:  u.DeactivatedAt,
	}
	return ret
}

func userFromModel(model models.User) UserV3 {
	suitable := false
	suitabilityDescription := "no matching suitability record found or loaded; assuming unsuitable"
	if model.Suitability != nil && model.Suitability.Suitable != nil && model.Suitability.Description != nil {
		suitable = *model.Suitability.Suitable
		suitabilityDescription = *model.Suitability.Description
	}
	ret := UserV3{
		CommonFields:           commonFieldsFromGormModel(model.Model),
		Email:                  model.Email,
		GoogleID:               model.GoogleID,
		GithubUsername:         model.GithubUsername,
		GithubID:               model.GithubID,
		SlackUsername:          model.SlackUsername,
		SlackID:                model.SlackID,
		DeactivatedAt:          model.DeactivatedAt,
		Suitable:               &suitable,
		SuitabilityDescription: &suitabilityDescription,
		userDirectlyEditableFields: userDirectlyEditableFields{
			Name:     model.Name,
			NameFrom: model.NameFrom,
		},
	}
	if len(model.Assignments) > 0 {
		ret.Assignments = utils.Map(model.Assignments, func(ra *models.RoleAssignment) *RoleAssignmentV3 {
			return utils.NilOrCall(roleAssignmentFromModel, ra)
		})
	}
	return ret
}

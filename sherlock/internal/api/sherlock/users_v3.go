package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

type UserV3 struct {
	CommonFields
	Email                  string  `json:"email" form:"email"`
	GoogleID               string  `json:"googleID" form:"googleID"`
	GithubUsername         *string `json:"githubUsername,omitempty" form:"githubUsername"`
	GithubID               *string `json:"githubID,omitempty" form:"githubID"`
	SlackUsername          *string `json:"slackUsername,omitempty" form:"slackUsername"`
	SlackID                *string `json:"slackID,omitempty" form:"slackID"`
	Suitable               *bool   `json:"suitable,omitempty" form:"suitable"`                             // Available only in responses; indicates whether the user is production-suitable
	SuitabilityDescription *string `json:"suitabilityDescription,omitempty" form:"suitabilityDescription"` // Available only in responses; describes the user's production-suitability
	userDirectlyEditableFields
}

type userDirectlyEditableFields struct {
	Name *string `json:"name,omitempty" form:"name"`
	// Controls whether Sherlock should automatically update the user's name based on a connected GitHub identity.
	// Will be set to true if the user account has no name and a GitHub account is linked.
	NameInferredFromGithub *bool   `json:"nameInferredFromGithub,omitempty" form:"nameInferredFromGithub"`
	NameFrom               *string `json:"nameFrom,omitempty" form:"nameFrom" enums:"sherlock,github,slack" binding:"omitempty,oneof=sherlock github slack"`
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
	}
	if u.NameFrom == nil && u.NameInferredFromGithub != nil {
		if *u.NameInferredFromGithub {
			githubString := "github"
			ret.NameFrom = &githubString
		} else {
			sherlockString := "sherlock"
			ret.NameFrom = &sherlockString
		}
	}
	return ret
}

func userFromModel(model models.User) UserV3 {
	suitable := model.Suitability().Suitable()
	suitabilityDescription := model.Suitability().Description()
	ret := UserV3{
		CommonFields:           commonFieldsFromGormModel(model.Model),
		Email:                  model.Email,
		GoogleID:               model.GoogleID,
		GithubUsername:         model.GithubUsername,
		GithubID:               model.GithubID,
		SlackUsername:          model.SlackUsername,
		SlackID:                model.SlackID,
		Suitable:               &suitable,
		SuitabilityDescription: &suitabilityDescription,
		userDirectlyEditableFields: userDirectlyEditableFields{
			Name:     model.Name,
			NameFrom: model.NameFrom,
		},
	}
	if model.NameFrom != nil {
		if *model.NameFrom == "github" {
			trueBool := true
			ret.NameInferredFromGithub = &trueBool
		} else {
			falseBool := false
			ret.NameInferredFromGithub = &falseBool
		}
	}
	return ret
}

package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

type UserV3 struct {
	commonFields
	Email                  string  `json:"email" form:"email"`
	GoogleID               string  `json:"googleID" form:"googleID"`
	GithubUsername         *string `json:"githubUsername,omitempty" form:"githubUsername"`
	GithubID               *string `json:"githubID,omitempty" form:"githubID"`
	Suitable               *bool   `json:"suitable,omitempty" form:"suitable"`                             // Available only in responses; indicates whether the user is production-suitable
	SuitabilityDescription *string `json:"suitabilityDescription,omitempty" form:"suitabilityDescription"` // Available only in responses; describes the user's production-suitability
	userDirectlyEditableFields
}

type userDirectlyEditableFields struct {
	Name *string `json:"name,omitempty" form:"name"`
	// Controls whether Sherlock should automatically update the user's name based on a connected GitHub identity.
	// Will be set to true if the user account has no name and a GitHub account is linked.
	NameInferredFromGithub *bool `json:"nameInferredFromGithub,omitempty" form:"nameInferredFromGithub"`
}

func (u UserV3) toModel() models.User {
	return models.User{
		Model:                  u.toGormModel(),
		Email:                  u.Email,
		GoogleID:               u.GoogleID,
		GithubUsername:         u.GithubUsername,
		GithubID:               u.GithubID,
		Name:                   u.Name,
		NameInferredFromGithub: u.NameInferredFromGithub,
	}
}

func userFromModel(model models.User) UserV3 {
	suitable := model.Suitability().Suitable()
	suitabilityDescription := model.Suitability().Description()
	return UserV3{
		commonFields:           commonFieldsFromGormModel(model.Model),
		Email:                  model.Email,
		GoogleID:               model.GoogleID,
		GithubUsername:         model.GithubUsername,
		GithubID:               model.GithubID,
		Suitable:               &suitable,
		SuitabilityDescription: &suitabilityDescription,
		userDirectlyEditableFields: userDirectlyEditableFields{
			Name:                   model.Name,
			NameInferredFromGithub: model.NameInferredFromGithub,
		},
	}
}

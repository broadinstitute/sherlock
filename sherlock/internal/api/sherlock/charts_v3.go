package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

type ChartV3 struct {
	CommonFields
	CiIdentifier *CiIdentifierV3 `json:"ciIdentifier,omitempty" form:"-"`
	ChartV3Create
}

type ChartV3Create struct {
	Name string `json:"name" form:"name"` // Required when creating
	ChartV3Edit
}

type ChartV3Edit struct {
	ChartRepo             *string `json:"chartRepo" form:"chartRepo" default:"terra-helm"`
	AppImageGitRepo       *string `json:"appImageGitRepo" form:"appImageGitRepo"`
	AppImageGitMainBranch *string `json:"appImageGitMainBranch" form:"appImageGitMainBranch"`
	ChartExposesEndpoint  *bool   `json:"chartExposesEndpoint" form:"chartExposesEndpoint" default:"false"` // Indicates if the default subdomain, protocol, and port fields are relevant for this chart
	LegacyConfigsEnabled  *bool   `json:"legacyConfigsEnabled" form:"legacyConfigsEnabled" default:"false"` // Indicates whether a chart requires config rendering from firecloud-develop
	DefaultSubdomain      *string `json:"defaultSubdomain" form:"defaultSubdomain"`                         // When creating, will default to the name of the chart
	DefaultProtocol       *string `json:"defaultProtocol" form:"defaultProtocol" default:"https"`
	DefaultPort           *uint   `json:"defaultPort" form:"defaultPort" default:"443"`
	Description           *string `json:"description" form:"description"`
	PlaybookURL           *string `json:"playbookURL" form:"playbookURL"`
	PactParticipant       *bool   `json:"pactParticipant" form:"pactParticipant" default:"false"`
}

func (c ChartV3) toModel() models.Chart {
	return models.Chart{
		Model:                 c.toGormModel(),
		Name:                  c.Name,
		ChartRepo:             c.ChartRepo,
		AppImageGitRepo:       c.AppImageGitRepo,
		AppImageGitMainBranch: c.AppImageGitMainBranch,
		ChartExposesEndpoint:  c.ChartExposesEndpoint,
		LegacyConfigsEnabled:  c.LegacyConfigsEnabled,
		DefaultSubdomain:      c.DefaultSubdomain,
		DefaultProtocol:       c.DefaultProtocol,
		DefaultPort:           c.DefaultPort,
		Description:           c.Description,
		PlaybookURL:           c.PlaybookURL,
		PactParticipant:       c.PactParticipant,
	}
}

func (c ChartV3Create) toModel() models.Chart {
	return ChartV3{ChartV3Create: c}.toModel()
}

func (c ChartV3Edit) toModel() models.Chart {
	return ChartV3Create{ChartV3Edit: c}.toModel()
}

func chartFromModel(model models.Chart) ChartV3 {
	var ciIdentifier *CiIdentifierV3
	if model.CiIdentifier != nil {
		ciIdentifier = utils.PointerTo(ciIdentifierFromModel(*model.CiIdentifier))
	}
	return ChartV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		CiIdentifier: ciIdentifier,
		ChartV3Create: ChartV3Create{
			Name: model.Name,
			ChartV3Edit: ChartV3Edit{
				ChartRepo:             model.ChartRepo,
				AppImageGitRepo:       model.AppImageGitRepo,
				AppImageGitMainBranch: model.AppImageGitMainBranch,
				ChartExposesEndpoint:  model.ChartExposesEndpoint,
				LegacyConfigsEnabled:  model.LegacyConfigsEnabled,
				DefaultSubdomain:      model.DefaultSubdomain,
				DefaultProtocol:       model.DefaultProtocol,
				DefaultPort:           model.DefaultPort,
				Description:           model.Description,
				PlaybookURL:           model.PlaybookURL,
				PactParticipant:       model.PactParticipant,
			},
		},
	}
}

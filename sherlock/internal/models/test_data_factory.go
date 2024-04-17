// Test Data Factories
// While test_data will populate the the database and return models from database, factories returns only the models so they can be sent to code to... be created in the db.
package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
)

func (td *testDataImpl) Environment_Swatomation_TestBee_Factory() Environment {
	td.environment_swatomation_testBee = Environment{
		Base:                      "bee",
		Lifecycle:                 "dynamic",
		Name:                      "swatomation-test-bee",
		ValuesName:                "swatomation",
		TemplateEnvironmentID:     utils.PointerTo(td.Environment_Swatomation().ID),
		AutoPopulateChartReleases: utils.PointerTo(true),
		DefaultNamespace:          "terra-swatomation-test-bee",
		DefaultClusterID:          utils.PointerTo(td.Cluster_TerraQaBees().ID),
		RequiresSuitability:       utils.PointerTo(false),
		BaseDomain:                utils.PointerTo("bee.envs-terra.bio"),
		NamePrefixesDomain:        utils.PointerTo(true),
		HelmfileRef:               utils.PointerTo("HEAD"),
		PreventDeletion:           utils.PointerTo(false),
		DeleteAfter:               sql.NullTime{Time: time.Now().Add(6 * time.Hour), Valid: true},
		Offline:                   utils.PointerTo(false),
	}
	return td.environment_swatomation_testBee
}

// A TestBee with as little configs as possible
func (td *testDataImpl) Environment_Swatomation_TestBee_Factory_Min() Environment {
	td.environment_swatomation_testBee = Environment{
		Lifecycle:                 "dynamic",
		TemplateEnvironmentID:     utils.PointerTo(td.Environment_Swatomation().ID),
		AutoPopulateChartReleases: utils.PointerTo(true),
		RequiresSuitability:       utils.PointerTo(false),
		BaseDomain:                utils.PointerTo("bee.envs-terra.bio"),
		NamePrefixesDomain:        utils.PointerTo(true),
		HelmfileRef:               utils.PointerTo("HEAD"),
		PreventDeletion:           utils.PointerTo(false),
		DeleteAfter:               sql.NullTime{Time: time.Now().Add(6 * time.Hour), Valid: true},
		Offline:                   utils.PointerTo(false),
	}
	return td.environment_swatomation_testBee
}

// updates leonardo from app version to V3 to V1. Relies on the existence of 2x chart version, 2x app version, 1x Chart Reelease, 1x Environment
// AppVersion_Leonardo_V1, ChartVersionLeonardo_V1
// AppVersion_Leonardo_V3, ChartVersionLeonardo_V3
// ChartRelease_LeonardoSwatomation
func (td *testDataImpl) Changeset_LeonardoSwatomation_TestBee_V3toV1_factory(chartReleaseID uint) Changeset {
	changeset_leonardoSwatomation_v1toV3_factory := Changeset{
		ChartReleaseID: chartReleaseID,
		To: ChartReleaseVersion{
			ResolvedAt:           utils.PointerTo(time.Now().Add(-(24 * time.Hour))),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo(td.AppVersion_Leonardo_V1().AppVersion),
			AppVersionBranch:     utils.PointerTo(td.AppVersion_Leonardo_V1().GitBranch),
			AppVersionCommit:     utils.PointerTo(td.AppVersion_Leonardo_V1().GitCommit),
			AppVersionID:         utils.PointerTo(td.AppVersion_Leonardo_V1().ID),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo(td.ChartVersion_Leonardo_V1().ChartVersion),
			ChartVersionID:       utils.PointerTo(td.ChartVersion_Leonardo_V1().ID),
			HelmfileRef:          utils.PointerTo(fmt.Sprintf("charts/leonardo-%s", td.ChartVersion_Leonardo_V1().ChartVersion)),
			HelmfileRefEnabled:   utils.PointerTo(false),
		},
		From: ChartReleaseVersion{
			ResolvedAt:           utils.PointerTo(time.Now().Add(-(19 * time.Hour))),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo(td.AppVersion_Leonardo_V3().AppVersion),
			AppVersionBranch:     utils.PointerTo(td.AppVersion_Leonardo_V3().GitBranch),
			AppVersionCommit:     utils.PointerTo(td.AppVersion_Leonardo_V3().GitCommit),
			AppVersionID:         utils.PointerTo(td.AppVersion_Leonardo_V3().ID),
			ChartVersionResolver: utils.PointerTo("latest"),
			ChartVersionExact:    utils.PointerTo(td.ChartVersion_Leonardo_V3().ChartVersion),
			ChartVersionID:       utils.PointerTo(td.ChartVersion_Leonardo_V3().ID),
			HelmfileRef:          utils.PointerTo(fmt.Sprintf("charts/leonardo-%s", td.ChartVersion_Leonardo_V3().ChartVersion)),
			HelmfileRefEnabled:   utils.PointerTo(false),
		},
	}
	return changeset_leonardoSwatomation_v1toV3_factory
}

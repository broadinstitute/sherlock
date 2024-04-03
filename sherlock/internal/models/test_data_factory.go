// Test Data Factories
// While test_data will populate the the database and return models from database, factories returns only the models so they can be sent to code to... be created in the db.
package models

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
)

// updates leonardo from app version to V1 to V3. Relies on the existence of 2x chart version, 2x app version, 1x Chart Reelease, 1x Environment
// AppVersion_Leonardo_V1, ChartVersionLeonardo_V1
// AppVersion_Leonardo_V3, ChartVersionLeonardo_V3
// ChartRelease_LeonardoSwatomation
func (td *testDataImpl) Changeset_LeonardoSwatomation_V1toV3_factory() Changeset {
	changeset_leonardoSwatomation_v1toV3_factory := Changeset{
		ChartReleaseID: td.ChartRelease_LeonardoSwatomation().ID,
		From: ChartReleaseVersion{
			ResolvedAt:           utils.PointerTo(time.Now().Add(-(24 * time.Hour))),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo(td.AppVersion_Leonardo_V1().AppVersion),
			AppVersionBranch:     utils.PointerTo(td.AppVersion_Leonardo_V1().GitBranch),
			AppVersionCommit:     utils.PointerTo(td.AppVersion_Leonardo_V1().GitCommit),
			AppVersionID:         utils.PointerTo(td.AppVersion_Leonardo_V1().ID),
			ChartVersionResolver: utils.PointerTo("latest"),
			ChartVersionExact:    utils.PointerTo(td.ChartVersion_Leonardo_V1().ChartVersion),
			ChartVersionID:       utils.PointerTo(td.ChartVersion_Leonardo_V1().ID),
			HelmfileRef:          utils.PointerTo(fmt.Sprintf("charts/leonardo-%s", td.ChartVersion_Leonardo_V1().ChartVersion)),
			HelmfileRefEnabled:   utils.PointerTo(false),
		},
		To: ChartReleaseVersion{
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
		AppliedAt:    utils.PointerTo(time.Now().Add(-(18 * time.Hour))),
		SupersededAt: nil,
		PlannedByID:  utils.PointerTo(td.User_Suitable().ID),
		AppliedByID:  utils.PointerTo(td.User_Suitable().ID),
	}
	return changeset_leonardoSwatomation_v1toV3_factory
}

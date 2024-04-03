package bee

import (
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const beeDefaultTemplate = "swatomation" // TODO: configurize this

func BeeUpsert(environmentCreateBody models.Environment, beeEdits []models.Changeset, db *gorm.DB) (beeModel models.Environment, err error) {
	// get a new bee
	beeModel, err = getBee(environmentCreateBody, db)

	// exit early if error retrieving bee
	if err != nil {
		return
	}

	// update the bee
	err = updateBee(beeEdits, db)

	// done
	return
}

// get an existing bee by name
func getEnvByName(envName string, db *gorm.DB) (envModel models.Environment, err error) {
	err = db.Preload(clause.Associations).First(&envModel, "environments.name = ?", envName).Error
	return
}

// update settings to a bee.
func updateBee(beeChangesets []models.Changeset, db *gorm.DB) (err error) {
	var createdChangesetIDs []uint

	// exit if plan fails
	if createdChangesetIDs, err = models.PlanChangesets(db, beeChangesets); err != nil {
		err = fmt.Errorf("error planning changesets: %w", err)
		return
	}

	// exit if nothing to update
	if len(createdChangesetIDs) == 0 {
		return
	}

	// Do the Apply
	if err = models.ApplyChangesets(db, createdChangesetIDs); err != nil {
		err = fmt.Errorf("error applying changesets: %w", err)
	}
	return
}

// returns either a bee of quality
// 1. existing
// 2. pooled
// 3. newly created
// 4. (nothing) error
func getBee(environmentCreateBody models.Environment, db *gorm.DB) (beeModel models.Environment, err error) {
	var noBee models.Environment // just an empty model

	// get existing bee if name given, else get fresh/pooled bee
	if environmentCreateBody.Name != "" {
		beeModel, err = getEnvByName(environmentCreateBody.Name, db)

		// check if returned bee (no err) matches expected template
		if err == nil && beeModel.TemplateEnvironment.Name != environmentCreateBody.TemplateEnvironment.Name {
			err = fmt.Errorf("(%s) request validation error: Template Mismatch", errors.BadRequest)
		}
	} else {
		// get pooled bee here
		beeModel, err = getPooledBee(environmentCreateBody.TemplateEnvironment.Name, db)

		// if no pooled bee, make a new bee
		if beeModel == noBee {
			err = db.Create(&environmentCreateBody).Error
		}
	}
	return
}

// placeholder method, will return pooled Bees one day.
func getPooledBee(templateName string, db *gorm.DB) (envModel models.Environment, err error) {
	//err = db.Preload(clause.Associations).Where().First(&envModel, templateName).Error
	err = fmt.Errorf("bee pool error: feature not implemented yet")
	return
}

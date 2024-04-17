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
	// validate struct
	if environmentCreateBody.Lifecycle != "dynamic" {
		err = fmt.Errorf("(%s) request validation error: Lifecycle is not \"dynamic\"", errors.BadRequest)
		return
	}

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

		// if has no err (bee successfully got)
		// else get a new-ish bee somehow
		if err == nil {
			// throw an error if bee has an expected template AND
			// expected template does not match retreived bee's template
			if environmentCreateBody.TemplateEnvironment != nil &&
				beeModel.TemplateEnvironment.Name != environmentCreateBody.TemplateEnvironment.Name {
				err = fmt.Errorf("(%s) request validation error: Template Mismatch", errors.BadRequest)
				return
			}
		} else {
			err = nil
			err = trySetDefaultTemplate(&environmentCreateBody, db)
			if err != nil {
				return
			}

			err = db.Omit(clause.Associations).Create(&environmentCreateBody).Error
			if err != nil {
				return
			}
			//beeModel, err = getEnvByName(environmentCreateBody.Name, db)
			err = db.Preload(clause.Associations).First(&beeModel, environmentCreateBody).Error
		}
	} else {
		// no logic here.
		err = trySetDefaultTemplate(&environmentCreateBody, db)
		if err != nil {
			return
		}

		beeModel, err = getPooledBee(environmentCreateBody.TemplateEnvironmentID, db)

		if beeModel == noBee {
			err = db.Omit(clause.Associations).Create(&environmentCreateBody).Error
			if err != nil {
				return
			}

			err = db.Preload(clause.Associations).First(&beeModel, environmentCreateBody).Error
		}
	}

	return
}

// update settings to a bee.
func updateBee(beeChangesets []models.Changeset, db *gorm.DB) (err error) {
	var createdChangesetIDs []uint

	if len(beeChangesets) == 0 {
		return
	}

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

// placeholder method, will return pooled Bees one day.
func getPooledBee(templateID *uint, db *gorm.DB) (envModel models.Environment, err error) {
	//err = db.Preload(clause.Associations).Where().First(&envModel, templateName).Error
	err = fmt.Errorf("bee pool error: feature not implemented yet")
	return
}

// try to set a default template, does nothing if template exists
func trySetDefaultTemplate(environmentCreateBody *models.Environment, db *gorm.DB) (err error) {
	if environmentCreateBody.TemplateEnvironmentID == nil {
		var templateEnvironment models.Environment
		templateEnvironment, err = getEnvByName(beeDefaultTemplate, db)
		environmentCreateBody.TemplateEnvironmentID = &templateEnvironment.ID
	}
	return
}

// get an existing bee by name
func getEnvByName(envName string, db *gorm.DB) (envModel models.Environment, err error) {
	err = db.Preload(clause.Associations).First(&envModel, "environments.name = ?", envName).Error
	return
}

// UNUSED allows unused variables to be included in Go programs
func UNUSED(x ...interface{}) {}

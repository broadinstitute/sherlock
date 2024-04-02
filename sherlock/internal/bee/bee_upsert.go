package bee

import (
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const beeDefaultTemplate = "swatomation"

func beeUpsert(environmentCreateBody models.Environment, beeEdits models.Environment, db *gorm.DB, ctx *gin.Context) (beeModel models.Environment, err error) {
	var noBee models.Environment // just an empty model

	// get bee if exists
	if environmentCreateBody.Name != "" {
		beeModel, err = getEnvByName(environmentCreateBody.Name, db)

		// check if returned bee (no err) matches expected template
		if err == nil && beeModel.TemplateEnvironment.Name != environmentCreateBody.TemplateEnvironment.Name {
			err = fmt.Errorf("(%s) request validation error: Template Mismatch", errors.BadRequest)
		}
	} else {
		// if bee does not exist, make it

		// get pooled bee here
		beeModel, err = getPooledBee(environmentCreateBody.TemplateEnvironment.Name, db)

		// if no pooled bee
		if beeModel == noBee {
			err = db.Create(&environmentCreateBody).Error
		}
	}

	// exit early if error retreiving bee
	if err != nil {
		return
	}

	// do things w/ the bee
	err = updateBee(&beeModel, beeEdits, db)

	// done
	return
}

// get an existing bee by name
func getEnvByName(envName string, db *gorm.DB) (envModel models.Environment, err error) {
	err = db.Preload(clause.Associations).First(&envModel, envName).Error
	return
}

// update settings to a bee.
func updateBee(beeModel *models.Environment, beeEdits models.Environment, db *gorm.DB) (err error) {
	err = db.Model(&beeModel).Omit(clause.Associations).Updates(&beeEdits).Error
	return
}

// somehow return a new-unused bee (maybe create it, maybe get from pool)
func newBee(environmentCreateBody models.Environment, db *gorm.DB) (beeModel models.Environment, err error) {
	var noBee models.Environment // just an empty model

	// get bee if exists
	if environmentCreateBody.Name != "" {
		beeModel, err = getEnvByName(environmentCreateBody.Name, db)

		// check if returned bee (no err) matches expected template
		if err == nil && beeModel.TemplateEnvironment.Name != environmentCreateBody.TemplateEnvironment.Name {
			err = fmt.Errorf("(%s) request validation error: Template Mismatch", errors.BadRequest)
		}
	} else {
		// if bee does not exist, make it

		// get pooled bee here
		beeModel, err = getPooledBee(environmentCreateBody.TemplateEnvironment.Name, db)

		// if no pooled bee
		if beeModel == noBee {
			err = db.Create(&environmentCreateBody).Error
		}
	}
}

// placeholder method, will return pooled Bees one day.
func getPooledBee(templateName string, db *gorm.DB) (envModel models.Environment, err error) {
	//err = db.Preload(clause.Associations).Where().First(&envModel, templateName).Error
	err = fmt.Errorf("bee pool error: feature not implemented yet")
	return
}

package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/knadh/koanf"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	explicitJoinTableRelationships := []struct {
		model             any
		relationFieldName string
		joinTableModel    any
	}{
		{&CiRun{}, "RelatedResources", &CiRunIdentifierJoin{}},
		{&CiIdentifier{}, "CiRuns", &CiRunIdentifierJoin{}},
	}
	for _, r := range explicitJoinTableRelationships {
		if err := db.SetupJoinTable(r.model, r.relationFieldName, r.joinTableModel); err != nil {
			return fmt.Errorf("failed to configure model %T relation %s to use join model %T: %v", r.model, r.relationFieldName, r.joinTableModel, err)
		}
	}

	if err := initDeployMatchers(); err != nil {
		return err
	}

	// If you're working on a new model and want to have
	// Gorm basically fudge the database schema from your
	// struct, here's a good place to do that.
	//
	//err = db.AutoMigrate(&SomeNewModel{})
	//if err != nil {
	//	return err
	//}

	return nil
}

func initDeployMatchers() error {
	deployMatchers = []CiRun{}
	for index, k := range config.Config.Slices("model.ciRuns.deployMatchers") {
		var partial CiRun
		if err := k.UnmarshalWithConf("", &partial, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
			return fmt.Errorf("error parsing model.ciRuns.deployMatchers[%d]: %w", index+1, err)
		} else {
			deployMatchers = append(deployMatchers, partial)
		}
	}
	return nil
}

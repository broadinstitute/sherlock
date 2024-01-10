package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *modelSuite) TestPagerdutyIntegrationPagerdutyIDValidationSql() {
	s.SetSuitableTestUserForDB()
	err := s.DB.Create(&PagerdutyIntegration{
		Name: utils.PointerTo("some-name"),
		Key:  utils.PointerTo("some-key"),
		Type: utils.PointerTo("some-type"),
	}).Error
	s.ErrorContains(err, "pagerduty_id_present")
}

func (s *modelSuite) TestPagerdutyIntegrationNameValidationSql() {
	s.SetSuitableTestUserForDB()
	err := s.DB.Create(&PagerdutyIntegration{
		PagerdutyID: "some-pagerduty-id",
		Key:         utils.PointerTo("some-key"),
		Type:        utils.PointerTo("some-type"),
	}).Error
	s.ErrorContains(err, "name_present")
}

func (s *modelSuite) TestPagerdutyIntegrationKeyValidationSql() {
	s.SetSuitableTestUserForDB()
	err := s.DB.Create(&PagerdutyIntegration{
		PagerdutyID: "some-pagerduty-id",
		Name:        utils.PointerTo("some-name"),
		Type:        utils.PointerTo("some-type"),
	}).Error
	s.ErrorContains(err, "key_present")
}

func (s *modelSuite) TestPagerdutyIntegrationTypeValidationSql() {
	s.SetSuitableTestUserForDB()
	err := s.DB.Create(&PagerdutyIntegration{
		PagerdutyID: "some-pagerduty-id",
		Name:        utils.PointerTo("some-name"),
		Key:         utils.PointerTo("some-key"),
	}).Error
	s.ErrorContains(err, "type_present")
}

func (s *modelSuite) TestPagerdutyIntegrationCreationForbidden() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&PagerdutyIntegration{
		PagerdutyID: "some-pagerduty-id",
		Name:        utils.PointerTo("some-name"),
		Key:         utils.PointerTo("some-key"),
		Type:        utils.PointerTo("some-type"),
	}).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestPagerdutyIntegrationEditForbidden() {
	i := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(s.DB.
		Model(&i).
		Updates(&PagerdutyIntegration{Key: utils.PointerTo("some fancy key")}).
		Error, errors.Forbidden)
}

func (s *modelSuite) TestPagerdutyIntegrationDeleteForbidden() {
	i := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	s.SetNonSuitableTestUserForDB()
	s.ErrorContains(s.DB.
		Delete(&i).
		Error, errors.Forbidden)
}

func (s *modelSuite) TestPagerdutyIntegrationDeleteWhileUsed() {
	i := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	s.TestData.Environment_Prod()
	s.SetSuitableTestUserForDB()
	s.ErrorContains(s.DB.
		Delete(&i).
		Error, errors.BadRequest)
}

func (s *modelSuite) TestPagerdutyIntegrationUniqueness() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	s.SetSuitableTestUserForDB()
	s.ErrorContains(s.DB.Create(&PagerdutyIntegration{
		PagerdutyID: pdi.PagerdutyID,
		Name:        utils.PointerTo("some-name"),
		Key:         utils.PointerTo("some-key"),
		Type:        utils.PointerTo("some-type"),
	}).Error, "pagerduty_integrations_pagerduty_id_unique_constraint")
}

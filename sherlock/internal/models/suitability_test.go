package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *modelSuite) TestSuitabilityForbiddenCreate() {
	s.SetSuitableTestUserForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestSuitabilityAllowedCreate() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestSuitabilityForbiddenEdit() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.NoError(err)
	s.SetSuitableTestUserForDB()
	err = s.DB.Where(&Suitability{
		Email: utils.PointerTo("email@example.com"),
	}).Updates(&Suitability{
		Suitable: utils.PointerTo(false),
	}).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestSuitabilityAllowedEdit() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.NoError(err)
	err = s.DB.Where(&Suitability{
		Email: utils.PointerTo("email@example.com"),
	}).Updates(&Suitability{
		Suitable: utils.PointerTo(false),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestSuitabilityForbiddenDelete() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.NoError(err)
	s.SetSuitableTestUserForDB()
	err = s.DB.Where(&Suitability{
		Email: utils.PointerTo("email@example.com"),
	}).Delete(&Suitability{}).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestSuitabilityAllowedDelete() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.NoError(err)
	err = s.DB.Where(&Suitability{
		Email: utils.PointerTo("email@example.com"),
	}).Delete(&Suitability{}).Error
	s.NoError(err)
}

func (s *modelSuite) TestSuitabilityCreateNullEmail() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Suitable:    utils.PointerTo(true),
		Description: utils.PointerTo("description"),
	}).Error
	s.ErrorContains(err, "email")
}

func (s *modelSuite) TestSuitabilityCreateNullSuitable() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:       utils.PointerTo("email@example.com"),
		Description: utils.PointerTo("description"),
	}).Error
	s.ErrorContains(err, "suitable")
}

func (s *modelSuite) TestSuitabilityCreateNullDescription() {
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&Suitability{
		Email:    utils.PointerTo("email@example.com"),
		Suitable: utils.PointerTo(true),
	}).Error
	s.ErrorContains(err, "description")
}

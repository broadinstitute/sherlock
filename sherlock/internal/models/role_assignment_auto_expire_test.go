package models

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"gorm.io/gorm"
	"time"
)

func (s *modelSuite) Test_doAutoExpiration() {
	s.TestData.User_Suitable()
	s.TestData.User_NonSuitable()

	ra := s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()

	s.SetSelfSuperAdminForDB()
	s.NoError(s.DB.Model(&ra).Updates(&RoleAssignment{
		RoleAssignmentFields: RoleAssignmentFields{
			ExpiresAt: utils.PointerTo(time.Now().Add(-time.Hour)),
		},
	}).Error)

	var ras []RoleAssignment
	s.NoError(s.DB.Where(&RoleAssignment{}).Find(&ras).Error)
	existingCount := len(ras)

	propagationFnCalled := false
	propagationFn := func(_ context.Context, _ *gorm.DB, roleID uint) {
		if roleID == s.TestData.Role_TerraSuitableEngineer().ID {
			propagationFnCalled = true
		}
	}

	doAutoExpiration(context.Background(), s.DB, propagationFn)

	s.Run("check that the expired role assignment was deleted", func() {
		var shouldStayEmpty []RoleAssignment
		s.NoError(s.DB.Where(&RoleAssignment{UserID: s.TestData.User_Suitable().ID, RoleID: s.TestData.Role_TerraSuitableEngineer().ID}).Find(&shouldStayEmpty).Error)
		s.Empty(shouldStayEmpty)
	})

	s.Run("check that nothing else got deleted", func() {
		s.NoError(s.DB.Where(&RoleAssignment{}).Find(&ras).Error)
		s.Len(ras, existingCount-1)
	})

	s.Run("check that the propagation function was called", func() {
		s.True(propagationFnCalled)
	})
}

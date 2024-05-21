package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
	"time"
)

func (s *handlerSuite) TestRoleAssignmentsV3List_none() {
	var got []RoleAssignmentV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/role-assignments/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 4) // 4 role assignments relied upon by base-level test users
}

func (s *handlerSuite) TestRoleAssignmentsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/role-assignments/v3?suspended=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/role-assignments/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/role-assignments/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3List() {
	s.TestData.RoleAssignment_SuperAdmin_SherlockSuperAdmin()
	s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	s.TestData.RoleAssignment_Suitable_TerraEngineer()
	s.TestData.RoleAssignment_NonSuitable_TerraEngineer()

	s.Run("all", func() {
		var got []RoleAssignmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/role-assignments/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 4)
	})
	s.Run("none", func() {
		var got []RoleAssignmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/role-assignments/v3?expiresAt="+time.Now().Format(time.RFC3339), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		s.SetSelfSuperAdminForDB()
		s.NoError(s.DB.Model(utils.PointerTo(s.TestData.RoleAssignment_Suitable_TerraEngineer())).Update("suspended", true).Error)
		var got []RoleAssignmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/role-assignments/v3?suspended=true", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []RoleAssignmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/role-assignments/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []RoleAssignmentV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/role-assignments/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.True(got1[0].RoleInfo.ID != got2[0].RoleInfo.ID || got1[0].UserInfo.ID != got2[0].UserInfo.ID)
	})
}

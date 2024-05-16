package sherlock

func (s *handlerSuite) Test_roleAssignmentFromModel() {
	model := s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	expected := RoleAssignmentV3{
		RoleAssignmentV3Edit: RoleAssignmentV3Edit{
			Suspended: model.Suspended,
			ExpiresAt: model.ExpiresAt,
		},
	}
	roleAssignment := roleAssignmentFromModel(model)
	s.Equalf(expected, roleAssignment, "roleAssignmentFromModel()")
}

package models

type CiRunIdentifierJoin struct {
	CiRunID        uint `gorm:"primaryKey"`
	CiIdentifierID uint `gorm:"primaryKey"`
	ResourceStatus *string
}

func (c *CiRunIdentifierJoin) TableName() string {
	return "ci_runs_for_identifiers"
}

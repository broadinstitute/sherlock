package models

import (
	"gorm.io/gorm"
)

type GitCommit struct {
	gorm.Model
	GitRepo      string
	GitCommit    string
	GitBranch    string
	IsMainBranch bool
	SecSincePrev *uint
}

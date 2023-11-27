package models

// SlackDeployHookState holds a SlackDeployHook's state for a specific CiRun,
// helping it keep updating the same message throughout the CiRun's lifecycle.
// This model is purely internal.
type SlackDeployHookState struct {
	CiRun             CiRun           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CiRunID           uint            `gorm:"primaryKey"`
	SlackDeployHook   SlackDeployHook `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SlackDeployHookID uint            `gorm:"primaryKey"`
	MessageTimestamp  string
	MessageChannel    string
}

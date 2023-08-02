package models

import "gorm.io/gorm"

type Cluster struct {
	gorm.Model
	CiIdentifier      *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:cluster"`
	Name              string
	Provider          string
	GoogleProject     string
	AzureSubscription string
	Location          string
	// Mutable
	Base                *string
	Address             *string
	RequiresSuitability *bool
	HelmfileRef         *string
}

func (c Cluster) TableName() string {
	return "v2_clusters"
}

func (c Cluster) GetCiIdentifier() *CiIdentifier {
	if c.CiIdentifier != nil {
		return c.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "cluster", ResourceID: c.ID}
	}
}

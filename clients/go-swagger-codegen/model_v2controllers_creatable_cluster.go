/*
 * Sherlock
 *
 * The Data Science Platform's source-of-truth service
 *
 * API version: development
 * Contact: dsp-devops@broadinstitute.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sherlock

// The subset of Cluster fields that can be set upon creation
type V2controllersCreatableCluster struct {
	// Required when creating
	Address string `json:"address,omitempty"`
	// Required when creating if providers is 'azure'
	AzureSubscription string `json:"azureSubscription,omitempty"`
	// Required when creating
	Base string `json:"base,omitempty"`
	// Required when creating if provider is 'google'
	GoogleProject string `json:"googleProject,omitempty"`
	// Required when creating
	Name string `json:"name"`
	Provider string `json:"provider,omitempty"`
	RequiresSuitability bool `json:"requiresSuitability,omitempty"`
}

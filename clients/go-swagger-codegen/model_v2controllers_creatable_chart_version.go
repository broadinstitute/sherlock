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

type V2controllersCreatableChartVersion struct {
	// Required when creating
	Chart string `json:"chart,omitempty"`
	// Required when creating
	ChartVersion string `json:"chartVersion,omitempty"`
}

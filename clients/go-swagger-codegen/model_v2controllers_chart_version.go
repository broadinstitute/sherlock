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

type V2controllersChartVersion struct {
	// Required when creating
	Chart string `json:"chart,omitempty"`
	ChartInfo *V2controllersChart `json:"chartInfo,omitempty"`
	// Required when creating
	ChartVersion string `json:"chartVersion,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Id int32 `json:"id,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersChangesetPlanRequestChartReleaseEntry v2controllers changeset plan request chart release entry
//
// swagger:model v2controllers.ChangesetPlanRequestChartReleaseEntry
type V2controllersChangesetPlanRequestChartReleaseEntry struct {

	// chart release
	ChartRelease string `json:"chartRelease,omitempty"`

	// to app version branch
	ToAppVersionBranch string `json:"toAppVersionBranch,omitempty"`

	// to app version commit
	ToAppVersionCommit string `json:"toAppVersionCommit,omitempty"`

	// to app version exact
	ToAppVersionExact string `json:"toAppVersionExact,omitempty"`

	// to app version follow chart release
	ToAppVersionFollowChartRelease string `json:"toAppVersionFollowChartRelease,omitempty"`

	// to app version resolver
	ToAppVersionResolver string `json:"toAppVersionResolver,omitempty"`

	// to chart version exact
	ToChartVersionExact string `json:"toChartVersionExact,omitempty"`

	// to chart version follow chart release
	ToChartVersionFollowChartRelease string `json:"toChartVersionFollowChartRelease,omitempty"`

	// to chart version resolver
	ToChartVersionResolver string `json:"toChartVersionResolver,omitempty"`

	// to firecloud develop ref
	ToFirecloudDevelopRef string `json:"toFirecloudDevelopRef,omitempty"`

	// to helmfile ref
	ToHelmfileRef string `json:"toHelmfileRef,omitempty"`

	// to helmfile ref enabled
	ToHelmfileRefEnabled bool `json:"toHelmfileRefEnabled,omitempty"`

	// use exact versions from other chart release
	UseExactVersionsFromOtherChartRelease string `json:"useExactVersionsFromOtherChartRelease,omitempty"`

	// If this is set, also copy the fc-dev ref from an OtherChartRelease
	UseOthersFirecloudDevelopRef bool `json:"useOthersFirecloudDevelopRef,omitempty"`

	// If this is set, also copy the helmfile ref from an OtherChartRelease
	UseOthersHelmfileRef bool `json:"useOthersHelmfileRef,omitempty"`
}

// Validate validates this v2controllers changeset plan request chart release entry
func (m *V2controllersChangesetPlanRequestChartReleaseEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2controllers changeset plan request chart release entry based on context it is used
func (m *V2controllersChangesetPlanRequestChartReleaseEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersChangesetPlanRequestChartReleaseEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersChangesetPlanRequestChartReleaseEntry) UnmarshalBinary(b []byte) error {
	var res V2controllersChangesetPlanRequestChartReleaseEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

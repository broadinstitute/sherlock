// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersCreatableChangeset v2controllers creatable changeset
//
// swagger:model v2controllers.CreatableChangeset
type V2controllersCreatableChangeset struct {

	// chart release
	ChartRelease string `json:"chartRelease,omitempty"`

	// to app version branch
	ToAppVersionBranch string `json:"toAppVersionBranch,omitempty"`

	// to app version commit
	ToAppVersionCommit string `json:"toAppVersionCommit,omitempty"`

	// to app version exact
	ToAppVersionExact string `json:"toAppVersionExact,omitempty"`

	// to app version resolver
	ToAppVersionResolver string `json:"toAppVersionResolver,omitempty"`

	// to chart version exact
	ToChartVersionExact string `json:"toChartVersionExact,omitempty"`

	// to chart version resolver
	ToChartVersionResolver string `json:"toChartVersionResolver,omitempty"`

	// to helmfile ref
	ToHelmfileRef string `json:"toHelmfileRef,omitempty"`
}

// Validate validates this v2controllers creatable changeset
func (m *V2controllersCreatableChangeset) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2controllers creatable changeset based on context it is used
func (m *V2controllersCreatableChangeset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCreatableChangeset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCreatableChangeset) UnmarshalBinary(b []byte) error {
	var res V2controllersCreatableChangeset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

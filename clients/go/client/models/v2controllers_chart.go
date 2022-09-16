// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersChart v2controllers chart
//
// swagger:model v2controllers.Chart
type V2controllersChart struct {

	// app image git main branch
	AppImageGitMainBranch string `json:"appImageGitMainBranch,omitempty"`

	// app image git repo
	AppImageGitRepo string `json:"appImageGitRepo,omitempty"`

	// chart repo
	ChartRepo *string `json:"chartRepo,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// default port
	DefaultPort *int64 `json:"defaultPort,omitempty"`

	// default protocol
	DefaultProtocol *string `json:"defaultProtocol,omitempty"`

	// When creating, will default to the name of the chart
	DefaultSubdomain string `json:"defaultSubdomain,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// Required when creating
	Name string `json:"name,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this v2controllers chart
func (m *V2controllersChart) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2controllers chart based on context it is used
func (m *V2controllersChart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersChart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersChart) UnmarshalBinary(b []byte) error {
	var res V2controllersChart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

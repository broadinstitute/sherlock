// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersEditableChart v2controllers editable chart
//
// swagger:model v2controllers.EditableChart
type V2controllersEditableChart struct {

	// app image git main branch
	AppImageGitMainBranch string `json:"appImageGitMainBranch,omitempty"`

	// app image git repo
	AppImageGitRepo string `json:"appImageGitRepo,omitempty"`

	// Indicates if the default subdomain, protocol, and port fields are relevant for this chart
	ChartExposesEndpoint *bool `json:"chartExposesEndpoint,omitempty"`

	// chart repo
	ChartRepo *string `json:"chartRepo,omitempty"`

	// default port
	DefaultPort *int64 `json:"defaultPort,omitempty"`

	// default protocol
	DefaultProtocol *string `json:"defaultProtocol,omitempty"`

	// When creating, will default to the name of the chart
	DefaultSubdomain string `json:"defaultSubdomain,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Indicates whether a chart requires config rendering from firecloud-develop
	LegacyConfigsEnabled *bool `json:"legacyConfigsEnabled,omitempty"`

	// playbook URL
	PlaybookURL string `json:"playbookURL,omitempty"`
}

// Validate validates this v2controllers editable chart
func (m *V2controllersEditableChart) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2controllers editable chart based on context it is used
func (m *V2controllersEditableChart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersEditableChart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersEditableChart) UnmarshalBinary(b []byte) error {
	var res V2controllersEditableChart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
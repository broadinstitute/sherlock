// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersCreatableAppVersion v2controllers creatable app version
//
// swagger:model v2controllers.CreatableAppVersion
type V2controllersCreatableAppVersion struct {

	// Required when creating
	AppVersion string `json:"appVersion,omitempty"`

	// Required when creating
	Chart string `json:"chart,omitempty"`

	// Generally the Git commit message
	Description string `json:"description,omitempty"`

	// git branch
	GitBranch string `json:"gitBranch,omitempty"`

	// git commit
	GitCommit string `json:"gitCommit,omitempty"`

	// parent app version
	ParentAppVersion string `json:"parentAppVersion,omitempty"`
}

// Validate validates this v2controllers creatable app version
func (m *V2controllersCreatableAppVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2controllers creatable app version based on context it is used
func (m *V2controllersCreatableAppVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCreatableAppVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCreatableAppVersion) UnmarshalBinary(b []byte) error {
	var res V2controllersCreatableAppVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
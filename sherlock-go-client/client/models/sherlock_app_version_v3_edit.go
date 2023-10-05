// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SherlockAppVersionV3Edit sherlock app version v3 edit
//
// swagger:model sherlock.AppVersionV3Edit
type SherlockAppVersionV3Edit struct {

	// Generally the Git commit message
	Description string `json:"description,omitempty"`
}

// Validate validates this sherlock app version v3 edit
func (m *SherlockAppVersionV3Edit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sherlock app version v3 edit based on context it is used
func (m *SherlockAppVersionV3Edit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockAppVersionV3Edit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockAppVersionV3Edit) UnmarshalBinary(b []byte) error {
	var res SherlockAppVersionV3Edit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

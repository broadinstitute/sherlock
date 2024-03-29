// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MiscConnectionCheckResponse misc connection check response
//
// swagger:model misc.ConnectionCheckResponse
type MiscConnectionCheckResponse struct {

	// Always true
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this misc connection check response
func (m *MiscConnectionCheckResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this misc connection check response based on context it is used
func (m *MiscConnectionCheckResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MiscConnectionCheckResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MiscConnectionCheckResponse) UnmarshalBinary(b []byte) error {
	var res MiscConnectionCheckResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

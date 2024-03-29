// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SherlockPagerdutyIntegrationV3Edit sherlock pagerduty integration v3 edit
//
// swagger:model sherlock.PagerdutyIntegrationV3Edit
type SherlockPagerdutyIntegrationV3Edit struct {

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this sherlock pagerduty integration v3 edit
func (m *SherlockPagerdutyIntegrationV3Edit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sherlock pagerduty integration v3 edit based on context it is used
func (m *SherlockPagerdutyIntegrationV3Edit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockPagerdutyIntegrationV3Edit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockPagerdutyIntegrationV3Edit) UnmarshalBinary(b []byte) error {
	var res SherlockPagerdutyIntegrationV3Edit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PagerdutySendAlertResponse pagerduty send alert response
//
// swagger:model pagerduty.SendAlertResponse
type PagerdutySendAlertResponse struct {

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this pagerduty send alert response
func (m *PagerdutySendAlertResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagerduty send alert response based on context it is used
func (m *PagerdutySendAlertResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagerdutySendAlertResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagerdutySendAlertResponse) UnmarshalBinary(b []byte) error {
	var res PagerdutySendAlertResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

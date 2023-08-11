// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SherlockSlackDeployHookV3Edit sherlock slack deploy hook v3 edit
//
// swagger:model sherlock.SlackDeployHookV3Edit
type SherlockSlackDeployHookV3Edit struct {

	// on failure
	OnFailure bool `json:"onFailure,omitempty"`

	// on success
	OnSuccess bool `json:"onSuccess,omitempty"`

	// slack channel
	SlackChannel string `json:"slackChannel,omitempty"`
}

// Validate validates this sherlock slack deploy hook v3 edit
func (m *SherlockSlackDeployHookV3Edit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sherlock slack deploy hook v3 edit based on context it is used
func (m *SherlockSlackDeployHookV3Edit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockSlackDeployHookV3Edit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockSlackDeployHookV3Edit) UnmarshalBinary(b []byte) error {
	var res SherlockSlackDeployHookV3Edit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SherlockSlackDeployHookV3 sherlock slack deploy hook v3
//
// swagger:model sherlock.SlackDeployHookV3
type SherlockSlackDeployHookV3 struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// mention people
	MentionPeople bool `json:"mentionPeople,omitempty"`

	// on chart release
	OnChartRelease string `json:"onChartRelease,omitempty"`

	// on environment
	OnEnvironment string `json:"onEnvironment,omitempty"`

	// on failure
	OnFailure bool `json:"onFailure,omitempty"`

	// on success
	OnSuccess bool `json:"onSuccess,omitempty"`

	// slack channel
	SlackChannel string `json:"slackChannel,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this sherlock slack deploy hook v3
func (m *SherlockSlackDeployHookV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SherlockSlackDeployHookV3) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockSlackDeployHookV3) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sherlock slack deploy hook v3 based on context it is used
func (m *SherlockSlackDeployHookV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockSlackDeployHookV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockSlackDeployHookV3) UnmarshalBinary(b []byte) error {
	var res SherlockSlackDeployHookV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

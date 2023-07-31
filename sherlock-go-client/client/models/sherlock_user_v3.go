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

// SherlockUserV3 sherlock user v3
//
// swagger:model sherlock.UserV3
type SherlockUserV3 struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// github ID
	GithubID string `json:"githubID,omitempty"`

	// github username
	GithubUsername string `json:"githubUsername,omitempty"`

	// google ID
	GoogleID string `json:"googleID,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Controls whether Sherlock should automatically update the user's name based on a connected GitHub identity.
	// Will be set to true if the user account has no name and a GitHub account is linked.
	NameInferredFromGithub bool `json:"nameInferredFromGithub,omitempty"`

	// Available only in responses; describes the user's production-suitability
	SuitabilityDescription string `json:"suitabilityDescription,omitempty"`

	// Available only in responses; indicates whether the user is production-suitable
	Suitable bool `json:"suitable,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this sherlock user v3
func (m *SherlockUserV3) Validate(formats strfmt.Registry) error {
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

func (m *SherlockUserV3) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockUserV3) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sherlock user v3 based on context it is used
func (m *SherlockUserV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockUserV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockUserV3) UnmarshalBinary(b []byte) error {
	var res SherlockUserV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
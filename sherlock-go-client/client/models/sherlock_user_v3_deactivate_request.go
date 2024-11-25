// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SherlockUserV3DeactivateRequest sherlock user v3 deactivate request
//
// swagger:model sherlock.UserV3DeactivateRequest
type SherlockUserV3DeactivateRequest struct {

	// suspend email handles across google workspace domains
	SuspendEmailHandlesAcrossGoogleWorkspaceDomains []string `json:"suspendEmailHandlesAcrossGoogleWorkspaceDomains"`

	// Domain of UserEmails that can be swapped out for the domains in SuspendEmailHandlesAcrossGoogleWorkspaceDomains
	UserEmailHomeDomain *string `json:"userEmailHomeDomain,omitempty"`

	// user emails
	UserEmails []string `json:"userEmails"`
}

// Validate validates this sherlock user v3 deactivate request
func (m *SherlockUserV3DeactivateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sherlock user v3 deactivate request based on context it is used
func (m *SherlockUserV3DeactivateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockUserV3DeactivateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockUserV3DeactivateRequest) UnmarshalBinary(b []byte) error {
	var res SherlockUserV3DeactivateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
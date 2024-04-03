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

// SherlockGithubActionsJobV3Create sherlock github actions job v3 create
//
// swagger:model sherlock.GithubActionsJobV3Create
type SherlockGithubActionsJobV3Create struct {

	// github actions attempt number
	GithubActionsAttemptNumber int64 `json:"githubActionsAttemptNumber,omitempty"`

	// github actions job ID
	GithubActionsJobID int64 `json:"githubActionsJobID,omitempty"`

	// github actions owner
	GithubActionsOwner string `json:"githubActionsOwner,omitempty"`

	// github actions repo
	GithubActionsRepo string `json:"githubActionsRepo,omitempty"`

	// github actions run ID
	GithubActionsRunID int64 `json:"githubActionsRunID,omitempty"`

	// job created at
	// Format: date-time
	JobCreatedAt strfmt.DateTime `json:"jobCreatedAt,omitempty"`

	// job started at
	// Format: date-time
	JobStartedAt strfmt.DateTime `json:"jobStartedAt,omitempty"`

	// job terminal at
	// Format: date-time
	JobTerminalAt strfmt.DateTime `json:"jobTerminalAt,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this sherlock github actions job v3 create
func (m *SherlockGithubActionsJobV3Create) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobTerminalAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SherlockGithubActionsJobV3Create) validateJobCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.JobCreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("jobCreatedAt", "body", "date-time", m.JobCreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockGithubActionsJobV3Create) validateJobStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.JobStartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("jobStartedAt", "body", "date-time", m.JobStartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockGithubActionsJobV3Create) validateJobTerminalAt(formats strfmt.Registry) error {
	if swag.IsZero(m.JobTerminalAt) { // not required
		return nil
	}

	if err := validate.FormatOf("jobTerminalAt", "body", "date-time", m.JobTerminalAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sherlock github actions job v3 create based on context it is used
func (m *SherlockGithubActionsJobV3Create) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockGithubActionsJobV3Create) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockGithubActionsJobV3Create) UnmarshalBinary(b []byte) error {
	var res SherlockGithubActionsJobV3Create
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
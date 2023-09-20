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

// SherlockChartV3 sherlock chart v3
//
// swagger:model sherlock.ChartV3
type SherlockChartV3 struct {

	// app image git main branch
	AppImageGitMainBranch string `json:"appImageGitMainBranch,omitempty"`

	// app image git repo
	AppImageGitRepo string `json:"appImageGitRepo,omitempty"`

	// Indicates if the default subdomain, protocol, and port fields are relevant for this chart
	ChartExposesEndpoint *bool `json:"chartExposesEndpoint,omitempty"`

	// chart repo
	ChartRepo *string `json:"chartRepo,omitempty"`

	// ci identifier
	CiIdentifier *SherlockCiIdentifierV3 `json:"ciIdentifier,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// default port
	DefaultPort *int64 `json:"defaultPort,omitempty"`

	// default protocol
	DefaultProtocol *string `json:"defaultProtocol,omitempty"`

	// When creating, will default to the name of the chart
	DefaultSubdomain string `json:"defaultSubdomain,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// Indicates whether a chart requires config rendering from firecloud-develop
	LegacyConfigsEnabled *bool `json:"legacyConfigsEnabled,omitempty"`

	// Required when creating
	Name string `json:"name,omitempty"`

	// pact participant
	PactParticipant *bool `json:"pactParticipant,omitempty"`

	// playbook URL
	PlaybookURL string `json:"playbookURL,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this sherlock chart v3
func (m *SherlockChartV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCiIdentifier(formats); err != nil {
		res = append(res, err)
	}

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

func (m *SherlockChartV3) validateCiIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.CiIdentifier) { // not required
		return nil
	}

	if m.CiIdentifier != nil {
		if err := m.CiIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ciIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ciIdentifier")
			}
			return err
		}
	}

	return nil
}

func (m *SherlockChartV3) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockChartV3) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sherlock chart v3 based on the context it is used
func (m *SherlockChartV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCiIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SherlockChartV3) contextValidateCiIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if m.CiIdentifier != nil {
		if err := m.CiIdentifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ciIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ciIdentifier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SherlockChartV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockChartV3) UnmarshalBinary(b []byte) error {
	var res SherlockChartV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
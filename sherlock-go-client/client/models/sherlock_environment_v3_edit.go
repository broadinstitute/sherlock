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

// SherlockEnvironmentV3Edit sherlock environment v3 edit
//
// swagger:model sherlock.EnvironmentV3Edit
type SherlockEnvironmentV3Edit struct {

	// base domain
	BaseDomain *string `json:"baseDomain,omitempty"`

	// default cluster
	DefaultCluster string `json:"defaultCluster,omitempty"`

	// should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	DefaultFirecloudDevelopRef *string `json:"defaultFirecloudDevelopRef,omitempty"`

	// If set, the BEE will be automatically deleted after this time (thelma checks this field)
	DeleteAfter string `json:"deleteAfter,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// name prefixes domain
	NamePrefixesDomain *bool `json:"namePrefixesDomain,omitempty"`

	// Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
	Offline *bool `json:"offline,omitempty"`

	// When enabled, the BEE will be slated to go offline around the begin time each day
	OfflineScheduleBeginEnabled bool `json:"offlineScheduleBeginEnabled,omitempty"`

	// Stored with timezone to determine day of the week
	// Format: date-time
	OfflineScheduleBeginTime strfmt.DateTime `json:"offlineScheduleBeginTime,omitempty"`

	// When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
	OfflineScheduleEndEnabled bool `json:"offlineScheduleEndEnabled,omitempty"`

	// Stored with timezone to determine day of the week
	// Format: date-time
	OfflineScheduleEndTime strfmt.DateTime `json:"offlineScheduleEndTime,omitempty"`

	// offline schedule end weekends
	OfflineScheduleEndWeekends bool `json:"offlineScheduleEndWeekends,omitempty"`

	// When creating, will default to you
	Owner string `json:"owner,omitempty"`

	// pact identifier
	PactIdentifier string `json:"pactIdentifier,omitempty"`

	// pagerduty integration
	PagerdutyIntegration string `json:"pagerdutyIntegration,omitempty"`

	// Used to protect specific BEEs from deletion (thelma checks this field)
	PreventDeletion *bool `json:"preventDeletion,omitempty"`

	// requires suitability
	RequiresSuitability *bool `json:"requiresSuitability,omitempty"`
}

// Validate validates this sherlock environment v3 edit
func (m *SherlockEnvironmentV3Edit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOfflineScheduleBeginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineScheduleEndTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SherlockEnvironmentV3Edit) validateOfflineScheduleBeginTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineScheduleBeginTime) { // not required
		return nil
	}

	if err := validate.FormatOf("offlineScheduleBeginTime", "body", "date-time", m.OfflineScheduleBeginTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockEnvironmentV3Edit) validateOfflineScheduleEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineScheduleEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("offlineScheduleEndTime", "body", "date-time", m.OfflineScheduleEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sherlock environment v3 edit based on context it is used
func (m *SherlockEnvironmentV3Edit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockEnvironmentV3Edit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockEnvironmentV3Edit) UnmarshalBinary(b []byte) error {
	var res SherlockEnvironmentV3Edit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

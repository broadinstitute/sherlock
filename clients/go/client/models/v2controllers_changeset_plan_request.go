// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersChangesetPlanRequest v2controllers changeset plan request
//
// swagger:model v2controllers.ChangesetPlanRequest
type V2controllersChangesetPlanRequest struct {

	// chart releases
	ChartReleases []*V2controllersChangesetPlanRequestChartReleaseEntry `json:"chartReleases"`

	// environments
	Environments []*V2controllersChangesetPlanRequestEnvironmentEntry `json:"environments"`
}

// Validate validates this v2controllers changeset plan request
func (m *V2controllersChangesetPlanRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartReleases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersChangesetPlanRequest) validateChartReleases(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartReleases) { // not required
		return nil
	}

	for i := 0; i < len(m.ChartReleases); i++ {
		if swag.IsZero(m.ChartReleases[i]) { // not required
			continue
		}

		if m.ChartReleases[i] != nil {
			if err := m.ChartReleases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chartReleases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chartReleases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2controllersChangesetPlanRequest) validateEnvironments(formats strfmt.Registry) error {
	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v2controllers changeset plan request based on the context it is used
func (m *V2controllersChangesetPlanRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChartReleases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersChangesetPlanRequest) contextValidateChartReleases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChartReleases); i++ {

		if m.ChartReleases[i] != nil {
			if err := m.ChartReleases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chartReleases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chartReleases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2controllersChangesetPlanRequest) contextValidateEnvironments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Environments); i++ {

		if m.Environments[i] != nil {
			if err := m.Environments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersChangesetPlanRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersChangesetPlanRequest) UnmarshalBinary(b []byte) error {
	var res V2controllersChangesetPlanRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

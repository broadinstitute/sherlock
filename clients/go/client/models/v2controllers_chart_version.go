// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2controllersChartVersion v2controllers chart version
//
// swagger:model v2controllers.ChartVersion
type V2controllersChartVersion struct {

	// Required when creating
	Chart string `json:"chart,omitempty"`

	// chart info
	ChartInfo *V2controllersChart `json:"chartInfo,omitempty"`

	// Required when creating
	ChartVersion string `json:"chartVersion,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// Generally the Git commit message
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// parent chart version
	ParentChartVersion string `json:"parentChartVersion,omitempty"`

	// parent chart version info
	ParentChartVersionInfo interface{} `json:"parentChartVersionInfo,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this v2controllers chart version
func (m *V2controllersChartVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersChartVersion) validateChartInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartInfo) { // not required
		return nil
	}

	if m.ChartInfo != nil {
		if err := m.ChartInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chartInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chartInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2controllers chart version based on the context it is used
func (m *V2controllersChartVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChartInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersChartVersion) contextValidateChartInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ChartInfo != nil {
		if err := m.ChartInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chartInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chartInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersChartVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersChartVersion) UnmarshalBinary(b []byte) error {
	var res V2controllersChartVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

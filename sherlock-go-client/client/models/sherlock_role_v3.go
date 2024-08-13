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
	"github.com/go-openapi/validate"
)

// SherlockRoleV3 sherlock role v3
//
// swagger:model sherlock.RoleV3
type SherlockRoleV3 struct {

	// assignments
	Assignments []*SherlockRoleAssignmentV3 `json:"assignments"`

	// When true, Sherlock will automatically assign all users to this role who do not already have a role assignment
	AutoAssignAllUsers bool `json:"autoAssignAllUsers,omitempty"`

	// can be glass broken by role
	CanBeGlassBrokenByRole int64 `json:"canBeGlassBrokenByRole,omitempty"`

	// can be glass broken by role info
	CanBeGlassBrokenByRoleInfo interface{} `json:"canBeGlassBrokenByRoleInfo,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// default glass break duration
	DefaultGlassBreakDuration string `json:"defaultGlassBreakDuration,omitempty"`

	// grants broad institute group
	GrantsBroadInstituteGroup string `json:"grantsBroadInstituteGroup,omitempty"`

	// grants dev azure group
	GrantsDevAzureGroup string `json:"grantsDevAzureGroup,omitempty"`

	// grants dev firecloud group
	GrantsDevFirecloudGroup string `json:"grantsDevFirecloudGroup,omitempty"`

	// grants prod azure group
	GrantsProdAzureGroup string `json:"grantsProdAzureGroup,omitempty"`

	// grants prod firecloud group
	GrantsProdFirecloudGroup string `json:"grantsProdFirecloudGroup,omitempty"`

	// grants qa firecloud group
	GrantsQaFirecloudGroup string `json:"grantsQaFirecloudGroup,omitempty"`

	// grants sherlock super admin
	GrantsSherlockSuperAdmin bool `json:"grantsSherlockSuperAdmin,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// When true, the "suspended" field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field
	SuspendNonSuitableUsers bool `json:"suspendNonSuitableUsers,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this sherlock role v3
func (m *SherlockRoleV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignments(formats); err != nil {
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

func (m *SherlockRoleV3) validateAssignments(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignments) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignments); i++ {
		if swag.IsZero(m.Assignments[i]) { // not required
			continue
		}

		if m.Assignments[i] != nil {
			if err := m.Assignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SherlockRoleV3) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockRoleV3) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sherlock role v3 based on the context it is used
func (m *SherlockRoleV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SherlockRoleV3) contextValidateAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Assignments); i++ {

		if m.Assignments[i] != nil {
			if err := m.Assignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SherlockRoleV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockRoleV3) UnmarshalBinary(b []byte) error {
	var res SherlockRoleV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2controllersCreatableChartRelease v2controllers creatable chart release
//
// swagger:model v2controllers.CreatableChartRelease
type V2controllersCreatableChartRelease struct {

	// When creating, will default to the app's mainline branch if no other app version info is present
	AppVersionBranch string `json:"appVersionBranch,omitempty"`

	// app version commit
	AppVersionCommit string `json:"appVersionCommit,omitempty"`

	// app version exact
	AppVersionExact string `json:"appVersionExact,omitempty"`

	// // When creating, will default to automatically reference any provided app version fields
	// Enum: [branch commit exact none]
	AppVersionResolver string `json:"appVersionResolver,omitempty"`

	// Required when creating
	Chart string `json:"chart,omitempty"`

	// chart version exact
	ChartVersionExact string `json:"chartVersionExact,omitempty"`

	// When creating, will default to automatically reference any provided chart version
	// Enum: [latest exact]
	ChartVersionResolver string `json:"chartVersionResolver,omitempty"`

	// When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Cluster string `json:"cluster,omitempty"`

	// Either this or cluster must be provided.
	Environment string `json:"environment,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// When creating, will be calculated if left empty
	Name string `json:"name,omitempty"`

	// When creating, will default to the environment's default namespace, if provided
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this v2controllers creatable chart release
func (m *V2controllersCreatableChartRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppVersionResolver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChartVersionResolver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v2controllersCreatableChartReleaseTypeAppVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["branch","commit","exact","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2controllersCreatableChartReleaseTypeAppVersionResolverPropEnum = append(v2controllersCreatableChartReleaseTypeAppVersionResolverPropEnum, v)
	}
}

const (

	// V2controllersCreatableChartReleaseAppVersionResolverBranch captures enum value "branch"
	V2controllersCreatableChartReleaseAppVersionResolverBranch string = "branch"

	// V2controllersCreatableChartReleaseAppVersionResolverCommit captures enum value "commit"
	V2controllersCreatableChartReleaseAppVersionResolverCommit string = "commit"

	// V2controllersCreatableChartReleaseAppVersionResolverExact captures enum value "exact"
	V2controllersCreatableChartReleaseAppVersionResolverExact string = "exact"

	// V2controllersCreatableChartReleaseAppVersionResolverNone captures enum value "none"
	V2controllersCreatableChartReleaseAppVersionResolverNone string = "none"
)

// prop value enum
func (m *V2controllersCreatableChartRelease) validateAppVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2controllersCreatableChartReleaseTypeAppVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2controllersCreatableChartRelease) validateAppVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.AppVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppVersionResolverEnum("appVersionResolver", "body", m.AppVersionResolver); err != nil {
		return err
	}

	return nil
}

var v2controllersCreatableChartReleaseTypeChartVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["latest","exact"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2controllersCreatableChartReleaseTypeChartVersionResolverPropEnum = append(v2controllersCreatableChartReleaseTypeChartVersionResolverPropEnum, v)
	}
}

const (

	// V2controllersCreatableChartReleaseChartVersionResolverLatest captures enum value "latest"
	V2controllersCreatableChartReleaseChartVersionResolverLatest string = "latest"

	// V2controllersCreatableChartReleaseChartVersionResolverExact captures enum value "exact"
	V2controllersCreatableChartReleaseChartVersionResolverExact string = "exact"
)

// prop value enum
func (m *V2controllersCreatableChartRelease) validateChartVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2controllersCreatableChartReleaseTypeChartVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2controllersCreatableChartRelease) validateChartVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateChartVersionResolverEnum("chartVersionResolver", "body", m.ChartVersionResolver); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v2controllers creatable chart release based on context it is used
func (m *V2controllersCreatableChartRelease) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersCreatableChartRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersCreatableChartRelease) UnmarshalBinary(b []byte) error {
	var res V2controllersCreatableChartRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

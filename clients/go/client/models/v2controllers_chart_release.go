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

// V2controllersChartRelease v2controllers chart release
//
// swagger:model v2controllers.ChartRelease
type V2controllersChartRelease struct {

	// When creating, will default to the app's mainline branch if no other app version info is present
	AppVersionBranch string `json:"appVersionBranch,omitempty"`

	// app version commit
	AppVersionCommit string `json:"appVersionCommit,omitempty"`

	// app version exact
	AppVersionExact string `json:"appVersionExact,omitempty"`

	// app version info
	AppVersionInfo *V2controllersAppVersion `json:"appVersionInfo,omitempty"`

	// app version reference
	AppVersionReference string `json:"appVersionReference,omitempty"`

	// // When creating, will default to automatically reference any provided app version fields
	// Enum: [branch commit exact none]
	AppVersionResolver string `json:"appVersionResolver,omitempty"`

	// Required when creating
	Chart string `json:"chart,omitempty"`

	// chart info
	ChartInfo *V2controllersChart `json:"chartInfo,omitempty"`

	// chart version exact
	ChartVersionExact string `json:"chartVersionExact,omitempty"`

	// chart version info
	ChartVersionInfo *V2controllersChartVersion `json:"chartVersionInfo,omitempty"`

	// chart version reference
	ChartVersionReference string `json:"chartVersionReference,omitempty"`

	// When creating, will default to automatically reference any provided chart version
	// Enum: [latest exact]
	ChartVersionResolver string `json:"chartVersionResolver,omitempty"`

	// When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Cluster string `json:"cluster,omitempty"`

	// cluster info
	ClusterInfo *V2controllersCluster `json:"clusterInfo,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// Calculated field
	DestinationType string `json:"destinationType,omitempty"`

	// Either this or cluster must be provided.
	Environment string `json:"environment,omitempty"`

	// environment info
	EnvironmentInfo *V2controllersEnvironment `json:"environmentInfo,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// When creating, will be calculated if left empty
	Name string `json:"name,omitempty"`

	// When creating, will default to the environment's default namespace, if provided
	Namespace string `json:"namespace,omitempty"`

	// When creating, will use the chart's default if left empty
	Port int64 `json:"port,omitempty"`

	// When creating, will use the chart's default if left empty
	Protocol string `json:"protocol,omitempty"`

	// When creating, will use the chart's default if left empty
	Subdomain string `json:"subdomain,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this v2controllers chart release
func (m *V2controllersChartRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppVersionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppVersionResolver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChartInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChartVersionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChartVersionResolver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersChartRelease) validateAppVersionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AppVersionInfo) { // not required
		return nil
	}

	if m.AppVersionInfo != nil {
		if err := m.AppVersionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appVersionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appVersionInfo")
			}
			return err
		}
	}

	return nil
}

var v2controllersChartReleaseTypeAppVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["branch","commit","exact","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2controllersChartReleaseTypeAppVersionResolverPropEnum = append(v2controllersChartReleaseTypeAppVersionResolverPropEnum, v)
	}
}

const (

	// V2controllersChartReleaseAppVersionResolverBranch captures enum value "branch"
	V2controllersChartReleaseAppVersionResolverBranch string = "branch"

	// V2controllersChartReleaseAppVersionResolverCommit captures enum value "commit"
	V2controllersChartReleaseAppVersionResolverCommit string = "commit"

	// V2controllersChartReleaseAppVersionResolverExact captures enum value "exact"
	V2controllersChartReleaseAppVersionResolverExact string = "exact"

	// V2controllersChartReleaseAppVersionResolverNone captures enum value "none"
	V2controllersChartReleaseAppVersionResolverNone string = "none"
)

// prop value enum
func (m *V2controllersChartRelease) validateAppVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2controllersChartReleaseTypeAppVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2controllersChartRelease) validateAppVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.AppVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppVersionResolverEnum("appVersionResolver", "body", m.AppVersionResolver); err != nil {
		return err
	}

	return nil
}

func (m *V2controllersChartRelease) validateChartInfo(formats strfmt.Registry) error {
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

func (m *V2controllersChartRelease) validateChartVersionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartVersionInfo) { // not required
		return nil
	}

	if m.ChartVersionInfo != nil {
		if err := m.ChartVersionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chartVersionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chartVersionInfo")
			}
			return err
		}
	}

	return nil
}

var v2controllersChartReleaseTypeChartVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["latest","exact"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2controllersChartReleaseTypeChartVersionResolverPropEnum = append(v2controllersChartReleaseTypeChartVersionResolverPropEnum, v)
	}
}

const (

	// V2controllersChartReleaseChartVersionResolverLatest captures enum value "latest"
	V2controllersChartReleaseChartVersionResolverLatest string = "latest"

	// V2controllersChartReleaseChartVersionResolverExact captures enum value "exact"
	V2controllersChartReleaseChartVersionResolverExact string = "exact"
)

// prop value enum
func (m *V2controllersChartRelease) validateChartVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2controllersChartReleaseTypeChartVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2controllersChartRelease) validateChartVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateChartVersionResolverEnum("chartVersionResolver", "body", m.ChartVersionResolver); err != nil {
		return err
	}

	return nil
}

func (m *V2controllersChartRelease) validateClusterInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterInfo) { // not required
		return nil
	}

	if m.ClusterInfo != nil {
		if err := m.ClusterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersChartRelease) validateEnvironmentInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvironmentInfo) { // not required
		return nil
	}

	if m.EnvironmentInfo != nil {
		if err := m.EnvironmentInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environmentInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environmentInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2controllers chart release based on the context it is used
func (m *V2controllersChartRelease) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppVersionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChartInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChartVersionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironmentInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2controllersChartRelease) contextValidateAppVersionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AppVersionInfo != nil {
		if err := m.AppVersionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appVersionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appVersionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersChartRelease) contextValidateChartInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V2controllersChartRelease) contextValidateChartVersionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ChartVersionInfo != nil {
		if err := m.ChartVersionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chartVersionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chartVersionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersChartRelease) contextValidateClusterInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterInfo != nil {
		if err := m.ClusterInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersChartRelease) contextValidateEnvironmentInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.EnvironmentInfo != nil {
		if err := m.EnvironmentInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environmentInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environmentInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2controllersChartRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersChartRelease) UnmarshalBinary(b []byte) error {
	var res V2controllersChartRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

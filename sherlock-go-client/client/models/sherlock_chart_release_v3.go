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

// SherlockChartReleaseV3 sherlock chart release v3
//
// swagger:model sherlock.ChartReleaseV3
type SherlockChartReleaseV3 struct {

	// When creating, will default to the app's mainline branch if no other app version info is present
	AppVersionBranch string `json:"appVersionBranch,omitempty"`

	// app version commit
	AppVersionCommit string `json:"appVersionCommit,omitempty"`

	// app version exact
	AppVersionExact string `json:"appVersionExact,omitempty"`

	// app version follow chart release
	AppVersionFollowChartRelease string `json:"appVersionFollowChartRelease,omitempty"`

	// app version info
	AppVersionInfo *SherlockAppVersionV3 `json:"appVersionInfo,omitempty"`

	// app version reference
	AppVersionReference string `json:"appVersionReference,omitempty"`

	// // When creating, will default to automatically reference any provided app version fields
	// Enum: [branch commit exact follow none]
	AppVersionResolver string `json:"appVersionResolver,omitempty"`

	// Required when creating
	Chart string `json:"chart,omitempty"`

	// chart info
	ChartInfo *SherlockChartV3 `json:"chartInfo,omitempty"`

	// chart version exact
	ChartVersionExact string `json:"chartVersionExact,omitempty"`

	// chart version follow chart release
	ChartVersionFollowChartRelease string `json:"chartVersionFollowChartRelease,omitempty"`

	// chart version info
	ChartVersionInfo *SherlockChartVersionV3 `json:"chartVersionInfo,omitempty"`

	// chart version reference
	ChartVersionReference string `json:"chartVersionReference,omitempty"`

	// When creating, will default to automatically reference any provided chart version
	// Enum: [latest exact follow]
	ChartVersionResolver string `json:"chartVersionResolver,omitempty"`

	// ci identifier
	CiIdentifier *SherlockCiIdentifierV3 `json:"ciIdentifier,omitempty"`

	// When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Cluster string `json:"cluster,omitempty"`

	// cluster info
	ClusterInfo *SherlockClusterV3 `json:"clusterInfo,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Calculated field
	DestinationType string `json:"destinationType,omitempty"`

	// Either this or cluster must be provided.
	Environment string `json:"environment,omitempty"`

	// environment info
	EnvironmentInfo *SherlockEnvironmentV3 `json:"environmentInfo,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// helmfile ref enabled
	HelmfileRefEnabled *bool `json:"helmfileRefEnabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// included in bulk changesets
	IncludedInBulkChangesets *bool `json:"includedInBulkChangesets,omitempty"`

	// When creating, will be calculated if left empty
	Name string `json:"name,omitempty"`

	// When creating, will default to the environment's default namespace, if provided
	Namespace string `json:"namespace,omitempty"`

	// pagerduty integration
	PagerdutyIntegration string `json:"pagerdutyIntegration,omitempty"`

	// pagerduty integration info
	PagerdutyIntegrationInfo *SherlockPagerdutyIntegrationV3 `json:"pagerdutyIntegrationInfo,omitempty"`

	// When creating, will use the chart's default if left empty
	Port int64 `json:"port,omitempty"`

	// When creating, will use the chart's default if left empty
	Protocol string `json:"protocol,omitempty"`

	// resolved at
	// Format: date-time
	ResolvedAt strfmt.DateTime `json:"resolvedAt,omitempty"`

	// When creating, will use the chart's default if left empty
	Subdomain string `json:"subdomain,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this sherlock chart release v3
func (m *SherlockChartReleaseV3) Validate(formats strfmt.Registry) error {
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

	if err := m.validateCiIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerdutyIntegrationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedAt(formats); err != nil {
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

func (m *SherlockChartReleaseV3) validateAppVersionInfo(formats strfmt.Registry) error {
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

var sherlockChartReleaseV3TypeAppVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["branch","commit","exact","follow","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sherlockChartReleaseV3TypeAppVersionResolverPropEnum = append(sherlockChartReleaseV3TypeAppVersionResolverPropEnum, v)
	}
}

const (

	// SherlockChartReleaseV3AppVersionResolverBranch captures enum value "branch"
	SherlockChartReleaseV3AppVersionResolverBranch string = "branch"

	// SherlockChartReleaseV3AppVersionResolverCommit captures enum value "commit"
	SherlockChartReleaseV3AppVersionResolverCommit string = "commit"

	// SherlockChartReleaseV3AppVersionResolverExact captures enum value "exact"
	SherlockChartReleaseV3AppVersionResolverExact string = "exact"

	// SherlockChartReleaseV3AppVersionResolverFollow captures enum value "follow"
	SherlockChartReleaseV3AppVersionResolverFollow string = "follow"

	// SherlockChartReleaseV3AppVersionResolverNone captures enum value "none"
	SherlockChartReleaseV3AppVersionResolverNone string = "none"
)

// prop value enum
func (m *SherlockChartReleaseV3) validateAppVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sherlockChartReleaseV3TypeAppVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SherlockChartReleaseV3) validateAppVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.AppVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppVersionResolverEnum("appVersionResolver", "body", m.AppVersionResolver); err != nil {
		return err
	}

	return nil
}

func (m *SherlockChartReleaseV3) validateChartInfo(formats strfmt.Registry) error {
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

func (m *SherlockChartReleaseV3) validateChartVersionInfo(formats strfmt.Registry) error {
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

var sherlockChartReleaseV3TypeChartVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["latest","exact","follow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sherlockChartReleaseV3TypeChartVersionResolverPropEnum = append(sherlockChartReleaseV3TypeChartVersionResolverPropEnum, v)
	}
}

const (

	// SherlockChartReleaseV3ChartVersionResolverLatest captures enum value "latest"
	SherlockChartReleaseV3ChartVersionResolverLatest string = "latest"

	// SherlockChartReleaseV3ChartVersionResolverExact captures enum value "exact"
	SherlockChartReleaseV3ChartVersionResolverExact string = "exact"

	// SherlockChartReleaseV3ChartVersionResolverFollow captures enum value "follow"
	SherlockChartReleaseV3ChartVersionResolverFollow string = "follow"
)

// prop value enum
func (m *SherlockChartReleaseV3) validateChartVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sherlockChartReleaseV3TypeChartVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SherlockChartReleaseV3) validateChartVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateChartVersionResolverEnum("chartVersionResolver", "body", m.ChartVersionResolver); err != nil {
		return err
	}

	return nil
}

func (m *SherlockChartReleaseV3) validateCiIdentifier(formats strfmt.Registry) error {
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

func (m *SherlockChartReleaseV3) validateClusterInfo(formats strfmt.Registry) error {
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

func (m *SherlockChartReleaseV3) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockChartReleaseV3) validateEnvironmentInfo(formats strfmt.Registry) error {
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

func (m *SherlockChartReleaseV3) validatePagerdutyIntegrationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PagerdutyIntegrationInfo) { // not required
		return nil
	}

	if m.PagerdutyIntegrationInfo != nil {
		if err := m.PagerdutyIntegrationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerdutyIntegrationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerdutyIntegrationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SherlockChartReleaseV3) validateResolvedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolvedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("resolvedAt", "body", "date-time", m.ResolvedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SherlockChartReleaseV3) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sherlock chart release v3 based on the context it is used
func (m *SherlockChartReleaseV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateCiIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironmentInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagerdutyIntegrationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SherlockChartReleaseV3) contextValidateAppVersionInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SherlockChartReleaseV3) contextValidateChartInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SherlockChartReleaseV3) contextValidateChartVersionInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SherlockChartReleaseV3) contextValidateCiIdentifier(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SherlockChartReleaseV3) contextValidateClusterInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SherlockChartReleaseV3) contextValidateEnvironmentInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SherlockChartReleaseV3) contextValidatePagerdutyIntegrationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PagerdutyIntegrationInfo != nil {
		if err := m.PagerdutyIntegrationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerdutyIntegrationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerdutyIntegrationInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SherlockChartReleaseV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockChartReleaseV3) UnmarshalBinary(b []byte) error {
	var res SherlockChartReleaseV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

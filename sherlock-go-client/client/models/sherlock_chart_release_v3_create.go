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

// SherlockChartReleaseV3Create sherlock chart release v3 create
//
// swagger:model sherlock.ChartReleaseV3Create
type SherlockChartReleaseV3Create struct {

	// When creating, will default to the app's mainline branch if no other app version info is present
	AppVersionBranch string `json:"appVersionBranch,omitempty"`

	// app version commit
	AppVersionCommit string `json:"appVersionCommit,omitempty"`

	// app version exact
	AppVersionExact string `json:"appVersionExact,omitempty"`

	// app version follow chart release
	AppVersionFollowChartRelease string `json:"appVersionFollowChartRelease,omitempty"`

	// // When creating, will default to automatically reference any provided app version fields
	// Enum: [branch commit exact follow none]
	AppVersionResolver string `json:"appVersionResolver,omitempty"`

	// Required when creating
	Chart string `json:"chart,omitempty"`

	// chart version exact
	ChartVersionExact string `json:"chartVersionExact,omitempty"`

	// chart version follow chart release
	ChartVersionFollowChartRelease string `json:"chartVersionFollowChartRelease,omitempty"`

	// When creating, will default to automatically reference any provided chart version
	// Enum: [latest exact follow]
	ChartVersionResolver string `json:"chartVersionResolver,omitempty"`

	// When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
	Cluster string `json:"cluster,omitempty"`

	// Either this or cluster must be provided.
	Environment string `json:"environment,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// helmfile ref enabled
	HelmfileRefEnabled *bool `json:"helmfileRefEnabled,omitempty"`

	// included in bulk changesets
	IncludedInBulkChangesets *bool `json:"includedInBulkChangesets,omitempty"`

	// When creating, will be calculated if left empty
	Name string `json:"name,omitempty"`

	// When creating, will default to the environment's default namespace, if provided
	Namespace string `json:"namespace,omitempty"`

	// pagerduty integration
	PagerdutyIntegration string `json:"pagerdutyIntegration,omitempty"`

	// When creating, will use the chart's default if left empty
	Port int64 `json:"port,omitempty"`

	// When creating, will use the chart's default if left empty
	Protocol string `json:"protocol,omitempty"`

	// When creating, will use the chart's default if left empty
	Subdomain string `json:"subdomain,omitempty"`
}

// Validate validates this sherlock chart release v3 create
func (m *SherlockChartReleaseV3Create) Validate(formats strfmt.Registry) error {
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

var sherlockChartReleaseV3CreateTypeAppVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["branch","commit","exact","follow","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sherlockChartReleaseV3CreateTypeAppVersionResolverPropEnum = append(sherlockChartReleaseV3CreateTypeAppVersionResolverPropEnum, v)
	}
}

const (

	// SherlockChartReleaseV3CreateAppVersionResolverBranch captures enum value "branch"
	SherlockChartReleaseV3CreateAppVersionResolverBranch string = "branch"

	// SherlockChartReleaseV3CreateAppVersionResolverCommit captures enum value "commit"
	SherlockChartReleaseV3CreateAppVersionResolverCommit string = "commit"

	// SherlockChartReleaseV3CreateAppVersionResolverExact captures enum value "exact"
	SherlockChartReleaseV3CreateAppVersionResolverExact string = "exact"

	// SherlockChartReleaseV3CreateAppVersionResolverFollow captures enum value "follow"
	SherlockChartReleaseV3CreateAppVersionResolverFollow string = "follow"

	// SherlockChartReleaseV3CreateAppVersionResolverNone captures enum value "none"
	SherlockChartReleaseV3CreateAppVersionResolverNone string = "none"
)

// prop value enum
func (m *SherlockChartReleaseV3Create) validateAppVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sherlockChartReleaseV3CreateTypeAppVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SherlockChartReleaseV3Create) validateAppVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.AppVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppVersionResolverEnum("appVersionResolver", "body", m.AppVersionResolver); err != nil {
		return err
	}

	return nil
}

var sherlockChartReleaseV3CreateTypeChartVersionResolverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["latest","exact","follow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sherlockChartReleaseV3CreateTypeChartVersionResolverPropEnum = append(sherlockChartReleaseV3CreateTypeChartVersionResolverPropEnum, v)
	}
}

const (

	// SherlockChartReleaseV3CreateChartVersionResolverLatest captures enum value "latest"
	SherlockChartReleaseV3CreateChartVersionResolverLatest string = "latest"

	// SherlockChartReleaseV3CreateChartVersionResolverExact captures enum value "exact"
	SherlockChartReleaseV3CreateChartVersionResolverExact string = "exact"

	// SherlockChartReleaseV3CreateChartVersionResolverFollow captures enum value "follow"
	SherlockChartReleaseV3CreateChartVersionResolverFollow string = "follow"
)

// prop value enum
func (m *SherlockChartReleaseV3Create) validateChartVersionResolverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sherlockChartReleaseV3CreateTypeChartVersionResolverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SherlockChartReleaseV3Create) validateChartVersionResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartVersionResolver) { // not required
		return nil
	}

	// value enum
	if err := m.validateChartVersionResolverEnum("chartVersionResolver", "body", m.ChartVersionResolver); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sherlock chart release v3 create based on context it is used
func (m *SherlockChartReleaseV3Create) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SherlockChartReleaseV3Create) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SherlockChartReleaseV3Create) UnmarshalBinary(b []byte) error {
	var res SherlockChartReleaseV3Create
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// V2controllersEnvironment v2controllers environment
//
// swagger:model v2controllers.Environment
type V2controllersEnvironment struct {

	// auto delete
	AutoDelete *EnvironmentAutoDelete `json:"autoDelete,omitempty"`

	// If true when creating, dynamic environments copy from template and template environments get the honeycomb chart
	AutoPopulateChartReleases *bool `json:"autoPopulateChartReleases,omitempty"`

	// Required when creating
	Base string `json:"base,omitempty"`

	// base domain
	BaseDomain *string `json:"baseDomain,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// default cluster
	DefaultCluster string `json:"defaultCluster,omitempty"`

	// default cluster info
	DefaultClusterInfo *V2controllersCluster `json:"defaultClusterInfo,omitempty"`

	// should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
	DefaultFirecloudDevelopRef *string `json:"defaultFirecloudDevelopRef,omitempty"`

	// When creating, will be calculated if left empty
	DefaultNamespace string `json:"defaultNamespace,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// helmfile ref
	HelmfileRef *string `json:"helmfileRef,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// lifecycle
	Lifecycle *string `json:"lifecycle,omitempty"`

	// When creating, will be calculated if dynamic, required otherwise
	Name string `json:"name,omitempty"`

	// Used for dynamic environment name generation only, to override using the owner email handle and template name
	NamePrefix string `json:"namePrefix,omitempty"`

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

	// When creating, will be set to your email
	Owner string `json:"owner,omitempty"`

	// pagerduty integration
	PagerdutyIntegration string `json:"pagerdutyIntegration,omitempty"`

	// pagerduty integration info
	PagerdutyIntegrationInfo *V2controllersPagerdutyIntegration `json:"pagerdutyIntegrationInfo,omitempty"`

	// Used to protect specific BEEs from deletion (thelma checks this field)
	PreventDeletion *bool `json:"preventDeletion,omitempty"`

	// requires suitability
	RequiresSuitability *bool `json:"requiresSuitability,omitempty"`

	// Required for dynamic environments
	TemplateEnvironment string `json:"templateEnvironment,omitempty"`

	// Single-layer recursive; provides info of the template environment if this environment has one
	TemplateEnvironmentInfo interface{} `json:"templateEnvironmentInfo,omitempty"`

	// When creating, will be calculated if left empty
	UniqueResourcePrefix string `json:"uniqueResourcePrefix,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// When creating, defaults to template name or environment name
	ValuesName string `json:"valuesName,omitempty"`
}

// Validate validates this v2controllers environment
func (m *V2controllersEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineScheduleBeginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineScheduleEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerdutyIntegrationInfo(formats); err != nil {
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

func (m *V2controllersEnvironment) validateAutoDelete(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoDelete) { // not required
		return nil
	}

	if m.AutoDelete != nil {
		if err := m.AutoDelete.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDelete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoDelete")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersEnvironment) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2controllersEnvironment) validateDefaultClusterInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultClusterInfo) { // not required
		return nil
	}

	if m.DefaultClusterInfo != nil {
		if err := m.DefaultClusterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultClusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultClusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersEnvironment) validateOfflineScheduleBeginTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineScheduleBeginTime) { // not required
		return nil
	}

	if err := validate.FormatOf("offlineScheduleBeginTime", "body", "date-time", m.OfflineScheduleBeginTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2controllersEnvironment) validateOfflineScheduleEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineScheduleEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("offlineScheduleEndTime", "body", "date-time", m.OfflineScheduleEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2controllersEnvironment) validatePagerdutyIntegrationInfo(formats strfmt.Registry) error {
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

func (m *V2controllersEnvironment) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v2controllers environment based on the context it is used
func (m *V2controllersEnvironment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoDelete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultClusterInfo(ctx, formats); err != nil {
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

func (m *V2controllersEnvironment) contextValidateAutoDelete(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoDelete != nil {
		if err := m.AutoDelete.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDelete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoDelete")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersEnvironment) contextValidateDefaultClusterInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultClusterInfo != nil {
		if err := m.DefaultClusterInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultClusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultClusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V2controllersEnvironment) contextValidatePagerdutyIntegrationInfo(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V2controllersEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2controllersEnvironment) UnmarshalBinary(b []byte) error {
	var res V2controllersEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

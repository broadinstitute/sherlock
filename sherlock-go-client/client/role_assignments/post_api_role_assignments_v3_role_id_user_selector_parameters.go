// Code generated by go-swagger; DO NOT EDIT.

package role_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
)

// NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParams creates a new PostAPIRoleAssignmentsV3RoleIDUserSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParams() *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	return &PostAPIRoleAssignmentsV3RoleIDUserSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParamsWithTimeout creates a new PostAPIRoleAssignmentsV3RoleIDUserSelectorParams object
// with the ability to set a timeout on a request.
func NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParamsWithTimeout(timeout time.Duration) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	return &PostAPIRoleAssignmentsV3RoleIDUserSelectorParams{
		timeout: timeout,
	}
}

// NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParamsWithContext creates a new PostAPIRoleAssignmentsV3RoleIDUserSelectorParams object
// with the ability to set a context for a request.
func NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParamsWithContext(ctx context.Context) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	return &PostAPIRoleAssignmentsV3RoleIDUserSelectorParams{
		Context: ctx,
	}
}

// NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParamsWithHTTPClient creates a new PostAPIRoleAssignmentsV3RoleIDUserSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParamsWithHTTPClient(client *http.Client) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	return &PostAPIRoleAssignmentsV3RoleIDUserSelectorParams{
		HTTPClient: client,
	}
}

/* PostAPIRoleAssignmentsV3RoleIDUserSelectorParams contains all the parameters to send to the API endpoint
   for the post API role assignments v3 role ID user selector operation.

   Typically these are written to a http.Request.
*/
type PostAPIRoleAssignmentsV3RoleIDUserSelectorParams struct {

	/* RoleAssignment.

	   The initial fields to set for the new RoleAssignment
	*/
	RoleAssignment *models.SherlockRoleAssignmentV3Edit

	/* RoleID.

	   The numeric ID of the role
	*/
	RoleID int64

	/* UserSelector.

	   The selector of the User, which can be either a numeric ID, the email, 'google-id/{google subject ID}', 'github/{github username}', or 'github-id/{github numeric ID}'.
	*/
	UserSelector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API role assignments v3 role ID user selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithDefaults() *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API role assignments v3 role ID user selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithTimeout(timeout time.Duration) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithContext(ctx context.Context) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithHTTPClient(client *http.Client) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleAssignment adds the roleAssignment to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithRoleAssignment(roleAssignment *models.SherlockRoleAssignmentV3Edit) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetRoleAssignment(roleAssignment)
	return o
}

// SetRoleAssignment adds the roleAssignment to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetRoleAssignment(roleAssignment *models.SherlockRoleAssignmentV3Edit) {
	o.RoleAssignment = roleAssignment
}

// WithRoleID adds the roleID to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithRoleID(roleID int64) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WithUserSelector adds the userSelector to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WithUserSelector(userSelector string) *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams {
	o.SetUserSelector(userSelector)
	return o
}

// SetUserSelector adds the userSelector to the post API role assignments v3 role ID user selector params
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) SetUserSelector(userSelector string) {
	o.UserSelector = userSelector
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RoleAssignment != nil {
		if err := r.SetBodyParam(o.RoleAssignment); err != nil {
			return err
		}
	}

	// path param role-id
	if err := r.SetPathParam("role-id", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	// path param user-selector
	if err := r.SetPathParam("user-selector", o.UserSelector); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
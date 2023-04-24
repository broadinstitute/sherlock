// Code generated by go-swagger; DO NOT EDIT.

package app_versions

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
)

// NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParams creates a new GetAPIV2ProceduresAppVersionsChildrenPathToParentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParams() *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresAppVersionsChildrenPathToParentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParamsWithTimeout creates a new GetAPIV2ProceduresAppVersionsChildrenPathToParentParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParamsWithTimeout(timeout time.Duration) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresAppVersionsChildrenPathToParentParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParamsWithContext creates a new GetAPIV2ProceduresAppVersionsChildrenPathToParentParams object
// with the ability to set a context for a request.
func NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParamsWithContext(ctx context.Context) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresAppVersionsChildrenPathToParentParams{
		Context: ctx,
	}
}

// NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParamsWithHTTPClient creates a new GetAPIV2ProceduresAppVersionsChildrenPathToParentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ProceduresAppVersionsChildrenPathToParentParamsWithHTTPClient(client *http.Client) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresAppVersionsChildrenPathToParentParams{
		HTTPClient: client,
	}
}

/* GetAPIV2ProceduresAppVersionsChildrenPathToParentParams contains all the parameters to send to the API endpoint
   for the get API v2 procedures app versions children path to parent operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2ProceduresAppVersionsChildrenPathToParentParams struct {

	/* Child.

	   The selector of the newer AppVersion for the changelog
	*/
	Child string

	/* Parent.

	   The selector of the older AppVersion for the changelog
	*/
	Parent string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 procedures app versions children path to parent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WithDefaults() *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 procedures app versions children path to parent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WithTimeout(timeout time.Duration) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WithContext(ctx context.Context) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WithHTTPClient(client *http.Client) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChild adds the child to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WithChild(child string) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	o.SetChild(child)
	return o
}

// SetChild adds the child to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) SetChild(child string) {
	o.Child = child
}

// WithParent adds the parent to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WithParent(parent string) *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the get API v2 procedures app versions children path to parent params
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) SetParent(parent string) {
	o.Parent = parent
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ProceduresAppVersionsChildrenPathToParentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param child
	qrChild := o.Child
	qChild := qrChild
	if qChild != "" {

		if err := r.SetQueryParam("child", qChild); err != nil {
			return err
		}
	}

	// query param parent
	qrParent := o.Parent
	qParent := qrParent
	if qParent != "" {

		if err := r.SetQueryParam("parent", qParent); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
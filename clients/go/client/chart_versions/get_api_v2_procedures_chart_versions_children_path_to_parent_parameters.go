// Code generated by go-swagger; DO NOT EDIT.

package chart_versions

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

// NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParams creates a new GetAPIV2ProceduresChartVersionsChildrenPathToParentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParams() *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresChartVersionsChildrenPathToParentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParamsWithTimeout creates a new GetAPIV2ProceduresChartVersionsChildrenPathToParentParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParamsWithTimeout(timeout time.Duration) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresChartVersionsChildrenPathToParentParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParamsWithContext creates a new GetAPIV2ProceduresChartVersionsChildrenPathToParentParams object
// with the ability to set a context for a request.
func NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParamsWithContext(ctx context.Context) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresChartVersionsChildrenPathToParentParams{
		Context: ctx,
	}
}

// NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParamsWithHTTPClient creates a new GetAPIV2ProceduresChartVersionsChildrenPathToParentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ProceduresChartVersionsChildrenPathToParentParamsWithHTTPClient(client *http.Client) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	return &GetAPIV2ProceduresChartVersionsChildrenPathToParentParams{
		HTTPClient: client,
	}
}

/* GetAPIV2ProceduresChartVersionsChildrenPathToParentParams contains all the parameters to send to the API endpoint
   for the get API v2 procedures chart versions children path to parent operation.

   Typically these are written to a http.Request.
*/
type GetAPIV2ProceduresChartVersionsChildrenPathToParentParams struct {

	/* Child.

	   The selector of the newer ChartVersion for the changelog
	*/
	Child string

	/* Parent.

	   The selector of the older ChartVersion for the changelog
	*/
	Parent string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 procedures chart versions children path to parent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WithDefaults() *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 procedures chart versions children path to parent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WithTimeout(timeout time.Duration) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WithContext(ctx context.Context) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WithHTTPClient(client *http.Client) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChild adds the child to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WithChild(child string) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	o.SetChild(child)
	return o
}

// SetChild adds the child to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) SetChild(child string) {
	o.Child = child
}

// WithParent adds the parent to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WithParent(parent string) *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the get API v2 procedures chart versions children path to parent params
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) SetParent(parent string) {
	o.Parent = parent
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ProceduresChartVersionsChildrenPathToParentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
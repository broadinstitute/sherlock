// Code generated by go-swagger; DO NOT EDIT.

package ci_runs

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

// NewGetAPICiRunsProceduresV3GithubInfoParams creates a new GetAPICiRunsProceduresV3GithubInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPICiRunsProceduresV3GithubInfoParams() *GetAPICiRunsProceduresV3GithubInfoParams {
	return &GetAPICiRunsProceduresV3GithubInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPICiRunsProceduresV3GithubInfoParamsWithTimeout creates a new GetAPICiRunsProceduresV3GithubInfoParams object
// with the ability to set a timeout on a request.
func NewGetAPICiRunsProceduresV3GithubInfoParamsWithTimeout(timeout time.Duration) *GetAPICiRunsProceduresV3GithubInfoParams {
	return &GetAPICiRunsProceduresV3GithubInfoParams{
		timeout: timeout,
	}
}

// NewGetAPICiRunsProceduresV3GithubInfoParamsWithContext creates a new GetAPICiRunsProceduresV3GithubInfoParams object
// with the ability to set a context for a request.
func NewGetAPICiRunsProceduresV3GithubInfoParamsWithContext(ctx context.Context) *GetAPICiRunsProceduresV3GithubInfoParams {
	return &GetAPICiRunsProceduresV3GithubInfoParams{
		Context: ctx,
	}
}

// NewGetAPICiRunsProceduresV3GithubInfoParamsWithHTTPClient creates a new GetAPICiRunsProceduresV3GithubInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPICiRunsProceduresV3GithubInfoParamsWithHTTPClient(client *http.Client) *GetAPICiRunsProceduresV3GithubInfoParams {
	return &GetAPICiRunsProceduresV3GithubInfoParams{
		HTTPClient: client,
	}
}

/* GetAPICiRunsProceduresV3GithubInfoParams contains all the parameters to send to the API endpoint
   for the get API ci runs procedures v3 github info operation.

   Typically these are written to a http.Request.
*/
type GetAPICiRunsProceduresV3GithubInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API ci runs procedures v3 github info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICiRunsProceduresV3GithubInfoParams) WithDefaults() *GetAPICiRunsProceduresV3GithubInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API ci runs procedures v3 github info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICiRunsProceduresV3GithubInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API ci runs procedures v3 github info params
func (o *GetAPICiRunsProceduresV3GithubInfoParams) WithTimeout(timeout time.Duration) *GetAPICiRunsProceduresV3GithubInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API ci runs procedures v3 github info params
func (o *GetAPICiRunsProceduresV3GithubInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API ci runs procedures v3 github info params
func (o *GetAPICiRunsProceduresV3GithubInfoParams) WithContext(ctx context.Context) *GetAPICiRunsProceduresV3GithubInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API ci runs procedures v3 github info params
func (o *GetAPICiRunsProceduresV3GithubInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API ci runs procedures v3 github info params
func (o *GetAPICiRunsProceduresV3GithubInfoParams) WithHTTPClient(client *http.Client) *GetAPICiRunsProceduresV3GithubInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API ci runs procedures v3 github info params
func (o *GetAPICiRunsProceduresV3GithubInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPICiRunsProceduresV3GithubInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

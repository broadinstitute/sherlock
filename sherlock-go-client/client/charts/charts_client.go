// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new charts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for charts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIChartsV3Selector(params *DeleteAPIChartsV3SelectorParams, opts ...ClientOption) (*DeleteAPIChartsV3SelectorOK, error)

	GetAPIChartsV3(params *GetAPIChartsV3Params, opts ...ClientOption) (*GetAPIChartsV3OK, error)

	GetAPIChartsV3Selector(params *GetAPIChartsV3SelectorParams, opts ...ClientOption) (*GetAPIChartsV3SelectorOK, error)

	PatchAPIChartsV3Selector(params *PatchAPIChartsV3SelectorParams, opts ...ClientOption) (*PatchAPIChartsV3SelectorOK, error)

	PostAPIChartsV3(params *PostAPIChartsV3Params, opts ...ClientOption) (*PostAPIChartsV3Created, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIChartsV3Selector deletes an individual chart

  Delete an individual Chart by its ID.
*/
func (a *Client) DeleteAPIChartsV3Selector(params *DeleteAPIChartsV3SelectorParams, opts ...ClientOption) (*DeleteAPIChartsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIChartsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIChartsV3Selector",
		Method:             "DELETE",
		PathPattern:        "/api/charts/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIChartsV3SelectorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAPIChartsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIChartsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIChartsV3 lists charts matching a filter

  List Charts matching a filter.
*/
func (a *Client) GetAPIChartsV3(params *GetAPIChartsV3Params, opts ...ClientOption) (*GetAPIChartsV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIChartsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIChartsV3",
		Method:             "GET",
		PathPattern:        "/api/charts/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIChartsV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIChartsV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIChartsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIChartsV3Selector gets an individual chart

  Get an individual Chart.
*/
func (a *Client) GetAPIChartsV3Selector(params *GetAPIChartsV3SelectorParams, opts ...ClientOption) (*GetAPIChartsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIChartsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIChartsV3Selector",
		Method:             "GET",
		PathPattern:        "/api/charts/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIChartsV3SelectorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIChartsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIChartsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIChartsV3Selector edits an individual chart

  Edit an individual Chart.
*/
func (a *Client) PatchAPIChartsV3Selector(params *PatchAPIChartsV3SelectorParams, opts ...ClientOption) (*PatchAPIChartsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIChartsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIChartsV3Selector",
		Method:             "PATCH",
		PathPattern:        "/api/charts/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIChartsV3SelectorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPIChartsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIChartsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIChartsV3 creates a chart

  Create a Chart.
*/
func (a *Client) PostAPIChartsV3(params *PostAPIChartsV3Params, opts ...ClientOption) (*PostAPIChartsV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIChartsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIChartsV3",
		Method:             "POST",
		PathPattern:        "/api/charts/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIChartsV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPIChartsV3Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIChartsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

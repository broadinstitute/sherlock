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
	DeleteAPIV2ChartsSelector(params *DeleteAPIV2ChartsSelectorParams, opts ...ClientOption) (*DeleteAPIV2ChartsSelectorOK, error)

	GetAPIV2Charts(params *GetAPIV2ChartsParams, opts ...ClientOption) (*GetAPIV2ChartsOK, error)

	GetAPIV2ChartsSelector(params *GetAPIV2ChartsSelectorParams, opts ...ClientOption) (*GetAPIV2ChartsSelectorOK, error)

	GetAPIV2SelectorsChartsSelector(params *GetAPIV2SelectorsChartsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsChartsSelectorOK, error)

	PatchAPIV2ChartsSelector(params *PatchAPIV2ChartsSelectorParams, opts ...ClientOption) (*PatchAPIV2ChartsSelectorOK, error)

	PostAPIV2Charts(params *PostAPIV2ChartsParams, opts ...ClientOption) (*PostAPIV2ChartsOK, *PostAPIV2ChartsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIV2ChartsSelector deletes a chart entry

  Delete an existing Chart entry via one of its "selectors": name or numeric ID.
*/
func (a *Client) DeleteAPIV2ChartsSelector(params *DeleteAPIV2ChartsSelectorParams, opts ...ClientOption) (*DeleteAPIV2ChartsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV2ChartsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIV2ChartsSelector",
		Method:             "DELETE",
		PathPattern:        "/api/v2/charts/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIV2ChartsSelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIV2ChartsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIV2ChartsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2Charts lists chart entries

  List existing Chart entries, ordered by most recently updated.
*/
func (a *Client) GetAPIV2Charts(params *GetAPIV2ChartsParams, opts ...ClientOption) (*GetAPIV2ChartsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ChartsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Charts",
		Method:             "GET",
		PathPattern:        "/api/v2/charts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2ChartsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ChartsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Charts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2ChartsSelector gets a chart entry

  Get an existing Chart entry via one of its "selectors": name or numeric ID.
*/
func (a *Client) GetAPIV2ChartsSelector(params *GetAPIV2ChartsSelectorParams, opts ...ClientOption) (*GetAPIV2ChartsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ChartsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2ChartsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/charts/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2ChartsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ChartsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2ChartsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2SelectorsChartsSelector lists chart selectors

  Validate a given Chart selector and provide any other selectors that would match the same Chart.
*/
func (a *Client) GetAPIV2SelectorsChartsSelector(params *GetAPIV2SelectorsChartsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsChartsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SelectorsChartsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SelectorsChartsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/selectors/charts/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2SelectorsChartsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2SelectorsChartsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SelectorsChartsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIV2ChartsSelector edits a chart entry

  Edit an existing Chart entry via one of its "selectors": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
*/
func (a *Client) PatchAPIV2ChartsSelector(params *PatchAPIV2ChartsSelectorParams, opts ...ClientOption) (*PatchAPIV2ChartsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2ChartsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2ChartsSelector",
		Method:             "PATCH",
		PathPattern:        "/api/v2/charts/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIV2ChartsSelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2ChartsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2ChartsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIV2Charts creates a new chart entry

  Create a new Chart entry. Note that some fields are immutable after creation; /edit lists mutable fields.
*/
func (a *Client) PostAPIV2Charts(params *PostAPIV2ChartsParams, opts ...ClientOption) (*PostAPIV2ChartsOK, *PostAPIV2ChartsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2ChartsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2Charts",
		Method:             "POST",
		PathPattern:        "/api/v2/charts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2ChartsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostAPIV2ChartsOK:
		return value, nil, nil
	case *PostAPIV2ChartsCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for charts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package chart_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new chart versions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for chart versions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV2ChartVersions(params *GetAPIV2ChartVersionsParams, opts ...ClientOption) (*GetAPIV2ChartVersionsOK, error)

	GetAPIV2ChartVersionsSelector(params *GetAPIV2ChartVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2ChartVersionsSelectorOK, error)

	GetAPIV2SelectorsChartVersionsSelector(params *GetAPIV2SelectorsChartVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsChartVersionsSelectorOK, error)

	PatchAPIV2ChartVersionsSelector(params *PatchAPIV2ChartVersionsSelectorParams, opts ...ClientOption) (*PatchAPIV2ChartVersionsSelectorOK, error)

	PostAPIV2ChartVersions(params *PostAPIV2ChartVersionsParams, opts ...ClientOption) (*PostAPIV2ChartVersionsOK, *PostAPIV2ChartVersionsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPIV2ChartVersions lists chart version entries

  List existing ChartVersion entries, ordered by most recently updated.
*/
func (a *Client) GetAPIV2ChartVersions(params *GetAPIV2ChartVersionsParams, opts ...ClientOption) (*GetAPIV2ChartVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ChartVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2ChartVersions",
		Method:             "GET",
		PathPattern:        "/api/v2/chart-versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2ChartVersionsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ChartVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2ChartVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2ChartVersionsSelector gets a chart version entry

  Get an existing ChartVersion entry via one its "selectors": chart/version or numeric ID.
*/
func (a *Client) GetAPIV2ChartVersionsSelector(params *GetAPIV2ChartVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2ChartVersionsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ChartVersionsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2ChartVersionsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/chart-versions/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2ChartVersionsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ChartVersionsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2ChartVersionsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2SelectorsChartVersionsSelector lists chart version selectors

  Validate a given ChartVersion selector and provide any other selectors that would match the same ChartVersion.
*/
func (a *Client) GetAPIV2SelectorsChartVersionsSelector(params *GetAPIV2SelectorsChartVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsChartVersionsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SelectorsChartVersionsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SelectorsChartVersionsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/selectors/chart-versions/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2SelectorsChartVersionsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2SelectorsChartVersionsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SelectorsChartVersionsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIV2ChartVersionsSelector edits a chart version entry

  Edit an existing ChartVersion entry via one its "selectors": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
*/
func (a *Client) PatchAPIV2ChartVersionsSelector(params *PatchAPIV2ChartVersionsSelectorParams, opts ...ClientOption) (*PatchAPIV2ChartVersionsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2ChartVersionsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2ChartVersionsSelector",
		Method:             "PATCH",
		PathPattern:        "/api/v2/chart-versions/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIV2ChartVersionsSelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2ChartVersionsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2ChartVersionsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIV2ChartVersions creates a new chart version entry

  Create a new ChartVersion entry. Note that fields are immutable after creation.
If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
*/
func (a *Client) PostAPIV2ChartVersions(params *PostAPIV2ChartVersionsParams, opts ...ClientOption) (*PostAPIV2ChartVersionsOK, *PostAPIV2ChartVersionsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2ChartVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2ChartVersions",
		Method:             "POST",
		PathPattern:        "/api/v2/chart-versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2ChartVersionsReader{formats: a.formats},
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
	case *PostAPIV2ChartVersionsOK:
		return value, nil, nil
	case *PostAPIV2ChartVersionsCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for chart_versions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

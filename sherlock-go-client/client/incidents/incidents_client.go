// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new incidents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for incidents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIIncidentsV3Selector(params *DeleteAPIIncidentsV3SelectorParams, opts ...ClientOption) (*DeleteAPIIncidentsV3SelectorOK, error)

	GetAPIIncidentsV3(params *GetAPIIncidentsV3Params, opts ...ClientOption) (*GetAPIIncidentsV3OK, error)

	GetAPIIncidentsV3Selector(params *GetAPIIncidentsV3SelectorParams, opts ...ClientOption) (*GetAPIIncidentsV3SelectorOK, error)

	PatchAPIIncidentsV3Selector(params *PatchAPIIncidentsV3SelectorParams, opts ...ClientOption) (*PatchAPIIncidentsV3SelectorOK, error)

	PostAPIIncidentsV3(params *PostAPIIncidentsV3Params, opts ...ClientOption) (*PostAPIIncidentsV3Created, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIIncidentsV3Selector deletes an individual incident

  Delete an individual Incident by its ID.
*/
func (a *Client) DeleteAPIIncidentsV3Selector(params *DeleteAPIIncidentsV3SelectorParams, opts ...ClientOption) (*DeleteAPIIncidentsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIIncidentsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIIncidentsV3Selector",
		Method:             "DELETE",
		PathPattern:        "/api/incidents/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIIncidentsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIIncidentsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIIncidentsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIIncidentsV3 lists incidents matching a filter

  List Incidents matching a filter.
*/
func (a *Client) GetAPIIncidentsV3(params *GetAPIIncidentsV3Params, opts ...ClientOption) (*GetAPIIncidentsV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIIncidentsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIIncidentsV3",
		Method:             "GET",
		PathPattern:        "/api/incidents/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIIncidentsV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIIncidentsV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIIncidentsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIIncidentsV3Selector gets an individual incident

  Get an individual Incident.
*/
func (a *Client) GetAPIIncidentsV3Selector(params *GetAPIIncidentsV3SelectorParams, opts ...ClientOption) (*GetAPIIncidentsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIIncidentsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIIncidentsV3Selector",
		Method:             "GET",
		PathPattern:        "/api/incidents/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIIncidentsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIIncidentsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIIncidentsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIIncidentsV3Selector edits an individual incident

  Edit an individual Incident.
*/
func (a *Client) PatchAPIIncidentsV3Selector(params *PatchAPIIncidentsV3SelectorParams, opts ...ClientOption) (*PatchAPIIncidentsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIIncidentsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIIncidentsV3Selector",
		Method:             "PATCH",
		PathPattern:        "/api/incidents/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIIncidentsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIIncidentsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIIncidentsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIIncidentsV3 creates a incident

  Create a Incident.
*/
func (a *Client) PostAPIIncidentsV3(params *PostAPIIncidentsV3Params, opts ...ClientOption) (*PostAPIIncidentsV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIIncidentsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIIncidentsV3",
		Method:             "POST",
		PathPattern:        "/api/incidents/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIIncidentsV3Reader{formats: a.formats},
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
	success, ok := result.(*PostAPIIncidentsV3Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIIncidentsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
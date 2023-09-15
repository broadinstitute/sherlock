// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIUsersV3(params *GetAPIUsersV3Params, opts ...ClientOption) (*GetAPIUsersV3OK, error)

	GetAPIUsersV3Selector(params *GetAPIUsersV3SelectorParams, opts ...ClientOption) (*GetAPIUsersV3SelectorOK, error)

	PutAPIUsersV3(params *PutAPIUsersV3Params, opts ...ClientOption) (*PutAPIUsersV3OK, *PutAPIUsersV3Created, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPIUsersV3 lists users matching a filter

  List Users matching a filter. The results will include suitability and other information.
Note that the suitability info can't directly be filtered for at this time.
*/
func (a *Client) GetAPIUsersV3(params *GetAPIUsersV3Params, opts ...ClientOption) (*GetAPIUsersV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIUsersV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIUsersV3",
		Method:             "GET",
		PathPattern:        "/api/users/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIUsersV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIUsersV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIUsersV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIUsersV3Selector gets an individual user

  Get an individual User. As a special case, "me" or "self" can be passed as the selector to get the current user.
*/
func (a *Client) GetAPIUsersV3Selector(params *GetAPIUsersV3SelectorParams, opts ...ClientOption) (*GetAPIUsersV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIUsersV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIUsersV3Selector",
		Method:             "GET",
		PathPattern:        "/api/users/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIUsersV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIUsersV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIUsersV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAPIUsersV3 updates the calling user s information

  Update the calling User's information. As with all authenticated Sherlock endpoints,
newly-observed callers will have a User record added, meaning that this endpoint
behaves like an upsert.
*/
func (a *Client) PutAPIUsersV3(params *PutAPIUsersV3Params, opts ...ClientOption) (*PutAPIUsersV3OK, *PutAPIUsersV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIUsersV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutAPIUsersV3",
		Method:             "PUT",
		PathPattern:        "/api/users/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAPIUsersV3Reader{formats: a.formats},
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
	case *PutAPIUsersV3OK:
		return value, nil, nil
	case *PutAPIUsersV3Created:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

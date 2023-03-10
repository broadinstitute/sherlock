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
	PostAPIV2ProceduresUsersLinkGithub(params *PostAPIV2ProceduresUsersLinkGithubParams, opts ...ClientOption) (*PostAPIV2ProceduresUsersLinkGithubOK, *PostAPIV2ProceduresUsersLinkGithubAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PostAPIV2ProceduresUsersLinkGithub updates the user s git hub account link

  Update the authenticated User's associated personal GitHub account
*/
func (a *Client) PostAPIV2ProceduresUsersLinkGithub(params *PostAPIV2ProceduresUsersLinkGithubParams, opts ...ClientOption) (*PostAPIV2ProceduresUsersLinkGithubOK, *PostAPIV2ProceduresUsersLinkGithubAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2ProceduresUsersLinkGithubParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2ProceduresUsersLinkGithub",
		Method:             "POST",
		PathPattern:        "/api/v2/procedures/users/link-github",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2ProceduresUsersLinkGithubReader{formats: a.formats},
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
	case *PostAPIV2ProceduresUsersLinkGithubOK:
		return value, nil, nil
	case *PostAPIV2ProceduresUsersLinkGithubAccepted:
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

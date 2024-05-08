// Code generated by go-swagger; DO NOT EDIT.

package role_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new role assignments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for role assignments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIRoleAssignmentsV3RoleIDUserSelector(params *DeleteAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*DeleteAPIRoleAssignmentsV3RoleIDUserSelectorOK, error)

	GetAPIRoleAssignmentsV3(params *GetAPIRoleAssignmentsV3Params, opts ...ClientOption) (*GetAPIRoleAssignmentsV3OK, error)

	GetAPIRoleAssignmentsV3RoleIDUserSelector(params *GetAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*GetAPIRoleAssignmentsV3RoleIDUserSelectorOK, error)

	PatchAPIRoleAssignmentsV3RoleIDUserSelector(params *PatchAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*PatchAPIRoleAssignmentsV3RoleIDUserSelectorOK, error)

	PostAPIRoleAssignmentsV3RoleIDUserSelector(params *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*PostAPIRoleAssignmentsV3RoleIDUserSelectorCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIRoleAssignmentsV3RoleIDUserSelector deletes a role assignment

  Delete the RoleAssignment between a given Role and User.
Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
*/
func (a *Client) DeleteAPIRoleAssignmentsV3RoleIDUserSelector(params *DeleteAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*DeleteAPIRoleAssignmentsV3RoleIDUserSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIRoleAssignmentsV3RoleIDUserSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIRoleAssignmentsV3RoleIDUserSelector",
		Method:             "DELETE",
		PathPattern:        "/api/role-assignments/v3/{role-id}/{user-selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIRoleAssignmentsV3RoleIDUserSelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIRoleAssignmentsV3RoleIDUserSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIRoleAssignmentsV3RoleIDUserSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIRoleAssignmentsV3 lists role assignments matching a filter

  List RoleAssignments matching a filter. The correct way to list RoleAssignments for a particular Role or User is to get that Role or User specifically, not to use this endpoint.
*/
func (a *Client) GetAPIRoleAssignmentsV3(params *GetAPIRoleAssignmentsV3Params, opts ...ClientOption) (*GetAPIRoleAssignmentsV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIRoleAssignmentsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIRoleAssignmentsV3",
		Method:             "GET",
		PathPattern:        "/api/role-assignments/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIRoleAssignmentsV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIRoleAssignmentsV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIRoleAssignmentsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIRoleAssignmentsV3RoleIDUserSelector gets a role assignment

  Get the RoleAssignment between a given Role and User.
*/
func (a *Client) GetAPIRoleAssignmentsV3RoleIDUserSelector(params *GetAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*GetAPIRoleAssignmentsV3RoleIDUserSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIRoleAssignmentsV3RoleIDUserSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIRoleAssignmentsV3RoleIDUserSelector",
		Method:             "GET",
		PathPattern:        "/api/role-assignments/v3/{role-id}/{user-selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIRoleAssignmentsV3RoleIDUserSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIRoleAssignmentsV3RoleIDUserSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIRoleAssignmentsV3RoleIDUserSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIRoleAssignmentsV3RoleIDUserSelector edits a role assignment

  Edit the RoleAssignment between a given Role and User.
Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
*/
func (a *Client) PatchAPIRoleAssignmentsV3RoleIDUserSelector(params *PatchAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*PatchAPIRoleAssignmentsV3RoleIDUserSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIRoleAssignmentsV3RoleIDUserSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIRoleAssignmentsV3RoleIDUserSelector",
		Method:             "PATCH",
		PathPattern:        "/api/role-assignments/v3/{role-id}/{user-selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIRoleAssignmentsV3RoleIDUserSelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIRoleAssignmentsV3RoleIDUserSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIRoleAssignmentsV3RoleIDUserSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIRoleAssignmentsV3RoleIDUserSelector creates a role assignment

  Create the RoleAssignment between a given Role and User.
Non-super-admins may only mutate RoleAssignments for themselves, only for roles they can break-glass into, and only with an expiry no further than the role's default break-glass duration in the future.
*/
func (a *Client) PostAPIRoleAssignmentsV3RoleIDUserSelector(params *PostAPIRoleAssignmentsV3RoleIDUserSelectorParams, opts ...ClientOption) (*PostAPIRoleAssignmentsV3RoleIDUserSelectorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIRoleAssignmentsV3RoleIDUserSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIRoleAssignmentsV3RoleIDUserSelector",
		Method:             "POST",
		PathPattern:        "/api/role-assignments/v3/{role-id}/{user-selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIRoleAssignmentsV3RoleIDUserSelectorReader{formats: a.formats},
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
	success, ok := result.(*PostAPIRoleAssignmentsV3RoleIDUserSelectorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIRoleAssignmentsV3RoleIDUserSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

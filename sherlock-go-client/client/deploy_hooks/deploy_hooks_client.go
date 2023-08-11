// Code generated by go-swagger; DO NOT EDIT.

package deploy_hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deploy hooks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deploy hooks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIDeployHooksGithubActionsV3Selector(params *DeleteAPIDeployHooksGithubActionsV3SelectorParams, opts ...ClientOption) (*DeleteAPIDeployHooksGithubActionsV3SelectorOK, error)

	DeleteAPIDeployHooksSlackV3Selector(params *DeleteAPIDeployHooksSlackV3SelectorParams, opts ...ClientOption) (*DeleteAPIDeployHooksSlackV3SelectorOK, error)

	GetAPIDeployHooksGithubActionsV3(params *GetAPIDeployHooksGithubActionsV3Params, opts ...ClientOption) (*GetAPIDeployHooksGithubActionsV3OK, error)

	GetAPIDeployHooksGithubActionsV3Selector(params *GetAPIDeployHooksGithubActionsV3SelectorParams, opts ...ClientOption) (*GetAPIDeployHooksGithubActionsV3SelectorOK, error)

	GetAPIDeployHooksSlackV3(params *GetAPIDeployHooksSlackV3Params, opts ...ClientOption) (*GetAPIDeployHooksSlackV3OK, error)

	GetAPIDeployHooksSlackV3Selector(params *GetAPIDeployHooksSlackV3SelectorParams, opts ...ClientOption) (*GetAPIDeployHooksSlackV3SelectorOK, error)

	PatchAPIDeployHooksGithubActionsV3Selector(params *PatchAPIDeployHooksGithubActionsV3SelectorParams, opts ...ClientOption) (*PatchAPIDeployHooksGithubActionsV3SelectorOK, error)

	PatchAPIDeployHooksSlackV3Selector(params *PatchAPIDeployHooksSlackV3SelectorParams, opts ...ClientOption) (*PatchAPIDeployHooksSlackV3SelectorOK, error)

	PostAPIDeployHooksGithubActionsV3(params *PostAPIDeployHooksGithubActionsV3Params, opts ...ClientOption) (*PostAPIDeployHooksGithubActionsV3Created, error)

	PostAPIDeployHooksSlackV3(params *PostAPIDeployHooksSlackV3Params, opts ...ClientOption) (*PostAPIDeployHooksSlackV3Created, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIDeployHooksGithubActionsV3Selector deletes an individual github actions deploy hook

  Delete an individual GithubActionsDeployHook by its ID.
*/
func (a *Client) DeleteAPIDeployHooksGithubActionsV3Selector(params *DeleteAPIDeployHooksGithubActionsV3SelectorParams, opts ...ClientOption) (*DeleteAPIDeployHooksGithubActionsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIDeployHooksGithubActionsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIDeployHooksGithubActionsV3Selector",
		Method:             "DELETE",
		PathPattern:        "/api/deploy-hooks/github-actions/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIDeployHooksGithubActionsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIDeployHooksGithubActionsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIDeployHooksGithubActionsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAPIDeployHooksSlackV3Selector deletes an individual slack deploy hook

  Delete an individual SlackDeployHook by its ID.
*/
func (a *Client) DeleteAPIDeployHooksSlackV3Selector(params *DeleteAPIDeployHooksSlackV3SelectorParams, opts ...ClientOption) (*DeleteAPIDeployHooksSlackV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIDeployHooksSlackV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIDeployHooksSlackV3Selector",
		Method:             "DELETE",
		PathPattern:        "/api/deploy-hooks/slack/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIDeployHooksSlackV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIDeployHooksSlackV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIDeployHooksSlackV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIDeployHooksGithubActionsV3 lists github actions deploy hooks matching a filter

  List GithubActionsDeployHooks matching a filter.
*/
func (a *Client) GetAPIDeployHooksGithubActionsV3(params *GetAPIDeployHooksGithubActionsV3Params, opts ...ClientOption) (*GetAPIDeployHooksGithubActionsV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIDeployHooksGithubActionsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIDeployHooksGithubActionsV3",
		Method:             "GET",
		PathPattern:        "/api/deploy-hooks/github-actions/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIDeployHooksGithubActionsV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIDeployHooksGithubActionsV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIDeployHooksGithubActionsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIDeployHooksGithubActionsV3Selector gets an individual github actions deploy hook

  Get an individual GithubActionsDeployHook by its ID.
*/
func (a *Client) GetAPIDeployHooksGithubActionsV3Selector(params *GetAPIDeployHooksGithubActionsV3SelectorParams, opts ...ClientOption) (*GetAPIDeployHooksGithubActionsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIDeployHooksGithubActionsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIDeployHooksGithubActionsV3Selector",
		Method:             "GET",
		PathPattern:        "/api/deploy-hooks/github-actions/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIDeployHooksGithubActionsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIDeployHooksGithubActionsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIDeployHooksGithubActionsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIDeployHooksSlackV3 lists slack deploy hooks matching a filter

  List SlackDeployHooks matching a filter.
*/
func (a *Client) GetAPIDeployHooksSlackV3(params *GetAPIDeployHooksSlackV3Params, opts ...ClientOption) (*GetAPIDeployHooksSlackV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIDeployHooksSlackV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIDeployHooksSlackV3",
		Method:             "GET",
		PathPattern:        "/api/deploy-hooks/slack/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIDeployHooksSlackV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIDeployHooksSlackV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIDeployHooksSlackV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIDeployHooksSlackV3Selector gets an individual slack deploy hook

  Get an individual SlackDeployHook by its ID.
*/
func (a *Client) GetAPIDeployHooksSlackV3Selector(params *GetAPIDeployHooksSlackV3SelectorParams, opts ...ClientOption) (*GetAPIDeployHooksSlackV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIDeployHooksSlackV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIDeployHooksSlackV3Selector",
		Method:             "GET",
		PathPattern:        "/api/deploy-hooks/slack/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIDeployHooksSlackV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIDeployHooksSlackV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIDeployHooksSlackV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIDeployHooksGithubActionsV3Selector edits an individual github actions deploy hook

  Edit an individual GithubActionsDeployHook by its ID.
*/
func (a *Client) PatchAPIDeployHooksGithubActionsV3Selector(params *PatchAPIDeployHooksGithubActionsV3SelectorParams, opts ...ClientOption) (*PatchAPIDeployHooksGithubActionsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIDeployHooksGithubActionsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIDeployHooksGithubActionsV3Selector",
		Method:             "PATCH",
		PathPattern:        "/api/deploy-hooks/github-actions/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIDeployHooksGithubActionsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIDeployHooksGithubActionsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIDeployHooksGithubActionsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIDeployHooksSlackV3Selector edits an individual slack deploy hook

  Edit an individual SlackDeployHook by its ID.
*/
func (a *Client) PatchAPIDeployHooksSlackV3Selector(params *PatchAPIDeployHooksSlackV3SelectorParams, opts ...ClientOption) (*PatchAPIDeployHooksSlackV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIDeployHooksSlackV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIDeployHooksSlackV3Selector",
		Method:             "PATCH",
		PathPattern:        "/api/deploy-hooks/slack/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIDeployHooksSlackV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIDeployHooksSlackV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIDeployHooksSlackV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIDeployHooksGithubActionsV3 creates a github actions deploy hook

  Create a GithubActionsDeployHook.
*/
func (a *Client) PostAPIDeployHooksGithubActionsV3(params *PostAPIDeployHooksGithubActionsV3Params, opts ...ClientOption) (*PostAPIDeployHooksGithubActionsV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIDeployHooksGithubActionsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIDeployHooksGithubActionsV3",
		Method:             "POST",
		PathPattern:        "/api/deploy-hooks/github-actions/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIDeployHooksGithubActionsV3Reader{formats: a.formats},
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
	success, ok := result.(*PostAPIDeployHooksGithubActionsV3Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIDeployHooksGithubActionsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIDeployHooksSlackV3 creates a slack deploy hook

  Create a SlackDeployHook.
*/
func (a *Client) PostAPIDeployHooksSlackV3(params *PostAPIDeployHooksSlackV3Params, opts ...ClientOption) (*PostAPIDeployHooksSlackV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIDeployHooksSlackV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIDeployHooksSlackV3",
		Method:             "POST",
		PathPattern:        "/api/deploy-hooks/slack/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIDeployHooksSlackV3Reader{formats: a.formats},
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
	success, ok := result.(*PostAPIDeployHooksSlackV3Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIDeployHooksSlackV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

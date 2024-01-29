// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new environments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for environments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIV2EnvironmentsSelector(params *DeleteAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*DeleteAPIV2EnvironmentsSelectorOK, error)

	GetAPIEnvironmentsV3(params *GetAPIEnvironmentsV3Params, opts ...ClientOption) (*GetAPIEnvironmentsV3OK, error)

	GetAPIEnvironmentsV3Selector(params *GetAPIEnvironmentsV3SelectorParams, opts ...ClientOption) (*GetAPIEnvironmentsV3SelectorOK, error)

	GetAPIV2Environments(params *GetAPIV2EnvironmentsParams, opts ...ClientOption) (*GetAPIV2EnvironmentsOK, error)

	GetAPIV2EnvironmentsSelector(params *GetAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*GetAPIV2EnvironmentsSelectorOK, error)

	GetAPIV2SelectorsEnvironmentsSelector(params *GetAPIV2SelectorsEnvironmentsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsEnvironmentsSelectorOK, error)

	PatchAPIV2EnvironmentsSelector(params *PatchAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*PatchAPIV2EnvironmentsSelectorOK, error)

	PostAPIV2Environments(params *PostAPIV2EnvironmentsParams, opts ...ClientOption) (*PostAPIV2EnvironmentsOK, *PostAPIV2EnvironmentsCreated, error)

	PostAPIV2ProceduresEnvironmentsTriggerIncidentSelector(params *PostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorParams, opts ...ClientOption) (*PostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorAccepted, error)

	PutAPIV2EnvironmentsSelector(params *PutAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*PutAPIV2EnvironmentsSelectorOK, *PutAPIV2EnvironmentsSelectorCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIV2EnvironmentsSelector deletes a environment entry

  Delete an existing Environment entry via one of its "selectors": name, numeric ID, or "resource-prefix/" + the unique resource prefix.
*/
func (a *Client) DeleteAPIV2EnvironmentsSelector(params *DeleteAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*DeleteAPIV2EnvironmentsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV2EnvironmentsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIV2EnvironmentsSelector",
		Method:             "DELETE",
		PathPattern:        "/api/v2/environments/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIV2EnvironmentsSelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIV2EnvironmentsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIV2EnvironmentsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIEnvironmentsV3 lists environments matching a filter

  List Environments matching a filter.
*/
func (a *Client) GetAPIEnvironmentsV3(params *GetAPIEnvironmentsV3Params, opts ...ClientOption) (*GetAPIEnvironmentsV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIEnvironmentsV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIEnvironmentsV3",
		Method:             "GET",
		PathPattern:        "/api/environments/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIEnvironmentsV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIEnvironmentsV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIEnvironmentsV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIEnvironmentsV3Selector gets an individual environment

  Get an individual Environment.
*/
func (a *Client) GetAPIEnvironmentsV3Selector(params *GetAPIEnvironmentsV3SelectorParams, opts ...ClientOption) (*GetAPIEnvironmentsV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIEnvironmentsV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIEnvironmentsV3Selector",
		Method:             "GET",
		PathPattern:        "/api/environments/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIEnvironmentsV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIEnvironmentsV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIEnvironmentsV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2Environments lists environment entries

  List existing Environment entries, ordered by most recently updated.
*/
func (a *Client) GetAPIV2Environments(params *GetAPIV2EnvironmentsParams, opts ...ClientOption) (*GetAPIV2EnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2EnvironmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Environments",
		Method:             "GET",
		PathPattern:        "/api/v2/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2EnvironmentsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2EnvironmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Environments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2EnvironmentsSelector gets a environment entry

  Get an existing Environment entry via one of its "selectors": name, numeric ID, or "resource-prefix/" + the unique resource prefix.
*/
func (a *Client) GetAPIV2EnvironmentsSelector(params *GetAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*GetAPIV2EnvironmentsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2EnvironmentsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2EnvironmentsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/environments/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2EnvironmentsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2EnvironmentsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2EnvironmentsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2SelectorsEnvironmentsSelector lists environment selectors

  Validate a given Environment selector and provide any other selectors that would match the same Environment.
*/
func (a *Client) GetAPIV2SelectorsEnvironmentsSelector(params *GetAPIV2SelectorsEnvironmentsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsEnvironmentsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SelectorsEnvironmentsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SelectorsEnvironmentsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/selectors/environments/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2SelectorsEnvironmentsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2SelectorsEnvironmentsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SelectorsEnvironmentsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIV2EnvironmentsSelector edits a environment entry

  Edit an existing Environment entry via one of its "selectors": name, numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create, or "resource-prefix/" + the unique resource prefix.
*/
func (a *Client) PatchAPIV2EnvironmentsSelector(params *PatchAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*PatchAPIV2EnvironmentsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2EnvironmentsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2EnvironmentsSelector",
		Method:             "PATCH",
		PathPattern:        "/api/v2/environments/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIV2EnvironmentsSelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2EnvironmentsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2EnvironmentsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIV2Environments creates a new environment entry

  Create a new Environment entry. Note that some fields are immutable after creation; /edit lists mutable fields.
Creating a dynamic environment based on a template will also copy ChartReleases from the template.
*/
func (a *Client) PostAPIV2Environments(params *PostAPIV2EnvironmentsParams, opts ...ClientOption) (*PostAPIV2EnvironmentsOK, *PostAPIV2EnvironmentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2EnvironmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2Environments",
		Method:             "POST",
		PathPattern:        "/api/v2/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2EnvironmentsReader{formats: a.formats},
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
	case *PostAPIV2EnvironmentsOK:
		return value, nil, nil
	case *PostAPIV2EnvironmentsCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIV2ProceduresEnvironmentsTriggerIncidentSelector triggers a pagerduty incident for a given environment

  Trigger an alert for the Pagerduty integration configured for a given Environment.
*/
func (a *Client) PostAPIV2ProceduresEnvironmentsTriggerIncidentSelector(params *PostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorParams, opts ...ClientOption) (*PostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2ProceduresEnvironmentsTriggerIncidentSelector",
		Method:             "POST",
		PathPattern:        "/api/v2/procedures/environments/trigger-incident/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorReader{formats: a.formats},
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
	success, ok := result.(*PostAPIV2ProceduresEnvironmentsTriggerIncidentSelectorAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIV2ProceduresEnvironmentsTriggerIncidentSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAPIV2EnvironmentsSelector creates or edit an environment entry

  Create or edit an Environment entry. Attempts to edit and will attempt to create upon an error.
If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
*/
func (a *Client) PutAPIV2EnvironmentsSelector(params *PutAPIV2EnvironmentsSelectorParams, opts ...ClientOption) (*PutAPIV2EnvironmentsSelectorOK, *PutAPIV2EnvironmentsSelectorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIV2EnvironmentsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutAPIV2EnvironmentsSelector",
		Method:             "PUT",
		PathPattern:        "/api/v2/environments/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAPIV2EnvironmentsSelectorReader{formats: a.formats},
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
	case *PutAPIV2EnvironmentsSelectorOK:
		return value, nil, nil
	case *PutAPIV2EnvironmentsSelectorCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

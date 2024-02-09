// Code generated by go-swagger; DO NOT EDIT.

package database_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new database instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for database instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIDatabaseInstancesV3Selector(params *DeleteAPIDatabaseInstancesV3SelectorParams, opts ...ClientOption) (*DeleteAPIDatabaseInstancesV3SelectorOK, error)

	DeleteAPIV2DatabaseInstancesSelector(params *DeleteAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*DeleteAPIV2DatabaseInstancesSelectorOK, error)

	GetAPIDatabaseInstancesV3(params *GetAPIDatabaseInstancesV3Params, opts ...ClientOption) (*GetAPIDatabaseInstancesV3OK, error)

	GetAPIDatabaseInstancesV3Selector(params *GetAPIDatabaseInstancesV3SelectorParams, opts ...ClientOption) (*GetAPIDatabaseInstancesV3SelectorOK, error)

	GetAPIV2DatabaseInstances(params *GetAPIV2DatabaseInstancesParams, opts ...ClientOption) (*GetAPIV2DatabaseInstancesOK, error)

	GetAPIV2DatabaseInstancesSelector(params *GetAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*GetAPIV2DatabaseInstancesSelectorOK, error)

	GetAPIV2SelectorsDatabaseInstancesSelector(params *GetAPIV2SelectorsDatabaseInstancesSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsDatabaseInstancesSelectorOK, error)

	PatchAPIDatabaseInstancesV3Selector(params *PatchAPIDatabaseInstancesV3SelectorParams, opts ...ClientOption) (*PatchAPIDatabaseInstancesV3SelectorOK, error)

	PatchAPIV2DatabaseInstancesSelector(params *PatchAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*PatchAPIV2DatabaseInstancesSelectorOK, error)

	PostAPIDatabaseInstancesV3(params *PostAPIDatabaseInstancesV3Params, opts ...ClientOption) (*PostAPIDatabaseInstancesV3Created, error)

	PostAPIV2DatabaseInstances(params *PostAPIV2DatabaseInstancesParams, opts ...ClientOption) (*PostAPIV2DatabaseInstancesOK, *PostAPIV2DatabaseInstancesCreated, error)

	PutAPIDatabaseInstancesV3(params *PutAPIDatabaseInstancesV3Params, opts ...ClientOption) (*PutAPIDatabaseInstancesV3OK, *PutAPIDatabaseInstancesV3Created, error)

	PutAPIV2DatabaseInstancesSelector(params *PutAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*PutAPIV2DatabaseInstancesSelectorOK, *PutAPIV2DatabaseInstancesSelectorCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPIDatabaseInstancesV3Selector deletes an individual database instance

  Delete an individual DatabaseInstance by its selector.
*/
func (a *Client) DeleteAPIDatabaseInstancesV3Selector(params *DeleteAPIDatabaseInstancesV3SelectorParams, opts ...ClientOption) (*DeleteAPIDatabaseInstancesV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIDatabaseInstancesV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIDatabaseInstancesV3Selector",
		Method:             "DELETE",
		PathPattern:        "/api/database-instances/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIDatabaseInstancesV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIDatabaseInstancesV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIDatabaseInstancesV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAPIV2DatabaseInstancesSelector deletes a database instance entry

  Delete an existing DatabaseInstance entry via one of its "selectors": numeric ID or 'chart-release/' followed by a chart release selector.
*/
func (a *Client) DeleteAPIV2DatabaseInstancesSelector(params *DeleteAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*DeleteAPIV2DatabaseInstancesSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV2DatabaseInstancesSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIV2DatabaseInstancesSelector",
		Method:             "DELETE",
		PathPattern:        "/api/v2/database-instances/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIV2DatabaseInstancesSelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIV2DatabaseInstancesSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIV2DatabaseInstancesSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIDatabaseInstancesV3 lists database instances matching a filter

  List DatabaseInstances matching a filter.
*/
func (a *Client) GetAPIDatabaseInstancesV3(params *GetAPIDatabaseInstancesV3Params, opts ...ClientOption) (*GetAPIDatabaseInstancesV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIDatabaseInstancesV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIDatabaseInstancesV3",
		Method:             "GET",
		PathPattern:        "/api/database-instances/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIDatabaseInstancesV3Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIDatabaseInstancesV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIDatabaseInstancesV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIDatabaseInstancesV3Selector gets an individual database instance

  Get an individual DatabaseInstance by its selector.
*/
func (a *Client) GetAPIDatabaseInstancesV3Selector(params *GetAPIDatabaseInstancesV3SelectorParams, opts ...ClientOption) (*GetAPIDatabaseInstancesV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIDatabaseInstancesV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIDatabaseInstancesV3Selector",
		Method:             "GET",
		PathPattern:        "/api/database-instances/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIDatabaseInstancesV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIDatabaseInstancesV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIDatabaseInstancesV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2DatabaseInstances lists database instance entries

  List existing DatabaseInstance entries, ordered by most recently updated.
*/
func (a *Client) GetAPIV2DatabaseInstances(params *GetAPIV2DatabaseInstancesParams, opts ...ClientOption) (*GetAPIV2DatabaseInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2DatabaseInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2DatabaseInstances",
		Method:             "GET",
		PathPattern:        "/api/v2/database-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2DatabaseInstancesReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2DatabaseInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2DatabaseInstances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2DatabaseInstancesSelector gets a database instance entry

  Get an existing DatabaseInstance entry via one of its "selectors": numeric ID or 'chart-release/' followed by a chart release selector.
*/
func (a *Client) GetAPIV2DatabaseInstancesSelector(params *GetAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*GetAPIV2DatabaseInstancesSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2DatabaseInstancesSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2DatabaseInstancesSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/database-instances/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2DatabaseInstancesSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2DatabaseInstancesSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2DatabaseInstancesSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2SelectorsDatabaseInstancesSelector lists database instance selectors

  Validate a given DatabaseInstance selector and provide any other selectors that would match the same DatabaseInstance.
*/
func (a *Client) GetAPIV2SelectorsDatabaseInstancesSelector(params *GetAPIV2SelectorsDatabaseInstancesSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsDatabaseInstancesSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SelectorsDatabaseInstancesSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SelectorsDatabaseInstancesSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/selectors/database-instances/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2SelectorsDatabaseInstancesSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2SelectorsDatabaseInstancesSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SelectorsDatabaseInstancesSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIDatabaseInstancesV3Selector edits an individual database instance

  Edit an individual DatabaseInstance by its selector.
*/
func (a *Client) PatchAPIDatabaseInstancesV3Selector(params *PatchAPIDatabaseInstancesV3SelectorParams, opts ...ClientOption) (*PatchAPIDatabaseInstancesV3SelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIDatabaseInstancesV3SelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIDatabaseInstancesV3Selector",
		Method:             "PATCH",
		PathPattern:        "/api/database-instances/v3/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIDatabaseInstancesV3SelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIDatabaseInstancesV3SelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIDatabaseInstancesV3Selector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPIV2DatabaseInstancesSelector edits a database instance entry

  Edit an existing DatabaseInstance entry via one of its "selectors": numeric ID or 'chart-release/' followed by a chart release selector. Note that only mutable fields are available here, immutable fields can only be set using /create.
*/
func (a *Client) PatchAPIV2DatabaseInstancesSelector(params *PatchAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*PatchAPIV2DatabaseInstancesSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2DatabaseInstancesSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2DatabaseInstancesSelector",
		Method:             "PATCH",
		PathPattern:        "/api/v2/database-instances/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIV2DatabaseInstancesSelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2DatabaseInstancesSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2DatabaseInstancesSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIDatabaseInstancesV3 creates a database instance

  Create a DatabaseInstance.
*/
func (a *Client) PostAPIDatabaseInstancesV3(params *PostAPIDatabaseInstancesV3Params, opts ...ClientOption) (*PostAPIDatabaseInstancesV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIDatabaseInstancesV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIDatabaseInstancesV3",
		Method:             "POST",
		PathPattern:        "/api/database-instances/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIDatabaseInstancesV3Reader{formats: a.formats},
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
	success, ok := result.(*PostAPIDatabaseInstancesV3Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIDatabaseInstancesV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIV2DatabaseInstances creates a new database instance entry

  Create a new DatabaseInstance entry. Note that some fields are immutable after creation; /edit lists mutable fields.
*/
func (a *Client) PostAPIV2DatabaseInstances(params *PostAPIV2DatabaseInstancesParams, opts ...ClientOption) (*PostAPIV2DatabaseInstancesOK, *PostAPIV2DatabaseInstancesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2DatabaseInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2DatabaseInstances",
		Method:             "POST",
		PathPattern:        "/api/v2/database-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2DatabaseInstancesReader{formats: a.formats},
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
	case *PostAPIV2DatabaseInstancesOK:
		return value, nil, nil
	case *PostAPIV2DatabaseInstancesCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for database_instances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAPIDatabaseInstancesV3 creates or edit a database instance

  Create or edit a DatabaseInstance, depending on whether one already exists for the chart release
*/
func (a *Client) PutAPIDatabaseInstancesV3(params *PutAPIDatabaseInstancesV3Params, opts ...ClientOption) (*PutAPIDatabaseInstancesV3OK, *PutAPIDatabaseInstancesV3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIDatabaseInstancesV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutAPIDatabaseInstancesV3",
		Method:             "PUT",
		PathPattern:        "/api/database-instances/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAPIDatabaseInstancesV3Reader{formats: a.formats},
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
	case *PutAPIDatabaseInstancesV3OK:
		return value, nil, nil
	case *PutAPIDatabaseInstancesV3Created:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for database_instances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAPIV2DatabaseInstancesSelector creates or edit a database instance entry

  Create or edit a DatabaseInstance entry. Attempts to edit and will attempt to create upon an error.
If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
*/
func (a *Client) PutAPIV2DatabaseInstancesSelector(params *PutAPIV2DatabaseInstancesSelectorParams, opts ...ClientOption) (*PutAPIV2DatabaseInstancesSelectorOK, *PutAPIV2DatabaseInstancesSelectorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIV2DatabaseInstancesSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutAPIV2DatabaseInstancesSelector",
		Method:             "PUT",
		PathPattern:        "/api/v2/database-instances/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAPIV2DatabaseInstancesSelectorReader{formats: a.formats},
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
	case *PutAPIV2DatabaseInstancesSelectorOK:
		return value, nil, nil
	case *PutAPIV2DatabaseInstancesSelectorCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for database_instances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

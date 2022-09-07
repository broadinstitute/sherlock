// Code generated by go-swagger; DO NOT EDIT.

package app_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new app versions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for app versions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV2AppVersions(params *GetAPIV2AppVersionsParams, opts ...ClientOption) (*GetAPIV2AppVersionsOK, error)

	GetAPIV2AppVersionsSelector(params *GetAPIV2AppVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2AppVersionsSelectorOK, error)

	GetAPIV2SelectorsAppVersionsSelector(params *GetAPIV2SelectorsAppVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsAppVersionsSelectorOK, error)

	PostAPIV2AppVersions(params *PostAPIV2AppVersionsParams, opts ...ClientOption) (*PostAPIV2AppVersionsOK, *PostAPIV2AppVersionsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPIV2AppVersions lists app version entries

  List existing AppVersion entries, ordered by most recently updated.
*/
func (a *Client) GetAPIV2AppVersions(params *GetAPIV2AppVersionsParams, opts ...ClientOption) (*GetAPIV2AppVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2AppVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2AppVersions",
		Method:             "GET",
		PathPattern:        "/api/v2/app-versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2AppVersionsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2AppVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2AppVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2AppVersionsSelector gets a app version entry

  Get an existing AppVersion entry via one its "selectors": chart/version or numeric ID.
*/
func (a *Client) GetAPIV2AppVersionsSelector(params *GetAPIV2AppVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2AppVersionsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2AppVersionsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2AppVersionsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/app-versions/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2AppVersionsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2AppVersionsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2AppVersionsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIV2SelectorsAppVersionsSelector lists app version selectors

  Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.
*/
func (a *Client) GetAPIV2SelectorsAppVersionsSelector(params *GetAPIV2SelectorsAppVersionsSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsAppVersionsSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SelectorsAppVersionsSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SelectorsAppVersionsSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/selectors/app-versions/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2SelectorsAppVersionsSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2SelectorsAppVersionsSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SelectorsAppVersionsSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPIV2AppVersions creates a new app version entry

  Create a new AppVersion entry. Note that fields are immutable after creation.
If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
*/
func (a *Client) PostAPIV2AppVersions(params *PostAPIV2AppVersionsParams, opts ...ClientOption) (*PostAPIV2AppVersionsOK, *PostAPIV2AppVersionsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2AppVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2AppVersions",
		Method:             "POST",
		PathPattern:        "/api/v2/app-versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2AppVersionsReader{formats: a.formats},
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
	case *PostAPIV2AppVersionsOK:
		return value, nil, nil
	case *PostAPIV2AppVersionsCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for app_versions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

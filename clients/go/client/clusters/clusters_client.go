// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIV2ClustersSelector(params *DeleteAPIV2ClustersSelectorParams, opts ...ClientOption) (*DeleteAPIV2ClustersSelectorOK, error)

	GetAPIV2Clusters(params *GetAPIV2ClustersParams, opts ...ClientOption) (*GetAPIV2ClustersOK, error)

	GetAPIV2ClustersSelector(params *GetAPIV2ClustersSelectorParams, opts ...ClientOption) (*GetAPIV2ClustersSelectorOK, error)

	GetAPIV2SelectorsClustersSelector(params *GetAPIV2SelectorsClustersSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsClustersSelectorOK, error)

	PatchAPIV2ClustersSelector(params *PatchAPIV2ClustersSelectorParams, opts ...ClientOption) (*PatchAPIV2ClustersSelectorOK, error)

	PostAPIV2Clusters(params *PostAPIV2ClustersParams, opts ...ClientOption) (*PostAPIV2ClustersOK, *PostAPIV2ClustersCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPIV2ClustersSelector deletes a cluster entry

Delete an existing Cluster entry via one of its "selectors": name or numeric ID.
*/
func (a *Client) DeleteAPIV2ClustersSelector(params *DeleteAPIV2ClustersSelectorParams, opts ...ClientOption) (*DeleteAPIV2ClustersSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV2ClustersSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIV2ClustersSelector",
		Method:             "DELETE",
		PathPattern:        "/api/v2/clusters/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIV2ClustersSelectorReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIV2ClustersSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIV2ClustersSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2Clusters lists cluster entries

List existing Cluster entries, ordered by most recently updated.
*/
func (a *Client) GetAPIV2Clusters(params *GetAPIV2ClustersParams, opts ...ClientOption) (*GetAPIV2ClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Clusters",
		Method:             "GET",
		PathPattern:        "/api/v2/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2ClustersReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2ClustersSelector gets a cluster entry

Get an existing Cluster entry via one of its "selectors": name or numeric ID.
*/
func (a *Client) GetAPIV2ClustersSelector(params *GetAPIV2ClustersSelectorParams, opts ...ClientOption) (*GetAPIV2ClustersSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ClustersSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2ClustersSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/clusters/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2ClustersSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ClustersSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2ClustersSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2SelectorsClustersSelector lists cluster selectors

Validate a given Cluster selector and provide any other selectors that would match the same Cluster.
*/
func (a *Client) GetAPIV2SelectorsClustersSelector(params *GetAPIV2SelectorsClustersSelectorParams, opts ...ClientOption) (*GetAPIV2SelectorsClustersSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SelectorsClustersSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SelectorsClustersSelector",
		Method:             "GET",
		PathPattern:        "/api/v2/selectors/clusters/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIV2SelectorsClustersSelectorReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2SelectorsClustersSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SelectorsClustersSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIV2ClustersSelector edits a cluster entry

Edit an existing Cluster entry via one of its "selectors": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
*/
func (a *Client) PatchAPIV2ClustersSelector(params *PatchAPIV2ClustersSelectorParams, opts ...ClientOption) (*PatchAPIV2ClustersSelectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2ClustersSelectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2ClustersSelector",
		Method:             "PATCH",
		PathPattern:        "/api/v2/clusters/{selector}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchAPIV2ClustersSelectorReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2ClustersSelectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2ClustersSelector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIV2Clusters creates a new cluster entry

Create a new Cluster entry. Note that some fields are immutable after creation; /edit lists mutable fields.
*/
func (a *Client) PostAPIV2Clusters(params *PostAPIV2ClustersParams, opts ...ClientOption) (*PostAPIV2ClustersOK, *PostAPIV2ClustersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2ClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2Clusters",
		Method:             "POST",
		PathPattern:        "/api/v2/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIV2ClustersReader{formats: a.formats},
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
	case *PostAPIV2ClustersOK:
		return value, nil, nil
	case *PostAPIV2ClustersCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

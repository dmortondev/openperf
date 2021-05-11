// Code generated by go-swagger; DO NOT EDIT.

package network_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network generator API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network generator API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BulkCreateNetworkGenerators(params *BulkCreateNetworkGeneratorsParams, opts ...ClientOption) (*BulkCreateNetworkGeneratorsOK, error)

	BulkCreateNetworkServers(params *BulkCreateNetworkServersParams, opts ...ClientOption) (*BulkCreateNetworkServersOK, error)

	BulkDeleteNetworkGenerators(params *BulkDeleteNetworkGeneratorsParams, opts ...ClientOption) (*BulkDeleteNetworkGeneratorsNoContent, error)

	BulkDeleteNetworkServers(params *BulkDeleteNetworkServersParams, opts ...ClientOption) (*BulkDeleteNetworkServersNoContent, error)

	BulkStartNetworkGenerators(params *BulkStartNetworkGeneratorsParams, opts ...ClientOption) (*BulkStartNetworkGeneratorsOK, error)

	BulkStopNetworkGenerators(params *BulkStopNetworkGeneratorsParams, opts ...ClientOption) (*BulkStopNetworkGeneratorsNoContent, error)

	CreateNetworkGenerator(params *CreateNetworkGeneratorParams, opts ...ClientOption) (*CreateNetworkGeneratorCreated, error)

	CreateNetworkServer(params *CreateNetworkServerParams, opts ...ClientOption) (*CreateNetworkServerCreated, error)

	DeleteNetworkGenerator(params *DeleteNetworkGeneratorParams, opts ...ClientOption) (*DeleteNetworkGeneratorNoContent, error)

	DeleteNetworkGeneratorResult(params *DeleteNetworkGeneratorResultParams, opts ...ClientOption) (*DeleteNetworkGeneratorResultNoContent, error)

	DeleteNetworkServer(params *DeleteNetworkServerParams, opts ...ClientOption) (*DeleteNetworkServerNoContent, error)

	GetNetworkGenerator(params *GetNetworkGeneratorParams, opts ...ClientOption) (*GetNetworkGeneratorOK, error)

	GetNetworkGeneratorResult(params *GetNetworkGeneratorResultParams, opts ...ClientOption) (*GetNetworkGeneratorResultOK, error)

	GetNetworkServer(params *GetNetworkServerParams, opts ...ClientOption) (*GetNetworkServerOK, error)

	ListNetworkGeneratorResults(params *ListNetworkGeneratorResultsParams, opts ...ClientOption) (*ListNetworkGeneratorResultsOK, error)

	ListNetworkGenerators(params *ListNetworkGeneratorsParams, opts ...ClientOption) (*ListNetworkGeneratorsOK, error)

	ListNetworkServers(params *ListNetworkServersParams, opts ...ClientOption) (*ListNetworkServersOK, error)

	StartNetworkGenerator(params *StartNetworkGeneratorParams, opts ...ClientOption) (*StartNetworkGeneratorCreated, error)

	StopNetworkGenerator(params *StopNetworkGeneratorParams, opts ...ClientOption) (*StopNetworkGeneratorNoContent, error)

	ToggleNetworkGenerators(params *ToggleNetworkGeneratorsParams, opts ...ClientOption) (*ToggleNetworkGeneratorsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BulkCreateNetworkGenerators bulks create network generators

  Create multiple network generators. Requests are processed in an
all-or-nothing manner, i.e. a single network generator creation failure
causes all creations for this request to fail.

*/
func (a *Client) BulkCreateNetworkGenerators(params *BulkCreateNetworkGeneratorsParams, opts ...ClientOption) (*BulkCreateNetworkGeneratorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkCreateNetworkGeneratorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkCreateNetworkGenerators",
		Method:             "POST",
		PathPattern:        "/network/generators/x/bulk-create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkCreateNetworkGeneratorsReader{formats: a.formats},
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
	success, ok := result.(*BulkCreateNetworkGeneratorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkCreateNetworkGenerators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkCreateNetworkServers bulks create network servers

  Create multiple network servers. Requests are processed in an
all-or-nothing manner, i.e. a single network server creation failure
causes all creations for this request to fail.

*/
func (a *Client) BulkCreateNetworkServers(params *BulkCreateNetworkServersParams, opts ...ClientOption) (*BulkCreateNetworkServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkCreateNetworkServersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkCreateNetworkServers",
		Method:             "POST",
		PathPattern:        "/network/servers/x/bulk-create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkCreateNetworkServersReader{formats: a.formats},
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
	success, ok := result.(*BulkCreateNetworkServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkCreateNetworkServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkDeleteNetworkGenerators bulks delete network generators

  Delete multiple network generators in a best-effort manner. Non-existant network generators ids
do not cause errors. Idempotent.

*/
func (a *Client) BulkDeleteNetworkGenerators(params *BulkDeleteNetworkGeneratorsParams, opts ...ClientOption) (*BulkDeleteNetworkGeneratorsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkDeleteNetworkGeneratorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkDeleteNetworkGenerators",
		Method:             "POST",
		PathPattern:        "/network/generators/x/bulk-delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkDeleteNetworkGeneratorsReader{formats: a.formats},
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
	success, ok := result.(*BulkDeleteNetworkGeneratorsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkDeleteNetworkGenerators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkDeleteNetworkServers bulks delete network servers

  Delete multiple network servers in a best-effort manner. Non-existant network server ids
do not cause errors. Idempotent.

*/
func (a *Client) BulkDeleteNetworkServers(params *BulkDeleteNetworkServersParams, opts ...ClientOption) (*BulkDeleteNetworkServersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkDeleteNetworkServersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkDeleteNetworkServers",
		Method:             "POST",
		PathPattern:        "/network/servers/x/bulk-delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkDeleteNetworkServersReader{formats: a.formats},
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
	success, ok := result.(*BulkDeleteNetworkServersNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkDeleteNetworkServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkStartNetworkGenerators tells multiple network generators to start

  Start multiple network generators.
*/
func (a *Client) BulkStartNetworkGenerators(params *BulkStartNetworkGeneratorsParams, opts ...ClientOption) (*BulkStartNetworkGeneratorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkStartNetworkGeneratorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkStartNetworkGenerators",
		Method:             "POST",
		PathPattern:        "/network/generators/x/bulk-start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkStartNetworkGeneratorsReader{formats: a.formats},
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
	success, ok := result.(*BulkStartNetworkGeneratorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkStartNetworkGenerators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkStopNetworkGenerators bulks stop network generators

  Best-effort stop multiple network generators. Non-existent network generator ids do not cause errors. Idempotent.
*/
func (a *Client) BulkStopNetworkGenerators(params *BulkStopNetworkGeneratorsParams, opts ...ClientOption) (*BulkStopNetworkGeneratorsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkStopNetworkGeneratorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BulkStopNetworkGenerators",
		Method:             "POST",
		PathPattern:        "/network/generators/x/bulk-stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkStopNetworkGeneratorsReader{formats: a.formats},
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
	success, ok := result.(*BulkStopNetworkGeneratorsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BulkStopNetworkGenerators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateNetworkGenerator creates a network generator

  Create a new network generator
*/
func (a *Client) CreateNetworkGenerator(params *CreateNetworkGeneratorParams, opts ...ClientOption) (*CreateNetworkGeneratorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkGeneratorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateNetworkGenerator",
		Method:             "POST",
		PathPattern:        "/network/generators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateNetworkGeneratorReader{formats: a.formats},
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
	success, ok := result.(*CreateNetworkGeneratorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateNetworkGenerator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateNetworkServer creates and run a network server

  Create a new network server.
*/
func (a *Client) CreateNetworkServer(params *CreateNetworkServerParams, opts ...ClientOption) (*CreateNetworkServerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateNetworkServer",
		Method:             "POST",
		PathPattern:        "/network/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateNetworkServerReader{formats: a.formats},
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
	success, ok := result.(*CreateNetworkServerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateNetworkServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkGenerator deletes a network generator

  Deletes an existing network generator. Idempotent.
*/
func (a *Client) DeleteNetworkGenerator(params *DeleteNetworkGeneratorParams, opts ...ClientOption) (*DeleteNetworkGeneratorNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkGeneratorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNetworkGenerator",
		Method:             "DELETE",
		PathPattern:        "/network/generators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworkGeneratorReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkGeneratorNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNetworkGenerator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkGeneratorResult deletes results from a network generator idempotent
*/
func (a *Client) DeleteNetworkGeneratorResult(params *DeleteNetworkGeneratorResultParams, opts ...ClientOption) (*DeleteNetworkGeneratorResultNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkGeneratorResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNetworkGeneratorResult",
		Method:             "DELETE",
		PathPattern:        "/network/generator-results/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworkGeneratorResultReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkGeneratorResultNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNetworkGeneratorResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkServer deletes a network server

  Deletes an existing network server. Idempotent.
*/
func (a *Client) DeleteNetworkServer(params *DeleteNetworkServerParams, opts ...ClientOption) (*DeleteNetworkServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNetworkServer",
		Method:             "DELETE",
		PathPattern:        "/network/servers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworkServerReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNetworkServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkGenerator gets a network generator

  Returns a network generator, by id.
*/
func (a *Client) GetNetworkGenerator(params *GetNetworkGeneratorParams, opts ...ClientOption) (*GetNetworkGeneratorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkGeneratorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkGenerator",
		Method:             "GET",
		PathPattern:        "/network/generators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworkGeneratorReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkGeneratorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkGenerator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkGeneratorResult gets a result from a network generator

  Returns results from a network generator by result id.
*/
func (a *Client) GetNetworkGeneratorResult(params *GetNetworkGeneratorResultParams, opts ...ClientOption) (*GetNetworkGeneratorResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkGeneratorResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkGeneratorResult",
		Method:             "GET",
		PathPattern:        "/network/generator-results/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworkGeneratorResultReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkGeneratorResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkGeneratorResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkServer gets a network server

  Returns a network server, by id.
*/
func (a *Client) GetNetworkServer(params *GetNetworkServerParams, opts ...ClientOption) (*GetNetworkServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkServer",
		Method:             "GET",
		PathPattern:        "/network/servers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworkServerReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNetworkGeneratorResults lists network generator results

  The `network-generator-results` endpoint returns all of the results produced by running network generators.
*/
func (a *Client) ListNetworkGeneratorResults(params *ListNetworkGeneratorResultsParams, opts ...ClientOption) (*ListNetworkGeneratorResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNetworkGeneratorResultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNetworkGeneratorResults",
		Method:             "GET",
		PathPattern:        "/network/generator-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNetworkGeneratorResultsReader{formats: a.formats},
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
	success, ok := result.(*ListNetworkGeneratorResultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListNetworkGeneratorResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNetworkGenerators lists network generators

  The `network-generators` endpoint returns all of the configured network generators.
*/
func (a *Client) ListNetworkGenerators(params *ListNetworkGeneratorsParams, opts ...ClientOption) (*ListNetworkGeneratorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNetworkGeneratorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNetworkGenerators",
		Method:             "GET",
		PathPattern:        "/network/generators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNetworkGeneratorsReader{formats: a.formats},
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
	success, ok := result.(*ListNetworkGeneratorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListNetworkGenerators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNetworkServers lists network servers

  Returns all network servers
*/
func (a *Client) ListNetworkServers(params *ListNetworkServersParams, opts ...ClientOption) (*ListNetworkServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNetworkServersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNetworkServers",
		Method:             "GET",
		PathPattern:        "/network/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNetworkServersReader{formats: a.formats},
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
	success, ok := result.(*ListNetworkServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListNetworkServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartNetworkGenerator starts a network generator

  Start an existing network generator.
*/
func (a *Client) StartNetworkGenerator(params *StartNetworkGeneratorParams, opts ...ClientOption) (*StartNetworkGeneratorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartNetworkGeneratorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartNetworkGenerator",
		Method:             "POST",
		PathPattern:        "/network/generators/{id}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartNetworkGeneratorReader{formats: a.formats},
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
	success, ok := result.(*StartNetworkGeneratorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartNetworkGenerator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StopNetworkGenerator stops a network generator

  Stop a running network generator. Idempotent.
*/
func (a *Client) StopNetworkGenerator(params *StopNetworkGeneratorParams, opts ...ClientOption) (*StopNetworkGeneratorNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopNetworkGeneratorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopNetworkGenerator",
		Method:             "POST",
		PathPattern:        "/network/generators/{id}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopNetworkGeneratorReader{formats: a.formats},
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
	success, ok := result.(*StopNetworkGeneratorNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopNetworkGenerator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ToggleNetworkGenerators replaces a running network generator with a stopped network generator

  Swap a running network generator with an idle network generator.
Upon success, the idle generator will be in the run state, the running generator
will be in the stopped state and all active TCP/UDP sessions will be transferred
to the newly running generator.
If the original network generator had a read/write load and the new network generator
does not have this type of load, these sessions will be immediately stopped.

*/
func (a *Client) ToggleNetworkGenerators(params *ToggleNetworkGeneratorsParams, opts ...ClientOption) (*ToggleNetworkGeneratorsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewToggleNetworkGeneratorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ToggleNetworkGenerators",
		Method:             "POST",
		PathPattern:        "/network/generators/x/toggle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ToggleNetworkGeneratorsReader{formats: a.formats},
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
	success, ok := result.(*ToggleNetworkGeneratorsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ToggleNetworkGenerators: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

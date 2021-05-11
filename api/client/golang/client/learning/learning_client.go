// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new learning API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for learning API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPacketGeneratorLearningResults(params *GetPacketGeneratorLearningResultsParams, opts ...ClientOption) (*GetPacketGeneratorLearningResultsOK, error)

	RetryPacketGeneratorLearning(params *RetryPacketGeneratorLearningParams, opts ...ClientOption) (*RetryPacketGeneratorLearningNoContent, error)

	StartPacketGeneratorLearning(params *StartPacketGeneratorLearningParams, opts ...ClientOption) (*StartPacketGeneratorLearningNoContent, error)

	StopPacketGeneratorLearning(params *StopPacketGeneratorLearningParams, opts ...ClientOption) (*StopPacketGeneratorLearningNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPacketGeneratorLearningResults gets detailed learning information

  Returns learning state and resolved addresses for
a packet generator attached to an interface, by id.

*/
func (a *Client) GetPacketGeneratorLearningResults(params *GetPacketGeneratorLearningResultsParams, opts ...ClientOption) (*GetPacketGeneratorLearningResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPacketGeneratorLearningResultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPacketGeneratorLearningResults",
		Method:             "GET",
		PathPattern:        "/packet/generators/{id}/learning",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPacketGeneratorLearningResultsReader{formats: a.formats},
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
	success, ok := result.(*GetPacketGeneratorLearningResultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPacketGeneratorLearningResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetryPacketGeneratorLearning retries m a c learning

  Used to retry MAC learning on a generator bound to an interface.
Performs MAC learning for only unresolved addresses.

*/
func (a *Client) RetryPacketGeneratorLearning(params *RetryPacketGeneratorLearningParams, opts ...ClientOption) (*RetryPacketGeneratorLearningNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryPacketGeneratorLearningParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetryPacketGeneratorLearning",
		Method:             "POST",
		PathPattern:        "/packet/generators/{id}/learning/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetryPacketGeneratorLearningReader{formats: a.formats},
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
	success, ok := result.(*RetryPacketGeneratorLearningNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetryPacketGeneratorLearning: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartPacketGeneratorLearning starts m a c learning

  Used to start MAC learning on a generator bound to an interface.
Clears previously resolved MAC table.

*/
func (a *Client) StartPacketGeneratorLearning(params *StartPacketGeneratorLearningParams, opts ...ClientOption) (*StartPacketGeneratorLearningNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPacketGeneratorLearningParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartPacketGeneratorLearning",
		Method:             "POST",
		PathPattern:        "/packet/generators/{id}/learning/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPacketGeneratorLearningReader{formats: a.formats},
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
	success, ok := result.(*StartPacketGeneratorLearningNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartPacketGeneratorLearning: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StopPacketGeneratorLearning stops m a c learning

  Used to stop MAC learning on a generator bound to
an interface.

*/
func (a *Client) StopPacketGeneratorLearning(params *StopPacketGeneratorLearningParams, opts ...ClientOption) (*StopPacketGeneratorLearningNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopPacketGeneratorLearningParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopPacketGeneratorLearning",
		Method:             "POST",
		PathPattern:        "/packet/generators/{id}/learning/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopPacketGeneratorLearningReader{formats: a.formats},
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
	success, ok := result.(*StopPacketGeneratorLearningNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopPacketGeneratorLearning: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

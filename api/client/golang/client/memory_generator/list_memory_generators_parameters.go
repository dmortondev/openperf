// Code generated by go-swagger; DO NOT EDIT.

package memory_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListMemoryGeneratorsParams creates a new ListMemoryGeneratorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMemoryGeneratorsParams() *ListMemoryGeneratorsParams {
	return &ListMemoryGeneratorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMemoryGeneratorsParamsWithTimeout creates a new ListMemoryGeneratorsParams object
// with the ability to set a timeout on a request.
func NewListMemoryGeneratorsParamsWithTimeout(timeout time.Duration) *ListMemoryGeneratorsParams {
	return &ListMemoryGeneratorsParams{
		timeout: timeout,
	}
}

// NewListMemoryGeneratorsParamsWithContext creates a new ListMemoryGeneratorsParams object
// with the ability to set a context for a request.
func NewListMemoryGeneratorsParamsWithContext(ctx context.Context) *ListMemoryGeneratorsParams {
	return &ListMemoryGeneratorsParams{
		Context: ctx,
	}
}

// NewListMemoryGeneratorsParamsWithHTTPClient creates a new ListMemoryGeneratorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMemoryGeneratorsParamsWithHTTPClient(client *http.Client) *ListMemoryGeneratorsParams {
	return &ListMemoryGeneratorsParams{
		HTTPClient: client,
	}
}

/* ListMemoryGeneratorsParams contains all the parameters to send to the API endpoint
   for the list memory generators operation.

   Typically these are written to a http.Request.
*/
type ListMemoryGeneratorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list memory generators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMemoryGeneratorsParams) WithDefaults() *ListMemoryGeneratorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list memory generators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMemoryGeneratorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list memory generators params
func (o *ListMemoryGeneratorsParams) WithTimeout(timeout time.Duration) *ListMemoryGeneratorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list memory generators params
func (o *ListMemoryGeneratorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list memory generators params
func (o *ListMemoryGeneratorsParams) WithContext(ctx context.Context) *ListMemoryGeneratorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list memory generators params
func (o *ListMemoryGeneratorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list memory generators params
func (o *ListMemoryGeneratorsParams) WithHTTPClient(client *http.Client) *ListMemoryGeneratorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list memory generators params
func (o *ListMemoryGeneratorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListMemoryGeneratorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

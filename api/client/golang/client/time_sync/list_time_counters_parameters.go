// Code generated by go-swagger; DO NOT EDIT.

package time_sync

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

// NewListTimeCountersParams creates a new ListTimeCountersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTimeCountersParams() *ListTimeCountersParams {
	return &ListTimeCountersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTimeCountersParamsWithTimeout creates a new ListTimeCountersParams object
// with the ability to set a timeout on a request.
func NewListTimeCountersParamsWithTimeout(timeout time.Duration) *ListTimeCountersParams {
	return &ListTimeCountersParams{
		timeout: timeout,
	}
}

// NewListTimeCountersParamsWithContext creates a new ListTimeCountersParams object
// with the ability to set a context for a request.
func NewListTimeCountersParamsWithContext(ctx context.Context) *ListTimeCountersParams {
	return &ListTimeCountersParams{
		Context: ctx,
	}
}

// NewListTimeCountersParamsWithHTTPClient creates a new ListTimeCountersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTimeCountersParamsWithHTTPClient(client *http.Client) *ListTimeCountersParams {
	return &ListTimeCountersParams{
		HTTPClient: client,
	}
}

/* ListTimeCountersParams contains all the parameters to send to the API endpoint
   for the list time counters operation.

   Typically these are written to a http.Request.
*/
type ListTimeCountersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list time counters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTimeCountersParams) WithDefaults() *ListTimeCountersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list time counters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTimeCountersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list time counters params
func (o *ListTimeCountersParams) WithTimeout(timeout time.Duration) *ListTimeCountersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list time counters params
func (o *ListTimeCountersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list time counters params
func (o *ListTimeCountersParams) WithContext(ctx context.Context) *ListTimeCountersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list time counters params
func (o *ListTimeCountersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list time counters params
func (o *ListTimeCountersParams) WithHTTPClient(client *http.Client) *ListTimeCountersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list time counters params
func (o *ListTimeCountersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListTimeCountersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

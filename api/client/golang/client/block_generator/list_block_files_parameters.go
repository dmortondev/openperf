// Code generated by go-swagger; DO NOT EDIT.

package block_generator

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

// NewListBlockFilesParams creates a new ListBlockFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBlockFilesParams() *ListBlockFilesParams {
	return &ListBlockFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBlockFilesParamsWithTimeout creates a new ListBlockFilesParams object
// with the ability to set a timeout on a request.
func NewListBlockFilesParamsWithTimeout(timeout time.Duration) *ListBlockFilesParams {
	return &ListBlockFilesParams{
		timeout: timeout,
	}
}

// NewListBlockFilesParamsWithContext creates a new ListBlockFilesParams object
// with the ability to set a context for a request.
func NewListBlockFilesParamsWithContext(ctx context.Context) *ListBlockFilesParams {
	return &ListBlockFilesParams{
		Context: ctx,
	}
}

// NewListBlockFilesParamsWithHTTPClient creates a new ListBlockFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBlockFilesParamsWithHTTPClient(client *http.Client) *ListBlockFilesParams {
	return &ListBlockFilesParams{
		HTTPClient: client,
	}
}

/* ListBlockFilesParams contains all the parameters to send to the API endpoint
   for the list block files operation.

   Typically these are written to a http.Request.
*/
type ListBlockFilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list block files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBlockFilesParams) WithDefaults() *ListBlockFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list block files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBlockFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list block files params
func (o *ListBlockFilesParams) WithTimeout(timeout time.Duration) *ListBlockFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list block files params
func (o *ListBlockFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list block files params
func (o *ListBlockFilesParams) WithContext(ctx context.Context) *ListBlockFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list block files params
func (o *ListBlockFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list block files params
func (o *ListBlockFilesParams) WithHTTPClient(client *http.Client) *ListBlockFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list block files params
func (o *ListBlockFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListBlockFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

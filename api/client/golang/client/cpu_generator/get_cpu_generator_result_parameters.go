// Code generated by go-swagger; DO NOT EDIT.

package cpu_generator

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

// NewGetCPUGeneratorResultParams creates a new GetCPUGeneratorResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCPUGeneratorResultParams() *GetCPUGeneratorResultParams {
	return &GetCPUGeneratorResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCPUGeneratorResultParamsWithTimeout creates a new GetCPUGeneratorResultParams object
// with the ability to set a timeout on a request.
func NewGetCPUGeneratorResultParamsWithTimeout(timeout time.Duration) *GetCPUGeneratorResultParams {
	return &GetCPUGeneratorResultParams{
		timeout: timeout,
	}
}

// NewGetCPUGeneratorResultParamsWithContext creates a new GetCPUGeneratorResultParams object
// with the ability to set a context for a request.
func NewGetCPUGeneratorResultParamsWithContext(ctx context.Context) *GetCPUGeneratorResultParams {
	return &GetCPUGeneratorResultParams{
		Context: ctx,
	}
}

// NewGetCPUGeneratorResultParamsWithHTTPClient creates a new GetCPUGeneratorResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCPUGeneratorResultParamsWithHTTPClient(client *http.Client) *GetCPUGeneratorResultParams {
	return &GetCPUGeneratorResultParams{
		HTTPClient: client,
	}
}

/* GetCPUGeneratorResultParams contains all the parameters to send to the API endpoint
   for the get Cpu generator result operation.

   Typically these are written to a http.Request.
*/
type GetCPUGeneratorResultParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Cpu generator result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCPUGeneratorResultParams) WithDefaults() *GetCPUGeneratorResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Cpu generator result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCPUGeneratorResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) WithTimeout(timeout time.Duration) *GetCPUGeneratorResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) WithContext(ctx context.Context) *GetCPUGeneratorResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) WithHTTPClient(client *http.Client) *GetCPUGeneratorResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) WithID(id string) *GetCPUGeneratorResultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get Cpu generator result params
func (o *GetCPUGeneratorResultParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCPUGeneratorResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

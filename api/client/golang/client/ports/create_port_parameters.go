// Code generated by go-swagger; DO NOT EDIT.

package ports

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

	"github.com/spirent/openperf/api/client/golang/models"
)

// NewCreatePortParams creates a new CreatePortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePortParams() *CreatePortParams {
	return &CreatePortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePortParamsWithTimeout creates a new CreatePortParams object
// with the ability to set a timeout on a request.
func NewCreatePortParamsWithTimeout(timeout time.Duration) *CreatePortParams {
	return &CreatePortParams{
		timeout: timeout,
	}
}

// NewCreatePortParamsWithContext creates a new CreatePortParams object
// with the ability to set a context for a request.
func NewCreatePortParamsWithContext(ctx context.Context) *CreatePortParams {
	return &CreatePortParams{
		Context: ctx,
	}
}

// NewCreatePortParamsWithHTTPClient creates a new CreatePortParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePortParamsWithHTTPClient(client *http.Client) *CreatePortParams {
	return &CreatePortParams{
		HTTPClient: client,
	}
}

/* CreatePortParams contains all the parameters to send to the API endpoint
   for the create port operation.

   Typically these are written to a http.Request.
*/
type CreatePortParams struct {

	/* Port.

	   New port
	*/
	Port *models.Port

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePortParams) WithDefaults() *CreatePortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create port params
func (o *CreatePortParams) WithTimeout(timeout time.Duration) *CreatePortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create port params
func (o *CreatePortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create port params
func (o *CreatePortParams) WithContext(ctx context.Context) *CreatePortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create port params
func (o *CreatePortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create port params
func (o *CreatePortParams) WithHTTPClient(client *http.Client) *CreatePortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create port params
func (o *CreatePortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPort adds the port to the create port params
func (o *CreatePortParams) WithPort(port *models.Port) *CreatePortParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the create port params
func (o *CreatePortParams) SetPort(port *models.Port) {
	o.Port = port
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Port != nil {
		if err := r.SetBodyParam(o.Port); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package packet_analyzers

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

// NewDeletePacketAnalyzerResultParams creates a new DeletePacketAnalyzerResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePacketAnalyzerResultParams() *DeletePacketAnalyzerResultParams {
	return &DeletePacketAnalyzerResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePacketAnalyzerResultParamsWithTimeout creates a new DeletePacketAnalyzerResultParams object
// with the ability to set a timeout on a request.
func NewDeletePacketAnalyzerResultParamsWithTimeout(timeout time.Duration) *DeletePacketAnalyzerResultParams {
	return &DeletePacketAnalyzerResultParams{
		timeout: timeout,
	}
}

// NewDeletePacketAnalyzerResultParamsWithContext creates a new DeletePacketAnalyzerResultParams object
// with the ability to set a context for a request.
func NewDeletePacketAnalyzerResultParamsWithContext(ctx context.Context) *DeletePacketAnalyzerResultParams {
	return &DeletePacketAnalyzerResultParams{
		Context: ctx,
	}
}

// NewDeletePacketAnalyzerResultParamsWithHTTPClient creates a new DeletePacketAnalyzerResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePacketAnalyzerResultParamsWithHTTPClient(client *http.Client) *DeletePacketAnalyzerResultParams {
	return &DeletePacketAnalyzerResultParams{
		HTTPClient: client,
	}
}

/* DeletePacketAnalyzerResultParams contains all the parameters to send to the API endpoint
   for the delete packet analyzer result operation.

   Typically these are written to a http.Request.
*/
type DeletePacketAnalyzerResultParams struct {

	/* ID.

	   Unique resource identifier

	   Format: string
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete packet analyzer result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePacketAnalyzerResultParams) WithDefaults() *DeletePacketAnalyzerResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete packet analyzer result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePacketAnalyzerResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) WithTimeout(timeout time.Duration) *DeletePacketAnalyzerResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) WithContext(ctx context.Context) *DeletePacketAnalyzerResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) WithHTTPClient(client *http.Client) *DeletePacketAnalyzerResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) WithID(id string) *DeletePacketAnalyzerResultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete packet analyzer result params
func (o *DeletePacketAnalyzerResultParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePacketAnalyzerResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

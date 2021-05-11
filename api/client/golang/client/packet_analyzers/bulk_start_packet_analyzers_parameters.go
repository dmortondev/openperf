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

// NewBulkStartPacketAnalyzersParams creates a new BulkStartPacketAnalyzersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkStartPacketAnalyzersParams() *BulkStartPacketAnalyzersParams {
	return &BulkStartPacketAnalyzersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkStartPacketAnalyzersParamsWithTimeout creates a new BulkStartPacketAnalyzersParams object
// with the ability to set a timeout on a request.
func NewBulkStartPacketAnalyzersParamsWithTimeout(timeout time.Duration) *BulkStartPacketAnalyzersParams {
	return &BulkStartPacketAnalyzersParams{
		timeout: timeout,
	}
}

// NewBulkStartPacketAnalyzersParamsWithContext creates a new BulkStartPacketAnalyzersParams object
// with the ability to set a context for a request.
func NewBulkStartPacketAnalyzersParamsWithContext(ctx context.Context) *BulkStartPacketAnalyzersParams {
	return &BulkStartPacketAnalyzersParams{
		Context: ctx,
	}
}

// NewBulkStartPacketAnalyzersParamsWithHTTPClient creates a new BulkStartPacketAnalyzersParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkStartPacketAnalyzersParamsWithHTTPClient(client *http.Client) *BulkStartPacketAnalyzersParams {
	return &BulkStartPacketAnalyzersParams{
		HTTPClient: client,
	}
}

/* BulkStartPacketAnalyzersParams contains all the parameters to send to the API endpoint
   for the bulk start packet analyzers operation.

   Typically these are written to a http.Request.
*/
type BulkStartPacketAnalyzersParams struct {

	/* Start.

	   Bulk start
	*/
	Start BulkStartPacketAnalyzersBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk start packet analyzers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkStartPacketAnalyzersParams) WithDefaults() *BulkStartPacketAnalyzersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk start packet analyzers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkStartPacketAnalyzersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) WithTimeout(timeout time.Duration) *BulkStartPacketAnalyzersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) WithContext(ctx context.Context) *BulkStartPacketAnalyzersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) WithHTTPClient(client *http.Client) *BulkStartPacketAnalyzersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStart adds the start to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) WithStart(start BulkStartPacketAnalyzersBody) *BulkStartPacketAnalyzersParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the bulk start packet analyzers params
func (o *BulkStartPacketAnalyzersParams) SetStart(start BulkStartPacketAnalyzersBody) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *BulkStartPacketAnalyzersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Start); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

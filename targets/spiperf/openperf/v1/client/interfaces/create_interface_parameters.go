// Code generated by go-swagger; DO NOT EDIT.

package interfaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Spirent/openperf/targets/spiperf/openperf/v1/models"
)

// NewCreateInterfaceParams creates a new CreateInterfaceParams object
// with the default values initialized.
func NewCreateInterfaceParams() *CreateInterfaceParams {
	var ()
	return &CreateInterfaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInterfaceParamsWithTimeout creates a new CreateInterfaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInterfaceParamsWithTimeout(timeout time.Duration) *CreateInterfaceParams {
	var ()
	return &CreateInterfaceParams{

		timeout: timeout,
	}
}

// NewCreateInterfaceParamsWithContext creates a new CreateInterfaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInterfaceParamsWithContext(ctx context.Context) *CreateInterfaceParams {
	var ()
	return &CreateInterfaceParams{

		Context: ctx,
	}
}

// NewCreateInterfaceParamsWithHTTPClient creates a new CreateInterfaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInterfaceParamsWithHTTPClient(client *http.Client) *CreateInterfaceParams {
	var ()
	return &CreateInterfaceParams{
		HTTPClient: client,
	}
}

/*CreateInterfaceParams contains all the parameters to send to the API endpoint
for the create interface operation typically these are written to a http.Request
*/
type CreateInterfaceParams struct {

	/*Interface
	  New network interface

	*/
	Interface *models.Interface

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create interface params
func (o *CreateInterfaceParams) WithTimeout(timeout time.Duration) *CreateInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create interface params
func (o *CreateInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create interface params
func (o *CreateInterfaceParams) WithContext(ctx context.Context) *CreateInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create interface params
func (o *CreateInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create interface params
func (o *CreateInterfaceParams) WithHTTPClient(client *http.Client) *CreateInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create interface params
func (o *CreateInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterface adds the interfaceVar to the create interface params
func (o *CreateInterfaceParams) WithInterface(interfaceVar *models.Interface) *CreateInterfaceParams {
	o.SetInterface(interfaceVar)
	return o
}

// SetInterface adds the interface to the create interface params
func (o *CreateInterfaceParams) SetInterface(interfaceVar *models.Interface) {
	o.Interface = interfaceVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Interface != nil {
		if err := r.SetBodyParam(o.Interface); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

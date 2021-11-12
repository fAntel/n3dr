// Code generated by go-swagger; DO NOT EDIT.

package security_management_realms

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

// NewGetRealmsParams creates a new GetRealmsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRealmsParams() *GetRealmsParams {
	return &GetRealmsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRealmsParamsWithTimeout creates a new GetRealmsParams object
// with the ability to set a timeout on a request.
func NewGetRealmsParamsWithTimeout(timeout time.Duration) *GetRealmsParams {
	return &GetRealmsParams{
		timeout: timeout,
	}
}

// NewGetRealmsParamsWithContext creates a new GetRealmsParams object
// with the ability to set a context for a request.
func NewGetRealmsParamsWithContext(ctx context.Context) *GetRealmsParams {
	return &GetRealmsParams{
		Context: ctx,
	}
}

// NewGetRealmsParamsWithHTTPClient creates a new GetRealmsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRealmsParamsWithHTTPClient(client *http.Client) *GetRealmsParams {
	return &GetRealmsParams{
		HTTPClient: client,
	}
}

/* GetRealmsParams contains all the parameters to send to the API endpoint
   for the get realms operation.

   Typically these are written to a http.Request.
*/
type GetRealmsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get realms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRealmsParams) WithDefaults() *GetRealmsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get realms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRealmsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get realms params
func (o *GetRealmsParams) WithTimeout(timeout time.Duration) *GetRealmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get realms params
func (o *GetRealmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get realms params
func (o *GetRealmsParams) WithContext(ctx context.Context) *GetRealmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get realms params
func (o *GetRealmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get realms params
func (o *GetRealmsParams) WithHTTPClient(client *http.Client) *GetRealmsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get realms params
func (o *GetRealmsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRealmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

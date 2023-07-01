// Code generated by go-swagger; DO NOT EDIT.

package product_licensing

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

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// NewSetLicenseParams creates a new SetLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetLicenseParams() *SetLicenseParams {
	return &SetLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetLicenseParamsWithTimeout creates a new SetLicenseParams object
// with the ability to set a timeout on a request.
func NewSetLicenseParamsWithTimeout(timeout time.Duration) *SetLicenseParams {
	return &SetLicenseParams{
		timeout: timeout,
	}
}

// NewSetLicenseParamsWithContext creates a new SetLicenseParams object
// with the ability to set a context for a request.
func NewSetLicenseParamsWithContext(ctx context.Context) *SetLicenseParams {
	return &SetLicenseParams{
		Context: ctx,
	}
}

// NewSetLicenseParamsWithHTTPClient creates a new SetLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetLicenseParamsWithHTTPClient(client *http.Client) *SetLicenseParams {
	return &SetLicenseParams{
		HTTPClient: client,
	}
}

/*
SetLicenseParams contains all the parameters to send to the API endpoint

	for the set license operation.

	Typically these are written to a http.Request.
*/
type SetLicenseParams struct {

	// Body.
	Body models.InputStream

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetLicenseParams) WithDefaults() *SetLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set license params
func (o *SetLicenseParams) WithTimeout(timeout time.Duration) *SetLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set license params
func (o *SetLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set license params
func (o *SetLicenseParams) WithContext(ctx context.Context) *SetLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set license params
func (o *SetLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set license params
func (o *SetLicenseParams) WithHTTPClient(client *http.Client) *SetLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set license params
func (o *SetLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set license params
func (o *SetLicenseParams) WithBody(body models.InputStream) *SetLicenseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set license params
func (o *SetLicenseParams) SetBody(body models.InputStream) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

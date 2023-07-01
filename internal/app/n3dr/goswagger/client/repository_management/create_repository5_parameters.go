// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

// NewCreateRepository5Params creates a new CreateRepository5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository5Params() *CreateRepository5Params {
	return &CreateRepository5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository5ParamsWithTimeout creates a new CreateRepository5Params object
// with the ability to set a timeout on a request.
func NewCreateRepository5ParamsWithTimeout(timeout time.Duration) *CreateRepository5Params {
	return &CreateRepository5Params{
		timeout: timeout,
	}
}

// NewCreateRepository5ParamsWithContext creates a new CreateRepository5Params object
// with the ability to set a context for a request.
func NewCreateRepository5ParamsWithContext(ctx context.Context) *CreateRepository5Params {
	return &CreateRepository5Params{
		Context: ctx,
	}
}

// NewCreateRepository5ParamsWithHTTPClient creates a new CreateRepository5Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository5ParamsWithHTTPClient(client *http.Client) *CreateRepository5Params {
	return &CreateRepository5Params{
		HTTPClient: client,
	}
}

/*
CreateRepository5Params contains all the parameters to send to the API endpoint

	for the create repository 5 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository5Params struct {

	// Body.
	Body *models.RawGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository5Params) WithDefaults() *CreateRepository5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 5 params
func (o *CreateRepository5Params) WithTimeout(timeout time.Duration) *CreateRepository5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 5 params
func (o *CreateRepository5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 5 params
func (o *CreateRepository5Params) WithContext(ctx context.Context) *CreateRepository5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 5 params
func (o *CreateRepository5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 5 params
func (o *CreateRepository5Params) WithHTTPClient(client *http.Client) *CreateRepository5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 5 params
func (o *CreateRepository5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 5 params
func (o *CreateRepository5Params) WithBody(body *models.RawGroupRepositoryAPIRequest) *CreateRepository5Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 5 params
func (o *CreateRepository5Params) SetBody(body *models.RawGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

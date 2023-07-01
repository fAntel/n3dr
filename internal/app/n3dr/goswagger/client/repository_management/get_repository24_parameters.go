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
)

// NewGetRepository24Params creates a new GetRepository24Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepository24Params() *GetRepository24Params {
	return &GetRepository24Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepository24ParamsWithTimeout creates a new GetRepository24Params object
// with the ability to set a timeout on a request.
func NewGetRepository24ParamsWithTimeout(timeout time.Duration) *GetRepository24Params {
	return &GetRepository24Params{
		timeout: timeout,
	}
}

// NewGetRepository24ParamsWithContext creates a new GetRepository24Params object
// with the ability to set a context for a request.
func NewGetRepository24ParamsWithContext(ctx context.Context) *GetRepository24Params {
	return &GetRepository24Params{
		Context: ctx,
	}
}

// NewGetRepository24ParamsWithHTTPClient creates a new GetRepository24Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepository24ParamsWithHTTPClient(client *http.Client) *GetRepository24Params {
	return &GetRepository24Params{
		HTTPClient: client,
	}
}

/*
GetRepository24Params contains all the parameters to send to the API endpoint

	for the get repository 24 operation.

	Typically these are written to a http.Request.
*/
type GetRepository24Params struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 24 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository24Params) WithDefaults() *GetRepository24Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 24 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository24Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 24 params
func (o *GetRepository24Params) WithTimeout(timeout time.Duration) *GetRepository24Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 24 params
func (o *GetRepository24Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 24 params
func (o *GetRepository24Params) WithContext(ctx context.Context) *GetRepository24Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 24 params
func (o *GetRepository24Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 24 params
func (o *GetRepository24Params) WithHTTPClient(client *http.Client) *GetRepository24Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 24 params
func (o *GetRepository24Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 24 params
func (o *GetRepository24Params) WithRepositoryName(repositoryName string) *GetRepository24Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 24 params
func (o *GetRepository24Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepository24Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

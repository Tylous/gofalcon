// Code generated by go-swagger; DO NOT EDIT.

package quarantine

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpdateQfByQueryParams creates a new UpdateQfByQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateQfByQueryParams() *UpdateQfByQueryParams {
	return &UpdateQfByQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateQfByQueryParamsWithTimeout creates a new UpdateQfByQueryParams object
// with the ability to set a timeout on a request.
func NewUpdateQfByQueryParamsWithTimeout(timeout time.Duration) *UpdateQfByQueryParams {
	return &UpdateQfByQueryParams{
		timeout: timeout,
	}
}

// NewUpdateQfByQueryParamsWithContext creates a new UpdateQfByQueryParams object
// with the ability to set a context for a request.
func NewUpdateQfByQueryParamsWithContext(ctx context.Context) *UpdateQfByQueryParams {
	return &UpdateQfByQueryParams{
		Context: ctx,
	}
}

// NewUpdateQfByQueryParamsWithHTTPClient creates a new UpdateQfByQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateQfByQueryParamsWithHTTPClient(client *http.Client) *UpdateQfByQueryParams {
	return &UpdateQfByQueryParams{
		HTTPClient: client,
	}
}

/* UpdateQfByQueryParams contains all the parameters to send to the API endpoint
   for the update qf by query operation.

   Typically these are written to a http.Request.
*/
type UpdateQfByQueryParams struct {

	// Body.
	Body *models.DomainQueriesPatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update qf by query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQfByQueryParams) WithDefaults() *UpdateQfByQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update qf by query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQfByQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update qf by query params
func (o *UpdateQfByQueryParams) WithTimeout(timeout time.Duration) *UpdateQfByQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update qf by query params
func (o *UpdateQfByQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update qf by query params
func (o *UpdateQfByQueryParams) WithContext(ctx context.Context) *UpdateQfByQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update qf by query params
func (o *UpdateQfByQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update qf by query params
func (o *UpdateQfByQueryParams) WithHTTPClient(client *http.Client) *UpdateQfByQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update qf by query params
func (o *UpdateQfByQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update qf by query params
func (o *UpdateQfByQueryParams) WithBody(body *models.DomainQueriesPatchRequest) *UpdateQfByQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update qf by query params
func (o *UpdateQfByQueryParams) SetBody(body *models.DomainQueriesPatchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateQfByQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

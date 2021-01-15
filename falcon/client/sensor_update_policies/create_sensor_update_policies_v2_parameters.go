// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// NewCreateSensorUpdatePoliciesV2Params creates a new CreateSensorUpdatePoliciesV2Params object
// with the default values initialized.
func NewCreateSensorUpdatePoliciesV2Params() *CreateSensorUpdatePoliciesV2Params {
	var ()
	return &CreateSensorUpdatePoliciesV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSensorUpdatePoliciesV2ParamsWithTimeout creates a new CreateSensorUpdatePoliciesV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSensorUpdatePoliciesV2ParamsWithTimeout(timeout time.Duration) *CreateSensorUpdatePoliciesV2Params {
	var ()
	return &CreateSensorUpdatePoliciesV2Params{

		timeout: timeout,
	}
}

// NewCreateSensorUpdatePoliciesV2ParamsWithContext creates a new CreateSensorUpdatePoliciesV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSensorUpdatePoliciesV2ParamsWithContext(ctx context.Context) *CreateSensorUpdatePoliciesV2Params {
	var ()
	return &CreateSensorUpdatePoliciesV2Params{

		Context: ctx,
	}
}

// NewCreateSensorUpdatePoliciesV2ParamsWithHTTPClient creates a new CreateSensorUpdatePoliciesV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSensorUpdatePoliciesV2ParamsWithHTTPClient(client *http.Client) *CreateSensorUpdatePoliciesV2Params {
	var ()
	return &CreateSensorUpdatePoliciesV2Params{
		HTTPClient: client,
	}
}

/*CreateSensorUpdatePoliciesV2Params contains all the parameters to send to the API endpoint
for the create sensor update policies v2 operation typically these are written to a http.Request
*/
type CreateSensorUpdatePoliciesV2Params struct {

	/*Body*/
	Body *models.RequestsCreateSensorUpdatePoliciesV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) WithTimeout(timeout time.Duration) *CreateSensorUpdatePoliciesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) WithContext(ctx context.Context) *CreateSensorUpdatePoliciesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) WithHTTPClient(client *http.Client) *CreateSensorUpdatePoliciesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) WithBody(body *models.RequestsCreateSensorUpdatePoliciesV2) *CreateSensorUpdatePoliciesV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create sensor update policies v2 params
func (o *CreateSensorUpdatePoliciesV2Params) SetBody(body *models.RequestsCreateSensorUpdatePoliciesV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSensorUpdatePoliciesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
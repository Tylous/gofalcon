// Code generated by go-swagger; DO NOT EDIT.

package sensor_visibility_exclusions

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

// NewCreateSVExclusionsV1Params creates a new CreateSVExclusionsV1Params object
// with the default values initialized.
func NewCreateSVExclusionsV1Params() *CreateSVExclusionsV1Params {
	var ()
	return &CreateSVExclusionsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSVExclusionsV1ParamsWithTimeout creates a new CreateSVExclusionsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSVExclusionsV1ParamsWithTimeout(timeout time.Duration) *CreateSVExclusionsV1Params {
	var ()
	return &CreateSVExclusionsV1Params{

		timeout: timeout,
	}
}

// NewCreateSVExclusionsV1ParamsWithContext creates a new CreateSVExclusionsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSVExclusionsV1ParamsWithContext(ctx context.Context) *CreateSVExclusionsV1Params {
	var ()
	return &CreateSVExclusionsV1Params{

		Context: ctx,
	}
}

// NewCreateSVExclusionsV1ParamsWithHTTPClient creates a new CreateSVExclusionsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSVExclusionsV1ParamsWithHTTPClient(client *http.Client) *CreateSVExclusionsV1Params {
	var ()
	return &CreateSVExclusionsV1Params{
		HTTPClient: client,
	}
}

/*CreateSVExclusionsV1Params contains all the parameters to send to the API endpoint
for the create s v exclusions v1 operation typically these are written to a http.Request
*/
type CreateSVExclusionsV1Params struct {

	/*Body*/
	Body *models.RequestsSvExclusionCreateReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) WithTimeout(timeout time.Duration) *CreateSVExclusionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) WithContext(ctx context.Context) *CreateSVExclusionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) WithHTTPClient(client *http.Client) *CreateSVExclusionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) WithBody(body *models.RequestsSvExclusionCreateReqV1) *CreateSVExclusionsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create s v exclusions v1 params
func (o *CreateSVExclusionsV1Params) SetBody(body *models.RequestsSvExclusionCreateReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSVExclusionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
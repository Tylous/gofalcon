// Code generated by go-swagger; DO NOT EDIT.

package message_center

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

// NewCaseAddAttachmentParams creates a new CaseAddAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCaseAddAttachmentParams() *CaseAddAttachmentParams {
	return &CaseAddAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCaseAddAttachmentParamsWithTimeout creates a new CaseAddAttachmentParams object
// with the ability to set a timeout on a request.
func NewCaseAddAttachmentParamsWithTimeout(timeout time.Duration) *CaseAddAttachmentParams {
	return &CaseAddAttachmentParams{
		timeout: timeout,
	}
}

// NewCaseAddAttachmentParamsWithContext creates a new CaseAddAttachmentParams object
// with the ability to set a context for a request.
func NewCaseAddAttachmentParamsWithContext(ctx context.Context) *CaseAddAttachmentParams {
	return &CaseAddAttachmentParams{
		Context: ctx,
	}
}

// NewCaseAddAttachmentParamsWithHTTPClient creates a new CaseAddAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCaseAddAttachmentParamsWithHTTPClient(client *http.Client) *CaseAddAttachmentParams {
	return &CaseAddAttachmentParams{
		HTTPClient: client,
	}
}

/* CaseAddAttachmentParams contains all the parameters to send to the API endpoint
   for the case add attachment operation.

   Typically these are written to a http.Request.
*/
type CaseAddAttachmentParams struct {

	/* CaseID.

	   Case ID
	*/
	CaseID string

	/* File.

	   File Body
	*/
	File runtime.NamedReadCloser

	/* UserUUID.

	   User UUID
	*/
	UserUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the case add attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CaseAddAttachmentParams) WithDefaults() *CaseAddAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the case add attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CaseAddAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the case add attachment params
func (o *CaseAddAttachmentParams) WithTimeout(timeout time.Duration) *CaseAddAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the case add attachment params
func (o *CaseAddAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the case add attachment params
func (o *CaseAddAttachmentParams) WithContext(ctx context.Context) *CaseAddAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the case add attachment params
func (o *CaseAddAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the case add attachment params
func (o *CaseAddAttachmentParams) WithHTTPClient(client *http.Client) *CaseAddAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the case add attachment params
func (o *CaseAddAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaseID adds the caseID to the case add attachment params
func (o *CaseAddAttachmentParams) WithCaseID(caseID string) *CaseAddAttachmentParams {
	o.SetCaseID(caseID)
	return o
}

// SetCaseID adds the caseId to the case add attachment params
func (o *CaseAddAttachmentParams) SetCaseID(caseID string) {
	o.CaseID = caseID
}

// WithFile adds the file to the case add attachment params
func (o *CaseAddAttachmentParams) WithFile(file runtime.NamedReadCloser) *CaseAddAttachmentParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the case add attachment params
func (o *CaseAddAttachmentParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithUserUUID adds the userUUID to the case add attachment params
func (o *CaseAddAttachmentParams) WithUserUUID(userUUID string) *CaseAddAttachmentParams {
	o.SetUserUUID(userUUID)
	return o
}

// SetUserUUID adds the userUuid to the case add attachment params
func (o *CaseAddAttachmentParams) SetUserUUID(userUUID string) {
	o.UserUUID = userUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CaseAddAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param case_id
	frCaseID := o.CaseID
	fCaseID := frCaseID
	if fCaseID != "" {
		if err := r.SetFormParam("case_id", fCaseID); err != nil {
			return err
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// form param user_uuid
	frUserUUID := o.UserUUID
	fUserUUID := frUserUUID
	if fUserUUID != "" {
		if err := r.SetFormParam("user_uuid", fUserUUID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

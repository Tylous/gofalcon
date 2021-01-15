// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// RTRCheckAdminCommandStatusReader is a Reader for the RTRCheckAdminCommandStatus structure.
type RTRCheckAdminCommandStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRCheckAdminCommandStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRCheckAdminCommandStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRTRCheckAdminCommandStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRCheckAdminCommandStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRCheckAdminCommandStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRCheckAdminCommandStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRCheckAdminCommandStatusOK creates a RTRCheckAdminCommandStatusOK with default headers values
func NewRTRCheckAdminCommandStatusOK() *RTRCheckAdminCommandStatusOK {
	return &RTRCheckAdminCommandStatusOK{}
}

/*RTRCheckAdminCommandStatusOK handles this case with default header values.

success
*/
type RTRCheckAdminCommandStatusOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainStatusResponseWrapper
}

func (o *RTRCheckAdminCommandStatusOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusOK  %+v", 200, o.Payload)
}

func (o *RTRCheckAdminCommandStatusOK) GetPayload() *models.DomainStatusResponseWrapper {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.DomainStatusResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckAdminCommandStatusUnauthorized creates a RTRCheckAdminCommandStatusUnauthorized with default headers values
func NewRTRCheckAdminCommandStatusUnauthorized() *RTRCheckAdminCommandStatusUnauthorized {
	return &RTRCheckAdminCommandStatusUnauthorized{}
}

/*RTRCheckAdminCommandStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type RTRCheckAdminCommandStatusUnauthorized struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRCheckAdminCommandStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRCheckAdminCommandStatusUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckAdminCommandStatusForbidden creates a RTRCheckAdminCommandStatusForbidden with default headers values
func NewRTRCheckAdminCommandStatusForbidden() *RTRCheckAdminCommandStatusForbidden {
	return &RTRCheckAdminCommandStatusForbidden{}
}

/*RTRCheckAdminCommandStatusForbidden handles this case with default header values.

Forbidden
*/
type RTRCheckAdminCommandStatusForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRCheckAdminCommandStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusForbidden  %+v", 403, o.Payload)
}

func (o *RTRCheckAdminCommandStatusForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckAdminCommandStatusTooManyRequests creates a RTRCheckAdminCommandStatusTooManyRequests with default headers values
func NewRTRCheckAdminCommandStatusTooManyRequests() *RTRCheckAdminCommandStatusTooManyRequests {
	return &RTRCheckAdminCommandStatusTooManyRequests{}
}

/*RTRCheckAdminCommandStatusTooManyRequests handles this case with default header values.

Too Many Requests
*/
type RTRCheckAdminCommandStatusTooManyRequests struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
	/*Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-RetryAfter
	xRateLimitRetryAfter, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-RetryAfter"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", response.GetHeader("X-RateLimit-RetryAfter"))
	}
	o.XRateLimitRetryAfter = xRateLimitRetryAfter

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckAdminCommandStatusDefault creates a RTRCheckAdminCommandStatusDefault with default headers values
func NewRTRCheckAdminCommandStatusDefault(code int) *RTRCheckAdminCommandStatusDefault {
	return &RTRCheckAdminCommandStatusDefault{
		_statusCode: code,
	}
}

/*RTRCheckAdminCommandStatusDefault handles this case with default header values.

success
*/
type RTRCheckAdminCommandStatusDefault struct {
	_statusCode int

	Payload *models.DomainStatusResponseWrapper
}

// Code gets the status code for the r t r check admin command status default response
func (o *RTRCheckAdminCommandStatusDefault) Code() int {
	return o._statusCode
}

func (o *RTRCheckAdminCommandStatusDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] RTR-CheckAdminCommandStatus default  %+v", o._statusCode, o.Payload)
}

func (o *RTRCheckAdminCommandStatusDefault) GetPayload() *models.DomainStatusResponseWrapper {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainStatusResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// PerformFirewallPoliciesActionReader is a Reader for the PerformFirewallPoliciesAction structure.
type PerformFirewallPoliciesActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformFirewallPoliciesActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformFirewallPoliciesActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformFirewallPoliciesActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformFirewallPoliciesActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformFirewallPoliciesActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPerformFirewallPoliciesActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformFirewallPoliciesActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformFirewallPoliciesActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformFirewallPoliciesActionOK creates a PerformFirewallPoliciesActionOK with default headers values
func NewPerformFirewallPoliciesActionOK() *PerformFirewallPoliciesActionOK {
	return &PerformFirewallPoliciesActionOK{}
}

/*PerformFirewallPoliciesActionOK handles this case with default header values.

OK
*/
type PerformFirewallPoliciesActionOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *PerformFirewallPoliciesActionOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformFirewallPoliciesActionOK) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionBadRequest creates a PerformFirewallPoliciesActionBadRequest with default headers values
func NewPerformFirewallPoliciesActionBadRequest() *PerformFirewallPoliciesActionBadRequest {
	return &PerformFirewallPoliciesActionBadRequest{}
}

/*PerformFirewallPoliciesActionBadRequest handles this case with default header values.

Bad Request
*/
type PerformFirewallPoliciesActionBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *PerformFirewallPoliciesActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformFirewallPoliciesActionBadRequest) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionForbidden creates a PerformFirewallPoliciesActionForbidden with default headers values
func NewPerformFirewallPoliciesActionForbidden() *PerformFirewallPoliciesActionForbidden {
	return &PerformFirewallPoliciesActionForbidden{}
}

/*PerformFirewallPoliciesActionForbidden handles this case with default header values.

Forbidden
*/
type PerformFirewallPoliciesActionForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *PerformFirewallPoliciesActionForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformFirewallPoliciesActionForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionNotFound creates a PerformFirewallPoliciesActionNotFound with default headers values
func NewPerformFirewallPoliciesActionNotFound() *PerformFirewallPoliciesActionNotFound {
	return &PerformFirewallPoliciesActionNotFound{}
}

/*PerformFirewallPoliciesActionNotFound handles this case with default header values.

Not Found
*/
type PerformFirewallPoliciesActionNotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *PerformFirewallPoliciesActionNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformFirewallPoliciesActionNotFound) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionTooManyRequests creates a PerformFirewallPoliciesActionTooManyRequests with default headers values
func NewPerformFirewallPoliciesActionTooManyRequests() *PerformFirewallPoliciesActionTooManyRequests {
	return &PerformFirewallPoliciesActionTooManyRequests{}
}

/*PerformFirewallPoliciesActionTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PerformFirewallPoliciesActionTooManyRequests struct {
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

func (o *PerformFirewallPoliciesActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformFirewallPoliciesActionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPerformFirewallPoliciesActionInternalServerError creates a PerformFirewallPoliciesActionInternalServerError with default headers values
func NewPerformFirewallPoliciesActionInternalServerError() *PerformFirewallPoliciesActionInternalServerError {
	return &PerformFirewallPoliciesActionInternalServerError{}
}

/*PerformFirewallPoliciesActionInternalServerError handles this case with default header values.

Internal Server Error
*/
type PerformFirewallPoliciesActionInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *PerformFirewallPoliciesActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformFirewallPoliciesActionInternalServerError) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformFirewallPoliciesActionDefault creates a PerformFirewallPoliciesActionDefault with default headers values
func NewPerformFirewallPoliciesActionDefault(code int) *PerformFirewallPoliciesActionDefault {
	return &PerformFirewallPoliciesActionDefault{
		_statusCode: code,
	}
}

/*PerformFirewallPoliciesActionDefault handles this case with default header values.

OK
*/
type PerformFirewallPoliciesActionDefault struct {
	_statusCode int

	Payload *models.ResponsesFirewallPoliciesV1
}

// Code gets the status code for the perform firewall policies action default response
func (o *PerformFirewallPoliciesActionDefault) Code() int {
	return o._statusCode
}

func (o *PerformFirewallPoliciesActionDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-actions/v1][%d] performFirewallPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformFirewallPoliciesActionDefault) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *PerformFirewallPoliciesActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
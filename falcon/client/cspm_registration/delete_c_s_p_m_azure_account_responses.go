// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// DeleteCSPMAzureAccountReader is a Reader for the DeleteCSPMAzureAccount structure.
type DeleteCSPMAzureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCSPMAzureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCSPMAzureAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteCSPMAzureAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCSPMAzureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCSPMAzureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCSPMAzureAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCSPMAzureAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCSPMAzureAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCSPMAzureAccountOK creates a DeleteCSPMAzureAccountOK with default headers values
func NewDeleteCSPMAzureAccountOK() *DeleteCSPMAzureAccountOK {
	return &DeleteCSPMAzureAccountOK{}
}

/*DeleteCSPMAzureAccountOK handles this case with default header values.

OK
*/
type DeleteCSPMAzureAccountOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationBaseResponseV1
}

func (o *DeleteCSPMAzureAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMAzureAccountOK) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountMultiStatus creates a DeleteCSPMAzureAccountMultiStatus with default headers values
func NewDeleteCSPMAzureAccountMultiStatus() *DeleteCSPMAzureAccountMultiStatus {
	return &DeleteCSPMAzureAccountMultiStatus{}
}

/*DeleteCSPMAzureAccountMultiStatus handles this case with default header values.

Multi-Status
*/
type DeleteCSPMAzureAccountMultiStatus struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationBaseResponseV1
}

func (o *DeleteCSPMAzureAccountMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMAzureAccountMultiStatus) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountBadRequest creates a DeleteCSPMAzureAccountBadRequest with default headers values
func NewDeleteCSPMAzureAccountBadRequest() *DeleteCSPMAzureAccountBadRequest {
	return &DeleteCSPMAzureAccountBadRequest{}
}

/*DeleteCSPMAzureAccountBadRequest handles this case with default header values.

Bad Request
*/
type DeleteCSPMAzureAccountBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationBaseResponseV1
}

func (o *DeleteCSPMAzureAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMAzureAccountBadRequest) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountForbidden creates a DeleteCSPMAzureAccountForbidden with default headers values
func NewDeleteCSPMAzureAccountForbidden() *DeleteCSPMAzureAccountForbidden {
	return &DeleteCSPMAzureAccountForbidden{}
}

/*DeleteCSPMAzureAccountForbidden handles this case with default header values.

Forbidden
*/
type DeleteCSPMAzureAccountForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DeleteCSPMAzureAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMAzureAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMAzureAccountTooManyRequests creates a DeleteCSPMAzureAccountTooManyRequests with default headers values
func NewDeleteCSPMAzureAccountTooManyRequests() *DeleteCSPMAzureAccountTooManyRequests {
	return &DeleteCSPMAzureAccountTooManyRequests{}
}

/*DeleteCSPMAzureAccountTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteCSPMAzureAccountTooManyRequests struct {
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

func (o *DeleteCSPMAzureAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMAzureAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMAzureAccountInternalServerError creates a DeleteCSPMAzureAccountInternalServerError with default headers values
func NewDeleteCSPMAzureAccountInternalServerError() *DeleteCSPMAzureAccountInternalServerError {
	return &DeleteCSPMAzureAccountInternalServerError{}
}

/*DeleteCSPMAzureAccountInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteCSPMAzureAccountInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

func (o *DeleteCSPMAzureAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMAzureAccountInternalServerError) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountDefault creates a DeleteCSPMAzureAccountDefault with default headers values
func NewDeleteCSPMAzureAccountDefault(code int) *DeleteCSPMAzureAccountDefault {
	return &DeleteCSPMAzureAccountDefault{
		_statusCode: code,
	}
}

/*DeleteCSPMAzureAccountDefault handles this case with default header values.

OK
*/
type DeleteCSPMAzureAccountDefault struct {
	_statusCode int

	Payload *models.RegistrationBaseResponseV1
}

// Code gets the status code for the delete c s p m azure account default response
func (o *DeleteCSPMAzureAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCSPMAzureAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] DeleteCSPMAzureAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCSPMAzureAccountDefault) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
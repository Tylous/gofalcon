// Code generated by go-swagger; DO NOT EDIT.

package host_group

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

// QueryGroupMembersReader is a Reader for the QueryGroupMembers structure.
type QueryGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryGroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryGroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryGroupMembersOK creates a QueryGroupMembersOK with default headers values
func NewQueryGroupMembersOK() *QueryGroupMembersOK {
	return &QueryGroupMembersOK{}
}

/*QueryGroupMembersOK handles this case with default header values.

OK
*/
type QueryGroupMembersOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryGroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryGroupMembersOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGroupMembersBadRequest creates a QueryGroupMembersBadRequest with default headers values
func NewQueryGroupMembersBadRequest() *QueryGroupMembersBadRequest {
	return &QueryGroupMembersBadRequest{}
}

/*QueryGroupMembersBadRequest handles this case with default header values.

Bad Request
*/
type QueryGroupMembersBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryGroupMembersBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGroupMembersForbidden creates a QueryGroupMembersForbidden with default headers values
func NewQueryGroupMembersForbidden() *QueryGroupMembersForbidden {
	return &QueryGroupMembersForbidden{}
}

/*QueryGroupMembersForbidden handles this case with default header values.

Forbidden
*/
type QueryGroupMembersForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersNotFound creates a QueryGroupMembersNotFound with default headers values
func NewQueryGroupMembersNotFound() *QueryGroupMembersNotFound {
	return &QueryGroupMembersNotFound{}
}

/*QueryGroupMembersNotFound handles this case with default header values.

Not Found
*/
type QueryGroupMembersNotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryGroupMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryGroupMembersNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGroupMembersTooManyRequests creates a QueryGroupMembersTooManyRequests with default headers values
func NewQueryGroupMembersTooManyRequests() *QueryGroupMembersTooManyRequests {
	return &QueryGroupMembersTooManyRequests{}
}

/*QueryGroupMembersTooManyRequests handles this case with default header values.

Too Many Requests
*/
type QueryGroupMembersTooManyRequests struct {
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

func (o *QueryGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGroupMembersInternalServerError creates a QueryGroupMembersInternalServerError with default headers values
func NewQueryGroupMembersInternalServerError() *QueryGroupMembersInternalServerError {
	return &QueryGroupMembersInternalServerError{}
}

/*QueryGroupMembersInternalServerError handles this case with default header values.

Internal Server Error
*/
type QueryGroupMembersInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryGroupMembersInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGroupMembersDefault creates a QueryGroupMembersDefault with default headers values
func NewQueryGroupMembersDefault(code int) *QueryGroupMembersDefault {
	return &QueryGroupMembersDefault{
		_statusCode: code,
	}
}

/*QueryGroupMembersDefault handles this case with default header values.

OK
*/
type QueryGroupMembersDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query group members default response
func (o *QueryGroupMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryGroupMembersDefault) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-group-members/v1][%d] queryGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryGroupMembersDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
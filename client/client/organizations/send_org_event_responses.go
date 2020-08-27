// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// SendOrgEventReader is a Reader for the SendOrgEvent structure.
type SendOrgEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendOrgEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendOrgEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSendOrgEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSendOrgEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSendOrgEventUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendOrgEventOK creates a SendOrgEventOK with default headers values
func NewSendOrgEventOK() *SendOrgEventOK {
	return &SendOrgEventOK{}
}

/*SendOrgEventOK handles this case with default header values.

Event has been registered
*/
type SendOrgEventOK struct {
}

func (o *SendOrgEventOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendOrgEventOK ", 200)
}

func (o *SendOrgEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendOrgEventForbidden creates a SendOrgEventForbidden with default headers values
func NewSendOrgEventForbidden() *SendOrgEventForbidden {
	return &SendOrgEventForbidden{}
}

/*SendOrgEventForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type SendOrgEventForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *SendOrgEventForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendOrgEventForbidden  %+v", 403, o.Payload)
}

func (o *SendOrgEventForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *SendOrgEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendOrgEventNotFound creates a SendOrgEventNotFound with default headers values
func NewSendOrgEventNotFound() *SendOrgEventNotFound {
	return &SendOrgEventNotFound{}
}

/*SendOrgEventNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type SendOrgEventNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *SendOrgEventNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendOrgEventNotFound  %+v", 404, o.Payload)
}

func (o *SendOrgEventNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *SendOrgEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendOrgEventUnprocessableEntity creates a SendOrgEventUnprocessableEntity with default headers values
func NewSendOrgEventUnprocessableEntity() *SendOrgEventUnprocessableEntity {
	return &SendOrgEventUnprocessableEntity{}
}

/*SendOrgEventUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type SendOrgEventUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *SendOrgEventUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendOrgEventUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SendOrgEventUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *SendOrgEventUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// UpdateUserGuideReader is a Reader for the UpdateUserGuide structure.
type UpdateUserGuideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserGuideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateUserGuideNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateUserGuideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateUserGuideUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateUserGuideDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserGuideNoContent creates a UpdateUserGuideNoContent with default headers values
func NewUpdateUserGuideNoContent() *UpdateUserGuideNoContent {
	return &UpdateUserGuideNoContent{}
}

/*UpdateUserGuideNoContent handles this case with default header values.

The guide progress has been updated.
*/
type UpdateUserGuideNoContent struct {
}

func (o *UpdateUserGuideNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/guide][%d] updateUserGuideNoContent ", 204)
}

func (o *UpdateUserGuideNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGuideNotFound creates a UpdateUserGuideNotFound with default headers values
func NewUpdateUserGuideNotFound() *UpdateUserGuideNotFound {
	return &UpdateUserGuideNotFound{}
}

/*UpdateUserGuideNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateUserGuideNotFound struct {
	Payload *models.ErrorPayload
}

func (o *UpdateUserGuideNotFound) Error() string {
	return fmt.Sprintf("[PUT /user/guide][%d] updateUserGuideNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserGuideNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserGuideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGuideUnprocessableEntity creates a UpdateUserGuideUnprocessableEntity with default headers values
func NewUpdateUserGuideUnprocessableEntity() *UpdateUserGuideUnprocessableEntity {
	return &UpdateUserGuideUnprocessableEntity{}
}

/*UpdateUserGuideUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateUserGuideUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *UpdateUserGuideUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /user/guide][%d] updateUserGuideUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserGuideUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserGuideUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGuideDefault creates a UpdateUserGuideDefault with default headers values
func NewUpdateUserGuideDefault(code int) *UpdateUserGuideDefault {
	return &UpdateUserGuideDefault{
		_statusCode: code,
	}
}

/*UpdateUserGuideDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateUserGuideDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the update user guide default response
func (o *UpdateUserGuideDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserGuideDefault) Error() string {
	return fmt.Sprintf("[PUT /user/guide][%d] updateUserGuide default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUserGuideDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserGuideDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

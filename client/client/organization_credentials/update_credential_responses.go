// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// UpdateCredentialReader is a Reader for the UpdateCredential structure.
type UpdateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewUpdateCredentialLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateCredentialUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCredentialOK creates a UpdateCredentialOK with default headers values
func NewUpdateCredentialOK() *UpdateCredentialOK {
	return &UpdateCredentialOK{}
}

/*UpdateCredentialOK handles this case with default header values.

Credential updated.
*/
type UpdateCredentialOK struct {
	Payload *UpdateCredentialOKBody
}

func (o *UpdateCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/credentials/{credential_id}][%d] updateCredentialOK  %+v", 200, o.Payload)
}

func (o *UpdateCredentialOK) GetPayload() *UpdateCredentialOKBody {
	return o.Payload
}

func (o *UpdateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateCredentialOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCredentialForbidden creates a UpdateCredentialForbidden with default headers values
func NewUpdateCredentialForbidden() *UpdateCredentialForbidden {
	return &UpdateCredentialForbidden{}
}

/*UpdateCredentialForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateCredentialForbidden struct {
	Payload *models.ErrorPayload
}

func (o *UpdateCredentialForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/credentials/{credential_id}][%d] updateCredentialForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCredentialForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCredentialNotFound creates a UpdateCredentialNotFound with default headers values
func NewUpdateCredentialNotFound() *UpdateCredentialNotFound {
	return &UpdateCredentialNotFound{}
}

/*UpdateCredentialNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateCredentialNotFound struct {
	Payload *models.ErrorPayload
}

func (o *UpdateCredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/credentials/{credential_id}][%d] updateCredentialNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCredentialNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCredentialLengthRequired creates a UpdateCredentialLengthRequired with default headers values
func NewUpdateCredentialLengthRequired() *UpdateCredentialLengthRequired {
	return &UpdateCredentialLengthRequired{}
}

/*UpdateCredentialLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type UpdateCredentialLengthRequired struct {
}

func (o *UpdateCredentialLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/credentials/{credential_id}][%d] updateCredentialLengthRequired ", 411)
}

func (o *UpdateCredentialLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCredentialUnprocessableEntity creates a UpdateCredentialUnprocessableEntity with default headers values
func NewUpdateCredentialUnprocessableEntity() *UpdateCredentialUnprocessableEntity {
	return &UpdateCredentialUnprocessableEntity{}
}

/*UpdateCredentialUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateCredentialUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *UpdateCredentialUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/credentials/{credential_id}][%d] updateCredentialUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateCredentialUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateCredentialUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCredentialDefault creates a UpdateCredentialDefault with default headers values
func NewUpdateCredentialDefault(code int) *UpdateCredentialDefault {
	return &UpdateCredentialDefault{
		_statusCode: code,
	}
}

/*UpdateCredentialDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateCredentialDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the update credential default response
func (o *UpdateCredentialDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCredentialDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/credentials/{credential_id}][%d] updateCredential default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCredentialDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateCredentialOKBody update credential o k body
swagger:model UpdateCredentialOKBody
*/
type UpdateCredentialOKBody struct {

	// data
	// Required: true
	Data *models.Credential `json:"data"`
}

// Validate validates this update credential o k body
func (o *UpdateCredentialOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCredentialOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateCredentialOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateCredentialOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCredentialOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCredentialOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateCredentialOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

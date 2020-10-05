// Code generated by go-swagger; DO NOT EDIT.

package organization_forms

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

// ValidateFormsFileReader is a Reader for the ValidateFormsFile structure.
type ValidateFormsFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateFormsFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateFormsFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewValidateFormsFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateFormsFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValidateFormsFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewValidateFormsFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateFormsFileOK creates a ValidateFormsFileOK with default headers values
func NewValidateFormsFileOK() *ValidateFormsFileOK {
	return &ValidateFormsFileOK{}
}

/*ValidateFormsFileOK handles this case with default header values.

The result of validating the provided configuration
*/
type ValidateFormsFileOK struct {
	Payload *ValidateFormsFileOKBody
}

func (o *ValidateFormsFileOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/validate][%d] validateFormsFileOK  %+v", 200, o.Payload)
}

func (o *ValidateFormsFileOK) GetPayload() *ValidateFormsFileOKBody {
	return o.Payload
}

func (o *ValidateFormsFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValidateFormsFileOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateFormsFileForbidden creates a ValidateFormsFileForbidden with default headers values
func NewValidateFormsFileForbidden() *ValidateFormsFileForbidden {
	return &ValidateFormsFileForbidden{}
}

/*ValidateFormsFileForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type ValidateFormsFileForbidden struct {
	Payload *models.ErrorPayload
}

func (o *ValidateFormsFileForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/validate][%d] validateFormsFileForbidden  %+v", 403, o.Payload)
}

func (o *ValidateFormsFileForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateFormsFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateFormsFileNotFound creates a ValidateFormsFileNotFound with default headers values
func NewValidateFormsFileNotFound() *ValidateFormsFileNotFound {
	return &ValidateFormsFileNotFound{}
}

/*ValidateFormsFileNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type ValidateFormsFileNotFound struct {
	Payload *models.ErrorPayload
}

func (o *ValidateFormsFileNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/validate][%d] validateFormsFileNotFound  %+v", 404, o.Payload)
}

func (o *ValidateFormsFileNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateFormsFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateFormsFileUnprocessableEntity creates a ValidateFormsFileUnprocessableEntity with default headers values
func NewValidateFormsFileUnprocessableEntity() *ValidateFormsFileUnprocessableEntity {
	return &ValidateFormsFileUnprocessableEntity{}
}

/*ValidateFormsFileUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type ValidateFormsFileUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *ValidateFormsFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/validate][%d] validateFormsFileUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ValidateFormsFileUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateFormsFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateFormsFileDefault creates a ValidateFormsFileDefault with default headers values
func NewValidateFormsFileDefault(code int) *ValidateFormsFileDefault {
	return &ValidateFormsFileDefault{
		_statusCode: code,
	}
}

/*ValidateFormsFileDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type ValidateFormsFileDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the validate forms file default response
func (o *ValidateFormsFileDefault) Code() int {
	return o._statusCode
}

func (o *ValidateFormsFileDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/validate][%d] validateFormsFile default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateFormsFileDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateFormsFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ValidateFormsFileOKBody validate forms file o k body
swagger:model ValidateFormsFileOKBody
*/
type ValidateFormsFileOKBody struct {

	// data
	// Required: true
	Data *models.FormsValidationResult `json:"data"`
}

// Validate validates this validate forms file o k body
func (o *ValidateFormsFileOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidateFormsFileOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("validateFormsFileOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validateFormsFileOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidateFormsFileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidateFormsFileOKBody) UnmarshalBinary(b []byte) error {
	var res ValidateFormsFileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization_roles

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

// CreateOrgRoleReader is a Reader for the CreateOrgRole structure.
type CreateOrgRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrgRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrgRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateOrgRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrgRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateOrgRoleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOrgRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrgRoleOK creates a CreateOrgRoleOK with default headers values
func NewCreateOrgRoleOK() *CreateOrgRoleOK {
	return &CreateOrgRoleOK{}
}

/*CreateOrgRoleOK handles this case with default header values.

New role created in the organization.
*/
type CreateOrgRoleOK struct {
	Payload *CreateOrgRoleOKBody
}

func (o *CreateOrgRoleOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/roles][%d] createOrgRoleOK  %+v", 200, o.Payload)
}

func (o *CreateOrgRoleOK) GetPayload() *CreateOrgRoleOKBody {
	return o.Payload
}

func (o *CreateOrgRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateOrgRoleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRoleForbidden creates a CreateOrgRoleForbidden with default headers values
func NewCreateOrgRoleForbidden() *CreateOrgRoleForbidden {
	return &CreateOrgRoleForbidden{}
}

/*CreateOrgRoleForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateOrgRoleForbidden struct {
	Payload *models.ErrorPayload
}

func (o *CreateOrgRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/roles][%d] createOrgRoleForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrgRoleForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOrgRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRoleNotFound creates a CreateOrgRoleNotFound with default headers values
func NewCreateOrgRoleNotFound() *CreateOrgRoleNotFound {
	return &CreateOrgRoleNotFound{}
}

/*CreateOrgRoleNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateOrgRoleNotFound struct {
	Payload *models.ErrorPayload
}

func (o *CreateOrgRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/roles][%d] createOrgRoleNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrgRoleNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOrgRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRoleUnprocessableEntity creates a CreateOrgRoleUnprocessableEntity with default headers values
func NewCreateOrgRoleUnprocessableEntity() *CreateOrgRoleUnprocessableEntity {
	return &CreateOrgRoleUnprocessableEntity{}
}

/*CreateOrgRoleUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateOrgRoleUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *CreateOrgRoleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/roles][%d] createOrgRoleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateOrgRoleUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOrgRoleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRoleDefault creates a CreateOrgRoleDefault with default headers values
func NewCreateOrgRoleDefault(code int) *CreateOrgRoleDefault {
	return &CreateOrgRoleDefault{
		_statusCode: code,
	}
}

/*CreateOrgRoleDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateOrgRoleDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the create org role default response
func (o *CreateOrgRoleDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrgRoleDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/roles][%d] createOrgRole default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrgRoleDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOrgRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateOrgRoleOKBody create org role o k body
swagger:model CreateOrgRoleOKBody
*/
type CreateOrgRoleOKBody struct {

	// data
	// Required: true
	Data *models.NewRole `json:"data"`
}

// Validate validates this create org role o k body
func (o *CreateOrgRoleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrgRoleOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createOrgRoleOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrgRoleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrgRoleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrgRoleOKBody) UnmarshalBinary(b []byte) error {
	var res CreateOrgRoleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

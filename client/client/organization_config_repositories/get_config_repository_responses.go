// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

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

// GetConfigRepositoryReader is a Reader for the GetConfigRepository structure.
type GetConfigRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetConfigRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetConfigRepositoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetConfigRepositoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConfigRepositoryOK creates a GetConfigRepositoryOK with default headers values
func NewGetConfigRepositoryOK() *GetConfigRepositoryOK {
	return &GetConfigRepositoryOK{}
}

/*GetConfigRepositoryOK handles this case with default header values.

Organization Config Repository.
*/
type GetConfigRepositoryOK struct {
	Payload *GetConfigRepositoryOKBody
}

func (o *GetConfigRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories/{config_repository_id}][%d] getConfigRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetConfigRepositoryOK) GetPayload() *GetConfigRepositoryOKBody {
	return o.Payload
}

func (o *GetConfigRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetConfigRepositoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigRepositoryForbidden creates a GetConfigRepositoryForbidden with default headers values
func NewGetConfigRepositoryForbidden() *GetConfigRepositoryForbidden {
	return &GetConfigRepositoryForbidden{}
}

/*GetConfigRepositoryForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetConfigRepositoryForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetConfigRepositoryForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories/{config_repository_id}][%d] getConfigRepositoryForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigRepositoryForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetConfigRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigRepositoryUnprocessableEntity creates a GetConfigRepositoryUnprocessableEntity with default headers values
func NewGetConfigRepositoryUnprocessableEntity() *GetConfigRepositoryUnprocessableEntity {
	return &GetConfigRepositoryUnprocessableEntity{}
}

/*GetConfigRepositoryUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetConfigRepositoryUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetConfigRepositoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories/{config_repository_id}][%d] getConfigRepositoryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetConfigRepositoryUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetConfigRepositoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigRepositoryDefault creates a GetConfigRepositoryDefault with default headers values
func NewGetConfigRepositoryDefault(code int) *GetConfigRepositoryDefault {
	return &GetConfigRepositoryDefault{
		_statusCode: code,
	}
}

/*GetConfigRepositoryDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetConfigRepositoryDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get config repository default response
func (o *GetConfigRepositoryDefault) Code() int {
	return o._statusCode
}

func (o *GetConfigRepositoryDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories/{config_repository_id}][%d] getConfigRepository default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigRepositoryDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetConfigRepositoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetConfigRepositoryOKBody get config repository o k body
swagger:model GetConfigRepositoryOKBody
*/
type GetConfigRepositoryOKBody struct {

	// data
	// Required: true
	Data *models.ConfigRepository `json:"data"`
}

// Validate validates this get config repository o k body
func (o *GetConfigRepositoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigRepositoryOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getConfigRepositoryOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getConfigRepositoryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigRepositoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigRepositoryOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigRepositoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

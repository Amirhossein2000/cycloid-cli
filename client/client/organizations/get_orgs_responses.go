// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// GetOrgsReader is a Reader for the GetOrgs structure.
type GetOrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetOrgsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgsOK creates a GetOrgsOK with default headers values
func NewGetOrgsOK() *GetOrgsOK {
	return &GetOrgsOK{}
}

/*GetOrgsOK handles this case with default header values.

List of the organizations which authenticated user has access.
*/
type GetOrgsOK struct {
	Payload *GetOrgsOKBody
}

func (o *GetOrgsOK) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrgsOK  %+v", 200, o.Payload)
}

func (o *GetOrgsOK) GetPayload() *GetOrgsOKBody {
	return o.Payload
}

func (o *GetOrgsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrgsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgsUnprocessableEntity creates a GetOrgsUnprocessableEntity with default headers values
func NewGetOrgsUnprocessableEntity() *GetOrgsUnprocessableEntity {
	return &GetOrgsUnprocessableEntity{}
}

/*GetOrgsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetOrgsUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrgsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetOrgsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgsDefault creates a GetOrgsDefault with default headers values
func NewGetOrgsDefault(code int) *GetOrgsDefault {
	return &GetOrgsDefault{
		_statusCode: code,
	}
}

/*GetOrgsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get orgs default response
func (o *GetOrgsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrgs default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrgsOKBody get orgs o k body
swagger:model GetOrgsOKBody
*/
type GetOrgsOKBody struct {

	// data
	// Required: true
	Data []*models.Organization `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get orgs o k body
func (o *GetOrgsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrgsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrgsOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getOrgsOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

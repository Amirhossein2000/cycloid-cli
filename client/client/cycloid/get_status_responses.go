// Code generated by go-swagger; DO NOT EDIT.

package cycloid

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

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*GetStatusOK handles this case with default header values.

General application status and services statuses.
*/
type GetStatusOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *GetStatusOKBody
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}

func (o *GetStatusOK) GetPayload() *GetStatusOKBody {
	return o.Payload
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusUnprocessableEntity creates a GetStatusUnprocessableEntity with default headers values
func NewGetStatusUnprocessableEntity() *GetStatusUnprocessableEntity {
	return &GetStatusUnprocessableEntity{}
}

/*GetStatusUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetStatusUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetStatusUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetStatusUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusDefault creates a GetStatusDefault with default headers values
func NewGetStatusDefault(code int) *GetStatusDefault {
	return &GetStatusDefault{
		_statusCode: code,
	}
}

/*GetStatusDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetStatusDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get status default response
func (o *GetStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetStatusDefault) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetStatusDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetStatusOKBody get status o k body
swagger:model GetStatusOKBody
*/
type GetStatusOKBody struct {

	// data
	// Required: true
	Data *models.GeneralStatus `json:"data"`
}

// Validate validates this get status o k body
func (o *GetStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getStatusOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStatusOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

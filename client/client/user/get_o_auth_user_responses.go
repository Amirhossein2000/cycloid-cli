// Code generated by go-swagger; DO NOT EDIT.

package user

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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetOAuthUserReader is a Reader for the GetOAuthUser structure.
type GetOAuthUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOAuthUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOAuthUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOAuthUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOAuthUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOAuthUserOK creates a GetOAuthUserOK with default headers values
func NewGetOAuthUserOK() *GetOAuthUserOK {
	return &GetOAuthUserOK{}
}

/*GetOAuthUserOK handles this case with default header values.

Used to know if a user from the platform exists on that 'social_type'.
*/
type GetOAuthUserOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *GetOAuthUserOKBody
}

func (o *GetOAuthUserOK) Error() string {
	return fmt.Sprintf("[GET /user/{social_type}/oauth][%d] getOAuthUserOK  %+v", 200, o.Payload)
}

func (o *GetOAuthUserOK) GetPayload() *GetOAuthUserOKBody {
	return o.Payload
}

func (o *GetOAuthUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetOAuthUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOAuthUserUnauthorized creates a GetOAuthUserUnauthorized with default headers values
func NewGetOAuthUserUnauthorized() *GetOAuthUserUnauthorized {
	return &GetOAuthUserUnauthorized{}
}

/*GetOAuthUserUnauthorized handles this case with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type GetOAuthUserUnauthorized struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetOAuthUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/{social_type}/oauth][%d] getOAuthUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOAuthUserUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOAuthUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOAuthUserDefault creates a GetOAuthUserDefault with default headers values
func NewGetOAuthUserDefault(code int) *GetOAuthUserDefault {
	return &GetOAuthUserDefault{
		_statusCode: code,
	}
}

/*GetOAuthUserDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOAuthUserDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get o auth user default response
func (o *GetOAuthUserDefault) Code() int {
	return o._statusCode
}

func (o *GetOAuthUserDefault) Error() string {
	return fmt.Sprintf("[GET /user/{social_type}/oauth][%d] getOAuthUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetOAuthUserDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOAuthUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOAuthUserOKBody get o auth user o k body
swagger:model GetOAuthUserOKBody
*/
type GetOAuthUserOKBody struct {

	// data
	// Required: true
	Data *models.UserOAuth `json:"data"`
}

// Validate validates this get o auth user o k body
func (o *GetOAuthUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOAuthUserOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOAuthUserOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOAuthUserOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOAuthUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOAuthUserOKBody) UnmarshalBinary(b []byte) error {
	var res GetOAuthUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
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

// GetOrgEventsReader is a Reader for the GetOrgEvents structure.
type GetOrgEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrgEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrgEventsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgEventsOK creates a GetOrgEventsOK with default headers values
func NewGetOrgEventsOK() *GetOrgEventsOK {
	return &GetOrgEventsOK{}
}

/*GetOrgEventsOK handles this case with default header values.

The list of events which fulfills the query parameters filter
*/
type GetOrgEventsOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *GetOrgEventsOKBody
}

func (o *GetOrgEventsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events][%d] getOrgEventsOK  %+v", 200, o.Payload)
}

func (o *GetOrgEventsOK) GetPayload() *GetOrgEventsOKBody {
	return o.Payload
}

func (o *GetOrgEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetOrgEventsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgEventsForbidden creates a GetOrgEventsForbidden with default headers values
func NewGetOrgEventsForbidden() *GetOrgEventsForbidden {
	return &GetOrgEventsForbidden{}
}

/*GetOrgEventsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetOrgEventsForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetOrgEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events][%d] getOrgEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgEventsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOrgEventsNotFound creates a GetOrgEventsNotFound with default headers values
func NewGetOrgEventsNotFound() *GetOrgEventsNotFound {
	return &GetOrgEventsNotFound{}
}

/*GetOrgEventsNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetOrgEventsNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetOrgEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events][%d] getOrgEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgEventsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOrgEventsUnprocessableEntity creates a GetOrgEventsUnprocessableEntity with default headers values
func NewGetOrgEventsUnprocessableEntity() *GetOrgEventsUnprocessableEntity {
	return &GetOrgEventsUnprocessableEntity{}
}

/*GetOrgEventsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetOrgEventsUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetOrgEventsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events][%d] getOrgEventsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetOrgEventsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOrgEventsDefault creates a GetOrgEventsDefault with default headers values
func NewGetOrgEventsDefault(code int) *GetOrgEventsDefault {
	return &GetOrgEventsDefault{
		_statusCode: code,
	}
}

/*GetOrgEventsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgEventsDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get org events default response
func (o *GetOrgEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgEventsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events][%d] getOrgEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgEventsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetOrgEventsOKBody The list of log lines
swagger:model GetOrgEventsOKBody
*/
type GetOrgEventsOKBody struct {

	// data
	// Required: true
	Data []*models.Event `json:"data"`
}

// Validate validates this get org events o k body
func (o *GetOrgEventsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgEventsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgEventsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrgEventsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgEventsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgEventsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgEventsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

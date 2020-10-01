// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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

// GetServiceCatalogConfigReader is a Reader for the GetServiceCatalogConfig structure.
type GetServiceCatalogConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceCatalogConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogConfigOK creates a GetServiceCatalogConfigOK with default headers values
func NewGetServiceCatalogConfigOK() *GetServiceCatalogConfigOK {
	return &GetServiceCatalogConfigOK{}
}

/*GetServiceCatalogConfigOK handles this case with default header values.

The config of the service catalog.
*/
type GetServiceCatalogConfigOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *GetServiceCatalogConfigOKBody
}

func (o *GetServiceCatalogConfigOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/config][%d] getServiceCatalogConfigOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogConfigOK) GetPayload() *GetServiceCatalogConfigOKBody {
	return o.Payload
}

func (o *GetServiceCatalogConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetServiceCatalogConfigOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogConfigForbidden creates a GetServiceCatalogConfigForbidden with default headers values
func NewGetServiceCatalogConfigForbidden() *GetServiceCatalogConfigForbidden {
	return &GetServiceCatalogConfigForbidden{}
}

/*GetServiceCatalogConfigForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogConfigForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/config][%d] getServiceCatalogConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogConfigForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogConfigNotFound creates a GetServiceCatalogConfigNotFound with default headers values
func NewGetServiceCatalogConfigNotFound() *GetServiceCatalogConfigNotFound {
	return &GetServiceCatalogConfigNotFound{}
}

/*GetServiceCatalogConfigNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetServiceCatalogConfigNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/config][%d] getServiceCatalogConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceCatalogConfigNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogConfigDefault creates a GetServiceCatalogConfigDefault with default headers values
func NewGetServiceCatalogConfigDefault(code int) *GetServiceCatalogConfigDefault {
	return &GetServiceCatalogConfigDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogConfigDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogConfigDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalog config default response
func (o *GetServiceCatalogConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogConfigDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/config][%d] getServiceCatalogConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogConfigDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServiceCatalogConfigOKBody get service catalog config o k body
swagger:model GetServiceCatalogConfigOKBody
*/
type GetServiceCatalogConfigOKBody struct {

	// data
	// Required: true
	Data interface{} `json:"data"`
}

// Validate validates this get service catalog config o k body
func (o *GetServiceCatalogConfigOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogConfigOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogConfigOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogConfigOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogConfigOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogConfigOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

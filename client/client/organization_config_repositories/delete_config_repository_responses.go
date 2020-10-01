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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// DeleteConfigRepositoryReader is a Reader for the DeleteConfigRepository structure.
type DeleteConfigRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfigRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConfigRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteConfigRepositoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConfigRepositoryNoContent creates a DeleteConfigRepositoryNoContent with default headers values
func NewDeleteConfigRepositoryNoContent() *DeleteConfigRepositoryNoContent {
	return &DeleteConfigRepositoryNoContent{}
}

/*DeleteConfigRepositoryNoContent handles this case with default header values.

Organization Config repository has been deleted
*/
type DeleteConfigRepositoryNoContent struct {
}

func (o *DeleteConfigRepositoryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/config_repositories/{config_repository_id}][%d] deleteConfigRepositoryNoContent ", 204)
}

func (o *DeleteConfigRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConfigRepositoryDefault creates a DeleteConfigRepositoryDefault with default headers values
func NewDeleteConfigRepositoryDefault(code int) *DeleteConfigRepositoryDefault {
	return &DeleteConfigRepositoryDefault{
		_statusCode: code,
	}
}

/*DeleteConfigRepositoryDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteConfigRepositoryDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete config repository default response
func (o *DeleteConfigRepositoryDefault) Code() int {
	return o._statusCode
}

func (o *DeleteConfigRepositoryDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/config_repositories/{config_repository_id}][%d] deleteConfigRepository default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConfigRepositoryDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteConfigRepositoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// EmailVerificationReader is a Reader for the EmailVerification structure.
type EmailVerificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailVerificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEmailVerificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmailVerificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailVerificationNoContent creates a EmailVerificationNoContent with default headers values
func NewEmailVerificationNoContent() *EmailVerificationNoContent {
	return &EmailVerificationNoContent{}
}

/*EmailVerificationNoContent handles this case with default header values.

Email address has been verified.
*/
type EmailVerificationNoContent struct {
}

func (o *EmailVerificationNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationNoContent ", 204)
}

func (o *EmailVerificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmailVerificationDefault creates a EmailVerificationDefault with default headers values
func NewEmailVerificationDefault(code int) *EmailVerificationDefault {
	return &EmailVerificationDefault{
		_statusCode: code,
	}
}

/*EmailVerificationDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type EmailVerificationDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the email verification default response
func (o *EmailVerificationDefault) Code() int {
	return o._statusCode
}

func (o *EmailVerificationDefault) Error() string {
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerification default  %+v", o._statusCode, o.Payload)
}

func (o *EmailVerificationDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EmailVerificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

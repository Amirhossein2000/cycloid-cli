// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// PauseJobReader is a Reader for the PauseJob structure.
type PauseJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPauseJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPauseJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPauseJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPauseJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPauseJobNoContent creates a PauseJobNoContent with default headers values
func NewPauseJobNoContent() *PauseJobNoContent {
	return &PauseJobNoContent{}
}

/*PauseJobNoContent handles this case with default header values.

Job has been paused.
*/
type PauseJobNoContent struct {
}

func (o *PauseJobNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause][%d] pauseJobNoContent ", 204)
}

func (o *PauseJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPauseJobForbidden creates a PauseJobForbidden with default headers values
func NewPauseJobForbidden() *PauseJobForbidden {
	return &PauseJobForbidden{}
}

/*PauseJobForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type PauseJobForbidden struct {
	Payload *models.ErrorPayload
}

func (o *PauseJobForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause][%d] pauseJobForbidden  %+v", 403, o.Payload)
}

func (o *PauseJobForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PauseJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseJobNotFound creates a PauseJobNotFound with default headers values
func NewPauseJobNotFound() *PauseJobNotFound {
	return &PauseJobNotFound{}
}

/*PauseJobNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type PauseJobNotFound struct {
	Payload *models.ErrorPayload
}

func (o *PauseJobNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause][%d] pauseJobNotFound  %+v", 404, o.Payload)
}

func (o *PauseJobNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PauseJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseJobDefault creates a PauseJobDefault with default headers values
func NewPauseJobDefault(code int) *PauseJobDefault {
	return &PauseJobDefault{
		_statusCode: code,
	}
}

/*PauseJobDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type PauseJobDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the pause job default response
func (o *PauseJobDefault) Code() int {
	return o._statusCode
}

func (o *PauseJobDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause][%d] pauseJob default  %+v", o._statusCode, o.Payload)
}

func (o *PauseJobDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PauseJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

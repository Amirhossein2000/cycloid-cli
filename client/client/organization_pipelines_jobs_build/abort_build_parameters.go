// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs_build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAbortBuildParams creates a new AbortBuildParams object
// with the default values initialized.
func NewAbortBuildParams() *AbortBuildParams {
	var ()
	return &AbortBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAbortBuildParamsWithTimeout creates a new AbortBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAbortBuildParamsWithTimeout(timeout time.Duration) *AbortBuildParams {
	var ()
	return &AbortBuildParams{

		timeout: timeout,
	}
}

// NewAbortBuildParamsWithContext creates a new AbortBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewAbortBuildParamsWithContext(ctx context.Context) *AbortBuildParams {
	var ()
	return &AbortBuildParams{

		Context: ctx,
	}
}

// NewAbortBuildParamsWithHTTPClient creates a new AbortBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAbortBuildParamsWithHTTPClient(client *http.Client) *AbortBuildParams {
	var ()
	return &AbortBuildParams{
		HTTPClient: client,
	}
}

/*AbortBuildParams contains all the parameters to send to the API endpoint
for the abort build operation typically these are written to a http.Request
*/
type AbortBuildParams struct {

	/*BuildID
	  A build id

	*/
	BuildID string
	/*InpathPipelineName
	  A pipeline name

	*/
	InpathPipelineName string
	/*JobName
	  A job name

	*/
	JobName string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*ProjectCanonical
	  A canonical of a project.

	*/
	ProjectCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the abort build params
func (o *AbortBuildParams) WithTimeout(timeout time.Duration) *AbortBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the abort build params
func (o *AbortBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the abort build params
func (o *AbortBuildParams) WithContext(ctx context.Context) *AbortBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the abort build params
func (o *AbortBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the abort build params
func (o *AbortBuildParams) WithHTTPClient(client *http.Client) *AbortBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the abort build params
func (o *AbortBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the abort build params
func (o *AbortBuildParams) WithBuildID(buildID string) *AbortBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the abort build params
func (o *AbortBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithInpathPipelineName adds the inpathPipelineName to the abort build params
func (o *AbortBuildParams) WithInpathPipelineName(inpathPipelineName string) *AbortBuildParams {
	o.SetInpathPipelineName(inpathPipelineName)
	return o
}

// SetInpathPipelineName adds the inpathPipelineName to the abort build params
func (o *AbortBuildParams) SetInpathPipelineName(inpathPipelineName string) {
	o.InpathPipelineName = inpathPipelineName
}

// WithJobName adds the jobName to the abort build params
func (o *AbortBuildParams) WithJobName(jobName string) *AbortBuildParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the abort build params
func (o *AbortBuildParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithOrganizationCanonical adds the organizationCanonical to the abort build params
func (o *AbortBuildParams) WithOrganizationCanonical(organizationCanonical string) *AbortBuildParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the abort build params
func (o *AbortBuildParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the abort build params
func (o *AbortBuildParams) WithProjectCanonical(projectCanonical string) *AbortBuildParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the abort build params
func (o *AbortBuildParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *AbortBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}

	// path param inpath_pipeline_name
	if err := r.SetPathParam("inpath_pipeline_name", o.InpathPipelineName); err != nil {
		return err
	}

	// path param job_name
	if err := r.SetPathParam("job_name", o.JobName); err != nil {
		return err
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param project_canonical
	if err := r.SetPathParam("project_canonical", o.ProjectCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
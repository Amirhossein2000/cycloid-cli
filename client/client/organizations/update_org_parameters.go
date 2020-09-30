// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// NewUpdateOrgParams creates a new UpdateOrgParams object
// with the default values initialized.
func NewUpdateOrgParams() *UpdateOrgParams {
	var ()
	return &UpdateOrgParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrgParamsWithTimeout creates a new UpdateOrgParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrgParamsWithTimeout(timeout time.Duration) *UpdateOrgParams {
	var ()
	return &UpdateOrgParams{

		timeout: timeout,
	}
}

// NewUpdateOrgParamsWithContext creates a new UpdateOrgParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrgParamsWithContext(ctx context.Context) *UpdateOrgParams {
	var ()
	return &UpdateOrgParams{

		Context: ctx,
	}
}

// NewUpdateOrgParamsWithHTTPClient creates a new UpdateOrgParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrgParamsWithHTTPClient(client *http.Client) *UpdateOrgParams {
	var ()
	return &UpdateOrgParams{
		HTTPClient: client,
	}
}

/*UpdateOrgParams contains all the parameters to send to the API endpoint
for the update org operation typically these are written to a http.Request
*/
type UpdateOrgParams struct {

	/*Body
	  The information of the organization to update.

	*/
	Body *models.UpdateOrganization
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update org params
func (o *UpdateOrgParams) WithTimeout(timeout time.Duration) *UpdateOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update org params
func (o *UpdateOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update org params
func (o *UpdateOrgParams) WithContext(ctx context.Context) *UpdateOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update org params
func (o *UpdateOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update org params
func (o *UpdateOrgParams) WithHTTPClient(client *http.Client) *UpdateOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update org params
func (o *UpdateOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update org params
func (o *UpdateOrgParams) WithBody(body *models.UpdateOrganization) *UpdateOrgParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update org params
func (o *UpdateOrgParams) SetBody(body *models.UpdateOrganization) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update org params
func (o *UpdateOrgParams) WithOrganizationCanonical(organizationCanonical string) *UpdateOrgParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update org params
func (o *UpdateOrgParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization_roles

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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// NewCreateOrgRoleParams creates a new CreateOrgRoleParams object
// with the default values initialized.
func NewCreateOrgRoleParams() *CreateOrgRoleParams {
	var ()
	return &CreateOrgRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgRoleParamsWithTimeout creates a new CreateOrgRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrgRoleParamsWithTimeout(timeout time.Duration) *CreateOrgRoleParams {
	var ()
	return &CreateOrgRoleParams{

		timeout: timeout,
	}
}

// NewCreateOrgRoleParamsWithContext creates a new CreateOrgRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrgRoleParamsWithContext(ctx context.Context) *CreateOrgRoleParams {
	var ()
	return &CreateOrgRoleParams{

		Context: ctx,
	}
}

// NewCreateOrgRoleParamsWithHTTPClient creates a new CreateOrgRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrgRoleParamsWithHTTPClient(client *http.Client) *CreateOrgRoleParams {
	var ()
	return &CreateOrgRoleParams{
		HTTPClient: client,
	}
}

/*CreateOrgRoleParams contains all the parameters to send to the API endpoint
for the create org role operation typically these are written to a http.Request
*/
type CreateOrgRoleParams struct {

	/*Body
	  The information of the organization's role to create.

	*/
	Body *models.NewRole
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create org role params
func (o *CreateOrgRoleParams) WithTimeout(timeout time.Duration) *CreateOrgRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org role params
func (o *CreateOrgRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org role params
func (o *CreateOrgRoleParams) WithContext(ctx context.Context) *CreateOrgRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org role params
func (o *CreateOrgRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org role params
func (o *CreateOrgRoleParams) WithHTTPClient(client *http.Client) *CreateOrgRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org role params
func (o *CreateOrgRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create org role params
func (o *CreateOrgRoleParams) WithBody(body *models.NewRole) *CreateOrgRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create org role params
func (o *CreateOrgRoleParams) SetBody(body *models.NewRole) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create org role params
func (o *CreateOrgRoleParams) WithOrganizationCanonical(organizationCanonical string) *CreateOrgRoleParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create org role params
func (o *CreateOrgRoleParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
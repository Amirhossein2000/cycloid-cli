// Code generated by go-swagger; DO NOT EDIT.

package organization_members

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

// NewInviteUserToOrgMemberParams creates a new InviteUserToOrgMemberParams object
// with the default values initialized.
func NewInviteUserToOrgMemberParams() *InviteUserToOrgMemberParams {
	var ()
	return &InviteUserToOrgMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInviteUserToOrgMemberParamsWithTimeout creates a new InviteUserToOrgMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInviteUserToOrgMemberParamsWithTimeout(timeout time.Duration) *InviteUserToOrgMemberParams {
	var ()
	return &InviteUserToOrgMemberParams{

		timeout: timeout,
	}
}

// NewInviteUserToOrgMemberParamsWithContext creates a new InviteUserToOrgMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewInviteUserToOrgMemberParamsWithContext(ctx context.Context) *InviteUserToOrgMemberParams {
	var ()
	return &InviteUserToOrgMemberParams{

		Context: ctx,
	}
}

// NewInviteUserToOrgMemberParamsWithHTTPClient creates a new InviteUserToOrgMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInviteUserToOrgMemberParamsWithHTTPClient(client *http.Client) *InviteUserToOrgMemberParams {
	var ()
	return &InviteUserToOrgMemberParams{
		HTTPClient: client,
	}
}

/*InviteUserToOrgMemberParams contains all the parameters to send to the API endpoint
for the invite user to org member operation typically these are written to a http.Request
*/
type InviteUserToOrgMemberParams struct {

	/*Body
	  The user's member invitation.

	*/
	Body *models.NewMemberInvitation
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invite user to org member params
func (o *InviteUserToOrgMemberParams) WithTimeout(timeout time.Duration) *InviteUserToOrgMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite user to org member params
func (o *InviteUserToOrgMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite user to org member params
func (o *InviteUserToOrgMemberParams) WithContext(ctx context.Context) *InviteUserToOrgMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite user to org member params
func (o *InviteUserToOrgMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite user to org member params
func (o *InviteUserToOrgMemberParams) WithHTTPClient(client *http.Client) *InviteUserToOrgMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite user to org member params
func (o *InviteUserToOrgMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invite user to org member params
func (o *InviteUserToOrgMemberParams) WithBody(body *models.NewMemberInvitation) *InviteUserToOrgMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invite user to org member params
func (o *InviteUserToOrgMemberParams) SetBody(body *models.NewMemberInvitation) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the invite user to org member params
func (o *InviteUserToOrgMemberParams) WithOrganizationCanonical(organizationCanonical string) *InviteUserToOrgMemberParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the invite user to org member params
func (o *InviteUserToOrgMemberParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *InviteUserToOrgMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
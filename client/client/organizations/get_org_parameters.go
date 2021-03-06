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
)

// NewGetOrgParams creates a new GetOrgParams object
// with the default values initialized.
func NewGetOrgParams() *GetOrgParams {
	var ()
	return &GetOrgParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgParamsWithTimeout creates a new GetOrgParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrgParamsWithTimeout(timeout time.Duration) *GetOrgParams {
	var ()
	return &GetOrgParams{

		timeout: timeout,
	}
}

// NewGetOrgParamsWithContext creates a new GetOrgParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrgParamsWithContext(ctx context.Context) *GetOrgParams {
	var ()
	return &GetOrgParams{

		Context: ctx,
	}
}

// NewGetOrgParamsWithHTTPClient creates a new GetOrgParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrgParamsWithHTTPClient(client *http.Client) *GetOrgParams {
	var ()
	return &GetOrgParams{
		HTTPClient: client,
	}
}

/*GetOrgParams contains all the parameters to send to the API endpoint
for the get org operation typically these are written to a http.Request
*/
type GetOrgParams struct {

	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get org params
func (o *GetOrgParams) WithTimeout(timeout time.Duration) *GetOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org params
func (o *GetOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org params
func (o *GetOrgParams) WithContext(ctx context.Context) *GetOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org params
func (o *GetOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org params
func (o *GetOrgParams) WithHTTPClient(client *http.Client) *GetOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org params
func (o *GetOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get org params
func (o *GetOrgParams) WithOrganizationCanonical(organizationCanonical string) *GetOrgParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get org params
func (o *GetOrgParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

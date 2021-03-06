// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCredentialParams creates a new DeleteCredentialParams object
// with the default values initialized.
func NewDeleteCredentialParams() *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCredentialParamsWithTimeout creates a new DeleteCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCredentialParamsWithTimeout(timeout time.Duration) *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{

		timeout: timeout,
	}
}

// NewDeleteCredentialParamsWithContext creates a new DeleteCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCredentialParamsWithContext(ctx context.Context) *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{

		Context: ctx,
	}
}

// NewDeleteCredentialParamsWithHTTPClient creates a new DeleteCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCredentialParamsWithHTTPClient(client *http.Client) *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{
		HTTPClient: client,
	}
}

/*DeleteCredentialParams contains all the parameters to send to the API endpoint
for the delete credential operation typically these are written to a http.Request
*/
type DeleteCredentialParams struct {

	/*CredentialID
	  A Credential id

	*/
	CredentialID uint32
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete credential params
func (o *DeleteCredentialParams) WithTimeout(timeout time.Duration) *DeleteCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete credential params
func (o *DeleteCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete credential params
func (o *DeleteCredentialParams) WithContext(ctx context.Context) *DeleteCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete credential params
func (o *DeleteCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete credential params
func (o *DeleteCredentialParams) WithHTTPClient(client *http.Client) *DeleteCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete credential params
func (o *DeleteCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the delete credential params
func (o *DeleteCredentialParams) WithCredentialID(credentialID uint32) *DeleteCredentialParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the delete credential params
func (o *DeleteCredentialParams) SetCredentialID(credentialID uint32) {
	o.CredentialID = credentialID
}

// WithOrganizationCanonical adds the organizationCanonical to the delete credential params
func (o *DeleteCredentialParams) WithOrganizationCanonical(organizationCanonical string) *DeleteCredentialParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the delete credential params
func (o *DeleteCredentialParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credential_id
	if err := r.SetPathParam("credential_id", swag.FormatUint32(o.CredentialID)); err != nil {
		return err
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

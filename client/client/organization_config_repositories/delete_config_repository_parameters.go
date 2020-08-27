// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

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

// NewDeleteConfigRepositoryParams creates a new DeleteConfigRepositoryParams object
// with the default values initialized.
func NewDeleteConfigRepositoryParams() *DeleteConfigRepositoryParams {
	var ()
	return &DeleteConfigRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfigRepositoryParamsWithTimeout creates a new DeleteConfigRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConfigRepositoryParamsWithTimeout(timeout time.Duration) *DeleteConfigRepositoryParams {
	var ()
	return &DeleteConfigRepositoryParams{

		timeout: timeout,
	}
}

// NewDeleteConfigRepositoryParamsWithContext creates a new DeleteConfigRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConfigRepositoryParamsWithContext(ctx context.Context) *DeleteConfigRepositoryParams {
	var ()
	return &DeleteConfigRepositoryParams{

		Context: ctx,
	}
}

// NewDeleteConfigRepositoryParamsWithHTTPClient creates a new DeleteConfigRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConfigRepositoryParamsWithHTTPClient(client *http.Client) *DeleteConfigRepositoryParams {
	var ()
	return &DeleteConfigRepositoryParams{
		HTTPClient: client,
	}
}

/*DeleteConfigRepositoryParams contains all the parameters to send to the API endpoint
for the delete config repository operation typically these are written to a http.Request
*/
type DeleteConfigRepositoryParams struct {

	/*ConfigRepositoryID
	  Organization Config Repositories ID

	*/
	ConfigRepositoryID uint32
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete config repository params
func (o *DeleteConfigRepositoryParams) WithTimeout(timeout time.Duration) *DeleteConfigRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete config repository params
func (o *DeleteConfigRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete config repository params
func (o *DeleteConfigRepositoryParams) WithContext(ctx context.Context) *DeleteConfigRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete config repository params
func (o *DeleteConfigRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete config repository params
func (o *DeleteConfigRepositoryParams) WithHTTPClient(client *http.Client) *DeleteConfigRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete config repository params
func (o *DeleteConfigRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigRepositoryID adds the configRepositoryID to the delete config repository params
func (o *DeleteConfigRepositoryParams) WithConfigRepositoryID(configRepositoryID uint32) *DeleteConfigRepositoryParams {
	o.SetConfigRepositoryID(configRepositoryID)
	return o
}

// SetConfigRepositoryID adds the configRepositoryId to the delete config repository params
func (o *DeleteConfigRepositoryParams) SetConfigRepositoryID(configRepositoryID uint32) {
	o.ConfigRepositoryID = configRepositoryID
}

// WithOrganizationCanonical adds the organizationCanonical to the delete config repository params
func (o *DeleteConfigRepositoryParams) WithOrganizationCanonical(organizationCanonical string) *DeleteConfigRepositoryParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the delete config repository params
func (o *DeleteConfigRepositoryParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfigRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param config_repository_id
	if err := r.SetPathParam("config_repository_id", swag.FormatUint32(o.ConfigRepositoryID)); err != nil {
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
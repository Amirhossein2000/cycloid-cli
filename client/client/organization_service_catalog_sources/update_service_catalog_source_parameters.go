// Code generated by go-swagger; DO NOT EDIT.

package organization_service_catalog_sources

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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// NewUpdateServiceCatalogSourceParams creates a new UpdateServiceCatalogSourceParams object
// with the default values initialized.
func NewUpdateServiceCatalogSourceParams() *UpdateServiceCatalogSourceParams {
	var ()
	return &UpdateServiceCatalogSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceCatalogSourceParamsWithTimeout creates a new UpdateServiceCatalogSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceCatalogSourceParamsWithTimeout(timeout time.Duration) *UpdateServiceCatalogSourceParams {
	var ()
	return &UpdateServiceCatalogSourceParams{

		timeout: timeout,
	}
}

// NewUpdateServiceCatalogSourceParamsWithContext creates a new UpdateServiceCatalogSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceCatalogSourceParamsWithContext(ctx context.Context) *UpdateServiceCatalogSourceParams {
	var ()
	return &UpdateServiceCatalogSourceParams{

		Context: ctx,
	}
}

// NewUpdateServiceCatalogSourceParamsWithHTTPClient creates a new UpdateServiceCatalogSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceCatalogSourceParamsWithHTTPClient(client *http.Client) *UpdateServiceCatalogSourceParams {
	var ()
	return &UpdateServiceCatalogSourceParams{
		HTTPClient: client,
	}
}

/*UpdateServiceCatalogSourceParams contains all the parameters to send to the API endpoint
for the update service catalog source operation typically these are written to a http.Request
*/
type UpdateServiceCatalogSourceParams struct {

	/*Body
	  The information of the organization to create.

	*/
	Body *models.UpdateServiceCatalogSource
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*ServiceCatalogSourceID
	  Organization Service Catalog Sources ID

	*/
	ServiceCatalogSourceID uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) WithTimeout(timeout time.Duration) *UpdateServiceCatalogSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) WithContext(ctx context.Context) *UpdateServiceCatalogSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) WithHTTPClient(client *http.Client) *UpdateServiceCatalogSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) WithBody(body *models.UpdateServiceCatalogSource) *UpdateServiceCatalogSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) SetBody(body *models.UpdateServiceCatalogSource) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) WithOrganizationCanonical(organizationCanonical string) *UpdateServiceCatalogSourceParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogSourceID adds the serviceCatalogSourceID to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) WithServiceCatalogSourceID(serviceCatalogSourceID uint32) *UpdateServiceCatalogSourceParams {
	o.SetServiceCatalogSourceID(serviceCatalogSourceID)
	return o
}

// SetServiceCatalogSourceID adds the serviceCatalogSourceId to the update service catalog source params
func (o *UpdateServiceCatalogSourceParams) SetServiceCatalogSourceID(serviceCatalogSourceID uint32) {
	o.ServiceCatalogSourceID = serviceCatalogSourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceCatalogSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param service_catalog_source_id
	if err := r.SetPathParam("service_catalog_source_id", swag.FormatUint32(o.ServiceCatalogSourceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

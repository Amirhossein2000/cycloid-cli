// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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

// NewGetServiceCatalogTerraformDiagramParams creates a new GetServiceCatalogTerraformDiagramParams object
// with the default values initialized.
func NewGetServiceCatalogTerraformDiagramParams() *GetServiceCatalogTerraformDiagramParams {
	var ()
	return &GetServiceCatalogTerraformDiagramParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceCatalogTerraformDiagramParamsWithTimeout creates a new GetServiceCatalogTerraformDiagramParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceCatalogTerraformDiagramParamsWithTimeout(timeout time.Duration) *GetServiceCatalogTerraformDiagramParams {
	var ()
	return &GetServiceCatalogTerraformDiagramParams{

		timeout: timeout,
	}
}

// NewGetServiceCatalogTerraformDiagramParamsWithContext creates a new GetServiceCatalogTerraformDiagramParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceCatalogTerraformDiagramParamsWithContext(ctx context.Context) *GetServiceCatalogTerraformDiagramParams {
	var ()
	return &GetServiceCatalogTerraformDiagramParams{

		Context: ctx,
	}
}

// NewGetServiceCatalogTerraformDiagramParamsWithHTTPClient creates a new GetServiceCatalogTerraformDiagramParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceCatalogTerraformDiagramParamsWithHTTPClient(client *http.Client) *GetServiceCatalogTerraformDiagramParams {
	var ()
	return &GetServiceCatalogTerraformDiagramParams{
		HTTPClient: client,
	}
}

/*GetServiceCatalogTerraformDiagramParams contains all the parameters to send to the API endpoint
for the get service catalog terraform diagram operation typically these are written to a http.Request
*/
type GetServiceCatalogTerraformDiagramParams struct {

	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*ServiceCatalogRef
	  A Service Catalog name

	*/
	ServiceCatalogRef string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) WithTimeout(timeout time.Duration) *GetServiceCatalogTerraformDiagramParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) WithContext(ctx context.Context) *GetServiceCatalogTerraformDiagramParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) WithHTTPClient(client *http.Client) *GetServiceCatalogTerraformDiagramParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) WithOrganizationCanonical(organizationCanonical string) *GetServiceCatalogTerraformDiagramParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogRef adds the serviceCatalogRef to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) WithServiceCatalogRef(serviceCatalogRef string) *GetServiceCatalogTerraformDiagramParams {
	o.SetServiceCatalogRef(serviceCatalogRef)
	return o
}

// SetServiceCatalogRef adds the serviceCatalogRef to the get service catalog terraform diagram params
func (o *GetServiceCatalogTerraformDiagramParams) SetServiceCatalogRef(serviceCatalogRef string) {
	o.ServiceCatalogRef = serviceCatalogRef
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceCatalogTerraformDiagramParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param service_catalog_ref
	if err := r.SetPathParam("service_catalog_ref", o.ServiceCatalogRef); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
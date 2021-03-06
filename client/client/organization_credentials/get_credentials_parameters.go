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

// NewGetCredentialsParams creates a new GetCredentialsParams object
// with the default values initialized.
func NewGetCredentialsParams() *GetCredentialsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetCredentialsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialsParamsWithTimeout creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCredentialsParamsWithTimeout(timeout time.Duration) *GetCredentialsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetCredentialsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetCredentialsParamsWithContext creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCredentialsParamsWithContext(ctx context.Context) *GetCredentialsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetCredentialsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetCredentialsParamsWithHTTPClient creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCredentialsParamsWithHTTPClient(client *http.Client) *GetCredentialsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetCredentialsParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetCredentialsParams contains all the parameters to send to the API endpoint
for the get credentials operation typically these are written to a http.Request
*/
type GetCredentialsParams struct {

	/*CredentialType
	  A Credential type

	*/
	CredentialType *string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*PageIndex
	  The page number to request. The first page is 1.

	*/
	PageIndex *uint32
	/*PageSize
	  The number of items at most which the response can have.

	*/
	PageSize *uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) WithTimeout(timeout time.Duration) *GetCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials params
func (o *GetCredentialsParams) WithContext(ctx context.Context) *GetCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials params
func (o *GetCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) WithHTTPClient(client *http.Client) *GetCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialType adds the credentialType to the get credentials params
func (o *GetCredentialsParams) WithCredentialType(credentialType *string) *GetCredentialsParams {
	o.SetCredentialType(credentialType)
	return o
}

// SetCredentialType adds the credentialType to the get credentials params
func (o *GetCredentialsParams) SetCredentialType(credentialType *string) {
	o.CredentialType = credentialType
}

// WithOrganizationCanonical adds the organizationCanonical to the get credentials params
func (o *GetCredentialsParams) WithOrganizationCanonical(organizationCanonical string) *GetCredentialsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get credentials params
func (o *GetCredentialsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get credentials params
func (o *GetCredentialsParams) WithPageIndex(pageIndex *uint32) *GetCredentialsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get credentials params
func (o *GetCredentialsParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get credentials params
func (o *GetCredentialsParams) WithPageSize(pageSize *uint32) *GetCredentialsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get credentials params
func (o *GetCredentialsParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CredentialType != nil {

		// query param credential_type
		var qrCredentialType string
		if o.CredentialType != nil {
			qrCredentialType = *o.CredentialType
		}
		qCredentialType := qrCredentialType
		if qCredentialType != "" {
			if err := r.SetQueryParam("credential_type", qCredentialType); err != nil {
				return err
			}
		}

	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if o.PageIndex != nil {

		// query param page_index
		var qrPageIndex uint32
		if o.PageIndex != nil {
			qrPageIndex = *o.PageIndex
		}
		qPageIndex := swag.FormatUint32(qrPageIndex)
		if qPageIndex != "" {
			if err := r.SetQueryParam("page_index", qPageIndex); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize uint32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatUint32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

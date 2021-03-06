// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization credentials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization credentials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCredential Create a new Credential, based on the type you will have to pass different parameters within the body:
* ssh: ssh_key
* aws: access_key, secret_key
* gcp: json_key
* azure: client_id, client_secret, subscription_id, tenant_id
* azure_storage: account_name, access_key
* basic_auth: username, password
* elasticsearch: username, password, ca_cert
* swift: auth_url, username, password, domain_id, tenant_id
*/
func (a *Client) CreateCredential(params *CreateCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCredential",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteCredential Delete the Credential.
*/
func (a *Client) DeleteCredential(params *DeleteCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCredentialNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCredential",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/credentials/{credential_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCredentialNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCredential Get the information of the Credential.
*/
func (a *Client) GetCredential(params *GetCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*GetCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredential",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/credentials/{credential_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCredentials Return all the Credentials
*/
func (a *Client) GetCredentials(params *GetCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredentials",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateCredential Update an existing Credential, based on the type you will have to pass different parameters within the body:
* ssh: ssh_key
* aws: access_key, secret_key
* gcp: json_key
* azure: client_id, client_secret, subscription_id, tenant_id
* azure_storage: account_name, access_key
* basic_auth: username, password
* elasticsearch: username, password, ca_cert
* swift: auth_url, username, password, domain_id, tenant_id
*/
func (a *Client) UpdateCredential(params *UpdateCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCredential",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/credentials/{credential_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

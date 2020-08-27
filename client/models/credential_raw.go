// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CredentialRaw Credential Raw
//
// All the possible fields inside it
// swagger:model CredentialRaw
type CredentialRaw struct {

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// account name
	AccountName string `json:"account_name,omitempty"`

	// auth url
	AuthURL string `json:"auth_url,omitempty"`

	// ca cert
	CaCert string `json:"ca_cert,omitempty"`

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// json key
	JSONKey string `json:"json_key,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// raw
	Raw interface{} `json:"raw,omitempty"`

	// secret key
	SecretKey string `json:"secret_key,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this credential raw
func (m *CredentialRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CredentialRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialRaw) UnmarshalBinary(b []byte) error {
	var res CredentialRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
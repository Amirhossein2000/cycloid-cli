// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BillingInformationAddress Billing Information Address
//
// The address of the billing
// swagger:model BillingInformationAddress
type BillingInformationAddress struct {

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// line1
	Line1 string `json:"line1,omitempty"`

	// line2
	Line2 string `json:"line2,omitempty"`

	// postal code
	PostalCode string `json:"postal_code,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this billing information address
func (m *BillingInformationAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingInformationAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingInformationAddress) UnmarshalBinary(b []byte) error {
	var res BillingInformationAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

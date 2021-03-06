// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillingInformation Organization Billing Information
//
// The billing information of the Organization
// swagger:model BillingInformation
type BillingInformation struct {

	// address
	Address *BillingInformationAddress `json:"address,omitempty"`

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *int64 `json:"created_at"`

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// language
	Language string `json:"language,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// phone
	Phone string `json:"phone,omitempty"`

	// source
	Source *BillingInformationSource `json:"source,omitempty"`

	// subscription
	Subscription *BillingInformationSubscription `json:"subscription,omitempty"`

	// tax information
	TaxInformation *BillingInformationTaxInformation `json:"tax_information,omitempty"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *int64 `json:"updated_at"`
}

// Validate validates this billing information
func (m *BillingInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingInformation) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *BillingInformation) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformation) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformation) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformation) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *BillingInformation) validateSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscription) { // not required
		return nil
	}

	if m.Subscription != nil {
		if err := m.Subscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription")
			}
			return err
		}
	}

	return nil
}

func (m *BillingInformation) validateTaxInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxInformation) { // not required
		return nil
	}

	if m.TaxInformation != nil {
		if err := m.TaxInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_information")
			}
			return err
		}
	}

	return nil
}

func (m *BillingInformation) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingInformation) UnmarshalBinary(b []byte) error {
	var res BillingInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

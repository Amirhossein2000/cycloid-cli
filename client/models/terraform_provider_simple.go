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

// TerraformProviderSimple TerraformProviderSimple
//
// Provider of infrastrucutre without the conifg
// swagger:model TerraformProviderSimple
type TerraformProviderSimple struct {

	// abbreviation
	// Required: true
	Abbreviation *string `json:"abbreviation"`

	// canonical
	// Required: true
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// cloud
	// Required: true
	Cloud *bool `json:"cloud"`

	// image
	// Required: true
	// Format: uri
	Image *strfmt.URI `json:"image"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this terraform provider simple
func (m *TerraformProviderSimple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbbreviation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformProviderSimple) validateAbbreviation(formats strfmt.Registry) error {

	if err := validate.Required("abbreviation", "body", m.Abbreviation); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderSimple) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderSimple) validateCloud(formats strfmt.Registry) error {

	if err := validate.Required("cloud", "body", m.Cloud); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderSimple) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderSimple) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraformProviderSimple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformProviderSimple) UnmarshalBinary(b []byte) error {
	var res TerraformProviderSimple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

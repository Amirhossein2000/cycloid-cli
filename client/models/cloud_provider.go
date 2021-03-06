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

// CloudProvider Cloud Provider
//
// CloudProvider represents a cloud provider. Those cloud providers are used to identify the scope of projects and/or stacks.
// swagger:model CloudProvider
type CloudProvider struct {

	// abbreviation
	// Max Length: 60
	// Min Length: 2
	Abbreviation string `json:"abbreviation,omitempty"`

	// canonical
	// Required: true
	// Max Length: 30
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// created at
	// Minimum: 0
	CreatedAt *int64 `json:"created_at,omitempty"`

	// id
	// Minimum: 1
	ID uint32 `json:"id,omitempty"`

	// name
	// Required: true
	// Max Length: 60
	// Min Length: 2
	Name *string `json:"name"`

	// updated at
	// Minimum: 0
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

// Validate validates this cloud provider
func (m *CloudProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbbreviation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *CloudProvider) validateAbbreviation(formats strfmt.Registry) error {

	if swag.IsZero(m.Abbreviation) { // not required
		return nil
	}

	if err := validate.MinLength("abbreviation", "body", string(m.Abbreviation), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("abbreviation", "body", string(m.Abbreviation), 60); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(*m.Canonical), 30); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 60); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProvider) UnmarshalBinary(b []byte) error {
	var res CloudProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ConfigRepository ConfigRepository
// swagger:model ConfigRepository
type ConfigRepository struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// created at
	// Minimum: 0
	CreatedAt *int64 `json:"created_at,omitempty"`

	// credential id
	// Minimum: 1
	CredentialID uint32 `json:"credential_id,omitempty"`

	// default
	// Required: true
	Default *bool `json:"default"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// updated at
	// Minimum: 0
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// url
	// Required: true
	// Pattern: (?:http|https|git|ssh|git@[-\w.]+):(\/\/)?(.*?)(\.git)?(\/?|\#[-\d\w._]+?)$
	URL *string `json:"url"`
}

// Validate validates this config repository
func (m *ConfigRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefault(formats); err != nil {
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

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigRepository) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRepository) validateCredentialID(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialID) { // not required
		return nil
	}

	if err := validate.MinimumInt("credential_id", "body", int64(m.CredentialID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRepository) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRepository) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRepository) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRepository) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRepository) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", string(*m.URL), `(?:http|https|git|ssh|git@[-\w.]+):(\/\/)?(.*?)(\.git)?(\/?|\#[-\d\w._]+?)$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigRepository) UnmarshalBinary(b []byte) error {
	var res ConfigRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

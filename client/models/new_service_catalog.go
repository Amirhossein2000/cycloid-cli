// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewServiceCatalog Service Catalog
//
// Represents the Service Catalog item
// swagger:model NewServiceCatalog
type NewServiceCatalog struct {

	// author
	// Required: true
	Author *string `json:"author"`

	// canonical
	// Required: true
	// Max Length: 30
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// created at
	// Minimum: 0
	CreatedAt *int64 `json:"created_at,omitempty"`

	// dependencies
	Dependencies []*ServiceCatalogDependency `json:"dependencies"`

	// description
	// Required: true
	Description *string `json:"description"`

	// image
	// Format: uri
	Image strfmt.URI `json:"image,omitempty"`

	// keywords
	// Required: true
	Keywords []string `json:"keywords"`

	// name
	// Required: true
	Name *string `json:"name"`

	// service catalog source id
	// Required: true
	// Minimum: 1
	ServiceCatalogSourceID *uint32 `json:"service_catalog_source_id"`

	// status
	Status string `json:"status,omitempty"`

	// technologies
	Technologies []*ServiceCatalogTechnology `json:"technologies"`

	// updated at
	// Minimum: 0
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

// Validate validates this new service catalog
func (m *NewServiceCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTechnologies(formats); err != nil {
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

func (m *NewServiceCatalog) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateCanonical(formats strfmt.Registry) error {

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

func (m *NewServiceCatalog) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewServiceCatalog) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateKeywords(formats strfmt.Registry) error {

	if err := validate.Required("keywords", "body", m.Keywords); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateServiceCatalogSourceID(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_source_id", "body", m.ServiceCatalogSourceID); err != nil {
		return err
	}

	if err := validate.MinimumInt("service_catalog_source_id", "body", int64(*m.ServiceCatalogSourceID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalog) validateTechnologies(formats strfmt.Registry) error {

	if swag.IsZero(m.Technologies) { // not required
		return nil
	}

	for i := 0; i < len(m.Technologies); i++ {
		if swag.IsZero(m.Technologies[i]) { // not required
			continue
		}

		if m.Technologies[i] != nil {
			if err := m.Technologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("technologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewServiceCatalog) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewServiceCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewServiceCatalog) UnmarshalBinary(b []byte) error {
	var res NewServiceCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

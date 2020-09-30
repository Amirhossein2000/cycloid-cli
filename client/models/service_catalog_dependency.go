// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ServiceCatalogDependency ServiceCatalogDependency
//
// Service Catalog Dependency identifies ServiceCatalog being dependency of other
// swagger:model ServiceCatalogDependency
type ServiceCatalogDependency struct {

	// ref
	Ref string `json:"ref,omitempty"`

	// required
	Required bool `json:"required,omitempty"`
}

// Validate validates this service catalog dependency
func (m *ServiceCatalogDependency) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCatalogDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCatalogDependency) UnmarshalBinary(b []byte) error {
	var res ServiceCatalogDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

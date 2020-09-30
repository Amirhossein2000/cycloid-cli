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

// ServiceCatalogSource ServiceCatalogSource
// swagger:model ServiceCatalogSource
type ServiceCatalogSource struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// Represents map of service catalogs changes during the update of a service catalog source. Used only for update action on a service catalog source.
	Changes *ServiceCatalogChanges `json:"changes,omitempty"`

	// created at
	// Minimum: 0
	CreatedAt *int64 `json:"created_at,omitempty"`

	// credential id
	// Minimum: 1
	CredentialID uint32 `json:"credential_id,omitempty"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization member that owns this service catalog source. When a user is the owner of a
	// service catalog source they has all the permissions on it.
	// In the event where the user has been deleted this field might be empty.
	//
	// Required: true
	Owner *MemberOrg `json:"owner"`

	// The last time the Service Catalog Source was (successfully) refreshed.
	// Minimum: 0
	RefreshedAt *int64 `json:"refreshed_at,omitempty"`

	// Represents list of service catalogs in the service catalog source. Not used during update action on a service catalog source.
	ServiceCatalogs []*ServiceCatalog `json:"service_catalogs"`

	// stack count
	// Required: true
	// Minimum: 0
	StackCount *uint32 `json:"stack_count"`

	// updated at
	// Minimum: 0
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// url
	// Required: true
	// Pattern: (?:http|https|git|ssh|git@[-\w.]+):(\/\/)?(.*?)(\.git)?(\/?|\#[-\d\w._]+?)$
	URL *string `json:"url"`
}

// Validate validates this service catalog source
func (m *ServiceCatalogSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackCount(formats); err != nil {
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

func (m *ServiceCatalogSource) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	if m.Changes != nil {
		if err := m.Changes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changes")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogSource) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateCredentialID(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialID) { // not required
		return nil
	}

	if err := validate.MinimumInt("credential_id", "body", int64(m.CredentialID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogSource) validateRefreshedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RefreshedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("refreshed_at", "body", int64(*m.RefreshedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateServiceCatalogs(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceCatalogs) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceCatalogs); i++ {
		if swag.IsZero(m.ServiceCatalogs[i]) { // not required
			continue
		}

		if m.ServiceCatalogs[i] != nil {
			if err := m.ServiceCatalogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_catalogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceCatalogSource) validateStackCount(formats strfmt.Registry) error {

	if err := validate.Required("stack_count", "body", m.StackCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("stack_count", "body", int64(*m.StackCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogSource) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", string(*m.URL), `(?:http|https|git|ssh|git@[-\w.]+):(\/\/)?(.*?)(\.git)?(\/?|\#[-\d\w._]+?)$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCatalogSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCatalogSource) UnmarshalBinary(b []byte) error {
	var res ServiceCatalogSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

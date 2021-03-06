// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateProject Update Project
//
// The entity which represents the information of the project to be updated.
// swagger:model UpdateProject
type UpdateProject struct {

	// The cloud provider canonical that this project is using - between the
	// supported ones.
	//
	// Enum: [aws google azurerm flexibleengine openstack]
	CloudProvider string `json:"cloud_provider,omitempty"`

	// The config_repository_id points to new Config Repository the project
	// will be using. If this value is filled and it's different from the
	// current one, the whole project will be migrated to new CR, meaning
	// configuration files will also be moved.
	// If the project didn't has config_repository_id set, this action will
	// only attach the project to the CR, it won't create/move any files.
	// In order to be sure everything works, make sure the
	// config_repository_id is pointing at the CR with the same git
	// repository that was used during project creation.
	//
	ConfigRepositoryID uint32 `json:"config_repository_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// environments
	Environments []string `json:"environments"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// User canonical that owns this project. Only the owner or an
	// organization admin can update such a field. When a user is the owner
	// of a project it has all the permission on it.
	//
	Owner string `json:"owner,omitempty"`

	// It's the ref of the Service Catalog, like 'cycloidio:stack-magento'
	// Required: true
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+:[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ServiceCatalogRef *string `json:"service_catalog_ref"`
}

// Validate validates this update project
func (m *UpdateProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateProjectTypeCloudProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws","google","azurerm","flexibleengine","openstack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateProjectTypeCloudProviderPropEnum = append(updateProjectTypeCloudProviderPropEnum, v)
	}
}

const (

	// UpdateProjectCloudProviderAws captures enum value "aws"
	UpdateProjectCloudProviderAws string = "aws"

	// UpdateProjectCloudProviderGoogle captures enum value "google"
	UpdateProjectCloudProviderGoogle string = "google"

	// UpdateProjectCloudProviderAzurerm captures enum value "azurerm"
	UpdateProjectCloudProviderAzurerm string = "azurerm"

	// UpdateProjectCloudProviderFlexibleengine captures enum value "flexibleengine"
	UpdateProjectCloudProviderFlexibleengine string = "flexibleengine"

	// UpdateProjectCloudProviderOpenstack captures enum value "openstack"
	UpdateProjectCloudProviderOpenstack string = "openstack"
)

// prop value enum
func (m *UpdateProject) validateCloudProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateProjectTypeCloudProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateProject) validateCloudProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudProviderEnum("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

func (m *UpdateProject) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {

		if err := validate.Pattern("environments"+"."+strconv.Itoa(i), "body", string(m.Environments[i]), `^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateProject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	return nil
}

func (m *UpdateProject) validateServiceCatalogRef(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_ref", "body", m.ServiceCatalogRef); err != nil {
		return err
	}

	if err := validate.Pattern("service_catalog_ref", "body", string(*m.ServiceCatalogRef), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+:[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProject) UnmarshalBinary(b []byte) error {
	var res UpdateProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

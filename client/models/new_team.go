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

// NewTeam Create Team
//
// The entity which represents the information of a new team.
// swagger:model NewTeam
type NewTeam struct {

	// canonical
	// Required: true
	// Max Length: 30
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// User canonical that owns this team. If omitted then the person
	// creating this team will be assigned as owner. When a user is the
	// owner of a team it has all the permissions on it.
	//
	Owner string `json:"owner,omitempty"`

	// The roles to be assigned to a team.
	// Required: true
	RolesID []uint32 `json:"roles_id"`
}

// Validate validates this new team
func (m *NewTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolesID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewTeam) validateCanonical(formats strfmt.Registry) error {

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

func (m *NewTeam) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	return nil
}

func (m *NewTeam) validateRolesID(formats strfmt.Registry) error {

	if err := validate.Required("roles_id", "body", m.RolesID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewTeam) UnmarshalBinary(b []byte) error {
	var res NewTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

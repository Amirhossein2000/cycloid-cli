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

// UpdateUserAccountEmail User's email
//
// The email address of a user to be updated.
// swagger:model UpdateUserAccountEmail
type UpdateUserAccountEmail struct {

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// This values are set by the application to indicate the purpose of the email address. At least there is always one which is the primary.
	// Required: true
	Purpose *string `json:"purpose"`
}

// Validate validates this update user account email
func (m *UpdateUserAccountEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserAccountEmail) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserAccountEmail) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("purpose", "body", m.Purpose); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserAccountEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserAccountEmail) UnmarshalBinary(b []byte) error {
	var res UpdateUserAccountEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
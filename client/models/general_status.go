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

// GeneralStatus GeneralStatus
// swagger:model GeneralStatus
type GeneralStatus struct {

	// List of all checks report type/name/status.
	// Required: true
	Checks []*CheckReport `json:"checks"`

	// Message providing information regarding the status.
	// Required: true
	// Min Length: 1
	Message *string `json:"message"`

	// The overall status for the application.
	// Required: true
	// Enum: [Unknown Success Error]
	Status *string `json:"status"`
}

// Validate validates this general status
func (m *GeneralStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralStatus) validateChecks(formats strfmt.Registry) error {

	if err := validate.Required("checks", "body", m.Checks); err != nil {
		return err
	}

	for i := 0; i < len(m.Checks); i++ {
		if swag.IsZero(m.Checks[i]) { // not required
			continue
		}

		if m.Checks[i] != nil {
			if err := m.Checks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GeneralStatus) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", string(*m.Message), 1); err != nil {
		return err
	}

	return nil
}

var generalStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Success","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		generalStatusTypeStatusPropEnum = append(generalStatusTypeStatusPropEnum, v)
	}
}

const (

	// GeneralStatusStatusUnknown captures enum value "Unknown"
	GeneralStatusStatusUnknown string = "Unknown"

	// GeneralStatusStatusSuccess captures enum value "Success"
	GeneralStatusStatusSuccess string = "Success"

	// GeneralStatusStatusError captures enum value "Error"
	GeneralStatusStatusError string = "Error"
)

// prop value enum
func (m *GeneralStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, generalStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GeneralStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralStatus) UnmarshalBinary(b []byte) error {
	var res GeneralStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
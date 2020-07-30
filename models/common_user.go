// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonUser common user
//
// swagger:model common.user
type CommonUser struct {

	// date joined
	// Format: date
	DateJoined strfmt.Date `json:"date_joined,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// has marketing access
	HasMarketingAccess bool `json:"has_marketing_access,omitempty"`

	// has sales access
	HasSalesAccess bool `json:"has_sales_access,omitempty"`

	// is active
	IsActive bool `json:"is_active,omitempty"`

	// is admin
	IsAdmin bool `json:"is_admin,omitempty"`

	// is staff
	IsStaff bool `json:"is_staff,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// profile pic
	ProfilePic string `json:"profile_pic,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this common user
func (m *CommonUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateJoined(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonUser) validateDateJoined(formats strfmt.Registry) error {

	if swag.IsZero(m.DateJoined) { // not required
		return nil
	}

	if err := validate.FormatOf("date_joined", "body", "date", m.DateJoined.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonUser) UnmarshalBinary(b []byte) error {
	var res CommonUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonAddress common address
//
// swagger:model common.address
type CommonAddress struct {

	// address line
	AddressLine string `json:"address_line,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// postcode
	Postcode string `json:"postcode,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// street
	Street string `json:"street,omitempty"`
}

// Validate validates this common address
func (m *CommonAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonAddress) UnmarshalBinary(b []byte) error {
	var res CommonAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

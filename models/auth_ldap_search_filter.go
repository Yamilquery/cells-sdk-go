// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AuthLdapSearchFilter auth ldap search filter
// swagger:model authLdapSearchFilter
type AuthLdapSearchFilter struct {

	// d ns
	DNs []string `json:"DNs"`

	// display attribute
	DisplayAttribute string `json:"DisplayAttribute,omitempty"`

	// filter
	Filter string `json:"Filter,omitempty"`

	// ID attribute
	IDAttribute string `json:"IDAttribute,omitempty"`

	// scope
	Scope string `json:"Scope,omitempty"`
}

// Validate validates this auth ldap search filter
func (m *AuthLdapSearchFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthLdapSearchFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthLdapSearchFilter) UnmarshalBinary(b []byte) error {
	var res AuthLdapSearchFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
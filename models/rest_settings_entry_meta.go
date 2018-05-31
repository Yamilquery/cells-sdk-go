// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestSettingsEntryMeta rest settings entry meta
// swagger:model restSettingsEntryMeta
type RestSettingsEntryMeta struct {

	// component
	Component string `json:"Component,omitempty"`

	// icon class
	IconClass string `json:"IconClass,omitempty"`

	// props
	Props string `json:"Props,omitempty"`
}

// Validate validates this rest settings entry meta
func (m *RestSettingsEntryMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestSettingsEntryMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSettingsEntryMeta) UnmarshalBinary(b []byte) error {
	var res RestSettingsEntryMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

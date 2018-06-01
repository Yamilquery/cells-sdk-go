// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EncryptionImport encryption import
// swagger:model encryptionImport
type EncryptionImport struct {

	// by
	By string `json:"By,omitempty"`

	// date
	Date int32 `json:"Date,omitempty"`
}

// Validate validates this encryption import
func (m *EncryptionImport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionImport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionImport) UnmarshalBinary(b []byte) error {
	var res EncryptionImport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EncryptionAdminImportKeyRequest encryption admin import key request
// swagger:model encryptionAdminImportKeyRequest
type EncryptionAdminImportKeyRequest struct {

	// key
	Key *EncryptionKey `json:"Key,omitempty"`

	// override
	Override bool `json:"Override,omitempty"`

	// str password
	StrPassword string `json:"StrPassword,omitempty"`
}

// Validate validates this encryption admin import key request
func (m *EncryptionAdminImportKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionAdminImportKeyRequest) validateKey(formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {
		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionAdminImportKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionAdminImportKeyRequest) UnmarshalBinary(b []byte) error {
	var res EncryptionAdminImportKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

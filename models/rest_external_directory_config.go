// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RestExternalDirectoryConfig rest external directory config
// swagger:model restExternalDirectoryConfig
type RestExternalDirectoryConfig struct {

	// config
	Config *AuthLdapServerConfig `json:"Config,omitempty"`

	// config Id
	ConfigID string `json:"ConfigId,omitempty"`
}

// Validate validates this rest external directory config
func (m *RestExternalDirectoryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestExternalDirectoryConfig) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestExternalDirectoryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestExternalDirectoryConfig) UnmarshalBinary(b []byte) error {
	var res RestExternalDirectoryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestResetPasswordRequest rest reset password request
// swagger:model restResetPasswordRequest
type RestResetPasswordRequest struct {

	// new password
	NewPassword string `json:"NewPassword,omitempty"`

	// reset password token
	ResetPasswordToken string `json:"ResetPasswordToken,omitempty"`

	// user login
	UserLogin string `json:"UserLogin,omitempty"`
}

// Validate validates this rest reset password request
func (m *RestResetPasswordRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestResetPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestResetPasswordRequest) UnmarshalBinary(b []byte) error {
	var res RestResetPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

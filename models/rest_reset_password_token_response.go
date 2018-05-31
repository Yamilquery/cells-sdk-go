// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestResetPasswordTokenResponse rest reset password token response
// swagger:model restResetPasswordTokenResponse
type RestResetPasswordTokenResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this rest reset password token response
func (m *RestResetPasswordTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestResetPasswordTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestResetPasswordTokenResponse) UnmarshalBinary(b []byte) error {
	var res RestResetPasswordTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

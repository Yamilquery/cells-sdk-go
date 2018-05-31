// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestRevokeResponse Rest response
// swagger:model restRevokeResponse
type RestRevokeResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this rest revoke response
func (m *RestRevokeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestRevokeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestRevokeResponse) UnmarshalBinary(b []byte) error {
	var res RestRevokeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

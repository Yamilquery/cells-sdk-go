// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MailerSendMailResponse mailer send mail response
// swagger:model mailerSendMailResponse
type MailerSendMailResponse struct {

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this mailer send mail response
func (m *MailerSendMailResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MailerSendMailResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailerSendMailResponse) UnmarshalBinary(b []byte) error {
	var res MailerSendMailResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

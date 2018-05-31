// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DocstorePutDocumentResponse docstore put document response
// swagger:model docstorePutDocumentResponse
type DocstorePutDocumentResponse struct {

	// document
	Document *DocstoreDocument `json:"Document,omitempty"`
}

// Validate validates this docstore put document response
func (m *DocstorePutDocumentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocstorePutDocumentResponse) validateDocument(formats strfmt.Registry) error {

	if swag.IsZero(m.Document) { // not required
		return nil
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Document")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocstorePutDocumentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocstorePutDocumentResponse) UnmarshalBinary(b []byte) error {
	var res DocstorePutDocumentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

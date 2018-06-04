// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TreeChangeLog tree change log
// swagger:model treeChangeLog
type TreeChangeLog struct {

	// Arbitrary additional data
	// Format: byte
	Data strfmt.Base64 `json:"Data,omitempty"`

	// Human-readable description of what happened
	Description string `json:"Description,omitempty"`

	// Event that triggered this change
	Event *TreeNodeChangeEvent `json:"Event,omitempty"`

	// Unix Timestamp
	MTime int64 `json:"MTime,string,omitempty"`

	// Who performed this action
	OwnerUUID string `json:"OwnerUuid,omitempty"`

	// Content Size at that moment
	Size int64 `json:"Size,string,omitempty"`

	// Unique commit ID
	UUID string `json:"Uuid,omitempty"`
}

// Validate validates this tree change log
func (m *TreeChangeLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeChangeLog) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if err := validate.FormatOf("Data", "body", "byte", m.Data.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TreeChangeLog) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Event")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreeChangeLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeChangeLog) UnmarshalBinary(b []byte) error {
	var res TreeChangeLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

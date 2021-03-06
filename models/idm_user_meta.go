// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdmUserMeta Piece of metadata attached to a node
// swagger:model idmUserMeta
type IdmUserMeta struct {

	// Json value
	JSONValue string `json:"JsonValue,omitempty"`

	// namespace
	Namespace string `json:"Namespace,omitempty"`

	// node Uuid
	NodeUUID string `json:"NodeUuid,omitempty"`

	// policies
	Policies []*ServiceResourcePolicy `json:"Policies"`

	// policies context editable
	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`

	// Uuid
	UUID string `json:"Uuid,omitempty"`
}

// Validate validates this idm user meta
func (m *IdmUserMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmUserMeta) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmUserMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmUserMeta) UnmarshalBinary(b []byte) error {
	var res IdmUserMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

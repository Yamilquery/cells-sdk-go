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

// RestCell Model for representing a shared room
// swagger:model restCell
type RestCell struct {

	// acls
	Acls RestCellAcls `json:"ACLs,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// label
	Label string `json:"Label,omitempty"`

	// policies
	Policies []*ServiceResourcePolicy `json:"Policies"`

	// policies context editable
	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`

	// root nodes
	RootNodes []*TreeNode `json:"RootNodes"`

	// Uuid
	UUID string `json:"Uuid,omitempty"`
}

// Validate validates this rest cell
func (m *RestCell) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestCell) validateAcls(formats strfmt.Registry) error {

	if swag.IsZero(m.Acls) { // not required
		return nil
	}

	if err := m.Acls.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ACLs")
		}
		return err
	}

	return nil
}

func (m *RestCell) validatePolicies(formats strfmt.Registry) error {

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

func (m *RestCell) validateRootNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.RootNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.RootNodes); i++ {
		if swag.IsZero(m.RootNodes[i]) { // not required
			continue
		}

		if m.RootNodes[i] != nil {
			if err := m.RootNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RootNodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestCell) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestCell) UnmarshalBinary(b []byte) error {
	var res RestCell
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

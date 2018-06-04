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

// ServiceQuery service query
// swagger:model serviceQuery
type ServiceQuery struct {

	// limit
	Limit int64 `json:"Limit,string,omitempty"`

	// offset
	Offset int64 `json:"Offset,string,omitempty"`

	// operation
	Operation ServiceOperationType `json:"Operation,omitempty"`

	// resource policy query
	ResourcePolicyQuery *ServiceResourcePolicyQuery `json:"ResourcePolicyQuery,omitempty"`

	// sub queries
	SubQueries []*ProtobufAny `json:"SubQueries"`

	// group by
	GroupBy int32 `json:"groupBy,omitempty"`
}

// Validate validates this service query
func (m *ServiceQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePolicyQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubQueries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceQuery) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if err := m.Operation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Operation")
		}
		return err
	}

	return nil
}

func (m *ServiceQuery) validateResourcePolicyQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourcePolicyQuery) { // not required
		return nil
	}

	if m.ResourcePolicyQuery != nil {
		if err := m.ResourcePolicyQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourcePolicyQuery")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceQuery) validateSubQueries(formats strfmt.Registry) error {

	if swag.IsZero(m.SubQueries) { // not required
		return nil
	}

	for i := 0; i < len(m.SubQueries); i++ {
		if swag.IsZero(m.SubQueries[i]) { // not required
			continue
		}

		if m.SubQueries[i] != nil {
			if err := m.SubQueries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SubQueries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceQuery) UnmarshalBinary(b []byte) error {
	var res ServiceQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

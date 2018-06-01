// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// UpdateUserMetaNamespaceRequestUserMetaNsOp update user meta namespace request user meta ns op
// swagger:model UpdateUserMetaNamespaceRequestUserMetaNsOp
type UpdateUserMetaNamespaceRequestUserMetaNsOp string

const (

	// UpdateUserMetaNamespaceRequestUserMetaNsOpPUT captures enum value "PUT"
	UpdateUserMetaNamespaceRequestUserMetaNsOpPUT UpdateUserMetaNamespaceRequestUserMetaNsOp = "PUT"

	// UpdateUserMetaNamespaceRequestUserMetaNsOpDELETE captures enum value "DELETE"
	UpdateUserMetaNamespaceRequestUserMetaNsOpDELETE UpdateUserMetaNamespaceRequestUserMetaNsOp = "DELETE"
)

// for schema
var updateUserMetaNamespaceRequestUserMetaNsOpEnum []interface{}

func init() {
	var res []UpdateUserMetaNamespaceRequestUserMetaNsOp
	if err := json.Unmarshal([]byte(`["PUT","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateUserMetaNamespaceRequestUserMetaNsOpEnum = append(updateUserMetaNamespaceRequestUserMetaNsOpEnum, v)
	}
}

func (m UpdateUserMetaNamespaceRequestUserMetaNsOp) validateUpdateUserMetaNamespaceRequestUserMetaNsOpEnum(path, location string, value UpdateUserMetaNamespaceRequestUserMetaNsOp) error {
	if err := validate.Enum(path, location, value, updateUserMetaNamespaceRequestUserMetaNsOpEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this update user meta namespace request user meta ns op
func (m UpdateUserMetaNamespaceRequestUserMetaNsOp) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUpdateUserMetaNamespaceRequestUserMetaNsOpEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
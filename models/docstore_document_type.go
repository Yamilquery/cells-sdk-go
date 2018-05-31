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

// DocstoreDocumentType docstore document type
// swagger:model docstoreDocumentType
type DocstoreDocumentType string

const (

	// DocstoreDocumentTypeJSON captures enum value "JSON"
	DocstoreDocumentTypeJSON DocstoreDocumentType = "JSON"

	// DocstoreDocumentTypeBINARY captures enum value "BINARY"
	DocstoreDocumentTypeBINARY DocstoreDocumentType = "BINARY"
)

// for schema
var docstoreDocumentTypeEnum []interface{}

func init() {
	var res []DocstoreDocumentType
	if err := json.Unmarshal([]byte(`["JSON","BINARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		docstoreDocumentTypeEnum = append(docstoreDocumentTypeEnum, v)
	}
}

func (m DocstoreDocumentType) validateDocstoreDocumentTypeEnum(path, location string, value DocstoreDocumentType) error {
	if err := validate.Enum(path, location, value, docstoreDocumentTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this docstore document type
func (m DocstoreDocumentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDocstoreDocumentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

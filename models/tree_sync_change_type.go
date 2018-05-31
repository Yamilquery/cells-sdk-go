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

// TreeSyncChangeType tree sync change type
// swagger:model treeSyncChangeType
type TreeSyncChangeType string

const (

	// TreeSyncChangeTypeUnknown captures enum value "unknown"
	TreeSyncChangeTypeUnknown TreeSyncChangeType = "unknown"

	// TreeSyncChangeTypeCreate captures enum value "create"
	TreeSyncChangeTypeCreate TreeSyncChangeType = "create"

	// TreeSyncChangeTypeDelete captures enum value "delete"
	TreeSyncChangeTypeDelete TreeSyncChangeType = "delete"

	// TreeSyncChangeTypePath captures enum value "path"
	TreeSyncChangeTypePath TreeSyncChangeType = "path"

	// TreeSyncChangeTypeContent captures enum value "content"
	TreeSyncChangeTypeContent TreeSyncChangeType = "content"
)

// for schema
var treeSyncChangeTypeEnum []interface{}

func init() {
	var res []TreeSyncChangeType
	if err := json.Unmarshal([]byte(`["unknown","create","delete","path","content"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		treeSyncChangeTypeEnum = append(treeSyncChangeTypeEnum, v)
	}
}

func (m TreeSyncChangeType) validateTreeSyncChangeTypeEnum(path, location string, value TreeSyncChangeType) error {
	if err := validate.Enum(path, location, value, treeSyncChangeTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tree sync change type
func (m TreeSyncChangeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTreeSyncChangeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

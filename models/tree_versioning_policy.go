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

// TreeVersioningPolicy tree versioning policy
// swagger:model treeVersioningPolicy
type TreeVersioningPolicy struct {

	// description
	Description string `json:"Description,omitempty"`

	// ignore files greater than
	IgnoreFilesGreaterThan int64 `json:"IgnoreFilesGreaterThan,omitempty"`

	// keep periods
	KeepPeriods []*TreeVersioningKeepPeriod `json:"KeepPeriods"`

	// max size per file
	MaxSizePerFile int64 `json:"MaxSizePerFile,omitempty"`

	// max total size
	MaxTotalSize int64 `json:"MaxTotalSize,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// Uuid
	UUID string `json:"Uuid,omitempty"`

	// versions data source bucket
	VersionsDataSourceBucket string `json:"VersionsDataSourceBucket,omitempty"`

	// versions data source name
	VersionsDataSourceName string `json:"VersionsDataSourceName,omitempty"`
}

// Validate validates this tree versioning policy
func (m *TreeVersioningPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeepPeriods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeVersioningPolicy) validateKeepPeriods(formats strfmt.Registry) error {

	if swag.IsZero(m.KeepPeriods) { // not required
		return nil
	}

	for i := 0; i < len(m.KeepPeriods); i++ {
		if swag.IsZero(m.KeepPeriods[i]) { // not required
			continue
		}

		if m.KeepPeriods[i] != nil {
			if err := m.KeepPeriods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("KeepPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreeVersioningPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeVersioningPolicy) UnmarshalBinary(b []byte) error {
	var res TreeVersioningPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

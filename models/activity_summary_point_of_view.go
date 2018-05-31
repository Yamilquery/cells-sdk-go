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

// ActivitySummaryPointOfView activity summary point of view
// swagger:model activitySummaryPointOfView
type ActivitySummaryPointOfView string

const (

	// ActivitySummaryPointOfViewGENERIC captures enum value "GENERIC"
	ActivitySummaryPointOfViewGENERIC ActivitySummaryPointOfView = "GENERIC"

	// ActivitySummaryPointOfViewACTOR captures enum value "ACTOR"
	ActivitySummaryPointOfViewACTOR ActivitySummaryPointOfView = "ACTOR"

	// ActivitySummaryPointOfViewSUBJECT captures enum value "SUBJECT"
	ActivitySummaryPointOfViewSUBJECT ActivitySummaryPointOfView = "SUBJECT"
)

// for schema
var activitySummaryPointOfViewEnum []interface{}

func init() {
	var res []ActivitySummaryPointOfView
	if err := json.Unmarshal([]byte(`["GENERIC","ACTOR","SUBJECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activitySummaryPointOfViewEnum = append(activitySummaryPointOfViewEnum, v)
	}
}

func (m ActivitySummaryPointOfView) validateActivitySummaryPointOfViewEnum(path, location string, value ActivitySummaryPointOfView) error {
	if err := validate.Enum(path, location, value, activitySummaryPointOfViewEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity summary point of view
func (m ActivitySummaryPointOfView) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivitySummaryPointOfViewEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

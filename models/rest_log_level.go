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

// RestLogLevel Frontend Log Level
// swagger:model restLogLevel
type RestLogLevel string

const (

	// RestLogLevelDEBUG captures enum value "DEBUG"
	RestLogLevelDEBUG RestLogLevel = "DEBUG"

	// RestLogLevelINFO captures enum value "INFO"
	RestLogLevelINFO RestLogLevel = "INFO"

	// RestLogLevelNOTICE captures enum value "NOTICE"
	RestLogLevelNOTICE RestLogLevel = "NOTICE"

	// RestLogLevelWARNING captures enum value "WARNING"
	RestLogLevelWARNING RestLogLevel = "WARNING"

	// RestLogLevelERROR captures enum value "ERROR"
	RestLogLevelERROR RestLogLevel = "ERROR"
)

// for schema
var restLogLevelEnum []interface{}

func init() {
	var res []RestLogLevel
	if err := json.Unmarshal([]byte(`["DEBUG","INFO","NOTICE","WARNING","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restLogLevelEnum = append(restLogLevelEnum, v)
	}
}

func (m RestLogLevel) validateRestLogLevelEnum(path, location string, value RestLogLevel) error {
	if err := validate.Enum(path, location, value, restLogLevelEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this rest log level
func (m RestLogLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRestLogLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

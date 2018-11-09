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

// JobsJob jobs job
// swagger:model jobsJob
type JobsJob struct {

	// Chain of actions to perform
	Actions []*JobsAction `json:"Actions"`

	// Remove job automatically once it is finished (success only)
	AutoClean bool `json:"AutoClean,omitempty"`

	// Start task as soon as job is inserted
	AutoStart bool `json:"AutoStart,omitempty"`

	// How the job will be triggered.
	// One of these must be set (not exclusive)
	// Listen to a given set of events
	EventNames []string `json:"EventNames"`

	// Unique ID for this Job
	ID string `json:"ID,omitempty"`

	// Admin can temporarily disable this job
	Inactive bool `json:"Inactive,omitempty"`

	// Human-readable Label
	Label string `json:"Label,omitempty"`

	// Optional list of languages detected in the context at launch time
	Languages []string `json:"Languages"`

	// Task properties
	MaxConcurrency int32 `json:"MaxConcurrency,omitempty"`

	// Who created this Job
	Owner string `json:"Owner,omitempty"`

	// Schedule a periodic repetition
	Schedule *JobsSchedule `json:"Schedule,omitempty"`

	// Filled with currently running tasks
	Tasks []*JobsTask `json:"Tasks"`

	// Do not send notification on task update
	TasksSilentUpdate bool `json:"TasksSilentUpdate,omitempty"`
}

// Validate validates this jobs job
func (m *JobsJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsJob) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schedule")
			}
			return err
		}
	}

	return nil
}

func (m *JobsJob) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsJob) UnmarshalBinary(b []byte) error {
	var res JobsJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

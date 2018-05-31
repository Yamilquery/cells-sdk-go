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

// RestShareLink Model for representing a public link
// swagger:model restShareLink
type RestShareLink struct {

	// access end
	AccessEnd int64 `json:"AccessEnd,omitempty"`

	// access start
	AccessStart int64 `json:"AccessStart,omitempty"`

	// current downloads
	CurrentDownloads int64 `json:"CurrentDownloads,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// label
	Label string `json:"Label,omitempty"`

	// link hash
	LinkHash string `json:"LinkHash,omitempty"`

	// link Url
	LinkURL string `json:"LinkUrl,omitempty"`

	// max downloads
	MaxDownloads int64 `json:"MaxDownloads,omitempty"`

	// password required
	PasswordRequired bool `json:"PasswordRequired,omitempty"`

	// permissions
	Permissions []RestShareLinkAccessType `json:"Permissions"`

	// policies
	Policies []*ServiceResourcePolicy `json:"Policies"`

	// policies context editable
	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`

	// restrict to target users
	RestrictToTargetUsers bool `json:"RestrictToTargetUsers,omitempty"`

	// root nodes
	RootNodes []*TreeNode `json:"RootNodes"`

	// target users
	TargetUsers RestShareLinkTargetUsers `json:"TargetUsers,omitempty"`

	// user login
	UserLogin string `json:"UserLogin,omitempty"`

	// user Uuid
	UserUUID string `json:"UserUuid,omitempty"`

	// Uuid
	UUID string `json:"Uuid,omitempty"`

	// view template name
	ViewTemplateName string `json:"ViewTemplateName,omitempty"`
}

// Validate validates this rest share link
func (m *RestShareLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestShareLink) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {

		if err := m.Permissions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Permissions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *RestShareLink) validatePolicies(formats strfmt.Registry) error {

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

func (m *RestShareLink) validateRootNodes(formats strfmt.Registry) error {

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

func (m *RestShareLink) validateTargetUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetUsers) { // not required
		return nil
	}

	if err := m.TargetUsers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TargetUsers")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestShareLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestShareLink) UnmarshalBinary(b []byte) error {
	var res RestShareLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

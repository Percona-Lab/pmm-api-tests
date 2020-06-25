// Code generated by go-swagger; DO NOT EDIT.

package ammodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VersionInfo version info
//
// swagger:model versionInfo
type VersionInfo struct {

	// branch
	// Required: true
	Branch *string `json:"branch"`

	// build date
	// Required: true
	BuildDate *string `json:"buildDate"`

	// build user
	// Required: true
	BuildUser *string `json:"buildUser"`

	// go version
	// Required: true
	GoVersion *string `json:"goVersion"`

	// revision
	// Required: true
	Revision *string `json:"revision"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this version info
func (m *VersionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionInfo) validateBranch(formats strfmt.Registry) error {

	if err := validate.Required("branch", "body", m.Branch); err != nil {
		return err
	}

	return nil
}

func (m *VersionInfo) validateBuildDate(formats strfmt.Registry) error {

	if err := validate.Required("buildDate", "body", m.BuildDate); err != nil {
		return err
	}

	return nil
}

func (m *VersionInfo) validateBuildUser(formats strfmt.Registry) error {

	if err := validate.Required("buildUser", "body", m.BuildUser); err != nil {
		return err
	}

	return nil
}

func (m *VersionInfo) validateGoVersion(formats strfmt.Registry) error {

	if err := validate.Required("goVersion", "body", m.GoVersion); err != nil {
		return err
	}

	return nil
}

func (m *VersionInfo) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	return nil
}

func (m *VersionInfo) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionInfo) UnmarshalBinary(b []byte) error {
	var res VersionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

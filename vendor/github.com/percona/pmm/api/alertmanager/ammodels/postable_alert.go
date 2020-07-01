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

// PostableAlert postable alert
//
// swagger:model postableAlert
type PostableAlert struct {

	// annotations
	Annotations LabelSet `json:"annotations,omitempty"`

	// ends at
	// Format: date-time
	EndsAt strfmt.DateTime `json:"endsAt,omitempty"`

	// starts at
	// Format: date-time
	StartsAt strfmt.DateTime `json:"startsAt,omitempty"`

	Alert
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostableAlert) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Annotations LabelSet `json:"annotations,omitempty"`

		EndsAt strfmt.DateTime `json:"endsAt,omitempty"`

		StartsAt strfmt.DateTime `json:"startsAt,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Annotations = dataAO0.Annotations

	m.EndsAt = dataAO0.EndsAt

	m.StartsAt = dataAO0.StartsAt

	// AO1
	var aO1 Alert
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Alert = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostableAlert) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Annotations LabelSet `json:"annotations,omitempty"`

		EndsAt strfmt.DateTime `json:"endsAt,omitempty"`

		StartsAt strfmt.DateTime `json:"startsAt,omitempty"`
	}

	dataAO0.Annotations = m.Annotations

	dataAO0.EndsAt = m.EndsAt

	dataAO0.StartsAt = m.StartsAt

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.Alert)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this postable alert
func (m *PostableAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with Alert
	if err := m.Alert.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAlert) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if err := m.Annotations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *PostableAlert) validateEndsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EndsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("endsAt", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostableAlert) validateStartsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startsAt", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostableAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostableAlert) UnmarshalBinary(b []byte) error {
	var res PostableAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

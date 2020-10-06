// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangeSettingsReader is a Reader for the ChangeSettings structure.
type ChangeSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeSettingsOK creates a ChangeSettingsOK with default headers values
func NewChangeSettingsOK() *ChangeSettingsOK {
	return &ChangeSettingsOK{}
}

/*ChangeSettingsOK handles this case with default header values.

A successful response.
*/
type ChangeSettingsOK struct {
	Payload *ChangeSettingsOKBody
}

func (o *ChangeSettingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Change][%d] changeSettingsOk  %+v", 200, o.Payload)
}

func (o *ChangeSettingsOK) GetPayload() *ChangeSettingsOKBody {
	return o.Payload
}

func (o *ChangeSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeSettingsDefault creates a ChangeSettingsDefault with default headers values
func NewChangeSettingsDefault(code int) *ChangeSettingsDefault {
	return &ChangeSettingsDefault{
		_statusCode: code,
	}
}

/*ChangeSettingsDefault handles this case with default header values.

An unexpected error response
*/
type ChangeSettingsDefault struct {
	_statusCode int

	Payload *ChangeSettingsDefaultBody
}

// Code gets the status code for the change settings default response
func (o *ChangeSettingsDefault) Code() int {
	return o._statusCode
}

func (o *ChangeSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Change][%d] ChangeSettings default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeSettingsDefault) GetPayload() *ChangeSettingsDefaultBody {
	return o.Payload
}

func (o *ChangeSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeSettingsBody change settings body
swagger:model ChangeSettingsBody
*/
type ChangeSettingsBody struct {

	// enable telemetry
	EnableTelemetry bool `json:"enable_telemetry,omitempty"`

	// disable telemetry
	DisableTelemetry bool `json:"disable_telemetry,omitempty"`

	// A number of full days for Prometheus and QAN data retention. Should have a suffix in JSON: 2592000s, 43200m, 720h.
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// Prometheus AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// remove alert manager url
	RemoveAlertManagerURL bool `json:"remove_alert_manager_url,omitempty"`

	// alert manager rules
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// remove alert manager rules
	RemoveAlertManagerRules bool `json:"remove_alert_manager_rules,omitempty"`

	// Enable Security Threat Tool
	EnableStt bool `json:"enable_stt,omitempty"`

	// Disable Security Threat Tool
	DisableStt bool `json:"disable_stt,omitempty"`

	// metrics resolutions
	MetricsResolutions *ChangeSettingsParamsBodyMetricsResolutions `json:"metrics_resolutions,omitempty"`
}

// Validate validates this change settings body
func (o *ChangeSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsBody) validateMetricsResolutions(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsDefaultBody change settings default body
swagger:model ChangeSettingsDefaultBody
*/
type ChangeSettingsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change settings default body
func (o *ChangeSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBody change settings OK body
swagger:model ChangeSettingsOKBody
*/
type ChangeSettingsOKBody struct {

	// settings
	Settings *ChangeSettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this change settings OK body
func (o *ChangeSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBody) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettings Settings represents PMM Server settings.
swagger:model ChangeSettingsOKBodySettings
*/
type ChangeSettingsOKBodySettings struct {

	// updates disabled
	UpdatesDisabled bool `json:"updates_disabled,omitempty"`

	// telemetry enabled
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// data retention
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// Prometheus AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// alert manager rules
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// Security Threat Tool enabled
	SttEnabled bool `json:"stt_enabled,omitempty"`

	// Percona Platform user's email, if this PMM instance is linked to the Platform.
	PlatformEmail string `json:"platform_email,omitempty"`

	// DBaaS enabled
	DbaasEnabled bool `json:"dbaas_enabled,omitempty"`

	// metrics resolutions
	MetricsResolutions *ChangeSettingsOKBodySettingsMetricsResolutions `json:"metrics_resolutions,omitempty"`
}

// Validate validates this change settings OK body settings
func (o *ChangeSettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBodySettings) validateMetricsResolutions(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettingsMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model ChangeSettingsOKBodySettingsMetricsResolutions
*/
type ChangeSettingsOKBodySettingsMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this change settings OK body settings metrics resolutions
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsParamsBodyMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model ChangeSettingsParamsBodyMetricsResolutions
*/
type ChangeSettingsParamsBodyMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this change settings params body metrics resolutions
func (o *ChangeSettingsParamsBodyMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodyMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

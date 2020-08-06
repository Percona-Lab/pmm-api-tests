// Code generated by go-swagger; DO NOT EDIT.

package security_checks

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

// GetSecurityCheckResultsReader is a Reader for the GetSecurityCheckResults structure.
type GetSecurityCheckResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityCheckResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityCheckResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSecurityCheckResultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSecurityCheckResultsOK creates a GetSecurityCheckResultsOK with default headers values
func NewGetSecurityCheckResultsOK() *GetSecurityCheckResultsOK {
	return &GetSecurityCheckResultsOK{}
}

/*GetSecurityCheckResultsOK handles this case with default header values.

A successful response.
*/
type GetSecurityCheckResultsOK struct {
	Payload *GetSecurityCheckResultsOKBody
}

func (o *GetSecurityCheckResultsOK) Error() string {
	return fmt.Sprintf("[GET /v1/management/SecurityChecks/GetCheckResults][%d] getSecurityCheckResultsOk  %+v", 200, o.Payload)
}

func (o *GetSecurityCheckResultsOK) GetPayload() *GetSecurityCheckResultsOKBody {
	return o.Payload
}

func (o *GetSecurityCheckResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSecurityCheckResultsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityCheckResultsDefault creates a GetSecurityCheckResultsDefault with default headers values
func NewGetSecurityCheckResultsDefault(code int) *GetSecurityCheckResultsDefault {
	return &GetSecurityCheckResultsDefault{
		_statusCode: code,
	}
}

/*GetSecurityCheckResultsDefault handles this case with default header values.

An unexpected error response
*/
type GetSecurityCheckResultsDefault struct {
	_statusCode int

	Payload *GetSecurityCheckResultsDefaultBody
}

// Code gets the status code for the get security check results default response
func (o *GetSecurityCheckResultsDefault) Code() int {
	return o._statusCode
}

func (o *GetSecurityCheckResultsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/management/SecurityChecks/GetCheckResults][%d] GetSecurityCheckResults default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecurityCheckResultsDefault) GetPayload() *GetSecurityCheckResultsDefaultBody {
	return o.Payload
}

func (o *GetSecurityCheckResultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSecurityCheckResultsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSecurityCheckResultsDefaultBody get security check results default body
swagger:model GetSecurityCheckResultsDefaultBody
*/
type GetSecurityCheckResultsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get security check results default body
func (o *GetSecurityCheckResultsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityCheckResultsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetSecurityCheckResults default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityCheckResultsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityCheckResultsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityCheckResultsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSecurityCheckResultsOKBody get security check results OK body
swagger:model GetSecurityCheckResultsOKBody
*/
type GetSecurityCheckResultsOKBody struct {

	// results
	Results []*ResultsItems0 `json:"results"`
}

// Validate validates this get security check results OK body
func (o *GetSecurityCheckResultsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityCheckResultsOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getSecurityCheckResultsOk" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityCheckResultsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityCheckResultsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityCheckResultsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ResultsItems0 STTCheckResult represents the check result returned from pmm-managed after running the check.
swagger:model ResultsItems0
*/
type ResultsItems0 struct {

	// summary
	Summary string `json:"summary,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// severity
	Severity int32 `json:"severity,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`
}

// Validate validates this results items0
func (o *ResultsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResultsItems0) UnmarshalBinary(b []byte) error {
	var res ResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

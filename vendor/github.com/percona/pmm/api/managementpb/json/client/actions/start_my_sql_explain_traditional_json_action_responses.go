// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// StartMySQLExplainTraditionalJSONActionReader is a Reader for the StartMySQLExplainTraditionalJSONAction structure.
type StartMySQLExplainTraditionalJSONActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLExplainTraditionalJSONActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartMySQLExplainTraditionalJSONActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewStartMySQLExplainTraditionalJSONActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLExplainTraditionalJSONActionOK creates a StartMySQLExplainTraditionalJSONActionOK with default headers values
func NewStartMySQLExplainTraditionalJSONActionOK() *StartMySQLExplainTraditionalJSONActionOK {
	return &StartMySQLExplainTraditionalJSONActionOK{}
}

/*StartMySQLExplainTraditionalJSONActionOK handles this case with default header values.

A successful response.
*/
type StartMySQLExplainTraditionalJSONActionOK struct {
	Payload *StartMySQLExplainTraditionalJSONActionOKBody
}

func (o *StartMySQLExplainTraditionalJSONActionOK) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartMySQLExplainTraditionalJSON][%d] startMySqlExplainTraditionalJsonActionOk  %+v", 200, o.Payload)
}

func (o *StartMySQLExplainTraditionalJSONActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLExplainTraditionalJSONActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLExplainTraditionalJSONActionDefault creates a StartMySQLExplainTraditionalJSONActionDefault with default headers values
func NewStartMySQLExplainTraditionalJSONActionDefault(code int) *StartMySQLExplainTraditionalJSONActionDefault {
	return &StartMySQLExplainTraditionalJSONActionDefault{
		_statusCode: code,
	}
}

/*StartMySQLExplainTraditionalJSONActionDefault handles this case with default header values.

An error response.
*/
type StartMySQLExplainTraditionalJSONActionDefault struct {
	_statusCode int

	Payload *StartMySQLExplainTraditionalJSONActionDefaultBody
}

// Code gets the status code for the start my SQL explain traditional JSON action default response
func (o *StartMySQLExplainTraditionalJSONActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLExplainTraditionalJSONActionDefault) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartMySQLExplainTraditionalJSON][%d] StartMySQLExplainTraditionalJSONAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMySQLExplainTraditionalJSONActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLExplainTraditionalJSONActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMySQLExplainTraditionalJSONActionBody start my SQL explain traditional JSON action body
swagger:model StartMySQLExplainTraditionalJSONActionBody
*/
type StartMySQLExplainTraditionalJSONActionBody struct {

	// The database name for EXPLAIN query.
	Database string `json:"database,omitempty"`

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// query example.
	Query string `json:"query,omitempty"`

	// Service ID for this Action.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this start my SQL explain traditional JSON action body
func (o *StartMySQLExplainTraditionalJSONActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainTraditionalJSONActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainTraditionalJSONActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainTraditionalJSONActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainTraditionalJSONActionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model StartMySQLExplainTraditionalJSONActionDefaultBody
*/
type StartMySQLExplainTraditionalJSONActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start my SQL explain traditional JSON action default body
func (o *StartMySQLExplainTraditionalJSONActionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainTraditionalJSONActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainTraditionalJSONActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainTraditionalJSONActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainTraditionalJSONActionOKBody start my SQL explain traditional JSON action OK body
swagger:model StartMySQLExplainTraditionalJSONActionOKBody
*/
type StartMySQLExplainTraditionalJSONActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL explain traditional JSON action OK body
func (o *StartMySQLExplainTraditionalJSONActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainTraditionalJSONActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainTraditionalJSONActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainTraditionalJSONActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

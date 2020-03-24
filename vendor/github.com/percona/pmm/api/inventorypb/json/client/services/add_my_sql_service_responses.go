// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMySQLServiceReader is a Reader for the AddMySQLService structure.
type AddMySQLServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMySQLServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMySQLServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMySQLServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMySQLServiceOK creates a AddMySQLServiceOK with default headers values
func NewAddMySQLServiceOK() *AddMySQLServiceOK {
	return &AddMySQLServiceOK{}
}

/*AddMySQLServiceOK handles this case with default header values.

A successful response.
*/
type AddMySQLServiceOK struct {
	Payload *AddMySQLServiceOKBody
}

func (o *AddMySQLServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMySQL][%d] addMySqlServiceOk  %+v", 200, o.Payload)
}

func (o *AddMySQLServiceOK) GetPayload() *AddMySQLServiceOKBody {
	return o.Payload
}

func (o *AddMySQLServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySQLServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMySQLServiceDefault creates a AddMySQLServiceDefault with default headers values
func NewAddMySQLServiceDefault(code int) *AddMySQLServiceDefault {
	return &AddMySQLServiceDefault{
		_statusCode: code,
	}
}

/*AddMySQLServiceDefault handles this case with default header values.

An unexpected error response
*/
type AddMySQLServiceDefault struct {
	_statusCode int

	Payload *AddMySQLServiceDefaultBody
}

// Code gets the status code for the add my SQL service default response
func (o *AddMySQLServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddMySQLServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMySQL][%d] AddMySQLService default  %+v", o._statusCode, o.Payload)
}

func (o *AddMySQLServiceDefault) GetPayload() *AddMySQLServiceDefaultBody {
	return o.Payload
}

func (o *AddMySQLServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySQLServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMySQLServiceBody add my SQL service body
swagger:model AddMySQLServiceBody
*/
type AddMySQLServiceBody struct {

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add my SQL service body
func (o *AddMySQLServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLServiceBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLServiceDefaultBody add my SQL service default body
swagger:model AddMySQLServiceDefaultBody
*/
type AddMySQLServiceDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add my SQL service default body
func (o *AddMySQLServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySQLServiceDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddMySQLService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLServiceOKBody add my SQL service OK body
swagger:model AddMySQLServiceOKBody
*/
type AddMySQLServiceOKBody struct {

	// mysql
	Mysql *AddMySQLServiceOKBodyMysql `json:"mysql,omitempty"`
}

// Validate validates this add my SQL service OK body
func (o *AddMySQLServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySQLServiceOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	if o.Mysql != nil {
		if err := o.Mysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySqlServiceOk" + "." + "mysql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLServiceOKBodyMysql MySQLService represents a generic MySQL instance.
swagger:model AddMySQLServiceOKBodyMysql
*/
type AddMySQLServiceOKBodyMysql struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add my SQL service OK body mysql
func (o *AddMySQLServiceOKBodyMysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLServiceOKBodyMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLServiceOKBodyMysql) UnmarshalBinary(b []byte) error {
	var res AddMySQLServiceOKBodyMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

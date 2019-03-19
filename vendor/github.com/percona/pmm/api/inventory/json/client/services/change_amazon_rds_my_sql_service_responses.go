// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeAmazonRDSMySQLServiceReader is a Reader for the ChangeAmazonRDSMySQLService structure.
type ChangeAmazonRDSMySQLServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeAmazonRDSMySQLServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeAmazonRDSMySQLServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewChangeAmazonRDSMySQLServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeAmazonRDSMySQLServiceOK creates a ChangeAmazonRDSMySQLServiceOK with default headers values
func NewChangeAmazonRDSMySQLServiceOK() *ChangeAmazonRDSMySQLServiceOK {
	return &ChangeAmazonRDSMySQLServiceOK{}
}

/*ChangeAmazonRDSMySQLServiceOK handles this case with default header values.

A successful response.
*/
type ChangeAmazonRDSMySQLServiceOK struct {
	Payload *ChangeAmazonRDSMySQLServiceOKBody
}

func (o *ChangeAmazonRDSMySQLServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/ChangeAmazonRDSMySQL][%d] changeAmazonRdsMySqlServiceOk  %+v", 200, o.Payload)
}

func (o *ChangeAmazonRDSMySQLServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeAmazonRDSMySQLServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeAmazonRDSMySQLServiceDefault creates a ChangeAmazonRDSMySQLServiceDefault with default headers values
func NewChangeAmazonRDSMySQLServiceDefault(code int) *ChangeAmazonRDSMySQLServiceDefault {
	return &ChangeAmazonRDSMySQLServiceDefault{
		_statusCode: code,
	}
}

/*ChangeAmazonRDSMySQLServiceDefault handles this case with default header values.

An error response.
*/
type ChangeAmazonRDSMySQLServiceDefault struct {
	_statusCode int

	Payload *ChangeAmazonRDSMySQLServiceDefaultBody
}

// Code gets the status code for the change amazon RDS my SQL service default response
func (o *ChangeAmazonRDSMySQLServiceDefault) Code() int {
	return o._statusCode
}

func (o *ChangeAmazonRDSMySQLServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/ChangeAmazonRDSMySQL][%d] ChangeAmazonRDSMySQLService default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeAmazonRDSMySQLServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeAmazonRDSMySQLServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeAmazonRDSMySQLServiceBody change amazon RDS my SQL service body
swagger:model ChangeAmazonRDSMySQLServiceBody
*/
type ChangeAmazonRDSMySQLServiceBody struct {

	// Instance endpoint (full DNS name). Required.
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Instance port. Required.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this change amazon RDS my SQL service body
func (o *ChangeAmazonRDSMySQLServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceBody) UnmarshalBinary(b []byte) error {
	var res ChangeAmazonRDSMySQLServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAmazonRDSMySQLServiceDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeAmazonRDSMySQLServiceDefaultBody
*/
type ChangeAmazonRDSMySQLServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change amazon RDS my SQL service default body
func (o *ChangeAmazonRDSMySQLServiceDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeAmazonRDSMySQLServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAmazonRDSMySQLServiceOKBody change amazon RDS my SQL service OK body
swagger:model ChangeAmazonRDSMySQLServiceOKBody
*/
type ChangeAmazonRDSMySQLServiceOKBody struct {

	// amazon rds mysql
	AmazonRDSMysql *ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql `json:"amazon_rds_mysql,omitempty"`
}

// Validate validates this change amazon RDS my SQL service OK body
func (o *ChangeAmazonRDSMySQLServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmazonRDSMysql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAmazonRDSMySQLServiceOKBody) validateAmazonRDSMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.AmazonRDSMysql) { // not required
		return nil
	}

	if o.AmazonRDSMysql != nil {
		if err := o.AmazonRDSMysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeAmazonRdsMySqlServiceOk" + "." + "amazon_rds_mysql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeAmazonRDSMySQLServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql AmazonRDSMySQLService represents a MySQL instance running on a single RemoteAmazonRDS Node
swagger:model ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql
*/
type ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql struct {

	// Instance endpoint (full DNS name).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Instance port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this change amazon RDS my SQL service OK body amazon RDS mysql
func (o *ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql) UnmarshalBinary(b []byte) error {
	var res ChangeAmazonRDSMySQLServiceOKBodyAmazonRDSMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package xtra_db_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListXtraDBClustersReader is a Reader for the ListXtraDBClusters structure.
type ListXtraDBClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListXtraDBClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListXtraDBClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListXtraDBClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListXtraDBClustersOK creates a ListXtraDBClustersOK with default headers values
func NewListXtraDBClustersOK() *ListXtraDBClustersOK {
	return &ListXtraDBClustersOK{}
}

/*ListXtraDBClustersOK handles this case with default header values.

A successful response.
*/
type ListXtraDBClustersOK struct {
	Payload *ListXtraDBClustersOKBody
}

func (o *ListXtraDBClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBClusters/List][%d] listXtraDbClustersOk  %+v", 200, o.Payload)
}

func (o *ListXtraDBClustersOK) GetPayload() *ListXtraDBClustersOKBody {
	return o.Payload
}

func (o *ListXtraDBClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListXtraDBClustersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListXtraDBClustersDefault creates a ListXtraDBClustersDefault with default headers values
func NewListXtraDBClustersDefault(code int) *ListXtraDBClustersDefault {
	return &ListXtraDBClustersDefault{
		_statusCode: code,
	}
}

/*ListXtraDBClustersDefault handles this case with default header values.

An unexpected error response
*/
type ListXtraDBClustersDefault struct {
	_statusCode int

	Payload *ListXtraDBClustersDefaultBody
}

// Code gets the status code for the list xtra DB clusters default response
func (o *ListXtraDBClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListXtraDBClustersDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBClusters/List][%d] ListXtraDBClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListXtraDBClustersDefault) GetPayload() *ListXtraDBClustersDefaultBody {
	return o.Payload
}

func (o *ListXtraDBClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListXtraDBClustersDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ClustersItems0 Cluster represents XtraDB cluster information.
swagger:model ClustersItems0
*/
type ClustersItems0 struct {

	// Cluster name.
	Name string `json:"name,omitempty"`

	// XtraDBClusterState represents XtraDB cluster CR state.
	//
	//  - XTRA_DB_CLUSTER_STATE_INVALID: XTRA_DB_CLUSTER_STATE_INVALID represents unknown state.
	//  - XTRA_DB_CLUSTER_STATE_CHANGING: XTRA_DB_CLUSTER_STATE_CHANGING represents a cluster being changed (initializing).
	//  - XTRA_DB_CLUSTER_STATE_READY: XTRA_DB_CLUSTER_STATE_READY represents a cluster without pending changes (ready).
	//  - XTRA_DB_CLUSTER_STATE_FAILED: XTRA_DB_CLUSTER_STATE_FAILED represents a failed cluster (error).
	//  - XTRA_DB_CLUSTER_STATE_DELETING: XTRA_DB_CLUSTER_STATE_DELETING represents a cluster being deleting.
	// Enum: [XTRA_DB_CLUSTER_STATE_INVALID XTRA_DB_CLUSTER_STATE_CHANGING XTRA_DB_CLUSTER_STATE_READY XTRA_DB_CLUSTER_STATE_FAILED XTRA_DB_CLUSTER_STATE_DELETING]
	State *string `json:"state,omitempty"`

	// operation
	Operation *ClustersItems0Operation `json:"operation,omitempty"`

	// params
	Params *ClustersItems0Params `json:"params,omitempty"`
}

// Validate validates this clusters items0
func (o *ClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clustersItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["XTRA_DB_CLUSTER_STATE_INVALID","XTRA_DB_CLUSTER_STATE_CHANGING","XTRA_DB_CLUSTER_STATE_READY","XTRA_DB_CLUSTER_STATE_FAILED","XTRA_DB_CLUSTER_STATE_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clustersItems0TypeStatePropEnum = append(clustersItems0TypeStatePropEnum, v)
	}
}

const (

	// ClustersItems0StateXTRADBCLUSTERSTATEINVALID captures enum value "XTRA_DB_CLUSTER_STATE_INVALID"
	ClustersItems0StateXTRADBCLUSTERSTATEINVALID string = "XTRA_DB_CLUSTER_STATE_INVALID"

	// ClustersItems0StateXTRADBCLUSTERSTATECHANGING captures enum value "XTRA_DB_CLUSTER_STATE_CHANGING"
	ClustersItems0StateXTRADBCLUSTERSTATECHANGING string = "XTRA_DB_CLUSTER_STATE_CHANGING"

	// ClustersItems0StateXTRADBCLUSTERSTATEREADY captures enum value "XTRA_DB_CLUSTER_STATE_READY"
	ClustersItems0StateXTRADBCLUSTERSTATEREADY string = "XTRA_DB_CLUSTER_STATE_READY"

	// ClustersItems0StateXTRADBCLUSTERSTATEFAILED captures enum value "XTRA_DB_CLUSTER_STATE_FAILED"
	ClustersItems0StateXTRADBCLUSTERSTATEFAILED string = "XTRA_DB_CLUSTER_STATE_FAILED"

	// ClustersItems0StateXTRADBCLUSTERSTATEDELETING captures enum value "XTRA_DB_CLUSTER_STATE_DELETING"
	ClustersItems0StateXTRADBCLUSTERSTATEDELETING string = "XTRA_DB_CLUSTER_STATE_DELETING"
)

// prop value enum
func (o *ClustersItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clustersItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ClustersItems0) validateState(formats strfmt.Registry) error {

	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(o.Operation) { // not required
		return nil
	}

	if o.Operation != nil {
		if err := o.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (o *ClustersItems0) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0) UnmarshalBinary(b []byte) error {
	var res ClustersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0Operation RunningOperation respresents a long-running operation.
swagger:model ClustersItems0Operation
*/
type ClustersItems0Operation struct {

	// Progress from 0.0 to 1.0; can decrease compared to the previous value.
	Progress float32 `json:"progress,omitempty"`

	// Text describing the current operation progress step.
	Message string `json:"message,omitempty"`
}

// Validate validates this clusters items0 operation
func (o *ClustersItems0Operation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0Operation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0Operation) UnmarshalBinary(b []byte) error {
	var res ClustersItems0Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0Params XtraDBClusterParams represents XtraDB cluster parameters that can be updated.
swagger:model ClustersItems0Params
*/
type ClustersItems0Params struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// proxysql
	Proxysql *ClustersItems0ParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	Pxc *ClustersItems0ParamsPxc `json:"pxc,omitempty"`
}

// Validate validates this clusters items0 params
func (o *ClustersItems0Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0Params) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *ClustersItems0Params) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(o.Pxc) { // not required
		return nil
	}

	if o.Pxc != nil {
		if err := o.Pxc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0Params) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0Params) UnmarshalBinary(b []byte) error {
	var res ClustersItems0Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0ParamsProxysql ProxySQL container parameters.
swagger:model ClustersItems0ParamsProxysql
*/
type ClustersItems0ParamsProxysql struct {

	// compute resources
	ComputeResources *ClustersItems0ParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this clusters items0 params proxysql
func (o *ClustersItems0ParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0ParamsProxysql) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0ParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0ParamsProxysql) UnmarshalBinary(b []byte) error {
	var res ClustersItems0ParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0ParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model ClustersItems0ParamsProxysqlComputeResources
*/
type ClustersItems0ParamsProxysqlComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this clusters items0 params proxysql compute resources
func (o *ClustersItems0ParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0ParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0ParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res ClustersItems0ParamsProxysqlComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0ParamsPxc PXC container parameters.
swagger:model ClustersItems0ParamsPxc
*/
type ClustersItems0ParamsPxc struct {

	// compute resources
	ComputeResources *ClustersItems0ParamsPxcComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this clusters items0 params pxc
func (o *ClustersItems0ParamsPxc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0ParamsPxc) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0ParamsPxc) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0ParamsPxc) UnmarshalBinary(b []byte) error {
	var res ClustersItems0ParamsPxc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0ParamsPxcComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model ClustersItems0ParamsPxcComputeResources
*/
type ClustersItems0ParamsPxcComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this clusters items0 params pxc compute resources
func (o *ClustersItems0ParamsPxcComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0ParamsPxcComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0ParamsPxcComputeResources) UnmarshalBinary(b []byte) error {
	var res ClustersItems0ParamsPxcComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListXtraDBClustersBody list xtra DB clusters body
swagger:model ListXtraDBClustersBody
*/
type ListXtraDBClustersBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`
}

// Validate validates this list xtra DB clusters body
func (o *ListXtraDBClustersBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListXtraDBClustersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListXtraDBClustersBody) UnmarshalBinary(b []byte) error {
	var res ListXtraDBClustersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListXtraDBClustersDefaultBody list xtra DB clusters default body
swagger:model ListXtraDBClustersDefaultBody
*/
type ListXtraDBClustersDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list xtra DB clusters default body
func (o *ListXtraDBClustersDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListXtraDBClustersDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListXtraDBClusters default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListXtraDBClustersDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListXtraDBClustersDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListXtraDBClustersDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListXtraDBClustersOKBody list xtra DB clusters OK body
swagger:model ListXtraDBClustersOKBody
*/
type ListXtraDBClustersOKBody struct {

	// XtraDB clusters information.
	Clusters []*ClustersItems0 `json:"clusters"`
}

// Validate validates this list xtra DB clusters OK body
func (o *ListXtraDBClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListXtraDBClustersOKBody) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(o.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(o.Clusters); i++ {
		if swag.IsZero(o.Clusters[i]) { // not required
			continue
		}

		if o.Clusters[i] != nil {
			if err := o.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listXtraDbClustersOk" + "." + "clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListXtraDBClustersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListXtraDBClustersOKBody) UnmarshalBinary(b []byte) error {
	var res ListXtraDBClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

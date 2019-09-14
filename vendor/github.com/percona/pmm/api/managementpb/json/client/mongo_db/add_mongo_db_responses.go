// Code generated by go-swagger; DO NOT EDIT.

package mongo_db

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMongoDBReader is a Reader for the AddMongoDB structure.
type AddMongoDBReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMongoDBOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMongoDBDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMongoDBOK creates a AddMongoDBOK with default headers values
func NewAddMongoDBOK() *AddMongoDBOK {
	return &AddMongoDBOK{}
}

/*AddMongoDBOK handles this case with default header values.

A successful response.
*/
type AddMongoDBOK struct {
	Payload *AddMongoDBOKBody
}

func (o *AddMongoDBOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/MongoDB/Add][%d] addMongoDbOk  %+v", 200, o.Payload)
}

func (o *AddMongoDBOK) GetPayload() *AddMongoDBOKBody {
	return o.Payload
}

func (o *AddMongoDBOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMongoDBDefault creates a AddMongoDBDefault with default headers values
func NewAddMongoDBDefault(code int) *AddMongoDBDefault {
	return &AddMongoDBDefault{
		_statusCode: code,
	}
}

/*AddMongoDBDefault handles this case with default header values.

An error response.
*/
type AddMongoDBDefault struct {
	_statusCode int

	Payload *AddMongoDBDefaultBody
}

// Code gets the status code for the add mongo DB default response
func (o *AddMongoDBDefault) Code() int {
	return o._statusCode
}

func (o *AddMongoDBDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/MongoDB/Add][%d] AddMongoDB default  %+v", o._statusCode, o.Payload)
}

func (o *AddMongoDBDefault) GetPayload() *AddMongoDBDefaultBody {
	return o.Payload
}

func (o *AddMongoDBDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMongoDBBody add mongo DB body
swagger:model AddMongoDBBody
*/
type AddMongoDBBody struct {

	// add node
	AddNode *AddMongoDBParamsBodyAddNode `json:"add_node,omitempty"`

	// Node and Service access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"node_id,omitempty"`

	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"node_name,omitempty"`

	// MongoDB password for exporter and QAN agent access.
	Password string `json:"password,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service Access port. Required.
	Port int64 `json:"port,omitempty"`

	// If true, adds qan-mongodb-profiler-agent for provided service.
	QANMongodbProfiler bool `json:"qan_mongodb_profiler,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MongoDB username for exporter and QAN agent access.
	Username string `json:"username,omitempty"`
}

// Validate validates this add mongo DB body
func (o *AddMongoDBBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBBody) validateAddNode(formats strfmt.Registry) error {

	if swag.IsZero(o.AddNode) { // not required
		return nil
	}

	if o.AddNode != nil {
		if err := o.AddNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "add_node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddMongoDBDefaultBody
*/
type AddMongoDBDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add mongo DB default body
func (o *AddMongoDBDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBOKBody add mongo DB OK body
swagger:model AddMongoDBOKBody
*/
type AddMongoDBOKBody struct {

	// mongodb exporter
	MongodbExporter *AddMongoDBOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`

	// qan mongodb profiler
	QANMongodbProfiler *AddMongoDBOKBodyQANMongodbProfiler `json:"qan_mongodb_profiler,omitempty"`

	// service
	Service *AddMongoDBOKBodyService `json:"service,omitempty"`
}

// Validate validates this add mongo DB OK body
func (o *AddMongoDBOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMongodbProfiler(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBOKBody) validateMongodbExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddMongoDBOKBody) validateQANMongodbProfiler(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMongodbProfiler) { // not required
		return nil
	}

	if o.QANMongodbProfiler != nil {
		if err := o.QANMongodbProfiler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbOk" + "." + "qan_mongodb_profiler")
			}
			return err
		}
	}

	return nil
}

func (o *AddMongoDBOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model AddMongoDBOKBodyMongodbExporter
*/
type AddMongoDBOKBodyMongodbExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add mongo DB OK body mongodb exporter
func (o *AddMongoDBOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbOkBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbOkBodyMongodbExporterTypeStatusPropEnum = append(addMongoDbOkBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMongoDBOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMongoDBOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMongoDBOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	AddMongoDBOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// AddMongoDBOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	AddMongoDBOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// AddMongoDBOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	AddMongoDBOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// AddMongoDBOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	AddMongoDBOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// AddMongoDBOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	AddMongoDBOKBodyMongodbExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMongoDBOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMongoDbOkBodyMongodbExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMongoDbOk"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res AddMongoDBOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBOKBodyQANMongodbProfiler QANMongoDBProfilerAgent runs within pmm-agent and sends MongoDB Query Analytics data to the PMM Server.
swagger:model AddMongoDBOKBodyQANMongodbProfiler
*/
type AddMongoDBOKBodyQANMongodbProfiler struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if query examples are disabled.
	// bool query_examples_disabled = 8; TODO https://jira.percona.com/browse/PMM-4650
	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MongoDB username for getting profiler data.
	Username string `json:"username,omitempty"`
}

// Validate validates this add mongo DB OK body QAN mongodb profiler
func (o *AddMongoDBOKBodyQANMongodbProfiler) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbOkBodyQanMongodbProfilerTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbOkBodyQanMongodbProfilerTypeStatusPropEnum = append(addMongoDbOkBodyQanMongodbProfilerTypeStatusPropEnum, v)
	}
}

const (

	// AddMongoDBOKBodyQANMongodbProfilerStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMongoDBOKBodyQANMongodbProfilerStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMongoDBOKBodyQANMongodbProfilerStatusSTARTING captures enum value "STARTING"
	AddMongoDBOKBodyQANMongodbProfilerStatusSTARTING string = "STARTING"

	// AddMongoDBOKBodyQANMongodbProfilerStatusRUNNING captures enum value "RUNNING"
	AddMongoDBOKBodyQANMongodbProfilerStatusRUNNING string = "RUNNING"

	// AddMongoDBOKBodyQANMongodbProfilerStatusWAITING captures enum value "WAITING"
	AddMongoDBOKBodyQANMongodbProfilerStatusWAITING string = "WAITING"

	// AddMongoDBOKBodyQANMongodbProfilerStatusSTOPPING captures enum value "STOPPING"
	AddMongoDBOKBodyQANMongodbProfilerStatusSTOPPING string = "STOPPING"

	// AddMongoDBOKBodyQANMongodbProfilerStatusDONE captures enum value "DONE"
	AddMongoDBOKBodyQANMongodbProfilerStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMongoDBOKBodyQANMongodbProfiler) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMongoDbOkBodyQanMongodbProfilerTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBOKBodyQANMongodbProfiler) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMongoDbOk"+"."+"qan_mongodb_profiler"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBOKBodyQANMongodbProfiler) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBOKBodyQANMongodbProfiler) UnmarshalBinary(b []byte) error {
	var res AddMongoDBOKBodyQANMongodbProfiler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBOKBodyService MongoDBService represents a generic MongoDB instance.
swagger:model AddMongoDBOKBodyService
*/
type AddMongoDBOKBodyService struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add mongo DB OK body service
func (o *AddMongoDBOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddMongoDBOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBParamsBodyAddNode AddNodeParams is a params to add new node to inventory while adding new service.
swagger:model AddMongoDBParamsBodyAddNode
*/
type AddMongoDBParamsBodyAddNode struct {

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Linux machine-id.
	// Must be unique across all Generic Nodes if specified.
	MachineID string `json:"machine_id,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_INVALID GENERIC_NODE CONTAINER_NODE REMOTE_NODE]
	NodeType *string `json:"node_type,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`
}

// Validate validates this add mongo DB params body add node
func (o *AddMongoDBParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbParamsBodyAddNodeTypeNodeTypePropEnum = append(addMongoDbParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddMongoDBParamsBodyAddNodeNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	AddMongoDBParamsBodyAddNodeNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// AddMongoDBParamsBodyAddNodeNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	AddMongoDBParamsBodyAddNodeNodeTypeGENERICNODE string = "GENERIC_NODE"

	// AddMongoDBParamsBodyAddNodeNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	AddMongoDBParamsBodyAddNodeNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// AddMongoDBParamsBodyAddNodeNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	AddMongoDBParamsBodyAddNodeNodeTypeREMOTENODE string = "REMOTE_NODE"
)

// prop value enum
func (o *AddMongoDBParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMongoDbParamsBodyAddNodeTypeNodeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"add_node"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddMongoDBParamsBodyAddNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

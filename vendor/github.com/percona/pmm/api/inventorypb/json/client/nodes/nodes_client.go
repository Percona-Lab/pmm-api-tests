// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new nodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddContainerNode adds container node adds container node
*/
func (a *Client) AddContainerNode(params *AddContainerNodeParams) (*AddContainerNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddContainerNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddContainerNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddContainer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddContainerNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddContainerNodeOK), nil

}

/*
AddGenericNode adds generic node adds generic node
*/
func (a *Client) AddGenericNode(params *AddGenericNodeParams) (*AddGenericNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddGenericNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddGenericNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddGeneric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddGenericNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddGenericNodeOK), nil

}

/*
AddRemoteAmazonRDSNode adds remote amazon RDS node adds amazon AWS RDS remote node
*/
func (a *Client) AddRemoteAmazonRDSNode(params *AddRemoteAmazonRDSNodeParams) (*AddRemoteAmazonRDSNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRemoteAmazonRDSNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddRemoteAmazonRDSNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddRemoteAmazonRDS",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRemoteAmazonRDSNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddRemoteAmazonRDSNodeOK), nil

}

/*
AddRemoteNode adds remote node adds remote node
*/
func (a *Client) AddRemoteNode(params *AddRemoteNodeParams) (*AddRemoteNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRemoteNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddRemoteNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddRemote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRemoteNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddRemoteNodeOK), nil

}

/*
GetNode gets node returns a single node by ID
*/
func (a *Client) GetNode(params *GetNodeParams) (*GetNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodeOK), nil

}

/*
ListNodes lists nodes returns a list of all nodes
*/
func (a *Client) ListNodes(params *ListNodesParams) (*ListNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListNodes",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodesOK), nil

}

/*
RemoveNode removes node removes node without any agents and services
*/
func (a *Client) RemoveNode(params *RemoveNodeParams) (*RemoveNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/Remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveNodeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
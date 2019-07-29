package inventory

import (
	"context"
	"testing"

	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestNodes(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		t.Parallel()

		remoteNode := addRemoteNode(t, pmmapitests.TestString(t, "Test Remote Node for List"))
		remoteNodeID := remoteNode.Remote.NodeID
		defer pmmapitests.RemoveNodes(t, remoteNodeID)
		genericNodeID := addGenericNode(t, pmmapitests.TestString(t, "Test Generic Node for List")).NodeID
		require.NotEmpty(t, genericNodeID)
		defer pmmapitests.RemoveNodes(t, genericNodeID)

		res, err := client.Default.Nodes.ListNodes(nil)
		require.NoError(t, err)
		require.NotZerof(t, len(res.Payload.Generic), "There should be at least one node")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.Generic {
				if v.NodeID == genericNodeID {
					return true
				}
			}
			return false
		}, "There should be generic node with id `%s`", genericNodeID)
		require.NotZerof(t, len(res.Payload.Remote), "There should be at least one node")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.Remote {
				if v.NodeID == remoteNodeID {
					return true
				}
			}
			return false
		}, "There should be remote node with id `%s`", remoteNodeID)
	})
}

func TestGetNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "TestGenericNode")
		nodeID := addGenericNode(t, nodeName).NodeID
		require.NotEmpty(t, nodeID)
		defer pmmapitests.RemoveNodes(t, nodeID)

		expectedResponse := nodes.GetNodeOK{
			Payload: &nodes.GetNodeOKBody{
				Generic: &nodes.GetNodeOKBodyGeneric{
					NodeID:   nodeID,
					NodeName: nodeName,
					Address:  "10.10.10.10",
				},
			},
		}

		params := &nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.GetNode(params)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse.Payload, res.Payload)
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		params := &nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: "pmm-not-found"},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.GetNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Node with ID \"pmm-not-found\" not found.")
		assert.Nil(t, res)
	})

	t.Run("EmptyNodeID", func(t *testing.T) {
		t.Parallel()

		params := &nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.GetNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeId: value '' must not be an empty string")
		assert.Nil(t, res)
	})
}

func TestGenericNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Test Generic Node")
		params := &nodes.AddGenericNodeParams{
			Body: nodes.AddGenericNodeBody{
				NodeName: nodeName,
				Address:  "10.10.10.10",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddGenericNode(params)
		assert.NoError(t, err)
		require.NotNil(t, res)
		require.NotNil(t, res.Payload.Generic)
		nodeID := res.Payload.Generic.NodeID
		defer pmmapitests.RemoveNodes(t, nodeID)

		// Check node exists in DB.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		expectedResponse := &nodes.GetNodeOK{
			Payload: &nodes.GetNodeOKBody{
				Generic: &nodes.GetNodeOKBodyGeneric{
					NodeID:   res.Payload.Generic.NodeID,
					NodeName: nodeName,
					Address:  "10.10.10.10",
				},
			},
		}
		require.Equal(t, expectedResponse, getNodeRes)

		// Check duplicates.
		res, err = client.Default.Nodes.AddGenericNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 409, codes.AlreadyExists, "Node with name %q already exists.", nodeName)
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.Generic.NodeID)
		}
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		t.Parallel()

		params := &nodes.AddGenericNodeParams{
			Body:    nodes.AddGenericNodeBody{NodeName: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddGenericNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeName: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.Generic.NodeID)
		}
	})
}

func TestContainerNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Test Container Node")
		params := &nodes.AddContainerNodeParams{
			Body: nodes.AddContainerNodeBody{
				NodeName:      nodeName,
				ContainerID:   "docker-id",
				ContainerName: "docker-name",
				MachineID:     "machine-id",
				Address:       "10.10.1.10",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddContainerNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.Container)
		nodeID := res.Payload.Container.NodeID
		defer pmmapitests.RemoveNodes(t, nodeID)

		// Check node exists in DB.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		expectedResponse := &nodes.GetNodeOK{
			Payload: &nodes.GetNodeOKBody{
				Container: &nodes.GetNodeOKBodyContainer{
					NodeID:        res.Payload.Container.NodeID,
					NodeName:      nodeName,
					ContainerID:   "docker-id",
					ContainerName: "docker-name",
					MachineID:     "machine-id",
					Address:       "10.10.1.10",
				},
			},
		}
		require.Equal(t, expectedResponse, getNodeRes)

		// Check duplicates.
		res, err = client.Default.Nodes.AddContainerNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 409, codes.AlreadyExists, "Node with name %q already exists.", nodeName)
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.Container.NodeID)
		}
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		t.Parallel()

		params := &nodes.AddContainerNodeParams{
			Body:    nodes.AddContainerNodeBody{NodeName: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddContainerNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeName: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.Container.NodeID)
		}
	})
}

func TestRemoteNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Test Remote Node")
		params := &nodes.AddRemoteNodeParams{
			Body: nodes.AddRemoteNodeBody{
				NodeName: nodeName,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.Remote)
		nodeID := res.Payload.Remote.NodeID
		defer pmmapitests.RemoveNodes(t, nodeID)

		// Check node exists in DB.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		expectedResponse := &nodes.GetNodeOK{
			Payload: &nodes.GetNodeOKBody{
				Remote: &nodes.GetNodeOKBodyRemote{
					NodeID:   res.Payload.Remote.NodeID,
					NodeName: nodeName,
				},
			},
		}
		require.Equal(t, expectedResponse, getNodeRes)

		// Check duplicates.
		res, err = client.Default.Nodes.AddRemoteNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 409, codes.AlreadyExists, "Node with name %q already exists.", nodeName)
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.Remote.NodeID)
		}
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		t.Parallel()

		params := &nodes.AddRemoteNodeParams{
			Body:    nodes.AddRemoteNodeBody{NodeName: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeName: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.Remote.NodeID)
		}
	})
}

func TestRemoteAmazonRDSNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Test RemoteAmazonRDS Node")
		instanceName := pmmapitests.TestString(t, "some-instance")
		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: nodeName,
				Instance: instanceName,
				Region:   "us-east-1",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.RemoteAmazonRDS)
		nodeID := res.Payload.RemoteAmazonRDS.NodeID
		defer pmmapitests.RemoveNodes(t, nodeID)

		// Check if the node saved in PMM-Managed.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		expectedResponse := &nodes.GetNodeOK{
			Payload: &nodes.GetNodeOKBody{
				RemoteAmazonRDS: &nodes.GetNodeOKBodyRemoteAmazonRDS{
					NodeID:   nodeID,
					NodeName: nodeName,
					Region:   "us-east-1",
					Instance: instanceName,
				},
			},
		}
		assert.Equal(t, expectedResponse, getNodeRes)

		// Check duplicates.
		res, err = client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 409, codes.AlreadyExists, "Node with name %q already exists.", nodeName)
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.RemoteAmazonRDS.NodeID)
		}
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: "",
				Instance: "some-instance-without-name",
				Region:   "us-east-1",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeName: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.RemoteAmazonRDS.NodeID)
		}
	})

	t.Run("AddInstanceEmpty", func(t *testing.T) {
		t.Parallel()

		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: pmmapitests.TestString(t, "Remote AmazonRDSNode without instance"),
				Region:   "us-west-1",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field Instance: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.RemoteAmazonRDS.NodeID)
		}
	})

	t.Run("AddRegionEmpty", func(t *testing.T) {
		t.Parallel()

		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: pmmapitests.TestString(t, "Remote AmazonRDSNode without instance"),
				Instance: "instance-without-region",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field Region: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.RemoteAmazonRDS.NodeID)
		}
	})
}

func TestRemoveNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Generic Node for basic remove test")
		node := addGenericNode(t, nodeName)
		nodeID := node.NodeID

		removeResp, err := client.Default.Nodes.RemoveNode(&nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: nodeID,
			},
			Context: context.Background(),
		})
		assert.NoError(t, err)
		assert.NotNil(t, removeResp)
	})

	t.Run("With service", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Generic Node for remove test")
		node := addGenericNode(t, nodeName)

		serviceName := pmmapitests.TestString(t, "MySQL Service for agent")
		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      node.NodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: serviceName,
		})
		serviceID := service.Mysql.ServiceID

		removeResp, err := client.Default.Nodes.RemoveNode(&nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: node.NodeID,
			},
			Context: context.Background(),
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `Node with ID %q has services.`, node.NodeID)
		assert.Nil(t, removeResp)

		// Check that node and service isn't removed.
		getServiceResp, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: node.NodeID},
			Context: pmmapitests.Context,
		})
		assert.NotNil(t, getServiceResp)
		assert.NoError(t, err)

		listAgentsOK, err := client.Default.Services.ListServices(&services.ListServicesParams{
			Body: services.ListServicesBody{
				NodeID: node.NodeID,
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &services.ListServicesOKBody{
			Mysql: []*services.MysqlItems0{
				{
					NodeID:      node.NodeID,
					ServiceID:   serviceID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, listAgentsOK.Payload)

		// Remove with force flag.
		params := &nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: node.NodeID,
				Force:  true,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.RemoveNode(params)
		assert.NoError(t, err)
		assert.NotNil(t, res)

		// Check that the node and agents are removed.
		getServiceResp, err = client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: node.NodeID},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Node with ID %q not found.", node.NodeID)
		assert.Nil(t, getServiceResp)

		listAgentsOK, err = client.Default.Services.ListServices(&services.ListServicesParams{
			Body: services.ListServicesBody{
				NodeID: node.NodeID,
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &services.ListServicesOKBody{}, listAgentsOK.Payload)
	})

	t.Run("With pmm-agent", func(t *testing.T) {
		t.Parallel()

		nodeName := pmmapitests.TestString(t, "Generic Node for remove test")
		node := addGenericNode(t, nodeName)

		_ = addPMMAgent(t, node.NodeID)

		removeResp, err := client.Default.Nodes.RemoveNode(&nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: node.NodeID,
			},
			Context: context.Background(),
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `Node with ID %q has pmm-agent.`, node.NodeID)
		assert.Nil(t, removeResp)

		// Remove with force flag.
		params := &nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: node.NodeID,
				Force:  true,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.RemoveNode(params)
		assert.NoError(t, err)
		assert.NotNil(t, res)

		// Check that the node and agents are removed.
		getServiceResp, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: node.NodeID},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Node with ID %q not found.", node.NodeID)
		assert.Nil(t, getServiceResp)

		listAgentsOK, err := client.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Body: agents.ListAgentsBody{
				NodeID: node.NodeID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Node with ID %q not found.", node.NodeID)
		assert.Nil(t, listAgentsOK)
	})

	t.Run("Not-exist node", func(t *testing.T) {
		t.Parallel()
		nodeID := "not-exist-node-id"
		removeResp, err := client.Default.Nodes.RemoveNode(&nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: nodeID,
			},
			Context: context.Background(),
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, `Node with ID %q not found.`, nodeID)
		assert.Nil(t, removeResp)
	})

	t.Run("Empty params", func(t *testing.T) {
		t.Parallel()
		removeResp, err := client.Default.Nodes.RemoveNode(&nodes.RemoveNodeParams{
			Body:    nodes.RemoveNodeBody{},
			Context: context.Background(),
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeId: value '' must not be an empty string")
		assert.Nil(t, removeResp)
	})
}

package inventory

import (
	"testing"

	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/nodes"
	"github.com/stretchr/testify/require"

	"github.com/Percona-Lab/pmm-api-tests"
)

func TestNodes(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		remoteNode := addRemoteNode(t, withUUID(t, "Test Remote Node for List"))
		remoteNodeID := remoteNode.Remote.NodeID
		defer removeNodes(t, remoteNodeID)
		genericNode := addGenericNode(t, withUUID(t, "Test Remote Node for List"))
		genericNodeID := genericNode.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		res, err := client.Default.Nodes.ListNodes(nil)
		require.NoError(t, err)
		require.NotZerof(t, len(res.Payload.Generic), "There should be at least one node")
		require.Condition(t, func() (success bool) {
			for _, v := range res.Payload.Generic {
				if v.NodeID == genericNodeID {
					return true
				}
			}
			return false
		}, "There should be generic node with id `pmm-server`")
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
		params := &nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: "pmm-server"},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.GetNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.Generic)
		require.Equal(t, res.Payload.Generic.NodeID, "pmm-server")
		require.Equal(t, res.Payload.Generic.NodeName, "PMM Server")
		require.Nil(t, res.Payload.Container)
		require.Nil(t, res.Payload.Remote)
		require.Nil(t, res.Payload.RemoteAmazonRDS)
	})

	t.Run("NotFound", func(t *testing.T) {
		params := &nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: "pmm-not-found"},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.GetNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 404)")
		require.Nil(t, res)
	})

	t.Run("EmptyNodeID", func(t *testing.T) {
		params := &nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.GetNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})
}

func TestGenericNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		params := &nodes.AddGenericNodeParams{
			Body:    nodes.AddGenericNodeBody{NodeName: withUUID(t, "Test Generic Node")},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddGenericNode(params)
		nodeID := res.Payload.Generic.NodeID
		defer removeNodes(t, nodeID)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.Generic)

		// Check node exists in DB.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getNodeRes.Payload.Generic)
		require.Equal(t, nodeID, getNodeRes.Payload.Generic.NodeID)
		require.Equal(t, params.Body.NodeName, getNodeRes.Payload.Generic.NodeName)

		// Check duplicates.
		res, err = client.Default.Nodes.AddGenericNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 409)")
		require.Nil(t, res)

		// Change node.
		changedNodeName := withUUID(t, "Changed Generic Node")
		changeRes, err := client.Default.Nodes.ChangeGenericNode(&nodes.ChangeGenericNodeParams{
			Body:    nodes.ChangeGenericNodeBody{NodeID: nodeID, NodeName: changedNodeName},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Generic)
		require.Equal(t, nodeID, changeRes.Payload.Generic.NodeID)
		require.Equal(t, changedNodeName, changeRes.Payload.Generic.NodeName)
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		params := &nodes.AddGenericNodeParams{
			Body:    nodes.AddGenericNodeBody{NodeName: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddGenericNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})
}

func TestContainerNode(t *testing.T) {
	t.Skip("Haven't implemented yet.")
	t.Run("Basic", func(t *testing.T) {
		params := &nodes.AddContainerNodeParams{
			Body: nodes.AddContainerNodeBody{
				NodeName:            withUUID(t, "Test Container Node"),
				DockerContainerID:   "docker-id",
				DockerContainerName: "docker-name",
				MachineID:           "machine-id",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddContainerNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.Container)
		defer removeNodes(t, res.Payload.Container.NodeID)

		// Check node exists in DB.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: res.Payload.Container.NodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getNodeRes.Payload.Container)
		require.Equal(t, getNodeRes.Payload.Container.NodeID, res.Payload.Container.NodeID)
		require.Equal(t, getNodeRes.Payload.Container.NodeName, params.Body.NodeName)

		// Check duplicates.
		res, err = client.Default.Nodes.AddContainerNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 409)")
		require.Nil(t, res)

		// Change node.
		changedNodeName := withUUID(t, "Changed Container Node")
		changeRes, err := client.Default.Nodes.ChangeContainerNode(&nodes.ChangeContainerNodeParams{
			Body:    nodes.ChangeContainerNodeBody{NodeID: res.Payload.Container.NodeID, NodeName: changedNodeName},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Container)
		require.Equal(t, getNodeRes.Payload.Container.NodeID, res.Payload.Container.NodeID)
		require.Equal(t, getNodeRes.Payload.Container.NodeName, changedNodeName)
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		params := &nodes.AddContainerNodeParams{
			Body:    nodes.AddContainerNodeBody{NodeName: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddContainerNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})
}

func TestRemoteNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		params := &nodes.AddRemoteNodeParams{
			Body: nodes.AddRemoteNodeBody{
				NodeName: withUUID(t, "Test Remote Node"),
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.Remote)
		nodeID := res.Payload.Remote.NodeID
		defer removeNodes(t, nodeID)

		// Check node exists in DB.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getNodeRes.Payload.Remote)
		require.Equal(t, getNodeRes.Payload.Remote.NodeID, nodeID)
		require.Equal(t, getNodeRes.Payload.Remote.NodeName, params.Body.NodeName)

		// Check duplicates.
		res, err = client.Default.Nodes.AddRemoteNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 409)")
		require.Nil(t, res)

		// Change node.
		changedNodeName := withUUID(t, "Changed Remote Node")
		changeRes, err := client.Default.Nodes.ChangeRemoteNode(&nodes.ChangeRemoteNodeParams{
			Body:    nodes.ChangeRemoteNodeBody{NodeID: nodeID, NodeName: changedNodeName},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Remote)
		require.Equal(t, changeRes.Payload.Remote.NodeID, nodeID)
		require.Equal(t, changeRes.Payload.Remote.NodeName, changedNodeName)
	})

	t.Run("AddNameEmpty", func(t *testing.T) {
		params := &nodes.AddRemoteNodeParams{
			Body:    nodes.AddRemoteNodeBody{NodeName: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})
}

func TestRemoteAmazonRDSNode(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: withUUID(t, "Test RemoteAmazonRDS Node"),
				Instance: withUUID(t, "some-instance"),
				Region:   "us-east-1",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		require.NoError(t, err)
		require.NotNil(t, res.Payload.RemoteAmazonRDS)
		nodeID := res.Payload.RemoteAmazonRDS.NodeID
		defer removeNodes(t, nodeID)

		// Check if the node saved in PMM-Managed.
		getNodeRes, err := client.Default.Nodes.GetNode(&nodes.GetNodeParams{
			Body:    nodes.GetNodeBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getNodeRes.Payload.RemoteAmazonRDS)
		require.Equal(t, getNodeRes.Payload.RemoteAmazonRDS.NodeID, nodeID)
		require.Equal(t, getNodeRes.Payload.RemoteAmazonRDS.NodeName, params.Body.NodeName)

		// Check duplicates.
		res, err = client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 409)")
		require.Nil(t, res)

		// Change node.
		changedNodeName := withUUID(t, "Changed RemoteAmazonRDS Node")
		changeRes, err := client.Default.Nodes.ChangeRemoteAmazonRDSNode(&nodes.ChangeRemoteAmazonRDSNodeParams{
			Body:    nodes.ChangeRemoteAmazonRDSNodeBody{NodeID: nodeID, NodeName: changedNodeName},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.RemoteAmazonRDS)
		require.Equal(t, changeRes.Payload.RemoteAmazonRDS.NodeID, nodeID)
		require.Equal(t, changeRes.Payload.RemoteAmazonRDS.NodeName, changedNodeName)
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
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})

	t.Run("AddInstanceEmpty", func(t *testing.T) {
		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: withUUID(t, "Remote AmazonRDSNode without instance"),
				Region:   "us-west-1",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})

	t.Run("AddRegionEmpty", func(t *testing.T) {
		params := &nodes.AddRemoteAmazonRDSNodeParams{
			Body: nodes.AddRemoteAmazonRDSNodeBody{
				NodeName: withUUID(t, "Remote AmazonRDSNode without instance"),
				Instance: "instance-without-region",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Nodes.AddRemoteAmazonRDSNode(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})
}

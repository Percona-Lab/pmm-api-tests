package inventory

import (
	"context"
	"testing"

	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/agents"
	"github.com/percona/pmm/api/inventory/json/client/nodes"
	"github.com/percona/pmm/api/inventory/json/client/services"
	"github.com/stretchr/testify/require"

	_ "github.com/Percona-Lab/pmm-api-tests" // init default client
)

func removeNodes(t *testing.T, nodeIDs ...string) {
	t.Helper()
	for _, nodeID := range nodeIDs {
		params := &nodes.RemoveNodeParams{
			Body:    nodes.RemoveNodeBody{NodeID: nodeID},
			Context: context.TODO(),
		}
		res, err := client.Default.Nodes.RemoveNode(params)
		require.NoError(t, err)
		require.NotNil(t, res)
	}
}

func addRemoteNode(t *testing.T, nodeName string) string {
	t.Helper()
	params := &nodes.AddRemoteNodeParams{
		Body: nodes.AddRemoteNodeBody{
			NodeName: nodeName,
		},
		Context: context.TODO(),
	}
	res, err := client.Default.Nodes.AddRemoteNode(params)
	require.NoError(t, err)
	require.NotNil(t, res.Payload.Remote)
	nodeID := res.Payload.Remote.NodeID
	return nodeID
}

func removeServices(t *testing.T, serviceIDs ...string) {
	t.Helper()
	for _, serviceID := range serviceIDs {
		params := &services.RemoveServiceParams{
			Body:    services.RemoveServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		}
		res, err := client.Default.Services.RemoveService(params)
		require.NoError(t, err)
		require.NotNil(t, res)
	}
}

func addMySQLService(t *testing.T, body services.AddMySQLServiceBody) string {
	t.Helper()
	params := &services.AddMySQLServiceParams{
		Body:    body,
		Context: context.TODO(),
	}
	res, err := client.Default.Services.AddMySQLService(params)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.Payload.Mysql)
	require.NotEmpty(t, res.Payload.Mysql.ServiceID)
	return res.Payload.Mysql.ServiceID
}

func removeAgents(t *testing.T, agentIDs ...string) {
	t.Helper()
	for _, agentID := range agentIDs {
		params := &agents.RemoveAgentParams{
			Body:    agents.RemoveAgentBody{AgentID: agentID},
			Context: context.TODO(),
		}
		res, err := client.Default.Agents.RemoveAgent(params)
		require.NoError(t, err)
		require.NotNil(t, res)
	}
}

func addPMMAgent(t *testing.T, node string) (*agents.AddPMMAgentOK, error) {
	t.Helper()
	res, err := client.Default.Agents.AddPMMAgent(&agents.AddPMMAgentParams{
		Body:    agents.AddPMMAgentBody{NodeID: node},
		Context: context.TODO(),
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.Payload)
	require.NotNil(t, res.Payload.PMMAgent)
	require.NotNil(t, res.Payload.PMMAgent.AgentID)
	return res, err
}

func addMySqldExporter(t *testing.T, body agents.AddMySqldExporterBody) *agents.AddMySqldExporterOKBody {
	t.Helper()
	agentRes, err := client.Default.Agents.AddMySqldExporter(&agents.AddMySqldExporterParams{
		Body:    body,
		Context: context.TODO(),
	})
	require.NoError(t, err)
	require.NotNil(t, agentRes)
	require.NotNil(t, agentRes.Payload.MysqldExporter)
	return agentRes.Payload
}

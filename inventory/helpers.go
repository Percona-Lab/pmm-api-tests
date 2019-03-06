package inventory

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/google/uuid"
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/agents"
	"github.com/percona/pmm/api/inventory/json/client/nodes"
	"github.com/percona/pmm/api/inventory/json/client/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/Percona-Lab/pmm-api-tests" // init default client
)

func withUUID(t *testing.T, name string) string {
	hostname, err := os.Hostname()
	require.NoError(t, err)
	random, err := uuid.NewRandom()
	require.NoError(t, err)

	return fmt.Sprintf("test-for-%s-%s-%s", hostname, name, random.String())
}

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

func addGenericNode(t *testing.T, nodeName string) *nodes.AddGenericNodeOKBody {
	t.Helper()
	params := &nodes.AddGenericNodeParams{
		Body: nodes.AddGenericNodeBody{
			NodeName: nodeName,
		},
		Context: context.TODO(),
	}
	res, err := client.Default.Nodes.AddGenericNode(params)
	require.NoError(t, err)
	require.NotNil(t, res.Payload.Generic)
	return res.Payload
}

func addRemoteNode(t *testing.T, nodeName string) *nodes.AddRemoteNodeOKBody {
	t.Helper()
	params := &nodes.AddRemoteNodeParams{
		Body: nodes.AddRemoteNodeBody{
			NodeName: nodeName,
		},
		Context: context.TODO(),
	}
	res, err := client.Default.Nodes.AddRemoteNode(params)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.Payload)
	require.NotNil(t, res.Payload.Remote)
	return res.Payload
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

func addMySQLService(t *testing.T, body services.AddMySQLServiceBody) *services.AddMySQLServiceOKBody {
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
	return res.Payload
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

func addPMMAgent(t *testing.T, node string) *agents.AddPMMAgentOKBody {
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
	return res.Payload
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

func assertEqualAPIError(t *testing.T, err error, expectedCode int) {
	t.Helper()
	expectedError := &runtime.APIError{
		OperationName: "unknown error",
		Code:          expectedCode,
	}
	assert.Error(t, err)
	err.(*runtime.APIError).Response = nil
	assert.Equal(t, expectedError, err)
}

package inventory

import (
	"testing"

	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/agents"
	"github.com/percona/pmm/api/inventory/json/client/services"
	"github.com/stretchr/testify/require"

	"github.com/Percona-Lab/pmm-api-tests"
)

func TestAgents(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		t.Parallel()

		node := addRemoteNode(t, withUUID(t, "Remote node for agents list"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: "MySQL Service for agent",
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		mySqldExporter := addMySqldExporter(t, agents.AddMySqldExporterBody{
			ServiceID:    serviceID,
			Username:     "username",
			Password:     "password",
			RunsOnNodeID: "pmm-server",
		})
		mySqldExporterID := mySqldExporter.MysqldExporter.AgentID
		defer removeAgents(t, mySqldExporterID)

		pmmAgent := addPMMAgent(t, nodeID)
		pmmAgentID := pmmAgent.PMMAgent.AgentID
		defer removeAgents(t, pmmAgentID)

		res, err := client.Default.Agents.ListAgents(&agents.ListAgentsParams{Context: pmmapitests.Context})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotZerof(t, len(res.Payload.MysqldExporter), "There should be at least one service")

		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.MysqldExporter {
				if v.AgentID == mySqldExporterID {
					return true
				}
			}
			return false
		}, "There should be MySQL agent with id `%s`", mySqldExporterID)
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.PMMAgent {
				if v.AgentID == pmmAgentID {
					return true
				}
			}
			return false
		}, "There should be PMM-agent with id `%s`", pmmAgentID)
	})

	t.Run("FilterList", func(t *testing.T) {
		t.Parallel()

		node := addRemoteNode(t, withUUID(t, "Remote node for agents filters"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "MySQL Service for filter test"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		mySqldExporter := addMySqldExporter(t, agents.AddMySqldExporterBody{
			ServiceID:    serviceID,
			Username:     "username",
			Password:     "password",
			RunsOnNodeID: "pmm-server",
		})
		mySqldExporterID := mySqldExporter.MysqldExporter.AgentID
		defer removeAgents(t, mySqldExporterID)

		pmmAgent := addPMMAgent(t, nodeID)
		pmmAgentID := pmmAgent.PMMAgent.AgentID
		defer removeAgents(t, pmmAgentID)

		// Filter by runs on node ID.
		res, err := client.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Body:    agents.ListAgentsBody{RunsOnNodeID: "pmm-server"},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotZerof(t, len(res.Payload.MysqldExporter), "There should be at least one service")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.MysqldExporter {
				if v.AgentID == mySqldExporterID {
					return true
				}
			}
			return false
		}, "There should be MySQL agent with id `%s`", mySqldExporterID)
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.PMMAgent {
				if v.AgentID == pmmAgentID {
					return false
				}
			}
			return true
		}, "There should not be PMM-agent with id `%s`", pmmAgentID)

		// Filter by node ID.
		res, err = client.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Body:    agents.ListAgentsBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotZerof(t, len(res.Payload.PMMAgent), "There should be at least one service")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.MysqldExporter {
				if v.AgentID == mySqldExporterID {
					return false
				}
			}
			return true
		}, "There should not be MySQL agent with id `%s`", mySqldExporterID)
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.PMMAgent {
				if v.AgentID == pmmAgentID {
					return true
				}
			}
			return false
		}, "There should be PMM-agent with id `%s`", pmmAgentID)

		// Filter by service ID.
		res, err = client.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Body:    agents.ListAgentsBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotZerof(t, len(res.Payload.MysqldExporter), "There should be at least one service")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.MysqldExporter {
				if v.AgentID == mySqldExporterID {
					return true
				}
			}
			return false
		}, "There should be MySQL agent with id `%s`", mySqldExporterID)
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.PMMAgent {
				if v.AgentID == pmmAgentID {
					return false
				}
			}
			return true
		}, "There should not be PMM-agent with id `%s`", pmmAgentID)
	})

	t.Run("TwoOrMoreFilters", func(t *testing.T) {
		t.Skip("it doesn't return error")
		t.Parallel()

		res, err := client.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Body: agents.ListAgentsBody{
				RunsOnNodeID: "pmm-server",
				NodeID:       "pmm-server",
				ServiceID:    "some-service-id",
			},
			Context: pmmapitests.Context,
		})
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestPMMAgent(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		node := addRemoteNode(t, withUUID(t, "Remote node for PMM-agent"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		res := addPMMAgent(t, nodeID)
		require.Equal(t, nodeID, res.PMMAgent.NodeID)
		agentID := res.PMMAgent.AgentID
		defer removeAgents(t, agentID)

		getAgentRes, err := client.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body:    agents.GetAgentBody{AgentID: agentID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getAgentRes)
		require.NotNil(t, getAgentRes.Payload)
		require.NotNil(t, getAgentRes.Payload.PMMAgent)
		require.Equal(t, agentID, getAgentRes.Payload.PMMAgent.AgentID)
		require.Equal(t, nodeID, getAgentRes.Payload.PMMAgent.NodeID)
	})

	t.Run("AddNodeIDEmpty", func(t *testing.T) {
		t.Parallel()

		res, err := client.Default.Agents.AddPMMAgent(&agents.AddPMMAgentParams{
			Body:    agents.AddPMMAgentBody{NodeID: ""},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})
}

func TestNodeExporter(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		node := addRemoteNode(t, withUUID(t, "Remote node for Node exporter"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		res, err := client.Default.Agents.AddNodeExporter(&agents.AddNodeExporterParams{
			Body: agents.AddNodeExporterBody{
				NodeID: nodeID,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotNil(t, res.Payload)
		require.NotNil(t, res.Payload.NodeExporter)
		require.NotNil(t, res.Payload.NodeExporter.AgentID)
		require.Equal(t, nodeID, res.Payload.NodeExporter.NodeID)
		agentID := res.Payload.NodeExporter.AgentID
		defer removeAgents(t, agentID)

		getAgentRes, err := client.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body:    agents.GetAgentBody{AgentID: agentID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getAgentRes)
		require.NotNil(t, getAgentRes.Payload)
		require.NotNil(t, getAgentRes.Payload.NodeExporter)
		require.Equal(t, agentID, getAgentRes.Payload.NodeExporter.AgentID)
		require.Equal(t, nodeID, getAgentRes.Payload.NodeExporter.NodeID)
	})

	t.Run("AddNodeIDEmpty", func(t *testing.T) {
		t.Parallel()

		res, err := client.Default.Agents.AddNodeExporter(&agents.AddNodeExporterParams{
			Body:    agents.AddNodeExporterBody{NodeID: ""},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})

	t.Run("NotExistNodeID", func(t *testing.T) {
		t.Parallel()

		res, err := client.Default.Agents.AddNodeExporter(&agents.AddNodeExporterParams{
			Body:    agents.AddNodeExporterBody{NodeID: "pmm-node-exporter-node"},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 404)")
		require.Nil(t, res)
	})
}

func TestMySQLdExporter(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		node := addRemoteNode(t, withUUID(t, "Remote node for Node exporter"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "MySQL Service for MySQLdExporter test"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		mySqldExporter := addMySqldExporter(t, agents.AddMySqldExporterBody{
			ServiceID:    serviceID,
			Username:     "username",
			Password:     "password",
			RunsOnNodeID: nodeID,
		})
		agentID := mySqldExporter.MysqldExporter.AgentID
		defer removeAgents(t, agentID)

		getAgentRes, err := client.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body:    agents.GetAgentBody{AgentID: agentID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getAgentRes)
		require.NotNil(t, getAgentRes.Payload)
		require.NotNil(t, getAgentRes.Payload.MysqldExporter)
		require.Equal(t, agentID, getAgentRes.Payload.MysqldExporter.AgentID)
		require.Equal(t, serviceID, getAgentRes.Payload.MysqldExporter.ServiceID)
		require.Equal(t, nodeID, getAgentRes.Payload.MysqldExporter.RunsOnNodeID)
	})

	t.Run("AddServiceIDEmpty", func(t *testing.T) {
		t.Parallel()

		res, err := client.Default.Agents.AddMySqldExporter(&agents.AddMySqldExporterParams{
			Body: agents.AddMySqldExporterBody{
				ServiceID:    "",
				RunsOnNodeID: "pmm-server",
			},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})

	t.Run("AddServiceIDEmpty", func(t *testing.T) {
		t.Parallel()

		res, err := client.Default.Agents.AddMySqldExporter(&agents.AddMySqldExporterParams{
			Body: agents.AddMySqldExporterBody{
				ServiceID:    "pmm-service-id",
				RunsOnNodeID: "",
			},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 400)")
		require.Nil(t, res)
	})

	t.Run("NotExistServiceID", func(t *testing.T) {
		t.Parallel()

		res, err := client.Default.Agents.AddMySqldExporter(&agents.AddMySqldExporterParams{
			Body: agents.AddMySqldExporterBody{
				ServiceID:    "pmm-service-id",
				RunsOnNodeID: "pmm-server",
			},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 404)")
		require.Nil(t, res)
	})

	t.Run("NotExistNodeID", func(t *testing.T) {
		t.Parallel()

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "MySQL Service for not exists node ID"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		res, err := client.Default.Agents.AddMySqldExporter(&agents.AddMySqldExporterParams{
			Body: agents.AddMySqldExporterBody{
				ServiceID:    serviceID,
				RunsOnNodeID: "pmm-not-exist-server",
			},
			Context: pmmapitests.Context,
		})
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 404)")
		require.Nil(t, res)
	})
}

func TestRDSExporter(t *testing.T) {
	t.Skip("Not implemented yet.")
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		node := addRemoteNode(t, withUUID(t, "Remote node for Node exporter"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "MySQL Service for RDSExporter test"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		res, err := client.Default.Agents.AddRDSExporter(&agents.AddRDSExporterParams{
			Body: agents.AddRDSExporterBody{
				RunsOnNodeID: nodeID,
				ServiceIds:   []string{serviceID},
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotNil(t, res.Payload.RDSExporter)
		agentID := res.Payload.RDSExporter.AgentID
		defer removeAgents(t, agentID)

		getAgentRes, err := client.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body:    agents.GetAgentBody{AgentID: agentID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, getAgentRes)
		require.NotNil(t, getAgentRes.Payload)
		require.NotNil(t, getAgentRes.Payload.RDSExporter)
		require.Equal(t, agentID, getAgentRes.Payload.RDSExporter.AgentID)
		require.Contains(t, getAgentRes.Payload.RDSExporter.ServiceIds, serviceID)
		require.Equal(t, nodeID, getAgentRes.Payload.RDSExporter.RunsOnNodeID)
	})
}

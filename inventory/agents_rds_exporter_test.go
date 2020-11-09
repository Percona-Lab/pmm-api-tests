package inventory

import (
	"testing"

	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestRDSExporter(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		genericNodeID := pmmapitests.AddGenericNode(t, pmmapitests.TestString(t, "")).NodeID
		require.NotEmpty(t, genericNodeID)
		defer pmmapitests.RemoveNodes(t, genericNodeID)

		node := addRemoteRDSNode(t, pmmapitests.TestString(t, "Remote node for RDS exporter"))
		nodeID := node.RemoteRDS.NodeID
		defer pmmapitests.RemoveNodes(t, nodeID)

		pmmAgent := pmmapitests.AddPMMAgent(t, genericNodeID)
		pmmAgentID := pmmAgent.PMMAgent.AgentID
		defer pmmapitests.RemoveAgents(t, pmmAgentID)

		rdsExporter := addRDSExporter(t, agents.AddRDSExporterBody{
			NodeID:     nodeID,
			PMMAgentID: pmmAgentID,
			CustomLabels: map[string]string{
				"custom_label_rds_exporter": "rds_exporter",
			},
			SkipConnectionCheck:    true,
			DisableBasicMetrics:    true,
			DisableEnhancedMetrics: true,
		})
		agentID := rdsExporter.RDSExporter.AgentID
		defer pmmapitests.RemoveAgents(t, agentID)

		getAgentRes, err := client.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body:    agents.GetAgentBody{AgentID: agentID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		assert.Equal(t, &agents.GetAgentOK{
			Payload: &agents.GetAgentOKBody{
				RDSExporter: &agents.GetAgentOKBodyRDSExporter{
					NodeID:     nodeID,
					AgentID:    agentID,
					PMMAgentID: pmmAgentID,
					CustomLabels: map[string]string{
						"custom_label_rds_exporter": "rds_exporter",
					},
					BasicMetricsDisabled:    true,
					EnhancedMetricsDisabled: true,
				},
			},
		}, getAgentRes)

		// Test change API.
		changeRDSExporterOK, err := client.Default.Agents.ChangeRDSExporter(&agents.ChangeRDSExporterParams{
			Body: agents.ChangeRDSExporterBody{
				AgentID: agentID,
				Common: &agents.ChangeRDSExporterParamsBodyCommon{
					Disable:            true,
					RemoveCustomLabels: true,
				},
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &agents.ChangeRDSExporterOK{
			Payload: &agents.ChangeRDSExporterOKBody{
				RDSExporter: &agents.ChangeRDSExporterOKBodyRDSExporter{
					NodeID:                  nodeID,
					AgentID:                 agentID,
					PMMAgentID:              pmmAgentID,
					Disabled:                true,
					BasicMetricsDisabled:    true,
					EnhancedMetricsDisabled: true,
				},
			},
		}, changeRDSExporterOK)

		changeRDSExporterOK, err = client.Default.Agents.ChangeRDSExporter(&agents.ChangeRDSExporterParams{
			Body: agents.ChangeRDSExporterBody{
				AgentID: agentID,
				Common: &agents.ChangeRDSExporterParamsBodyCommon{
					Enable: true,
					CustomLabels: map[string]string{
						"new_label": "rds_exporter",
					},
				},
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &agents.ChangeRDSExporterOK{
			Payload: &agents.ChangeRDSExporterOKBody{
				RDSExporter: &agents.ChangeRDSExporterOKBodyRDSExporter{
					NodeID:     nodeID,
					AgentID:    agentID,
					PMMAgentID: pmmAgentID,
					Disabled:   false,
					CustomLabels: map[string]string{
						"new_label": "rds_exporter",
					},
					BasicMetricsDisabled:    true,
					EnhancedMetricsDisabled: true,
				},
			},
		}, changeRDSExporterOK)
	})

	t.Run("AddNodeIDEmpty", func(t *testing.T) {
		t.Parallel()

		genericNodeID := pmmapitests.AddGenericNode(t, pmmapitests.TestString(t, "")).NodeID
		require.NotEmpty(t, genericNodeID)
		defer pmmapitests.RemoveNodes(t, genericNodeID)

		pmmAgent := pmmapitests.AddPMMAgent(t, genericNodeID)
		pmmAgentID := pmmAgent.PMMAgent.AgentID
		defer pmmapitests.RemoveAgents(t, pmmAgentID)

		res, err := client.Default.Agents.AddRDSExporter(&agents.AddRDSExporterParams{
			Body: agents.AddRDSExporterBody{
				NodeID:     "",
				PMMAgentID: pmmAgentID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field NodeId: value '' must not be an empty string")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveNodes(t, res.Payload.RDSExporter.AgentID)
		}
	})

	t.Run("NotExistNodeID", func(t *testing.T) {
		t.Parallel()

		genericNodeID := pmmapitests.AddGenericNode(t, pmmapitests.TestString(t, "")).NodeID
		require.NotEmpty(t, genericNodeID)
		defer pmmapitests.RemoveNodes(t, genericNodeID)

		pmmAgent := pmmapitests.AddPMMAgent(t, genericNodeID)
		pmmAgentID := pmmAgent.PMMAgent.AgentID
		defer pmmapitests.RemoveAgents(t, pmmAgentID)

		res, err := client.Default.Agents.AddRDSExporter(&agents.AddRDSExporterParams{
			Body: agents.AddRDSExporterBody{
				NodeID:     "pmm-node-id",
				PMMAgentID: pmmAgentID,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Node with ID \"pmm-node-id\" not found.")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveAgents(t, res.Payload.RDSExporter.AgentID)
		}
	})

	t.Run("NotExistPMMAgentID", func(t *testing.T) {
		t.Parallel()
		genericNodeID := pmmapitests.AddGenericNode(t, pmmapitests.TestString(t, "")).NodeID
		require.NotEmpty(t, genericNodeID)
		defer pmmapitests.RemoveNodes(t, genericNodeID)

		res, err := client.Default.Agents.AddRDSExporter(&agents.AddRDSExporterParams{
			Body: agents.AddRDSExporterBody{
				NodeID:     "nodeID",
				PMMAgentID: "pmm-not-exist-server",
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Agent with ID \"pmm-not-exist-server\" not found.")
		if !assert.Nil(t, res) {
			pmmapitests.RemoveAgents(t, res.Payload.RDSExporter.AgentID)
		}
	})

	t.Run("With PushMetrics", func(t *testing.T) {
		genericNodeID := pmmapitests.AddGenericNode(t, pmmapitests.TestString(t, "")).NodeID
		require.NotEmpty(t, genericNodeID)
		defer pmmapitests.RemoveNodes(t, genericNodeID)

		node := addRemoteRDSNode(t, pmmapitests.TestString(t, "Remote node for RDS exporter"))
		nodeID := node.RemoteRDS.NodeID
		defer pmmapitests.RemoveNodes(t, nodeID)

		pmmAgent := pmmapitests.AddPMMAgent(t, genericNodeID)
		pmmAgentID := pmmAgent.PMMAgent.AgentID
		defer pmmapitests.RemoveAgents(t, pmmAgentID)

		rdsExporter := addRDSExporter(t, agents.AddRDSExporterBody{
			NodeID:     nodeID,
			PMMAgentID: pmmAgentID,
			CustomLabels: map[string]string{
				"custom_label_rds_exporter": "rds_exporter",
			},
			SkipConnectionCheck:    true,
			DisableBasicMetrics:    true,
			DisableEnhancedMetrics: true,
		})
		agentID := rdsExporter.RDSExporter.AgentID
		defer pmmapitests.RemoveAgents(t, agentID)

		getAgentRes, err := client.Default.Agents.GetAgent(&agents.GetAgentParams{
			Body:    agents.GetAgentBody{AgentID: agentID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		assert.Equal(t, &agents.GetAgentOK{
			Payload: &agents.GetAgentOKBody{
				RDSExporter: &agents.GetAgentOKBodyRDSExporter{
					NodeID:     nodeID,
					AgentID:    agentID,
					PMMAgentID: pmmAgentID,
					CustomLabels: map[string]string{
						"custom_label_rds_exporter": "rds_exporter",
					},
					BasicMetricsDisabled:    true,
					EnhancedMetricsDisabled: true,
				},
			},
		}, getAgentRes)

		// Test change API.
		changeRDSExporterOK, err := client.Default.Agents.ChangeRDSExporter(&agents.ChangeRDSExporterParams{
			Body: agents.ChangeRDSExporterBody{
				AgentID: agentID,
				Common: &agents.ChangeRDSExporterParamsBodyCommon{
					EnablePushMetrics: true,
				},
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &agents.ChangeRDSExporterOK{
			Payload: &agents.ChangeRDSExporterOKBody{
				RDSExporter: &agents.ChangeRDSExporterOKBodyRDSExporter{
					NodeID:     nodeID,
					AgentID:    agentID,
					PMMAgentID: pmmAgentID,
					CustomLabels: map[string]string{
						"custom_label_rds_exporter": "rds_exporter",
					},
					BasicMetricsDisabled:    true,
					EnhancedMetricsDisabled: true,
					PushMetricsEnabled:      true,
				},
			},
		}, changeRDSExporterOK)

		changeRDSExporterOK, err = client.Default.Agents.ChangeRDSExporter(&agents.ChangeRDSExporterParams{
			Body: agents.ChangeRDSExporterBody{
				AgentID: agentID,
				Common: &agents.ChangeRDSExporterParamsBodyCommon{
					DisablePushMetrics: true,
				},
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &agents.ChangeRDSExporterOK{
			Payload: &agents.ChangeRDSExporterOKBody{
				RDSExporter: &agents.ChangeRDSExporterOKBodyRDSExporter{
					NodeID:     nodeID,
					AgentID:    agentID,
					PMMAgentID: pmmAgentID,
					CustomLabels: map[string]string{
						"custom_label_rds_exporter": "rds_exporter",
					},
					BasicMetricsDisabled:    true,
					EnhancedMetricsDisabled: true,
				},
			},
		}, changeRDSExporterOK)
		_, err = client.Default.Agents.ChangeRDSExporter(&agents.ChangeRDSExporterParams{
			Body: agents.ChangeRDSExporterBody{
				AgentID: agentID,
				Common: &agents.ChangeRDSExporterParamsBodyCommon{
					EnablePushMetrics:  true,
					DisablePushMetrics: true,
				},
			},
			Context: pmmapitests.Context,
		})

		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "expected one of  param: enable_push_metrics or disable_push_metrics")
	})
}

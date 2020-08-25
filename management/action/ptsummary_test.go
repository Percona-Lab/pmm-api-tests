package action

import (
	"context"
	"testing"
	"time"

	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/actions"
	"github.com/percona/pmm/api/managementpb/json/client/node"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func registerGenericNode(t pmmapitests.TestingT, body node.RegisterNodeBody) (string, string) {
	t.Helper()

	params := node.RegisterNodeParams{
		Context: pmmapitests.Context,
		Body:    body,
	}
	registerOK, err := client.Default.Node.RegisterNode(&params)
	require.NoError(t, err)
	require.NotNil(t, registerOK)
	require.NotNil(t, registerOK.Payload.PMMAgent)
	require.NotNil(t, registerOK.Payload.PMMAgent.AgentID)
	require.NotNil(t, registerOK.Payload.GenericNode)
	require.NotNil(t, registerOK.Payload.GenericNode.NodeID)
	return registerOK.Payload.GenericNode.NodeID, registerOK.Payload.PMMAgent.AgentID
}

func TestPTSummary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	explainActionOK, err := client.Default.Actions.StartPTSummaryAction(&actions.StartPTSummaryActionParams{
		Context: ctx,
		Body: actions.StartPTSummaryActionBody{
			NodeID: "pmm-server",
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, explainActionOK.Payload.ActionID)

	for {
		actionOK, err := client.Default.Actions.GetAction(&actions.GetActionParams{
			Context: ctx,
			Body: actions.GetActionBody{
				ActionID: explainActionOK.Payload.ActionID,
			},
		})
		require.NoError(t, err)

		if !actionOK.Payload.Done {
			time.Sleep(1 * time.Second)
			continue
		}

		require.True(t, actionOK.Payload.Done)
		require.Empty(t, actionOK.Payload.Error)
		require.NotEmpty(t, actionOK.Payload.Output)
		t.Log(actionOK.Payload.Output)

		break
	}
}

package action

import (
	"testing"
	"time"

	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/actions"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestPTSummary(t *testing.T) {
	t.Skip("not implemented yet")

	explainActionOK, err := client.Default.Actions.StartPTSummaryAction(&actions.StartPTSummaryActionParams{
		Context: pmmapitests.Context,
		Body: actions.StartPTSummaryActionBody{
			PMMAgentID: "/agent_id/12856e6b-5f4a-49cb-9c70-54edd0e0a074",
			NodeID:     "/node_id/5a9a7aa6-7af4-47be-817c-6d88e955bff2",
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, explainActionOK.Payload.ActionID)

	time.Sleep(2 * time.Second)

	actionOK, err := client.Default.Actions.GetAction(&actions.GetActionParams{
		Context: pmmapitests.Context,
		Body: actions.GetActionBody{
			ActionID: explainActionOK.Payload.ActionID,
		},
	})
	require.NoError(t, err)
	require.Empty(t, actionOK.Payload.Error)
	t.Log(actionOK.Payload.Output)
}

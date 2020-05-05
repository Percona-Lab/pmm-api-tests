package server

import (
	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
	"github.com/percona/pmm/api/alertmanager/amclient/alert"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/percona/pmm/api/alertmanager/amclient"
	"github.com/stretchr/testify/require"
)

func TestAlertmanager(t *testing.T) {

	t.Run("GetAlerts", func(t *testing.T) {
		res, err := amclient.Default.Alert.GetAlerts(&alert.GetAlertsParams{
			Filter:  []string{"node_name=pmm-server"},
			Context: pmmapitests.Context,
		})
		t.Logf("Response is %v", res)
		t.Logf("Error is %v", err)
		require.NoError(t, err)
		assert.Empty(t, res.Payload, "test")
	})
}

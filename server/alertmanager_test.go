package server

import (
	"github.com/percona/pmm/api/alertmanager/amclient/alert"
	"testing"

	"github.com/percona/pmm/api/alertmanager/amclient"
	"github.com/stretchr/testify/require"
)

func TestAlertmanager(t *testing.T) {

	t.Run("GetAlerts", func(t *testing.T) {
		params := alert.NewGetAlertsParams()
		params.Filter = []string{"node_name=pmm-server"}

		res, err := amclient.Default.Alert.GetAlerts(params)
		t.Logf("Params are %v", params)
		t.Logf("Response is %v", res)
		t.Logf("Error is %v", err)
		require.NoError(t, err)
	})
}

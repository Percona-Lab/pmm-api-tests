package server

import (
	"os"
	"testing"
	"time"

	"github.com/percona/pmm/api/alertmanager/amclient"
	"github.com/percona/pmm/api/alertmanager/amclient/alert"
	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAlertManager(t *testing.T) {
	res, err := serverClient.Default.Server.GetSettings(nil)
	require.NoError(t, err)
	assert.True(t, res.Payload.Settings.TelemetryEnabled)
	err = os.Setenv("PERCONA_TEST_ALERTMANAGER_RESEND_INTERVAL", "10s")
	require.NoError(t, err)

	t.Run("TestEndsAtForFailedChecksAlerts", func(t *testing.T) {
		if !pmmapitests.RunSTTTests {
			t.Skip("Skipping STT tests until we have environment: https://jira.percona.com/browse/PMM-5106")
		}

		defer restoreSettingsDefaults(t)

		// Enabling STT
		res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
			Body: server.ChangeSettingsBody{
				EnableStt: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.True(t, res.Payload.Settings.TelemetryEnabled)

		const resendInterval = 10

		// 120 sec ping for failed checks alerts to appear in alertmanager
		var alertsCount int
		for i := 0; i < 120; i++ {
			res, err := amclient.Default.Alert.GetAlerts(&alert.GetAlertsParams{
				Filter:  []string{"stt_check=1"},
				Context: pmmapitests.Context,
			})
			require.NoError(t, err)
			if len(res.Payload) == 0 {
				time.Sleep(1 * time.Second)
				continue
			}

			// TODO: Expand this test once we are silencing
			// removing alerts.
			for _, v := range res.Payload {
				delta := time.Duration(3 * resendInterval)
				// Since the `EndsAt` timestamp is always 3 times the
				// `resendInterval` in the future from `UpdatedAt`
				// we check whether they lie in that time delta.
				assert.WithinDuration(t, time.Time(*v.UpdatedAt), time.Time(*v.EndsAt), delta)
			}
			alertsCount = len(res.Payload)
			break
		}
		assert.Greater(t, alertsCount, 0, "No alerts met")
	})
}
